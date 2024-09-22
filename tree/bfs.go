package tree

import "container/list"


func FindNodeBFS(root *TreeNode, id int) *TreeNode {

if(root == nil) {
    return nil
  }

  queue := list.New()
  queue.PushBack(root)

  for queue.Len() > 0 {

    currNode := queue.Remove(queue.Front()).(*TreeNode)

    if currNode.ID == id {
      return currNode
    }

  for _, child := range currNode.Children {
      queue.PushBack(child)
    }
  }

  return nil
}

