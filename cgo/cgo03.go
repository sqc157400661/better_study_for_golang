package main

/*
#include <stdio.h>
void SayHello(const char* s);
*/
import "C"

func main()  {
	C.SayHello(C.CString("Hello,World\n"))
}