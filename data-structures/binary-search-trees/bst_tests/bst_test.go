package bst_test

import (
  "testing"
  bst "golang-practice/data-structures/binary-search-trees/bst"
)

func TestBST(t *testing.T) {
  bt := *bst.NewBST()
  if bt.Root != nil {
    t.Error("Expected root to be nil upon instantiation ", bt.Root)
  }
  bt.Insert(45)
  if bt.Root.Data != 45 {
    t.Error("Expected root to have a data value of 45")
  }
  bt.Insert(32)

  if bt.Root.Left.Data != 32 {
    t.Error("Expected the left leaf of root to be 32: ", bt.Root.Left.Data)
  }
  bt.Insert(11)
  bt.Insert(101)
  bt.Insert(55)
}
