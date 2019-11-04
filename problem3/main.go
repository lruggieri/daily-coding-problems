package main

import (
	"errors"
	"github.com/lruggieri/daily-coding-problems/util"
	"strings"
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
func Deserialize(iSerializedNode string)(oNode *Node, oErr error){
	util.Logger.Debug("analyzing string "+iSerializedNode)
	if len(iSerializedNode) == 0{return nil,nil}

	currentNode := &Node{}

	/*
		1) everything before a [ is the current node value
		2) everything between the outermost set of [] are children node
			1) every SameLevelSeparator that is not inside another [] pair separates the 2 children
		3) for every children call recursively
	*/
	childrenStartIndex := strings.Index(iSerializedNode,"[") //from the first [
	childrenEndIndex := strings.LastIndex(iSerializedNode,"]") //to the last ]

	//1) everything before a [ is the current node value


	if childrenStartIndex > 0{
		//current node has some children

		currentNode.Val = iSerializedNode[:childrenStartIndex]

		if childrenStartIndex == len(iSerializedNode)-1{
			return nil,errors.New("invalid serialized input")
		}
		if childrenEndIndex <= childrenStartIndex {
			return nil, errors.New("invalid serialized input")
		}

		//2) everything between the outermost set of [] are children node
		childrenString := iSerializedNode[childrenStartIndex+1:childrenEndIndex]
		util.Logger.Debug("children string: "+childrenString)

		/*
		2)
			1) every SameLevelSeparator that is not inside another [] pair separates the 2 children
		*/
		separationSubStrings := strings.Split(childrenString, SameLevelSeparator)
		if len(separationSubStrings) < 2{
			return nil, errors.New("invalid serialized input")
		}
		for separationSubStringIdx := 1 ; separationSubStringIdx <  len(separationSubStrings) ; separationSubStringIdx++{
			/*
			if we have the same amount of [ and ] chars in the current separationSubString, it means that this SameLevelSeparator
			is separating children on this level
			*/
			leftSideOfSeparator := strings.Join(separationSubStrings[:separationSubStringIdx],SameLevelSeparator)
			rightSideOfSeparator := strings.Join(separationSubStrings[separationSubStringIdx:],SameLevelSeparator)
			util.Logger.Debug("leftSideOfSeparator: "+leftSideOfSeparator)
			util.Logger.Debug("rightSideOfSeparator: "+rightSideOfSeparator)

			leftSideIsClosed := strings.Count(leftSideOfSeparator,"[") == strings.Count(leftSideOfSeparator,"]")
			rightSideIsClosed := strings.Count(rightSideOfSeparator,"[") == strings.Count(rightSideOfSeparator,"]")

			if leftSideIsClosed && rightSideIsClosed{
				var err error
				util.Logger.Debug("left child: "+leftSideOfSeparator)
				util.Logger.Debug("right child: "+rightSideOfSeparator)
				currentNode.Left,err = Deserialize(leftSideOfSeparator)
				if err != nil{
					return nil, err
				}
				currentNode.Right,err = Deserialize(rightSideOfSeparator)
				if err != nil{
					return nil, err
				}
				break
			}
		}

		//
	}else{
		currentNode.Val = iSerializedNode
	}


	return currentNode,nil
}


func compareNodes(iNode1, iNode2 *Node) bool{
	return Serialize(iNode1)==Serialize(iNode2)
}