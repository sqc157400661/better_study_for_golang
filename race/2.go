package main

import (
	"fmt"
	"strings"
)

func main() {
	aa := `POST /verify2 HTTP/1.1
User-Agent: curl/7.22.0 (x86_64-unknown-linux-gnu) libcurl/7.22.0 OpenSSL/1.0.2k zlib/1.2.7
Host: 10.0.20.22:5052
Accept: */*
Content-Length: 457
Content-Type: application/x-www-form-urlencoded

reqtype=order_query_check&userid=133223647&cookie_user=MDpoel9jb2JyYTo6Tm9uZTo1MDA6MTQzMjIzNjQ3OjcsMTExMTExMTExMTEsNDA7NDQsMTEsNDA7NiwxLDQwOzUsMSw0MDsxLDEwMSw0MDsyLDEsNDA7MywxLDQwOzUsMSw0MDs4LDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAxLDQwOzEwMiwxLDQwOjo6OjEzMzIyMzY0NzoxNjA4MTAwOTcyOjo6MTMzNzQ0MTY0MDoyMjc4Mjg6MDoxZTIzZjhkMzY0MzllNGI1NDgzOTZkMzhhZDA4NjI4M2M6ZGVmYXVsdF83OjA%3D&cookie_ticket=abf4dd51d5eaf64cf6ec9de8710c1d02&sidlist=194,193&app_flag=0B1U&show_enddate=1
`
	o:=strings.Index(aa,"\n\n")
fmt.Println(o)
fmt.Println(aa[o+1:])
}
