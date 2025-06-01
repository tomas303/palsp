package lsp

import "encoding/json"

// LSP Request structure
type LSPRequest struct {
	JsonRPC string          `json:"jsonrpc"`
	ID      *int            `json:"id,omitempty"`
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
	SearchFolders  []string `json:"searchFolders,omitempty"`
	UnitScopeNames []string `json:"unitScopeNames,omitempty"`
	PrefetchUnits  bool     `json:"prefetchUnits,omitempty"`
	Defines        []string `json:"defines,omitempty"` // Add compiler defines support
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

// DidChangeTextDocumentParams structure - corrected
type DidChangeTextDocumentParams struct {
	TextDocument   VersionedTextDocumentIdentifier  `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

// TextDocumentContentChangeEvent structure
type TextDocumentContentChangeEvent struct {
	Range       *TextRange `json:"range,omitempty"`       // Optional: specific range that changed
	RangeLength *int       `json:"rangeLength,omitempty"` // Optional: length of range being replaced
	Text        string     `json:"text"`                  // The new text content
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

// DefinitionParams structure
type DefinitionParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
	Position     Position         `json:"position"`
}

// Location structure to return definition locations
type Location struct {
	URI   string    `json:"uri"`
	Range TextRange `json:"range"`
}

// TextRange structure for text ranges
type TextRange struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

// DumpScopesParams structure for the dump definition command
type DumpScopesParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}

// DumpDBScopesParams structure for the dump database scopes command
type DumpDBScopesParams struct {
	TextDocument struct {
		URI string `json:"uri"`
	} `json:"textDocument"`
}

// ExecuteSQLQueryParams structure for the execute SQL query command
type ExecuteSQLQueryParams struct {
	SQLQuery string `json:"sqlQuery"`
}

// Add VersionedTextDocumentIdentifier structure
type VersionedTextDocumentIdentifier struct {
	URI     string `json:"uri"`
	Version int    `json:"version"`
}
