package decisionTree

func CreateDecisionTree() chan string {
	tree := make(chan string)
	go func() {
		for request := range tree {
			r := ToRequest(request)
			var result string

			switch r.request {
			case "winner":
				result = GetWinnerWorldCup()
			}
			tree <- result
		}
	}()

	return tree
}
