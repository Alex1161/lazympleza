package decisionTree

import (
	"fmt"
	"lazympleza/logic"
	"lazympleza/memoize"
	"strconv"
)

func SimulateGroup(group string) (string, string) {
	groupA := [4]string{"Qatar", "Ecuador", "Senegal", "Netherlands"}
	groupB := [4]string{"England", "Iran", "USA", "Wales"}
	groupC := [4]string{"Argentina", "Saudi Arabia", "Mexico", "Poland"}
	groupD := [4]string{"France", "Australia", "Denmark", "Tunisia"}
	groupE := [4]string{"Spain", "Costa Rica", "Germany", "Japan"}
	groupF := [4]string{"Belgium", "Canada", "Morocco", "Croatia"}
	groupG := [4]string{"Brazil", "Serbia", "Switzerland", "Cameroon"}
	groupH := [4]string{"Portugal", "Ghana", "Uruguay", "South Korea"}

	var groupSelected [4]string

	switch group {
	case "A":
		groupSelected = groupA
	case "B":
		groupSelected = groupB
	case "C":
		groupSelected = groupC
	case "D":
		groupSelected = groupD
	case "E":
		groupSelected = groupE
	case "F":
		groupSelected = groupF
	case "G":
		groupSelected = groupG
	case "H":
		groupSelected = groupH
	}

	groupPoints := [4]int{0, 0, 0, 0}

	for i := 0; i < len(groupSelected); i++ {
		for j := i; j < len(groupSelected); j++ {
			if i == j {
				continue
			}
			win, proba := logic.GetWinnerBetween(groupSelected[i], groupSelected[j])
			if proba < 0 {
				groupPoints[i] = groupPoints[i] + 1
				groupPoints[j] = groupPoints[j] + 1
			} else {
				index := findIndex(groupSelected, win)
				groupPoints[index] = groupPoints[index] + 3
			}
		}
	}

	first := ""
	firstPoints := 0
	second := ""
	secondPoints := 0

	for p, v := range groupPoints {
		if v > firstPoints {
			secondPoints = firstPoints
			second = first
			firstPoints = v
			first = groupSelected[p]
		} else if v > secondPoints {
			secondPoints = v
			second = groupSelected[p]
		}
	}

	fmt.Println("")
	fmt.Println("Grupo " + group + ":")
	fmt.Println(first + " | " + strconv.FormatInt(int64(firstPoints), 10) + "pts")
	fmt.Println(second + " | " + strconv.FormatInt(int64(secondPoints), 10) + "pts")

	return first, second
}

func _GetWinnerWorldCup(args ...string) string {
	//Fase de grupos
	firstA, secondA := SimulateGroup("A")
	firstB, secondB := SimulateGroup("B")
	firstC, secondC := SimulateGroup("C")
	firstD, secondD := SimulateGroup("D")
	firstE, secondE := SimulateGroup("E")
	firstF, secondF := SimulateGroup("F")
	firstG, secondG := SimulateGroup("G")
	firstH, secondH := SimulateGroup("H")

	//Octavos de final
	winnerAB, prob := logic.GetWinnerBetween(firstA, secondB)
	winnerCD, prob := logic.GetWinnerBetween(firstC, secondD)
	winnerEF, prob := logic.GetWinnerBetween(firstE, secondF)
	winnerGH, prob := logic.GetWinnerBetween(firstG, secondH)
	winnerBA, prob := logic.GetWinnerBetween(firstB, secondA)
	winnerDC, prob := logic.GetWinnerBetween(firstD, secondC)
	winnerFE, prob := logic.GetWinnerBetween(firstF, secondE)
	winnerHG, prob := logic.GetWinnerBetween(firstH, secondG)

	fmt.Println("")
	fmt.Println("---- Octavos de final ----")
	fmt.Println("El ganador entre " + firstA + " y " + secondB + " es: " + winnerAB)
	fmt.Println("El ganador entre " + firstC + " y " + secondD + " es: " + winnerCD)
	fmt.Println("El ganador entre " + firstE + " y " + secondF + " es: " + winnerEF)
	fmt.Println("El ganador entre " + firstG + " y " + secondH + " es: " + winnerGH)
	fmt.Println("El ganador entre " + firstB + " y " + secondA + " es: " + winnerBA)
	fmt.Println("El ganador entre " + firstD + " y " + secondC + " es: " + winnerDC)
	fmt.Println("El ganador entre " + firstF + " y " + secondE + " es: " + winnerFE)
	fmt.Println("El ganador entre " + firstH + " y " + secondG + " es: " + winnerHG)

	//Cuartos de final
	winnerQ1, prob := logic.GetWinnerBetween(winnerAB, winnerCD)
	winnerQ2, prob := logic.GetWinnerBetween(winnerEF, winnerGH)
	winnerQ3, prob := logic.GetWinnerBetween(winnerBA, winnerDC)
	winnerQ4, prob := logic.GetWinnerBetween(winnerFE, winnerHG)

	fmt.Println("")
	fmt.Println("---- Cuartos de final ----")
	fmt.Println("El ganador entre " + winnerAB + " y " + winnerCD + " es: " + winnerQ1)
	fmt.Println("El ganador entre " + winnerEF + " y " + winnerGH + " es: " + winnerQ2)
	fmt.Println("El ganador entre " + winnerBA + " y " + winnerDC + " es: " + winnerQ3)
	fmt.Println("El ganador entre " + winnerFE + " y " + winnerHG + " es: " + winnerQ4)

	//Semifinal
	finalist1, prob := logic.GetWinnerBetween(winnerQ1, winnerQ2)
	finalist2, prob := logic.GetWinnerBetween(winnerQ3, winnerQ4)

	fmt.Println("")
	fmt.Println("---- Semifinal ----")
	fmt.Println("El ganador entre " + winnerQ1 + " y " + winnerQ2 + " es: " + finalist1)
	fmt.Println("El ganador entre " + winnerQ3 + " y " + winnerQ4 + " es: " + finalist2)

	//Final
	WCWinner, prob := logic.GetWinnerBetween(finalist1, finalist2)

	fmt.Println("")
	fmt.Println("---- Final ----")
	fmt.Println("El ganador entre " + finalist1 + " y " + finalist2 + " es: " + WCWinner)
	fmt.Println(prob)

	return WCWinner
}

func GetWinnerWorldCup() memoize.MemoizedFunction {
	return memoize.Memoized(_GetWinnerWorldCup)
}

// No estoy pudiendo usar esta misma funcion pero dejandola en el archivo util
func findIndex(array [4]string, value string) int {
	for p, v := range array {
		if v == value {
			return p
		}
	}
	return -1
}
