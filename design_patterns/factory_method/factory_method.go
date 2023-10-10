package main

import (
	"fmt"
	"math"
)

// Shape インターフェースは、各種形状の共通のメソッドを定義します。
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle は円を表す構造体です。
type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle は長方形を表す構造体です。
type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// ShapeFactory インターフェースは、Shape オブジェクトを生成するためのファクトリメソッドを定義します。
type ShapeFactory interface {
	CreateShape() Shape
}

// CircleFactory は Circle オブジェクトを生成するファクトリです。
type CircleFactory struct {
	Radius float64
}

func (cf *CircleFactory) CreateShape() Shape {
	return &Circle{Radius: cf.Radius}
}

// RectangleFactory は Rectangle オブジェクトを生成するファクトリです。
type RectangleFactory struct {
	Width  float64
	Height float64
}

func (rf *RectangleFactory) CreateShape() Shape {
	return &Rectangle{Width: rf.Width, Height: rf.Height}
}

func main() {
	// 円を生成して面積と周囲の長さを計算
	circleFactory := &CircleFactory{Radius: 5}
	circle := circleFactory.CreateShape()
	fmt.Printf("円の面積: %.2f\n", circle.Area())
	fmt.Printf("円の周囲の長さ: %.2f\n", circle.Perimeter())

	// 長方形を生成して面積と周囲の長さを計算
	rectangleFactory := &RectangleFactory{Width: 4, Height: 6}
	rectangle := rectangleFactory.CreateShape()
	fmt.Printf("長方形の面積: %.2f\n", rectangle.Area())
	fmt.Printf("長方形の周囲の長さ: %.2f\n", rectangle.Perimeter())
}
