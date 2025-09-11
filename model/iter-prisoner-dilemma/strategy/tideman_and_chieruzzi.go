package strategy

// This file is generate by deepseek
// Awesome

import (
	"math"

	"github.com/chengshusss/iter-prisoner-dilemma/common"
)

// Tideman 实现策略
type Tideman struct {
	index       int
	totalRounds int

	// 每个对手的状态
	opponentStates map[int]*OpponentState
}

// OpponentState 跟踪每个对手
type OpponentState struct {
	opponentHistory       []common.ActEnum
	myHistory             []common.ActEnum
	consecutiveDefections int
	runsOfDefections      int
	punishmentCount       int
	lastFreshStartRound   int
	inDefectionRun        bool
	myScore               int
	opponentScore         int
}

// Act 根据当前状态决定行动
func (s *Tideman) Act(opponentIdx, currentRound int) common.ActEnum {
	state := s.getOpponentState(opponentIdx)

	// 最后两轮自动背叛
	if currentRound >= s.totalRounds-2 {
		return common.ActDefeat
	}

	// 检查是否需要重新开始
	if s.shouldFreshStart(opponentIdx, currentRound) {
		// 重新开始: 先合作两次，然后像游戏刚开始一样
		if len(state.myHistory) == 0 ||
			(len(state.myHistory) == 1 && state.myHistory[0] == common.ActCooperate) {
			return common.ActCooperate
		}
		// 两次合作后，回到初始策略
		state.lastFreshStartRound = currentRound
		// 重置相关状态
		state.consecutiveDefections = 0
		state.runsOfDefections = 0
		state.punishmentCount = 0
		state.inDefectionRun = false
	}

	// 第一轮合作
	if len(state.opponentHistory) == 0 {
		return common.ActCooperate
	}

	// 实施惩罚逻辑
	if state.punishmentCount > 0 {
		state.punishmentCount--
		return common.ActDefeat
	}

	// 否则以牙还牙
	return state.opponentHistory[len(state.opponentHistory)-1]
}

// Update 更新对手行为记录
func (s *Tideman) Update(currentRound, opponentIndex int, opponentAct common.ActEnum) {
	state := s.getOpponentState(opponentIndex)

	// 记录对手行为
	state.opponentHistory = append(state.opponentHistory, opponentAct)

	// 更新连续背叛计数
	if opponentAct == common.ActDefeat {
		if !state.inDefectionRun {
			state.inDefectionRun = true
		}
		state.consecutiveDefections++
	} else {
		// 背叛运行结束
		if state.inDefectionRun {
			state.inDefectionRun = false
			state.runsOfDefections++

			// 当对手完成第二次连续背叛后，实施额外惩罚
			if state.runsOfDefections >= 2 {
				state.punishmentCount = state.runsOfDefections - 1
			}
		}
		state.consecutiveDefections = 0
	}

	// 更新得分（假设标准囚徒困境得分）
	myLastAct := common.ActCooperate
	if len(state.myHistory) > 0 {
		myLastAct = state.myHistory[len(state.myHistory)-1]
	}

	s.updateScores(state, myLastAct, opponentAct)
}

// shouldFreshStart 判断是否应该给对手重新开始的机会
func (s *Tideman) shouldFreshStart(opponentIdx, currentRound int) bool {
	state := s.getOpponentState(opponentIdx)

	// 对手落后10分或更多
	if state.opponentScore-state.myScore >= -10 {
		return false
	}

	// 对手没有刚刚开始一轮背叛
	if state.consecutiveDefections == 1 {
		return false
	}

	// 至少20轮 since a fresh start
	if currentRound-state.lastFreshStartRound < 20 {
		return false
	}

	// 至少10轮剩余
	if s.totalRounds-currentRound < 10 {
		return false
	}

	// 背叛次数与50-50随机生成器相差至少3.0标准差
	n := len(state.opponentHistory)
	if n == 0 {
		return false
	}

	defects := 0
	for _, act := range state.opponentHistory {
		if act == common.ActDefeat {
			defects++
		}
	}

	expected := float64(n) * 0.5
	stdDev := math.Sqrt(float64(n) * 0.5 * 0.5)
	diff := math.Abs(float64(defects) - expected)

	return diff >= 3.0*stdDev
}

// updateScores 更新得分（假设标准囚徒困境得分规则）
func (s *Tideman) updateScores(state *OpponentState, myAct, opponentAct common.ActEnum) {
	// 标准囚徒困境得分
	if myAct == common.ActCooperate && opponentAct == common.ActCooperate {
		state.myScore += 3
		state.opponentScore += 3
	} else if myAct == common.ActCooperate && opponentAct == common.ActDefeat {
		state.opponentScore += 5
	} else if myAct == common.ActDefeat && opponentAct == common.ActCooperate {
		state.myScore += 5
	} else { // 双方都背叛
		state.myScore += 1
		state.opponentScore += 1
	}
}

// getOpponentState 获取或创建对手状态
func (s *Tideman) getOpponentState(opponentIdx int) *OpponentState {
	if state, exists := s.opponentStates[opponentIdx]; exists {
		return state
	}

	state := &OpponentState{
		lastFreshStartRound: -20, // 确保初始满足20轮条件
	}
	s.opponentStates[opponentIdx] = state
	return state
}

// Index 返回策略索引
func (s *Tideman) Index() int {
	return s.index
}

// Reset 重置策略状态
func (s *Tideman) Reset() {
	s.opponentStates = make(map[int]*OpponentState)
}

// NewTidesStrategy 创建新策略实例
func NewTideman(index, totalRounds int) *Tideman {
	return &Tideman{
		index:          index,
		totalRounds:    totalRounds,
		opponentStates: make(map[int]*OpponentState),
	}
}
