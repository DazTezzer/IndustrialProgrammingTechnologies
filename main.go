package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"strings"
	"sort"
	"unicode"
)

func main() {
	for {
		tasksList()
		choose()
	}
}

func tasksList() {
	fmt.Println("Выберите задание:")
	fmt.Println("1 - Перевод чисел из одной системы счисления в другую")
	fmt.Println("2 - Решение квадратного уравнения")
	fmt.Println("3 - Сортировка чисел по модулю")
	fmt.Println("4 - Слияние двух отсортированных массивов")
	fmt.Println("5 - Нахождение подстроки в строке без использования встроенных функций")
	fmt.Println("6 - Калькулятор с расширенными операциями")
	fmt.Println("7 - Проверка палиндрома")
	fmt.Println("8 - Нахождение пересечения трех отрезков")
	fmt.Println("9 - Выбор самого длинного слова в предложении")
	fmt.Println("10 - Проверка высокосного года")
	fmt.Println("11 - Числа Фибоначчи до определенного числа")
	fmt.Println("12 - Определение простых чисел в диапазоне")
	fmt.Println("13 - Числа Армстронга в заданном диапазоне")
	fmt.Println("14 - Реверс строки")
	fmt.Println("15 - Нахождение наибольшего общего делителя")

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
	var number string
	var fromBase, toBase int

	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	fmt.Print("Введите исходную систему (от 2 до 36): ")
	fmt.Scan(&fromBase)
	fmt.Print("Введите конечную систему (от 2 до 36): ")
	fmt.Scan(&toBase)

	decimalValue, err := toDecimal(number, fromBase)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	result := fromDecimal(decimalValue, toBase)
	fmt.Printf("Число %s в системе счисления %d равно %s в системе счисления %d\n", number, fromBase, result, toBase)
}

func toDecimal(number string, base int) (int, error) {
	result := 0
	for i, char := range strings.ToUpper(number) {
		var value int
		switch {
		case char >= '0' && char <= '9':
			value = int(char - '0')
		case char >= 'A' && char <= 'Z':
			value = int(char - 'A' + 10)
		default:
			return 0, fmt.Errorf("invalid character: %c", char)
		}

		if value >= base {
			return 0, fmt.Errorf("digit %c not valid in base %d", char, base)
		}
		result += value * int(math.Pow(float64(base), float64(len(number)-i-1)))
	}
	return result, nil
}

func fromDecimal(number, base int) string {
	if number == 0 {
		return "0"
	}

	var result strings.Builder
	for number > 0 {
		remainder := number % base
		var char rune
		if remainder < 10 {
			char = rune('0' + remainder)
		} else {
			char = rune('A' + remainder - 10)
		}
		result.WriteRune(char)
		number /= base
	}

	runes := []rune(result.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func task2() {
	var a, b, c float64

	fmt.Print("Введите коэффициент a: ")
	fmt.Scan(&a)
	fmt.Print("Введите коэффициент b: ")
	fmt.Scan(&b)
	fmt.Print("Введите коэффициент c: ")
	fmt.Scan(&c)

	if a == 0 {
		fmt.Println("Коэффициент a не должен быть равен нулю.")
		return
	}

	delta := b*b - 4*a*c
	sqrtDelta := cmplx.Sqrt(complex(delta, 0))

	root1 := complex(-b, 0) + sqrtDelta / complex(2*a, 0)
	root2 := complex(-b, 0) - sqrtDelta / complex(2*a, 0)

	fmt.Printf("Корни квадратного уравнения: \nКорень 1: %v\nКорень 2: %v\n", root1, root2)
}

func task3() {
	arr := []int{-3, 1, -2, 4, -1, 0, 2}
	sort.Slice(arr, func(i, j int) bool {
		return abs(arr[i]) < abs(arr[j])
	})
	fmt.Println("Отсортированный массив по модулю:", arr)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func task4() {
	var n1, n2 int

	fmt.Print("Введите количество элементов первого массива: ")
	fmt.Scan(&n1)
	arr1 := make([]int, n1)
	fmt.Println("Введите элементы первого массива (в порядке возрастания):")
	for i := 0; i < n1; i++ {
		fmt.Scan(&arr1[i])
	}

	fmt.Print("Введите количество элементов второго массива: ")
	fmt.Scan(&n2)
	arr2 := make([]int, n2)
	fmt.Println("Введите элементы второго массива (в порядке возрастания):")
	for i := 0; i < n2; i++ {
		fmt.Scan(&arr2[i])
	}

	merged := func(arr1, arr2 []int) []int {
		var result []int
		i, j := 0, 0

		for i < len(arr1) && j < len(arr2) {
			if arr1[i] < arr2[j] {
				result = append(result, arr1[i])
				i++
			} else {
				result = append(result, arr2[j])
				j++
			}
		}

		for i < len(arr1) {
			result = append(result, arr1[i])
			i++
		}
		for j < len(arr2) {
			result = append(result, arr2[j])
			j++
		}

		return result
	}(arr1, arr2)

	fmt.Println("Слитый отсортированный массив:", merged)
}

func task5() {
	var haystack, needle string

	fmt.Println("Введите строку (главная):")
	fmt.Scanln(&haystack)
	fmt.Println("Введите подстроку для поиска:")
	fmt.Scanln(&needle)

	index := findSubstring(haystack, needle)
	fmt.Println("Индекс первого вхождения:", index)
}

func findSubstring(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		j := 0
		for j < len(needle) && haystack[i+j] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}

func task6() {
	var a, b float64
	var operator string

	fmt.Println("Введите первое число:")
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Ошибка ввода числа:", err)
		return
	}

	fmt.Println("Введите второе число:")
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("Ошибка ввода числа:", err)
		return
	}

	fmt.Println("Введите оператор (+, -, *, /, ^, %):")
	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("Ошибка ввода оператора:", err)
		return
	}

	var result float64
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: Деление на ноль!")
			return
		}
		result = a / b
	case "^":
		result = math.Pow(a, b)
	case "%":
		result = math.Mod(a, b)
	default:
		fmt.Println("Ошибка: Недопустимая операция")
		return
	}

	fmt.Printf("Результат операции %f %s %f = %f\n", a, operator, b, result)
}

func task7() {
	var input string

	fmt.Println("Введите строку для проверки на палиндром:")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	isPalindrome := func(s string) bool {
		var filteredRunes []rune
		for _, r := range s {
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				filteredRunes = append(filteredRunes, unicode.ToLower(r))
			}
		}

		for i, j := 0, len(filteredRunes)-1; i < j; i, j = i+1, j-1 {
			if filteredRunes[i] != filteredRunes[j] {
				return false
			}
		}
		return true
	}

	if isPalindrome(input) {
		fmt.Println("Строка является палиндромом.")
	} else {
		fmt.Println("Строка не является палиндромом.")
	}
}

func task8() {
	var a1, a2, b1, b2, c1, c2 int

	fmt.Println("Введите начальные и конечные точки трех отрезков (a1, a2, b1, b2, c1, c2):")
	_, err := fmt.Scanf("%d %d %d %d %d %d", &a1, &a2, &b1, &b2, &c1, &c2)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	isIntersecting := func(a1, a2, b1, b2, c1, c2 int) bool {
		max := func(x, y int) int {
			if x > y {
				return x
			}
			return y
		}

		min := func(x, y int) int {
			if x < y {
				return x
			}
			return y
		}

		start := max(max(a1, b1), c1)
		end := min(min(a2, b2), c2)
		return start <= end
	}

	if isIntersecting(a1, a2, b1, b2, c1, c2) {
		fmt.Println("Существует область пересечения всех трех отрезков.")
	} else {
		fmt.Println("Нет области пересечения всех трех отрезков.")
	}
}

func task9() {
	fmt.Println("Введите предложение:")
	var input string
	fmt.Scanln(&input)

	longestWord := ""
	words := strings.Fields(input)

	for _, word := range words {
		cleanedWord := ""
		for _, char := range word {
			if unicode.IsLetter(char) || unicode.IsNumber(char) {
				cleanedWord += string(char)
			}
		}

		if len(cleanedWord) > len(longestWord) {
			longestWord = cleanedWord
		}
	}

	fmt.Println("Самое длинное слово:", longestWord)
}

func task10() {
	fmt.Println("Введите год:")
	var year int
	fmt.Scan(&year)
	isLeapYear := func(y int) bool {
		return (y%4 == 0 && y%100 != 0) || (y%400 == 0)
	}

	fmt.Println(isLeapYear(year))
}

func task11() {
	fmt.Println("Введите целое число:")
	var n int
	fmt.Scan(&n)

	fib := func(limit int) {
		a, b := 0, 1
		for a <= limit {
			fmt.Print(a, " ")
			a, b = b, a+b
		}
	}

	fib(n)
}

func task12() {
	var start, end int
	fmt.Println("Введите два целых числа (начало и конец диапазона):")
	fmt.Scan(&start, &end)

	func(start, end int) {
		for num := start; num <= end; num++ {
			if num < 2 {
				continue
			}
			isPrime := true
			for i := 2; i*i <= num; i++ {
				if num%i == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				fmt.Print(num, " ")
			}
		}
	}(start, end)
}

func task13() {
	var start, end int
	fmt.Println("Введите два целых числа (начало и конец диапазона):")
	fmt.Scan(&start, &end)

	func(start, end int) {
		for num := start; num <= end; num++ {
			digits := 0
			temp := num
			for temp > 0 {
				temp /= 10
				digits++
			}

			sum := 0
			temp = num
			for temp > 0 {
				digit := temp % 10
				sum += int(math.Pow(float64(digit), float64(digits)))
				temp /= 10
			}

			if sum == num {
				fmt.Print(num, " ")
			}
		}
	}(start, end)
}

func task14() {
	var input string
	fmt.Println("Введите строку:")
	fmt.Scan(&input)

	reversed := ""
	for i := len(input) - 1; i >= 0; i-- {
		reversed += string(input[i])
	}

	fmt.Println("Преобразованная строка:", reversed)
}

func task15() {
	var a, b int
	fmt.Println("Введите два числа:")
	fmt.Scan(&a, &b)

	for b != 0 {
		a, b = b, a%b
	}

	fmt.Println("Наибольший общий делитель (НОД):", a)
}
