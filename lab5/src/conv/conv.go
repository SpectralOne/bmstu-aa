package conv

import (
	"fmt"
	"time"
	"math/big"
	"github.com/brianvoe/gofakeit"
)

// Represents random computation distribution
// such as random seed generator,
// unique check time and so.
// Education purpose only...
func randomOps() {
	_ = new(big.Int).MulRange(1, 136637) 
}

func createQueue(amount int) *Queue {
	q := new(Queue)
	q.q = make([](*Person), amount, amount)
	q.l = -1
	return q
}

func (q *Queue) push(p *Person) {
	if q.l != len(q.q)-1 {
		q.q[q.l+1] = p
		q.l++
	}
}
func (q *Queue) pop() *Person {
	p := q.q[0]
	q.q = q.q[1:]
	q.l--
	return p
}

func name(p *Person, qEmail *Queue) {
	p.haveName = true

	p.startName = time.Now()
	randomOps()
	p.name = gofakeit.Name()
	randomOps()
	p.doneName = time.Now()

	qEmail.push(p)
}

func email(p *Person, qBeer *Queue) {
	p.haveEmail = true

	p.startEmail = time.Now()
	randomOps()
	p.email = gofakeit.Email()
	randomOps()
	p.doneEmail = time.Now()

	qBeer.push(p)
}

func beer(p *Person, finished *Queue) {
	p.haveBeer = true

	p.startBeer = time.Now()
	randomOps()
	p.beer = gofakeit.BeerName()
	randomOps()
	p.doneBeer = time.Now()

	finished.push(p)
}

func Linear(amount int) *Queue {
	qEmail := createQueue(amount)
	qBeer := createQueue(amount)
	finished := createQueue(amount)
	i := 0
	for i != -1 {
		p := new(Person)
		p.num = i
		name(p, qEmail)
		if qEmail.l >= 0 {
			email(qEmail.pop(), qBeer)
		}
		if qBeer.l >= 0 {
			beer(qBeer.pop(), finished)
		}
		if finished.q[len(finished.q)-1] != nil {
			return finished
		}
		i++
	}
	return finished
}

func Parallel(amount int, wait chan int) *Queue {
	f := make(chan *Person, 5)
	s := make(chan *Person, 5)
	t := make(chan *Person, 5)
	line := createQueue(amount)
	name := func() {
		for {
			select {
			case p := <-f:
				p.haveName = true

				p.startName = time.Now()
				randomOps()
				p.name = gofakeit.Name()
				randomOps()
				p.doneName = time.Now()

				s <- p
			}
		}
	}

	email := func() {
		for {
			select {
			case p := <-s:
				p.haveEmail = true

				p.startEmail = time.Now()
				randomOps()
				p.email = gofakeit.Email()
				randomOps()
				p.doneEmail = time.Now()

				t <- p
			}
		}
	}

	beer := func() {
		for {
			select {
			case p := <-t:
				p.haveBeer = true

				p.startBeer = time.Now()
				randomOps()
				p.beer = gofakeit.BeerName()
				randomOps()
				p.doneBeer = time.Now()

				line.push(p)
				if p.num == amount {
					wait <- 0
				}

			}
		}
	}

	go name()
	go email()
	go beer()
	for i := 0; i <= amount; i++ {
		p := new(Person)
		p.num = i
		f <- p
	}

	return line
}

func Log(q *Queue, logPerson bool) {
	var (
		fDowntime time.Duration
		sDowntime time.Duration
		tDowntime time.Duration
	)

	line := q.q
	start := line[0].startName
	_ = line[len(line) - 1] // omit boundary checks
	fmt.Printf("Starting time\n")
	for i := range line {
		if line[i] != nil {
			fmt.Println(i, line[i].startName.Sub(start), line[i].startEmail.Sub(start), line[i].startBeer.Sub(start))
		}
	}

	fmt.Printf("Finishing time\n")
	for i := range line {
		if line[i] != nil {
			fmt.Println(i, line[i].doneName.Sub(start), line[i].doneEmail.Sub(start), line[i].doneBeer.Sub(start))
		}
	}

	fmt.Printf("Линии простаивали\n")
	for i := 0; i < len(line)-1; i++ {
		fDowntime += line[i+1].startName.Sub(start) - line[i].doneName.Sub(start)
		sDowntime += line[i+1].startEmail.Sub(start) - line[i].doneEmail.Sub(start)
		tDowntime += line[i+1].startBeer.Sub(start) - line[i].doneBeer.Sub(start)
	}
	fmt.Println(fDowntime, sDowntime, tDowntime)

	if logPerson {
		for i := range line {
			fmt.Printf("Person: %d\nName: %s\nEmail: %s\nBeer: %s\n\n", line[i].num + 1, line[i].name, line[i].email, line[i].beer)
		}
	}
}