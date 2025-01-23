package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strings"
)

// LSP Request structure
type LSPRequest struct {
	JsonRPC string          `json:"jsonrpc"`
	ID      int             `json:"id,omitempty"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

// LSP Response structure
type LSPResponse struct {
	JsonRPC string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   *LSPError   `json:"error,omitempty"`
}

// LSP Error structure
type LSPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Initialize Response
type InitializeResult struct {
	Capabilities map[string]interface{} `json:"capabilities"`
}

// Handle incoming JSON-RPC requests
func handleRequest(request LSPRequest) LSPResponse {
	fmt.Println("Received request:", request.Method)

	switch request.Method {
	case "initialize":
		return LSPResponse{
			JsonRPC: "2.0",
			ID:      request.ID,
			Result: InitializeResult{
				Capabilities: map[string]interface{}{
					"textDocumentSync": 1, // Full document sync
				},
			},
		}

	case "textDocument/didOpen":
		fmt.Println("File opened:", string(request.Params))
		return LSPResponse{JsonRPC: "2.0", ID: request.ID, Result: nil}

	default:
		return LSPResponse{
			JsonRPC: "2.0",
			ID:      request.ID,
			Error:   &LSPError{Code: -32601, Message: "Method not found"},
		}
	}
}

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
				fmt.Println("Read error:", err)
			}
			return
		}

		// Extract content length
		if !strings.HasPrefix(header, "Content-Length:") {
			continue
		}
		var contentLength int
		fmt.Sscanf(header, "Content-Length: %d", &contentLength)

		// Read empty line
		reader.ReadString('\n')

		// Read JSON payload
		body := make([]byte, contentLength)
		_, err = io.ReadFull(reader, body)
		if err != nil {
			fmt.Println("Failed to read body:", err)
			return
		}

		// Parse JSON-RPC request
		var request LSPRequest
		err = json.Unmarshal(body, &request)
		if err != nil {
			fmt.Println("JSON parse error:", err)
			return
		}

		// Handle request
		response := handleRequest(request)

		// Encode and send response
		responseJSON, _ := json.Marshal(response)
		responseStr := fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(responseJSON), responseJSON)
		conn.Write([]byte(responseStr))
	}
}

// Start LSP server on TCP
func startServer(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
	fmt.Println("LSP Server listening on", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go processConnection(conn)
	}
}

func main() {
	startServer("127.0.0.1:8080")
}
