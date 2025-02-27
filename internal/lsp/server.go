package lsp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
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
					fmt.Println("Client disconnected")
				} else {
					fmt.Println("Error reading header:", err.Error())
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
			fmt.Println("Invalid Content-Length")
			continue
		}

		// Read the content
		content := make([]byte, contentLength)
		_, err := io.ReadFull(reader, content)
		if err != nil {
			fmt.Println("Error reading content:", err.Error())
			break
		}

		// Unmarshal the request
		var request LSPRequest
		if err := json.Unmarshal(content, &request); err != nil {
			fmt.Println("Error unmarshalling request:", err.Error())
			continue
		}

		// Handle the request
		response := handleRequest(request)
		responseBytes, err := json.Marshal(response)
		if err != nil {
			fmt.Println("Error marshalling response:", err.Error())
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
		// Parse headers according to LSP spec
		contentLength := 0
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					fmt.Println("Client disconnected")
				} else {
					fmt.Println("Error reading header:", err.Error())
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
			fmt.Println("Invalid Content-Length")
			continue
		}

		// Read the content
		content := make([]byte, contentLength)
		_, err := io.ReadFull(reader, content)
		if err != nil {
			fmt.Println("Error reading content:", err.Error())
			break
		}

		// Unmarshal the request
		var request LSPRequest
		if err := json.Unmarshal(content, &request); err != nil {
			fmt.Println("Error unmarshalling request:", err.Error())
			continue
		}

		// Handle the request
		response := handleRequest(request)
		responseBytes, err := json.Marshal(response)
		if err != nil {
			fmt.Println("Error marshalling response:", err.Error())
			continue
		}

		// Write the response
		fmt.Fprintf(writer, "Content-Length: %d\r\n\r\n", len(responseBytes))
		writer.Write(responseBytes)
		writer.Flush()
	}
}

// Start the LSP server based on the provided port
func StartServer(port string) {
	if port == "" {
		fmt.Println("Starting LSP server on stdio")
		processStdio()
	} else {
		listener, err := net.Listen("tcp", "localhost:"+port)
		if err != nil {
			fmt.Println("Error starting server:", err.Error())
			return
		}
		defer listener.Close()
		fmt.Println("LSP server started on localhost:" + port)

		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error accepting connection:", err.Error())
				continue
			}
			go processConnection(conn)
		}
	}
}
