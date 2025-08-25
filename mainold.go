package main

// "example.com/myproject/utils" //导入本地包

//func main() {

// // utils.Greet("World")
// // result := utils.Add(10, 30)
// // println("Result:", result)
// // 十六进制
// var a uint8 = 0xF
// var b uint8 = 0xf

// // 八进制
// var c uint8 = 017
// var d uint8 = 0o17
// var e uint8 = 0o17

// // 二进制
// var f uint8 = 0b1111
// var g uint8 = 0b1111

// // 十进制
// var h uint8 = 15
// //浮点数
// var float1 float32 = 10
// float2 := 10.0
// fmt.Println(float1 == float32(float2))

// fmt.Println(a == b)
// fmt.Println(b == c)
// fmt.Println(c == d)
// fmt.Println(d == e)
// fmt.Println(e == f)
// fmt.Println(f == g)
// fmt.Println(g == h)

// var s string = "abc，你好，世界！"
// var runes []rune = []rune(s)
// fmt.Println(runes)
// fmt.Println(len(runes))

// var c1 complex64
// c1 = 1.10 + 0.1i
// c2 := 1.10 + 0.1i
// c3 := complex(1.10, 0.1) // c2与c3是等价的
// fmt.Println(c2 == c3)
// fmt.Println(c1)
// a := real(c1)
// b := imag(c1)
// fmt.Println(a, b)

// var s string = "Hello, world!"
// var bytes []byte = []byte(s)
// fmt.Println("convert \"Hello, world!\" to bytes: ", bytes)
// 	var bytes []byte = []byte{72, 101, 108, 108, 111, 44, 32, 119, 111, 114, 108, 100, 33}
// 	var s string = string(bytes)
// 	fmt.Println(s)
//
// var r1 rune = 'a'
// var r2 rune = '世'
// var s string = "abc，你好，世界！"
// var runes []rune = []rune(s)
// fmt.Println(runes)
// fmt.Println(len(runes))
// var s string = "Hello\nworld!\n"
// fmt.Println(s)
// var s string = "Go语言"
// var bytes []byte = []byte(s)
// var runes []rune = []rune(s)

// // fmt.Println("string length: ", len(s))
// // fmt.Println("bytes length: ", len(bytes))
// // fmt.Println("runes length: ", len(runes))
// fmt.Println("string sub: ", s[0:7])
// fmt.Println("bytes sub: ", string(bytes[0:7]))
// fmt.Println("runes sub: ", string(runes[0:3]))
// var a int = 10
// b := "你好"
// fmt.Println(a)
// fmt.Print(b)
// var (
// 	s string = "你好"
// 	a int    = 2
// )
// c := 3
// fmt.Println(s, a, c)
// 	d, e, f := method1(3, 4, 5)
// 	fmt.Println(d, e, f)
// }
// func method1(a, b, c int) (d, e, f int) {

// 	return a + 1, b + 2, c + 3
// a, b, c := 1, 2, 3
// fmt.Println(a, b, c)
// var p1 *int
// var p2 *string

// i := 1
// s := "Hello"
// // 基础类型数据，必须使用变量名获取指针，无法直接通过字面量获取指针
// // 因为字面量会在编译期被声明为成常量，不能获取到内存中的指针信息
// p1 = &i
// p2 = &s

// p3 := &p2         //将P2的地址赋值给P3
// fmt.Println(*p1)  //1
// fmt.Println(*p2)  //hello
// fmt.Println(**p3) //P2的地址

// 	a := 2
// 	var p *int
// 	fmt.Println(&a) //a的地址
// 	p = &a
// 	fmt.Println(p, &a) //p的地址，a的地址

// 	var pp **int
// 	pp = &p
// 	fmt.Println(pp, p) //pp的地址，p的地址
// 	**pp = 3
// 	fmt.Println(pp, *pp, p) //p的地址，p的地址，a的地址
// 	fmt.Println(**pp, *p)   //3，3
// 	fmt.Println(a, &a)      //2，a的地址
//
// type Person struct {
// 	Name  string
// 	Age   int
// 	Call  func() byte
// 	Map   map[string]string
// 	Ch    chan string
// 	Arr   [32]uint8
// 	Slice []interface{}
// 	Ptr   *int
// 	once  sync.Once
// }
// type Custom1 struct {
// 	field1, field2, field3 byte
// }
// //type Other struct {}

// type Student struct {
// 	Name  string            `json:"name" gorm:"column:<name>"`
// 	Age   int               `json:"age" gorm:"column:<name>"`
// 	Call  func()            `json:"-" gorm:"column:<name>"`
// 	Map   map[string]string `json:"map" gorm:"column:<name>"`
// 	Ch    chan string       `json:"-" gorm:"column:<name>"`
// 	Arr   [32]uint8         `json:"arr" gorm:"column:<name>"`
// 	Slice []interface{}     `json:"slice" gorm:"column:<name>"`
// 	Ptr   *int              `json:"-"`
// 	// O     Other              `json:"-"`
// }

// type Custom2 struct {
// 	int
// 	string
// 	Other string
// }
// custom2 := Custom2{
// 	int:    1,
// 	string: "Hello",
// 	Other:  "World",
// }
// var noName = struct {
// 	Field1 int
// 	Field2 string
// 	Field3 float64
// }{}
// var noNameStruct = make(chan struct{}, 0)
// fmt.Println(custom2)
// type A struct {
// 	a string
// }

// type B struct {
// 	A
// 	b string
// }

// type C struct {
// 	A
// 	B
// 	a string
// 	b string
// 	c string
// }
// var d = A{a: "hello"}
// var e = B{A: A{a: "world"}}
// var f = C{A: A{a: "!"}}
// var g = C{B: B{A: A{a: "!"}, b: "Go"}}
// fmt.Println(d.a)
// fmt.Println(e.A.a)
// fmt.Println(f.A.a)
// fmt.Println(g.B.b)

// c := NewC()
// cp := &c
// fmt.Println(c.string())
// fmt.Println(c.stringA())
// fmt.Println(c.stringB())

// fmt.Println(cp.string())
// fmt.Println(cp.stringA())
// fmt.Println(cp.stringB())

// c.setA("1a")
// fmt.Println("------------------c.setA")
// fmt.Println(c.A.a)
// fmt.Println(cp.A.a)

// cp.setA("2a")
// fmt.Println("------------------cp.setA")
// fmt.Println(c.A.a)
// fmt.Println(cp.A.a)

// c.setPA("3a")
// fmt.Println("------------------c.setPA")
// fmt.Println(c.A.a)
// fmt.Println(cp.A.a)

// cp.setPA("4a")
// fmt.Println("------------------cp.setPA")
// fmt.Println(c.A.a)
// fmt.Println(cp.A.a)

// cp.modityD()
// fmt.Println("------------------cp.modityD")
// fmt.Println(cp.d)
// const a int = 1
// const(
// 	b int = 2
// 	c int = 3
// 	d int = 4
// )

// const (
// 	Male   = "Male"
// 	Female = "Female"
// )

// g := Male
// fmt.Println(g.String())
// fmt.Println(g.IsMale())

// type ConnState int
// const (
// 	StateNew ConnState = iota
// 	StateActive
// 	StateIdle
// 	StateHijacked
// 	StateClosed
// )
// fmt.Println(StateNew)
// fmt.Println(StateActive)
// fmt.Println(StateIdle)
// fmt.Println(StateHijacked)
// fmt.Println(StateClosed)

// a, b := 1, 2
// sum := a + b
// sub := a - b
// mul := a * b
// div := a / b
// mod := a % b

// fmt.Println(sum, sub, mul, div, mod)

// a := 10 + 0.1
// b := byte(1) + 1
// fmt.Println(a, b)

// sum := a + float64(b)
// fmt.Println(sum)

// sub := byte(a) - b
// fmt.Println(sub)

// mul := a * float64(b)
// div := int(a) / int(b)

// fmt.Println(mul, div)

// a, b := 1, 2
// var c int
// c = a + b
// fmt.Println("c = a + b, c =", c)
// leftMoveAssignment(c, a)
// rightMoveAssignment(c, a)
// andAssignment(c, a)
// orAssignment(c, a)
// norAssignment(c, a)
// a := 4
// var ptr *int
// fmt.Println(a)

// ptr = &a
// fmt.Printf("*ptr 为 %d\n", *ptr)
// var a int = 10
// if b := 1; a > 10 {
// 	b = 2
// 	// c = 2
// 	fmt.Println("a > 10")
// } else if c := 3; b > 1 {
// 	b = 3
// 	fmt.Println("b > 1")
// } else {
// 	fmt.Println("其他")
// 	if c == 3 {
// 		fmt.Println("c == 3")
// 	}
// 	fmt.Println(b) //1
// 	fmt.Println(c)
// }

// 1. 基本用法
// switch a {
// case "test":
// 	fmt.Println("a = ", a)
// case "s":
// 	fmt.Println("a = ", a)
// case "t", "test string": // 可以匹配多个值，只要一个满足条件即可
// 	fmt.Println("catch in a test, a = ", a)
// case "n":
// 	fmt.Println("a = not")
// default:
// 	fmt.Println("default case")
// }

// 变量b仅在当前switch代码块内有效
// switch b := 5; b {
// case 1:
// 	fmt.Println("b = 1")
// case 2:
// 	fmt.Println("b = 2")
// case 3, 4:
// 	fmt.Println("b = 3 or 4")
// case 5:
// 	fmt.Println("b = 5")
// default:
// 	fmt.Println("b = ", b)
// }
// a := "t1"
// // 不指定判断变量，直接在case中添加判定条件
// b := 5
// switch {
// case a == "t":
// 	fmt.Println("a = t")
// case b == 3:
// 	fmt.Println("b = 3")
// case b == 5, a == "test string":
// 	fmt.Println("a = test string; or b = 5")
// default:
// 	fmt.Println("default case")
// }
// type CustomType struct{}
// var d interface{}
// var e byte = 1
// e := CustomType{}
// d = &e

// switch t := d.(type) {
// case byte:
// 	fmt.Println("d is byte type, ", t)
// case *byte:
// 	fmt.Println("d is byte point type, ", t)
// case *int:
// 	fmt.Println("d is int type, ", t)
// case *string:
// 	fmt.Println("d is string type, ", t)
// case *CustomType:
// 	fmt.Println("d is CustomType pointer type, ", t)
// case CustomType:
// 	fmt.Println("d is CustomType type, ", t)
// default:
// 	fmt.Println("d is unknown type, ", t)
// }

// 方式1
// for i := 0; i < 10; i++ {
// 	fmt.Println("方式1，第", i+1, "次循环")
// }

// // 方式2
// b := 1
// for b < 10 {
// 	fmt.Println("方式2，第", b, "次循环")
// 	b++
// }

// 方式3，无限循环
// ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
// var started bool
// var stopped atomic.Bool
// for {
// 	if !started {
// 		started = true
// 		go func() {
// 			for {
// 				select {
// 				case <-ctx.Done():
// 					fmt.Println("ctx done")
// 					stopped.Store(true)
// 					return
// 				}
// 			}
// 		}()
// 	}
// 	fmt.Println("main")
// 	if stopped.Load() {
// 		break
// 	}
// }

// 遍历数组
// var a [10]string
// a[0] = "Hello"
// for i := range a {
// 	fmt.Println("当前下标：", i, a[i])
// }
// for i, e := range a {
// 	fmt.Println("a[", i, "] = ", e)
// }

// 遍历切片
// s := make([]string, 8)
// s[0] = "Hello"
// for i := range s {
// 	fmt.Println("当前下标：", i, s[i])
// }
// for i, e := range s {
// 	fmt.Println("s[", i, "] = ", e)
// }

// m := make(map[string]string)
// m["b"] = "Hello, b"
// m["a"] = "Hello, a"
// m["c"] = "Hello, c"
// for i := range m {
// 	fmt.Println("当前key：", i, m[i])
// }
// for k, v := range m {
// 	fmt.Println("m[", k, "] = ", v)
//

// for i := 0; i < 5; i++ {
// 	if i == 3 {
// 		break
// 	}
// 	fmt.Println("第", i, "次循环")
// }
// switch i := 2; i {
// case 1:
// 	fmt.Println("进入case 1")
// 	if i == 1 {
// 		break
// 	}
// 	fmt.Println("i等于1")
// case 2:
// 	fmt.Println("i等于2")
// default:
// 	fmt.Println("default case")
// }

// select {
// case <-time.After(time.Second * 2):
// 	fmt.Println("过了2秒")
// case <-time.After(time.Second * 3):
// 	fmt.Println("经过了3秒")
// 	if true {
// 		break
// 	}
// 	fmt.Println("break 之后")
// }
// 不使用标记
// for i := 1; i <= 3; i++ {
// 	fmt.Printf("不使用标记,外部循环, i = %d\n", i) //123
// 	for j := 5; j <= 10; j++ {
// 		fmt.Printf("不使用标记,内部循环 j = %d\n", j) //5
// 		// break
// 	}
// }
// outter:
// 	for i := 1; i <= 3; i++ {
// 		fmt.Printf("使用标记,外部循环, i = %d\n", i)
// 		for j := 5; j <= 10; j++ {
// 			fmt.Printf("使用标记,内部循环 j = %d\n", j)
// 			break outter
// 		}
// 	}
// 中断for循环
// for i := 0; i < 5; i++ {
// 	if i == 3 {
// 		break
// 	}
// 	fmt.Println("第", i, "次循环")
// }

// 不使用标记

// 使用标记
// outter1:
// 	for i := 1; i <= 3; i++ {
// 		fmt.Printf("使用标记,外部循环, i = %d\n", i)
// 		for j := 5; j <= 10; j++ {
// 			fmt.Printf("使用标记,内部循环 j = %d\n", j)
// 			if j >= 7 {
// 				continue outter1
// 			}
// 			fmt.Println("不使用标记，内部循环，在continue之后执行")
// 		}
// 	}

// 	// 使用标记
// outter2:
// 	for i := 1; i <= 3; i++ {
// 		fmt.Printf("使用标记,外部循环, i = %d\n", i)
// 		for j := 5; j <= 10; j++ {
// 			fmt.Printf("使用标记,内部循环 j = %d\n", j)
// 			if j >= 7 {
// 				continue outter2
// 			}
// 		}
// 	}
// a := A{1}
// 把方法赋值给函数变量
// function1 = a.add //2

// // 声明一个闭包并直接执行
// // 此闭包返回值是另外一个闭包（带参闭包）
// returnFunc := func() func(int, string) (int, string) {
// 	fmt.Println("this is a anonymous function")
// 	return func(i int, s string) (int, string) {
// 		return i, s
// 	}
// }()

// // 执行returnFunc闭包并传递参数
// ret1, ret2 := returnFunc(1, "test")
// fmt.Println("call closure function, return1 = ", ret1, "; return2 = ", ret2)

// fmt.Println("a.i = ", a.i)
// fmt.Println("after call function1, a.i = ", function1(1))
// fmt.Println("a.i = ", a.i)
// 仅声明
// var a [5]int
// fmt.Println("a = ", a)

// var marr [2]map[string]string
// fmt.Println("marr = ", marr)
// // map的零值是nil，虽然打印出来是非空值，但真实的值是nil
// // marr[0]["test"] = "1"

// // 声明以及初始化
// var b [5]int = [5]int{1, 2, 3, 4, 5}
// fmt.Println("b = ", b)

// // 类型推导声明方式
// var c = []string{"c1", "c2", "c3", "c4", "c5"}
// fmt.Println("c = ", c)

// d := []int{3, 2, 1}
// fmt.Println("d = ", d)

// // 使用 ... 代替数组长度
// autoLen := [...]string{"auto1", "auto2", "auto3"}
// fmt.Println("autoLen = ", autoLen)

// // 声明时初始化指定下标的元素值
// positionInit := [5]string{1: "position1", 3: "position3"}
// fmt.Println("positionInit = ", positionInit)

// 初始化时，元素个数不能超过数组声明的长度
//overLen := [2]int{1, 2, 3}

// a := [5]int{5, 4, 3, 2, 1}

// // 方式1，使用下标读取数据
// element := a[2]
// fmt.Println("element = ", element) //3

// // 方式2，使用range遍历
// for i, v := range a {
// 	fmt.Println("index = ", i, "value = ", v)
// }

// for i := range a {
// 	fmt.Println("only index, index = ", i) //01234
// }

// // 读取数组长度
// fmt.Println("len(a) = ", len(a))
// // 使用下标，for循环遍历数组
// for i := 0; i < len(a); i++ {
// 	fmt.Println("use len(), index = ", i, "value = ", a[i])
// }

// a := [5]int{5, 4, 3, 2, 1}
// fmt.Println("before all, a = ", a) //54321
// for i := range carr {
// 	fmt.Printf("in main func, carr[%d] = %p, value = %v \n", i, &carr[i], *carr[i]) //01234、五个地址、678910
// }
// printFuncParamPointer(carr)

// receiveArray(a)                            //54321、5-5321
// fmt.Println("after receiveArray, a = ", a) //54321

// receiveArrayPointer(&a)                           //地址
// fmt.Println("after receiveArrayPointer, a = ", a) //地址

// 方式1，声明并初始化一个空的切片
// var s1 []int = []int{}

// // 方式2，类型推导，并初始化一个空的切片
// var s2 = []int{}

// // 方式3，与方式2等价
// s3 := []int{}

// // 方式4，与方式1、2、3 等价，可以在大括号中定义切片初始元素
// s4 := []int{1, 2, 3, 4}

// // 方式5，用make()函数创建切片，创建[]int类型的切片，指定切片初始长度为0
// s5 := make([]int, 0)

// // 方式6，用make()函数创建切片，创建[]int类型的切片，指定切片初始长度为2，指定容量参数4
// s6 := make([]int, 2, 4)

// // 方式7，引用一个数组，初始化切片
// a := [5]int{6, 5, 4, 3, 2}
// // 从数组下标2开始，直到数组的最后一个元素
// s7 := arr[2:]
// // 从数组下标1开始，直到数组下标3的元素，创建一个新的切片
// s8 := arr[1:3]
// // 从0到下标2的元素，创建一个新的切片
// s9 := arr[:2]

// src1 := []int{1, 2, 3}
// dst1 := make([]int, 4, 5)

// src2 := []int{1, 2, 3, 4, 5}
// dst2 := make([]int, 3, 3)

// fmt.Println("before copy, src1 = ", src1) //123
// fmt.Println("before copy, dst1 = ", dst1) //0000

// fmt.Println("before copy, src2 = ", src2) //12345
// fmt.Println("before copy, dst2 = ", dst2) //000

// copy(dst1, src1)
// copy(dst2, src2)

// fmt.Println("before copy, src1 = ", src1) //123
// fmt.Println("before copy, dst1 = ", dst1) //1230

// fmt.Println("before copy, src2 = ", src2) //12345
// fmt.Println("before copy, dst2 = ", dst2) //123
// fmt.Println(len(dst1), cap(dst1))
// var m1 map[string]string
// fmt.Println("m1 length:", len(m1)) //0

// m2 := make(map[string]string)
// fmt.Println("m2 length:", len(m2)) //0
// fmt.Println("m2 =", m2)            //map[]

// m3 := make(map[string]string, 10)
// fmt.Println("m3 length:", len(m3)) //0
// fmt.Println("m3 =", m3)            //map[]

// m4 := map[string]string{}
// fmt.Println("m4 length:", len(m4)) //0
// fmt.Println("m4 =", m4)            //map[]

// m5 := map[string]string{
// 	"key1": "value1",
// 	"key2": "value2",
// }
// fmt.Println("m5 length:", len(m5)) //2
// fmt.Println("m5 =", m5)            //map[key1:value1 key2:value2]

// m := make(map[string]int, 10)

// m["1"] = int(1)
// m["2"] = int(2)
// m["3"] = int(3)
// m["4"] = int(4)
// m["5"] = int(5)
// m["6"] = int(6)

// // 获取元素
// value1 := m["1"]
// fmt.Println("m[\"1\"] =", value1) //1

// value1, exist := m["1"]
// fmt.Println("m[\"1\"] =", value1, ", exist =", exist) //1、true

// valueUnexist, exist := m["10"]
// fmt.Println("m[\"10\"] =", valueUnexist, ", exist =", exist) //0、false

// // 修改值
// fmt.Println("before modify, m[\"2\"] =", m["2"]) //2
// m["2"] = 20
// fmt.Println("after modify, m[\"2\"] =", m["2"]) //20

// // 获取map的长度
// fmt.Println("before add, len(m) =", len(m)) //6
// m["10"] = 10
// fmt.Println("after add, len(m) =", len(m)) //7

// // 遍历map集合main
// for key, value := range m {
// 	fmt.Println("iterate map, m[", key, "] =", value) //1_1、2_20、3_3、4_4、5_5、6_6、10_10
// }

// // 使用内置函数删除指定的key
// _, exist_10 := m["10"]
// fmt.Println("before delete, exist 10: ", exist_10) //true
// delete(m, "10")
// _, exist_10 = m["10"]
// fmt.Println("after delete, exist 10: ", exist_10) //false

// // 在遍历时，删除map中的key
// for key := range m {
// 	fmt.Println("iterate map, will delete key:", key)
// 	delete(m, key)
// }
// fmt.Println("m = ", m) //map[]
// 	m := make(map[string]int)
// 	m["a"] = 1
// 	receiveMap(m)         //a,1
// 	fmt.Println("m =", m) //a,2  b,3
// }
// func receiveMap(param map[string]int) {
// 	fmt.Println("before modify, in receiveMap func: param[\"a\"] = ", param["a"])
// 	param["a"] = 2
// 	param["b"] = 3
// m := make(map[string]int)
// var lock sync.Mutex
// go func() {
// 	for {
// 		lock.Lock()
// 		m["a"]++
// 		lock.Unlock()
// 	}
// }()

// go func() {
// 	for {
// 		lock.Lock()
// 		m["a"]++
// 		fmt.Println(m["a"])
// 		lock.Unlock()
// 	}
// }()

// select {
// case <-time.After(time.Second * 5):
// 	fmt.Println("timeout, stopping")
// }
// str1 := "abc123"
// for index := range str1 {
// 	fmt.Printf("str1 -- index:%d, value:%d\n", index, str1[index]) // ASCII
// }

// str2 := "测试中文"
// for index := range str2 {
// 	fmt.Printf("str2 -- index:%d, value:%d\n", index, str2[index]) //utf-8
// }
// fmt.Printf("len(str2) = %d\n", len(str2)) //12

// runesFromStr2 := []rune(str2)
// bytesFromStr2 := []byte(str2)
// fmt.Printf("len(runesFromStr2) = %d\n", len(runesFromStr2)) //4
// fmt.Printf("len(bytesFromStr2) = %d\n", len(bytesFromStr2)) //12

// str1 := "a1中文"
// for index, value := range str1 {
// 	fmt.Printf("str1 -- index:%d, index value:%d\n", index, str1[index]) //byte
// 	fmt.Printf("str1 -- index:%d, range value:%d\n", index, value)       //runes utf-8
// }
// array := [...]int{1, 2, 3}
// slice := []int{4, 5, 6}

// // 方法1：只拿到数组的下标索引
// // for index := range array {
// // 	fmt.Printf("array -- index=%d value=%d \n", index, array[index]) // array元素0-1,1-2,2-3
// // }
// // for index := range slice {
// // 	fmt.Printf("slice -- index=%d value=%d \n", index, slice[index]) // slice元素0-4,1-5,2-6
// // }
// // fmt.Println()

// //方法2：同时拿到数组的下标索引和对应的值
// for index, value := range array {
// 	fmt.Printf("array -- index=%d index value=%d \n", index, array[index])
// 	fmt.Printf("array -- index=%d range value=%d \n", index, value)
// }
// for index, value := range slice {
// 	fmt.Printf("slice -- index=%d index value=%d \n", index, slice[index])
// 	fmt.Printf("slice -- index=%d range value=%d \n", index, value)
// }
// fmt.Println()
// 	ch := make(chan int, 10)

// 	go addData(ch)

// 	for a := range ch {
// 		fmt.Println(a)
// 		time.Sleep(1 * time.Second)
// 	}
// }
// func addData(ch chan int) {
// 	size := cap(ch)
// 	for i := 0; i < size; i++ {
// 		ch <- i

// 	}
// 	close(ch)

// a := map[string]int{
// 	"a": 1,
// 	"f": 2,
// 	"z": 3,
// 	"c": 4,
// }

// for key := range a {
// 	fmt.Printf("key=%s, value=%d\n", key, a[key])
// }

// for key, value := range a {
// 	fmt.Printf("key=%s, value=%d\n", key, value)
// }
// var i int32 = 17
// var b byte = 5
// var f float32

// // 数字类型可以直接强转

// fmt.Println("i 的值为: ", float32(i)) //17
// fmt.Println("b 的值为: ", float32(b)) //5
// f = float32(i) / float32(b)
// fmt.Printf("f 的值为: %f\n", f)

// // 当int32类型强转成byte时，高位被直接舍弃
// var i2 int32 = 256
// var b2 byte = byte(i2)
// fmt.Printf("b2 的值为: %d\n", b2) //0
// str := "hello, 123, 你好"
// var bytes []byte = []byte(str)
// var runes []rune = []rune(str)
// fmt.Printf("bytes 的值为: %v \n", bytes)
// fmt.Printf("runes 的值为: %v \n", runes)

// str2 := string(bytes)
// str3 := string(runes)
// fmt.Printf("str2 的值为: %v \n", str2)
// fmt.Printf("str3 的值为: %v \n", str3)
//var str string
// str := "123"
// num, err := strconv.Atoi(str)
// if err != nil {
// 	panic(err)
// }
// fmt.Printf("字符串转换为int: %d \n", num)
// str1 := strconv.Itoa(num)
// fmt.Printf("int转换为字符串: %s \n", str1)

// ui64, err := strconv.ParseUint(str, 10, 32)
// fmt.Printf("字符串转换为uint64: %d \n", ui64)

// str2 := strconv.FormatUint(ui64, 10)
// fmt.Printf("uint64转换为字符串: %s \n", str2)

// var i interface{} = 3
// a, ok := i.(int)
// if ok {
// 	fmt.Printf("'%d' is a int \n", a)
// } else {
// 	fmt.Println("conversion failed")
// }
// var i interface{} = "3"
// switch v := i.(type) {
// case int:
// 	fmt.Println("i is a int", v)
// case string:
// 	fmt.Println("i is a string", v)
// default:
// 	fmt.Println("i is unknown type", v)
// }
// 	var a Supplier = &DigitSupplier{value: 1}
// 	fmt.Println(a.Get())
// 	var a Supplier = &DigitSupplier{value: 1}
// 	b, ok := (a).(*DigitSupplier)
// 	fmt.Println(b, ok)
// }

// type Supplier interface {
// 	Get() string
// }

// type DigitSupplier struct {
// 	value int
// }

// func (i *DigitSupplier) Get() string {
// 	return fmt.Sprintf("%d", i.value)

// a := SameFieldA{
// 	name:  "a",
// 	value: 1,
// }

// b := SameFieldB(a)
// fmt.Printf("conver SameFieldA to SameFieldB, value is : %d \n", b.getValue())

// //只能结构体类型实例之间相互转换，指针不可以相互转换
// // 	var a Supplier = &DigitSupplier{value: 1}
// // 	b, ok := (a).(*DigitSupplier)
// var c interface{} = &a
// _, ok := c.(*SameFieldB)
// fmt.Printf("c is *SameFieldB: %v \n", ok)
//}

// type SameFieldA struct {
// 	name  string
// 	value int
// }

// type SameFieldB struct {
// 	name  string
// 	value int
// }

// func (s *SameFieldB) getValue() int {
// 	return s.value
// }

// func rightMoveAssignment(c, a int) {
// 	c >>= a // c = c >> a
// 	fmt.Println("c >>= a, c =", c)//1
// }

// func andAssignment(c, a int) {
// 	c &= a // c = c & a
// 	fmt.Println("c &= a, c =", c) // 1
// }

// func orAssignment(c, a int) {
// 	c |= a // c = c | a
// 	fmt.Println("c |= a, c =", c) // 3
// }

// func norAssignment(c, a int) {
// 	c ^= a // c = c ^ a
// 	fmt.Println("c ^= a, c =", c)//2
// }
// type A struct {
// 	i int
// }

// func (a *A) add(v int) int {
// 	a.i += v
// 	return a.i
// }

// 声明函数变量
// var function1 func(int) int

// 声明闭包
// var squart2 func(int) int = func(p int) int {
// 	p *= p
// 	return p
// }
// func receiveArray(param [5]int) {
// 	fmt.Println("in receiveArray func, before modify, param = ", param)
// 	param[1] = -5
// 	fmt.Println("in receiveArray func, after modify, param = ", param)
// }

// func receiveArrayPointer(param *[5]int) {
// 	fmt.Println("in receiveArrayPointer func, before modify, param = ", param)
// 	param[1] = -5
// 	fmt.Println("in receiveArrayPointer func, after modify, param = ", param)
// }

// func printFuncParamPointer(param [5]*Custom) {
// 	for i := range param {
// 		fmt.Printf("in printFuncParamPointer func, param[%d] = %p, value = %v \n", i, &param[i], *param[i])
// 	}

// }

// type Custom struct {
// 	i int
// }

// var carr [5]*Custom = [5]*Custom{
// 	{6},
// 	{7},
// 	{8},
// 	{9},
// 	{10},
// }
