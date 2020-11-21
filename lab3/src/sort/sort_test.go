package sort

import (
	"testing"
)

// Special case benchs

func BenchmarkBubble1(b *testing.B) {
	a := Generate(1);

	for i := 0; i < b.N; i++ {
		BubbleSort(a)
	}
}

func BenchmarkInsert1(b *testing.B) {
	a := Generate(1);

	for i := 0; i < b.N; i++ {
		InsertSort(a)
	}
}

func BenchmarkQuick1(b *testing.B) {
	a := Generate(1);

	for i := 0; i < b.N; i++ {
		QuickSort(a)
	}
}

func BenchmarkBubble10(b *testing.B) {
	a := Generate(10);

	for i := 0; i < b.N; i++ {
		BubbleSort(a)
	}
}

func BenchmarkInsert10(b *testing.B) {
	a := Generate(10);

	for i := 0; i < b.N; i++ {
		InsertSort(a)
	}
}

func BenchmarkQuick10(b *testing.B) {
	a := Generate(10);

	for i := 0; i < b.N; i++ {
		QuickSort(a)
	}
}

func BenchmarkBubble100(b *testing.B) {
	a := Generate(100);

	for i := 0; i < b.N; i++ {
		BubbleSort(a)
	}
}

func BenchmarkInsert100(b *testing.B) {
	a := Generate(100);

	for i := 0; i < b.N; i++ {
		InsertSort(a)
	}
}

func BenchmarkQuick100(b *testing.B) {
	a := Generate(100);

	for i := 0; i < b.N; i++ {
		QuickSort(a)
	}
}

func BenchmarkBubble1k(b *testing.B) {
	a := Generate(1000);

	for i := 0; i < b.N; i++ {
		BubbleSort(a)
	}
}

func BenchmarkInsert1k(b *testing.B) {
	a := Generate(1000);

	for i := 0; i < b.N; i++ {
		InsertSort(a)
	}
}

func BenchmarkQuick1k(b *testing.B) {
	a := Generate(1000);

	for i := 0; i < b.N; i++ {
		QuickSort(a)
	}
}

func BenchmarkBubble10k(b *testing.B) {
	a := Generate(10000);

	for i := 0; i < b.N; i++ {
		BubbleSort(a)
	}
}

func BenchmarkInsert10k(b *testing.B) {
	a := Generate(10000);

	for i := 0; i < b.N; i++ {
		InsertSort(a)
	}
}

func BenchmarkQuick10k(b *testing.B) {
	a := Generate(10000);

	for i := 0; i < b.N; i++ {
		QuickSort(a)
	}
}

func BenchmarkBubble100k(b *testing.B) {
	a := Generate(100000);

	for i := 0; i < b.N; i++ {
		BubbleSort(a)
	}
}

func BenchmarkInsert100k(b *testing.B) {
	a := Generate(100000);

	for i := 0; i < b.N; i++ {
		InsertSort(a)
	}
}

func BenchmarkQuick100k(b *testing.B) {
	a := Generate(100000);

	for i := 0; i < b.N; i++ {
		QuickSort(a)
	}
}

// Regular benchs

func BenchmarkBubble10kRandom(b *testing.B) {
	a := CreateRandom(10000);

	for i := 0; i < b.N; i++ {
		BubbleSort(a)
	}
}

func BenchmarkBubble20kRandom(b *testing.B) {
	a := CreateRandom(20000);

	for i := 0; i < b.N; i++ {
		BubbleSort(a)
	}
}

func BenchmarkBubble30kRandom(b *testing.B) {
	a := CreateRandom(30000);

	for i := 0; i < b.N; i++ {
		BubbleSort(a)
	}
}

func BenchmarkBubble40kRandom(b *testing.B) {
	a := CreateRandom(40000);

	for i := 0; i < b.N; i++ {
		BubbleSort(a)
	}
}

func BenchmarkBubble50kRandom(b *testing.B) {
	a := CreateRandom(50000);

	for i := 0; i < b.N; i++ {
		BubbleSort(a)
	}
}

// BubbleSort - sorted

func BenchmarkBubble10kSorted(b *testing.B) {
	a := CreateSorted(10000, false);

	for i := 0; i < b.N; i++ {
		BubbleSort(a)
	}
}

func BenchmarkBubble20kSorted(b *testing.B) {
	a := CreateSorted(20000, false);

	for i := 0; i < b.N; i++ {
		BubbleSort(a)
	}
}

func BenchmarkBubble30kSorted(b *testing.B) {
	a := CreateSorted(30000, false);

	for i := 0; i < b.N; i++ {
		BubbleSort(a)
	}
}

func BenchmarkBubble40kSorted(b *testing.B) {
	a := CreateSorted(40000, false);

	for i := 0; i < b.N; i++ {
		BubbleSort(a)
	}
}

func BenchmarkBubble50kSorted(b *testing.B) {
	a := CreateSorted(50000, false);

	for i := 0; i < b.N; i++ {
		BubbleSort(a)
	}
}

// BubbleSort - revers

func BenchmarkBubble10kReverse(b *testing.B) {
	a := CreateSorted(10000, true);

	for i := 0; i < b.N; i++ {
		BubbleSort(a)
	}
}

func BenchmarkBubble20kReverse(b *testing.B) {
	a := CreateSorted(20000, true);

	for i := 0; i < b.N; i++ {
		BubbleSort(a)
	}
}

func BenchmarkBubble30kReverse(b *testing.B) {
	a := CreateSorted(30000, true);

	for i := 0; i < b.N; i++ {
		BubbleSort(a)
	}
}

func BenchmarkBubble40kReverse(b *testing.B) {
	a := CreateSorted(40000, true);

	for i := 0; i < b.N; i++ {
		BubbleSort(a)
	}
}

func BenchmarkBubble50kReverse(b *testing.B) {
	a := CreateSorted(50000, true);

	for i := 0; i < b.N; i++ {
		BubbleSort(a)
	}
}

// InsertSort - random

func BenchmarkInsert10kRandom(b *testing.B) {
	a := CreateRandom(10000);

	for i := 0; i < b.N; i++ {
		InsertSort(a)
	}
}

func BenchmarkInsert20kRandom(b *testing.B) {
	a := CreateRandom(20000);

	for i := 0; i < b.N; i++ {
		InsertSort(a)
	}
}

func BenchmarkInsert30kRandom(b *testing.B) {
	a := CreateRandom(30000);

	for i := 0; i < b.N; i++ {
		InsertSort(a)
	}
}

func BenchmarkInsert40kRandom(b *testing.B) {
	a := CreateRandom(40000);

	for i := 0; i < b.N; i++ {
		InsertSort(a)
	}
}

func BenchmarkInsert50kRandom(b *testing.B) {
	a := CreateRandom(50000);

	for i := 0; i < b.N; i++ {
		InsertSort(a)
	}
}

// InsertSort - sorted

func BenchmarkInsert10kSorted(b *testing.B) {
	a := CreateSorted(10000, false);

	for i := 0; i < b.N; i++ {
		InsertSort(a)
	}
}

func BenchmarkInsert20kSorted(b *testing.B) {
	a := CreateSorted(20000, false);

	for i := 0; i < b.N; i++ {
		InsertSort(a)
	}
}

func BenchmarkInsert30kSorted(b *testing.B) {
	a := CreateSorted(30000, false);

	for i := 0; i < b.N; i++ {
		InsertSort(a)
	}
}

func BenchmarkInsert40kSorted(b *testing.B) {
	a := CreateSorted(40000, false);

	for i := 0; i < b.N; i++ {
		InsertSort(a)
	}
}

func BenchmarkInsert50kSorted(b *testing.B) {
	a := CreateSorted(50000, false);

	for i := 0; i < b.N; i++ {
		InsertSort(a)
	}
}

// InsertSort - reverse

func BenchmarkInsert10kReverse(b *testing.B) {
	a := CreateSorted(10000, true);

	for i := 0; i < b.N; i++ {
		InsertSort(a)
	}
}

func BenchmarkInsert20kReverse(b *testing.B) {
	a := CreateSorted(20000, true);

	for i := 0; i < b.N; i++ {
		InsertSort(a)
	}
}

func BenchmarkInsert30kReverse(b *testing.B) {
	a := CreateSorted(30000, true);

	for i := 0; i < b.N; i++ {
		InsertSort(a)
	}
}

func BenchmarkInsert40kReverse(b *testing.B) {
	a := CreateSorted(40000, true);

	for i := 0; i < b.N; i++ {
		InsertSort(a)
	}
}

func BenchmarkInsert50kReverse(b *testing.B) {
	a := CreateSorted(50000, true);

	for i := 0; i < b.N; i++ {
		InsertSort(a)
	}
}

// QuickSort - Random

func BenchmarkQuick10kRandom(b *testing.B) {
	a := CreateRandom(10000);

	for i := 0; i < b.N; i++ {
		QuickSort(a)
	}
}

func BenchmarkQuick20kRandom(b *testing.B) {
	a := CreateRandom(20000);

	for i := 0; i < b.N; i++ {
		QuickSort(a)
	}
}

func BenchmarkQuick30kRandom(b *testing.B) {
	a := CreateRandom(30000);

	for i := 0; i < b.N; i++ {
		QuickSort(a)
	}
}

func BenchmarkQuick40kRandom(b *testing.B) {
	a := CreateRandom(40000);

	for i := 0; i < b.N; i++ {
		QuickSort(a)
	}
}

func BenchmarkQuick50kRandom(b *testing.B) {
	a := CreateRandom(50000);

	for i := 0; i < b.N; i++ {
		QuickSort(a)
	}
}

// QuickSort - sorted

func BenchmarkQuick10kSorted(b *testing.B) {
	a := CreateSorted(10000, false);

	for i := 0; i < b.N; i++ {
		QuickSort(a)
	}
}

func BenchmarkQuick20kSorted(b *testing.B) {
	a := CreateSorted(20000, false);

	for i := 0; i < b.N; i++ {
		QuickSort(a)
	}
}

func BenchmarkQuick30kSorted(b *testing.B) {
	a := CreateSorted(30000, false);

	for i := 0; i < b.N; i++ {
		QuickSort(a)
	}
}

func BenchmarkQuick40kSorted(b *testing.B) {
	a := CreateSorted(40000, false);

	for i := 0; i < b.N; i++ {
		QuickSort(a)
	}
}

func BenchmarkQuick50kSorted(b *testing.B) {
	a := CreateSorted(50000, false);

	for i := 0; i < b.N; i++ {
		QuickSort(a)
	}
}

// QuickSort - reverse

func BenchmarkQuick10kReverse(b *testing.B) {
	a := CreateSorted(10000, true);

	for i := 0; i < b.N; i++ {
		QuickSort(a)
	}
}

func BenchmarkQuick20kReverse(b *testing.B) {
	a := CreateSorted(20000, true);

	for i := 0; i < b.N; i++ {
		QuickSort(a)
	}
}

func BenchmarkQuick30kReverse(b *testing.B) {
	a := CreateSorted(30000, true);

	for i := 0; i < b.N; i++ {
		QuickSort(a)
	}
}

func BenchmarkQuick40kReverse(b *testing.B) {
	a := CreateSorted(40000, true);

	for i := 0; i < b.N; i++ {
		QuickSort(a)
	}
}

func BenchmarkQuick50kReverse(b *testing.B) {
	a := CreateSorted(50000, true);

	for i := 0; i < b.N; i++ {
		QuickSort(a)
	}
}
