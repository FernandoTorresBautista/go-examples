package main

import (
	"fmt"
	"strings"
)

func main() {
	// SECTION 1
	fmt.Println("hello word")

	// SECTION 2
	// int
	testInt()
	fmt.Println("--")
	// string
	testString()
	fmt.Println("--")
	// array
	todo()
	fmt.Println("--")
	// pointers
	pointers()
	fmt.Println("--")
	// structures, abstract data type
	testStructure()
	fmt.Println("--")
	// interfaces 1
	testInterface()
	fmt.Println("--")
	// interfaces 2
	testInterface2()
	fmt.Println("--")
	// Control structure
	testControlstructures() // if for
	fmt.Println("--")
	// Control structure 2
	testControlstructures2() // switch

	// SECTION 3
}

func testControlstructures2() {
	// continue, break, switch case
	fmt.Println("continue")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("break")
	for i := 0; i < 10; i++ {
		if i > 2 {
			break
		}
		fmt.Println(i)
	}
	//

	day := "Fri"

	switch day {
	case "Fry":
		fmt.Println("Fry day")
	case "Mon", "Tue", "Wed":
		fmt.Println("boring")
	default:
		fmt.Println("boring")
	}

	switch {
	case day == "Fri":
		fmt.Println("Second case")
		break
	}
}

//
func testControlstructures() {
	fmt.Println("Control structures in go")
	// if else,  for, switch case, break continue
	f := true
	flag := &f
	if flag == nil {
		fmt.Println("nil")
	} else if *flag {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	for i := 0; i < 2; i++ {
		fmt.Println("i: ", i)
	}
	/*for { // loop infinito
	  }*/

	arr := []string{"a", "b", "c"}

	for i, v := range arr {
		fmt.Println("i, v: ", i, v)
	}
	mmap := make(map[string]interface{})
	//mmap := map[string]interface{}{}
	mmap["name"] = "name"
	mmap["age"] = 20
	for i, v := range mmap {
		fmt.Println(i, v)
	}

	//

}

// interfaces 2
func testInterface2() {
	Anything(1)
	Anything("string")
	Anything(10.1)
	Anything(struct{}{})

	mmap := make(map[string]interface{})
	mmap["name"] = "name"
	mmap["age"] = 15
	fmt.Println(mmap)
}

//interface{} una interface sin  definir es de cualqueir tipo
func Anything(anything interface{}) {
	fmt.Println(anything)
}

// interfaces 1
func testInterface() {
	l := Lambo{"gallardo"}
	c := Chevy{"c396"}
	l.Drive()
	c.Drive()
}
func NewModel(arg string) CarI {
	return &Lambo{arg}
}

// CarI es la interface que define los métodos de
type CarI interface {
	Drive()
	Stop()
}
type Lambo struct {
	LamboModel string
}
type Chevy struct {
	ChevyModel string
}

func (l *Lambo) Drive() {
	fmt.Println("on move lambo")
	fmt.Println(l.LamboModel)
}
func (l *Chevy) Drive() {
	fmt.Println("on move chevy")
	fmt.Println(l.ChevyModel)
}
func (l *Lambo) Stop() {
	fmt.Println("stop lambo")
	fmt.Println(l.LamboModel)
}
func (l *Chevy) Stop() {
	fmt.Println("stop chevy")
	fmt.Println(l.ChevyModel)
}

func testStructure() {
	car := Car{"name", 18, 1}
	car2 := Car{
		Name:    "chevy",
		Age:     18,
		ModelNo: 1,
	}

	fmt.Println(car)
	car2.Print()
	fmt.Println(car2.GetName())
}

// method Print for the struct Car
func (c Car) Print() {
	fmt.Print(c)
}

func (c Car) GetName() string {
	return c.Name
}

// structire -> encapsulation data
type Car struct {
	Name    string
	Age     int
	ModelNo int
}

func swap(m1, m2 *int) {
	var temp int
	temp = *m2
	*m2 = *m1
	*m1 = temp
}

func pointers() {
	m1 := 2
	ptr := &m1
	fmt.Println(ptr)  // direccion of memory
	fmt.Println(*ptr) // value of the memory
	m2 := 3
	swap(&m1, &m2)
	fmt.Println(m1, m2)
}

func testString() {
	var str1 string
	str2 := "new string"
	str1 = str2 + " new word"
	fmt.Println(str1)
	fmt.Println(strings.Contains(str1, str2))
	fmt.Println(strings.ReplaceAll(str1, "new", "old"))
	fmt.Println(strings.Split(str2, " "))
}

func testInt() {
	var varint int
	var varint32 int32
	var varint64 int64
	varint = 10
	varint32 = 10
	varint64 = 10
	fmt.Println(varint, varint32, varint64)
	intother := 10
	fmt.Println(intother)
}

func todo() {
	var arr []int
	arr = []int{1, 2, 3, 4}
	arr2 := []string{"hi", "my", "name"}
	fmt.Println(arr, arr2)
	arr2 = append(arr2, "is", "example")
	fmt.Println(arr2)
}
