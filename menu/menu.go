package menu

import (
	"cms/customer"
	"fmt"
)

type Menu struct {

}

func NewMenu()*Menu{
	return &Menu{}
}

func (menu *Menu)ShowMenu(){
	fmt.Println("1.添加客户")
	fmt.Println("2.修改客户")
	fmt.Println("3.删除客户")
	fmt.Println("4.客户列表")
	fmt.Println("5.退出")
	menu.Operate()
}
func (menu *Menu) Operate() {

	var customers = customer.NewCustomers()

	for {
		fmt.Println("请输入操作的功能编号：")
		var catagory int
		fmt.Scanln(&catagory)
		var isContinue bool = true
		switch catagory {
		case 1:
			var customerName string
			var customerPhone string
			fmt.Println("请输入客户姓名:")
			fmt.Scanln(&customerName)
			fmt.Println("请输入客户电话:")
			fmt.Scanln(&customerPhone)
			customer := customer.NewCustomer(customerName,"男",30,customerPhone,"zhulixiang@sina.com")
			customers.AddCustomer(customer)
		case 2:
			var customerNo = 0
			var customerName string
			var customerPhone string
			fmt.Println("请输入修改的客户编号：")
			fmt.Scanln(&customerNo)
			fmt.Println("请输入客户姓名:")
			fmt.Scanln(&customerName)
			fmt.Println("请输入客户电话:")
			fmt.Scanln(&customerPhone)
			customer := customer.NewCustomer(customerName,"男",30,customerPhone,"zhulixiang@sina.com")
			customers.UpdateCustomer(customerNo,customer)
		case 3:
			var customerNo = 0
			fmt.Println("请输入删除的客户编号：")
			fmt.Scanln(&customerNo)
            customers.DelCustomer(customerNo)
		case 4:
			customers.QueryCustomers()
		case 5:
			fmt.Println("退出")
			isContinue = false
		default:
		//	fmt.Println("没有该功能，退出")
		//	isContinue = false
		}
		if !isContinue {
			break
		}
	}
}