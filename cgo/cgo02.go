// 使用自己的C函数
package main


/*
#include <stdio.h>
static void SayHello(const char* s){
	puts(s);
}
 */
import "C"

func main()  {
	C.SayHello(C.CString("Hello,World\n"))
}