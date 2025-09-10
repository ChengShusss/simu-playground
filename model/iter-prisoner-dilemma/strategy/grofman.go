package strategy

import (
	"math/rand"

	"github.com/chengshusss/iter-prisoner-dilemma/common"
)

type Grofman struct {
	Idx int

	// Different from TitForTat, this map store act of oppo
	// and self in last round
	OppoActMap map[int][2]common.ActEnum
}

func (g *Grofman) Act(oppo int, _ int) common.ActEnum {
	actPair, ok := g.OppoActMap[oppo]
	act := common.ActCooperate
	if ok && actPair[0] != actPair[1] && rand.Intn(7) >= 2 {
		act = common.ActDefeat
	}

	actPair[0] = act
	g.OppoActMap[oppo] = actPair
	return act
}

func (g *Grofman) Update(_, oppoIndex int, oppoAct common.ActEnum) {
	actPair := g.OppoActMap[oppoIndex]
	actPair[1] = oppoAct
	g.OppoActMap[oppoIndex] = actPair
}

func (g *Grofman) Index() int {
	return g.Idx
}

func (g *Grofman) Reset() {
	g.OppoActMap = make(map[int][2]common.ActEnum)
}

func NewGrofman(idx int, totalRound int) *Grofman {
	return &Grofman{
		Idx:        idx,
		OppoActMap: make(map[int][2]common.ActEnum),
	}
}
