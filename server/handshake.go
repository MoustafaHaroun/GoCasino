package main

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
)

// generateAcceptHeader generates the value for the Sec-WebSocket-Accept header.
func generateAcceptHeader(key string) string {
	const magicString = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
	hash := sha1.New()
	hash.Write([]byte(key + magicString))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

// handshake performs the WebSocket handshake as specified in RFC 6455.
// It validates the incoming HTTP request and, if valid, returns the handshake response.
// If the request is invalid, it returns an error.
//
// References:
// - RFC 6455: https://tools.ietf.org/html/rfc6455
func handshake(r *http.Request) (string, error) {
	if r.Method != http.MethodGet {
		return "", errors.New("invalid method")
	}

	secWebSocketKey := r.Header.Get("Sec-WebSocket-Key")
	if secWebSocketKey == "" {
		return "", errors.New("no Sec-WebSocket-Key header")
	}

	acceptHeader := generateAcceptHeader(secWebSocketKey)
	response := fmt.Sprintf(
		"HTTP/1.1 101 Switching Protocols\r\n"+
			"Upgrade: websocket\r\n"+
			"Connection: Upgrade\r\n"+
			"Sec-WebSocket-Accept: %v\r\n\r\n", acceptHeader)

	return response, nil
}
