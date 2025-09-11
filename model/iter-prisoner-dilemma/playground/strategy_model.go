package playground

import (
	"github.com/chengshusss/iter-prisoner-dilemma/common"
	"github.com/chengshusss/iter-prisoner-dilemma/strategy"
)

// StrategyEntry is actual player in iteration-prisoner-dilemma,
// which means its imple should act to specific opponent, updating
// act history or reaction memory at meanwhile.
type StrategyEntry interface {
	Act(opponent, currentRound int) common.ActEnum
	Update(currentRound, opponentIndex int, opponentAct common.ActEnum)
	Index() int
	Reset()
}

func getStrategyFromName(name string, idx, totalRound int) StrategyEntry {
	switch name {
	case "tit":
		return strategy.NewTitForTat(idx, totalRound)
	case "tideman":
		return strategy.NewTideman(idx, totalRound)
	case "nydegger":
		return strategy.NewNydegger(idx, totalRound)
	case "grofman":
		return strategy.NewGrofman(idx, totalRound)
	case "shubik":
		return strategy.NewShubik(idx, totalRound)
	case "stein":
		return strategy.NewStein(idx, totalRound)
	case "friedman":
		return strategy.NewFriedman(idx, totalRound)
	case "davis":
		return strategy.NewDavis(idx, totalRound)

	case "random":
		return strategy.NewRandom(idx, totalRound)
	default:
		return nil
	}
}
