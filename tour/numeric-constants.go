package tour

const (
	Big = 1 << 100
	Small = Big >> 99
)

func NeedInt(x int) int  {
	return x*10 + 1
}

func NeedFloat(x float64) float64  {
	return x * 0.1
}
