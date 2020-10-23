package sqctool

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHTTPMethod(t *testing.T) {
	payload := "GET /test HTTP/1.1\r\n\r\n"
	if method, err := HTTPMethod(payload); err != nil {
		t.Errorf(err.Error())
	} else if method != "GET" {
		t.Errorf("invalid method: %s", method)
	}
}

func TestHTTPPath(t *testing.T) {
	payload := "GET /test HTTP/1.1\r\n\r\n"
	if path, err := HTTPPath(payload); err != nil {
		t.Errorf(err.Error())
	} else if path != "/test" {
		t.Errorf("invalid path: %s", path)
	}

	if newPayload, err := SetHTTPPath(payload, "/"); err != nil {
		t.Errorf(err.Error())
	} else if newPayload != "GET / HTTP/1.1\r\n\r\n" {
		t.Errorf("invalid newPayload: %s", newPayload)
	}

	if newPayload, err := SetHTTPPath(payload, "/new/test"); err != nil {
		t.Errorf(err.Error())
	} else if newPayload != "GET /new/test HTTP/1.1\r\n\r\n" {
		t.Errorf("invalid newPayload: %s", newPayload)
	}
}

func TestHTTPPathParam(t *testing.T) {
	payload := "GET /test HTTP/1.1\r\n\r\n"
	if params, err := HTTPPathParam(payload, "test"); err != nil {
		t.Errorf(err.Error())
	} else if !reflect.DeepEqual(params, []string{}) {
		t.Errorf("invalid params: %v", params)
	}

	payload, err := SetHTTPPathParam(payload, "test", "123")
	if err != nil {
		t.Errorf(err.Error())
	}
	if params, err2 := HTTPPathParam(payload, "test"); err != nil {
		t.Errorf(err2.Error())
	} else if !reflect.DeepEqual(params, []string{"123"}) {
		t.Errorf("invalid params: %v", params)
	}

	payload, err = SetHTTPPathParam(payload, "qwer", "ty")
	if err != nil {
		t.Errorf(err.Error())
	}
	if params, err := HTTPPathParam(payload, "qwer"); err != nil {
		t.Errorf(err.Error())
	} else if !reflect.DeepEqual(params, []string{"ty"}) {
		t.Errorf("invalid params: %v", params)
	}
}

func TestHTTPHeader(t *testing.T) {
	payload := "GET / HTTP/1.1\r\nHost: localhost:3000\r\nUser-Agent: Golang\r\nContent-Length:5\r\n\r\nhello"
	expected := map[string]string{
		"host":           "localhost:3000",
		"User-Agent":     "Golang",
		"Content-Length": "5",
	}
	for name, value := range expected {
		header, err := HTTPHeader(payload, name)
		if err != nil {
			t.Errorf(err.Error())
		}
		if header == nil || header["value"] != value {
			t.Errorf("invalid header %s value %s", name, header["value"])
		}
	}
}

func TestSetHTTPHeader(t *testing.T) {
	payload := "GET / HTTP/1.1\r\nUser-Agent: Golang\r\nContent-Length: 5\r\n\r\nhello"
	uas := []string{"", "1", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_0)"}
	expected := "GET / HTTP/1.1\r\nUser-Agent: %s\r\nContent-Length: 5\r\n\r\nhello"
	for _, ua := range uas {
		newPayload, err := SetHTTPHeader(payload, "user-agent", ua)
		if err != nil {
			t.Errorf(err.Error())
		}
		if newPayload != fmt.Sprintf(expected, ua) {
			t.Errorf("invalid payload after set http header: %s", newPayload)
		}
	}

	expected = "GET / HTTP/1.1\r\nX-Test: test\r\nUser-Agent: Golang\r\nContent-Length: 5\r\n\r\nhello"
	newPayload, err := SetHTTPHeader(payload, "X-Test", "test")
	if err != nil {
		t.Errorf(err.Error())
	} else if newPayload != expected {
		t.Errorf("invalid payload after set http header: %s", newPayload)
	}

	expected = "GET / HTTP/1.1\r\nX-Test2: test2\r\nX-Test: test\r\nUser-Agent: Golang\r\nContent-Length: 5\r\n\r\nhello"
	newPayload, err = SetHTTPHeader(newPayload, "X-Test2", "test2")
	if err != nil {
		t.Errorf(err.Error())
	} else if newPayload != expected {
		t.Errorf("invalid payload after set http header: %s", newPayload)
	}
}
