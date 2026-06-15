package btree

import (
	"testing"
)

func TestBTreeCreation(t *testing.T) {
	key := 5
	order := 4
	_, err := CreateBTree(key, order)
	if err != nil {
		t.Fatalf("Creation of the B Tree failed")
	}
}

func TestBTreeNoSplit(t *testing.T) {
	key := 5
	order := 4
	btree, _ := CreateBTree(key, order)

	btree.Insert(2, btree.Root)
	btree.Insert(10, btree.Root)
	btree.Insert(8, btree.Root)
	if len(btree.Root.keys) != 4 {
		t.Fatalf("Wrong number of keys")
	}
	if len(btree.Root.children) != 0 {
		t.Fatalf("Wrong number of children")
	}
}

func TestBTreeSplitSimple(t *testing.T) {
	key := 5
	order := 4
	btree, _ := CreateBTree(key, order)

	btree.Insert(2, btree.Root)
	btree.Insert(10, btree.Root)
	btree.Insert(8, btree.Root)
	btree.Insert(16, btree.Root)
	if btree.Root.keys[0] != 8 {
		t.Fatalf("Wrong root key")
	}
	if len(btree.Root.keys) != 1 {
		t.Fatalf("Wrong number of keys")
	}
	if len(btree.Root.children) != 2 {
		t.Fatalf("Wrong number of children")
	}
}
