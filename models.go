package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Todo struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	IsDone uint8  `json:"isDone"`
}

func (b *Todo) TableName() string {
	return "todo"
}

func GetAllTodosFromDB(todo *[]Todo) (err error) {
	if err = DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateATodoInDB(todo *Todo) (err error) {
	if err = DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetATodoFromDB(todo *Todo, id string) (err error) {
	if err := DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateATodoInDB(todo *Todo, id string) (err error) {
	fmt.Println(todo)
	DB.Save(todo)
	return nil
}

func DeleteATodoFromDB(todo *Todo, id string) (err error) {
	DB.Where("id = ?", id).Delete(todo)
	return nil
}
