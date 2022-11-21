package decisionTree

type Prediction func(args ...string) string

func _GetWinnerWorldCup(args ...string) string {
	return "Argentina (anulo mufa)"
}

func GetWinnerWorldCup() Prediction {
	return _GetWinnerWorldCup
}
