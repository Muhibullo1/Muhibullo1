package main

import (
	"fmt"
)

func main() {
/*	num := 15

	if num > 9 && num < 100{
		fmt.Println("Двузначное")
	}

	n := 7 
	if n %2==1 {
		fmt.Println("Число нечетное")
	} else {
		fmt.Println("Число четное")
	}

	m := 78
	if m >=0 && m <10{
		fmt.Println("Число однозначние")
	} else if m >= 10 && m < 100 {
		fmt.Println("Число двузначное")
	} else if m >=100 && m <1000 {
		fmt.Println("Число трехзначное")
	} else {
		fmt.Println("Число слишком большое")
	}

	nn := 1
	
	switch nn {
	case 1:
		fmt.Println("Positive")
		fallthrough
	case -1:
		fmt.Println("Negative")
		fallthrough
	case 0:
		fmt.Println("Unsinged")
	}
/*
	  var num1, num2 = 7, 8
	if num1 > num2 {
		fmt.Println("ПЕрвое число больше")
	} else {
		fmt.Println("Второе чтсло больше")	
	}

	a := 54
	b := 989
	c := 498
  if a > b && b >c {
	fmt.Println("a+b=", a+b)
  } else if b > a && a > c {
	fmt.Println("b+a=", b+a)
  } else if c > a && a > b {
	fmt.Println("c+a=", c+a )
  } else if a > c && c > b {
fmt.Println("a+c=", a+c)
  } else if b > c && c > a {
	fmt.Println("b+c=", b+c)
  } else {
	fmt.Println("c+b=", c+b)
  } 


  var god int= 2034
  if god %4==0  {
   
	fmt.Println("Год високосный")
  } else {
    fmt.Print("Год невисокосный")
  }
*/
var nomer int 
fmt.Scan(&nomer)
switch nomer {
case 1 :
	fmt.Println("В январье 31 дней")
case 2:
	fmt.Print("В феврале 28 или 29 дней")
case 3:
	fmt.Println("В марте 31 дней")
case 4 :
	fmt.Println("В апреле 31 дней")
case 5:
	fmt.Println("В майе 30 дней")
case 6:
	fmt.Println("В июне 30 дней")
case 7:
	fmt.Println("В июле 31 дней")
case 8:
	fmt.Println("В августе 31 дней")
case 9:
	fmt.Println("В сентябре 30 дней")	
case 10:
	fmt.Println("В октябрье 31 дней")
case 11:
	fmt.Print("В ноябрье 30 дней")
case 12:
	fmt.Println("В декабре 31 дней")

	}
}

