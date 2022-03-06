package main

import (
	"fmt"
	"os"
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

func Load() (error, []Reversort) {
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

	cj := []Reversort{}

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

		cj = append(cj, Reversort{X: x, Y: y, S: s, CaseNumber: i})
	}

	return nil, cj

}

type Reversort struct {
	CaseNumber int

	LenElements int
	Elements    []int
}

func (c Reversort) String() string {

	if Debug {
		return fmt.Sprintf("Case #%d: %d (CJ=%d JC:%d S:%v -> S:%v)", c.CaseNumber, 0, c.X, c.Y, c.S, "")
	}

	return fmt.Sprintf("Case #%d: %d", c.CaseNumber, "")
}
