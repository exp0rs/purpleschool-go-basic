package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readOperation() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Выберите операцию (AVG - среднее, SUM - сумма, MED - медиана)")
	operation, _ := reader.ReadString('\n')
	operation = strings.TrimSpace(operation)

	return normalizeOperation(operation)
}

func readTransactions() ([]float64, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите трансакции через запятую (пример: 9,1,-5)")
	transactions, _ := reader.ReadString('\n')
	transactions = strings.TrimSpace(transactions)

	return normalizeTransactions(transactions)
}

func normalizeOperation(operation string) (string, error) {
	op := strings.TrimSpace(operation)
	if op == "" {
		return "", fmt.Errorf("операция не может быть пустой")
	}
	op = strings.ToUpper(op)

	switch op {
	case "AVG", "MED", "SUM":
		return op, nil
	default:
		return "", fmt.Errorf("неизвестная операция: %s", op)
	}
}

func normalizeTransactions(transactions string) ([]float64, error) {
	trimmed := strings.TrimSpace(transactions)
	if trimmed == "" {
		return nil, fmt.Errorf("список трансакций не может быть пустым")
	}

	parts := strings.Split(trimmed, ",")
	nums := make([]float64, 0, len(parts))

	for i, part := range parts {
		trimmedPart := strings.TrimSpace(part)
		if trimmedPart == "" {
			return nil, fmt.Errorf("элемент %d: пустое значение", i+1)
		}

		num, err := strconv.ParseFloat(trimmedPart, 64)
		if err != nil {
			return nil, fmt.Errorf("элемент %d: некорректное число '%s'", i+1, trimmedPart)
		}
		nums = append(nums, num)
	}

	return nums, nil
}

func calculateSum(nums []float64) float64 {
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func calculateAvg(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	return calculateSum(nums) / float64(len(nums))
}

func calculateMed(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}

	// Создаем копию чтобы не менять оригинальный слайс
	sorted := make([]float64, len(nums))
	copy(sorted, nums)
	sort.Float64s(sorted)

	n := len(sorted)
	if n%2 == 0 {
		// Четное количество элементов
		return (sorted[n/2-1] + sorted[n/2]) / 2
	}
	// Нечетное количество элементов
	return sorted[n/2]
}

func main() {
	// Сначала запрашиваем и проверяем операцию
	operation, err := readOperation()
	if err != nil {
		fmt.Printf("Ошибка в операции: %v\n", err)
		return
	}

	// Только если операция корректна, запрашиваем транзакции
	transactions, err := readTransactions()
	if err != nil {
		fmt.Printf("Ошибка в трансакциях: %v\n", err)
		return
	}

	// Выполняем вычисление
	var result float64
	switch operation {
	case "SUM":
		result = calculateSum(transactions)
	case "AVG":
		result = calculateAvg(transactions)
	case "MED":
		result = calculateMed(transactions)
	}

	// Форматируем вывод без лишних нулей
	fmt.Printf("Результат: %g\n", result)
}
