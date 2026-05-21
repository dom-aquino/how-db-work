package btree

import (
	"fmt"
	"slices"
	"strings"
)

type BTree struct {
	Root  *Node
	order int // Maximum keys per node
}

type Node struct {
	keys     []int
	children []*Node
}

type splitResult struct {
	leftNode    *Node
	promotedKey int
	rightNode   *Node
}

func CreateBTree(rootNodeKey int, order int) (*BTree, error) {
	var rootNode Node
	rootNode.keys = append(rootNode.keys, rootNodeKey)
	btree := BTree{&rootNode, order}
	fmt.Printf("B-Tree Structure (order: %d)\n", btree.order)
	return &btree, nil
}

func (btree *BTree) SplitNode(node *Node) (*splitResult, error) {
	fmt.Println("Splitting node:", node.keys)
	var middleKey int = (len(node.keys) / 2)

	var leftNode Node
	leftNode.keys = slices.Clone(node.keys[0:middleKey])
	var rightNode Node
	rightNode.keys = slices.Clone(node.keys[middleKey+1:])

	var result splitResult
	result.leftNode = &leftNode
	result.promotedKey = node.keys[middleKey]
	result.rightNode = &rightNode

	return &result, nil
}

func (btree *BTree) Insert(key int, node *Node) (*splitResult, error) {
	fmt.Printf("Adding key %d to %d\n", key, node.keys)
	// Leaf node
	if len(node.children) == 0 {
		fmt.Printf("Leaf node\n\n")
		node.keys = append(node.keys, key)
		slices.Sort(node.keys)
		if len(node.keys) <= btree.order {
			return nil, nil
		} else {
			result, _ := btree.SplitNode(node)
			if result != nil {
				if btree.Root == node {
					var newRootNode Node
					newRootNode.keys = append(newRootNode.keys, result.promotedKey)
					newRootNode.children = append(newRootNode.children, result.leftNode, result.rightNode)
					btree.Root = &newRootNode
				} else {
					node.keys = append(node.keys, result.promotedKey)
					node.children = append(node.children, result.leftNode, result.rightNode)
				}
				return result, nil
			}
		}
	} else {
		fmt.Printf("Non-Leaf node\n\n")
		for i, nodeKey := range node.keys {
			if key < nodeKey {
				result, _ := btree.Insert(key, node.children[i])
				if result != nil {
					node.children = slices.Delete(node.children, i, i+1)
					node.children = slices.Insert(node.children, i, result.leftNode, result.rightNode)
					node.keys = append(node.keys, result.promotedKey)
					slices.Sort(node.keys)
				}
				break
			}
		}
	}
	return nil, nil
}

func (bt *BTree) ViewTree() {
	if bt.Root == nil {
		fmt.Println("Tree is empty")
		return
	}
	bt.printNode(bt.Root, 0, "")
}

func (bt *BTree) printNode(node *Node, level int, prefix string) {
	// Print current node's keys with indentation
	indent := strings.Repeat("  ", level)
	fmt.Printf("%s%s[Keys: %v]\n", indent, prefix, node.keys)

	// Print children recursively
	if len(node.children) > 0 {
		for i, child := range node.children {
			childPrefix := fmt.Sprintf("Child %d: ", i)
			bt.printNode(child, level+1, childPrefix)
		}
	}
}
