package tree

import "fmt"

func DisplayDocument(ct *CausalTree) {

	clearTerminal()
	traverse(ct.Root)
}

func clearTerminal() {
	fmt.Print("\033[H\033[2J") // ANSI escape codes to clear the screen
}
