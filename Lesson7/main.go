package main

import "fmt"
/*
type Age int

func Vozrast (Age Age) string {
	if Age >= 18 {
		return "Человек совершеннолетный"
	} else {
		return "Нет"
	}
}

type Number int 

func Chetnost (Number Number) string {
if Number %2 == 0 {
	return "Число четное"
} else {
	return "Число нечетное"
}
}

type Score int 

func Grade (grd Score) string {
	if grd >= 0 && grd <= 100 {
		return "Находиться"
	} else {
		return "Ненааходиться"
	}
}


 
type Comparator func(int, int) bool 

func equal( a int, b int) bool {
	return a == b
}

func greater(a int, b int) bool {
	return a > b
}

func less(a int, b int) bool {
	return a < b
}

type UnaryOperation func(int) int

func Double(a int) int {
	return a*2
}

func Triple(a int) int {
	return a*3
}
*/
type Count int 

func Otchet(c Count)  {
	for i := c; i >= 0; i--{
		fmt.Println(i)
	}
}


func main() {
/*var MyAge int

fmt.Scan(&MyAge)

fmt.Println(Vozrast(Age(MyAge)))

var Number1 int

fmt.Scan(&Number1)

fmt.Println(Chetnost(Number(Number1)))
*/
var n Count
fmt.Scan(&n)
Otchet(n)

}


