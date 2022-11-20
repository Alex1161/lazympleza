package decisionTree

import (
	"lazympleza/memoize"
)

func _GetWinnerWorldCup(args ...string) string {
	return "Argentina (anulo mufa)"
}

func GetWinnerWorldCup() memoize.MemoizedFunction {
	return memoize.Memoized(_GetWinnerWorldCup)
}
