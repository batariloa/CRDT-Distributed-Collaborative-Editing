package main

import (
	"log"

	"github.com/batariloa/tree"
	"github.com/eiannone/keyboard"
)



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

    inputChan <- char
  }
}

  func HandleInputs(inputChan <-chan rune, quitChan <-chan struct{}, ct *tree.CausalTree) {

    parentID := ct.Root.ID // Start with the root node as the initial parent

    for {

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
