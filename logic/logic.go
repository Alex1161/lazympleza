package logic

import (
	"encoding/csv"
	"math/rand"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetMostWorldCups(home string, away string) float64 {
	f, err := os.Open("./datasets/WorldCups.csv")
	check(err)

	fileReader := csv.NewReader(f)
	records, err := fileReader.ReadAll()
	check(err)
	f.Close()

	home_wins := 0
	away_wins := 0

	/*
		2 Winner
	*/
	for i := 1; i < len(records); i++ {
		if records[i][2] == home {
			home_wins++
		} else if records[i][2] == away {
			away_wins++
		}
	}

	if home_wins > away_wins {
		return -1
	} else if away_wins > home_wins {
		return 1
	} else {
		return 0
	}
}

func GetMostWinMatches(home string, away string) float64 {
	f, err := os.Open("./datasets/international_matches.csv")
	check(err)

	fileReader := csv.NewReader(f)
	records, err := fileReader.ReadAll()
	check(err)
	f.Close()

	home_wins := 0
	away_wins := 0

	/*
		1 Home Team Name
		9 Home Team Goals
		10 Away Team Goals
		2 Away Team Name
	*/
	for i := 1; i < len(records); i++ {
		if (records[i][1] == home && records[i][9] > records[i][10]) || (records[i][2] == home && records[i][10] > records[i][9]) {
			home_wins++
		} else if (records[i][1] == away && records[i][9] > records[i][10]) || (records[i][2] == away && records[i][10] > records[i][9]) {
			away_wins++
		}
	}

	if home_wins > away_wins {
		return -1
	} else if away_wins > home_wins {
		return 1
	} else {
		return 0
	}
}

func GetRecordWinner(home string, away string) float64 {
	f, err := os.Open("./datasets/international_matches.csv")
	check(err)

	fileReader := csv.NewReader(f)
	records, err := fileReader.ReadAll()
	check(err)
	f.Close()

	home_wins := 0
	away_wins := 0

	/*
		1 Home Team Name
		9 Home Team Goals
		10 Away Team Goals
		2 Away Team Name
	*/
	for i := 1; i < len(records); i++ {
		if records[i][1] == home && records[i][2] == away {
			if records[i][9] > records[i][10] {
				home_wins++
			} else if records[i][9] < records[i][10] {
				away_wins++
			}
		} else if records[i][1] == away && records[i][2] == home {
			if records[i][9] > records[i][10] {
				away_wins++
			} else if records[i][9] < records[i][10] {
				home_wins++
			}
		}
	}

	if home_wins > away_wins {
		return -1
	} else if away_wins > home_wins {
		return 1
	} else {
		return 0
	}
}

func GetWinnerBetween(home string, away string) (string, float64) {
	mwc := GetMostWorldCups(home, away)
	mwm := GetMostWinMatches(home, away)
	rw := GetRecordWinner(home, away)

	analysis := 0.5*mwc + 0.25*mwm + 0.25*rw

	if analysis < -0.25 {
		probability := (-0.25 - analysis) / 0.75
		return home, probability
	} else if analysis > 0.25 {
		probability := (analysis - 0.25) / 0.75
		return away, probability
	} else {
		penaltyWinner := rand.Float64()
		if penaltyWinner < 0.5 {
			return home, -1
		} else {
			return away, -1
		}
	}
}
