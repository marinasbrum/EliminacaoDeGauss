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

func initializeInputs() {
	fmt.Println("\nInitializing...")
	for col := 0; col < N; col++ {
		for row := 0; row < N; row++ {
			A[row][col] = rand.Float64()
		}
		B[col] = rand.Float64()
		X[col] = 0.0
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

	initializeInputs()

	printInputs()

	fmt.Println("\nStarting clock.")
	gauss()

	printX()

	elapsed := time.Since(start)
	//fmt.Printf("\nElapsed time = %g ms.\n", float64(elapsed.Milliseconds()))
	//fmt.Printf("(CPU times are accurate to the nearest %g ms)\n", float64(time.Second)/float64(time.Millisecond))

	elapsedCPU := time.Since(startCPU)
	totalCPU := elapsedCPU.Seconds() * 1000
	//fmt.Printf("My total CPU time = %g ms.\n", totalCPU)
	//fmt.Printf("--------------------------------------------\n")

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("=================================================\n"))
	file.WriteString(fmt.Sprintf("Matrix dimension N = %d\n", N))
	file.WriteString(fmt.Sprintf("Elapsed time = %.3f ms.\n", float64(elapsed.Milliseconds())/1000)) // Convertendo para segundos
	file.WriteString(fmt.Sprintf("(CPU times are accurate to the nearest %.3f ms)\n", float64(time.Second)/float64(time.Millisecond)))
	file.WriteString(fmt.Sprintf("My total CPU time = %.3f ms.\n", float64(totalCPU)/1000)) // Convertendo para segundos
	file.WriteString(fmt.Sprintf("=================================================\n"))
}
