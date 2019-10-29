package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/lruggieri/daily-coding-problems/util"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestSerialize(t *testing.T) {
	type test struct {
		root *Node
		expectedResult string
	}

	tests := []test{
		{
			root:&Node{
				Val:"root",
				Left:&Node{Val:"left"},
				Right:&Node{Val:"right"},
			},
			expectedResult:"root[left||right]",
		},
		{
			root:&Node{
				Val:"root",
			},
			expectedResult:"root",
		},
		{
			root:&Node{
				Val:"root",
				Left:&Node{
					Val:"left1",
					Right:&Node{
						Val:"right2",
					},
				},
				Right:&Node{Val:"right1"},
			},
			expectedResult:"root[left1[||right2]||right1]",
		},
		{
			root:&Node{
				Val:"root",
				Left:&Node{
					Val:"left",
					Left:&Node{
						Val:"left_left",
					},
					Right:&Node{
						Val:"left_right",
					},
				},
				Right:&Node{
					Val:"right",
					Left:&Node{
						Val:"right_left",
					},
					Right:&Node{
						Val:"right_right",
					},
				},
			},
			expectedResult:"root[left[left_left||left_right]||right[right_left||right_right]]",
		},
	}

	for tstIdx,tst := range tests{
		res := Serialize(tst.root)
		if res != tst.expectedResult{
			t.Error("test #",tstIdx,"expected",tst.expectedResult,"but got",res)
		}
	}
}

func TestDeserialize(t *testing.T) {
	type test struct {
		serializedTree string
		expectedResult *Node
	}

	tests := []test{
		{
			serializedTree:"root",
			expectedResult:&Node{
				Val:"root",
			},
		},
		{
			serializedTree:"root[left||right]",
			expectedResult:&Node{
				Val:"root",
				Left:&Node{Val:"left"},
				Right:&Node{Val:"right"},
			},
		},
		{
			serializedTree:"root[left[left_left||left_right]||right[right_left||right_right]]",
			expectedResult:&Node{
				Val:"root",
				Left:&Node{
					Val:"left",
					Left:&Node{
						Val:"left_left",
					},
					Right:&Node{
						Val:"left_right",
					},
				},
				Right:&Node{
					Val:"right",
					Left:&Node{
						Val:"right_left",
					},
					Right:&Node{
						Val:"right_right",
					},
				},
			},
		},
	}

	util.Logger.Level = logrus.DebugLevel
	for tstIdx, tst := range tests{
		res,err := Deserialize(tst.serializedTree)
		if err != nil{
			t.Error(err)
			continue
		}
		if !compareNodes(tst.expectedResult,res){
			spew.Dump(res)
			t.Error("test #",tstIdx,"fails")
		}
	}
}