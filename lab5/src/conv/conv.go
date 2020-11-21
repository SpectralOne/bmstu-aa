package conv

import (
	"fmt"
	"time"
)

func createQueue(amount int) *Queue {
	q := new(Queue)
	q.q = make([](*Cake), amount, amount)
	q.l = -1
	return q
}
func (q *Queue) push(cake *Cake) {
	if q.l != len(q.q)-1 {
		q.q[q.l+1] = cake
		q.l++
	}
}
func (q *Queue) pop() *Cake {
	cake := q.q[0]
	q.q = q.q[1:]
	q.l--
	return cake
}

func first(cake *Cake, qBake *Queue) {
	cake.prepare = true

	cake.sPrepare = time.Now()
	time.Sleep(time.Duration(200) * time.Millisecond)
	cake.fPrepare = time.Now()

	qBake.push(cake)
}

func second(cake *Cake, qFinalize *Queue) {
	cake.bake = true

	cake.sBake = time.Now()
	time.Sleep(time.Duration(200) * time.Millisecond)
	cake.fBake = time.Now()

	qFinalize.push(cake)
}

func third(cake *Cake, finished *Queue) {
	cake.finalize = true

	cake.sFinalize = time.Now()
	time.Sleep(time.Duration(200) * time.Millisecond)
	cake.fFinalize = time.Now()

	finished.push(cake)
}

func Parallel(amount int, wait chan int) *Queue {
	f := make(chan *Cake, 5)
	s := make(chan *Cake, 5)
	t := make(chan *Cake, 5)
	line := createQueue(amount)
	first := func() {
		for {
			select {
			case cake := <-f:
				cake.prepare = true

				cake.sPrepare = time.Now()
				time.Sleep(time.Duration(200) * time.Millisecond)
				cake.fPrepare = time.Now()

				s <- cake
			}
		}
	}

	second := func() {
		for {
			select {
			case cake := <-s:
				cake.bake = true

				cake.sBake = time.Now()
				time.Sleep(time.Duration(200) * time.Millisecond)
				cake.fBake = time.Now()

				t <- cake
			}
		}
	}

	third := func() {
		for {
			select {
			case cake := <-t:
				cake.finalize = true

				cake.sFinalize = time.Now()
				time.Sleep(time.Duration(200) * time.Millisecond)
				cake.fFinalize = time.Now()

				line.push(cake)
				if cake.num == amount {
					wait <- 0
				}

			}
		}
	}

	go first()
	go second()
	go third()
	for i := 0; i <= amount; i++ {
		cake := new(Cake)
		cake.num = i
		f <- cake
	}

	return line
}

func Linear(amount int) *Queue {
	qBake := createQueue(amount)
	qFinalize := createQueue(amount)
	finished := createQueue(amount)
	i := 0
	for i != -1 {
		cake := new(Cake)
		cake.num = i
		first(cake, qBake)
		if qBake.l >= 0 {
			second(qBake.pop(), qFinalize)
		}
		if qFinalize.l >= 0 {
			third(qFinalize.pop(), finished)
		}
		if finished.q[len(finished.q)-1] != nil {
			return finished
		}
		i++
	}
	return finished
}

func Log(q *Queue) {
	var (
		fDowntime time.Duration
		sDowntime time.Duration
		tDowntime time.Duration
	)

	line := q.q
	start := line[0].sPrepare
	fmt.Printf("Starting time\n")
	for i := 0; i < len(line); i++ {
		if line[i] != nil {
			fmt.Println(i, line[i].sPrepare.Sub(start), line[i].sBake.Sub(start), line[i].sFinalize.Sub(start))
		}
	}

	fmt.Printf("Finishing time\n")
	for i := 0; i < len(line); i++ {
		if line[i] != nil {
			fmt.Println(i, line[i].fPrepare.Sub(start), line[i].fBake.Sub(start), line[i].fFinalize.Sub(start))
		}
	}

	fmt.Printf("Линии простаивали\n")
	for i := 0; i < len(line)-1; i++ {
		fDowntime += line[i+1].sPrepare.Sub(start) - line[i].fPrepare.Sub(start)
		sDowntime += line[i+1].sBake.Sub(start) - line[i].fBake.Sub(start)
		tDowntime += line[i+1].sFinalize.Sub(start) - line[i].fFinalize.Sub(start)
	}
	fmt.Println(fDowntime, sDowntime, tDowntime)
}
