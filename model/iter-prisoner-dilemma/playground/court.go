package playground

import (
	"fmt"

	"github.com/chengshusss/iter-prisoner-dilemma/strategy"
	"github.com/chengshusss/iter-prisoner-dilemma/utils"
)

type Court struct {
	StrategyNames     []string
	Protagonists      []StrategyEntry
	ShadowCompetitors []StrategyEntry
	TotalRewardMap    map[string]int
	DetailRewardMap   map[string]map[string]float64

	Round int
}

func NewCourt(strategyNames []string, round int) *Court {
	n := len(strategyNames)
	c := Court{
		StrategyNames:     make([]string, n),
		Protagonists:      make([]StrategyEntry, n),
		ShadowCompetitors: make([]StrategyEntry, n),
		TotalRewardMap:    make(map[string]int, n),
		DetailRewardMap:   make(map[string]map[string]float64, n),
		Round:             round,
	}

	for i, name := range strategyNames {
		p := getStrategyFromName(name, i)
		if p == nil {
			panic("do not have strategy imple " + name)
		}
		c.StrategyNames[i] = name
		c.Protagonists[i] = p
		c.ShadowCompetitors[i] = getStrategyFromName(name, n*10+i)
	}

	return &c
}

func (c *Court) FullSimulate() {
	for i, protagonist := range c.Protagonists {
		for j, competitor := range c.ShadowCompetitors {
			u := NewUmpire(protagonist, competitor)

			rewards := make([]int, c.Round)
			for k := 0; k < c.Round; k++ {
				reward, _ := u.ConductOnce()
				rewards[k] = reward
			}

			nameA := c.StrategyNames[i]
			nameB := c.StrategyNames[j]
			avgReward := utils.Avg(rewards)

			rewardMap, ok := c.DetailRewardMap[nameA]
			if !ok {
				rewardMap = make(map[string]float64, len(c.StrategyNames))
			}
			rewardMap[nameB] = avgReward
			c.DetailRewardMap[nameA] = rewardMap
		}
	}
}

func (c *Court) Output() {
	for _, nameA := range c.StrategyNames {
		for _, nameB := range c.StrategyNames {
			rewardMap, ok := c.DetailRewardMap[nameA]
			if !ok {
				panic(fmt.Sprintf("lack of simulation for [%s]\n", nameA))
			}
			reward, ok := rewardMap[nameB]
			if !ok {
				panic(fmt.Sprintf("lack of simulation between [%s] and [%s]\n", nameA, nameB))
			}

			fmt.Printf("%15s vs %15s: %.2f\n", nameA, nameB, reward)
		}
	}
}

func getStrategyFromName(name string, idx int) StrategyEntry {
	switch name {
	case "tit_for_tat":
		return strategy.NewTitForTat(idx)
	case "random":
		return strategy.NewRandom(idx)

	default:
		return nil
	}
}
