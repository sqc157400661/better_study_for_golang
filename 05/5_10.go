/**
	空接口的概念
 */
package main
import "fmt"

// 接口定一个别名类型
type Element interface{}

//定义一个容器类型的结构体 Vector，它包含一个 Element 类型元素的切片
type Vector struct {
	a []Element
}

/*
    说明
    1、Vector 里能放任何类型的变量，因为空接口可以被任何类型实现
    2、Vector 里放的每个元素可以是不同类型的变量
    3、Vector 中存储的所有元素都是 Element 类型，要得到它们的原始类型，需要用到类型断言
*/

// 我们为它定义一个 Get() 方法用于返回第 i 个元素：
func (p *Vector) Get(i int) Element {
	return p.a[i]
}

//再定一个 Set() 方法用于设置第 i 个元素的值：
func (p *Vector) Set(i int, e Element) {
	p.a[i] = e
}
func main()  {
	struct1 := &Vector{make([]Element,10)}
	fmt.Println(struct1)
	i :=1
	struct1.Set(i,"shiqc")
	fmt.Printf("index %d:%T \n",i,struct1.Get(i))
	i=2
	struct1.Set(i,250)
	fmt.Printf("index %d:%T \n",i,struct1.Get(i))
	i=3
	struct1.Set(i,false)
	fmt.Printf("index %d:%T \n",i,struct1.Get(i))
	i=4
	struct1.Set(i,1.02)
	fmt.Printf("index %d:%T \n",i,struct1.Get(i))
}



