package _const

import (
	"os"
	"strconv"
)

var MachineLen = 1

func init() {
	MachineLenStr := os.Getenv("MachineLen")
	if MachineLenStr == "" {
		return
	}
	if id, err := strconv.ParseInt(MachineLenStr, 10, 64); err != nil {
		MachineLen = 0
	} else {
		MachineLen = int(id)
	}
}
