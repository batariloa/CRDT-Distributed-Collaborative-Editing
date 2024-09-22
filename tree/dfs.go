package tree

import "fmt"


func traverse(root *TreeNode) string{

  var doc string
  stack := []*TreeNode{}

  // first push roots children to top of stack
  stack = append(stack, root)

  for len(stack) > 0{

    current := stack[len(stack)-1]
    stack = stack[:len(stack)-1]



    if current.Operation == Insert {

      doc += current.Character
    }


    for i := len(current.Children) - 1; i>=0; i-- {

      stack = append(stack, current.Children[i])
    }
  }

    fmt.Printf("\rDocument: %s\n", doc)

    return doc
}

