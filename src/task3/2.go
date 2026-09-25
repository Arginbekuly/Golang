package main

import (
	"bytes"
	"strings"
)

func processPayloadOriginal(data []byte) string {
	trimmed := string(bytes.TrimSpace(data))
	clean := strings.ReplaceAll(trimmed, "\r", "")
	return clean
}

func processPayload(data []byte) string {
	trimmed := bytes.TrimSpace(data)
	clean := bytes.ReplaceAll(trimmed, []byte("\r"), []byte(""))
	return string(clean)
}
