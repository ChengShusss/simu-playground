package playground

import (
	"github.com/chengshusss/iter-prisoner-dilemma/common"
)

// StrategyEntry is actual player in iteration-prisoner-dilemma,
// which means its imple should act to specific opponent, updating
// act history or reaction memory at meanwhile.
type StrategyEntry interface {
	Act(opponent int) common.ActEnum
	Update(opponentIndex int, opponentAct common.ActEnum)
	Index() int
}
