package tree

import "github.com/batariloa/util"
import "github.com/batariloa/config"

type TreeNode struct {
  ID int
  Character string
  Operation Operation
  Parent *TreeNode
  Children []*TreeNode
  VectorClock []int
}

type CausalTree struct {
  Root *TreeNode
  VectorClock []int
}

// NewCausalTree creates a new CausalTree with an initialized root node and vector clock
func NewCausalTree(numInstances int, rootCharacter string) *CausalTree {
	// Initialize the vector clock with zeros, one entry per instance
	vectorClock := make([]int, numInstances)

	// Create the root node
	rootNode := &TreeNode{
		ID:          0, // Typically, the root node can have a fixed ID like 0
		Character:   rootCharacter,
		Operation:   Insert, // Assuming the root node is an "insert" operation
		Parent:      nil,
		Children:    []*TreeNode{},
		VectorClock: vectorClock,
	}

	// Return the new CausalTree with the root node and initialized vector clock
	return &CausalTree{
		Root:        rootNode,
		VectorClock: vectorClock,
	}
}

type Operation int
const (
	Insert Operation = iota
	Delete
)

func (ct *CausalTree) getNodeById(id int) *TreeNode {
  return FindNodeBFS(ct.Root, id)
}

func (ct *CausalTree) AddNode(character string, parentId int) {

  parentNode := ct.getNodeById(parentId)

  ct.VectorClock[config.InstanceNum]++
  newNode := ct.CreateNode(
    parentNode, 
    character, 
    Insert,
    append([]int(nil), ct.VectorClock...))

  parentNode.Children = append(parentNode.Children, newNode)
}

func (ct *CausalTree) CreateNode(parent *TreeNode, character string,
  operation Operation,vc []int) *TreeNode {

  uniquteId := util.GetUniqueId()

  return &TreeNode {
    ID: uniquteId, 
    Parent:  parent,
    Character: character,
    Operation: operation,
    VectorClock: vc,
    Children: []*TreeNode{},
  }
}
