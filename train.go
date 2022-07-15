package main
import "fmt"
type Person struct{
	num1 float64
	num2 float64
}
func (p Person)plus()float64{
	return p.num1+p.num2
}
func (p Person)minus()float64{
	return p.num1-p.num2
}
func (p Person)divide()float64{
	return p.num1/p.num2
}
func (p Person)multiply()float64{
	return p.num1*p.num2
}
func main(){
var a float64
var b float64
var c string
for i:=0;;i++{
fmt.Println("Введите первое число")
fmt.Scan(&a)
fmt.Println("Введите второе число")
fmt.Scan(&b)
fmt.Println("Введите действие+,-,*,/")
fmt.Scan(&c)
result:=Person{a,b}
if c=="+" {
	fmt.Println("Результат равен",result.plus())
}else if c=="-"{
	fmt.Println("Результат равен",result.minus())
}else if c=="/"{
	if a==0||b==0 {
		fmt.Println("На ноль делить нельзя")
	}
	fmt.Println("Результат равен",float64(result.divide()))
    }else if c=="*"{
	fmt.Println("Результат равен",result.multiply())
}
}
}