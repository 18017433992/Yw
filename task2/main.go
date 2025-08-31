package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。
func Sum(num *int) {
	result := *num + 10
	fmt.Println("result", result)
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func mul(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] *= 2
	}
	fmt.Println("arr", *arr)
}

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
type count struct {
	mu    sync.Mutex
	count int
}

func (c *count) saveCount() {
	start := time.Now()
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
	elapsed := time.Since(start)
	fmt.Printf("耗时: %.10f 毫秒\n", float64(elapsed.Milliseconds()))

}
func (c *count) getCount() int {
	return c.count
}

// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，
// 实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
type shape interface {
	Area() float64      //面积
	Perimeter() float64 //周长
}
type Rectangle struct {
	Width  float64 //矩形的宽
	Height float64 //矩形的高
}
type Circle struct {
	Radius float64 // 圆的半径
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，
// 组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
type Person struct {
	Name string
	age  int
}
type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d, EmployeeID: %d\n", e.Name, e.age, e.EmployeeID)
}

// 向通道写入数据
func wirtechann(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Println("写入数据:", i)
	}
	close(ch)
}

// 读取通道数据
func readchan(ch <-chan int) {
	for v := range ch {
		fmt.Println("读到数据:", v)
	}

}

// 向通道写入100个数据
func writechan100(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Println("写入数据:", i)
		time.Sleep(time.Millisecond)
	}
	close(ch)
}

// 读取通道写入的100个数据
func readchan100(ch <-chan int) {
	for v := range ch {
		fmt.Println("读取数据：", v)
		time.Sleep(time.Second)
	}
}

// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func main() {
	var wg sync.WaitGroup
	var counter int64

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 1; j <= 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("最终的计数器值:", counter)
}

//编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// func main() {
// 	var mu sync.Mutex
// 	var wg sync.WaitGroup
// 	count := 0

// 	for i := 1; i <= 10; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for j := 1; j <= 1000; j++ {
// 				mu.Lock()
// 				count += j
// 				mu.Unlock()
// 			}
// 		}()

// 	}
// 	wg.Wait()
// 	fmt.Println("最终的i值:", count)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan int, 100)
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		writechan100(ch)
// 	}()
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		readchan100(ch)
// 	}()
// 	wg.Wait()
// 	fmt.Println("所有操作完成")
// }

// func demonstrateBufferedChannel(bufferSize int) {
// 	fmt.Printf("\n=== 缓冲通道演示 (大小: %d) ===\n", bufferSize)

// 	ch := make(chan int, bufferSize)
// 	var wg sync.WaitGroup

// 	// 快速生产者
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for i := 1; i <= 8; i++ {
// 			ch <- i
// 			fmt.Printf("快速生产者 → 写入: %d (通道: %d/%d)\n",
// 				i, len(ch), cap(ch))
// 			time.Sleep(100 * time.Millisecond) // 快速生产
// 		}
// 		close(ch)
// 	}()

// 	// 慢速消费者
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for data := range ch {
// 			fmt.Printf("慢速消费者 ← 读取: %d (通道: %d/%d)\n",
// 				data, len(ch), cap(ch))
// 			time.Sleep(300 * time.Millisecond) // 慢速消费
// 		}
// 		fmt.Println("消费者: 通道已关闭")
// 	}()

// 	wg.Wait()
// }

// func main() {
// 	fmt.Println("通道缓冲机制演示")
// 	fmt.Println("=================")

// 	// 演示不同缓冲大小的影响
// 	demonstrateBufferedChannel(0) // 无缓冲
// 	demonstrateBufferedChannel(2) // 小缓冲
// 	demonstrateBufferedChannel(5) // 大缓冲

// 	// 详细演示无缓冲通道
// 	// demonstrateUnbufferedChannel()
// }

// var wg sync.WaitGroup
// ch := make(chan int)

// wg.Add(1)
// go func() {
// 	defer wg.Done()
// 	wirtechann(ch)
// }()
// wg.Add(1)
// go func() {
// 	defer wg.Done()
// 	readchan(ch)
// }()
// wg.Wait()
// emp := Employee{Person: Person{Name: "张先生", age: 32}, EmployeeID: 12345}
// emp.PrintInfo()

// c := count{}
// var wg sync.WaitGroup
// for i := 0; i < 100; i++ {
// 	wg.Add(1) // 等待1个goroutine
// 	go func() {
// 		defer wg.Done()
// 		c.saveCount()
// 	}()
// }
// wg.Wait() // 等待所有goroutine完成
// fmt.Println("count:", c.getCount())
// value := 20
// Sum(&value)
// arr := []int{1, 2, 3, 4, 5}
// mul(&arr)
// var wg sync.WaitGroup
// wg.Add(2) // 等待2个goroutine
// go func() {
// 	defer wg.Done() //奇数goroutine完成时调用
// 	for i := 1; i <= 10; i++ {
// 		if i%2 == 1 {
// 			fmt.Println("奇数：", i)
// 		}
// 	}
// }()
// go func() {
// 	defer wg.Done() //偶数goroutine完成时调用
// 	for i := 1; i <= 10; i++ {
// 		if i%2 == 0 {
// 			fmt.Println("偶数：", i)
// 		}
// 	}
// }()
// wg.Wait() // 等待所有goroutine完成
