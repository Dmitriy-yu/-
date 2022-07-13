package main

import (
	"fmt"
)
func main(){
var value1 int
var value2 int 
var result string
fmt.Println("Введите первое число") 
fmt.Scan(&value1)
fmt.Println("Введите второе число")
fmt.Scan(&value2)
fmt.Println("Введите действие +,-,*,/,")
fmt.Scan(&result)
switch result {
case "+":
fmt.Println(value1+value2)
}

}	


