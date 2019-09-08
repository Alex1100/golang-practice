package main

import (
  "fmt"
  bst "golang-practice/data-structures/binary-search-trees/bst"
)

func main() {
  bt := *bst.NewBST()
  bt.Insert(45)
  bt.Insert(32)
  bt.Insert(11)
  bt.Insert(101)
  bt.Insert(55)
  pre_order := bt.DepthFirstSearch("pre_order")
  in_order := bt.DepthFirstSearch("in_order")
  post_order := bt.DepthFirstSearch("post_order")
  fmt.Println(pre_order)
  fmt.Println(in_order)
  fmt.Println(post_order)
}
