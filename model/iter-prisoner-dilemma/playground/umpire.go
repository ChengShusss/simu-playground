package playground

import "github.com/chengshusss/iter-prisoner-dilemma/common"

var resultMap = map[common.ActEnum]map[common.ActEnum][2]int{
	common.ActCooperea: {
		common.ActCooperea: [2]int{3, 3},
		common.ActDefeat:   [2]int{0, 5},
	},
	common.ActDefeat: {
		common.ActCooperea: [2]int{5, 0},
		common.ActDefeat:   [2]int{1, 1},
	},
}

// Umpire is designed to judege result for two player. It would conduct
// one round game, include tell two player who is faced, gather act
// from them and give result to each.
type Umpire struct {
	PlayerA StrategyEntry
	PlayerB StrategyEntry
}

func NewUmpire(playerA, playerB StrategyEntry) *Umpire {
	return &Umpire{
		PlayerA: playerA,
		PlayerB: playerB,
	}
}

// Note: ConductOnce is not idempotent since strategy entry would
// update self after conduct
func (u *Umpire) ConductOnce() (int, int) {
	actA := u.PlayerA.Act(u.PlayerB.Index())
	actB := u.PlayerB.Act(u.PlayerA.Index())

	u.PlayerA.Update(u.PlayerB.Index(), actB)
	u.PlayerB.Update(u.PlayerA.Index(), actA)

	reward := resultMap[actA][actB]

	return reward[0], reward[1]
}
