package lsp

import (
	"encoding/json"
	"fmt"
	edit "palsp/internal/documents"
)

// Handle incoming JSON-RPC requests
func handleRequest(request LSPRequest) (response LSPResponse) {
	// Add defer recover to catch any panics and convert them to error responses
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic occurred while handling '%s': %v\n", request.Method, r)
			panicMsg := fmt.Sprintf("Internal error: %v", r)
			response = LSPResponse{
				JsonRPC: "2.0",
				ID:      request.ID,
				Error:   &LSPError{Code: -32603, Message: panicMsg},
			}
		}
	}()

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

	// Extract paths from workspace folders
	workspaceFolderPaths := make([]string, 0)
	if params.WorkspaceFolders != nil {
		for _, folder := range params.WorkspaceFolders {
			workspaceFolderPaths = append(workspaceFolderPaths, folder.URI)
		}
	}

	// Combine workspace folder paths with search folders from initialization options
	searchFolders := params.InitializationOptions.SearchFolders
	allFolders := append(workspaceFolderPaths, searchFolders...)

	opRes := edit.Mgr.Init(allFolders)
	return opResultToLSPResponse(id, opRes)
}

// Modified Handle textDocument/didOpen request
func handleDidOpen(params DidOpenTextDocumentParams, id int) LSPResponse {
	fmt.Println("File opened:", params.TextDocument.URI)
	opRes := edit.Mgr.DidOpen(params.TextDocument.URI, params.TextDocument.Text, params.TextDocument.Version)
	return opResultToLSPResponse(id, opRes)
}

// Modified Handle textDocument/didChange request
func handleDidChange(params DidChangeTextDocumentParams, id int) LSPResponse {
	fmt.Println("File changed:", params.TextDocument.URI)
	opRes := edit.Mgr.DidChange(params.TextDocument.URI, params.TextDocument.Text, params.TextDocument.Version)
	return opResultToLSPResponse(id, opRes)
}

// Modified Handle textDocument/didClose request
func handleDidClose(params DidCloseTextDocumentParams, id int) LSPResponse {
	fmt.Println("File closed:", params.TextDocument.URI)
	opRes := edit.Mgr.DidClose(params.TextDocument.URI)
	return opResultToLSPResponse(id, opRes)
}

// Modified Handle textDocument/completion request
func handleCompletion(params CompletionParams, id int) LSPResponse {
	fmt.Println("Completion requested for:", params.TextDocument.URI)
	opRes := edit.Mgr.Completion(params.TextDocument.URI, params.Position.Line, params.Position.Character)
	return opResultToLSPResponse(id, opRes)
}

// Modified Handle textDocument/hover request
func handleHover(params HoverParams, id int) LSPResponse {
	fmt.Println("Hover requested for:", params.TextDocument.URI)
	// Pass line and character from params.Position
	opRes := edit.Mgr.Hover(params.TextDocument.URI, params.TextDocument.Text, params.TextDocument.Version, params.Position.Line+1, params.Position.Character)
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
