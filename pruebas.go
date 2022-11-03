package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
	"time"
)

func getWorldCupsData() {
	f, err := os.Open("./WorldCups.csv")
	check(err)

	fileReader := csv.NewReader(f)
	records, error := fileReader.ReadAll()
	check(error)

	fmt.Println("La cantidad de datos de mundiales es:", len(records)-1)

	f.Close()
}

func getWorldCupsMatches() {
	f, err := os.Open("./WorldCupMatches.csv")
	check(err)

	fileReader := csv.NewReader(f)
	records, error := fileReader.ReadAll()
	check(error)

	fmt.Println("La cantidad de datos de partidos es:", len(records)-1)

	f.Close()
}

func getWorldCupsPlayers() {
	f, err := os.Open("./WorldCupPlayers.csv")
	check(err)

	fileReader := csv.NewReader(f)
	records, error := fileReader.ReadAll()
	check(error)

	fmt.Println("La cantidad de datos de jugadores es:", len(records)-1)

	f.Close()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Pruebas sin go keyword")

	start := time.Now()
	getWorldCupsData()
	getWorldCupsMatches()
	getWorldCupsPlayers()
	t := time.Now()
	fmt.Println("Tardo:", t.Sub(start))

	fmt.Println("------------")

	fmt.Println("Pruebas con go keyword")
	start_go := time.Now()

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		f, err := os.Open("./WorldCups.csv")
		check(err)

		fileReader := csv.NewReader(f)
		records, error := fileReader.ReadAll()
		check(error)

		fmt.Println("La cantidad de datos de mundiales es:", len(records)-1)

		f.Close()
	}()
	go func() {
		defer wg.Done()
		f, err := os.Open("./WorldCupMatches.csv")
		check(err)

		fileReader := csv.NewReader(f)
		records, error := fileReader.ReadAll()
		check(error)

		fmt.Println("La cantidad de datos de partidos es:", len(records)-1)

		f.Close()
	}()
	go func() {
		defer wg.Done()
		f, err := os.Open("./WorldCupPlayers.csv")
		check(err)

		fileReader := csv.NewReader(f)
		records, error := fileReader.ReadAll()
		check(error)

		fmt.Println("La cantidad de datos de jugadores es:", len(records)-1)

		f.Close()
	}()
	wg.Wait()
	t_go := time.Now()
	fmt.Println("Tardo:", t_go.Sub(start_go))
}
