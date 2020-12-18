package conv

import "time"

type Person struct {
	num        int
	haveName   bool
	haveEmail  bool
	haveBeer   bool
	startName  time.Time
	doneName   time.Time
	startEmail time.Time
	doneEmail  time.Time
	startBeer  time.Time
	doneBeer   time.Time

	name  string
	email string
	beer  string
}

type Queue struct {
	q [](*Person)
	l int
}