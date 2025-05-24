package lsp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	edit "palsp/internal/documents"
	"palsp/internal/log"
	"strings"
)

// Read and process incoming LSP messages
func processConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	processRequest(reader, writer)
}

// Read and process incoming LSP messages from stdio
func processStdio() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	processRequest(reader, writer)
}

func processRequest(reader *bufio.Reader, writer *bufio.Writer) {
	for {
		log.Main.Info().Msg("awaiting request")
		// Parse headers according to LSP spec
		contentLength := 0
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					log.Main.Info().Msg("Client disconnected")
				} else {
					log.Main.Error().Err(err).Msg("Error reading header")
				}
				return
			}

			// Trim trailing CR and LF
			line = strings.TrimRight(line, "\r\n")

			log.Main.Info().Str("header", line).Msg("Header received")

			// Empty line indicates end of headers
			if line == "" {
				break
			}

			// Parse Content-Length header
			if strings.HasPrefix(line, "Content-Length: ") {
				fmt.Sscanf(line, "Content-Length: %d", &contentLength)
			}
		}

		if contentLength == 0 {
			log.Main.Error().Msg("Invalid Content-Length")
			continue
		}

		// Read the content
		log.Main.Info().Msg("reading content")
		content := make([]byte, contentLength)
		_, err := io.ReadFull(reader, content)
		if err != nil {
			log.Main.Error().Err(err).Msg("Error reading content")
			break
		}

		// Unmarshal the request
		log.Main.Info().Msg("unmarshalling request")

		// Log the raw content (safely handling UTF-8)
		contentStr := string(content)
		if len(contentStr) > 1000 {
			log.Main.Info().Str("content", contentStr[:1000]+"...").Msg("REQUEST (truncated)")
		} else {
			log.Main.Info().Str("content", contentStr).Msg("REQUEST")
		}

		var request LSPRequest
		if err := json.Unmarshal(content, &request); err != nil {
			log.Main.Error().Err(err).Msg("Error unmarshalling request")
			continue
		}

		// Handle the request
		result := handleRequest(request)
		log.Main.Info().Msgf("RESULT: %v", result)

		if request.ID != nil {
			response := opResultToLSPResponse(*request.ID, result)
			responseBytes, err := json.Marshal(response)
			if err != nil {
				log.Main.Error().Err(err).Msg("Error marshalling response")
				continue
			}
			fmt.Fprintf(writer, "Content-Length: %d\r\n\r\n", len(responseBytes))
			writer.Write(responseBytes)
			writer.Flush()
		}

	}

}

// Start the LSP server based on the provided port
func StartServer(port string) {
	if port == "" {
		log.Main.Info().Msg("Starting LSP server on stdio")
		processStdio()
	} else {
		log.Main.Info().Msgf("Starting LSP server on port %s", port)
		listener, err := net.Listen("tcp", "localhost:"+port)
		if err != nil {
			log.Main.Error().Err(err).Msg("Error starting server")
			return
		}
		defer listener.Close()
		log.Main.Info().Str("port", port).Msg("LSP server started")

		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Main.Error().Err(err).Msg("Error accepting connection")
				continue
			}
			go processConnection(conn)
		}
	}
}

// Helper: convert edit.OpResult to LSPResponse.
func opResultToLSPResponse(id int, opResult edit.OpResult) LSPResponse {
	if !opResult.Success {
		var msg string
		if opResult.Error != nil {
			msg = fmt.Sprintf("%s (%v)", opResult.Message, opResult.Error)
		} else {
			msg = opResult.Message
		}
		return LSPResponse{
			JsonRPC: "2.0",
			ID:      id,
			Error:   &LSPError{Code: -32000, Message: msg},
		}
	}
	return LSPResponse{
		JsonRPC: "2.0",
		ID:      id,
		Result:  opResult.Result,
	}
}
