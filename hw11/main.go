package main 

import (
"fmt"
"strings"
)

func main() {
  // 1 Zadacha
	s1 := "Hello, "
	s2 := "world!"
	result := concatenateStrings(s1, s2)
	fmt.Println(result) // Output: Hello, world!

	// 2 Zadacha
	str := "Hello, world!"
	length := getStringLength(str)
	fmt.Println(length) // Output: 13

	// 3 Zadachaa
	substr1 := "world"
	substr2 := "Go"

	fmt.Printf("Does the string contain '%s'? %v\n", substr1, containsSubstring(str, substr1)) // Output: true
	fmt.Printf("Does the string contain '%s'? %v\n", substr2, containsSubstring(str, substr2)) // Output: false

	// 4 Zadacha
	index1 := indexOfSubstring(str, substr1)
	index2 := indexOfSubstring(str, substr2)

	fmt.Printf("Index of '%s' in '%s': %d\n", substr1, str, index1) // Output: 7
	fmt.Printf("Index of '%s' in '%s': %d\n", substr2, str, index2) // Output: -1

	// 5 Zadacha

	originalString := "Hello world! Welcome to the world of Go."
    oldSubstring := "world"
    newSubstring := "universe"

    reslt := ReplaceSubstring(originalString, oldSubstring, newSubstring)
    fmt.Println(reslt) // Вывод: Hello universe! Welcome to the universe of Go.
}


//1.	 Конкатенация строк
 
func concatenateStrings(str1, str2 string) string {
	return str1 + str2
}

//2.	 Длина строки

func getStringLength(s string) int {
	return len(s)
}

//3.	 Проверка подстроки

func containsSubstring(str, substr string) bool {
	return strings.Contains(str, substr)
}


//4.	 Поиск подстроки

func indexOfSubstring(str, substr string) int {
	return strings.Index(str, substr)
}

//5.	Замена подстроки

func ReplaceSubstring(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}