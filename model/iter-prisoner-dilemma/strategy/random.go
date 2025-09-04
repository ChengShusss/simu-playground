package strategy

import (
	"math/rand"

	"github.com/chengshusss/iter-prisoner-dilemma/common"
)

type Random struct {
	Idx int
}

func (t *Random) Act(oppo int) common.ActEnum {
	if rand.Intn(2)%2 == 0 {
		return common.ActCooperea
	}
	return common.ActDefeat
}

func (t *Random) Update(_ int, _ common.ActEnum) {
	// Do nothing
}

func (t *Random) Index() int {
	return t.Idx
}

func NewRandom(idx int) *Random {
	return &Random{
		Idx: idx,
	}
}
