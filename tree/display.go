package tree

import (
	"fmt"
	"strings"
)

var flat []*TreeNode

func DisplayDocument(ct *CausalTree) {
	clearTerminal()
	flat = FlattenTree(ct.Root)

	doc := buildDocFromFlatTree(flat)
	fmt.Printf("\rDocument: %s\n", doc)
}

func buildDocFromFlatTree(ft []*TreeNode) string {
	var builder strings.Builder

	for _, node := range ft {
		builder.WriteString(node.Character)
	}

	return builder.String()
}

func clearTerminal() {
	fmt.Print("\033[H\033[2J") // ANSI escape codes to clear the screen
}

func GetNodeIdAtPos(pos int) int {
	return flat[pos].ID
}
