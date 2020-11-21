package levenshtein

import (
	"testing"
)

// LevRec method benchmarks.

func BenchmarkRecursiveLen5(b *testing.B) {
	s1 := "about"
	s2 := "above"

	for i := 0; i < b.N; i++ {
		LevRec(s1, s2)
	}
}

func BenchmarkRecursiveLen10(b *testing.B) {
	s1 := "abbanition"
	s2 := "abaptiston"

	for i := 0; i < b.N; i++ {
		LevRec(s1, s2)
	}
}

func BenchmarkRecursiveLen15(b *testing.B) {
	s1 := "characteristics"
	s2 := "recommendations"

	for i := 0; i < b.N; i++ {
		LevRec(s1, s2)
	}
}

// LevMtrRec method becnhmarks.

func BenchmarkRecursiveMatrixLen10(b *testing.B) {
	s1 := "abbanition"
	s2 := "abaptiston"

	for i := 0; i < b.N; i++ {
		LevMtrRec(s1, s2)
	}
}

func BenchmarkRecursiveMatrixLen20(b *testing.B) {
	s1 := "abdominohysterectomy"
	s2 := "acetylcholinesterase"

	for i := 0; i < b.N; i++ {
		LevMtrRec(s1, s2)
	}
}
func BenchmarkRecursiveMatrixLen30(b *testing.B) {
	s1 := "chlorobenzylidenemalononitrile"
	s2 := "abdominalexternalobliquemuscle"

	for i := 0; i < b.N; i++ {
		LevMtrRec(s1, s2)
	}
}

func BenchmarkRecursiveMatrixLen50(b *testing.B) {
	s1 := "chlorobenzylidenemalononitrileabdominohysterectomy"
	s2 := "abdominalexternalobliquemuscleacetylcholinesterase"

	for i := 0; i < b.N; i++ {
		LevMtrRec(s1, s2)
	}
}

func BenchmarkRecursiveMatrixLen100(b *testing.B) {
	s1 := "chlorobenzylidenemalononitrileabdominohysterectomy" +
		"chlorobenzylidenemalononitrileabdominohysterectomy"
	s2 := "abdominalexternalobliquemuscleacetylcholinesterase" +
		"abdominalexternalobliquemuscleacetylcholinesterase"

	for i := 0; i < b.N; i++ {
		LevMtrRec(s1, s2)
	}
}

func BenchmarkRecursiveMatrixLen200(b *testing.B) {
	s1 := "chlorobenzylidenemalononitrileabdominohysterectomy" +
		"chlorobenzylidenemalononitrileabdominohysterectomy" +
		"chlorobenzylidenemalononitrileabdominohysterectomy" +
		"chlorobenzylidenemalononitrileabdominohysterectomy"
	s2 := "abdominalexternalobliquemuscleacetylcholinesterase" +
		"abdominalexternalobliquemuscleacetylcholinesterase" +
		"abdominalexternalobliquemuscleacetylcholinesterase" +
		"abdominalexternalobliquemuscleacetylcholinesterase"

	for i := 0; i < b.N; i++ {
		LevMtrRec(s1, s2)
	}
}

// LevMtr method benchmarks.

func BenchmarkIterativeMatrixLen10(b *testing.B) {
	s1 := "abbanition"
	s2 := "abaptiston"

	for i := 0; i < b.N; i++ {
		LevMtr(s1, s2)
	}
}

func BenchmarkIterativeMatrixLen20(b *testing.B) {
	s1 := "abdominohysterectomy"
	s2 := "acetylcholinesterase"

	for i := 0; i < b.N; i++ {
		LevMtr(s1, s2)
	}
}

func BenchmarkIterativeMatrixLen30(b *testing.B) {
	s1 := "chlorobenzylidenemalononitrile"
	s2 := "abdominalexternalobliquemuscle"

	for i := 0; i < b.N; i++ {
		LevMtr(s1, s2)
	}
}

func BenchmarkIterativeMatrixLen50(b *testing.B) {
	s1 := "chlorobenzylidenemalononitrileabdominohysterectomy"
	s2 := "abdominalexternalobliquemuscleacetylcholinesterase"

	for i := 0; i < b.N; i++ {
		LevMtr(s1, s2)
	}
}

func BenchmarkIterativeMatrixLen100(b *testing.B) {
	s1 := "chlorobenzylidenemalononitrileabdominohysterectomy" +
		"chlorobenzylidenemalononitrileabdominohysterectomy"
	s2 := "abdominalexternalobliquemuscleacetylcholinesterase" +
		"abdominalexternalobliquemuscleacetylcholinesterase"

	for i := 0; i < b.N; i++ {
		LevMtr(s1, s2)
	}
}

func BenchmarkIterativeMatrixLen200(b *testing.B) {
	s1 := "chlorobenzylidenemalononitrileabdominohysterectomy" +
		"chlorobenzylidenemalononitrileabdominohysterectomy" +
		"chlorobenzylidenemalononitrileabdominohysterectomy" +
		"chlorobenzylidenemalononitrileabdominohysterectomy"
	s2 := "abdominalexternalobliquemuscleacetylcholinesterase" +
		"abdominalexternalobliquemuscleacetylcholinesterase" +
		"abdominalexternalobliquemuscleacetylcholinesterase" +
		"abdominalexternalobliquemuscleacetylcholinesterase"

	for i := 0; i < b.N; i++ {
		LevMtr(s1, s2)
	}
}

// DamLevMtr method benchmarks.

func BenchmarkDamerauLevenshteinLen10(b *testing.B) {
	s1 := "abbanition"
	s2 := "abaptiston"

	for i := 0; i < b.N; i++ {
		DamLevMtr(s1, s2)
	}
}

func BenchmarkDamerauLevenshteinLen20(b *testing.B) {
	s1 := "abdominohysterectomy"
	s2 := "acetylcholinesterase"

	for i := 0; i < b.N; i++ {
		DamLevMtr(s1, s2)
	}
}

func BenchmarkDamerauLevenshteinLen30(b *testing.B) {
	s1 := "chlorobenzylidenemalononitrile"
	s2 := "abdominalexternalobliquemuscle"

	for i := 0; i < b.N; i++ {
		DamLevMtr(s1, s2)
	}
}

func BenchmarkDamerauLevenshteinLen50(b *testing.B) {
	s1 := "chlorobenzylidenemalononitrileabdominohysterectomy"
	s2 := "abdominalexternalobliquemuscleacetylcholinesterase"

	for i := 0; i < b.N; i++ {
		DamLevMtr(s1, s2)
	}
}

func BenchmarkDamerauLevenshteinLen100(b *testing.B) {
	s1 := "chlorobenzylidenemalononitrileabdominohysterectomy" +
		"chlorobenzylidenemalononitrileabdominohysterectomy"
	s2 := "abdominalexternalobliquemuscleacetylcholinesterase" +
		"abdominalexternalobliquemuscleacetylcholinesterase"

	for i := 0; i < b.N; i++ {
		DamLevMtr(s1, s2)
	}
}

func BenchmarkDamerauLevenshteinLen200(b *testing.B) {
	s1 := "chlorobenzylidenemalononitrileabdominohysterectomy" +
		"chlorobenzylidenemalononitrileabdominohysterectomy" +
		"chlorobenzylidenemalononitrileabdominohysterectomy" +
		"chlorobenzylidenemalononitrileabdominohysterectomy"
	s2 := "abdominalexternalobliquemuscleacetylcholinesterase" +
		"abdominalexternalobliquemuscleacetylcholinesterase" +
		"abdominalexternalobliquemuscleacetylcholinesterase" +
		"abdominalexternalobliquemuscleacetylcholinesterase"

	for i := 0; i < b.N; i++ {
		DamLevMtr(s1, s2)
	}
}
