package strategy

import (
	"github.com/chengshusss/iter-prisoner-dilemma/common"
	"github.com/chengshusss/iter-prisoner-dilemma/utils"
)

type Nydegger struct {
	Idx        int
	TotalRound int
	Memory     map[int][3][2]common.ActEnum
}

var defeatScoreList = []int{
	1, 6, 7, 17, 22, 23, 26, 29, 30, 31, 33, 38, 39, 45, 49, 54, 55, 58, 61,
}

func (n *Nydegger) Act(oppo int, currentRound int) common.ActEnum {

	history, ok := n.Memory[oppo]
	if !ok {
		history = [3][2]common.ActEnum{}
	}

	act := common.ActCooperate

	switch currentRound {
	case 0:
		// The first move of nydegger must be cooperate
		history[0][0] = common.ActCooperate
	case 1:
		act = history[0][1]
		history[1][0] = act
	case 2:
		//
		act = history[1][1]
		if history[0][1] == common.ActDefeat && history[1][1] == common.ActCooperate &&
			history[0][0] == common.ActCooperate && history[1][0] == common.ActDefeat {
			// It would choose defeat in this situation
			act = common.ActDefeat
		}
		history[2][0] = act
	default:
		score := getNydeggerScore(history)
		if utils.Contain(defeatScoreList, score) {
			act = common.ActDefeat
		}
		// forgive last third memory and put current act
		history[0] = history[1]
		history[1] = history[2]
		history[2][0] = act
	}

	n.Memory[oppo] = history

	return act
}

func (n *Nydegger) Update(currentRound, oppoIndex int, oppoAct common.ActEnum) {
	history := n.Memory[oppoIndex]
	if currentRound <= 1 {
		history[currentRound][1] = oppoAct
	} else {
		history[2][1] = oppoAct
	}
	n.Memory[oppoIndex] = history
}

func (n *Nydegger) Index() int {
	return n.Idx
}

func (n *Nydegger) Reset() {
	n.Memory = make(map[int][3][2]common.ActEnum)
}

func getNydeggerScore(memory [3][2]common.ActEnum) int {
	score := 0

	for i := 0; i < 3; i++ {
		weight := 16 >> (i * 2)
		if memory[i][0] == common.ActDefeat {
			score += weight
		}
		if memory[i][1] == common.ActDefeat {
			score += weight * 2
		}
	}

	return score
}

func NewNydegger(idx, totalRound int) *Nydegger {
	return &Nydegger{
		Idx:        idx,
		Memory:     make(map[int][3][2]common.ActEnum),
		TotalRound: totalRound,
	}
}
