package strategy

import "github.com/chengshusss/iter-prisoner-dilemma/common"

type TitForTat struct {
	OppoActMap map[int]common.ActEnum
	Idx        int
}

func (t *TitForTat) Act(oppo int) common.ActEnum {
	act, ok := t.OppoActMap[oppo]
	if !ok {
		return common.ActCooperea
	}

	return act
}

func (t *TitForTat) Update(oppoIndex int, opponentAct common.ActEnum) {
	t.OppoActMap[oppoIndex] = opponentAct
}

func (t *TitForTat) Index() int {
	return t.Idx
}

func NewTitForTat(idx int) *TitForTat {
	return &TitForTat{
		OppoActMap: make(map[int]common.ActEnum),
		Idx:        idx,
	}
}
