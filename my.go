package main

import (
	"fmt"
	
)
func main(){
var value1 int
var value2 int 
var result string
for i:=0;;i++{
fmt.Println("Введите первое число") 
fmt.Scan(&value1)
fmt.Println("Введите второе число")
fmt.Scan(&value2)
fmt.Println("Введите действие +,-,*,/,")
fmt.Scan(&result)
switch result {
case  "+":
fmt.Println(value1+value2)
case  "-":
	fmt.Println(value1-value2)
case  "*":
      fmt.Println(value1*value2)
	
case  "/":
	if value1==0||value2==0 {
		fmt.Println("На ноль делить нельзя")
		
	} 
	if  value1>0&&value2>0{
		fmt.Println(value1/value2)
	}
}
}

}	


