package main

import (
	"github.com/chengshusss/iter-prisoner-dilemma/playground"
)

func main() {
	court := playground.NewCourt([]string{
		"tit_for_tat",
		"random",
	}, 200)

	court.FullSimulate()
	court.Output()
}
