package multiplication

import (
	"testing"
	"../matrix"
)

// Std benchmarks

func BenchmarkStandart3x3(b *testing.B) {
	m1 := matrix.Zeros(3, 3)
	m2 := matrix.Zeros(3, 3)

	for i := 0; i < b.N; i++ {
		Multiply(m1, m2)
	}
}

func BenchmarkStandart9x9(b *testing.B) {
	m1 := matrix.Zeros(9, 9)
	m2 := matrix.Zeros(9, 9)

	for i := 0; i < b.N; i++ {
		Multiply(m1, m2)
	}
}

func BenchmarkStandart4x4(b *testing.B) {
	m1 := matrix.Zeros(4, 4)
	m2 := matrix.Zeros(4, 4)

	for i := 0; i < b.N; i++ {
		Multiply(m1, m2)
	}
}

func BenchmarkStandart8x8(b *testing.B) {
	m1 := matrix.Zeros(8, 8)
	m2 := matrix.Zeros(8, 8)
	for i := 0; i < b.N; i++ {
		Multiply(m1, m2)
	}
}

// Winograd benchmarks

func BenchmarkWinograd3x3(b *testing.B) {
	m1 := matrix.Zeros(3, 3)
	m2 := matrix.Zeros(3, 3)

	for i := 0; i < b.N; i++ {
		Winograd(m1, m2)
	}
}

func BenchmarkWinograd9x9(b *testing.B) {
	m1 := matrix.Zeros(9, 9)
	m2 := matrix.Zeros(9, 9)

	for i := 0; i < b.N; i++ {
		Winograd(m1, m2)
	}
}

func BenchmarkWinograd4x4(b *testing.B) {
	m1 := matrix.Zeros(4, 4)
	m2 := matrix.Zeros(4, 4)

	for i := 0; i < b.N; i++ {
		Winograd(m1, m2)
	}
}

func BenchmarkWinograd8x8(b *testing.B) {
	m1 := matrix.Zeros(8, 8)
	m2 := matrix.Zeros(8, 8)
	for i := 0; i < b.N; i++ {
		Winograd(m1, m2)
	}
}

// ImpWinograd benchmarks

func BenchmarkImpWinograd3x3(b *testing.B) {
	m1 := matrix.Zeros(3, 3)
	m2 := matrix.Zeros(3, 3)

	for i := 0; i < b.N; i++ {
		ImpWinograd(m1, m2)
	}
}

func BenchmarkImpWinograd9x9(b *testing.B) {
	m1 := matrix.Zeros(9, 9)
	m2 := matrix.Zeros(9, 9)
	for i := 0; i < b.N; i++ {
		ImpWinograd(m1, m2)
	}
}

func BenchmarkImpWinograd4x4(b *testing.B) {
	m1 := matrix.Zeros(4, 4)
	m2 := matrix.Zeros(4, 4)

	for i := 0; i < b.N; i++ {
		ImpWinograd(m1, m2)
	}
}

func BenchmarkImpWinograd8x8(b *testing.B) {
	m1 := matrix.Zeros(8, 8)
	m2 := matrix.Zeros(8, 8)
	for i := 0; i < b.N; i++ {
		ImpWinograd(m1, m2)
	}
}
//------------------------------------------------------------
func BenchmarkClassic100x100(b *testing.B) {
	m1 := matrix.Zeros(100, 100)
	m2 := matrix.Zeros(100, 100)
	for i := 0; i < b.N; i++ {
		Multiply(m1, m2)
	}
}

func BenchmarkClassic200x200(b *testing.B) {
	m1 := matrix.Zeros(200, 200)
	m2 := matrix.Zeros(200, 200)
	for i := 0; i < b.N; i++ {
		Multiply(m1, m2)
	}
}

func BenchmarkClassic300x300(b *testing.B) {
	m1 := matrix.Zeros(300,300)
	m2 := matrix.Zeros(300,300)
	for i := 0; i < b.N; i++ {
		Multiply(m1, m2)
	}
}

func BenchmarkClassic400x400(b *testing.B) {
	m1 := matrix.Zeros(400, 400)
	m2 := matrix.Zeros(400, 400)
	for i := 0; i < b.N; i++ {
		Multiply(m1, m2)
	}
}

func BenchmarkClassic500x500(b *testing.B) {
	m1 := matrix.Zeros(500, 500)
	m2 := matrix.Zeros(500, 500)
	for i := 0; i < b.N; i++ {
		Multiply(m1, m2)
	}
}

func BenchmarkClassic600x600(b *testing.B) {
	m1 := matrix.Zeros(600, 600)
	m2 := matrix.Zeros(600, 600)
	for i := 0; i < b.N; i++ {
		Multiply(m1, m2)
	}
}
//------------------------------------------------------------
func BenchmarkWinograd100x100(b *testing.B) {
	m1 := matrix.Zeros(100, 100)
	m2 := matrix.Zeros(100, 100)
	for i := 0; i < b.N; i++ {
		Winograd(m1, m2)
	}
}

func BenchmarkWinograd200x200(b *testing.B) {
	m1 := matrix.Zeros(200, 200)
	m2 := matrix.Zeros(200, 200)
	for i := 0; i < b.N; i++ {
		Winograd(m1, m2)
	}
}

func BenchmarkWinograd300x300(b *testing.B) {
	m1 := matrix.Zeros(300,300)
	m2 := matrix.Zeros(300,300)
	for i := 0; i < b.N; i++ {
		Winograd(m1, m2)
	}
}

func BenchmarkWinograd400x400(b *testing.B) {
	m1 := matrix.Zeros(400, 400)
	m2 := matrix.Zeros(400, 400)
	for i := 0; i < b.N; i++ {
		Winograd(m1, m2)
	}
}

func BenchmarkWinograd500x500(b *testing.B) {
	m1 := matrix.Zeros(500, 500)
	m2 := matrix.Zeros(500, 500)
	for i := 0; i < b.N; i++ {
		Winograd(m1, m2)
	}
}

func BenchmarkWinograd600x600(b *testing.B) {
	m1 := matrix.Zeros(600, 600)
	m2 := matrix.Zeros(600, 600)
	for i := 0; i < b.N; i++ {
		Winograd(m1, m2)
	}
}
//------------------------------------------------------------
func BenchmarkImpWinograd100x100(b *testing.B) {
	m1 := matrix.Zeros(100, 100)
	m2 := matrix.Zeros(100, 100)
	for i := 0; i < b.N; i++ {
		ImpWinograd(m1, m2)
	}
}

func BenchmarkImpWinograd200x200(b *testing.B) {
	m1 := matrix.Zeros(200, 200)
	m2 := matrix.Zeros(200, 200)
	for i := 0; i < b.N; i++ {
		ImpWinograd(m1, m2)
	}
}

func BenchmarkImpWinograd300x300(b *testing.B) {
	m1 := matrix.Zeros(300,300)
	m2 := matrix.Zeros(300,300)
	for i := 0; i < b.N; i++ {
		ImpWinograd(m1, m2)
	}
}

func BenchmarkImpWinograd400x400(b *testing.B) {
	m1 := matrix.Zeros(400, 400)
	m2 := matrix.Zeros(400, 400)
	for i := 0; i < b.N; i++ {
		ImpWinograd(m1, m2)
	}
}

func BenchmarkImpWinograd500x500(b *testing.B) {
	m1 := matrix.Zeros(500, 500)
	m2 := matrix.Zeros(500, 500)
	for i := 0; i < b.N; i++ {
		ImpWinograd(m1, m2)
	}
}

func BenchmarkImpWinograd600x600(b *testing.B) {
	m1 := matrix.Zeros(600, 600)
	m2 := matrix.Zeros(600, 600)
	for i := 0; i < b.N; i++ {
		ImpWinograd(m1, m2)
	}
}
//--
//------------------------------------------------------------
func BenchmarkClassic101x101(b *testing.B) {
	m1 := matrix.Zeros(101, 101)
	m2 := matrix.Zeros(101, 101)
	for i := 0; i < b.N; i++ {
		Multiply(m1, m2)
	}
}

func BenchmarkClassic201x201(b *testing.B) {
	m1 := matrix.Zeros(201, 201)
	m2 := matrix.Zeros(201, 201)
	for i := 0; i < b.N; i++ {
		Multiply(m1, m2)
	}
}

func BenchmarkClassic301x301(b *testing.B) {
	m1 := matrix.Zeros(301,301)
	m2 := matrix.Zeros(301,301)
	for i := 0; i < b.N; i++ {
		Multiply(m1, m2)
	}
}

func BenchmarkClassic401x401(b *testing.B) {
	m1 := matrix.Zeros(401, 401)
	m2 := matrix.Zeros(401, 401)
	for i := 0; i < b.N; i++ {
		Multiply(m1, m2)
	}
}

func BenchmarkClassic501x501(b *testing.B) {
	m1 := matrix.Zeros(501, 501)
	m2 := matrix.Zeros(501, 501)
	for i := 0; i < b.N; i++ {
		Multiply(m1, m2)
	}
}

func BenchmarkClassic601x601(b *testing.B) {
	m1 := matrix.Zeros(601, 601)
	m2 := matrix.Zeros(601, 601)
	for i := 0; i < b.N; i++ {
		Multiply(m1, m2)
	}
}
//------------------------------------------------------------
func BenchmarkWinograd101x101(b *testing.B) {
	m1 := matrix.Zeros(101, 101)
	m2 := matrix.Zeros(101, 101)
	for i := 0; i < b.N; i++ {
		Winograd(m1, m2)
	}
}

func BenchmarkWinograd201x201(b *testing.B) {
	m1 := matrix.Zeros(201, 201)
	m2 := matrix.Zeros(201, 201)
	for i := 0; i < b.N; i++ {
		Winograd(m1, m2)
	}
}

func BenchmarkWinograd301x301(b *testing.B) {
	m1 := matrix.Zeros(301,301)
	m2 := matrix.Zeros(301,301)
	for i := 0; i < b.N; i++ {
		Winograd(m1, m2)
	}
}

func BenchmarkWinograd401x401(b *testing.B) {
	m1 := matrix.Zeros(401, 401)
	m2 := matrix.Zeros(401, 401)
	for i := 0; i < b.N; i++ {
		Winograd(m1, m2)
	}
}

func BenchmarkWinograd501x501(b *testing.B) {
	m1 := matrix.Zeros(501, 501)
	m2 := matrix.Zeros(501, 501)
	for i := 0; i < b.N; i++ {
		Winograd(m1, m2)
	}
}

func BenchmarkWinograd601x601(b *testing.B) {
	m1 := matrix.Zeros(601, 601)
	m2 := matrix.Zeros(601, 601)
	for i := 0; i < b.N; i++ {
		Winograd(m1, m2)
	}
}
//------------------------------------------------------------
func BenchmarkImpWinograd101x101(b *testing.B) {
	m1 := matrix.Zeros(101, 101)
	m2 := matrix.Zeros(101, 101)
	for i := 0; i < b.N; i++ {
		ImpWinograd(m1, m2)
	}
}

func BenchmarkImpWinograd201x201(b *testing.B) {
	m1 := matrix.Zeros(201, 201)
	m2 := matrix.Zeros(201, 201)
	for i := 0; i < b.N; i++ {
		ImpWinograd(m1, m2)
	}
}

func BenchmarkImpWinograd301x301(b *testing.B) {
	m1 := matrix.Zeros(301,301)
	m2 := matrix.Zeros(301,301)
	for i := 0; i < b.N; i++ {
		ImpWinograd(m1, m2)
	}
}

func BenchmarkImpWinograd401x401(b *testing.B) {
	m1 := matrix.Zeros(401, 401)
	m2 := matrix.Zeros(401, 401)
	for i := 0; i < b.N; i++ {
		ImpWinograd(m1, m2)
	}
}

func BenchmarkImpWinograd501x501(b *testing.B) {
	m1 := matrix.Zeros(501, 501)
	m2 := matrix.Zeros(501, 501)
	for i := 0; i < b.N; i++ {
		ImpWinograd(m1, m2)
	}
}

func BenchmarkImpWinograd601x601(b *testing.B) {
	m1 := matrix.Zeros(601, 601)
	m2 := matrix.Zeros(601, 601)
	for i := 0; i < b.N; i++ {
		ImpWinograd(m1, m2)
	}
}
