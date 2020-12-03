package dict

import "testing"

var d = Create(100000)
var fa = FreqAnal(d)

func BenchmarkBrute_A(b *testing.B) {
	w := pick(d, "A")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_B(b *testing.B) {
	w := pick(d, "B")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_C(b *testing.B) {
	w := pick(d, "C")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_D(b *testing.B) {
	w := pick(d, "D")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_E(b *testing.B) {
	w := pick(d, "E")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_F(b *testing.B) {
	w := pick(d, "F")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_G(b *testing.B) {
	w := pick(d, "G")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_H(b *testing.B) {
	w := pick(d, "H")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_I(b *testing.B) {
	w := pick(d, "I")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_J(b *testing.B) {
	w := pick(d, "J")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_K(b *testing.B) {
	w := pick(d, "K")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_L(b *testing.B) {
	w := pick(d, "L")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_M(b *testing.B) {
	w := pick(d, "M")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_N(b *testing.B) {
	w := pick(d, "N")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_O(b *testing.B) {
	w := pick(d, "O")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_P(b *testing.B) {
	w := pick(d, "P")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_Q(b *testing.B) {
	w := pick(d, "Q")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_R(b *testing.B) {
	w := pick(d, "R")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_S(b *testing.B) {
	w := pick(d, "S")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_T(b *testing.B) {
	w := pick(d, "T")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_U(b *testing.B) {
	w := pick(d, "U")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_V(b *testing.B) {
	w := pick(d, "V")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_W(b *testing.B) {
	w := pick(d, "W")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_X(b *testing.B) {
	w := pick(d, "X")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_Y(b *testing.B) {
	w := pick(d, "Y")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

func BenchmarkBrute_Z(b *testing.B) {
	w := pick(d, "Z")
	for i := 0; i < b.N; i++ {
		Brute(d, w)
	}
}

// Combined benchmarks

func BenchmarkCombined_A(b *testing.B) {
	w := pick(d, "A")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_B(b *testing.B) {
	w := pick(d, "B")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_C(b *testing.B) {
	w := pick(d, "C")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_D(b *testing.B) {
	w := pick(d, "D")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_E(b *testing.B) {
	w := pick(d, "E")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_F(b *testing.B) {
	w := pick(d, "F")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_G(b *testing.B) {
	w := pick(d, "G")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_H(b *testing.B) {
	w := pick(d, "H")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_I(b *testing.B) {
	w := pick(d, "I")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_J(b *testing.B) {
	w := pick(d, "J")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_K(b *testing.B) {
	w := pick(d, "K")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_L(b *testing.B) {
	w := pick(d, "L")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_M(b *testing.B) {
	w := pick(d, "M")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_N(b *testing.B) {
	w := pick(d, "N")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_O(b *testing.B) {
	w := pick(d, "O")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_P(b *testing.B) {
	w := pick(d, "P")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_Q(b *testing.B) {
	w := pick(d, "Q")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_R(b *testing.B) {
	w := pick(d, "R")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_S(b *testing.B) {
	w := pick(d, "S")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_T(b *testing.B) {
	w := pick(d, "T")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_U(b *testing.B) {
	w := pick(d, "U")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_V(b *testing.B) {
	w := pick(d, "V")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_W(b *testing.B) {
	w := pick(d, "W")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_X(b *testing.B) {
	w := pick(d, "X")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_Y(b *testing.B) {
	w := pick(d, "Y")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}

func BenchmarkCombined_Z(b *testing.B) {
	w := pick(d, "Z")
	for i := 0; i < b.N; i++ {
		Combined(fa, w)
	}
}
