package controller

import (
	"bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodoHandler(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	err := models.CreateTodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetTodosHandler(c *gin.Context) {
	todoList, err := models.SelectTodos()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateTodoHandler(c *gin.Context) {
	idStr, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}

	todo, err := models.GetById(idStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)
	err = models.UpdateATodo(todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
	/*//将请求体的内容输出到todo变量中,传过来的只有修改的status字段，没有id和title字段
	var todo models.Todo
	c.BindJSON(&todo)
	id, _ := strconv.Atoi(idStr)
	todo.ID = id
	err := models.UpdateTodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}*/
}

func DeleteTodoHandler(c *gin.Context) {
	idStr, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	err := models.DeleteTodo(idStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"id": "deleted"})
	}
}
