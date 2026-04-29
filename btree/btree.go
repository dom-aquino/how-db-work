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
	for childNode := range parentNode.children {
		fmt.Printf("%T", childNode)
	}

	keyToMove := parentNode.keys[0]       // Get the leftmost key
	parentNode.keys = parentNode.keys[1:] // Re-slice excluding the leftmost key

	var newNode Node
	newNode.keys = append(newNode.keys, keyToMove)
	parentNode.children = append(parentNode.children, &newNode)
}

func (btree *BTree) Insert(key int, node *Node) {
	if len(node.keys) < btree.order {
		node.keys = append(node.keys, key)
		slices.Sort(node.keys)
	} else {
		node.keys = append(node.keys, key)
		slices.Sort(node.keys)
		CreateOrUpdateChildNode(key, node)
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
