package strategy

import "github.com/chengshusss/iter-prisoner-dilemma/common"

type Friedman struct {
	OppoActMap map[int]struct{}
	Idx        int
}

func (f *Friedman) Act(oppo int, _ int) common.ActEnum {
	_, ok := f.OppoActMap[oppo]
	if ok {
		return common.ActDefeat
	}
	return common.ActCooperate
}

func (f *Friedman) Update(_, oppoIndex int, opponentAct common.ActEnum) {
	if opponentAct == common.ActDefeat {
		f.OppoActMap[oppoIndex] = struct{}{}
	}
}

func (f *Friedman) Index() int {
	return f.Idx
}

func (f *Friedman) Reset() {
	n := len(f.OppoActMap)
	f.OppoActMap = make(map[int]struct{}, n)
}

func NewFriedman(idx int, _ int) *Friedman {
	return &Friedman{
		OppoActMap: make(map[int]struct{}),
		Idx:        idx,
	}
}
