package main

import (
	"github.com/chengshusss/iter-prisoner-dilemma/playground"
)

func main() {
	court := playground.NewCourt([]string{
		"tit",
		"tideman",
		"nydegger",
		"grofman",
		"shubik",
		"stein",
		"friedman",
		"davis",
		"random",
	}, 50, 200)

	court.FullSimulate()
	court.Output()
}
