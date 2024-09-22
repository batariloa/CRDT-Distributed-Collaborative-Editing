package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/batariloa/input"
	"github.com/batariloa/tree"
)


func main() {

 numInstancesStr := os.Getenv("NUM_INSTANCES")
    if numInstancesStr == "" {
        log.Fatal("NUM_INSTANCES environment variable not set")
    }

    numInstances, err := strconv.Atoi(numInstancesStr)
    if err != nil {
        log.Fatalf("Invalid NUM_INSTANCES value: %v", err)
    }

    // Initialize the CausalTree with the retrieved number of instances
    treeInstance := tree.NewCausalTree(numInstances, "a")

    fmt.Printf("CausalTree initialized with %d instances.\n", numInstances)

    input.EditDocument(treeInstance)



}
