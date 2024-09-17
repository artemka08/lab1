package main

import (
	"fmt"
	"time"
)

func SumAndDifference(a, b float64) (summa float64, diff float64) {
	summa = a + b
	diff = a - b
	return
}

func main() {
	fmt.Println("Задание 1")
	// Получаем текущее системное время
	currentTime := time.Now()

	// Форматируем дату и время в удобочитаемый формат
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	// Выводим дату и время на экран
	fmt.Println("Текущая дата и время:", formattedTime)

	fmt.Println("Задание 2")

	var a int = 5
	var b float32 = 5.32
	var c string = "Hello, world!!!"
	var d bool = true

	fmt.Println("Вывод переменной с тиом int:", a)
	fmt.Println("Вывод переменной с тиом float32:", b)
	fmt.Println("Вывод переменной с тиом string:", c)
	fmt.Println("Вывод переменной с тиом bool:", d)

	fmt.Println("Задание 3")
	a1 := 5
	b1 := 5.32
	c1 := "Hello, world!!!"
	d1 := true

	fmt.Println("Вывод переменной с тиом int с краткой формой объявления:", a1)
	fmt.Println("Вывод переменной с тиом float32 с краткой формой объявления:", b1)
	fmt.Println("Вывод переменной с тиом string с краткой формой объявления:", c1)
	fmt.Println("Вывод переменной с тиом bool с краткой формой объявления:", d1)

	fmt.Println("Задание 4")

	var x int

	fmt.Print("Введите первое число: ")
	fmt.Scan(&x)
	var y int

	fmt.Print("Введите второе число: ")
	fmt.Scan(&y)

	sum := x + y
	difference := x - y
	difference1 := y - x
	product := x * y
	quotient := float32(x) / float32(y)
	quotient1 := float32(y) / float32(x)

	fmt.Println("1 - Сложение \n2 - вычитание 2-го числа из 1-го\n3 - вычитание 1-го числа из 2-го\n4 - умножение\n5 - деление 1-го на 2-ое\n6 - деление 2-го числа на 1-ое")
	var action int
	fmt.Print("Выберите действие:")
	fmt.Scan(&action)
	switch action {
	case 1:
		fmt.Println("Сумма чисел", fmt.Sprint(sum))

	case 2:
		fmt.Println("Вычитание второго числа из первого:", fmt.Sprint(difference))
	case 3:
		fmt.Println("Вычитание первого числа из второго:", fmt.Sprint(difference1))
	case 4:
		fmt.Println("умножение чисел:", fmt.Sprint(product))
	case 5:
		if y != 0 {
			fmt.Println("Деление первого числа на второе: %.2f\n", fmt.Sprint(quotient))
		} else {
			fmt.Println("Ошибка:делить на ноль - нельзя")
		}
	case 6:
		if x != 0 {
			fmt.Println("Деление второго числа на первое: %.2f\n", fmt.Sprint(quotient1))
		} else {
			fmt.Println("Ошибка:делить на ноль - нельзя")
		}
	default:
		fmt.Println("Данного действия не существует")
	}

	fmt.Println("Задание 5")
	var x1, y1 float64
	var summa, diff float64
	fmt.Print("Введите 1-ое число:")
	fmt.Scan(&x1)
	fmt.Print("Введите 2-ое число:")
	fmt.Scan(&y1)
	summa, diff = SumAndDifference(x1, y1)

	fmt.Println("Сумма:", summa)
	fmt.Println("Разность:", diff)

	fmt.Println("Задание 6")
	var q, w, e float32
	fmt.Print("Введите 1-ое число:")
	fmt.Scan(&q)
	fmt.Print("Введите 2-ое число:")
	fmt.Scan(&w)
	fmt.Print("Введите 3-ее число:")
	fmt.Scan(&e)
	average := (q + w + e) / 3
	fmt.Println("Среднее значение трех чисел:", average)

}
