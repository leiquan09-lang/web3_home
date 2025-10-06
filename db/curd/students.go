package curd

type Students struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"size:50"`
	Age   int
	Grade string `gorm:"size:20"`
}
