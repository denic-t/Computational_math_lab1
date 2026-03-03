package solver

import (
	"fmt"
	"math"
)

func GaussSeidel(A [][]float64, b []float64, eps float64, maxIter int) ([]float64, int, []float64, error) {
	n := len(b)
	x := make([]float64, n)
	errors := make([]float64, n)

	if !makeDiagonallyDominant(A, b) {
		fmt.Println("[-] ВНИМАНИЕ: Невозможно достичь строгого диагонального преобладания путем перестановки строк. Сходимость не гарантируется!")
	} else {
		fmt.Println("[+] Диагональное преобладание успешно достигнуто (матрица переставлена).")
	}

	normC := calculateNormC(A)
	fmt.Printf("[i] Норма матрицы ||C|| = %.5f\n", normC)
	if normC > 1 {
		fmt.Println("[-] Внимание: Норма ||C|| > 1. Достаточное условие сходимости нарушено!")
	} else if normC == 1 {
		fmt.Println("[i] Норма ||C|| = 1. Выполняется ослабленное условие сходимости (есть нестрогие неравенства).")
	} else {
		fmt.Println("[+] Норма ||C|| < 1. Условие сходимости выполняется строго.")
	}

	for iter := 1; iter <= maxIter; iter++ {
		maxError := 0.0
		for i := 0; i < n; i++ {
			sum := 0.0
			for j := 0; j < n; j++ {
				if i != j {

					sum += A[i][j] * x[j]
				}
			}

			newX_i := (b[i] - sum) / A[i][i]

			err := math.Abs(newX_i - x[i])
			errors[i] = err
			if err > maxError {
				maxError = err
			}

			x[i] = newX_i
		}

		if maxError <= eps {
			return x, iter, errors, nil
		}
	}

	return x, maxIter, errors, fmt.Errorf("Превышено максимальное число итераций: %d", maxIter)
}
