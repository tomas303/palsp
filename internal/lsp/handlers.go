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
			log.Main.Error().Err(err).Msg("Recovered from panic in LSP handler")
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

	case "textDocument/definition":
		var params DefinitionParams
		if err := json.Unmarshal(request.Params, &params); err != nil {
			return edit.OpFailure("Invalid params", err)
		}
		return handleDefinition(params)

	case "pascal/dumpScopes":
		var params DumpScopesParams
		if err := json.Unmarshal(request.Params, &params); err != nil {
			return edit.OpFailure("Invalid params", err)
		}
		return handleDumpScopes(params)

	case "pascal/dumpDBScopes":
		var params DumpDBScopesParams
		if err := json.Unmarshal(request.Params, &params); err != nil {
			return edit.OpFailure("Invalid params", err)
		}
		return handleDumpDBScopes(params)

	case "pascal/executeSQLQuery":
		var params ExecuteSQLQueryParams
		if err := json.Unmarshal(request.Params, &params); err != nil {
			return edit.OpFailure("Invalid params", err)
		}
		return handleExecuteSQLQuery(params)

	default:
		return edit.OpFailure("Method not found", fmt.Errorf("Method %s not found", request.Method))
	}
}

// Handle initialize request
func handleInitialize(params InitializeParams) edit.OpResult {

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

	return edit.GetManager().Init(allFolders, params.InitializationOptions.UnitScopeNames, params.InitializationOptions.PrefetchUnits, params.InitializationOptions.Defines)
}

// Modified Handle textDocument/didOpen request
func handleDidOpen(params DidOpenTextDocumentParams) edit.OpResult {
	return edit.GetManager().DidOpen(params.TextDocument.URI, params.TextDocument.Text, params.TextDocument.Version)
}

// Modified Handle textDocument/didChange request
func handleDidChange(params DidChangeTextDocumentParams) edit.OpResult {
	// For full document sync (most common case)
	if len(params.ContentChanges) > 0 {
		// Take the last change which should contain full document for full sync
		lastChange := params.ContentChanges[len(params.ContentChanges)-1]

		// If no range specified, it's a full document update
		if lastChange.Range == nil {
			return edit.GetManager().DidChange(
				params.TextDocument.URI,
				lastChange.Text,
				params.TextDocument.Version,
			)
		}

		// TODO: Handle incremental changes if needed
		// For now, fallback to treating as full document
		return edit.GetManager().DidChange(
			params.TextDocument.URI,
			lastChange.Text,
			params.TextDocument.Version,
		)
	}

	return edit.OpFailure("No content changes provided", nil)
}

// Modified Handle textDocument/didClose request
func handleDidClose(params DidCloseTextDocumentParams) edit.OpResult {
	return edit.GetManager().DidClose(params.TextDocument.URI)
}

// Modified Handle textDocument/completion request
func handleCompletion(params CompletionParams) edit.OpResult {
	return edit.GetManager().Completion(params.TextDocument.URI, params.Position.Line, params.Position.Character)
}

// Modified Handle textDocument/hover request
func handleHover(params HoverParams) edit.OpResult {
	return edit.GetManager().Hover(params.TextDocument.URI, params.Position.Line, params.Position.Character)
}

// Modified Handle textDocument/definition request
func handleDefinition(params DefinitionParams) edit.OpResult {
	return edit.GetManager().Definition(params.TextDocument.URI, params.Position.Line, params.Position.Character)
}

// Handle dump definition request
func handleDumpScopes(params DumpScopesParams) edit.OpResult {
	return edit.GetManager().DumpScopes(params.TextDocument.URI)
}

// Handle dump database scopes request
func handleDumpDBScopes(params DumpDBScopesParams) edit.OpResult {
	return edit.GetManager().DumpDBScopes(params.TextDocument.URI)
}

// Handle execute SQL query request
func handleExecuteSQLQuery(params ExecuteSQLQueryParams) edit.OpResult {
	return edit.GetManager().ExecuteSQLQuery(params.SQLQuery)
}
