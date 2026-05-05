package btree

import (
	"testing"
)

func TestBTree(t *testing.T) {
	rootInitialKey := 5
	order := 4
	btree, err := CreateBTree(rootInitialKey, order)
	if err != nil {
		t.Fatalf("Creation of the B Tree failed")
	}

	btree.Insert(2, btree.Root)
	btree.Insert(10, btree.Root)
	btree.Insert(8, btree.Root)

	// Overflow starts here
	// A new root node with 8 as a lone key should be created
	// Two children nodes should be created as well
	btree.Insert(16, btree.Root)
	if len(btree.Root.keys) != 1 || btree.Root.keys[0] != 8 {
		t.Fatalf("Wrong root node")
	}

	//if len(btree.Root.children) != 2 {
	//	fmt.Println("Wrong root node")
	//}
}
