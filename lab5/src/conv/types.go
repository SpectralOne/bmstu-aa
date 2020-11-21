package conv

import "time"

type Cake struct {
	num       int
	prepare   bool
	bake      bool
	finalize  bool
	sPrepare  time.Time
	fPrepare  time.Time
	sBake     time.Time
	fBake     time.Time
	sFinalize time.Time
	fFinalize time.Time
}

type Queue struct {
	q [](*Cake)
	l int
}
