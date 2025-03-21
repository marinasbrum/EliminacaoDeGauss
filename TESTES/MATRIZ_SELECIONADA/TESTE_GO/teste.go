package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const MAXN = 2000

var N int
var A [MAXN][MAXN]float64
var B [MAXN]float64
var X [MAXN]float64

func timeSeed() uint32 {
	return uint32(time.Now().UnixNano() % int64(^uint(0)>>1))
}

func parameters() {
	seed := 0
	rand.Seed(int64(timeSeed()))

	if len(os.Args) == 3 {
		var err error
		seed, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid seed input.")
			os.Exit(1)
		}
		rand.Seed(int64(seed))
		fmt.Printf("Random seed = %d\n", seed)
	}

	if len(os.Args) >= 2 {
		var err error
		N, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Invalid matrix dimension input.")
			os.Exit(1)
		}
		if N < 1 || N > MAXN {
			fmt.Printf("N = %d is out of range.\n", N)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Usage: %s <matrix_dimension> [random seed]\n", os.Args[0])
		os.Exit(1)
	}
	fmt.Printf("\nMatrix dimension N = %d.\n", N)
}

// Function to initialize inputs (Matrix A and vector B)
func initializeTestData() {
	A[0][0] = 2.0
	A[0][1] = 1.0
	A[0][2] = 1.0
	A[0][3] = 0.0
	A[0][4] = 0.0

	A[1][0] = 1.0
	A[1][1] = 3.0
	A[1][2] = 0.0
	A[1][3] = 1.0
	A[1][4] = 0.0

	A[2][0] = 1.0
	A[2][1] = 0.0
	A[2][2] = 2.0
	A[2][3] = 1.0
	A[2][4] = 0.0

	A[3][0] = 0.0
	A[3][1] = 1.0
	A[3][2] = 1.0
	A[3][3] = 3.0
	A[3][4] = 1.0

	A[4][0] = 0.0
	A[4][1] = 0.0
	A[4][2] = 0.0
	A[4][3] = 1.0
	A[4][4] = 2.0

	B[0] = 5.0
	B[1] = 7.0
	B[2] = 6.0
	B[3] = 8.0
	B[4] = 4.0

	for i := 0; i < MAXN; i++ {
		X[i] = 0.0
	}
}

func printInputs() {
	if N < 10 {
		fmt.Println("\nA =")
		for row := 0; row < N; row++ {
			fmt.Printf("\t")
			for col := 0; col < N; col++ {
				fmt.Printf("%5.2f", A[row][col])
				if col < N-1 {
					fmt.Print(", ")
				} else {
					fmt.Println(";")
				}
			}
		}
		fmt.Print("\nB = [")
		for col := 0; col < N; col++ {
			fmt.Printf("%5.2f", B[col])
			if col < N-1 {
				fmt.Print("; ")
			} else {
				fmt.Println("]")
			}
		}
	}
}

func printX() {
	if N < 100 {
		fmt.Print("\nX = [")
		for row := 0; row < N; row++ {
			fmt.Printf("%5.2f", X[row])
			if row < N-1 {
				fmt.Print("; ")
			} else {
				fmt.Println("]")
			}
		}
	}
}

func gauss() {
	var multiplier float64
	fmt.Println("Computing Serially.")
	for norm := 0; norm < N-1; norm++ {
		for row := norm + 1; row < N; row++ {
			multiplier = A[row][norm] / A[norm][norm]
			for col := norm; col < N; col++ {
				A[row][col] -= A[norm][col] * multiplier
			}
			B[row] -= B[norm] * multiplier
		}
	}
	for row := N - 1; row >= 0; row-- {
		X[row] = B[row]
		for col := N - 1; col > row; col-- {
			X[row] -= A[row][col] * X[col]
		}
		X[row] /= A[row][row]
	}
}

func main() {
	start := time.Now()

	startCPU := time.Now()

	parameters()

	initializeTestData()

	printInputs()

	fmt.Println("\nStarting clock.")
	gauss()

	printX()

	elapsed := time.Since(start)
	fmt.Printf("\nElapsed time = %g ms.\n", float64(elapsed.Milliseconds()))
	fmt.Printf("(CPU times are accurate to the nearest %g ms)\n", float64(time.Second)/float64(time.Millisecond))

	elapsedCPU := time.Since(startCPU)
	totalCPU := elapsedCPU.Seconds() * 1000
	fmt.Printf("My total CPU time = %g ms.\n", totalCPU)
}
