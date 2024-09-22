package input

import (
	"log"
	"sync"

	"github.com/batariloa/tree"
	"github.com/eiannone/keyboard"
)

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

    log.Printf("Input char %d", char)

    inputChan <- char
  }
}

  func HandleInputs(inputChan <-chan rune, quitChan <-chan struct{}, ct *tree.CausalTree) {

    parentID := ct.Root.ID // Start with the root node as the initial parent

    for {

    tree.DisplayDocument(ct)
      select {

         case inputChar := <- inputChan:

            ct.AddNode(string(inputChar), parentID) 
            childID, exists := ct.GetLastChildId(parentID)  

            if exists {
                // Update parentID to the new child's ID for sequential insertion
                parentID = childID
            } else {
                // If no child exists, reset parentID to root's ID
                parentID = ct.Root.ID
            }

         case <- quitChan:

          return;
      }
    }

}
