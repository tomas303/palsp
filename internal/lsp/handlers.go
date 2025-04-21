package lsp

import (
	"encoding/json"
	"fmt"
	edit "palsp/internal/documents"
	"palsp/internal/log"
)

// Handle incoming JSON-RPC requests
func handleRequest(request LSPRequest) (response edit.OpResult) {
	// Add defer recover to catch any panics and convert them to error responses
	defer func() {
		if r := recover(); r != nil {
			var err error
			switch v := r.(type) {
			case error:
				err = v
			default:
				// It's some other type
				err = fmt.Errorf("%v", v)
			}
			response = edit.OpFailure("", err)
			log.Logger.Error().Err(err).Msg("Recovered from panic in LSP handler")
		}
	}()

	switch request.Method {
	case "initialize":
		var params InitializeParams
		if err := json.Unmarshal(request.Params, &params); err != nil {
			return edit.OpFailure("Invalid params", err)
		}
		return handleInitialize(params)

	case "textDocument/didOpen":
		var params DidOpenTextDocumentParams
		if err := json.Unmarshal(request.Params, &params); err != nil {
			return edit.OpFailure("Invalid params", err)
		}
		return handleDidOpen(params)

	case "textDocument/didChange":
		var params DidChangeTextDocumentParams
		if err := json.Unmarshal(request.Params, &params); err != nil {
			return edit.OpFailure("Invalid params", err)
		}
		return handleDidChange(params)

	case "textDocument/didClose":
		var params DidCloseTextDocumentParams
		if err := json.Unmarshal(request.Params, &params); err != nil {
			return edit.OpFailure("Invalid params", err)
		}
		return handleDidClose(params)

	case "textDocument/completion":
		var params CompletionParams
		if err := json.Unmarshal(request.Params, &params); err != nil {
			return edit.OpFailure("Invalid params", err)
		}
		return handleCompletion(params)

	case "textDocument/hover":
		var params HoverParams
		if err := json.Unmarshal(request.Params, &params); err != nil {
			return edit.OpFailure("Invalid params", err)
		}
		return handleHover(params)

	default:
		return edit.OpFailure("Method not found", fmt.Errorf("Method %s not found", request.Method))
	}
}

// Handle initialize request
func handleInitialize(params InitializeParams) edit.OpResult {
	log.Logger.Debug().Msg("Initialize request received")

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

	return edit.Mgr.Init(allFolders)
}

// Modified Handle textDocument/didOpen request
func handleDidOpen(params DidOpenTextDocumentParams) edit.OpResult {
	log.Logger.Debug().Str("file", params.TextDocument.URI).Msg("File opened")
	return edit.Mgr.DidOpen(params.TextDocument.URI, params.TextDocument.Text, params.TextDocument.Version)
}

// Modified Handle textDocument/didChange request
func handleDidChange(params DidChangeTextDocumentParams) edit.OpResult {
	log.Logger.Debug().Str("file", params.TextDocument.URI).Msg("File changed")
	return edit.Mgr.DidChange(params.TextDocument.URI, params.TextDocument.Text, params.TextDocument.Version)
}

// Modified Handle textDocument/didClose request
func handleDidClose(params DidCloseTextDocumentParams) edit.OpResult {
	log.Logger.Debug().Str("file", params.TextDocument.URI).Msg("File closed")
	return edit.Mgr.DidClose(params.TextDocument.URI)
}

// Modified Handle textDocument/completion request
func handleCompletion(params CompletionParams) edit.OpResult {
	log.Logger.Debug().Str("file", params.TextDocument.URI).Msg("Completion requested")
	return edit.Mgr.Completion(params.TextDocument.URI, params.TextDocument.Text, params.TextDocument.Version, params.Position.Line+1, params.Position.Character)
}

// Modified Handle textDocument/hover request
func handleHover(params HoverParams) edit.OpResult {
	log.Logger.Debug().Str("file", params.TextDocument.URI).Msg("Hover requested")
	// Pass line and character from params.Position
	return edit.Mgr.Hover(params.TextDocument.URI, params.TextDocument.Text, params.TextDocument.Version, params.Position.Line+1, params.Position.Character)
}
