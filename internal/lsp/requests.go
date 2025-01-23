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
