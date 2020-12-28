package main

import (
	"context"
	"fmt"
)

func main(){
	ctx :=context.Background()
	aa,_ := ctx.Value("sdfsadf").(map[string]bool)

	if aa["aa"] {
		aa["aa"] = true
	}
	fmt.Println(aa["bb"])
}
