package main

import (
	"fmt"
	"os"
	"strings"
)

var Debug bool

//Moons and Umbrellas (5pts, 11pts, 1pts)
//https://codingcompetitions.withgoogle.com/codejam/round/000000000043580a/00000000006d1145
func main() {

	args := os.Args
	if len(args) > 1 {
		if args[1] == "--debug" {
			Debug = true
			fmt.Println("-- Debug Mode --")
		}
	}

	_, cj := Load()

	for _, v := range cj {
		fmt.Println(v)
	}

}

func Load() (error, []CodyJamal) {
	//se sono in modalità debug verrà letto il file con l'input al posto che lo stdin
	stdin := os.Stdin
	if Debug {
		f, err := os.Open("test")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		stdin = f
	}

	cj := []CodyJamal{}

	var testCases int
	_, err := fmt.Fscanf(stdin, "%d", &testCases)
	if err != nil {
		return err, nil
	}

	for i := 1; i <= testCases; i++ {
		var x int
		var y int
		var s string

		_, err := fmt.Fscanf(stdin, "%d %d %v", &x, &y, &s)
		if err != nil {
			return err, nil
		}

		cj = append(cj, CodyJamal{X: x, Y: y, S: s, CaseNumber: i})
	}

	return nil, cj

}

type CodyJamal struct {
	CaseNumber int

	X int
	Y int
	S string
}

func (c CodyJamal) SplitStringIntoArray(S string) []rune {
	return []rune(S)
}

func (c CodyJamal) CalculateCost(StringArray []rune) int {

	//each time CJ appears in the mural, Cody-Jamal must pay X
	//each time JC appears in the mural, he must pay Y
	var cost int
	lenArray := len(StringArray)

	for i := range StringArray {
		if lenArray == i+1 {
			break
		}

		var sb strings.Builder
		sb.WriteRune(StringArray[i])
		sb.WriteRune(StringArray[i+1])

		switch sb.String() {
		case "CJ":
			cost = cost + c.X
		case "JC":
			cost = cost + c.Y
		}
	}

	return cost
}

func (c CodyJamal) Solve(StringArray []rune) []rune {

	//each time CJ appears in the mural, Cody-Jamal must pay X
	//each time JC appears in the mural, he must pay Y
	var result []rune = StringArray
	lenArray := len(StringArray)

	//mi posiziono sul primo ? che abbiamo prima o dopo una C o una J
	//se non ci sono caratteri diversi da ? inizializzo il primo carattere con quello con il costo più basso se negativo
	//se nessun carattere ha un ?, termino il loop e restituisco il  risultato
	//calcolo il valore minimo da assegnare al ?
	//ripeto il tutto in un loop infinito

	for {

		pos, pos2 := c.Position(result)
		if pos == -1 {
			break
		}
		prev := '?'
		next := '?'
		size := 1

		if pos-1 >= 0 {
			prev = StringArray[pos-1]
		}

		if pos+1 <= lenArray-1 {
			next = StringArray[pos+1]
		}

		if pos2 != -1 {
			size = 2

			if pos-1 >= 0 {
				prev = StringArray[pos-1]
			}

			if pos2+1 <= lenArray-1 {
				next = StringArray[pos2+1]
			}

		}

		if size == 1 {
			result[pos], _ = c.MinimumValue(prev, next, size)
		} else {
			result[pos], result[pos2] = c.MinimumValue(prev, next, size)
		}
	}

	return result
}

//calcola il valore minimo considerando il carattere precedente e quello successivo
//se non esiste il carattere precedente, successivo o entrambi devono essere impostati con un '?'
func (c CodyJamal) MinimumValue(previous rune, next rune, size int) (rune, rune) {
	//each time CJ appears in the mural, Cody-Jamal must pay X
	//each time JC appears in the mural, he must pay Y
	if size == 1 {
		//se non ci sono caratteri, inizializzo con la coppia dal costo minore
		if previous == '?' && next == '?' {
			if c.X < c.Y {
				return 'C', ' '
			} else {
				return 'J', ' '
			}
		}

		jCase := []rune{previous, 'J', next}
		cCase := []rune{previous, 'C', next}
		if c.CalculateCost(jCase) <= c.CalculateCost(cCase) {
			return 'J', ' '
		}
		return 'C', ' '

	} else {
		test := [][]rune{{'J', 'C'}, {'J', 'J'}, {'C', 'J'}, {'C', 'C'}}
		min := 0
		var r1 rune
		var r2 rune
		for i, v := range test {
			testCase := []rune{previous, v[0], v[1], next}
			if i == 0 {
				min = c.CalculateCost(testCase)
				r1 = v[0]
				r2 = v[1]
			} else {
				if c.CalculateCost(testCase) <= min {
					min = c.CalculateCost(testCase)
					r1 = v[0]
					r2 = v[1]
				}
			}
		}
		return r1, r2
	}
}

//restituisce la prima posizione con un ? che abbia C o J vicini
//se non è presente restituisce -1
func (c CodyJamal) Position(StringArray []rune) (int, int) {
	lenArray := len(StringArray)
	result := -1
	result2 := -1

	for i := range StringArray {
		//se non mi trovo un un ? continuo
		if StringArray[i] != '?' {
			continue
		}

		//se l'array ha un solo carattere
		if lenArray == 1 {
			if StringArray[i] == '?' {
				result = i
				break
			}
		}

		//se sono nella prima posizione
		if i == 0 {
			if StringArray[i+1] != '?' {
				result = i
				break
			}
			continue
		}

		//se sono nell'ultima posizione
		if i == lenArray-1 {
			if StringArray[i-1] != '?' {
				result = i
				break
			}
			//se arrivato all'ultima posizione ho trovato solo ? restituisco la prima posizione
			result = 0
			break
		}

		//in tutti gli altri casi
		if StringArray[i+1] != '?' {
			result = i
			break
		}
		if StringArray[i-1] != '?' {
			result = i
			break
		}

	}

	if result != -1 {
		if result-1 >= 0 && StringArray[result-1] == '?' {
			result2 = result - 1
		}

		if result+1 <= lenArray-1 && StringArray[result+1] == '?' {
			result2 = result + 1
		}
	}

	if result < result2 || result2 == -1 {
		return result, result2
	} else {
		return result2, result
	}
}

func (c CodyJamal) String() string {
	arr := c.SplitStringIntoArray(c.S)
	solve := c.Solve(arr)

	if Debug {
		return fmt.Sprintf("Case #%d: %d (CJ=%d JC:%d S:%v -> S:%v)", c.CaseNumber, c.CalculateCost(solve), c.X, c.Y, c.S, string(solve))
	}

	return fmt.Sprintf("Case #%d: %d", c.CaseNumber, c.CalculateCost(solve))
}
