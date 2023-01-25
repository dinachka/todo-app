package todo

type TodoList struct {
	Id          int64  `json:"id" db:"id"`
	Title       string `json:"title"  db:"title" binding:"required"`
	Description string `json:"description"  db:"description"`
}

type UsersList struct {
	Id     int64
	UserId int64
	ListId int64
}

type TodoItem struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListItem struct {
	Id      int
	ListId int
	ItemId  int
}
