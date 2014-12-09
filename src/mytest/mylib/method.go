package mylib

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func TestArea() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{24}

	fmt.Println("Area of r1 is :", r1.Area())
	fmt.Println("Area of r2 is :", r2.Area())
	fmt.Println("Area of r3 is :", c1.Area())
	fmt.Println("Area of r4 is :", c2.Area())

}

/*******测试自定义类型**********/
const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box //a slice boxes

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

func (b1 BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range b1 {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (b1 BoxList) PaintItBlack() {
	for i, _ := range b1 {
		b1[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	string := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return string[c]
}

func TestBox() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 100, 13, BLUE},
		Box{10, 30, 3, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set \n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm3")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestColor().String())

	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())
	fmt.Println("Obviously,now ,the biggest one is", boxes.BiggestColor().String())
}

/*********** method继承 ************/

type HumanNew struct {
	name  string
	age   int
	phone string
}

type StudentNew struct {
	HumanNew
	school string
}

type EmployeeNew struct {
	HumanNew
	company string
}

func (h *HumanNew) SayHi() {
	fmt.Printf("Hi,I am %s you cann call me on %s \n", h.name, h.phone)
}

//Employee的method 重写Human的method
func (e *EmployeeNew) SayHi() {
	fmt.Printf("Hi,I am %s,I work at %s.Call me on %s\n", e.name, e.company, e.phone)
}
func TestExtends() {
	mark := StudentNew{HumanNew{"Mark", 25, "222-222-YYYYY"}, "MIT"}
	sam := EmployeeNew{HumanNew{"Sam", 45, "111-888-xxxxx"}, "Golang Inc"}
	mark.SayHi()
	sam.SayHi()
}
