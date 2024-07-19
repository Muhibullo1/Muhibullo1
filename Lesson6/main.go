package main

import "fmt"

func main() {

	/*var x int= 10
	var ptr *int 
	fmt.Println(ptr)

	ptr = &x

	fmt.Println("Значение х:", x)
	fmt.Println("Адрес х:", &x)
	fmt.Println("Значение ptr:", ptr)
	fmt.Println("Значение на которое указывает ptr:", *ptr)

var num1 int 
fmt.Scan(&num1)
updateValue(&num1)

var x, y int 
fmt.Scan(&x, &y)
swap(&x, &y)
fmt.Println(x, y)
*/
var math, physics, chemistry int
mathPtr := &math
physicsPtr := &physics
chemistryPtr := &chemistry
addGrade(mathPtr, 85)
addGrade(physicsPtr, 90)
addGrade(chemistryPtr, 78)
fmt.Println("Средний балл по предметам: \n", printAverageGrade(mathPtr, physicsPtr, chemistryPtr))
}
/*
func updateValue(a *int)  {
	*a = *a + 10
	fmt.Println(*a)
}

func swap(b, c *int) {
	*b, *c = *c, *b

}

func vote(candidatePtr *string, votesPtr *int) string {
   *votesPtr++

}
func winner(candidate1VotesPtr *int , candidate2VotesPtr *int) string {
	if *candidate1VotesPtr > *candidate2VotesPtr {
		return "Кандидат 1"
	} else if *candidate2VotesPtr > *candidate1VotesPtr {
		return "Кандидат 2"
	} else {
		return "Ничья"
	}
}
*/

func addGrade(mathPtr *int, grade int) {
	*mathPtr = grade
}

func printAverageGrade(mathPtr *int , physicsPtr *int, chemistryPtr *int) float64 {
	sum := *mathPtr + *physicsPtr + *chemistryPtr
	return float64(sum) / 3.0
}


