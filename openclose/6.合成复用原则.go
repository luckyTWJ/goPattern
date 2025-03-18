package main

import "fmt"

type Cat struct {
	name string
}

func (cat *Cat) Eat() {

}

func (cat *Cat) Sleep() {
	fmt.Println("cat is sleeping...")
}

// 组合的方式实现
type CatC struct {
	Cat *Cat
}

func (catC *CatC) Sleep() {
	catC.Cat.Sleep()
}

func (c *CatC) Eat() {
	fmt.Println("cat is eating...")
	c.Cat.Eat()
}

func main() {
	cat := &Cat{}

	catC := &CatC{
		Cat: cat,
	}
	catC.Sleep()
}
