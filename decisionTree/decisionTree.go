package decisionTree

import "fmt"

func CreateDecisionTree() chan string {
	tree := make(chan string)
	go func() {
		for request := range tree {
			fmt.Println(request)
			tree <- "Argentina (anulo mufa)"
		}
	}()

	return tree
}
