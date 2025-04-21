package lsp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"palsp/internal/log"
	"strings"
)

// Read and process incoming LSP messages
func processConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		// Parse headers according to LSP spec
		contentLength := 0
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					log.Logger.Info().Msg("Client disconnected")
				} else {
					log.Logger.Error().Err(err).Msg("Error reading header")
				}
				return
			}

			// Trim trailing CR and LF
			line = strings.TrimRight(line, "\r\n")

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
			log.Logger.Error().Msg("Invalid Content-Length")
			continue
		}

		// Read the content
		content := make([]byte, contentLength)
		_, err := io.ReadFull(reader, content)
		if err != nil {
			log.Logger.Error().Err(err).Msg("Error reading content")
			break
		}

		// Unmarshal the request
		var request LSPRequest
		if err := json.Unmarshal(content, &request); err != nil {
			log.Logger.Error().Err(err).Msg("Error unmarshalling request")
			continue
		}

		// Handle the request
		response := handleRequest(request)
		responseBytes, err := json.Marshal(response)
		if err != nil {
			log.Logger.Error().Err(err).Msg("Error marshalling response")
			continue
		}

		// Write the response
		fmt.Fprintf(conn, "Content-Length: %d\r\n\r\n", len(responseBytes))
		conn.Write(responseBytes)
	}
}

// Read and process incoming LSP messages from stdio
func processStdio() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for {
		log.Logger.Info().Msg("awaiting request")
		// Parse headers according to LSP spec
		contentLength := 0
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					log.Logger.Info().Msg("Client disconnected")
				} else {
					log.Logger.Error().Err(err).Msg("Error reading header")
				}
				return
			}

			// Trim trailing CR and LF
			line = strings.TrimRight(line, "\r\n")

			log.Logger.Info().Str("header", line).Msg("Header received")

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
			log.Logger.Error().Msg("Invalid Content-Length")
			continue
		}

		// Read the content
		log.Logger.Info().Msg("reading content")
		content := make([]byte, contentLength)
		_, err := io.ReadFull(reader, content)
		if err != nil {
			log.Logger.Error().Err(err).Msg("Error reading content")
			break
		}

		// Unmarshal the request
		log.Logger.Info().Msg("unmarshalling request")

		// Log the raw content (safely handling UTF-8)
		contentStr := string(content)
		if len(contentStr) > 1000 {
			log.Logger.Info().Str("content", contentStr[:1000]+"...").Msg("Request content (truncated)")
		} else {
			log.Logger.Info().Str("content", contentStr).Msg("Request content")
		}

		var request LSPRequest
		if err := json.Unmarshal(content, &request); err != nil {
			log.Logger.Error().Err(err).Msg("Error unmarshalling request")
			continue
		}

		// Handle the request
		log.Logger.Info().Msgf("HANDLING REQUEST WITH ID {%d}", request.ID)
		response := handleRequest(request)
		log.Logger.Info().Msgf("MARSHALLING RESPONSE WITH ID {%d}", response.ID)
		responseBytes, err := json.Marshal(response)
		if err != nil {
			log.Logger.Error().Err(err).Msg("Error marshalling response")
			continue
		}

		// Write the response
		log.Logger.Info().Msg("writing response")
		fmt.Fprintf(writer, "Content-Length: %d\r\n\r\n", len(responseBytes))
		writer.Write(responseBytes)
		writer.Flush()
		log.Logger.Info().Msg("writing response end and flushing")
	}
}

// Start the LSP server based on the provided port
func StartServer(port string) {
	if port == "" {
		log.Logger.Info().Msg("Starting LSP server on stdio")
		processStdio()
	} else {
		log.Logger.Info().Msgf("Starting LSP server on port %s", port)
		listener, err := net.Listen("tcp", "localhost:"+port)
		if err != nil {
			log.Logger.Error().Err(err).Msg("Error starting server")
			return
		}
		defer listener.Close()
		log.Logger.Info().Str("port", port).Msg("LSP server started")

		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Logger.Error().Err(err).Msg("Error accepting connection")
				continue
			}
			go processConnection(conn)
		}
	}
}
