package bst

type OrderType string

type BST_Node struct {
  Parent *BST_Node
  Left *BST_Node
  Right *BST_Node
  Data int
}

type BST struct {
  Root *BST_Node
  LeafCount int
}

func NewBST() *BST {
  return &BST{
    Root: nil,
    LeafCount: 0,
  }
}

func (bt *BST) Insert(data int) {
  if bt.LeafCount == 0 {
    bt.Root = &BST_Node{
      Parent: nil,
      Left: nil,
      Right: nil,
      Data: data,
    }
    bt.LeafCount++
  } else {
    added_leaf := bt.Root.InsertLeaf(data)

    if added_leaf == 1 {
      bt.LeafCount++
    }

  }
}

func (b_node *BST_Node) InsertLeaf(data int) int {
  if data < b_node.Data {
    if b_node.Left == nil {
      b_node.Left = &BST_Node{
        Parent: b_node,
        Left: nil,
        Right: nil,
        Data: data,
      }
      return 1
    } else {
      b_node.Left.InsertLeaf(data)
    }
  } else if data > b_node.Data {
    if b_node.Right == nil {
      b_node.Right = &BST_Node{
        Parent: b_node,
        Left: nil,
        Right: nil,
        Data: data,
      }
      return 1
    } else {
      b_node.Right.InsertLeaf(data)
    }
  } else {
    return 0
  }

  return 1
}

func (b_node *BST_Node) TraverseTree(order_type OrderType, result []int) []int {

  if order_type == "pre_order" {
    result = append(result, b_node.Data)
    if b_node.Left != nil {
      result = b_node.Left.TraverseTree(order_type, result)
    }

    if b_node.Right != nil {
      result = b_node.Right.TraverseTree(order_type, result)
    }
  } else if order_type == "in_order" {
    if b_node.Left != nil {
      result = b_node.Left.TraverseTree(order_type, result)
    }

    result = append(result, b_node.Data)

    if b_node.Right != nil {
      result = b_node.Right.TraverseTree(order_type, result)
    }
  } else if order_type == "post_order" {
    if b_node.Left != nil {
      result = b_node.Left.TraverseTree(order_type, result)
    }

    if b_node.Right != nil {
      result = b_node.Right.TraverseTree(order_type, result)
    }

    result = append(result, b_node.Data)
  }

  return result
}

func (bt *BST) DepthFirstSearch(order_type OrderType) []int {
  var result = []int{}
  if bt.LeafCount == 0 {
    return result
  }

  b_node := bt.Root
  result = b_node.TraverseTree(order_type, result)
  return result
}
