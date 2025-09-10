package strategy

import (
	"math/rand"

	"github.com/chengshusss/iter-prisoner-dilemma/common"
)

type Random struct {
	Idx int
}

func (t *Random) Act(oppo int, _ int) common.ActEnum {
	if rand.Intn(2)%2 == 0 {
		return common.ActCooperate
	}
	return common.ActDefeat
}

func (t *Random) Update(_, _ int, _ common.ActEnum) {
	// Do nothing
}

func (t *Random) Index() int {
	return t.Idx
}

func (t *Random) Reset() {
	// Do nothing
}

func NewRandom(idx int, totalRound int) *Random {
	return &Random{
		Idx: idx,
	}
}
