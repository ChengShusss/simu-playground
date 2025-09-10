package strategy

import (
	"fmt"
	"testing"

	"github.com/chengshusss/iter-prisoner-dilemma/common"
)

func Test2DArry(t *testing.T) {

	a := [3][2]common.ActEnum{}

	a[0][1] = common.ActCooperate

	fmt.Printf("%v\n", a[0])
	fmt.Printf("%v\n", a[0][1])
}

// It turn out that customed enum does support Logic OR, even ENUM is int actually.
// func TestEnumLogicOr(t *testing.T) {
// 	co := common.ActCooperate
// 	de := common.ActDefeat

// 	fmt.Printf("OR: %v, AND: %v\n Same OR: %v Same AND: %v\n", co || de, )
// }

// Check whether getNydeggerScore is correct
func TestGetNydeggerScore(t *testing.T) {
	cases := []struct {
		input  [3][2]common.ActEnum
		Except int
	}{
		{
			[3][2]common.ActEnum{
				{common.ActCooperate, common.ActCooperate},
				{common.ActCooperate, common.ActCooperate},
				{common.ActCooperate, common.ActCooperate},
			},
			0,
		},
		{
			[3][2]common.ActEnum{
				{common.ActDefeat, common.ActDefeat},
				{common.ActDefeat, common.ActDefeat},
				{common.ActDefeat, common.ActDefeat},
			},
			63,
		},
		{
			[3][2]common.ActEnum{
				{common.ActDefeat, common.ActCooperate},
				{common.ActCooperate, common.ActDefeat},
				{common.ActCooperate, common.ActCooperate},
			},
			24,
		},
		{
			[3][2]common.ActEnum{
				{common.ActDefeat, common.ActCooperate},
				{common.ActCooperate, common.ActDefeat},
				{common.ActDefeat, common.ActCooperate},
			},
			25,
		},
	}

	for _, c := range cases {
		score := getNydeggerScore(c.input)
		if score != c.Except {
			t.Fatalf("Input: %v, expect: %v, gotten: %v\n", c.input, c.Except, score)
		}
	}
}

func Test2DArrayShift(t *testing.T) {
	a := [3][2]common.ActEnum{}

	a[0][1] = common.ActCooperate
	a[1][0] = common.ActDefeat
	a[2][0] = common.ActCooperate
	fmt.Printf("%v\n", a)

	a[0] = a[1]
	a[1] = a[2]

	fmt.Printf("%v\n", a)
}
