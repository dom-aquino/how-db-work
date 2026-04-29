package btree

import (
	"testing"
)

func TestBTree(t *testing.T) {
	rootInitialKey := 5
	order := 4
	btree, err := CreateBTree(rootInitialKey, order)
	if err != nil {
		t.Fatalf("BTree initialization failed")
	}

	btree.Insert(2, btree.Root)
	btree.Insert(10, btree.Root)
	btree.Insert(8, btree.Root)
	btree.Insert(7, btree.Root)

	if btree.Root.keys[0] != 5 {
		t.Fatalf("Splitting error")
	}

	if len(btree.Root.children[0].keys) == 0 {
		t.Fatalf("Child node is not created")
	}
}
