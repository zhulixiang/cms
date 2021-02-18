package main

import (
	"cms/menu"
	"fmt"
)

func main(){
	fmt.Println("客户信息管理系统")
	menu.NewMenu().ShowMenu()
}
