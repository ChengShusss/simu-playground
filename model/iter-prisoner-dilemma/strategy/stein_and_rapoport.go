package strategy

// This file is generate by deepseek
// Awesome

import (
	"math"

	"github.com/chengshusss/iter-prisoner-dilemma/common"
)

// Stein 策略实现
type Stein struct {
	index       int
	totalRounds int
	history     map[int][]common.ActEnum // 存储每个对手的历史行动
}

// Act 方法决定当前回合的行动。
// opponent: 对手索引，currentRound: 当前回合数（从1开始）。
func (s *Stein) Act(opponent, currentRound int) common.ActEnum {
	// 前四步合作
	if currentRound <= 3 {
		return common.ActCooperate
	}
	// 最后两步背叛
	if currentRound >= s.totalRounds-2 {
		return common.ActDefeat
	}
	// 每15步检查对手是否随机
	if currentRound%15 == 14 {
		oppHistory, exists := s.history[opponent]
		if !exists || len(oppHistory) < 2 {
			// 历史不足，使用 tit for tat
			return s.titForTat(opponent)
		}
		// 检查随机性
		if s.isRandom(oppHistory) {
			return common.ActDefeat
		}
	}
	// 其他情况使用 tit for tat
	return s.titForTat(opponent)
}

// Update 方法更新对手的历史行动。
// currentRound: 当前回合数，opponentIndex: 对手索引，opponentAct: 对手行动。
func (s *Stein) Update(currentRound, opponentIndex int, opponentAct common.ActEnum) {
	s.history[opponentIndex] = append(s.history[opponentIndex], opponentAct)
}

// Index 返回策略索引。
func (s *Stein) Index() int {
	return s.index
}

// Reset 重置策略状态，清空所有历史。
func (s *Stein) Reset() {
	s.history = make(map[int][]common.ActEnum)
}

// titForTat 根据对手上一次行动决定行动。
func (s *Stein) titForTat(opponent int) common.ActEnum {
	oppHistory, exists := s.history[opponent]
	if !exists || len(oppHistory) == 0 {
		return common.ActCooperate
	}
	// 返回对手最后一次行动
	return oppHistory[len(oppHistory)-1]
}

// isRandom 检查对手历史行动是否随机。
// 使用卡方检验转移概率和交替模式检查。
func (s *Stein) isRandom(history []common.ActEnum) bool {
	n := len(history)
	if n < 2 {
		return false
	}
	// 计算转移次数：CC, CD, DC, DD
	var CC, CD, DC, DD int
	for i := 0; i < n-1; i++ {
		prev := history[i]
		next := history[i+1]
		if prev == common.ActCooperate && next == common.ActCooperate {
			CC++
		} else if prev == common.ActCooperate && next == common.ActDefeat {
			CD++
		} else if prev == common.ActDefeat && next == common.ActCooperate {
			DC++
		} else if prev == common.ActDefeat && next == common.ActDefeat {
			DD++
		}
	}
	totalTrans := n - 1
	// 计算卡方值
	E := float64(totalTrans) / 4.0
	chi2 := math.Pow(float64(CC)-E, 2)/E +
		math.Pow(float64(CD)-E, 2)/E +
		math.Pow(float64(DC)-E, 2)/E +
		math.Pow(float64(DD)-E, 2)/E
	// 卡方临界值（自由度=3，显著性水平0.05）
	criticalValue := 7.815
	if chi2 > criticalValue {
		// 卡方检验显著，不随机
		return false
	}
	// 检查交替模式比例（CD或DC转移）
	altTrans := CD + DC
	altRatio := float64(altTrans) / float64(totalTrans)
	// 如果交替比例超过90%，则认为不随机
	return altRatio <= 0.9
}

// NewStein 构造函数，用于创建 Stein 实例。
// index: 策略索引，totalRounds: 总回合数。
func NewStein(index int, totalRounds int) *Stein {
	return &Stein{
		index:       index,
		totalRounds: totalRounds,
		history:     make(map[int][]common.ActEnum),
	}
}
