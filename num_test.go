package num_test

import (
	"testing"

	"github.com/sp301415/num"
)

func TestAbs(t *testing.T) {
	testCases := []struct {
		x, want int64
	}{
		{0, 0},
		{1, 1},
		{-1, 1},
	}
	for _, test := range testCases {
		if got := num.Abs(test.x); got != test.want {
			t.Errorf("Abs(%d) = %d; want %d", test.x, got, test.want)
		}
	}
}

func TestBit(t *testing.T) {
	testCases := []struct {
		x       int64
		i, want int
	}{
		{0b0, 0, 0},
		{0b10000, 4, 1},
		{0b101010, 4, 0},
	}
	for _, test := range testCases {
		if got := num.Bit(test.x, test.i); got != test.want {
			t.Errorf("Bit(%d, %d) = %d; want %d", test.x, test.i, got, test.want)
		}
	}
}

func TestBitPanic(t *testing.T) {
	failureMsg := "test failure"
	defer func() {
		msg := recover()
		if msg == nil || msg == failureMsg {
			panic(msg)
		}
		t.Log(msg)
	}()
	num.Bit(0, -1)
	panic(failureMsg)
}

func TestCmp(t *testing.T) {
	testCases := []struct {
		x, y int64
		want int
	}{
		{0, 0, 0},
		{-1, 1, -1},
		{1, -1, 1},
	}
	for _, test := range testCases {
		if got := num.Cmp(test.x, test.y); got != test.want {
			t.Errorf("Cmp(%d, %d) = %d; want %d", test.x, test.y, got, test.want)
		}
	}
}

func TestCmpAbs(t *testing.T) {
	testCases := []struct {
		x, y int64
		want int
	}{
		{0, 0, 0},
		{-1, 1, 0},
		{1, -1, 0},
		{1, -2, -1},
		{-2, 1, 1},
	}
	for _, test := range testCases {
		if got := num.CmpAbs(test.x, test.y); got != test.want {
			t.Errorf("CmpAbs(%d, %d) = %d; want %d", test.x, test.y, got, test.want)
		}
	}
}

func TestPow(t *testing.T) {
	testCases := []struct {
		x, y, want int64
	}{
		{0, 0, 1},
		{0, 1, 0},
		{1, 0, 1},
		{1, 1, 1},
		{2, 10, 1024},
		{-2, 10, 1024},
		{3, 4, 81},
		{10, 2, 100},
	}
	for _, test := range testCases {
		if got := num.Pow(test.x, test.y); got != test.want {
			t.Errorf("Pow(%d, %d) = %d; want %d", test.x, test.y, got, test.want)
		}
	}
}

func TestPowPanic(t *testing.T) {
	failureMsg := "test failure"
	defer func() {
		msg := recover()
		if msg == nil || msg == failureMsg {
			panic(msg)
		}
		t.Log(msg)
	}()
	num.Pow(1, -1)
	panic(failureMsg)
}

func TestPowMod(t *testing.T) {
	testCases := []struct {
		x, y, m, want int64
	}{
		{0, 0, 1, 0}, // 0**0 mod 1 = 1 mod 1 = 0
		{1, 1, 1, 0}, // 1**1 mod 1 = 1 mod 1 = 0
		{1, 1, 2, 1}, // 1**1 mod 2 = 1 mod 2 = 1
		{2, 2, 1, 0}, // 4 mod 1 = 0
		{3, 2, 1, 0}, // 9 mod 1 = 0
		{12345, 12345, 4, 1},
	}
	for _, test := range testCases {
		if got := num.PowMod(test.x, test.y, test.m); got != test.want {
			t.Errorf("PowMod(%d, %d, %d) = %d; want %d", test.x, test.y, test.m, got, test.want)
		}
	}
}

func TestPowModPanicZero(t *testing.T) {
	failureMsg := "test failure"
	defer func() {
		msg := recover()
		if msg == nil || msg == failureMsg {
			panic(msg)
		}
		t.Log(msg)
	}()
	num.PowMod(1, 1, 0)
	panic(failureMsg)
}

func TestPowModPanicNegative(t *testing.T) {
	failureMsg := "test failure"
	defer func() {
		msg := recover()
		if msg == nil || msg == failureMsg {
			panic(msg)
		}
		t.Log(msg)
	}()
	num.PowMod(1, -1, 2)
	panic(failureMsg)
}

func TestGCD(t *testing.T) {
	testCases := []struct {
		x, y, want int64
	}{
		{0, 0, 0},
		{0, 2, 2},
		{2, 0, 2},
		{2, 3, 1},
		{3, 6, 3},
	}
	for _, test := range testCases {
		if got := num.GCD(test.x, test.y); got != test.want {
			t.Errorf("GCD(%d, %d) = %d; want %d", test.x, test.y, got, test.want)
		}
	}
}

func TestXGCD(t *testing.T) {
	testCases := []struct {
		x, y, r int64
		s, t    int
	}{
		{0, 0, 0, 0, 0},
		{0, 2, 2, 0, 1},
		{0, -2, 2, 0, -1},
		{2, 0, 2, 1, 0},
		{-2, 0, 2, -1, 0},
		{240, 46, 2, -9, 47},
	}
	for _, test := range testCases {
		if got_r, got_s, got_t := num.XGCD(test.x, test.y); got_r != test.r || got_s != test.s || got_t != test.t {
			t.Errorf("XGCD(%d, %d) = %d, %d, %d; want %d, %d, %d", test.x, test.y, got_r, got_s, got_t, test.r, test.s, test.t)
		}
	}
}

func TestMax(t *testing.T) {
	testCases := []struct {
		x, y, want int64
	}{
		{0, 0, 0},
		{1, 0, 1},
		{-1, 0, 0},
	}
	for _, test := range testCases {
		if got := num.Max(test.x, test.y); got != test.want {
			t.Errorf("Max(%d, %d) = %d; want %d", test.x, test.y, got, test.want)
		}
	}
}

func TestMin(t *testing.T) {
	testCases := []struct {
		x, y, want int64
	}{
		{0, 0, 0},
		{1, 0, 0},
		{-1, 0, -1},
	}
	for _, test := range testCases {
		if got := num.Min(test.x, test.y); got != test.want {
			t.Errorf("Min(%d, %d) = %d; want %d", test.x, test.y, got, test.want)
		}
	}
}

func TestIsPrime(t *testing.T) {
	testCases := []struct {
		x    int64
		want bool
	}{
		{0, false},
		{1, false},
		{2, true},
		{997, true},
	}
	for _, test := range testCases {
		if got := num.IsPrime(test.x); got != test.want {
			t.Errorf("IsPrime(%d) = %t; want %t", test.x, got, test.want)
		}
	}
}

func TestSign(t *testing.T) {
	testCases := []struct {
		x    int64
		want int
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
	}
	for _, test := range testCases {
		if got := num.Sign(test.x); got != test.want {
			t.Errorf("Sign(%d) = %d; want %d", test.x, got, test.want)
		}
	}
}

func TestSqrt(t *testing.T) {
	testCases := []struct {
		x, want int64
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{2000000, 1414},
	}
	for _, test := range testCases {
		if got := num.Sqrt(test.x); got != test.want {
			t.Errorf("Sqrt(%d) = %d; want %d", test.x, got, test.want)
		}
	}
}

func TestSqrtPanic(t *testing.T) {
	failureMsg := "test failure"
	defer func() {
		msg := recover()
		if msg == nil || msg == failureMsg {
			panic(msg)
		}
		t.Log(msg)
	}()
	num.Sqrt(-1)
	panic(failureMsg)
}
