package main

import (

	"encoding/json"
	"fmt"
	"github.com/yudai/gojsondiff/formatter"
	"os"
	"replaymid/sqctool"
	"github.com/yudai/gojsondiff"
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
	//req, _ := kwargs[0].(*sqctool.GorMessage)
	resp, _ := kwargs[1].(*sqctool.GorMessage)
	//fmt.Fprintf(os.Stderr, "request raw http: %s\n", req.HTTP)
	//fmt.Fprintf(os.Stderr, "response raw http: %s\n", resp.HTTP)
	//fmt.Fprintf(os.Stderr, "replay raw http: %s\n", msg.HTTP)

	resp1,_ := sqctool.HTTPResp(string(resp.HTTP))
	resp2,_ := sqctool.HTTPResp(string(msg.HTTP))
	//fmt.Fprintf(os.Stderr, "replay1 [%v] \n", param)
	getDiffWithJson([]byte(resp1),[]byte(resp2))
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

func getDiffWithJson(aString,bString []byte){

	differ := gojsondiff.New()
	d, err := differ.Compare(aString, bString)
	if err != nil {
		fmt.Printf("Failed to unmarshal file: %s\n", err.Error())
		os.Exit(3)
	}
	var diffString string
	if d.Modified(){
		var aJson map[string]interface{}
		json.Unmarshal(aString, &aJson)
		config := formatter.AsciiFormatterConfig{
			ShowArrayIndex: true,
			Coloring:true,
		}
		formatters := formatter.NewAsciiFormatter(aJson, config)
		diffString, err = formatters.Format(d)
		if err != nil {
			// No error can occur
		}
	}
	if diffString != ""{
		fmt.Fprintf(os.Stderr, "差异化数据: \n%v\n",diffString)
	}else{
		fmt.Fprintf(os.Stderr, "same body | ")
	}

}

