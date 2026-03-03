package main

import (
	"Computational_mathematics/internal/Input"
	"Computational_mathematics/internal/solver"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var provider Input.Provider
	reader := bufio.NewReader(os.Stdin)
	var A [][]float64
	var B []float64
	var epsilon float64
	var err error
	maxIterations := 10000
	var matrixData *Input.MatrixData
	fmt.Println("Запуск решения системы линейных алгебраических уравнений СЛАУ")

	for {
		fmt.Println("Выберите '1' или '2' для чтения данных:")
		fmt.Println("'1' - пользовательский ввод с консоли")
		fmt.Println("'2' - чтение из файла")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "1" {
			provider = Input.ConsoleProvider{}
			matrixData, err = provider.Read()
			if err != nil {
				fmt.Printf("Ошибка ввода: %v\n", err)
				continue
			}
			break
		} else if input == "2" {
			provider = Input.FileProvider{}
			matrixData, err = provider.Read()
			if err != nil {
				fmt.Printf("Ошибка чтения файла %v\n\n", err)
				continue
			}
			break
		} else {
			fmt.Println("Ошибка: неверный выбор. Пожалуйста, введите 1 или 2.\n")
		}
	}

	fmt.Println("\n=== Решение СЛАУ методом Гаусса-Зейделя ===")

	A = matrixData.A
	B = matrixData.B
	epsilon = matrixData.Eps

	fmt.Printf("\n===Результаты работы программы===\n")

	x, iters, errVec, err := solver.GaussSeidel(A, B, epsilon, maxIterations)
	if err != nil {
		fmt.Printf("\n[!] Алгоритм завершился с ошибкой: %v\n", err)
	}

	fmt.Printf("Количество итераций: %d\n", iters)

	fmt.Println("Вектор неизвестных (x_1, x_2, ..., x_n):")
	for i, val := range x {
		fmt.Printf("  x_%d = %.6f\n", i+1, val)
	}

	fmt.Println("Вектор погрешностей |x_i^(k) - x_i^(k-1)|:")
	for i, val := range errVec {
		fmt.Printf("  e_%d = %.6f\n", i+1, val)
	}
}
