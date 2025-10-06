package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	sqlxcurd "github.com/leiquan09-lang/web3_home/db/sqlxCurd"
)

func main() {
	roots := "root:123456@127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := roots
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err) // 或 log.Fatal(err)
	// }
	// 创建表
	// db.AutoMigrate(&curd.Students{})
	// 插入数据
	// db.Create(&curd.Students{Name: "张三", Age: 20, Grade: "三年级"})
	// 查询年纪>18
	// var list []curd.Students
	// db.Where("age >?", 18).Find(&list)
	// fmt.Println("年龄>18:", list)
	//3 更新张三年级
	// db.Model(&curd.Students{}).Where("name =?", "张三").Update("grade", "4️⃣年级")
	// 4删除
	// db.Where("age<?", 15).Delete(&curd.Students{})

	// // 建表
	// db.AutoMigrate(&curd.Account{}, &curd.Transaction{})

	// // 初始化测试数据
	// db.Create(&curd.Account{ID: 1, Balance: 1000})
	// db.Create(&curd.Account{ID: 2, Balance: 500})

	// 执行转账
	// if err := curd.Transfer(db, 1, 2, 100); err != nil {
	// 	fmt.Println("转账失败:", err)
	// } else {
	// 	fmt.Println("转账成功")
	// }

	db, err := sqlx.Connect("mysql", roots)
	if err != nil {
		log.Fatalf("connect failed: %v", err)
	}
	defer db.Close()
	// 2. 查询所有技术部员工
	var emps []sqlxcurd.Employee
	err = db.Select(&emps, "SELECT * FROM employees WHERE department = ?", "技术部")
	if err != nil {
		log.Fatalf("select tech failed: %v", err)
	}
	fmt.Println("技术部员工：", emps)

	// 3. 查询工资最高的员工（单条）
	var top sqlxcurd.Employee
	err = db.Get(&top, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("select top failed: %v", err)
	}
	fmt.Println("工资最高：", top)

	// 查询价格 > 50 的书籍
	var books []sqlxcurd.Book
	err = db.Select(&books, "SELECT * FROM books WHERE price > ?", 50)
	if err != nil {
		log.Fatalf("select books failed: %v", err)
	}

	for _, b := range books {
		log.Printf("id=%d title=%s author=%s price=%.2f\n",
			b.ID, b.Title, b.Author, b.Price)
	}
}
