package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите выражение в формате 'число операция число', например '3 + 5' или 'III * VI':")

	// Чтение ввода с консоли
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	input = strings.ToUpper(input)

	// Удаление лишних пробелов из ввода
	input = strings.TrimSpace(input)

	// Проверка, что ввод не пустой
	if input == "" {
		panic("Неправильный формат ввода")
	}

	// Поиск позиции оператора
	operatorPos := -1
	for i, char := range input {
		if char == '+' || char == '-' || char == '*' || char == '/' {
			if operatorPos != -1 {
				// Если уже был найден оператор, выдаем панику о неправильном формате
				panic("В вводе должны быть только два операнда и один оператор (+, -, /, *)")
			}
			operatorPos = i
		}
	}

	if operatorPos == -1 {
		panic("В вводе должны быть только два операнда и один оператор (+, -, /, *)")
	}

	// Разделение ввода на левую и правую части
	leftStr := strings.TrimSpace(input[:operatorPos])
	rightStr := strings.TrimSpace(input[operatorPos+1:])

	// Проверка, что обе части не пустые
	if leftStr == "" || rightStr == "" {
		panic("В вводе должны быть только два операнда и один оператор (+, -, /, *)")
	}

	// Преобразование чисел в числовой формат
	leftNum, leftIsRoman, err := parseNumber(leftStr)
	if err != nil {
		panic(err)
	}

	rightNum, rightIsRoman, err := parseNumber(rightStr)
	if err != nil {
		panic(err)
	}

	// Проверка на смешивание систем чисел
	if leftIsRoman && !rightIsRoman || !leftIsRoman && rightIsRoman {
		panic("Невозможно смешивать арабские и римские числа")
	}

	// Выполнение операции
	operator := string(input[operatorPos])
	result := calculate(leftNum, rightNum, operator)

	// Вывод результата
	if leftIsRoman || rightIsRoman {
		// Преобразование результата в римское число, если хотя бы один операнд был римским
		fmt.Println("Результат:", intToRoman(result.(int)))
	} else {
		// Вывод арабского числа, если оба операнда были арабскими
		fmt.Println("Результат:", result.(int))
	}
}

// Функция для парсинга числа из строки
func parseNumber(str string) (interface{}, bool, error) {
	// Проверка на арабские числа
	arabic, err := strconv.Atoi(str)
	if err == nil {
		if arabic < 1 || arabic > 10 {
			return 0, false, fmt.Errorf("аргумент '%s' вне диапазона (1-10)", str)
		}
		return arabic, false, nil
	}

	// Проверка на римские числа
	roman := map[rune]int{
		'I': 1, 'V': 5, 'X': 10,
	}
	sum := 0
	lastValue := 0
	for _, c := range str {
		value, ok := roman[c]
		if !ok {
			return 0, false, fmt.Errorf("неверный формат числа '%s'", str)
		}
		sum += value
		if lastValue < value {
			sum -= 2 * lastValue
		}
		lastValue = value
	}

	// Проверка на отрицательные римские числа
	if sum <= 0 {
		return 0, false, fmt.Errorf("в римской системе нет отрицательных чисел: '%s'", str)
	}

	return sum, true, nil
}

// Функция для выполнения арифметической операции
func calculate(a, b interface{}, op string) interface{} {
	switch op {
	case "+":
		return a.(int) + b.(int)
	case "-":
		return a.(int) - b.(int)
	case "*":
		return a.(int) * b.(int)
	case "/":
		if b.(int) == 0 {
			panic("деление на ноль")
		}
		return a.(int) / b.(int)
	default:
		panic("неподдерживаемая операция: " + op)
	}
}

// Функция для преобразования целого числа в римское
func intToRoman(num int) string {
	if num <= 0 || num > 3999 {
		panic("невозможно преобразовать число в римскую систему")
	}

	vals := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, v := range vals {
		for num >= v.Value {
			roman += v.Symbol
			num -= v.Value
		}
	}

	return roman
}

// go run main.go
