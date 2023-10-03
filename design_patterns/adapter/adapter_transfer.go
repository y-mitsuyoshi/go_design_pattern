package main

import "fmt"

// Taro 構造体はTaroを表します。Adaptee
type Taro struct{}

// EnjoyWithAllClassmate はTaroのメソッドで、みんなで楽しむことを表します。
func (t *Taro) EnjoyWithAllClassmate() {
	fmt.Println("みんなで楽しむ")
}

// Chairperson インターフェースはChairpersonを表します。Target
type Chairperson interface {
	OrganizeClass()
}

// Hanako 構造体はChairpersonインターフェースを実装します。Adapter
type Hanako struct {
	taro *Taro
}

// OrganizeClass はChairpersonインターフェースのメソッドで、クラスを組織することを表します。
func (h *Hanako) OrganizeClass() {
	h.taro.EnjoyWithAllClassmate()
}

// Teacher 構造体はTeacherを表します。Client
type Teacher struct {
}

// Order はTeacherのメソッドで、Hanakoに命令ができます。
func (te *Teacher) Order() {
	var chairperson Chairperson = &Hanako{&Taro{}}
	chairperson.OrganizeClass()
}

func main() {
	teacher := &Teacher{}
	teacher.Order()
}
