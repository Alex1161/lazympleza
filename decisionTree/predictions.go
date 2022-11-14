package decisionTree

import (
	"lazympleza/lazy"
	"time"
)

func _GetWinnerWorldCup() string {
	time.Sleep(10 * time.Second)
	return "Argentina (anulo mufa)"
}

func GetWinnerWorldCup() lazy.LazyFunction {
	return lazy.Lazy(_GetWinnerWorldCup)
}
