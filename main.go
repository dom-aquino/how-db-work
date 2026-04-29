package main

import "github.com/dom-aquino/how-db-work/btree"

func main() {
	btree, _ := btree.CreateBTree(5, 4)
	btree.Insert(2, btree.Root)
	btree.Insert(10, btree.Root)
	btree.Insert(8, btree.Root)
	btree.Insert(7, btree.Root)
	btree.ViewTree()
}
