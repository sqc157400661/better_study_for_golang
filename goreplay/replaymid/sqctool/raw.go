/*
This middleware made for auth system that randomly generate access tokens, which used later for accessing secure content. Since there is no pre-defined token value, naive approach without middleware (or if middleware use only request payloads) will fail, because replayed server have own tokens, not synced with origin. To fix this, our middleware should take in account responses of replayed and origin server, store `originalToken -> replayedToken` aliases and rewrite all requests using this token to use replayed alias. See `middleware_test.go#TestTokenMiddleware` test for examples of using this middleware.

How middleware works:

                   Original request      +--------------+
+-------------+----------STDIN---------->+              |
|  Gor input  |                          |  Middleware  |
+-------------+----------STDIN---------->+              |
                   Original response     +------+---+---+
                                                |   ^
+-------------+    Modified request             v   |
| Gor output  +<---------STDOUT-----------------+   |
+-----+-------+                                     |
      |                                             |
      |            Replayed response                |
      +------------------STDIN----------------->----+
*/

package sqctool

import (
	"bufio"
	"encoding/hex"
	"os"
)

// requestID -> originalToken
var originalTokens map[string][]byte

// originalToken -> replayedToken
var tokenAliases map[string][]byte

func Raw() {
	originalTokens = make(map[string][]byte)
	tokenAliases = make(map[string][]byte)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		encoded := scanner.Bytes()
		buf := make([]byte, len(encoded)/2)
		hex.Decode(buf, encoded)
		process(buf)
	}
}

