package poly

import (
	"math"
	"math/rand"
)

type Poly struct {
	deg   int
	coeff []int
}

func NewPoly(coeff []int) Poly {
	p := Poly{len(coeff) - 1, coeff}
	return p
}

func ZeroPoly(deg int) Poly {
	p := Poly{deg, make([]int, deg+1)}
	return p
}

func (p *Poly) AddElem(elem int) {
	for idx := 0; idx <= p.deg; idx++ {
		p.coeff[idx] += elem
	}
}

/* set's coefficient of given degree to elem */
func (p *Poly) SetElem(elem int, deg int) bool {
	if p.deg-deg < 0 {
		return false
	}
	p.coeff[p.deg-deg] = elem
	return true
}

/* Return's the coefficient of given degree in poslynomial */
func (p *Poly) GetElem(deg int) (int, bool) {
	if p.deg-deg < 0 {
		return 0, false
	}
	return p.coeff[p.deg-deg], true
}

func (p *Poly) GetDeg() int {
	return p.deg
}

/* subtracts individual coefficient of polynomial with given elem */
func (p *Poly) SubElem(elem int) {
	for idx := 0; idx <= p.deg; idx++ {
		p.coeff[idx] -= elem
	}
}

/* Modulos individual coefficient of polynomial with given elem */
func (p *Poly) ModElem(mod int) {
	for idx := 0; idx <= p.deg; idx++ {
		p.coeff[idx] += mod
		p.coeff[idx] %= mod
	}
}

/* Multilies individual coefficient of polynomial with given elem */
func (p *Poly) MulElem(elem float64) {
	for idx := 0; idx <= p.deg; idx++ {
		p.coeff[idx] = int(math.Round(float64(p.coeff[idx]) * elem))
	}
}

/* Divides individual coefficient of polynomial with given elem */
func (p *Poly) DivElem(elem float64) {
	for idx := 0; idx <= p.deg; idx++ {
		p.coeff[idx] = int(math.Round(float64(p.coeff[idx]) / elem))
	}
}

/*
Performs Polynomial Addition
Argument:
	p,q : polynomials to be added
Returns:
	resultant polynomial
*/
func PolyAdd(p Poly, q Poly) Poly {
	ans := ZeroPoly(int(math.Max(float64(q.deg), float64(p.deg))))

	for idx := ans.deg; idx >= 0; idx-- {
		if p.deg-(ans.deg-idx) >= 0 {
			ans.coeff[idx] += p.coeff[p.deg-(ans.deg-idx)]
		}
		if q.deg-(ans.deg-idx) >= 0 {
			ans.coeff[idx] += q.coeff[q.deg-(ans.deg-idx)]
		}
	}
	return ans
}

/*
Performs Polynomial Multiplication
Argument:
	p,q : polynomials to be multiplied
Returns:
	resultant polynomial
*/
func PolyMul(p Poly, q Poly) Poly {
	mul_Poly := ZeroPoly(p.deg + q.deg)

	for idx := 0; idx <= mul_Poly.deg; idx++ {
		req_deg := mul_Poly.deg - idx
		for idx2, val := range p.coeff {
			if (p.deg - idx2) > req_deg {
				continue
			}
			if req_deg-(p.deg-idx2) > q.deg || req_deg-(p.deg-idx2) < 0 {
				continue
			}
			req_idx := q.deg - (req_deg - (p.deg - idx2))
			mul_Poly.coeff[idx] += val * q.coeff[req_idx]
		}
	}
	return mul_Poly
}

/*
Helper function to convert an array of integers to float64
*/
func ConvertToFloat(arr []int) []float64 {
	converted := make([]float64, len(arr))

	for idx, val := range arr {
		converted[idx] = float64(val)
	}
	return converted
}

/*
Helper function to convert an array of float64 to integers
*/
func ConvertToInt(arr []float64) []int {
	converted := make([]int, len(arr))

	for idx, val := range arr {
		converted[idx] = int(math.Round(val))
	}
	return converted
}

/*
Performs Polynomial Division
Argument:
	p : dividend polynomial
 	q : divisor polynomial
Returns:
	a tuple of (qutoient, remainder)
*/
func PolyDiv(p Poly, q Poly) (Poly, Poly) {
	quo := []float64{}
	divisor := ConvertToFloat(q.coeff)
	dividend := ConvertToFloat(p.coeff)

	for idx, val := range dividend {
		if val != float64(0) {
			dividend = dividend[idx:]
			break
		}
		if idx == len(dividend)-1 {
			return ZeroPoly(0), ZeroPoly(0)
		}
	}

	for idx, val := range divisor {
		if val != float64(0) {
			divisor = divisor[idx:]
			break
		}
	}

	p_deg := len(dividend) - 1
	q_deg := len(divisor) - 1

	for p_deg >= q_deg {
		mx_elem_divisor := divisor[0]
		mx_elem_dividend := dividend[0]
		elem_to_mul := mx_elem_dividend / mx_elem_divisor

		to_sub := []float64{}
		quo = append(quo, elem_to_mul)
		for _, val := range divisor {
			to_sub = append(to_sub, elem_to_mul*val)
		}

		for i := 0; i < p_deg-q_deg; i++ {
			to_sub = append(to_sub, 0)
		}

		left_dividend := []float64{}
		for idx, val := range dividend {
			left_dividend = append(left_dividend, val-to_sub[idx])
		}

		for idx, val := range left_dividend {
			if val != float64(0) {
				left_dividend = left_dividend[idx:]
				break
			}
		}

		dividend = left_dividend
		p_deg = len(dividend) - 1
	}

	rem := dividend

	quo_Poly_list := ConvertToInt(quo)
	rem_Poly_list := ConvertToInt(rem)

	return NewPoly(quo_Poly_list), NewPoly(rem_Poly_list)
}

/*
Performs Polynomial Multiplication around the a polynomial as module
Argument :
	p,q : polynomial to be multiplied
	poly_mod : polynomial used as modulo
returns :
	result of polynomial multipication of p and q around poly_mod as modulo
*/
func PolyMulMod(p, q, poly_mod Poly) Poly {
	ans := PolyMul(p, q)
	_, rem := PolyDiv(ans, poly_mod)
	return rem
}

/*
Performs Polynomial Multiplication around the a polynomial as module and also
Coeffecient modulo around given mod
Argument :
	p,q : polynomial to be multiplied
	poly_mod : polynomial used as modulo
returns :
	result of polynomial multipication of p and q around poly_mod as modulo and
	coeffecient modulo with mod
*/
func PolyMulModCoeffMod(p, q, poly_mod Poly, mod int) Poly {
	rem := PolyMulMod(p, q, poly_mod)
	rem.ModElem(mod)
	return rem
}

/*
Performs Polynomial Addition around the a polynomial as module
Argument :
	p,q : polynomial to be added
	poly_mod : polynomial used as modulo
returns :
	result of polynomial addition of p and q around poly_mod as modulo
*/
func PolyAddMod(p, q, poly_mod Poly) Poly {
	ans := PolyAdd(p, q)
	_, rem := PolyDiv(ans, poly_mod)
	return rem
}

/*
Performs Polynomial Addition around the a polynomial as module and also
Coeffecient modulo around given mod
Argument :
	p,q : polynomial to be added
	poly_mod : polynomial used as modulo
returns :
	result of polynomial Addition of p and q around poly_mod as modulo and
	coeffecient modulo with mod
*/
func PolyAddModCoeffMod(p, q, poly_mod Poly, mod int) Poly {
	rem := PolyAddMod(p, q, poly_mod)
	rem.ModElem(mod)
	return rem
}

/*
Generates a binary polynomial with given degree
*/
func GenBinaryPoly(deg int) Poly {
	coeff_list := []int{}
	for idx := 0; idx <= deg; idx++ {
		coeff_list = append(coeff_list, rand.Intn(2))
	}
	return NewPoly(coeff_list)
}

/*
Generates a uniform polynomial with given degree and coeffecient between
0 and mod-1
*/
func GenUniformPoly(deg int, modulus int) Poly {
	coeff_list := []int{}

	for idx := 0; idx <= deg; idx++ {
		coeff_list = append(coeff_list, int(rand.Float64())%modulus)
	}
	return NewPoly(coeff_list)
}

/*
Generates a normal polynomial with given mean and standard deviation
*/
func GenNormalPoly(deg int, mean int, std int) Poly {
	coeff_list := []int{}

	for idx := 0; idx <= deg; idx++ {
		rand_num := rand.NormFloat64()
		rand_num *= float64(std)
		rand_num += float64(mean)
		coeff_list = append(coeff_list, int(math.Round(rand_num)))
	}
	return NewPoly(coeff_list)
}
