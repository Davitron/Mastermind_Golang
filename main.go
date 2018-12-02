package main

import (
	"strings"
	"os"
	"bufio"
	"fmt"
	"time"
	"math/rand"
	"regexp"
	"sort"
)

// This is the ruberic for mastermind game
// 

func main() {
	gameColors := []string{"R", "O", "Y", "G", "B", "I", "V"}
	gameRound := 8

	var clue []string

	gameCode := generateGameCode(gameColors)
	fmt.Println("game code", gameCode)
	fmt.Println("Game code is generated, You have 8 tries to crack the code")

	for gameRound >= 1 {
		
		fmt.Printf("You have %d more tries to crack the code\n", gameRound)
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Submit you Code(With so whitespaces)")
		userInput, _ := reader.ReadString('\n')
		userInput= strings.TrimSuffix(userInput, "\n")
		userInput = strings.ToUpper(userInput)

		isInputValid, _ := regexp.MatchString("^[A-Za-z]{5}$", userInput)

		fmt.Println(isInputValid)

		if isInputValid == false {

			fmt.Println("Invalid Entry", len(userInput))
			continue
		}

		fmt.Println("Confirm Entry with 'Y' or cancel with any other key")

		input, _ := reader.ReadString('\n')
		confirmed, _ := regexp.MatchString("^[Yy]", input)

		if confirmed {
			userInputArr := strings.Split(userInput, "")[0:5]
			clue = runCodeCheck(userInputArr, gameCode)
		} else {
			continue
		}

		if !contains("W", clue) || !contains("X", clue) {
			fmt.Println("You Win")
			break
		}
		
		sort.Strings(clue)
		fmt.Println(clue)
		gameRound--
				

	}
}

func generateGameCode(arr []string) []string {
	fmt.Println("arg", arr)
	size := len(arr)
	rand.Seed( time.Now().UnixNano())
	for i := size - 1; i > 0; i-- {
		index := rand.Intn(i + 1)
		tmp := arr[index]
		arr[index] = arr[i]
		arr[i] = tmp
	}
	return arr[0:5] 
}

func contains(element string, arr []string) bool {
	exists := false
	for _, item := range arr {
		if item == element {
			exists = true
		}
	}
	return exists
}

func runCodeCheck(input []string, code []string) []string {
	var clue []string
	for index := range code {
		if input[index] == code[index] {
			clue = append(clue, "B")
		} else if contains(input[index], code) {
			clue = append(clue, "W")
		} else {
			clue = append(clue, "X")
		}
	}
	return clue
}


