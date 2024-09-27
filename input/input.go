package input

import (
	"log"
	"sync"

	"github.com/batariloa/tree"
	"github.com/eiannone/keyboard"
)

var cursorPos = 0

func EditDocument(ct *tree.CausalTree) {
	inputChan := make(chan rune, 5)
	quitChan := make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		ListenKeys(inputChan, quitChan)
	}()

	go func() {
		defer wg.Done()
		HandleInputs(inputChan, quitChan, ct)
	}()

	wg.Wait()
}

func ListenKeys(inputChan chan<- rune, quitChan chan<- struct{}) {
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	for {

		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Print(err)
			close(quitChan)
			return
		}

		if key == keyboard.KeyEsc {
			close(quitChan)
			return
		}

		if key == keyboard.KeySpace {
			inputChan <- ' '
			continue
		}

		if key == keyboard.KeyArrowLeft {
			inputChan <- '←'
			continue
		}

		if key == keyboard.KeyArrowRight {
			inputChan <- '→'
			continue
		}

		log.Printf("Input char %d", char)

		inputChan <- char
	}
}

func HandleInputs(inputChan <-chan rune, quitChan <-chan struct{}, ct *tree.CausalTree) {
	for {

		tree.DisplayDocument(ct)
		select {

		case inputChar := <-inputChan:

			if !handleCursorInputs(inputChar) {
				// Only insert the character if it's not a cursor input
				parentID := tree.GetNodeIdAtPos(cursorPos)
				ct.AddInsertNode(string(inputChar), parentID)
				cursorPos++
			}

		case <-quitChan:

			return
		}
	}
}

func handleCursorInputs(inputChar rune) bool {
	switch inputChar {
	case '←':

		cursorPos--
		return true

	case '→':

		cursorPos++
		return true
	}

	return false
}
