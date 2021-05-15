package main

import "testing"

func TestStrStr(t *testing.T) {
	type caseOp struct {
		str string
		sub string
		ret int
	}
	testCaseOps := []caseOp{
		{"hello", "ll", 2},
		{"good afternoon", "noon", 10},
		{"up coffee", "", 0},
	}

	for _, caseOp := range testCaseOps {
		ret := StrStr(caseOp.str, caseOp.sub)
		if caseOp.ret == ret {
			t.Logf("expect ret: %2d, got ret: %2d", caseOp.ret, ret)
		} else {
			t.Errorf("oops...")
		}
	}
}
