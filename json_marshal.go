package main

import (
	"encoding/json"
	"fmt"
)

type ComputeSpec struct {
	Profile Profile `json:"profile"`

	Size ComputeSize `json:"size"`
}

type ComputeSize string
type Profile string

func main() {
  computeSpec := &ComputeSpec{Profile: "kubernetes_node", Size: "small"}
	b, err := json.Marshal(computeSpec)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
