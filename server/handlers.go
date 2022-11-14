package server

import (
	"encoding/csv"
	"fmt"
	"lazympleza/logic"
	"net/http"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetWinnerBetween(w http.ResponseWriter, r *http.Request, home string, away string) {
	mwc := logic.GetMostWorldCups(home, away)
	mwm := logic.GetMostWinMatches(home, away)
	rw := logic.GetRecordWinner(home, away)

	analysis := 0.5*mwc + 0.25*mwm + 0.25*rw

	if analysis < -0.25 {
		probability := (-0.25 - analysis) / 0.75
		fmt.Fprintf(w, "El ganador entre "+home+" y "+away+" es: "+home+" con una certeza de "+strconv.FormatInt(int64(probability*100), 10)+"%")
	} else if analysis > 0.25 {
		probability := (analysis - 0.25) / 0.75
		fmt.Fprintf(w, "El ganador entre "+home+" y "+away+" es: "+away+" con una certeza de "+strconv.FormatInt(int64(probability*100), 10)+"%")
	} else {
		fmt.Fprintf(w, "El partido entre "+home+" y "+away+" termina en un empate despues de un partido muy peleado")
	}
}

func GetWorldCupsData(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("../WorldCups.csv")
	check(err)

	fileReader := csv.NewReader(f)
	records, error := fileReader.ReadAll()
	check(error)

	fmt.Fprintf(w, "La cantidad de datos de mundiales es:"+strconv.FormatInt(int64((len(records)-1)), 10))

	f.Close()
}

func GetWorldCupsMatches(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("../WorldCupMatches.csv")
	check(err)

	fileReader := csv.NewReader(f)
	records, error := fileReader.ReadAll()
	check(error)

	fmt.Fprintf(w, "La cantidad de datos de partidos es:"+strconv.FormatInt(int64((len(records)-1)), 10))

	f.Close()
}

func GetWorldCupsPlayers(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("../WorldCupPlayers.csv")
	check(err)

	fileReader := csv.NewReader(f)
	records, error := fileReader.ReadAll()
	check(error)

	fmt.Fprintf(w, "La cantidad de datos de jugadores es:"+strconv.FormatInt(int64((len(records)-1)), 10))

	f.Close()
}
