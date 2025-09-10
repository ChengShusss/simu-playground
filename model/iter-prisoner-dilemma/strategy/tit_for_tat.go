package strategy

import "github.com/chengshusss/iter-prisoner-dilemma/common"

type Tit struct {
	OppoActMap map[int]common.ActEnum
	Idx        int
}

func (t *Tit) Act(oppo int, _ int) common.ActEnum {
	act, ok := t.OppoActMap[oppo]
	if !ok {
		return common.ActCooperate
	}

	return act
}

func (t *Tit) Update(_, oppoIndex int, opponentAct common.ActEnum) {
	t.OppoActMap[oppoIndex] = opponentAct
}

func (t *Tit) Index() int {
	return t.Idx
}

func (t *Tit) Reset() {
	n := len(t.OppoActMap)
	t.OppoActMap = make(map[int]common.ActEnum, n)
}

func NewTitForTat(idx int, _ int) *Tit {
	return &Tit{
		OppoActMap: make(map[int]common.ActEnum),
		Idx:        idx,
	}
}
