package lsp

import (
	"encoding/json"
	"fmt"
	"palsp/internal/edit"
)

// Handle incoming JSON-RPC requests
func handleRequest(request LSPRequest) LSPResponse {
	fmt.Println("Received request:", request.Method)

	switch request.Method {
	case "initialize":
		var params InitializeParams
		if err := json.Unmarshal(request.Params, &params); err != nil {
			return LSPResponse{
				JsonRPC: "2.0",
				ID:      request.ID,
				Error:   &LSPError{Code: -32602, Message: "Invalid params"},
			}
		}
		return handleInitialize(request.ID, params)

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

	case "textDocument/didChange":
		var params DidChangeTextDocumentParams
		if err := json.Unmarshal(request.Params, &params); err != nil {
			return LSPResponse{
				JsonRPC: "2.0",
				ID:      request.ID,
				Error:   &LSPError{Code: -32602, Message: "Invalid params"},
			}
		}
		return handleDidChange(params, request.ID)

	case "textDocument/didClose":
		var params DidCloseTextDocumentParams
		if err := json.Unmarshal(request.Params, &params); err != nil {
			return LSPResponse{
				JsonRPC: "2.0",
				ID:      request.ID,
				Error:   &LSPError{Code: -32602, Message: "Invalid params"},
			}
		}
		return handleDidClose(params, request.ID)

	case "textDocument/completion":
		var params CompletionParams
		if err := json.Unmarshal(request.Params, &params); err != nil {
			return LSPResponse{
				JsonRPC: "2.0",
				ID:      request.ID,
				Error:   &LSPError{Code: -32602, Message: "Invalid params"},
			}
		}
		return handleCompletion(params, request.ID)

	case "textDocument/hover":
		var params HoverParams
		if err := json.Unmarshal(request.Params, &params); err != nil {
			return LSPResponse{
				JsonRPC: "2.0",
				ID:      request.ID,
				Error:   &LSPError{Code: -32602, Message: "Invalid params"},
			}
		}
		return handleHover(params, request.ID)

	default:
		return LSPResponse{
			JsonRPC: "2.0",
			ID:      request.ID,
			Error:   &LSPError{Code: -32601, Message: "Method not found"},
		}
	}
}

// Handle initialize request
func handleInitialize(id int, params InitializeParams) LSPResponse {
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

// Modified Handle textDocument/didOpen request
func handleDidOpen(params DidOpenTextDocumentParams, id int) LSPResponse {
	fmt.Println("File opened:", params.TextDocument.URI)
	opRes := edit.Lspi.DidOpen(params.TextDocument.URI, params.TextDocument.Text)
	return opResultToLSPResponse(id, opRes)
}

// Modified Handle textDocument/didChange request
func handleDidChange(params DidChangeTextDocumentParams, id int) LSPResponse {
	fmt.Println("File changed:", params.TextDocument.URI)
	opRes := edit.Lspi.DidChange(params.TextDocument.URI, params.TextDocument.Text)
	return opResultToLSPResponse(id, opRes)
}

// Modified Handle textDocument/didClose request
func handleDidClose(params DidCloseTextDocumentParams, id int) LSPResponse {
	fmt.Println("File closed:", params.TextDocument.URI)
	opRes := edit.Lspi.DidClose(params.TextDocument.URI)
	return opResultToLSPResponse(id, opRes)
}

// Modified Handle textDocument/completion request
func handleCompletion(params CompletionParams, id int) LSPResponse {
	fmt.Println("Completion requested for:", params.TextDocument.URI)
	opRes := edit.Lspi.Completion(params.TextDocument.URI, params.Position.Line, params.Position.Character)
	return opResultToLSPResponse(id, opRes)
}

// Modified Handle textDocument/hover request
func handleHover(params HoverParams, id int) LSPResponse {
	fmt.Println("Hover requested for:", params.TextDocument.URI)
	// Pass line and character from params.Position
	opRes := edit.Lspi.Hover(params.TextDocument.URI, params.Position.Line, params.Position.Character)
	return opResultToLSPResponse(id, opRes)
}

// Helper: convert edit.OpResult to LSPResponse.
func opResultToLSPResponse(id int, opResult edit.OpResult) LSPResponse {
	if !opResult.Success {
		return LSPResponse{
			JsonRPC: "2.0",
			ID:      id,
			Error:   &LSPError{Code: -32000, Message: opResult.Message},
		}
	}
	return LSPResponse{
		JsonRPC: "2.0",
		ID:      id,
		Result:  opResult.Result,
	}
}
