package cls2


import "fmt"
import "../base"

type Class2 struct {
}

func (c *Class2) Do() {
	fmt.Println("Class2")
}

func init() {

	// 在启动时注册类2工厂
	base.Register("Class2", func() base.Class {
		return new(Class2)
	})

}

