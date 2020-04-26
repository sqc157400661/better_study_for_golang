/**
	学习接口
 */
package main

import "fmt"

func main() {
	// 在处理来自于外部的、类型未知的数据时，比如解析诸如 JSON 或 XML 编码的数据，类型测试和转换会非常有用。
	classifier(13, -14.3, "BELGIUM", complex(1, 2), nil, false)
}
func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("Param #%d is a bool\n", i)
		case float64:
			fmt.Printf("Param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("Param #%d is a int\n", i)
		case nil:
			fmt.Printf("Param #%d is a nil\n", i)
		case string:
			fmt.Printf("Param #%d is a string\n", i)
		default:
			fmt.Printf("Param #%d is unknown\n", i)
		}
	}
}




