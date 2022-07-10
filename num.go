// package num implements various integer functions, analogous to math/big.
package num

import (
	"math/big"

	"golang.org/x/exp/constraints"
)

// Abs returns the absolute value of x.
func Abs[T constraints.Integer](x T) T {
	if x > 0 {
		return x
	}
	return -x
}

// Bit returns the value of the i'th bit of x.
func Bit[T constraints.Integer](x T, i int) int {
	return int((uint64(x) >> i) & 1)
}

// Cmp compares x and y and returns:
//
//	-1 if x < y
//	 0 if x = y
//	+1 if x > y
//
func Cmp[T constraints.Integer](x, y T) int {
	if x < y {
		return -1
	} else if x > y {
		return 1
	}
	return 0
}

// CmpAbs compares the absolute values of x and y and returns:
//
//	-1 if |x| < |y|
//	 0 if |x| =	|y|
//	+1 if |x| > |y|
//
func CmpAbs[T constraints.Integer](x, y T) int {
	ax := Abs(x)
	ay := Abs(y)
	if ax < ay {
		return -1
	} else if ax > ay {
		return 1
	}
	return 0
}

// Pow returns x**y. If x == y == 0 or y < 0, a runtime panic occurs.
func Pow[T constraints.Integer](x, y T) T {
	if x == 0 && y == 0 {
		return 1
	}
	if y < 0 {
		panic("negative exponent")
	}

	switch x {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1 << y
	}

	var r T = 1
	for y > 0 {
		if y&1 == 1 {
			r *= x
		}
		x *= x
		y >>= 1
	}
	return r
}

// PowMod returns x**y mod m.
// If m == 0, it is equivalant to Pow(x, y).
// If x == y == 0 or y < 0, a runtime panic occurs.
func PowMod[T constraints.Integer](x, y, m T) T {
	if y < 0 {
		panic("negative exponent")
	}
	if m == 0 {
		panic("modulo by zero")
	}

	switch x {
	case 0:
		return 0
	case 1:
		if m == 1 {
			return 0
		} else {
			return 1
		}
	case 2:
		return (1 << y) % m
	}

	var r T = 1
	for y > 0 {
		if y&1 == 1 {
			r = (r * x) % m
		}
		x = (x * x) % m
		y >>= 1
	}
	return r
}

// GCD returns the greatest common divisor of x and y.
//
// If x == y == 0, GCD returns 0.
// If x == 0 and y != 0, GCD returns |y|.
// If x != 0 and y == 0, GCD returns |x|.
func GCD[T constraints.Integer](x, y T) T {
	x = Abs(x)
	y = Abs(y)

	if x == 0 && y == 0 {
		return 0
	} else if x == 0 {
		return y
	} else if y == 0 {
		return x
	}

	for y > 0 {
		x, y = y, x%y
	}
	return x
}

// XGCD returns the greatest common divisor of x and y, along with a and b, such that a*x+b*y==g.
//
// If x == y == 0, g equals 0.
//
// If x == 0 and y != 0, g equals |y|.
//
// If x != 0 and y == 0, g equals |x|.
func XGCD[T constraints.Integer](x, y T) (T, int, int) {
	if x == 0 && y == 0 {
		return 0, 0, 0
	} else if x == 0 {
		return Abs(y), 0, Sign(y)
	} else if y == 0 {
		return Abs(x), Sign(x), 0
	}

	// Taken from WikiPedia: https://en.wikipedia.org/wiki/Extended_Euclidean_algorithm#Pseudocode
	var oldr, r T = x, y
	var olds, s T = 1, 0
	var oldt, t T = 0, 1

	for r != 0 {
		var q T = oldr / r
		oldr, r = r, oldr-q*r
		olds, s = s, olds-q*s
		oldt, t = t, oldt-q*t
	}

	return oldr, int(olds), int(oldt)
}

// Max returns the larger integer between x and y.
func Max[T constraints.Integer](x, y T) T {
	if x > y {
		return x
	}
	return y
}

// Min returns the smaller integer between x and y.
func Min[T constraints.Integer](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// IsPrime checks if |x| is prime.
// It is equivalant to ProbabalyPrime(0) in math/big, so it is guaranteed to be correct for inputs less than 2^64.
func IsPrime[T constraints.Integer](x T) bool {
	return new(big.Int).SetUint64(uint64(Abs(x))).ProbablyPrime(0)
}

// Sign returns +1 if x > 0, -1 if x < 0, 0 if x == 0.
func Sign[T constraints.Integer](x T) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}

// Sqrt returns the integer square root of x, i.e. the smallest integer i such that i*i <= x.
// If x is negative, a runtime panic occurs.
func Sqrt[T constraints.Integer](x T) T {
	if x < 0 {
		panic("square root of negative number")
	}

	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}

	// Taken from WikiPedia: https://en.wikipedia.org/wiki/Integer_square_root#Using_only_integer_division
	var x0 T = x / 2
	var x1 T = (x0 + x/x0) / 2

	for x1 < x0 {
		x0 = x1
		x1 = (x0 + x/x0) / 2
	}

	return x0
}
