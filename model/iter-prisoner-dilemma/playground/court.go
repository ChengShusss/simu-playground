package playground

import (
	"fmt"

	"github.com/chengshusss/iter-prisoner-dilemma/utils"
)

type Court struct {
	StrategyNames     []string
	Protagonists      []StrategyEntry
	ShadowCompetitors []StrategyEntry
	TotalRewardMap    map[string]int
	DetailRewardMap   map[string]map[string]float64

	MatchCount int
	RoundCount int
}

func NewCourt(strategyNames []string, matchCount, roundCount int) *Court {
	n := len(strategyNames)
	c := Court{
		StrategyNames:     make([]string, n),
		Protagonists:      make([]StrategyEntry, n),
		ShadowCompetitors: make([]StrategyEntry, n),
		TotalRewardMap:    make(map[string]int, n),
		DetailRewardMap:   make(map[string]map[string]float64, n),
		MatchCount:        matchCount,
		RoundCount:        roundCount,
	}

	for i, name := range strategyNames {
		p := getStrategyFromName(name, i, roundCount)
		if p == nil {
			panic("do not have strategy imple " + name)
		}
		c.StrategyNames[i] = name
		c.Protagonists[i] = p
		c.ShadowCompetitors[i] = getStrategyFromName(name, n*10+i, roundCount)
	}

	return &c
}

func (c *Court) simulateOnce() {
	for i, protagonist := range c.Protagonists {
		for j, competitor := range c.ShadowCompetitors {
			u := NewUmpire(protagonist, competitor)

			rewards := make([]int, c.RoundCount)
			for k := 0; k < c.RoundCount; k++ {
				reward, _ := u.ConductOnce(k, c.RoundCount)
				rewards[k] = reward
			}

			nameA := c.StrategyNames[i]
			nameB := c.StrategyNames[j]

			rewardMap, ok := c.DetailRewardMap[nameA]
			if !ok {
				rewardMap = make(map[string]float64, len(c.StrategyNames))
			}
			rewardMap[nameB] += float64(utils.Sum(rewards))
			c.DetailRewardMap[nameA] = rewardMap
		}
	}
}

func (c *Court) reset() {
	for _, p := range c.Protagonists {
		p.Reset()
	}

	for _, s := range c.ShadowCompetitors {
		s.Reset()
	}
}

func (c *Court) FullSimulate() {
	for range c.MatchCount {
		c.simulateOnce()
		c.reset()
	}

	for i := range c.Protagonists {
		nameA := c.StrategyNames[i]
		rewardMap, ok := c.DetailRewardMap[nameA]
		if !ok {
			panic(fmt.Sprintf("missing reward map for %s", nameA))
		}

		for j := range c.ShadowCompetitors {
			nameB := c.StrategyNames[j]

			rewardMap[nameB] /= float64(c.MatchCount)
		}

		c.DetailRewardMap[nameA] = rewardMap
	}
}

func (c *Court) toDetailString() string {
	s := "        "
	for _, name := range c.StrategyNames {
		s += fmt.Sprintf("%10s", name)
	}
	s += "\n"

	for _, nameA := range c.StrategyNames {
		s += fmt.Sprintf("%10s", nameA)
		rewardMap, ok := c.DetailRewardMap[nameA]
		if !ok {
			panic(fmt.Sprintf("lack of simulation for [%s]\n", nameA))
		}
		for _, nameB := range c.StrategyNames {
			reward, ok := rewardMap[nameB]
			if !ok {
				panic(fmt.Sprintf("lack of simulation between [%s] and [%s]\n", nameA, nameB))
			}

			s += fmt.Sprintf("  %.2f  ", reward)
		}
		s += "\n"
	}

	return s
}

func (c *Court) Output() {

	fmt.Print(c.toDetailString())

	fmt.Print("===========================================================\nAvg Reward:")

	for _, nameA := range c.StrategyNames {
		rewardMap := c.DetailRewardMap[nameA]
		reward := utils.Avg(utils.ToSlice(rewardMap))
		fmt.Printf(" %.2f   ", reward)
	}
}
