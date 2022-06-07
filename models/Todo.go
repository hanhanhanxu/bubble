package models

import "bubble/dao"

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// todo 增删改查

func CreateTodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func SelectTodos() (todoList []*Todo, err error) {
	err = dao.DB.Find(&todoList).Error
	return
}

func UpdateTodo(todo *Todo) (err error) {
	id := todo.ID
	//错误原因:这里查询时已经将数据库的数据映射到todo变量上了，所以覆盖了前端传过来的修改的字段。所以下面修改时，会发现没有什么可修改的。
	err = dao.DB.Where("id=?", id).First(todo).Error
	if err != nil {
		return err
	}
	err = dao.DB.Save(todo).Error
	return
}

func GetById(id string) (todo *Todo, err error) {
	todo = new(Todo)
	err = dao.DB.Where("id=?", id).First(todo).Error
	if err != nil {
		return nil, err
	}
	return
}
func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteTodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(Todo{}).Error
	return
}
