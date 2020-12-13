package ant

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func (e *Env) newAnt(pos int) *Ant {
	ant := new(Ant)
	ant.env = e
	ant.visited = make([][]int, len(e.weight))

	for i := 0; i < len(e.weight); i++ {
		ant.visited[i] = make([]int, len(e.weight))
		for j := 0; j < len(e.weight[i]); j++ {
			ant.visited[i][j] = e.weight[i][j]
		}
	}
	ant.position = pos
	ant.been = make([][]bool, len(e.weight))

	for i := range ant.been {
		ant.been[i] = make([]bool, len(e.weight))
	}
	return ant
}

func newEnv(f string) *Env {
	e := new(Env)
	e.weight = getWeights(f)
	e.pheromon = make([][]float64, len(e.weight), len(e.weight))
	for i := 0; i < len(e.pheromon); i++ {
		e.pheromon[i] = make([]float64, len(e.weight[i]))
		for j := range e.pheromon[i] {
			e.pheromon[i][j] = 0.5
		}
	}
	e.alpha = 3.0
	e.betta = 7.0
	e.q = 20.0
	e.p = 0.6
	return e
}

// getWeights - read graph from file
func getWeights(f string) [][]int {
	g := make([][]int, 0)
	file, err := os.Open(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		str = strings.TrimSuffix(str, "\n")
		str = strings.TrimSuffix(str, "\r")
		str = strings.TrimRight(str, " ")
		//fmt.Printf(str)
		cur := strings.Split(str, " ")
		path := make([]int, 0)
		for _, i := range cur {
			i, err := strconv.Atoi(i)
			if err != nil {
				fmt.Println(err)
			}
			path = append(path, i)
		}
		g = append(g, path)
	}
	return g
}

// getProbability - probability of path being choosen
func (ant *Ant) getProbability() []float64 {
	p := make([]float64, 0)
	var sum float64
	for i, l := range ant.visited[ant.position] {
		if l != 0 {
			d := math.Pow((float64(1)/float64(l)), ant.env.alpha) * math.Pow(ant.env.pheromon[ant.position][i], ant.env.betta)
			p = append(p, d)
			sum += d
		} else {
			p = append(p, 0)
		}
	}
	for _, l := range p {
		l = l / sum
	}
	return p
}

// pickWay - choose way with probability
func pickWay(p []float64) int {
	var sum float64
	for _, j := range p {
		sum += j
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rn := r.Float64() * sum
	sum = 0
	for i, j := range p {
		if rn > sum && rn < sum+j {
			return i
		}
		sum += j
	}
	return -1
}

// renewPheromon - renew pheromon after move
func (ant *Ant) renewPheromon() {
	var delta float64
	delta = 0
	for k := 0; k < len(ant.env.pheromon); k++ {
		for i, j := range ant.env.pheromon[k] {
			if ant.env.weight[k][i] != 0 {
				if ant.been[k][i] {
					delta = ant.env.q / float64(ant.env.weight[k][i])
				} else {
					delta = 0
				}
				ant.env.pheromon[k][i] = (1 - ant.env.p) * (float64(j) + delta)
			}
			if ant.env.pheromon[k][i] <= 0 {
				ant.env.pheromon[k][i] = 0.1
			}
		}
	}
}

// moveWay - perform move
func (ant *Ant) moveWay(path int) {
	for i := range ant.visited {
		ant.visited[i][ant.position] = 0
	}
	ant.been[ant.position][path] = true
	ant.position = path
}

// moveAnt - start moving
func (ant *Ant) moveAnt() {
	for {
		prob := ant.getProbability()
		way := pickWay(prob)
		if way == -1 {
			break
		}
		ant.moveWay(way)
		ant.renewPheromon()
	}
}

// getDistance - distance travelled
func (ant *Ant) getDistance() int {
	var distance int
	for i, j := range ant.been {
		for k, z := range j {
			if z {
				distance += ant.env.weight[i][k]
			}
		}
	}
	return distance
}

// StartAnt - ant
func StartAnt(env *Env, days int) []int {
	short := make([]int, len(env.weight))
	for n := 0; n < days; n++ {
		for i := 0; i < len(env.weight); i++ {
			ant := env.newAnt(i)
			ant.moveAnt()
			cur := ant.getDistance()
			if (short[i] == 0) || (cur < short[i]) {
				short[i] = cur
			}
		}
	}
	return short
}

// Brute - brute
func Brute(f string) []int {
	weight := getWeights(f)
	path := make([]int, 0)
	res := make([]int, len(weight))
	// for every node
	for k := 0; k < len(weight); k++ {
		ways := make([][]int, 0)
		_ = getRoutes(k, weight, path, &ways)
		sum := 1000
		curr := 0
		ind := 0
		for i := 0; i < len(ways); i++ {
			curr = 0
			for j := 0; j < len(ways[i])-1; j++ {
				curr += weight[ways[i][j]][ways[i][j+1]]
			}
			if curr < sum {
				sum = curr
				ind = i
			}
		}
		res[k] = sum
		ind = 0
		_ = ind
	}
	return res
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// getRoutes - find all ways
func getRoutes(pos int, weight [][]int, path []int, routes *[][]int) []int {
	path = append(path, pos)
	if len(path) < len(weight) {
		for i := 0; i < len(weight); i++ {
			if !(contains(path, i)) {
				_ = getRoutes(i, weight, path, routes)
			}
		}
	} else {
		*routes = append(*routes, path)
	}
	return path
}
