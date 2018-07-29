package main

import (
	"fmt"
	"reflect"
)

type ControllerInterface interface {
	Init(action string, method string)
}
type Controller struct {
	Action string
	Method string
	Tag    string `json:"tag"`
}

func (c *Controller) Init(action string, method string) {
	c.Action = action
	c.Method = method
	//增加fmt打印，便于看是否调用
	fmt.Println("Init() is run.")
	fmt.Println("c:", c)
}

//增加一个无参数的Func
func (c *Controller) Test() {
	fmt.Println("Test() is run.")
}

func main() {
	fmt.Println("start。。。。。")
	c := &Controller{
		Action: "index",
		Method: "Get",
	}

	// 实现接口方法
	var i ControllerInterface

	i = c

	v := reflect.ValueOf(i)
	fmt.Printf("value: %v \n", v)

	t := reflect.TypeOf(i)

	fmt.Printf("type: %v \n",t)

	// Elem返回值v包含的接口
	controllerValue := v.Elem()
	//fmt.Println("controllerType(reflect.Value):",controllerType)
	//获取存储在第一个字段里面的值
	fmt.Println("Action:", controllerValue.Field(0).String())

	fmt.Println("by name:", controllerValue.FieldByName("Action"))
	// Elem返回类型的元素类型。
	controllerType := t.Elem()
	tag := controllerType.Field(2).Tag //Field(第几个字段,index从0开始)
	fmt.Println("Tag:", tag)
}
