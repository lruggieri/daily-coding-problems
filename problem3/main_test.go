package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/lruggieri/daily-coding-problems/util"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerialize(t *testing.T) {
	assert.Equal(t, "root[left||right]",Serialize(&Node{
		Val:"root",
		Left:&Node{Val:"left"},
		Right:&Node{Val:"right"},
	}))

	assert.Equal(t, "root",Serialize(&Node{
		Val:"root",
	}))

	assert.Equal(t, "root[left1[||right2]||right1]",Serialize(&Node{
		Val:"root",
		Left:&Node{
			Val:"left1",
			Right:&Node{
				Val:"right2",
			},
		},
		Right:&Node{Val:"right1"},
	}))

	assert.Equal(t, "root[left[left_left||left_right]||right[right_left||right_right]]",Serialize(&Node{
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
	}))
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
		{
			serializedTree:"root[left[left_left||left_right[left_right_left||left_right_right]]||right[right_left||right_right]]",
			expectedResult:&Node{
				Val:"root",
				Left:&Node{
					Val:"left",
					Left:&Node{
						Val:"left_left",
					},
					Right:&Node{
						Val:"left_right",
						Left:&Node{
							Val:"left_right_left",
						},
						Right:&Node{
							Val:"left_right_right",
						},
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

	util.Logger.Level = logrus.InfoLevel
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

func TestCommon(t *testing.T){
	tests := []string{
		"root",
		"root[left||right]",
		"root[left[left_left||left_right]||right[right_left||right_right]]",
		"root[left[left_left||left_right[left_right_left||left_right_right]]||right[right_left||right_right]]",
		"root[left[left_left||left_right[left_right_left||left_right_right]]||right[right_left[right_left_left||right_left_right]||right_right]]",
	}

	for tstIdx,tst := range tests{
		serialization,err := Deserialize(tst)
		if err != nil{
			t.Error("test #",tstIdx,"error:",err)
			continue
		}
		if tst != Serialize(serialization){
			t.Error("test #",tstIdx,"serialization of deserialization doesn't match original string")
			continue
		}
	}
}
