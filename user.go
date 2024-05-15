package todo_app

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}
