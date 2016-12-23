package main

import(
	"fmt"
	"io/ioutil"
	"strings"
)

type field [][]string

func check(err error) {
	if err != nil {
		fmt.Print(err)
	}
}

func (f *field) start(tick int) {
	for i := 0 ; i < tick ; i++ {
		f.calcTick()
	}
}

func (f *field) calcTick() {
	nextGenField := make([][]string, cap(*f))

	for i := 0 ; i < len(*f) ; i++ {
		nextGenField[i] = make([]string, cap((*f)[i]))
		for j := 0 ; j < len(*f) ; j++ {
			aliveNeighbors := f.calcNumberOfNeighbors(j,i)
			switch aliveNeighbors {
					/* 1. Any live cell with fewer than two live neighbours dies, as if by loneliness. */
				case 0: fallthrough
				case 1: nextGenField[i][j] = "-"
					/* 3. Any live cell with two or three live neighbours lives, unchanged, to the next generation. */
				case 2: nextGenField[i][j] = (*f)[i][j]
					/* 4. Any dead cell with exactly three live neighbours comes to life. */
				case 3: nextGenField[i][j] = "X"
					/* 2. Any live cell with more than three live neighbours dies, as if by overcrowding. */
				case 4: fallthrough
				case 5: fallthrough
				case 6: fallthrough
				case 7: fallthrough
				case 8: fallthrough
				case 9: nextGenField[i][j] = "-"
			}
		}
	}
	
	*f = field(nextGenField)
	f.print()
}

func (f field) calcNumberOfNeighbors(x, y int) int {
	count := 0
	
	for i := -1 ; i <= 1 ; i++ {
	
		if (y == 0 && i == -1) || (y == cap(f)-1 && i == 1) {
			continue
		}
	
		for j := -1 ; j <= 1 ; j++ {
		
			if (x == 0 && j == -1) || (x == cap(f[y+i])-1 && j == 1) || (i == 0 && j == 0) {
				continue
			}
		
			if f[y+i][x+j] == "X" {
				count++
			}
		}
	}
	return count
}

func (f field) print() {
	for i := 0 ; i < len(f) ; i++ {
		for j := 0 ; j < len(f[i]) ; j++ {
			fmt.Printf("%v ",f[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func main() {

	fmt.Print("Enter matrix size: ")
	var inputSize int
	//_, err := fmt.Scanf("%d\n", &inputSize)

	//fmt.Println("Debug: input: ", inputSize, "length: ", len(inputField))
	
	matrix, err := ioutil.ReadFile("matrix.txt")
	check(err)
	
	matrixstr := string(matrix)

	inStrings := strings.Split(matrixstr,"\n")
	fmt.Printf("instringleng %v \n", len(inStrings))
	stringLength := len(inStrings[0])
	fmt.Println("Debug: string0: ", len(inStrings[0]))
	for i := 1 ; i < len(inStrings)-1 ; i++ {
		fmt.Println("stringother: ", len(inStrings[i]))
		if stringLength != len(inStrings[i]) {
			fmt.Println("Matrix width is different")
			return
		}

	}
	inputSize = len(inStrings[0])/2
	fmt.Println("inputSize: ", inputSize)
	inputField := make([][]string, inputSize)


	fmt.Println("Debug: matrixstr: ",matrixstr)
	inputString := strings.Fields(matrixstr)
	fmt.Println("Debug: inputString: ", inputString)
	for i := 0 ; i < inputSize ; i++ {
		inputField[i] = make([]string, inputSize)
		for j := 0 ; j < inputSize ; j++ {
			val := inputString[i*inputSize+j:i*inputSize+j+1]
			
			if val[0] == "1" {
				inputField[i][j] = "X"
			} else if val[0] == "0" {
				inputField[i][j] = "-"
			} else {
				panic("wrong input")
			}
			
		}
	}
	//fmt.Println("Debug: inputField: ", inputField)
			
	fmt.Print("Enter amount of generations: ")
	var inputTicks int
	_, err = fmt.Scanf("%d\n", &inputTicks)
	check(err)

	/*
	game := field(
		[][]string{
		[]string{"-","-","-"},
		[]string{"X","X","X"},
		[]string{"-","-","-"},
		})*/
	
	game := field(inputField)
	game.print()
	game.start(inputTicks)
}