package main

import (
	"fmt"

	"github.com/chengshusss/iter-prisoner-dilemma/utils"
)

const MaxScore = 64

var defeatScoreList = []int{
	1, 6, 7, 17, 22, 23, 26, 29, 30, 31, 33, 38, 39, 45, 49, 54, 55, 58, 61,
}

type Act int

const (
	ActCooperate Act = 0
	ActDefect    Act = 1
)

func (a Act) String() string {
	switch a {
	case ActCooperate:
		return "Co"
	case ActDefect:
		return "De"
	default:
		return "UNKNOW"
	}
}

type Match struct {
	SelfAct     Act
	OpponentAct Act
}

func (m Match) String() string {
	return fmt.Sprintf("{self: %v, oppo: %v}", m.SelfAct, m.OpponentAct)
}

type Case struct {
	No     int
	Rounds [3]Match
}

func (c Case) String() string {
	return fmt.Sprintf(
		"Case %d\nMatch 1: %v, Match 2: %v, Match 3: %v",
		c.No, c.Rounds[0], c.Rounds[1], c.Rounds[2],
	)
}

func (c *Case) NextAct() Act {
	if utils.Contain(defeatScoreList, c.No) {
		return ActDefect
	}
	return ActCooperate
}

func (c *Case) Format2Lines() []string {
	lines := make([]string, 4)

	// lines[0] = fmt.Sprintf("|/////Case %d\\\\\\\\\\|", c.No)
	lines[0] = "- - - - - - - - - - - - - - -  "
	lines[1] = fmt.Sprintf(
		"|-Round %2d-|%v   |%v   |%v   | ",
		c.No,
		c.Rounds[0].SelfAct, c.Rounds[1].SelfAct, c.Rounds[2].SelfAct,
	)
	// lines[1] = "- - - - - - - - -  "
	lines[2] = fmt.Sprintf(
		"|NextAct %v|   %v|   %v|   %v| ",
		c.NextAct(),
		c.Rounds[0].OpponentAct, c.Rounds[1].OpponentAct, c.Rounds[2].OpponentAct,
	)
	lines[3] = "- - - - - - - - - - - - - - -  "
	return lines
}

func NewCase(no int) *Case {

	if no < 0 || no >= MaxScore {
		panic(fmt.Sprintf("wrong no of [%d]", no))
	}

	c := &Case{
		No:     no,
		Rounds: [3]Match{},
	}

	for i := 0; i < 3; i++ {
		k := no % 4
		c.Rounds[2-i].OpponentAct = Act(k / 2)
		c.Rounds[2-i].SelfAct = Act(k % 2)

		no = no / 4
	}

	return c
}

func main() {
	cases := make([]*Case, MaxScore)
	for i := range MaxScore {
		cases[i] = NewCase(i)
	}

	const Cols = 4

	indexTable := make([][]int, (MaxScore+(Cols)-1)/Cols)
	for i := range MaxScore {

		col := i % Cols
		row := i / Cols
		if col == 0 {
			indexTable[i/Cols] = make([]int, Cols)
		}

		indexTable[row][col] = i
	}

	results := make([]string, 1, 4*MaxScore/Cols+1)
	results[0] += "           "
	for range Cols {
		results[0] += "|---INFO---| R1  | R2  | R3  |   "
	}
	for _, row := range indexTable {
		tempLines := []string{
			"         ",
			"Self Act ",
			"Oppo Act ",
			"         ",
		}

		for _, index := range row {
			lines := cases[index].Format2Lines()
			for i := range 4 {
				tempLines[i] += "  " + lines[i]
			}
		}
		results = append(results, tempLines...)
	}

	for _, line := range results {
		fmt.Println(line)
	}
}
