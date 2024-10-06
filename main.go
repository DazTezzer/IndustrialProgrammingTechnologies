package main

import (
	"fmt"
	"math"
)

func main() {
	for {
		tasksList()
		choose()
	}
}

func tasksList() {
	fmt.Println("Выберите задание:")
	fmt.Println("1 - Сумма цифр числа")
	fmt.Println("2 - Преобразование температуры")
	fmt.Println("3 - Удвоение каждого элемента массива")
	fmt.Println("4 - Объединение строк")
	fmt.Println("5 - Расчет расстояния между двумя точками")
	fmt.Println("6 - Проверка на четность/нечетность")
	fmt.Println("7 - Проверка высокосного года")
	fmt.Println("8 - Определение наибольшего из трех чисел")
	fmt.Println("9 - Категория возраста")
	fmt.Println("10 - Проверка делимости на 3 и 5")
	fmt.Println("11 - Факториал числа")
	fmt.Println("12 - Числа Фибоначчи")
	fmt.Println("13 - Реверс массива")
	fmt.Println("14 - Поиск простых чисел")
	fmt.Println("15 - Сумма чисел в массиве")

}

func choose() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	switch num {
	case 1:
		task1()
	case 2:
		task2()
	case 3:
		task3()
	case 4:
		task4()
	case 5:
		task5()
	case 6:
		task6()
	case 7:
		task7()
	case 8:
		task8()
	case 9:
		task9()
	case 10:
		task10()
	case 11:
		task11()
	case 12:
		task12()
	case 13:
		task13()
	case 14:
		task14()
	case 15:
		task15()
	default:
		fmt.Println("Неправильный код")
		return
	}
}

func task1() {
	var number int
	fmt.Print("Введите целое число: ")
	fmt.Scan(&number)
	if number < 0 {
		number = -number
	}
	sum := sumDigits(number)
	fmt.Println("Сумма цифр числа:", sum)
}

func sumDigits(number int) int {
	sum := 0
	for number > 0 {
		sum += number % 10
		number /= 10
	}
	return sum
}

func task2() {
	var C float64
	fmt.Print("Введите температуру в градусах Цельсия: ")
	fmt.Scan(&C)
	F := C*9/5 + 32
	fmt.Printf("%.2f градусов Цельсия = %.2f градусов Фаренгейта\n", C, F)
}

func task3() {
	var size int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&size)

	numbers := make([]int, size)
	fmt.Println("Введите числа:")
	for i := 0; i < size; i++ {
		fmt.Scan(&numbers[i])
	}

	doubled := make([]int, len(numbers))
	for i, num := range numbers {
		doubled[i] = num * 2
	}

	fmt.Println("Удвоенные числа:", doubled)
}

func task4() {
	var n int
	fmt.Print("Введите количество строк: ")
	fmt.Scan(&n)
	arr := make([]string, n)
	fmt.Println("Введите строки:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	var result string
	for i := 0; i < n; i++ {
		result += arr[i]
		if i < n-1 {
			result += " "
		}
	}
	fmt.Println("Объединенная строка:", result)
}

func task5() {
	var x1, y1, x2, y2 float64
	fmt.Print("Введите координаты первой точки (x1, y1): ")
	fmt.Scan(&x1, &y1)
	fmt.Print("Введите координаты второй точки (x2, y2): ")
	fmt.Scan(&x2, &y2)
	distance := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	fmt.Printf("Расстояние между точками (%.1f, %.1f) и (%.1f, %.1f) равно %.1f\n", x1, y1, x2, y2, distance)
}

func task6() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	if number%2 == 0 {
		fmt.Printf("%d является четным.\n", number)
	} else {
		fmt.Printf("%d является нечетным.\n", number)
	}
}

func task7() {
	var year int
	var isLeapYear bool
	fmt.Print("Введите год: ")
	fmt.Scan(&year)
	isLeapYear = year%4 == 0 && (year%100 != 0 || year%400 == 0)
	if isLeapYear {
		fmt.Printf("%d является високосным годом.\n", year)
	} else {
		fmt.Printf("%d не является високосным годом.\n", year)
	}
}

func task8() {
	var num1, num2, num3 int
	fmt.Print("Введите три числа: ")
	fmt.Scan(&num1, &num2, &num3)
	max := num1
	if num2 > max {
		max = num2
	}
	if num3 > max {
		max = num3
	}
	fmt.Printf("Наибольшее число: %d\n", max)
}

func task9() {
	var age int
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)
	if age < 14 {
		fmt.Println("ребенок")
	} else if age < 18 {
		fmt.Println("подросток")
	} else if age < 65 {
		fmt.Println("взрослый")
	} else {
		fmt.Println("пожилой")
	}
}

func task10() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	if number%3 == 0 && number%5 == 0 {
		fmt.Printf("Число %d делится на 3 и 5.\n", number)
	} else {
		fmt.Printf("Число %d не делится одновременно на 3 и 5.\n", number)
	}
}

func task11() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	if number < 0 {
		fmt.Println("Нет факториал отрицательного числа")
	} else {
		result := 1
		for i := 1; i <= number; i++ {
			result *= i
		}
		fmt.Printf("Факториал числа %d равен %d\n", number, result)
	}
}

func task12() {
	var number int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&number)
	if number <= 0 {
		fmt.Println("Введено отрицательное число")
	} else {
		fib := make([]int, number)
		if number > 0 {
			fib[0] = 0
		}
		if number > 1 {
			fib[1] = 1
		}
		for i := 2; i < number; i++ {
			fib[i] = fib[i-1] + fib[i-2]
		}
		fmt.Printf("Первые %d чисел Фибоначчи: %v\n", number, fib)
	}
}

func task13() {
	var number int
	fmt.Print("Введите количество элементов массива: ")
	fmt.Scan(&number)
	arr := make([]int, number)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < number; i++ {
		fmt.Scan(&arr[i])
	}
	reversed := make([]int, number)
	for i := 0; i < number; i++ {
		reversed[i] = arr[number-i-1]
	}
	fmt.Println("Реверсированный массив:", reversed)
}

func task14() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	primes := []int{}
	for i := 2; i <= n; i++ {
		prime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				prime = false
				break
			}
		}
		if prime {
			primes = append(primes, i)
		}
	}
	fmt.Println("Простые числа до", n, ":", primes)
}

func task15() {
	var numbers int
	fmt.Print("Введите количество элементов массива: ")
	fmt.Scan(&numbers)
	arr := make([]int, numbers)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < numbers; i++ {
		fmt.Scan(&arr[i])
	}
	sum := 0
	for _, number := range arr {
		sum += number
	}
	fmt.Println("Сумма чисел в массиве:", sum)
}
