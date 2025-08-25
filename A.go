package main

// type A struct {
// 	a string
// }

// func (a A) string() string {
// 	return a.a
// }

// func (a A) stringA() string {
// 	return a.a
// }

// func (a A) setA(v string) {
// 	a.a = v
// }

// func (a *A) stringPA() string {
// 	return a.a
// }

// func (a *A) setPA(v string) {
// 	a.a = v
// }

// type B struct {
// 	A
// 	b string
// }

// func (b B) string() string {
// 	return b.b
// }

// func (b B) stringB() string {
// 	return b.b
// }

// type C struct {
// 	B
// 	a string
// 	b string
// 	c string
// 	d []byte
// }

// func (c C) string() string {
// 	return c.c
// }

// func (c C) modityD() {
// 	c.d[2] = 3
// }

// func callStructMethod() {
// 	var a A
// 	a = A{
// 		a: "a",
// 	}
// 	a.string()
// }

// func NewC() C {
// 	return C{
// 		B: B{
// 			A: A{
// 				a: "ba",
// 			},
// 			b: "b",
// 		},
// 		a: "ca",
// 		b: "cb",
// 		c: "c",
// 		d: []byte{1, 2, 3},
// 	}
// }
// type Gender string

// const (
// 	Male   Gender = "Male"
// 	Female Gender = "Female"
// )

// func (g Gender) String() string {
// 	switch g {
// 	case Male:
// 		return "Male"
// 	case Female:
// 		return "Female"
// 	default:
// 		return "Unknown"
// 	}
// }

// func (g Gender) IsMale() bool {
// 	return g == Male
// }
