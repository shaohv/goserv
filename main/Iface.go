package main

import "fmt"

// Idog ..
type Idog interface {
	Wang()
}

// Dog ..
type Dog struct {
	Name string
}

func (d *Dog) Wang() {
	fmt.Println(" ", d.Name, "wang ...")
}

func Eat(dog Idog) {
	d := dog.(*Dog)
	fmt.Println(" ", d.Name, "eat ...")
}

func main() {
	d := &Dog{
		Name: "XiaoMing",
	}

	Eat(d)
}
