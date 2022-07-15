package main
import "fmt"
type Numbers struct{
	num1 int
	num2 int
}
func(n Numbers) plus()int{
	return n.num1 + n.num2
}
func(n Numbers)minus()int{
	return num1+num2
}
fuc main(){
var a int
var b int
var result string
fmt.Println("Введите первое число")
fmt.Scan(&a)
fmt.Println("ВВедите второе число")
fmt.Scan(&b)
fmt.Println("Введите действие")
fmt.Scan(&result)
numbers:=Numbers{a,b}
switch result{
case "+":
	fmt.Printf("Сумма=:%d",numbers.plus())
case "-":
	fmt.Printf("Сумма=:%d",minus())
}
}