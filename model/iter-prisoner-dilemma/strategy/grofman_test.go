package strategy

import (
	"fmt"
	"testing"

	"github.com/chengshusss/iter-prisoner-dilemma/common"
)

func TestGrofman(t *testing.T) {
	g := NewGrofman(1, 100)

	fmt.Printf("%+v\n", g)
	act := g.Act(2, 0)
	fmt.Printf("Act: %+v\n", act)
	fmt.Printf("%+v\n", g)
	g.Update(0, 2, common.ActCooperate)
	fmt.Printf("%+v\n", g)

}
