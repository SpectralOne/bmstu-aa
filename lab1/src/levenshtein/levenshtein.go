package levenshtein

// LevRec - recursive levenstein
func LevRec(a, b string) int {
	s1, s2 := []rune(a), []rune(b)

	lenS1, lenS2 := len(s1), len(s2)
	// swap to save some memory O(min(a,b)) instead of O(a)
	if lenS1 > lenS2 {
		s1, s2 = s2, s1
		lenS1, lenS2 = lenS2, lenS1
	}

	return getDistRec(s1, s2, lenS1, lenS2)
}

// LevMtr - levenstein with matrix
func LevMtr(a, b string) (d int, mtr [][]int) {
	s1, s2 := []rune(a), []rune(b)
	lenS1, lenS2 := len(s1)+1, len(s2)+1

	mtr = createMatrix(lenS1, lenS2, false)

	penalty := 0
	for i := 1; i < lenS1; i++ {
		for j := 1; j < lenS2; j++ {
			if s1[i-1] == s2[j-1] {
				penalty = 0
			} else {
				penalty = 1
			}

			mtr[i][j] = minOf(
				mtr[i-1][j]+1,
				mtr[i][j-1]+1,
				mtr[i-1][j-1]+penalty)
		}
	}
	d = mtr[lenS1-1][lenS2-1]

	return
}

// LevMtrRec - Recursive Levenstein with matrix
func LevMtrRec(a, b string) (d int, mtr [][]int) {
	s1, s2 := []rune(a), []rune(b)
	lenS1, lenS2 := len(s1), len(s2)

	mtr = createMatrix(lenS1, lenS2, true)
	d = getDistMtr(s1, s2, lenS1, lenS2, mtr)

	return
}

// DamLevMtr - Damerau - Levenshtein with matrix
func DamLevMtr(a, b string) (d int, mtr [][]int) {
	s1, s2 := []rune(a), []rune(b)
	lenS1, lenS2 := len(s1)+1, len(s2)+1

	mtr = createMatrix(lenS1, lenS2, false)

	penalty := 0
	for i := 1; i < lenS1; i++ {
		for j := 1; j < lenS2; j++ {
			if s1[i-1] == s2[j-1] {
				penalty = 0
			} else {
				penalty = 1
			}
			mtr[i][j] = minOf(
				mtr[i-1][j]+1,
				mtr[i][j-1]+1,
				mtr[i-1][j-1]+penalty)

			if i > 1 && j > 1 &&
				s1[i-1] == s2[j-2] &&
				s1[i-2] == s2[j-1] {
				mtr[i][j] = minOf(mtr[i][j], mtr[i-2][j-2]+1)
			}
		}
	}

	d = mtr[lenS1-1][lenS2-1]

	return
}

func getDistRec(a, b []rune, lenA, lenB int) int {
	if lenA == 0 {
		return lenB
	}
	if lenB == 0 && lenA > 0 {
		return lenA
	}
	penalty := 1
	if a[lenA-1] == b[lenB-1] {
		penalty = 0
	}

	return minOf(
		getDistRec(a, b, lenA, lenB-1)+1,
		getDistRec(a, b, lenA-1, lenB)+1,
		getDistRec(a, b, lenA-1, lenB-1)+penalty)
}

func getDistMtr(a, b []rune, lenA, lenB int, mtr [][]int) int {
	if lenA == 0 {
		return lenB
	}
	if lenB == 0 {
		return lenA
	}

	if mtr[lenA-1][lenB-1] != -1 {
		return mtr[lenA-1][lenB-1]
	}

	if a[lenA-1] == b[lenB-1] {
		mtr[lenA-1][lenB-1] = getDistMtr(a, b, lenA-1, lenB-1, mtr)
		return mtr[lenA-1][lenB-1]
	}

	mtr[lenA-1][lenB-1] = 1 + minOf(
		getDistMtr(a, b, lenA, lenB-1, mtr),
		getDistMtr(a, b, lenA-1, lenB, mtr),
		getDistMtr(a, b, lenA-1, lenB-1, mtr))

	return mtr[lenA-1][lenB-1]
}
