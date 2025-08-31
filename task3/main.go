package main

import (
	"example.com/myproject/lesson"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:Zk12345678@tcp(127.0.0.1:3306)/learndb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//lesson.InitCreateT(db)

	// // {{"Name", "张三"}, {"Name", "李四"}, {"Name", "王五"}}

	// // lesson.InitCreateI(db, u)
	// lesson.InitCreateI(db)
	// u := &lesson.User1{}
	// u.ID = 4
	// lesson.FindUserPostAndComment(u, db)
	//lesson.InitCreateI(db)
	c := &lesson.Comment{PostID: 1}
	if err := db.Debug().Where("post_id", c.PostID).Delete(c).Error; err != nil {
		return
	}

}

// dbconfig := mysql.Config{clear
// 	User:                 "root",
// 	Passwd:               "Zk12345678",
// 	Net:                  "tcp",
// 	Addr:                 "127.0.0.1:3306",
// 	DBName:               "learndb",
// 	AllowNativePasswords: true,
// 	Params: map[string]string{
// 		"charset":   "utf8mb4",
// 		"parseTime": "True",
// 		"loc":       "Local",
// 	},
// }

// db, err := sqlx.Connect("mysql", dbconfig.FormatDSN())
// if err != nil {
// 	log.Fatal(err)
// }
// defer db.Close()
// lesson.InitEmployeeTable(db)
// emp := []lesson.Employee{
// 	{Name: "张三", Department: "研发部", Salary: 8000.0},
// 	{Name: "李四", Department: "研发部", Salary: 7000.0},
// 	{Name: "王五", Department: "技术部", Salary: 6000.0},
// 	{Name: "赵六", Department: "技术部", Salary: 5000.0},
// }
// lesson.CreateEmployees(db, emp)
// emp, _ := lesson.FindHighSalary(db)
// fmt.Println(emp)
// lesson.BookCreateTable(db)
// booklist := &[]lesson.Book{{Title: "书名1", Author: "作者1", Price: 100.0},
// 	{Title: "书名2", Author: "作者2", Price: 200.0},
// 	{Title: "书名3", Author: "作者3", Price: 30.0}}
// lesson.BookCreateBooks(db, *booklist)
// booklist, _ := lesson.FindBooks(db)
// for _, book := range *booklist {
// 	log.Printf("书籍ID: %d, 书名: %s, 作者: %s, 价格: %.2f\n", book.ID, book.Title, book.Author, book.Price)
// }
//lesson.RunCreateAndInit(db)
//lesson.RunGetBalance(db)
//lesson.RunInsert(db)
//lesson.RunSelect(db)
//lesson.RunUpdate(db)
//lesson.RunDelete(db)
// result := lesson.RunTransaction(db, 2, 1, 1001)
// if result != nil {
// 	fmt.Println("转账失败:", result)
// }
//lesson.RunEmployees(db)
// emp := lesson.RunFindEmployees(db)
// for _, e := range *emp {
// 	fmt.Printf("员工ID: %d, 姓名: %s, 部门: %s, 工资: %.2f\n", e.ID, e.Name, e.Department, e.Salary)
// }
//
//lesson.BookInit(db)
// booklist := lesson.BookFind(db)
// fmt.Println(booklist)

//lesson.Run(db)
// db.Debug().First(&lesson.Blog2{})
// fmt.Println("查询成功")
// blog2 := &lesson.Blog2{}
// db.Debug().Where("id = ? AND author_name = ?", 22, "John Doe").First(blog2)
// fmt.Println(blog2)
