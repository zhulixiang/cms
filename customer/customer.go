package customer

import (
	"fmt"
)

type Customers struct {
	Customerslice
}

type Customerslice []Customer
type Customer struct {
	Name string
	Sex string
	Age int
	Phone string
	Mail string
}

//创建客户列表
func NewCustomers()*Customers{
	return &Customers{}
}

//创建客户
func NewCustomer(name string ,sex string ,age int ,phone string ,mail string)Customer{
	return Customer{Name: name,Sex: sex,Age: age,Phone: phone,Mail: mail}
}

//添加客户
func (customers *Customers)AddCustomer(customer Customer){
	customers.Customerslice = append(customers.Customerslice,customer)
}

//删除客户
func (customers *Customers)DelCustomer(index int){
	defer func() {
		if err:= recover();err != nil {
			fmt.Println("删除客户出错",err)
		}
	}()
	if index >= 0 && index < len(customers.Customerslice)  {
		if index + 1 >= len(customers.Customerslice) {
			customers.Customerslice = append(customers.Customerslice[:index])
		}else {
			customers.Customerslice = append(customers.Customerslice[:index],customers.Customerslice[index+1])
		}

	}else {
		fmt.Println("删除的客户编号不存在")
	}
}
//修改客户
func (customers *Customers)UpdateCustomer(index int,customer Customer){
	if index >= 0 && index < len(customers.Customerslice)  {
		customers.Customerslice[index].Name = customer.Name
		customers.Customerslice[index].Age = customer.Age
		customers.Customerslice[index].Sex = customer.Sex
		customers.Customerslice[index].Phone = customer.Phone
		customers.Customerslice[index].Mail = customer.Mail
	}else {
		fmt.Println("修改的客户编号不存在")
	}
}

//修改客户
func (customers *Customers)QueryCustomers(){
	if len(customers.Customerslice) <= 0 {
		fmt.Println("没有客户，请先添加客户")
	}else{
		fmt.Printf("客户编号\t 客户姓名 \t客户性别\t客户年龄\t客户电话\t客户邮箱\n ")
		for index,val := range customers.Customerslice{
			fmt.Printf("%v\t %v \t %v\t%v\t%v\t%v\n",index,val.Name,val.Sex,val.Age,val.Phone,val.Mail)
		}
	}
}








