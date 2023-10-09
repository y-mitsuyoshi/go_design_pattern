package main

import "fmt"

// AbstractClass はテンプレートメソッドを持つ抽象クラスです。
type AbstractClass interface {
	TemplateMethod()
	step1()
	step2()
	step3()
}

// ConcreteClassA は AbstractClass の具象クラスです。
type ConcreteClassA struct{}

func (c *ConcreteClassA) TemplateMethod() {
	c.step1()
	c.step2()
	c.step3()
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
type ConcreteClassB struct{}

func (c *ConcreteClassB) TemplateMethod() {
	c.step1()
	c.step2()
	c.step3()
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
	concreteA.TemplateMethod()

	fmt.Println("\nUsing ConcreteClassB:")
	concreteB := &ConcreteClassB{}
	concreteB.TemplateMethod()
}
