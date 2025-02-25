package edit

type OpResult struct {
	Success bool
	Message string
	Error   error
	Result  interface{}
}

func OpSuccess() OpResult {
	return OpResult{Success: true}
}

func OpSuccessWith(result interface{}) OpResult {
	return OpResult{Success: true, Result: result}
}

func OpFailure(msg string, err error) OpResult {
	return OpResult{Success: false, Message: msg, Error: err}
}

// LSP Hover response structures based on the LSP standard
type Hover struct {
	Contents MarkupContent `json:"contents"`
	Range    *Range        `json:"range,omitempty"`
}

type MarkupContent struct {
	Kind  string `json:"kind"`
	Value string `json:"value"`
}

type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

// LSP Completion response structures based on the LSP standard
type CompletionItem struct {
	Label         string `json:"label"`
	Kind          int    `json:"kind,omitempty"`
	Detail        string `json:"detail,omitempty"`
	Documentation string `json:"documentation,omitempty"`
}

type CompletionList struct {
	IsIncomplete bool             `json:"isIncomplete"`
	Items        []CompletionItem `json:"items"`
}
