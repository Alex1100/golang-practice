package main

import (
  "fmt"
)

type OrderType string

type BST_Node struct {
  parent *BST_Node
  left *BST_Node
  right *BST_Node
  data int
}

type BST struct {
  root *BST_Node
  leaf_count int
}

func NewBST() *BST {
  return &BST{
    root: nil,
    leaf_count: 0,
  }
}

func (bt *BST) Insert(data int) {
  if bt.leaf_count == 0 {
    bt.root = &BST_Node{
      parent: nil,
      left: nil,
      right: nil,
      data: data,
    }
    bt.leaf_count++
  } else {
    added_leaf := bt.root.InsertLeaf(data)

    if added_leaf == 1 {
      bt.leaf_count++
    }

  }
}

func (b_node *BST_Node) InsertLeaf(data int) int {
  if data < b_node.data {
    if b_node.left == nil {
      b_node.left = &BST_Node{
        parent: b_node,
        left: nil,
        right: nil,
        data: data,
      }
      return 1
    } else {
      b_node.left.InsertLeaf(data)
    }
  } else if data > b_node.data {
    if b_node.right == nil {
      b_node.right = &BST_Node{
        parent: b_node,
        left: nil,
        right: nil,
        data: data,
      }
      return 1
    } else {
      b_node.right.InsertLeaf(data)
    }
  } else {
    return 0
  }

  return 1
}

func (b_node *BST_Node) TraverseTree(order_type OrderType, result []int) []int {

  if order_type == "pre_order" {
    result = append(result, b_node.data)
    if b_node.left != nil {
      result = b_node.left.TraverseTree(order_type, result)
    }

    if b_node.right != nil {
      result = b_node.right.TraverseTree(order_type, result)
    }
  } else if order_type == "in_order" {
    fmt.Println("HERE....", b_node.data)
    if b_node.left != nil {
      result = b_node.left.TraverseTree(order_type, result)
    }

    result = append(result, b_node.data)
    fmt.Println("RESULTIS NOW: ", result)

    if b_node.right != nil {
      result = b_node.right.TraverseTree(order_type, result)
    }
  } else if order_type == "post_order" {
    if b_node.left != nil {
      result = b_node.left.TraverseTree(order_type, result)
    }

    if b_node.right != nil {
      result = b_node.right.TraverseTree(order_type, result)
    }

    result = append(result, b_node.data)
  }

  return result
}

func (bt *BST) DepthFirstSearch(order_type OrderType) []int {
  var result = []int{}
  if bt.leaf_count == 0 {
    return result
  }

  b_node := bt.root
  result = b_node.TraverseTree(order_type, result)
  return result
}


func main() {
  bst := NewBST()
  bst.Insert(45)
  bst.Insert(32)
  bst.Insert(11)
  bst.Insert(101)
  bst.Insert(55)
  pre_order := bst.DepthFirstSearch("pre_order")
  in_order := bst.DepthFirstSearch("in_order")
  post_order := bst.DepthFirstSearch("post_order")
  fmt.Println(pre_order)
  fmt.Println(in_order)
  fmt.Println(post_order)
}
