package lesson

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
type Student struct {
	ID    uint `gorm:"primaryKey;autoIncrement:true"`
	Name  string
	Age   uint8
	Grade string
	gorm.Model
}

//编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。

func RunInsert(db *gorm.DB) {
	db.AutoMigrate(&Student{})
	student := &Student{Name: "张三", Age: 20, Grade: "三年级"}
	db.Create(student)
	fmt.Println("插入成功")
}

//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。

func RunSelect(db *gorm.DB) {
	student := &[]Student{}
	db.Where("age > ?", 18).Find(student)
	fmt.Println(student)
}

//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。

func RunUpdate(db *gorm.DB) {
	student := &[]Student{}
	db.Where("name = ?", "张三").Find(student).Update("grade", "四年级")
	fmt.Println(student)
}

//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

func RunDelete(db *gorm.DB) {
	student := &[]Student{}
	db.Where("age < ?", 15).Find(student).Delete(student)
	fmt.Println(student)
}

// Accounts 表（包含字段 id 主键， balance 账户余额）
type Account struct {
	ID      uint `gorm:"primaryKey;autoIncrement:true"`
	Balance float64
}

// Aransactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
type Transactions struct {
	ID            uint    `gorm:"primaryKey;autoIncrement:true"`
	FromAccountID uint    `gorm:"not null;column:"from_account_id"`
	Amount        float64 `gorm:"not null"`
	ToAccountID   uint    `gorm:"not null;column:"to_account_id"`
}

// 建表，初始化数据
func RunCreateAndInit(db *gorm.DB) {
	db.AutoMigrate(&Account{}, &Transactions{})
	account1 := &Account{Balance: 1000.0}
	account2 := &Account{Balance: 2000.0}
	db.Create(account1)
	db.Create(account2)
	fmt.Println("插入成功")
}

// 查询初始化的数据
func RunGetBalance(db *gorm.DB) {
	accounts := &[]Account{}
	db.Find(&accounts)
	for _, acc := range *accounts {
		fmt.Printf("账户ID: %d, 余额: %.2f\n", acc.ID, acc.Balance)
	}
}

// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，
// 如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。
// 如果余额不足，则回滚事务。
func RunTransaction(db *gorm.DB, fromAccountID uint, toAccountID uint, amount float64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		fromAccount := &Account{}
		toAccount := &Account{}
		tx.Where("id=?", fromAccountID).First(fromAccount)
		tx.Where("id=?", toAccountID).First(toAccount)
		if err := tx.First(fromAccount, fromAccountID).Error; err != nil {
			return fmt.Errorf("账户A不存在")
		}
		if err := tx.First(toAccount, toAccountID).Error; err != nil {
			return fmt.Errorf("账户B不存在")
		}
		if fromAccount.Balance < amount {
			return fmt.Errorf("余额不足")
		}
		fromAccount.Balance -= amount
		toAccount.Balance += amount
		if err := tx.Save(fromAccount).Error; err != nil {
			return err
		}
		if err := tx.Save(toAccount).Error; err != nil {
			return err
		}
		transaction := &Transactions{
			FromAccountID: fromAccount.ID,
			ToAccountID:   toAccount.ID,
			Amount:        amount,
		}
		if err := tx.Create(transaction).Error; err != nil {
			return err
		}
		fmt.Println("转账成功")
		return nil
	})
}

// 定义一个员工结构体
type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

// 建表
func InitEmployeeTable(db *sqlx.DB) error {
	createTableSql := `CREATE TABLE IF NOT EXISTS employee(
			id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		department VARCHAR(255) NOT NULL,
		salary DECIMAL(10, 2) NOT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
	_, err := db.Exec(createTableSql)
	return err
}

// 建员工
func CreateEmployees(db *sqlx.DB, emp []Employee) error {
	_, err := db.NamedExec(`INSERT INTO employee (name, department, salary) 
		  VALUES (:name, :department, :salary)`, emp)
	return err
}

// 查询技术部所有员工
func FindJS(db *sqlx.DB) (emp1 *[]Employee, err error) {
	employees := &[]Employee{}
	err1 := db.Select(employees, "SELECT * FROM employee WHERE department = ?", "技术部")
	if err1 != nil {
		return nil, err
	}
	return employees, err1
}

// 查询工资最高的人
func FindHighSalary(db *sqlx.DB) (emp1 *Employee, err error) {
	employees := &Employee{}
	err = db.Get(employees, "SELECT * FROM employee ORDER BY salary DESC LIMIT 1")
	if err != nil {
		return nil, err
	}
	return employees, nil
}

// 建表，初始化数据
func RunEmployees(db *gorm.DB) {
	db.AutoMigrate(&Employee{})
	employees := &[]Employee{{Name: "张三", Department: "研发部", Salary: 8000.0},
		{Name: "张三", Department: "研发部", Salary: 7000.0},
		{Name: "张三", Department: "技术部", Salary: 6000.0},
		{Name: "张三", Department: "技术部", Salary: 5000.0}}
	db.Create(employees)
	fmt.Println("插入成功")
}

// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中
func RunFindEmployees(db *gorm.DB) (emp *[]Employee) {
	emp1 := &[]Employee{}
	db.Debug().Where("Department", "技术部").Find(emp1)
	// for _, emp := range *emp1 {
	// 	fmt.Printf("员工ID: %d, 姓名: %s, 部门: %s, 工资: %.2f\n", emp.ID, emp.Name, emp.Department, emp.Salary)
	// }
	return emp1
}

// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
func RunFindHighestSalaryEmployee(db *gorm.DB) (emp *Employee) {
	emp1 := &Employee{}
	db.Debug().Order("salary DESC").First(emp1)
	return emp1
}

// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price
type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

// 定义一个 Book 结构体，包含与 books 表对应的字段。
type Book1 struct {
	Book
}

// 建表
func BookCreateTable(db *sqlx.DB) error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS books (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		author VARCHAR(255) NOT NULL,
		price DECIMAL(10, 2) NOT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
	_, err := db.Exec(createTableSQL)
	return err
}

// 初始化数据
func BookCreateBooks(db *sqlx.DB, books []Book) error {
	_, err := db.NamedExec(`INSERT INTO books (title, author, price) 
		  VALUES (:title, :author, :price)`, books)
	return err
}

// 查询所有书籍
func FindBooks(db *sqlx.DB) (booklist *[]Book, err error) {
	booklist = &[]Book{}
	err = db.Select(booklist, "SELECT * FROM books where price > ?", 50)
	if err != nil {
		return nil, err
	}
	return booklist, nil
}

// 建表，初始化数据
func BookInit(db *gorm.DB) {
	db.AutoMigrate(&Book{})
	bookList := &[]Book{
		{Title: "书名1", Author: "作者1", Price: 100.0},
		{Title: "书名2", Author: "作者2", Price: 200.0},
		{Title: "书名3", Author: "作者3", Price: 30.0},
	}
	db.Create(bookList)
	fmt.Println("插入成功")
}

// 查询价格大于50元的数据
func BookFind(db *gorm.DB) (booklist *[]Book) {
	booklist1 := &[]Book{}
	db.Debug().Where("Price > ?", 50).Find(booklist1)
	return booklist1
}

// 用户模型
type User1 struct {
	Name string
	gorm.Model
	Posts     []Post `gorm:"foreignKey:UserID"`
	PostCount uint
}

// 文章模型
type Post struct {
	ARTICLE  string
	UserID   uint
	User     User1     `gorm:"foreignKey:UserID"`
	Comments []Comment `gorm:"foreignKey:PostID"`
	gorm.Model
}

// 评论模型
type Comment struct {
	Comments string
	PostID   uint
	post     Post `gorm:"foreignKey:PostID"`
	gorm.Model
	CommentStatus string `gorm:"comment:有评论、无评论;default:有评论"`
}

// 创建文章后，更新数量
func (p *Post) AfterCreate(db *gorm.DB) error {
	err := db.Debug().Model(&User1{}).Where("id=?", p.UserID).UpdateColumn("post_count", gorm.Expr("post_count+ ?", 1)).Error
	if err != nil {
		return err
	}
	fmt.Println("文章数量更新成功")
	return nil
}

// 删除评论后,更新文章评论状态
func (c *Comment) AfterDelete(db *gorm.DB) error {
	var cnt int64

	fmt.Println(&c.PostID)
	err := db.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&cnt).Error
	if err != nil {
		return err
	}
	if cnt == 0 {
		err1 := db.Debug().Model(&Comment{}).Unscoped().Where("post_id", c.PostID).UpdateColumn("CommentStatus", "无评论").Error
		if err1 != nil {
			return err1
		}
	}

	return nil
}

// 建表
func InitCreateT(db *gorm.DB) error {
	err := db.AutoMigrate(&User1{}, &Post{}, &Comment{})
	return err
}

// 初始化表数据
func InitCreateI(db *gorm.DB) error {
	user01 := User1{Name: "张三"}
	user02 := User1{Name: "李四"}
	user03 := User1{Name: "王五"}

	if err1 := db.Create(&user01).Error; err1 != nil {
		return err1
	}
	if err1 := db.Create(&user02).Error; err1 != nil {
		return err1
	}
	if err1 := db.Create(&user03).Error; err1 != nil {
		return err1
	}
	p1 := Post{ARTICLE: "文章1", UserID: user01.ID}
	p2 := Post{ARTICLE: "文章2", UserID: user02.ID}
	p3 := Post{ARTICLE: "文章3", UserID: user03.ID}
	p4 := Post{ARTICLE: "文章4", UserID: user03.ID}

	if err2 := db.Create(&p1).Error; err2 != nil {
		return err2
	}
	if err2 := db.Create(&p2).Error; err2 != nil {
		return err2
	}
	if err2 := db.Create(&p3).Error; err2 != nil {
		return err2
	}
	if err2 := db.Create(&p4).Error; err2 != nil {
		return err2
	}

	c1 := Comment{Comments: "评论1", PostID: p1.ID}
	c2 := Comment{Comments: "评论2", PostID: p1.ID}
	c3 := Comment{Comments: "评论3", PostID: p2.ID}
	c4 := Comment{Comments: "评论4", PostID: p3.ID}

	if err3 := db.Create(&c1).Error; err3 != nil {
		return err3
	}
	if err3 := db.Create(&c2).Error; err3 != nil {
		return err3
	}
	if err3 := db.Create(&c3).Error; err3 != nil {
		return err3
	}
	if err3 := db.Create(&c4).Error; err3 != nil {
		return err3
	}
	return nil
}

// 查询某个用户的文章及评论
func FindUserPostAndComment(u *User1, db *gorm.DB) error {
	var post []Post
	//fmt.Println(u.ID)
	err := db.Debug().Preload("Comments").Where("user_id = ?", u.ID).Find(&post).Error
	if err != nil {
		return err
	}
	for _, p := range post {
		fmt.Println(p.ID, p.ARTICLE, p.UserID)
		for _, c := range p.Comments {
			fmt.Println(c.Comments)
		}
	}
	fmt.Println("没有查询到相关数据")
	return nil
}

// 查找评论数最多的文章
func FindMaxComments(db *gorm.DB) error {
	var comments []Comment
	var post []Post
	//var count int64
	err := db.Debug().Model(&Comment{}).Select("post_id, count(*) as count").Group("post_id").Order("count desc").Find(&comments).Error
	if err != nil {
		return err
	}
	maxcount := &comments[0]
	// fmt.Println(maxcount.PostID)
	// err1 := db.Debug().Preload("Comments").Select("posts.*,comments.*").Joins("Left join comments on posts.id = comments.post_id").Where("comments.post_id=?", maxcount.PostID).Find(&post).Error
	err1 := db.Debug().
		Preload("Comments").                    // 让 GORM 处理关联加载
		Where("posts.id = ?", maxcount.PostID). // 直接使用帖子ID过滤
		Find(&post).Error

	if err1 != nil {
		return err1
	}
	for _, p := range post {
		fmt.Println(p.ID, p.ARTICLE, p.UserID)
		for _, c := range p.Comments {
			fmt.Println(c.Comments)
		}
	}
	return nil
}
