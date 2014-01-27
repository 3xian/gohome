package geometry

const (
	EPS = 1e-7
)

type Line struct {
	p1, p2 complex128
}

func CrossProduct(a complex128, b complex128) float64 {
	return real(a)*imag(b) - real(b)*imag(a)
}

func Roatate(p complex128, a float64) complex128 {
	return p * complex(math.Cos(a), math.Sin(a))
}

func Intersact(a Line, b Line) (intersaction complex128, isIntersact bool) {
	var (
		s1 = CrossProduct(a.p2-a.p1, b.p1-a.p1)
		s2 = CrossProduct(a.p2-a.p1, b.p2-a.p1)
		s  = s2 - s1
	)
	if math.Abs(s) >= EPS {
		isIntersact = true
		intersaction = complex((real(b.p1)*s2-real(b.p2)*s1)/s, (imag(b.p1)*s2-imag(b.p2)*s1)/s)
	}
	return
}
