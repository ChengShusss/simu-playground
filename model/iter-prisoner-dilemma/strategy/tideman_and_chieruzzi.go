package strategy

import "github.com/chengshusss/iter-prisoner-dilemma/common"

type Tideman struct {
	DefeatResidueMap map[int]int
	Idx              int
}

func (t *Tideman) Act(oppo int, _ int) common.ActEnum {
	act, ok := t.DefeatResidueMap[oppo]
	if !ok {
		return common.ActCooperate
	}
	if act > 0 {
		t.DefeatResidueMap[oppo] -= 1
		return common.ActDefeat
	}

	return common.ActCooperate
}

func (t *Tideman) Update(_, oppoIndex int, opponentAct common.ActEnum) {
	if opponentAct == common.ActDefeat {
		t.DefeatResidueMap[oppoIndex] += 2
	}
}

func (t *Tideman) Index() int {
	return t.Idx
}

func (t *Tideman) Reset() {
	n := len(t.DefeatResidueMap)
	t.DefeatResidueMap = make(map[int]int, n)
}

func NewTideman(idx int, _ int) *Tideman {
	return &Tideman{
		DefeatResidueMap: make(map[int]int),
		Idx:              idx,
	}
}
