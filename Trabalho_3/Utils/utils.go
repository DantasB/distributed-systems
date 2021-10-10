package utils

import (
	"fmt"
	"strings"
)

// GenMessage receives a message, a process and returns an error.
// It will build a message according the message code.
// It will also write the process number to this built string.
// It will returns the built string.
func GenMessage(message uint32, process uint32) string {
	var b strings.Builder
	switch messageCode := (message & Message_mask); messageCode {
	case Request_message:
		b.WriteString("[Request Message] ")
	case Grant_message:
		b.WriteString("[Grant Message] ")
	case Release_message:
		b.WriteString("[Release Message] ")
	case Error_message:
		b.WriteString("[Error Message] ")
	}
	fmt.Fprintf(&b, "Process: %v\n", process)
	return b.String()
}
