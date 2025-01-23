package main

import (
	"flag"
	"fmt"

	"palsp/internal/lsp"
)

func main() {
	port := flag.String("port", "", "Port to run the LSP server on (leave empty for stdio)")
	flag.Parse()

	if *port == "" {
		fmt.Println("Starting LSP server on stdio")
	} else {
		fmt.Println("Starting LSP server on port", *port)
	}

	lsp.StartServer(*port)
}
