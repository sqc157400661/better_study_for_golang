package main
import(
	"fmt"
)
func main() {
	var (
		name string
		sex byte
		age byte
		isVip bool
		account float64
	)
	//fmt.Println("请输入姓名")
	//fmt.Scanln(&name)
	//
	//fmt.Println("请输入性别")
	//fmt.Scanln(&sex)
	//
	//fmt.Println("请输入年龄")
	//fmt.Scanln(&age)
	//
	//fmt.Println("是否是会员")
	//fmt.Scanln(&isVip)
	//
	//fmt.Println("系统余额")
	//fmt.Scanln(&account)



	//fmt.Printf("姓名：%v\t性别：%v\t年龄：%v\t是否是会员：%v\t系统余额：%v\t", name, sex, age, isVip, account)

	fmt.Println("请输入姓名 性别 年龄 是否是会员 系统余额 按空格隔开")
	fmt.Scanf("%s %d %t %f", &name, &sex, &age, &isVip, &account)
	fmt.Printf("姓名：%v\t性别：%v\t年龄：%v\t是否是会员：%v\t系统余额：%v\t", name, sex, age, isVip, account)
}