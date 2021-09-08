package main

import (
	"fmt"
	"strings"
)

func genMessage(message uint32, process uint32) string {
	var b strings.Builder
	switch messageCode := (message & message_mask); messageCode {
	case request_message:
		b.WriteString("[Request Message] ")
	case grant_message:
		b.WriteString("[Grant Message] ")
	case release_message:
		b.WriteString("[Release Message] ")
	case error_message:
		b.WriteString("[Error Message] ")
	}
	fmt.Fprintf(&b, "Process: %v\n", process)
	return b.String()
}
