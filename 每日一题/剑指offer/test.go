package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := `[2021-04-07 13:52:06.703] [TRACE] [1347853059814160AF-3205-4C7E-8889-435CD09440B3] [upass] [upass_8000_unisdk_unicom] [{"req_app_flag":"0E","req_cip":"172.16.216.28","req_client_id":"99166000000000001347","req_token":"3aa5c1429ff74013aa9ee85efa96abbe","req_utoa":"1","req_mtuserid":"602599871","req_did":"B480D5F8ED2749418EF16BB7DA883C6C","req_dpass":"Xs6sHVit0jerV5EyNYQlxVTXNf9xvLnESOh+wVAbMlm8CFtEaYXTKB1cuO1Cebzp","req_sdtis":"c26","req_ttype":"IOS","req_devname":"h%E7%9A%84 iPhone","sources":"unisdk_unicom","upass_cip":"172.17.80.43","x-real-ip":"172.17.80.43","client-ip":"172.17.80.43, 172.17.80.43, 172.17.80.43","response_code":0,"response_errmsg":"\u767b\u5f55\u6210\u529f"}]`
	matched, err := regexp.MatchString(`[:^((\"response_code\":(\"|')?(402|0|200|\-11202|\-11204|\-3|3050|201|\-11206|\-11207)))):]`, str)
	fmt.Println(matched, err)
}