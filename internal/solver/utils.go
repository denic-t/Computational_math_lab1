package solver

import "math"

func calculateNormC(A [][]float64) float64 {
	n := len(A)
	maxNorm := 0.0
	for i := 0; i < n; i++ {
		sum := 0.0
		for j := 0; j < n; j++ {
			if i != j {
				sum += math.Abs(A[i][j] / A[i][i])
			}
		}
		if sum > maxNorm {
			maxNorm = sum
		}
	}
	return maxNorm
}

func makeDiagonallyDominant(A [][]float64, B []float64) bool {
	n := len(A)
	usedRows := make([]bool, n)
	newA := make([][]float64, n)
	newB := make([]float64, n)

	for i := 0; i < n; i++ {
		newA[i] = make([]float64, n)
	}

	hasStrictInequality := false

	for i := 0; i < n; i++ {
		found := false
		for r := 0; r < n; r++ {
			if usedRows[r] {
				continue
			}

			sum := 0.0
			for c := 0; c < n; c++ {
				if c != i {
					sum += math.Abs(A[r][c])
				}
			}

			if math.Abs(A[r][i]) >= sum && A[r][i] != 0 {
				copy(newA[i], A[r])
				newB[i] = B[r]
				usedRows[r] = true
				found = true
			}

			if math.Abs(A[r][i]) > sum {
				hasStrictInequality = true
			}
			break
		}

		if !found {
			return false
		}
	}

	if !hasStrictInequality {
		return false
	}
	for i := 0; i < n; i++ {
		copy(A[i], newA[i])
		B[i] = newB[i]
	}
	return true
}
