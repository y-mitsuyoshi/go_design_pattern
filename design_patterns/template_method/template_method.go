package main

import "fmt"

// AbstractClass はテンプレートメソッドを持つ抽象クラスです。
type IAbstractClass interface {
	step1()
	step2()
	step3()
}
type AbstractClass struct {
	iAbstractClass IAbstractClass
}

func (c *AbstractClass) TemplateMethod() {
	c.iAbstractClass.step1()
	c.iAbstractClass.step2()
	c.iAbstractClass.step3()
}

// ConcreteClassA は AbstractClass の具象クラスです。
type ConcreteClassA struct {
	AbstractClass
}

func (c *ConcreteClassA) step1() {
	fmt.Println("ConcreteClassA: Step 1")
}

func (c *ConcreteClassA) step2() {
	fmt.Println("ConcreteClassA: Step 2")
}

func (c *ConcreteClassA) step3() {
	fmt.Println("ConcreteClassA: Step 3")
}

// ConcreteClassB は AbstractClass の別の具象クラスです。
type ConcreteClassB struct {
	AbstractClass
}

func (c *ConcreteClassB) step1() {
	fmt.Println("ConcreteClassB: Step 1")
}

func (c *ConcreteClassB) step2() {
	fmt.Println("ConcreteClassB: Step 2")
}

func (c *ConcreteClassB) step3() {
	fmt.Println("ConcreteClassB: Step 3")
}

func main() {
	fmt.Println("Using ConcreteClassA:")
	concreteA := &ConcreteClassA{}
	a := AbstractClass{
		iAbstractClass: concreteA,
	}
	a.TemplateMethod()

	fmt.Println("Using ConcreteClassB:")
	concreteB := &ConcreteClassB{}
	b := AbstractClass{
		iAbstractClass: concreteB,
	}
	b.TemplateMethod()
}
