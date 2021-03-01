package main

import (
	"fmt"
	"os"
	"replaymid/sqctool"
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

	//resp1,_ := sqctool.HTTPResp(string(resp.HTTP))
	//resp2,_ := sqctool.HTTPResp(string(msg.HTTP))
	//fmt.Fprintf(os.Stderr, "replay1 [%v] \n", param)
	//getDiffWithJson([]byte(resp1),[]byte(resp2))
	respStatus, _ := sqctool.HTTPStatus(string(resp.HTTP))
	replayStatus, _ := sqctool.HTTPStatus(string(msg.HTTP))
	if respStatus != replayStatus {
		fmt.Fprintf(os.Stderr, "replay status [%s] diffs from response status [%s]\n", replayStatus, respStatus)
	} else {
		fmt.Fprintf(os.Stderr, "same response status\n")
	}
	return msg
}

func main() {
	gor := sqctool.CreateGor()
	gor.On("request", OnRequest, "")
	gor.Run()
}



