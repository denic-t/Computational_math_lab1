package Input

import "fmt"

type ConsoleProvider struct{}

func (c ConsoleProvider) Read() (*MatrixData, error) {
	var n int
	fmt.Print("Введите размерность матрицы (n): ")
	if _, err := fmt.Scan(&n); err != nil || n <= 0 {
		return nil, fmt.Errorf("неверная размерность")
	}

	A := make([][]float64, n)
	B := make([]float64, n)

	fmt.Println("Введите расширенную матрицу построчно (коэффициенты A и свободный член B через пробел):")
	for i := 0; i < n; i++ {
		A[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&A[i][j])
		}
		fmt.Scan(&B[i])
	}

	var eps float64
	fmt.Print("Введите точность epsilon: ")
	if _, err := fmt.Scan(&eps); err != nil || eps <= 0 {
		return nil, fmt.Errorf("Неверная точность")
	}

	return &MatrixData{A, B, eps}, nil
}
