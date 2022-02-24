package main

import (
	"fmt"

	hc_input "github.com/ismtabo/hashcode-2022-qr/pkg/hc2022/input"
	hc_solution "github.com/ismtabo/hashcode-2022-qr/pkg/hc2022/solution"
)

func main() {
	input := hc_input.GetInput()
	problem := hc_input.ParseInput(input)
	// pp.Printf("Input:\n%+v\n", problem)
	solution := hc_solution.Resolve(problem)
	fmt.Println(solution.String())
}
