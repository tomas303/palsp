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
		// Read LSP headers
		header, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected")
			} else {
				fmt.Println("Error reading header:", err.Error())
			}
			break
		}

		// Read LSP content length
		if !strings.HasPrefix(header, "Content-Length: ") {
			fmt.Println("Invalid header:", header)
			continue
		}
		var contentLength int
		fmt.Sscanf(header, "Content-Length: %d", &contentLength)

		// Read the blank line
		_, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading blank line:", err.Error())
			break
		}

		// Read the content
		content := make([]byte, contentLength)
		_, err = io.ReadFull(reader, content)
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
		// Read LSP headers
		header, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected")
			} else {
				fmt.Println("Error reading header:", err.Error())
			}
			break
		}

		// Read LSP content length
		if !strings.HasPrefix(header, "Content-Length: ") {
			fmt.Println("Invalid header:", header)
			continue
		}
		var contentLength int
		fmt.Sscanf(header, "Content-Length: %d", &contentLength)

		// Read the blank line
		_, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading blank line:", err.Error())
			break
		}

		// Read the content
		content := make([]byte, contentLength)
		_, err = io.ReadFull(reader, content)
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
