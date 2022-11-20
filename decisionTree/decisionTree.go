package decisionTree

import "lazympleza/memoize"

type DecisionTree memoize.MemoizedFunction

func CreateDecisionTree() memoize.MemoizedFunction {
	return GetWinnerWorldCup()
}
