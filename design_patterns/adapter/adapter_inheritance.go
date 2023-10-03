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

// NewTaro 構造体はTaroのサブクラスで、Chairpersonインターフェースを実装します。Adapter
type NewTaro struct {
	*Taro
}

// OrganizeClass はChairpersonインターフェースのメソッドで、クラスを組織することを表します。
func (nt *NewTaro) OrganizeClass() {
	nt.EnjoyWithAllClassmate()
}

// Teacher 構造体はTeacherを表します。Client
type Teacher struct {
}

// Order はTeacherのメソッドで、Taroに命令ができます。
func (te *Teacher) Order() {
	var chairperson Chairperson = &NewTaro{&Taro{}}
	chairperson.OrganizeClass()
}

func main() {
	teacher := &Teacher{}
	teacher.Order()
}
