# GoReplay Middleware



Sample code:

```go
package main

import (
	"fmt"
	"replaymid/sqctool"
	"os"
)

func OnRequest(gor *sqctool.Gor, msg *sqctool.GorMessage, kwargs ...interface{}) *sqctool.GorMessage {
	gor.On("response", OnResponse, msg.ID, msg)
	return msg
}

func OnResponse(gor *sqctool.Gor, msg *sqctool.GorMessage, kwargs ...interface{}) *sqctool.GorMessage {
	req, _ := kwargs[0].(*sqctool.GorMessage)
	gor.On("replay", OnReplay, req.ID, req, msg)
	return msg
}

func OnReplay(gor *sqctool.Gor, msg *sqctool.GorMessage, kwargs ...interface{}) *sqctool.GorMessage {
	req, _ := kwargs[0].(*sqctool.GorMessage)
	resp, _ := kwargs[1].(*sqctool.GorMessage)
	fmt.Fprintf(os.Stderr, "request raw http: %s\n", req.HTTP)
	fmt.Fprintf(os.Stderr, "response raw http: %s\n", resp.HTTP)
	fmt.Fprintf(os.Stderr, "replay raw http: %s\n", msg.HTTP)
	respStatus, _ := sqctool.HTTPStatus(string(resp.HTTP))
	replayStatus, _ := sqctool.HTTPStatus(string(msg.HTTP))
	if respStatus != replayStatus {
		fmt.Fprintf(os.Stderr, "replay status [%s] diffs from response status [%s]\n", replayStatus, respStatus)
	} else {
		fmt.Fprintf(os.Stderr, "replay status is same as response status\n")
	}
	return msg
}

func main() {
	gor := sqctool.CreateGor()
	gor.On("request", OnRequest, "")
	gor.Run()
}
```
