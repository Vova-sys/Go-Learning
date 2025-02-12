package main

// import (
// 	"fmt"
// )

// func main() {
// 	s, i, f := "Happiness", 42, 42.42
// 	fmt.Printf("%v is of type %T\n", i, i)
// 	fmt.Printf("%v is of type %T\n", f, f)
// 	fmt.Printf("%v is of type %T\n", s, s)
// }

/*
var for zero value
short declaration operator
multiple initializations
var when specificity is required
blank identifier
*/

// func main() {
// 	var y int
// 	fmt.Println(y)

// 	z := 42
// 	fmt.Println(z)

// 	a, b := 43, "Happiness"
// 	fmt.Println(a, b)

// 	var c float32 = 42.42
// 	fmt.Printf("%v is of this type %T\n", c, c)

// 	e, f, _ := 45, 46, 47
// 	fmt.Println(e, f)

// }

// func main() {
// 	fmt.Printf("%d \t %b\n", 1, 1)
// 	fmt.Printf("%d \t %b\n", 1<<1, 1<<1)
// 	fmt.Printf("%d \t %b\n", 1<<2, 1<<2)
// 	fmt.Printf("%d \t %b\n", 1<<3, 1<<3)
// 	fmt.Printf("%d \t %b\n", 1<<4, 1<<4)
// 	fmt.Printf("%d \t %b\n", 1<<5, 1<<5)
// }

// const (
// 	Flag1 = 1 << iota // 1 << 0 = 1
// 	Flag2             // 1 << 1 = 2
// 	Flag3             // 1 << 2 = 4
// 	Flag4             // 1 << 3 = 8
// 	Flag5             // 1 << 4 = 16
// )

// func main() {
// 	fmt.Println(Flag1, Flag2, Flag3, Flag4)
// }

// func split(sum int) (x, y int) {
// 	x = sum * 4 / 2
// 	y = sum - x
// 	return
// }

// func main() {
// 	fmt.Println(split(10))
// }

// func main() {
// 	var i int
// 	j := i // j is an int

// 	fmt.Printf("i, j %T", j)
// }

// func add(x, y int) int {
// 	return x + y
// }

// func sayHello() {
// 	fmt.Println("Hello")
// }

// func main() {
// 	var a, b int
// 	fmt.Scan(&a, &b)
// 	fmt.Println(add(a, b))
// 	sayHello()
// }

// func main() {
// 	fmt.Printf("Now you have %g problems.\n", math.Sqrt(9))
// }

// func main() {
// 	b, c, d, e := 1, 2, 3, 4
// 	fmt.Println(b, c, d, e)
// }

// func main() {
// 	fmt.Println("😍")
// 	fmt.Println(`By Serena Alagappan
// This morning, the wind whipped the trees
// so hard one sycamore head snapped,
// its leaves strewn on my…`)
// }

// func main() {
// 	const name, age = "Kim", 22
// 	fmt.Printf("%s is %d years old.  \t and the type is %T and %T", name, age, name, age)
// }

// func main() {
// 	var age uint8 = 23 // age := 23 Автоматическое определение типа когда переменная инициализируется
// 	var number float64 = 233.333
// 	fmt.Println(age)
// 	fmt.Println(number)
// 	name := "Vova"
// 	fmt.Println(name)
// }

// ВВОД ДАННЫХ С КОНСОЛИ

// func main() {
// 	var name string
// 	var age uint8
// 	fmt.Println("What is your name?")
// 	fmt.Scan((&name))
// 	fmt.Println("Hello", name)
// 	fmt.Println("How old are you?")
// 	fmt.Scan(&age)
// 	fmt.Println("You are " + fmt.Sprint(age) + " years old") // fmt.Sprint(age) - преобразует число в строку
// }

// ПРЕОБРАЗОВАНИЕ ТИПОВ

//	func main() {
//		var a int8 = 2
//		var b float64 = float64(a)
//	}

// УСЛОВНЫЕ ОПЕРАТОРЫ
// func main() {
// 	num := 3
// 	if num > 0 {
// 		fmt.Println("Number is graeter than 0")
// 	} else if num < 0 {
// 		fmt.Println("Number is less than 0")
// 	} else if num == 3 {
// 		fmt.Println("Number is equal to 3")
// 	}

// 	fmt.Scan(&num)
// }

// ЦИКЛЫ

// func main() {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 	}
// }

// func main() {  //Бесконечный цикл
// 	for {
// 		fmt.Println("Hello")
// 	}
// }

// func main() {
// 	for i := 1; i <= 100; i++ {
// 		if i%2 == 0 {
// 			fmt.Println(i)
// 		}
// 	}
// }

// func main() {
// 	for i := 1; i <= 100; i++ {
// 		if i%2 == 0 {
// 			fmt.Println(i)
// 		}
// 	}
// }

// func main() {
// 	for i := 1; i <= 100; i++ {
// 		if i%2 != 0 {
// 			continue //Текущая итерация прерывается и начинается следующая
// 		}
// 		fmt.Println(i)
// 	}
// }

// func main() {
// 	for i := 1; i <= 100; i++ {
// 		if i == 50 {
// 			break //Прерывание цикла
// 		}

// 		fmt.Println(i)
// 	}
// }

// func main() {
// 	nums := []int{1, 2, 3, 4, 5} //Срез

// 	for i := 0; i < len(nums); i++ { // Итератор меньше длины среза
// 		fmt.Println(nums[i])
// 	}
// }

// func main() {
// 	nums := []int{1, 2, 3, 4, 5}

// 	for index, element := range nums {
// 		fmt.Printf("Index: %d Element: %d\n", index, element)
// 	}
// }

// func main() {
// 	nums := []int{1, 2, 3, 4, 5}

// 	for _, element := range nums { // _ - игнорирование индекса
// 		fmt.Printf("Element: %d\n", element)
// 	}
// }

// Двумерные массивы

// func main() {
// 	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

// 	for _, row := range matrix {
// 		for _, element := range row {
// 			fmt.Printf("%d ", element)
// 		}
// 		fmt.Println()
// 	}

// }

// SWITCH (Если выполнилось первое условие, то switch прерывается; default - если ни одно условие не выполняется; fallthrough - переход к следующему условию)

// func main() {
// 	name := "Andrew"

// 	switch name {
// 	case "Kris":
// 		fmt.Println("Hello Kris")
// 	case "Kate":
// 		fmt.Println("Hello Kate")
// 	case "Vova":
// 		fmt.Println("Hello Vova")

// 	default:
// 		fmt.Println("Hello Stranger")
// 	}
// }

// func main() {
// 	number := 10

// 	switch {
// 	case number > 1:
// 		fmt.Println("Number is greater than 1")
// 		fallthrough
// 	case number < 11:
// 		fmt.Println("Number is less than 11")
// 	}
// }
