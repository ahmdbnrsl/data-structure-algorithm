package main 

import "fmt"

type TNode struct {
    value int
    left, right *TNode
}

type Bst struct {
    Root *TNode
}

func Node(value int) *TNode {
    return &TNode{value: value, left: nil, right: nil}
}

func (bst *Bst) Insert(value int) {
    if bst.Root == nil {
        bst.Root = Node(value)
        return
    }
    current := bst.Root
    for {
        if value < (*current).value {
            if (*current).left == nil {
                current.left = Node(value)
                return
            }
            current = current.left
        } else if value > (*current).value  {
            if (*current).right == nil {
                current.right = Node(value)
                return
            }
            current = current.right
        } else {
            return
        }
    }
}

func (bst *Bst) Inorder(node *TNode) {
    if node == nil {
        return
    }
    bst.Inorder(node.left)
    fmt.Println(node.value)
    bst.Inorder(node.right)
}

func (bst *Bst) Search(node *TNode, target int) bool {
    if node == nil {
        fmt.Println("Empty")
        return false
    }
    if target == node.value {
        fmt.Println("Found ", target)
        return true
    }
    if target < node.value {
        return bst.Search(node.left, target)
    } else {
        return bst.Search(node.right, target)
    }
}

func (bst *Bst) DeleteValue(node *TNode, target int) *TNode {
    if node == nil {
        fmt.Println("Value Not Found")
        return nil
    }

    if target < node.value {
        node.left = bst.DeleteValue(node.left, target)
    } else if target > node.value {
        node.right = bst.DeleteValue(node.right, target)
    } else {
        if node.left == nil {
            return node.right
        } else if node.right == nil {
            return node.left
        }
        
        successor := node.right
        for successor.left != nil {
            successor = successor.left
        }
        node.value = successor.value
        node.right = bst.DeleteValue(node.right, successor.value)
    }
    return node
}

func (bst *Bst) Clear() {
    bst.Root = nil
}

func BST() Bst {
    return Bst{Root:nil}
}
func main() {
    bst := BST()
    bst.Insert(9)
    bst.Insert(6)
    bst.Insert(15)
    bst.Insert(8)
    bst.Insert(16)
    bst.Insert(10)
    bst.Insert(7)
    bst.Insert(11)
    bst.Insert(20)
    bst.Insert(5)
    bst.Inorder(bst.Root)
    bst.DeleteValue(bst.Root, 7)
    bst.Inorder(bst.Root)
    bst.DeleteValue(bst.Root, 9)
    fmt.Println(bst.Root.right.left)
    bst.Inorder(bst.Root)
}