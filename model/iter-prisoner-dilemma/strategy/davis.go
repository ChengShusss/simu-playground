package strategy

import "github.com/chengshusss/iter-prisoner-dilemma/common"

type Davis struct {
	OppoActMap map[int]struct{}
	Idx        int
}

func (d *Davis) Act(oppo int, _ int) common.ActEnum {
	_, ok := d.OppoActMap[oppo]
	if ok {
		return common.ActDefeat
	}
	return common.ActCooperate
}

func (d *Davis) Update(currentRound, oppoIndex int, opponentAct common.ActEnum) {
	if currentRound < 10 {
		return
	}
	if opponentAct == common.ActDefeat {
		d.OppoActMap[oppoIndex] = struct{}{}
	}
}

func (d *Davis) Index() int {
	return d.Idx
}

func (d *Davis) Reset() {
	n := len(d.OppoActMap)
	d.OppoActMap = make(map[int]struct{}, n)
}

func NewDavis(idx int, _ int) *Davis {
	return &Davis{
		OppoActMap: make(map[int]struct{}),
		Idx:        idx,
	}
}
