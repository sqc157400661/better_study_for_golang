package sqctool

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/buger/goreplay/proto"
	"os"
)

// GorMessage stores data and parsed information in a incoming request
type GorMessage struct {
	ID      string
	Type    string
	Meta    [][]byte // Meta is an array size of 4, containing: request type, uuid, timestamp, latency
	RawMeta []byte   //
	HTTP    []byte   // Raw HTTP payload
}


func process(buf []byte) {
	// First byte indicate payload type, possible values:
	//  1 - Request
	//  2 - Response
	//  3 - ReplayedResponse
	payloadType := buf[0]
	headerSize := bytes.IndexByte(buf, '\n') + 1
	header := buf[:headerSize-1]

	// Header contains space separated values of: request type, request id, and request start time (or round-trip time for responses)
	meta := bytes.Split(header, []byte(" "))
	// For each request you should receive 3 payloads (request, response, replayed response) with same request id
	reqID := string(meta[1])
	payload := buf[headerSize:]

	Debug("Received payload:", string(buf))
	fmt.Println("12121212121122")
	switch payloadType {
	case '1': // Request
		if bytes.Equal(proto.Path(payload), []byte("/token")) {
			originalTokens[reqID] = []byte{}
			Debug("Found token request:", reqID)
		} else {
			token, vs, _ := proto.PathParam(payload, []byte("token"))

			if vs != -1 { // If there is GET token param
				if alias, ok := tokenAliases[string(token)]; ok {
					// Rewrite original token to alias
					payload = proto.SetPathParam(payload, []byte("token"), alias)

					// Copy modified payload to our buffer
					buf = append(buf[:headerSize], payload...)
				}
			}
		}
		// Emitting data back
		os.Stdout.Write(encode(buf))
	case '2': // Original response
		if _, ok := originalTokens[reqID]; ok {
			// Token is inside response body
			secureToken := proto.Body(payload)
			originalTokens[reqID] = secureToken
			Debug("Remember origial token:", string(secureToken))
		}
	case '3': // Replayed response
		if originalToken, ok := originalTokens[reqID]; ok {
			delete(originalTokens, reqID)
			secureToken := proto.Body(payload)
			tokenAliases[string(originalToken)] = secureToken

			Debug("Create alias for new token token, was:", string(originalToken), "now:", string(secureToken))
		}
	}
}

// ParseMessage parses string middleware dataflow into a GorMessage
func  ParseMessage(payload []byte) (*GorMessage, error) {

	metaPos := bytes.Index(payload, []byte("\n"))
	metaRaw := payload[:metaPos]
	metaArr := bytes.Split(metaRaw, []byte(" "))
	ptype, pid := metaArr[0], string(metaArr[1])
	httpPayload := payload[metaPos+1:]
	return &GorMessage{
		ID:      pid,
		Type:    string(ptype),
		Meta:    metaArr,
		RawMeta: metaRaw,
		HTTP:    httpPayload,
	}, nil
}

func encode(buf []byte) []byte {
	dst := make([]byte, len(buf)*2+1)
	hex.Encode(dst, buf)
	dst[len(dst)-1] = '\n'
	return dst
}

func Debug(args ...interface{}) {
	if os.Getenv("GOR_TEST") == "" { // if we are not testing
		fmt.Fprint(os.Stderr, "[DEBUG][TOKEN-MOD] ")
		fmt.Fprintln(os.Stderr, args...)
	}
}