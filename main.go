package main

import (
	"design-patterns/solid-in-go/lsp"
	"design-patterns/solid-in-go/ocp"
)

func main() {

	// OCP
	ocp.UseFilterBad()
	ocp.UseFilterOK()

	// LSP
	lsp.UseSizedBad()
}
