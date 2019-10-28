
package main

import "fmt"
type person struct {
	name   string
	age    int
	gender string
}


type dove interface {
	gugugu()
}
type Argumentative  interface {
	buguoruci()
}
type Really  interface {
	wpbuhuichiniyikoudongxi()
}
type repeater interface {
	repeat(string)
}
func (p *person) repeat(word string) {
	fmt.Println(word)
}

func (p *person) gugugu() {
	fmt.Println(p.name, "又鸽了")
}
func (p *person) buguoruci() {
	fmt.Println(p.name, "酸了")
}
func (p *person) wobuhuichiniyikoudongxi() {
	fmt.Println(p.name, "真香")
}
func main() {

	p := &person{
		name:   "lcm",
		age:    18,
		gender: "male",
	}
	p.gugugu()
p.wobuhuichiniyikoudongxi()
	p.buguoruci()
	var r repeater

	r = p
	r.repeat("helloworld")

}