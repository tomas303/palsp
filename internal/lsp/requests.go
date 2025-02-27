package lsp

import "encoding/json"

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

// InitializeParams structure
type InitializeParams struct {
	ProcessID    int                `json:"processId"`
	RootURI      string             `json:"rootUri"`
	Capabilities ClientCapabilities `json:"capabilities"`
	// Replace generic options with structured options
	InitializationOptions InitializationOptions `json:"initializationOptions,omitempty"`
	Trace                 string                `json:"trace,omitempty"`
	WorkspaceFolders      []WorkspaceFolder     `json:"workspaceFolders,omitempty"`
}

// Add new custom initialization options struct with SearchFolders as a subkey.
type InitializationOptions struct {
	SearchFolders []string `json:"searchFolders,omitempty"`
}

// WorkspaceFolder structure based on LSP standard
type WorkspaceFolder struct {
	URI  string `json:"uri"`
	Name string `json:"name"`
}

// ClientCapabilities structure
type ClientCapabilities struct {
	TextDocument TextDocumentClientCapabilities `json:"textDocument"`
}

// TextDocumentClientCapabilities structure
type TextDocumentClientCapabilities struct {
	Synchronization TextDocumentSyncClientCapabilities `json:"synchronization"`
}

// TextDocumentSyncClientCapabilities structure
type TextDocumentSyncClientCapabilities struct {
	DidSave bool `json:"didSave"`
}

// DidOpenTextDocumentParams structure
type DidOpenTextDocumentParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}

// TextDocumentItem structure
type TextDocumentItem struct {
	URI        string `json:"uri"`
	LanguageID string `json:"languageId"`
	Version    int    `json:"version"`
	Text       string `json:"text"`
}

// DidChangeTextDocumentParams structure
type DidChangeTextDocumentParams struct {
	TextDocument   TextDocumentItem                 `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

// TextDocumentContentChangeEvent structure
type TextDocumentContentChangeEvent struct {
	Text string `json:"text"`
}

// DidCloseTextDocumentParams structure
type DidCloseTextDocumentParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}

// CompletionParams structure
type CompletionParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
	Position     Position         `json:"position"`
}

// Position structure
type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

// CompletionItem structure
type CompletionItem struct {
	Label string `json:"label"`
	Kind  int    `json:"kind"`
}

// HoverParams structure
type HoverParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
	Position     Position         `json:"position"`
}

// Hover structure
type Hover struct {
	Contents []MarkedString `json:"contents"`
}

// MarkedString structure
type MarkedString struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}
