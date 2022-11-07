package server

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
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
