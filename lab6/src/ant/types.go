package ant

// Env - enviroments struct
type Env struct {
	weight   [][]int
	pheromon [][]float64
	alpha    float64 // приоретет пути
	betta    float64 // приоретет феромона
	q        float64 // переносимый муравьем феромона
	p        float64 // коэффициент испарения феромона
}

// Ant - ant struct
type Ant struct {
	env      *Env
	visited  [][]int
	been     [][]bool
	position int
}
