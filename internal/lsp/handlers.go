package lsp

import (
	"encoding/json"
	"fmt"
	"palsp/internal/discover"
)

// Handle incoming JSON-RPC requests
func handleRequest(request LSPRequest) LSPResponse {
	fmt.Println("Received request:", request.Method)

	switch request.Method {
	case "initialize":
		var params map[string]interface{}
		if err := json.Unmarshal(request.Params, &params); err != nil {
			return LSPResponse{
				JsonRPC: "2.0",
				ID:      request.ID,
				Error:   &LSPError{Code: -32602, Message: "Invalid params"},
			}
		}
		return handleInitialize(request.ID)

	case "textDocument/didOpen":
		var params DidOpenTextDocumentParams
		if err := json.Unmarshal(request.Params, &params); err != nil {
			return LSPResponse{
				JsonRPC: "2.0",
				ID:      request.ID,
				Error:   &LSPError{Code: -32602, Message: "Invalid params"},
			}
		}
		return handleDidOpen(params, request.ID)

	default:
		return LSPResponse{
			JsonRPC: "2.0",
			ID:      request.ID,
			Error:   &LSPError{Code: -32601, Message: "Method not found"},
		}
	}
}

// Handle initialize request
func handleInitialize(id int) LSPResponse {
	fmt.Println("Initialize request received")
	return LSPResponse{
		JsonRPC: "2.0",
		ID:      id,
		Result: InitializeResult{
			Capabilities: map[string]interface{}{
				"textDocumentSync": 1, // Full document sync
			},
		},
	}
}

// Handle textDocument/didOpen request
func handleDidOpen(params DidOpenTextDocumentParams, id int) LSPResponse {
	fmt.Println("File opened:", params.TextDocument.URI)
	discover.HandleDidOpen(params.TextDocument.URI, params.TextDocument.Text)
	return LSPResponse{JsonRPC: "2.0", ID: id, Result: nil}
}
