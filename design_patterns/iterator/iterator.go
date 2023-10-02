package main

import "fmt"

// イテレーターを表すインターフェース
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// ToDoアイテムを表す構造体
type TodoItem struct {
	Title string
}

// ToDoリストを表す構造体
type TodoList struct {
	items []TodoItem
}

// ToDoリストにアイテムを追加
func (tl *TodoList) AddItem(item TodoItem) {
	tl.items = append(tl.items, item)
}

// イテレーターを生成
func (tl *TodoList) CreateTodoListIterator() Iterator {
	return &TodoListIterator{
		todoList: tl,
		index:    0,
	}
}

// イテレーターを表す構造体
type TodoListIterator struct {
	todoList *TodoList
	index    int
}

// 次のアイテムが存在するか確認
func (ti *TodoListIterator) HasNext() bool {
	return ti.index < len(ti.todoList.items)
}

// 次のアイテムを取得
func (ti *TodoListIterator) Next() interface{} {
	if !ti.HasNext() {
		return nil
	}
	item := ti.todoList.items[ti.index]
	ti.index++
	return item
}

func main() {
	todoList := &TodoList{}
	todoList.AddItem(TodoItem{Title: "洗濯をする"})
	todoList.AddItem(TodoItem{Title: "買い物に行く"})
	todoList.AddItem(TodoItem{Title: "仕事をする"})

	iterator := todoList.CreateTodoListIterator()

	fmt.Println("ToDoリスト:")
	for iterator.HasNext() {
		item := iterator.Next().(TodoItem)
		fmt.Printf("- %s\n", item.Title)
	}
}
