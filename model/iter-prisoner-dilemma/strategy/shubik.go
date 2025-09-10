package strategy

import "github.com/chengshusss/iter-prisoner-dilemma/common"

type Shubik struct {
	DefeatResidueMap map[int]int
	Idx              int
}

func (s *Shubik) Act(oppo int, _ int) common.ActEnum {
	act, ok := s.DefeatResidueMap[oppo]
	if !ok {
		return common.ActCooperate
	}
	if act > 0 {
		s.DefeatResidueMap[oppo] -= 1
		return common.ActDefeat
	}

	return common.ActCooperate
}

func (s *Shubik) Update(_, oppoIndex int, opponentAct common.ActEnum) {
	if opponentAct == common.ActDefeat {
		s.DefeatResidueMap[oppoIndex] += 2
	}
}

func (s *Shubik) Index() int {
	return s.Idx
}

func (s *Shubik) Reset() {
	n := len(s.DefeatResidueMap)
	s.DefeatResidueMap = make(map[int]int, n)
}

func NewShubik(idx int, _ int) *Shubik {
	return &Shubik{
		DefeatResidueMap: make(map[int]int),
		Idx:              idx,
	}
}
