package main
import "fmt"
type Person struct{
	num1 int
	num2 int
}
func (p Person)plus()int{
	return p.num1+p.num2
}
func (p Person)minus()int{
	return p.num1-p.num2
}
func (p Person)divide()int{
	return p.num1/p.num2
}
func (p Person)multiply()int{
	return p.num1*p.num2
}
func main(){
var a int
var b int
var c string
for i:=0;;i++{
fmt.Println("Введите первое число")
fmt.Scan(&a)
fmt.Println("Введите второе число")
fmt.Scan(&b)
fmt.Println("Введите действие")
fmt.Scan(&c)
result:=Person{a,b}
if c=="+" {
	fmt.Println("Результат равен",result.plus())
}else if c=="-"{
	fmt.Println("Результат равен",result.minus())
}else if c=="/"{
	fmt.Println("Результат равен",result.divide())
}else if c=="*"{
	fmt.Println("Результат равен",result.multiply())
}
}
}