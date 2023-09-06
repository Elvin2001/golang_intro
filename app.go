package main

import "fmt"

type Person struct {
	NAME   string
	AGE    int
	SEX    string
	SALARY int
}

type Car struct {
	MODEL              string
	NEW                bool
	YEAR_OF_PRODUCTION int
	PRICE              int
}

func main() {
	//a := "hello"
	//b := 12
	//
	//c := &a
	//d := &b
	//
	//fmt.Println(c)
	//fmt.Println(d)
	//fmt.Println(*c)
	//fmt.Println(*d)

	workingWithCar()
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
