package main

import "fmt"

type Person struct {
	NAME   string
	AGE    int
	SEX    string
	SALARY int
}

type Book struct {
	TITLE              string
	AUTHOR             string
	COUNT_OF_LISTS     int
	YEAR_OF_PRODUCTION int
}

type Author struct {
	NAME string
	AGE  int
	SEX  string
	BOOK Book
}

type Car struct {
	MODEL              string
	NEW                bool
	YEAR_OF_PRODUCTION int
	PRICE              int
}

func main() {

	// arrays
	var arr [4]string
	arr[0] = "Elvin"
	arr[1] = "Eduard"
	arr[2] = "Evgeniy"
	arr[3] = "Sonya"

	elements := arr[0:2]
	fmt.Println(elements)
	fmt.Println(arr)

	strings := []string{"el", "wl", "wq", "xl", "al", "gl"}

	if strings == nil {
		fmt.Println("Nil")
	} else {
		fmt.Println("Not nil")
	}
	str := strings[:]
	fmt.Println(str)

	names := [...]string{"Elvin", "Evgeniy", "Denis", "Alexander", "Vladimir", "Sofia", "Polina"}

	for i := 0; i < len(names); i++ {
		fmt.Println("Current elem: " + names[i])
	}

	// slices
	colors := []string{
		"green",
		"yellow",
		"black",
		"blue",
		"white",
	}

	colors = append(colors, "violet")

	for i := range colors {
		fmt.Println(colors[i])
	}

	createSlice := make([]string, 3)
	createSlice[0] = "a"
	createSlice[1] = "b"
	createSlice[2] = "c"
	a := "hello"
	b := 12

	c := &a
	d := &b

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(*c)
	fmt.Println(*d)
	book := Book{"The green mile", "Steven King", 245, 2004}
	author := Author{"Steven King", 57, "M", book}
	fmt.Println(book.TITLE)
	fmt.Println(book.AUTHOR)
	book.method()
	workingWithCar()
	fmt.Println(author)
}

func workingWithPerson() {
	myPerson := Person{
		NAME:   "Elvin",
		AGE:    22,
		SEX:    "M",
		SALARY: 1000,
	}
	fmt.Println("Name of person: " + myPerson.NAME)
	fmt.Println(fmt.Sprintf("Age of person: %d", myPerson.AGE))
	fmt.Println(fmt.Sprintf("Salary of person: %d", myPerson.SALARY))
}

func workingWithCar() {
	car := Car{"Bugatti Veyron", true, 2019, 24559}
	getCarInfo(car)
}

func getCarInfo(car Car) {
	fmt.Println(fmt.Sprintf("Model: %s, new: %t, year: %d, price: %d", car.MODEL, car.NEW, car.YEAR_OF_PRODUCTION, car.PRICE))
}

func (book Book) method() {
	fmt.Println(book.TITLE)
	fmt.Println(book.AUTHOR)
}
