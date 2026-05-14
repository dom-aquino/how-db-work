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

func CreateBTree(rootNodeKey int, order int) (*BTree, error) {
	var rootNode Node
	rootNode.keys = append(rootNode.keys, rootNodeKey)
	btree := BTree{&rootNode, order}
	fmt.Printf("B-Tree Structure (order: %d)\n", btree.order)
	return &btree, nil
}

func CreateOrUpdateChildNode(key int, parentNode *Node) {
	if len(parentNode.children) == 0 {
		var newNode Node
		newNode.keys = append(newNode.keys, key)
		parentNode.children = append(parentNode.children, &newNode)
	}
}

func (btree *BTree) SplitNode(node *Node) {
	// Create new root code using the middle key of the current node
	var middleKey int = (len(node.keys) / 2)
	var newRootNode Node
	newRootNode.keys = append(newRootNode.keys, node.keys[middleKey])
	btree.Root = &newRootNode

	var leftNode Node
	leftNode.keys = node.keys[0:middleKey]

	var rightNode Node
	rightNode.keys = node.keys[middleKey+1:]

	newRootNode.children = append(newRootNode.children, &leftNode)
	newRootNode.children = append(newRootNode.children, &rightNode)
}

func (btree *BTree) Insert(key int, node *Node) {
	if len(node.children) == 0 {
		node.keys = append(node.keys, key)
		slices.Sort(node.keys)
	}

	if len(node.keys) > btree.order {
		btree.SplitNode(node)
	} else {
		for _, child := range node.children {
			if child.keys[len(child.keys)-1] > key {
				child.keys = append(child.keys, key)
				slices.Sort(child.keys)
				break
			}
		}
	}
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
