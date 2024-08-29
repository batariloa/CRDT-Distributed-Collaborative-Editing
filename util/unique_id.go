package util

import (
	"fmt"
	"strconv"

	cfg "github.com/batariloa/config"
)
var currMaxId int32


func init() {
  currMaxId = 1
}

func IncrementId() {
  currMaxId+=1
}

func GetUniqueId() int {

  IncrementId()

  instanceId := cfg.InstanceNum
  prefixStr := fmt.Sprintf("%d", instanceId)
  numberStr := fmt.Sprintf("%d", currMaxId)

  newId := prefixStr + numberStr

  newIdInt, err := strconv.Atoi(newId)
  if(err!=nil) {
    panic("aaa")
  }

  return newIdInt
}
