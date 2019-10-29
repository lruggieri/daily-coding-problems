package main

import (
	"github.com/lruggieri/daily-coding-problems/util"
)

const(
	SameLevelSeparator = "||"
)

type Node struct{
	Val string
	Left *Node
	Right *Node
}

func Serialize(iNode *Node) (oSerializedNode string){
	if iNode == nil{return ""}
	oSerializedNode += iNode.Val

	if !(iNode.Left == nil && iNode.Right == nil){
		oSerializedNode += "["

		if iNode.Left != nil{
			oSerializedNode += Serialize(iNode.Left)
		}
		oSerializedNode += SameLevelSeparator
		if iNode.Right != nil{
			oSerializedNode += Serialize(iNode.Right)
		}

		oSerializedNode += "]"
	}
	return
}

//TODO it should work with runes, not strings
func Deserialize(iSerializedNode string, father *Node)(oNode *Node, oErr error){
	util.Logger.Debug("analyzing string "+iSerializedNode)
	if len(iSerializedNode) == 0{return nil,nil}

	currentNode := &Node{}

	/*
		1) everything before a [ if the current node value
		2) everything between the outermost set of [] are children node
			1) every SameLevelSeparator that is not inside another [] pair separates the 2 children
		3) for every children call recursively
	*/

	return currentNode,nil
}


func compareNodes(iNode1, iNode2 *Node) bool{
	return Serialize(iNode1)==Serialize(iNode2)
}