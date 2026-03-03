package Input

type MatrixData struct {
	A   [][]float64
	B   []float64
	Eps float64
}

type Provider interface {
	Read() (*MatrixData, error)
}
