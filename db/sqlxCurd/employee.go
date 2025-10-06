package sqlxcurd

// 1. 定义结构体，字段 tag 统一用 db 指明列名
type Employee struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}
