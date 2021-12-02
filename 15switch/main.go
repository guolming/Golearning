package main

import (
	"fmt"
	"reflect"
)

func main() {
	var money interface{} = 10
	//var money interface{} =10.0
	//var money interface{} = "10"
	switch newMoney := money.(type) {
	case int:
		tmpMoney := newMoney + 3
		fmt.Println("money is int", tmpMoney)
	case int64:
		fmt.Println("money is int64")
	case float64:
		fmt.Println("money is float64")
	case float32:
		fmt.Println("money is float32")
	default:
		fmt.Println("money is unknow type")
	}
	fmt.Println("结束", money, reflect.TypeOf(money))
	money = 32
	fmt.Println("结束", money, reflect.TypeOf(money))
}
