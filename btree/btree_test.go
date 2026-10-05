package btree

import (
	"math/rand"
	"sort"
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

func TestBTreeBulkInsert(t *testing.T) {
	order := 4
	btree, _ := CreateBTree(0, order)

	const n = 10000
	keys := make([]int, n)
	for i := range keys {
		keys[i] = i + 1
	}
	rand.New(rand.NewSource(42)).Shuffle(n, func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	for _, key := range keys {
		btree.Insert(key, btree.Root)
	}

	validateBTreeStructure(t, btree.Root, order)

	got := collectInOrder(btree.Root)
	want := append([]int{0}, keys...)
	sort.Ints(want)

	if len(got) != len(want) {
		t.Fatalf("expected %d keys in tree, got %d", len(want), len(got))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("mismatch at position %d: want %d, got %d", i, want[i], got[i])
		}
	}
}

func validateBTreeStructure(t *testing.T, node *Node, order int) {
	t.Helper()
	if node == nil {
		t.Fatalf("encountered nil node")
	}
	if len(node.keys) > order {
		t.Fatalf("node %v exceeds order %d", node.keys, order)
	}
	if len(node.children) > 0 && len(node.children) != len(node.keys)+1 {
		t.Fatalf("node %v has %d keys but %d children", node.keys, len(node.keys), len(node.children))
	}
	for i := 1; i < len(node.keys); i++ {
		if node.keys[i-1] >= node.keys[i] {
			t.Fatalf("node keys not strictly sorted: %v", node.keys)
		}
	}
	for _, child := range node.children {
		validateBTreeStructure(t, child, order)
	}
}

func collectInOrder(node *Node) []int {
	if node == nil {
		return nil
	}
	if len(node.children) == 0 {
		return append([]int{}, node.keys...)
	}
	var result []int
	for i, key := range node.keys {
		result = append(result, collectInOrder(node.children[i])...)
		result = append(result, key)
	}
	result = append(result, collectInOrder(node.children[len(node.children)-1])...)
	return result
}
