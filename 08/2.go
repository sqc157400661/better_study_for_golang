/**
	以下代码有什么问题，说明原因。
 */
package main

func main() {
	pase_student()
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	for k,v:=range m{
		println(k,"=>",v.Name)
	}

}

/**
	zhou => wang
	li => wang
	wang => wang

	解答：
	for range 都是使用副本的方式
	所以m[stu.Name]=&stu实际上一致指向同一个指针， 最终该指针的值为遍历的最后一个struct的值拷贝

	正确方式：
	for i:=0;i<len(stus);i++  {
        m[stus[i].Name] = &stus[i]
    }

 */