package main

func replaceSpace(s string) string {
	re :=[]byte{}
	for i:=0;i<len(s);i++{
		if s[i] ==' ' {
			re = append(re,[]byte{'%','2','0'}...)
		}else{
			re = append(re,s[i])
		}
	}
	return string(re)
}
