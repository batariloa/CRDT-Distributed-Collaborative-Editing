package tree

func FlattenTree(root *TreeNode) []*TreeNode {
	var nodes []*TreeNode
	stack := []*TreeNode{}

	// first push roots children to top of stack
	stack = append(stack, root)

	for len(stack) > 0 {

		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.Operation == Insert && !isNodeDeleted(current) {
			nodes = append(nodes, current)
		}

		for i := len(current.Children) - 1; i >= 0; i-- {
			stack = append(stack, current.Children[i])
		}
	}

	return nodes
}

func isNodeDeleted(node *TreeNode) bool {
	for _, child := range node.Children {
		if child.Operation == Delete {
			return true
		}
	}

	return false
}
