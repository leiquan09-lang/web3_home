package sqlxcurd

// 1. 定义 Book 结构体，字段与表列一一对应
type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}
