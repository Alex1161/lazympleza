package decisionTree

import (
	"fmt"
	"math/rand"
	"strconv"
)

func SimulateGroup(group string, tree decision_tree) (string, string) {
	groupA := []string{"Qatar", "Ecuador", "Senegal", "Netherlands"}
	groupB := []string{"England", "Iran", "USA", "Wales"}
	groupC := []string{"Argentina", "Saudi Arabia", "Mexico", "Poland"}
	groupD := []string{"France", "Australia", "Denmark", "Tunisia"}
	groupE := []string{"Spain", "Costa Rica", "Germany", "Japan"}
	groupF := []string{"Belgium", "Canada", "Morocco", "Croatia"}
	groupG := []string{"Brazil", "Serbia", "Switzerland", "Cameroon"}
	groupH := []string{"Portugal", "Ghana", "Uruguay", "South Korea"}

	var groupSelected []string

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
			win, proba := GetWinnerBetween(groupSelected[i], groupSelected[j], tree)
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

func GetWinnerBetween(home string, away string, tree decision_tree) (string, float64) {
	match_info := make(map[string]string)

	setMatchInfo(match_info, home, away)

	home_result := predict(match_info, tree)

	if home_result == "Win" {
		return home, 1
	} else if home_result == "Lose" {
		return away, 1
	} else {
		source := rand.NewSource(45)
		rnd := rand.New(source)
		simulation := rnd.Float64()
		if simulation <= 0.5 {
			return home, -1
		} else {
			return away, -1
		}
	}
}

func GetWinnerWorldCup(tree decision_tree) string {
	//Fase de grupos
	firstA, secondA := SimulateGroup("A", tree)
	firstB, secondB := SimulateGroup("B", tree)
	firstC, secondC := SimulateGroup("C", tree)
	firstD, secondD := SimulateGroup("D", tree)
	firstE, secondE := SimulateGroup("E", tree)
	firstF, secondF := SimulateGroup("F", tree)
	firstG, secondG := SimulateGroup("G", tree)
	firstH, secondH := SimulateGroup("H", tree)

	//Octavos de final
	winnerAB, _ := GetWinnerBetween(firstA, secondB, tree)
	winnerCD, _ := GetWinnerBetween(firstC, secondD, tree)
	winnerEF, _ := GetWinnerBetween(firstE, secondF, tree)
	winnerGH, _ := GetWinnerBetween(firstG, secondH, tree)
	winnerBA, _ := GetWinnerBetween(firstB, secondA, tree)
	winnerDC, _ := GetWinnerBetween(firstD, secondC, tree)
	winnerFE, _ := GetWinnerBetween(firstF, secondE, tree)
	winnerHG, _ := GetWinnerBetween(firstH, secondG, tree)

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
	winnerQ1, _ := GetWinnerBetween(winnerAB, winnerCD, tree)
	winnerQ2, _ := GetWinnerBetween(winnerEF, winnerGH, tree)
	winnerQ3, _ := GetWinnerBetween(winnerBA, winnerDC, tree)
	winnerQ4, _ := GetWinnerBetween(winnerFE, winnerHG, tree)

	fmt.Println("")
	fmt.Println("---- Cuartos de final ----")
	fmt.Println("El ganador entre " + winnerAB + " y " + winnerCD + " es: " + winnerQ1)
	fmt.Println("El ganador entre " + winnerEF + " y " + winnerGH + " es: " + winnerQ2)
	fmt.Println("El ganador entre " + winnerBA + " y " + winnerDC + " es: " + winnerQ3)
	fmt.Println("El ganador entre " + winnerFE + " y " + winnerHG + " es: " + winnerQ4)

	//Semifinal
	finalist1, _ := GetWinnerBetween(winnerQ1, winnerQ2, tree)
	finalist2, _ := GetWinnerBetween(winnerQ3, winnerQ4, tree)

	fmt.Println("")
	fmt.Println("---- Semifinal ----")
	fmt.Println("El ganador entre " + winnerQ1 + " y " + winnerQ2 + " es: " + finalist1)
	fmt.Println("El ganador entre " + winnerQ3 + " y " + winnerQ4 + " es: " + finalist2)

	//Final
	WCWinner, _ := GetWinnerBetween(finalist1, finalist2, tree)

	fmt.Println("")
	fmt.Println("---- Final ----")
	fmt.Println("El ganador entre " + finalist1 + " y " + finalist2 + " es: " + WCWinner)

	return WCWinner
}

func setMatchInfo(match_info map[string]string, home string, away string) {
	countries_info := map[string]map[string]string{
		"Qatar": {"continent": "Asia"}, "Ecuador": {"continent": "South America"}, "Senegal": {"continent": "Africa"}, "Netherlands": {"continent": "Europe"},
		"England": {"continent": "Europe"}, "Iran": {"continent": "Asia"}, "USA": {"continent": "North America"}, "Wales": {"continent": "Europe"},
		"Argentina": {"continent": "South America"}, "Saudi Arabia": {"continent": "Asia"}, "Mexico": {"continent": "North America"}, "Poland": {"continent": "Europe"},
		"France": {"continent": "Europe"}, "Australia": {"continent": "Oceania"}, "Denmark": {"continent": "Europe"}, "Tunisia": {"continent": "Africa"},
		"Spain": {"continent": "Europe"}, "Costa Rica": {"continent": "North America"}, "Germany": {"continent": "Europe"}, "Japan": {"continent": "Asia"},
		"Belgium": {"continent": "Europe"}, "Canada": {"continent": "North America"}, "Morocco": {"continent": "Africa"}, "Croatia": {"continent": "Europe"},
		"Brazil": {"continent": "South America"}, "Serbia": {"continent": "Europe"}, "Switzerland": {"continent": "Europe"}, "Cameroon": {"continent": "Africa"},
		"Portugal": {"continent": "Europe"}, "Ghana": {"continent": "Africa"}, "Uruguay": {"continent": "South America"}, "South Korea": {"continent": "Asia"}}

	match_info["tournament"] = "FIFA World Cup"
	match_info["home_team"] = home
	match_info["away_team"] = away
	match_info["home_team_continent"] = countries_info[home]["continent"]
	match_info["away_team_continent"] = countries_info[away]["continent"]
}

// No estoy pudiendo usar esta misma funcion pero dejandola en el archivo util
func findIndex(array []string, value string) int {
	for p, v := range array {
		if v == value {
			return p
		}
	}
	return -1
}
