package decisionTree

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func get_entropy(resultados series.Series) (int, float64) {
	total := resultados.Len()
	win, _ := resultados.Compare(series.Eq, "Win").Bool()
	draw, _ := resultados.Compare(series.Eq, "Draw").Bool()
	lose, _ := resultados.Compare(series.Eq, "Lose").Bool()

	win_filtered := 0
	draw_filtered := 0
	lose_filtered := 0
	for i := 0; i < total; i++ {
		if win[i] {
			win_filtered++
		}

		if draw[i] {
			draw_filtered++
		}

		if lose[i] {
			lose_filtered++
		}
	}

	p_win := float64(win_filtered) / float64(total)
	p_draw := float64(draw_filtered) / float64(total)
	p_lose := float64(lose_filtered) / float64(total)

	calc_aux_win := 0.
	if win_filtered != 0 {
		calc_aux_win = -(p_win * math.Log2(p_win))
	} else {
		calc_aux_win = 0.
	}

	calc_aux_draw := 0.
	if draw_filtered != 0 {
		calc_aux_draw = -(p_draw * math.Log2(p_draw))
	} else {
		calc_aux_draw = 0.
	}

	calc_aux_lose := 0.
	if lose_filtered != 0 {
		calc_aux_lose = -(p_lose * math.Log2(p_lose))
	} else {
		calc_aux_lose = 0.
	}

	return total, calc_aux_win + calc_aux_draw + calc_aux_lose
}

func setFromList(list []string) (set []string) {
	ks := make(map[string]bool)

	for _, e := range list {
		if _, v := ks[e]; !v {
			ks[e] = true
			set = append(set, e)
		}
	}
	return
}

func uniqueGotaSeries(s series.Series) series.Series {
	return series.New(setFromList(s.Records()), s.Type(), s.Name)
}

func get_entropy_values(df dataframe.DataFrame, column string, value string, target string, ch chan entropy_info) {
	df_values := df.Filter(
		dataframe.F{
			Colname:    column,
			Comparator: series.Eq,
			Comparando: value,
		},
	)

	total, entropy := get_entropy(df_values.Col(target))

	ch <- entropy_info{column, value, total, entropy}
}

func get_info_gain_column(df dataframe.DataFrame, column string, target string, ch chan info_gain_info, initial_entropy_info entropy_info) {
	ch_entropy := make(chan entropy_info)

	values := uniqueGotaSeries(df.Col(column))
	for _, value := range values.Records() {
		go get_entropy_values(df, column, value, target, ch_entropy)
	}

	entropys_per_value := []entropy_info{}
	for i := 0; i < len(values.Records()); i++ {
		entropy := <-ch_entropy
		entropys_per_value = append(entropys_per_value, entropy)
	}

	info_gain := initial_entropy_info.entropy

	for _, entropy_info := range entropys_per_value {
		info_gain -= entropy_info.entropy * float64(entropy_info.total) / float64(initial_entropy_info.total)
	}

	ch <- info_gain_info{column, info_gain}
}

func fit(df dataframe.DataFrame, columns []string, target string, depth int, tree *decision_tree) {
	if depth == 0 {
		return
	}

	ch := make(chan info_gain_info)

	total, initial_entropy := get_entropy(df.Col(target))
	initial_entropy_info := entropy_info{"init", "init", total, initial_entropy}

	for _, column := range columns {
		go get_info_gain_column(df, column, target, ch, initial_entropy_info)
	}

	results := []info_gain_info{}
	for i := 0; i < len(columns); i++ {
		result := <-ch
		results = append(results, result)
	}

	maxValue := 0.0
	maxColumn := ""
	for _, info_gain_per_column := range results {
		if info_gain_per_column.info_gain > maxValue {
			maxValue = info_gain_per_column.info_gain
			maxColumn = info_gain_per_column.column
		}
	}

	if maxColumn == "" {
		columnIndex := 0
		for i, row := range df.Records() {
			if i == 0 {
				columnIndex = findIndex(row, target)
				continue
			}
			tree.column = row[columnIndex]
			break
		}
		tree.childs = nil

	} else {
		tree.column = maxColumn

		values := uniqueGotaSeries(df.Col(maxColumn))
		value_dt_array := []value_decision_tree{}
		for i := 0; i < len(values.Records()); i++ {
			dt := decision_tree{"", nil}
			vdt := value_decision_tree{values.Records()[i], dt}
			value_dt_array = append(value_dt_array, vdt)
		}
		tree.childs = value_dt_array

		for i, child := range tree.childs {
			//Le paso el dataframe filtrado por el valor value
			df_filtered := df.Filter(
				dataframe.F{
					Colname:    maxColumn,
					Comparator: series.Eq,
					Comparando: child.value,
				},
			)

			fit(df_filtered, columns, target, depth-1, &tree.childs[i].column)
		}
	}
}

type decision_tree struct {
	column string
	childs []value_decision_tree
}

type value_decision_tree struct {
	value  string
	column decision_tree
}

type info_gain_info struct {
	column    string
	info_gain float64
}

type entropy_info struct {
	column  string
	value   string
	total   int
	entropy float64
}

func _createDecisionTree() decision_tree {
	file, err := os.Open("./datasets/international_matches.csv")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	df := dataframe.ReadCSV(file)

	COLS := []string{"home_team", "away_team", "home_team_continent", "away_team_continent", "tournament"}

	tree := decision_tree{"", nil}

	fit(df, COLS, "home_team_result", 4, &tree)

	return tree
}

func predict(match_info map[string]string, tree decision_tree) string {
	poda := tree

	for len(poda.childs) > 0 {
		attr := poda.column
		info, _ := match_info[attr]
		found := false
		for i, child := range poda.childs {
			if child.value == info {
				found = true
				poda = poda.childs[i].column
			}
		}
		if !found {
			//Simulamos el resultado del partido ya que es un partido bastante peleado (le damos mas peso al empate)
			source := rand.NewSource(45)
			rnd := rand.New(source)
			simulation := rnd.Float64()
			if simulation <= 0.25 {
				poda.column = "Win"
			} else if simulation <= 0.75 {
				poda.column = "Draw"
			} else {
				poda.column = "Lose"
			}
			poda.childs = nil
			break
		}
	}

	return poda.column
}

func CreateDecisionTree() chan string {
	tree := make(chan string)

	go func() {
		dt := _createDecisionTree()

		for request := range tree {
			fmt.Println(request)
			r := ToRequest(request)
			var result string

			switch r.request {
			case "winner":
				result = GetWinnerWorldCup(dt)

			case "winner_between":
				match_info := make(map[string]string)
				match_info["home_team"] = r.params[0]
				match_info["away_team"] = r.params[1]
				match_info["home_continent"] = "South America"
				match_info["away_team"] = "South America"
				match_info["tournament"] = "FIFA World Cup"
				result = predict(match_info, dt)
			}
			tree <- result
		}
	}()

	return tree
}
