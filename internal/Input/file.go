package Input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FileProvider struct {
	Path string
}

func (f FileProvider) Read() (*MatrixData, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите название файла: ")
	filename, _ := reader.ReadString('\n')
	filename = strings.TrimSpace(filename)

	file, err := os.Open(filename)

	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	var n int
	if _, err := fmt.Fscan(file, &n); err != nil || n <= 0 {
		return nil, fmt.Errorf("Не удалось считать размерность матрицы из файла или размерность <= 0")
	}

	A := make([][]float64, n)
	B := make([]float64, n)
	var eps float64

	for i := 0; i < n; i++ {
		A[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			if _, err := fmt.Fscan(file, &A[i][j]); err != nil {
				return nil, fmt.Errorf("Ошибка чтения элемента A[%d][%d]", i, j)
			}
		}

		if _, err := fmt.Fscan(file, &B[i]); err != nil {
			return nil, fmt.Errorf("Ошибка чтения элемента B[%d]", i)
		}
	}

	if _, err := fmt.Fscan(file, &eps); err != nil || eps <= 0 {
		return nil, fmt.Errorf("Ошибка чтения eps или eps <= 0")
	}

	return &MatrixData{A, B, eps}, nil

}
