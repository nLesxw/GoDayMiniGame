package main

import "fmt"

type people struct {
}

func (p *people) ShowA() {
	fmt.Println("ShowA")
	p.ShowB()
}

func (p *people) ShowB(){
	fmt.Println("ShowB")
}

type teacher struct {
	people
}

func (t *teacher) ShowB() {
	fmt.Println("teacher ShowB")
}

func main() {
	t := &teacher{}
	t.ShowA()
}