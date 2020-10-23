package sqctool

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

func invalidPayloadError(payload string) (string, error) {
	return "", fmt.Errorf("invalid payload: %s", payload)
}

// HTTPMethod returns http method from payload string
func HTTPMethod(payload string) (string, error) {
	pend := strings.IndexByte(payload, ' ')
	if pend == -1 {
		return invalidPayloadError(payload)
	}
	return payload[:pend], nil
}

// HTTPPath returns http path from payload string
func HTTPPath(payload string) (string, error) {
	payload, err := url.PathUnescape(payload)
	if err != nil {
		return "", err
	}
	pstart := strings.IndexByte(payload, ' ')
	if pstart == -1 {
		return invalidPayloadError(payload)
	}
	pend := strings.IndexByte(payload[pstart+1:], ' ')
	if pend == -1 {
		return invalidPayloadError(payload)
	}
	return payload[pstart+1 : pstart+pend+1], nil
}

// SetHTTPPath sets new path returns new payload string
func SetHTTPPath(payload, newPath string) (string, error) {
	pstart := strings.IndexByte(payload, ' ')
	if pstart == -1 {
		return invalidPayloadError(payload)
	}
	pend := strings.IndexByte(payload[pstart+1:], ' ')
	if pend == -1 {
		return invalidPayloadError(payload)
	}
	return payload[:pstart+1] + newPath + payload[pstart+pend+1:], nil
}

// HTTPStatus returns http status of the payload
// HTTP response have status code in the same position as path for requests
func HTTPStatus(payload string) (string, error) {
	return HTTPPath(payload)
}

// SetHTTPStatus set new status and return new payload string
func SetHTTPStatus(payload, newStatus string) (string, error) {
	return SetHTTPPath(payload, newStatus)
}

// HTTPPathParam gets path param with specific name, params with same name will
// be all returned in the slice
func HTTPPathParam(payload, name string) ([]string, error) {
	path, err := HTTPPath(payload)
	if err != nil {
		return []string{}, err
	}
	u, err := url.Parse(path)
	if err != nil {
		return []string{}, err
	}
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return []string{}, err
	}
	value, ok := m[name]
	if !ok {
		return []string{}, nil
	}
	return value, nil
}

// SetHTTPPathParam sets path param new value with key and returns new payload string
func SetHTTPPathParam(payload, name, value string) (string, error) {
	pathQs, err := HTTPPath(payload)
	if err != nil {
		return "", err
	}
	reg := regexp.MustCompile(name + "=([^&$]+)")
	newParam := name + "=" + value
	newPath := reg.ReplaceAllString(pathQs, newParam)
	if newPath == pathQs {
		if strings.IndexByte(newPath, '?') == -1 {
			newPath += "?"
		} else {
			newPath += "&"
		}
		newPath += name + "=" + value
	}
	return SetHTTPPath(payload, url.PathEscape(newPath))
}

// HTTPHeader returns the value of key with specific name in http headers
func HTTPHeader(payload, name string) (map[string]interface{}, error) {
	currentLine := 0
	idx := 0
	header := map[string]interface{}{
		"start":  -1,
		"end":    -1,
		"vstart": -1,
	}
	for idx < len(payload) {
		c := payload[idx]
		if c == '\n' {
			currentLine++
			idx++
			header["end"] = idx
			start := header["start"].(int)
			vstart := header["vstart"].(int)
			if currentLine > 0 && start > 0 && vstart > 0 {
				key := payload[start : vstart-1]
				if strings.Compare(strings.ToLower(key), strings.ToLower(name)) == 0 {
					header["name"] = strings.ToLower(name)
					end := header["end"].(int)
					header["value"] = strings.TrimSpace(payload[vstart:end])
					return header, nil
				}
			}
			header["start"] = -1
			header["vstart"] = -1
		} else if c == '\r' {
			idx++
			continue
		} else if c == ':' {
			if header["vstart"].(int) == -1 {
				idx++
				header["vstart"] = idx
				continue
			}
		}
		if header["start"].(int) == -1 {
			header["start"] = idx
		}
		idx++
	}
	return nil, nil
}

// SetHTTPHeader sets variable in http header with specific key value pair
func SetHTTPHeader(payload, name, value string) (string, error) {
	header, err := HTTPHeader(payload, name)
	if err != nil {
		return "", err
	}
	if header == nil {
		headerStart := strings.IndexByte(payload, '\n') + 1
		if headerStart == 0 {
			return invalidPayloadError(payload)
		}
		return payload[:headerStart] + name + ": " + value + "\r\n" + payload[headerStart:], nil
	}
	return payload[:header["vstart"].(int)] + " " + value + "\r\n" + payload[header["end"].(int):], nil
}
