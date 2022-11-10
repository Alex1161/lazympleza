package logic

import (
	"encoding/csv"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetMostWorldCups(home string, away string) float64 {
	f, err := os.Open("../WorldCups.csv")
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
	f, err := os.Open("../WorldCupMatches.csv")
	check(err)

	fileReader := csv.NewReader(f)
	records, err := fileReader.ReadAll()
	check(err)
	f.Close()

	home_wins := 0
	away_wins := 0

	/*
		5 Home Team Name
		6 Home Team Goals
		7 Away Team Goals
		8 Away Team Name
	*/
	for i := 1; i < len(records); i++ {
		if (records[i][5] == home && records[i][6] > records[i][7]) || (records[i][8] == home && records[i][7] > records[i][6]) {
			home_wins++
		} else if (records[i][5] == away && records[i][6] > records[i][7]) || (records[i][8] == away && records[i][7] > records[i][6]) {
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
	f, err := os.Open("../WorldCupMatches.csv")
	check(err)

	fileReader := csv.NewReader(f)
	records, err := fileReader.ReadAll()
	check(err)
	f.Close()

	home_wins := 0
	away_wins := 0

	/*
		5 Home Team Name
		6 Home Team Goals
		7 Away Team Goals
		8 Away Team Name
	*/
	for i := 1; i < len(records); i++ {
		if records[i][5] == home && records[i][8] == away {
			if records[i][6] > records[i][7] {
				home_wins++
			} else if records[i][6] < records[i][7] {
				away_wins++
			}
		} else if records[i][5] == away && records[i][8] == home {
			if records[i][6] > records[i][7] {
				away_wins++
			} else if records[i][6] < records[i][7] {
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
