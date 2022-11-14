package decisionTree

import "lazympleza/lazy"

type DecisionTree lazy.LazyFunction

func CreateDecisionTree() lazy.LazyFunction {
	return GetWinnerWorldCup()
}
