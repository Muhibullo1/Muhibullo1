package main

import (
	"fmt"
	
)

func main() {
	//var arr [5]int

  // arr2 := [3]string{"a", "b", "c"}
 
  // arr3 := [...]int{1, 2, 3, 4, 5}

//   slice1 := []int{1, 2, 3}
    
//    slice2 := []int{4,5}

//    slice3 := append(slice2, slice1...)

//    fmt.Println(len(slice3), cap(slice3))

/*
// 1 promlem
 arr := [7]int{1, 2, 3, 4, 5, 6, 7}

fmt.Println(arr)

// 2 Promlem
arr1 := [5]string{"a", "b", "c", "d", "e"}

fmt.Println(arr1)

// 3 Problem
arr2 := [4]bool{true, false, true, false}

fmt.Println(arr2)

// 4 Promlem
var arr3 [10]int

fmt.Println(arr3)

// 5 Promlem
var arr4 [4]bool

fmt.Println(arr4[1], arr4[3])

// 6 Promlem

var arr5 [3]bool

arr5[0] = false

fmt.Println(arr5)
*/

// 7 Promlem 

 
/*
a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11}

idx := 0

for i, v := range a {

	if a[0] < v && v <= a[9]{
		idx = i
	}
}
  
fmt.Println(idx)
*/

b := [10]int{3, 32, 43, 5, 45, 55, 65, 34, 67, 89}

for i := 0 ;i < 10 ; i++{
	if b[i]- b[i] == b[i]-b[i-1]{
		fmt.Println(b[i+1]-b[i])
	} else {
		fmt.Println("NO")
	}
}


}