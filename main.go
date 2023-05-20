package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func getRandomPosition(gameDashboardArray []string, poition int) int {
	var randomArray []int
	for index, value := range gameDashboardArray {

		if value != "*" && value != "O" && index != poition {
			intVal, _ := strconv.Atoi(value)
			randomArray = append(randomArray, intVal)
		}
	}
	rand.Seed(time.Now().Unix())
	return randomArray[rand.Intn(len(randomArray))]
}
func Contains(randomArray []int, value int) bool {
	for i := 0; i < len(randomArray); i++ {
		if randomArray[i] == value {
			return true
		}
	}
	return false
}

func stringToInt(value string) int {
	intVal, _ := strconv.Atoi(value)
	return intVal
}
func getDefencePosition(gameDashboardArray []string, poition int, opponentSymbol string) int {
	var randomArray []int
	for index, value := range gameDashboardArray {

		if value != "*" && value != "O" && index != poition-1 {
			intVal, _ := strconv.Atoi(value)
			randomArray = append(randomArray, intVal)
		}
	}
	//For 1 as common
	if gameDashboardArray[1-1] == gameDashboardArray[2-1] && gameDashboardArray[1-1] == opponentSymbol {
		//fmt.Println("Tested 1-2", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[3-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	if gameDashboardArray[1-1] == gameDashboardArray[3-1] && gameDashboardArray[1-1] == opponentSymbol {
		//fmt.Println("Tested 1-3", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[2-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	if gameDashboardArray[1-1] == gameDashboardArray[4-1] && gameDashboardArray[1-1] == opponentSymbol {
		//fmt.Println("Tested 1-4", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[7-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	if gameDashboardArray[1-1] == gameDashboardArray[5-1] && gameDashboardArray[1-1] == opponentSymbol {
		//fmt.Println("Tested 1-5", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[9-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}

	if gameDashboardArray[1-1] == gameDashboardArray[7-1] && gameDashboardArray[1-1] == opponentSymbol {
		//fmt.Println("Tested 1-7", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[4-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}

	if gameDashboardArray[1-1] == gameDashboardArray[9-1] && gameDashboardArray[1-1] == opponentSymbol {
		//fmt.Println("Tested 1-9", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[5-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	//For 2 as common
	if gameDashboardArray[2-1] == gameDashboardArray[3-1] && gameDashboardArray[2-1] == opponentSymbol {
		//fmt.Println("Tested 2-3", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[1-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}

	if gameDashboardArray[2-1] == gameDashboardArray[5-1] && gameDashboardArray[2-1] == opponentSymbol {
		//fmt.Println("Tested 2-5", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[8-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}

	if gameDashboardArray[2-1] == gameDashboardArray[8-1] && gameDashboardArray[2-1] == opponentSymbol {
		//fmt.Println("Tested 2-8", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[5-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}

	//For 3 common

	if gameDashboardArray[3-1] == gameDashboardArray[5-1] && gameDashboardArray[3-1] == opponentSymbol {
		//fmt.Println("Tested 3-5", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[7-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	if gameDashboardArray[3-1] == gameDashboardArray[6-1] && gameDashboardArray[3-1] == opponentSymbol {
		//fmt.Println("Tested 3-6", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[9-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	if gameDashboardArray[3-1] == gameDashboardArray[7-1] && gameDashboardArray[3-1] == opponentSymbol {
		//fmt.Println("Tested 3-5", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[5-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	if gameDashboardArray[3-1] == gameDashboardArray[9-1] && gameDashboardArray[3-1] == opponentSymbol {
		//fmt.Println("Tested 3-9", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[6-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	//For 4 common
	if gameDashboardArray[4-1] == gameDashboardArray[5-1] && gameDashboardArray[4-1] == opponentSymbol {
		//fmt.Println("Tested 4-5", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[6-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	if gameDashboardArray[4-1] == gameDashboardArray[6-1] && gameDashboardArray[4-1] == opponentSymbol {
		//fmt.Println("Tested 4-6", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[5-1]
		//fmt.Println(randomArray)
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	if gameDashboardArray[4-1] == gameDashboardArray[7-1] && gameDashboardArray[4-1] == opponentSymbol {
		//fmt.Println("Tested 4-7", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[1-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}

	//For 5 common
	if gameDashboardArray[5-1] == gameDashboardArray[6-1] && gameDashboardArray[5-1] == opponentSymbol {
		//fmt.Println("Tested 5-6", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[4-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	if gameDashboardArray[5-1] == gameDashboardArray[7-1] && gameDashboardArray[5-1] == opponentSymbol {
		//fmt.Println("Tested 5-6", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[3-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	if gameDashboardArray[5-1] == gameDashboardArray[8-1] && gameDashboardArray[5-1] == opponentSymbol {
		//fmt.Println("Tested 5-8", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[2-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	if gameDashboardArray[5-1] == gameDashboardArray[9-1] && gameDashboardArray[5-1] == opponentSymbol {
		//fmt.Println("Tested 5-9", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[1-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	//For 6 common
	if gameDashboardArray[6-1] == gameDashboardArray[9-1] && gameDashboardArray[6-1] == opponentSymbol {
		//fmt.Println("Tested 6-9", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[3-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	//For 7 common
	if gameDashboardArray[7-1] == gameDashboardArray[8-1] && gameDashboardArray[7-1] == opponentSymbol {
		//fmt.Println("Tested 7-8", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[9-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	if gameDashboardArray[7-1] == gameDashboardArray[9-1] && gameDashboardArray[7-1] == opponentSymbol {
		//fmt.Println("Tested 7-8", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[8-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	//For 8 common
	if gameDashboardArray[8-1] == gameDashboardArray[9-1] && gameDashboardArray[8-1] == opponentSymbol {
		//fmt.Println("Tested 8-9", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[7-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	//Odd conditions

	return -1
}
func getDefencePositionOddeConditions(gameDashboardArray []string, poition int, opponentSymbol string) int {
	var randomArray []int
	for index, value := range gameDashboardArray {

		if value != "*" && value != "O" && index != poition-1 {
			intVal, _ := strconv.Atoi(value)
			randomArray = append(randomArray, intVal)
		}
	}
	if gameDashboardArray[1-1] == gameDashboardArray[6-1] && gameDashboardArray[1-1] == opponentSymbol {
		//fmt.Println("Tested 1-7", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[3-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
		considerValue = gameDashboardArray[9-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	if gameDashboardArray[1-1] == gameDashboardArray[8-1] && gameDashboardArray[1-1] == opponentSymbol {
		//fmt.Println("Tested 1-9", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[7-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
		considerValue = gameDashboardArray[9-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	if gameDashboardArray[2-1] == gameDashboardArray[4-1] && gameDashboardArray[2-1] == opponentSymbol {
		//fmt.Println("Tested 2-5", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[1-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
		considerValue = gameDashboardArray[7-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	if gameDashboardArray[2-1] == gameDashboardArray[6-1] && gameDashboardArray[2-1] == opponentSymbol {
		//fmt.Println("Tested 2-8", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[3-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
		considerValue = gameDashboardArray[9-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	if gameDashboardArray[2-1] == gameDashboardArray[7-1] && gameDashboardArray[2-1] == opponentSymbol {
		//fmt.Println("Tested 2-8", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[1-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
		considerValue = gameDashboardArray[3-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	if gameDashboardArray[2-1] == gameDashboardArray[9-1] && gameDashboardArray[2-1] == opponentSymbol {
		//fmt.Println("Tested 2-8", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[3-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
		considerValue = gameDashboardArray[5-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	if gameDashboardArray[3-1] == gameDashboardArray[4-1] && gameDashboardArray[3-1] == opponentSymbol {
		//fmt.Println("Tested 3-5", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[1-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
		considerValue = gameDashboardArray[5-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
		considerValue = gameDashboardArray[7-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	if gameDashboardArray[3-1] == gameDashboardArray[8-1] && gameDashboardArray[3-1] == opponentSymbol {
		//fmt.Println("Tested 3-5", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[9-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
		considerValue = gameDashboardArray[7-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
		considerValue = gameDashboardArray[5-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
		considerValue = gameDashboardArray[1-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	if gameDashboardArray[4-1] == gameDashboardArray[8-1] && gameDashboardArray[4-1] == opponentSymbol {
		//fmt.Println("Tested 4-7", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[7-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
		considerValue = gameDashboardArray[5-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
		considerValue = gameDashboardArray[1-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	if gameDashboardArray[4-1] == gameDashboardArray[9-1] && gameDashboardArray[4-1] == opponentSymbol {
		//fmt.Println("Tested 4-7", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[1-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
		considerValue = gameDashboardArray[5-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
		considerValue = gameDashboardArray[7-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
		considerValue = gameDashboardArray[6-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	if gameDashboardArray[6-1] == gameDashboardArray[7-1] && gameDashboardArray[6-1] == opponentSymbol {
		//fmt.Println("Tested 6-9", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[1-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
		considerValue = gameDashboardArray[5-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
		considerValue = gameDashboardArray[9-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
		considerValue = gameDashboardArray[3-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	if gameDashboardArray[6-1] == gameDashboardArray[8-1] && gameDashboardArray[6-1] == opponentSymbol {
		//fmt.Println("Tested 6-9", gameDashboardArray, opponentSymbol)
		considerValue := gameDashboardArray[9-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
		considerValue = gameDashboardArray[5-1]
		if Contains(randomArray, stringToInt(considerValue)) {
			return stringToInt(considerValue)
		}
	}
	return -1
}
func mediumLevel(gameDashboardArray []string, poition int, opponentSymbol string) int {
	returnedInt := getDefencePosition(gameDashboardArray, poition, opponentSymbol)
	if returnedInt != -1 {
		return returnedInt
	}
	return getRandomPosition(gameDashboardArray, poition)
}
func hardLevel(gameDashboardArray []string, poition int, symbol string, opponentSymbol string) int {
	selfWinPosition := getDefencePosition(gameDashboardArray, poition, symbol)
	if selfWinPosition != -1 {
		return selfWinPosition
	}
	blockOppoentWinPosition := getDefencePosition(gameDashboardArray, poition, opponentSymbol)
	if blockOppoentWinPosition != -1 {
		return blockOppoentWinPosition
	}
	selfWinOddPosition := getDefencePositionOddeConditions(gameDashboardArray, poition, symbol)
	if selfWinOddPosition != -1 {
		return blockOppoentWinPosition
	}
	opponentWinOddPosition := getDefencePositionOddeConditions(gameDashboardArray, poition, opponentSymbol)
	if opponentWinOddPosition != -1 {
		return opponentWinOddPosition
	}
	return getRandomPosition(gameDashboardArray, poition)
}
func updateDashBoard(dashboard string, gameDashboardArray []string, poition int, symbol string, humanMode bool, level int) (string, []string) {
	gameDashboardArray[poition-1] = symbol
	var randomPotion int
	if humanMode == false {
		dashboard = fmt.Sprintf("| %s | %s | %s |\n__________\n| %s | %s | %s |\n__________\n| %s | %s | %s |", gameDashboardArray[0], gameDashboardArray[1], gameDashboardArray[2], gameDashboardArray[3], gameDashboardArray[4], gameDashboardArray[5], gameDashboardArray[6], gameDashboardArray[7], gameDashboardArray[8])
		if CheckWin(gameDashboardArray, dashboard, symbol) {
			os.Exit(0)
		}
		opponentSymbol := symbol
		if symbol == "O" {
			symbol = "*"
		} else {
			symbol = "O"
		}
		if level == 1 {
			randomPotion = getRandomPosition(gameDashboardArray, poition)

		} else if level == 2 {
			randomPotion = mediumLevel(gameDashboardArray, poition, opponentSymbol)

		} else if level == 3 {
			randomPotion = hardLevel(gameDashboardArray, poition, symbol, opponentSymbol)
		}
		gameDashboardArray[randomPotion-1] = symbol
		dashboard = fmt.Sprintf("| %s | %s | %s |\n__________\n| %s | %s | %s |\n__________\n| %s | %s | %s |", gameDashboardArray[0], gameDashboardArray[1], gameDashboardArray[2], gameDashboardArray[3], gameDashboardArray[4], gameDashboardArray[5], gameDashboardArray[6], gameDashboardArray[7], gameDashboardArray[8])
		if CheckWin(gameDashboardArray, dashboard, symbol) {
			//fmt.Println(fmt.Sprintf("| %s | %s | %s |\n__________\n| %s | %s | %s |\n__________\n| %s | %s | %s |", gameDashboardArray[0], gameDashboardArray[1], gameDashboardArray[2], gameDashboardArray[3], gameDashboardArray[4], gameDashboardArray[5], gameDashboardArray[6], gameDashboardArray[7], gameDashboardArray[8]))
			os.Exit(0)
		}
		if symbol == "O" {
			symbol = "*"
		} else {
			symbol = "O"
		}
	}

	return fmt.Sprintf("| %s | %s | %s |\n__________\n| %s | %s | %s |\n__________\n| %s | %s | %s |", gameDashboardArray[0], gameDashboardArray[1], gameDashboardArray[2], gameDashboardArray[3], gameDashboardArray[4], gameDashboardArray[5], gameDashboardArray[6], gameDashboardArray[7], gameDashboardArray[8]), gameDashboardArray
}

func checkDraw(gameDashboardArray []string) bool {
	count := 0
	for i := 0; i < len(gameDashboardArray); i++ {
		if gameDashboardArray[i] == "O" || gameDashboardArray[i] == "*" {
			count += 1
		}
	}
	if count == 9 {
		return true
	}
	return false
}
func CheckWin(gameDashboardArray []string, gameDashboard string, symbolChoice string) bool {
	//For 1 common
	if gameDashboardArray[1-1] == gameDashboardArray[2-1] && gameDashboardArray[1-1] == gameDashboardArray[3-1] {

		fmt.Println("Game Over:" + symbolChoice + " Win")
		println(gameDashboard)
		return true
	}
	if gameDashboardArray[1-1] == gameDashboardArray[4-1] && gameDashboardArray[1-1] == gameDashboardArray[7-1] {

		fmt.Println("Game Over:" + symbolChoice + " Win")
		println(gameDashboard)
		return true
	}
	if gameDashboardArray[1-1] == gameDashboardArray[5-1] && gameDashboardArray[1-1] == gameDashboardArray[9-1] {

		fmt.Println("Game Over:" + symbolChoice + " Win")
		println(gameDashboard)
		return true
	}
	if gameDashboardArray[1-1] == gameDashboardArray[2-1] && gameDashboardArray[1-1] == gameDashboardArray[3-1] {

		fmt.Println("Game Over:" + symbolChoice + " Win")
		println(gameDashboard)
		return true
	}
	//For 2 common
	if gameDashboardArray[2-1] == gameDashboardArray[5-1] && gameDashboardArray[2-1] == gameDashboardArray[8-1] {

		fmt.Println("Game Over:" + symbolChoice + " Win")
		println(gameDashboard)
		return true
	}
	//For 3 common
	if gameDashboardArray[3-1] == gameDashboardArray[5-1] && gameDashboardArray[3-1] == gameDashboardArray[7-1] {

		fmt.Println("Game Over:" + symbolChoice + " Win")
		println(gameDashboard)
		return true
	}
	if gameDashboardArray[3-1] == gameDashboardArray[6-1] && gameDashboardArray[3-1] == gameDashboardArray[9-1] {

		fmt.Println("Game Over:" + symbolChoice + " Win")
		println(gameDashboard)
		return true
	}
	if gameDashboardArray[3-1] == gameDashboardArray[5-1] && gameDashboardArray[3-1] == gameDashboardArray[7-1] {

		fmt.Println("Game Over:" + symbolChoice + " Win")
		println(gameDashboard)
		return true
	}
	//For 4 coomon
	if gameDashboardArray[4-1] == gameDashboardArray[5-1] && gameDashboardArray[4-1] == gameDashboardArray[6-1] {

		fmt.Println("Game Over:" + symbolChoice + " Win")
		println(gameDashboard)
		return true
	}
	//For 7 common
	if gameDashboardArray[7-1] == gameDashboardArray[8-1] && gameDashboardArray[7-1] == gameDashboardArray[9-1] {

		fmt.Println("Game Over:" + symbolChoice + " Win")
		println(gameDashboard)
		return true
	}
	//Draw
	if checkDraw(gameDashboardArray) {
		println(gameDashboard)
		fmt.Println("Game Over:Match Draw")
		return true
	}
	return false
}
func gameStatus(gameDashboardArray []string, gameDashboard string, symbolChoice string, humanMode bool, level int) bool {
	var currentPosition int
	fmt.Println("Enter your position")
	fmt.Scanln(&currentPosition)
	//fmt.Println("current:", gameDashboardArray[currentPosition-1], "O")
	if gameDashboardArray[currentPosition-1] != "O" {
		if gameDashboardArray[currentPosition-1] != "*" {
			gameDashboard, gameDashboardArray = updateDashBoard(gameDashboard, gameDashboardArray, currentPosition, symbolChoice, humanMode, level)
			if CheckWin(gameDashboardArray, gameDashboard, symbolChoice) {
				return true
			}
			if humanMode == true {
				if symbolChoice == "O" {
					symbolChoice = "*"
				} else {
					symbolChoice = "O"
				}
			}

			//fmt.Println("Updated")
		} else {
			fmt.Println("Invalid Choice")
		}
	} else {
		fmt.Println("Invalid Choice")
	}
	println(gameDashboard)
	gameStatus(gameDashboardArray, gameDashboard, symbolChoice, humanMode, level)
	return true
}
func humanMode(gameDashboard string, gameDashboardArray []string, humanMode bool, level int) {
	var symbolChoice string
	fmt.Println("Select your symbol: \n1.O \n2.*\n3.Any other choice for Exit")
	fmt.Println("Enter your Symbol:")
	fmt.Scanln(&symbolChoice)
	if symbolChoice != "O" {
		if symbolChoice != "*" {
			os.Exit(0)
		}
	}
	fmt.Println("Game Started")
	fmt.Println(gameDashboard)
	gameStatus(gameDashboardArray, gameDashboard, symbolChoice, humanMode, level)
}

func main() {

	var playmode int
	level := 0
	gameDashboardArray := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	gameDashboard := fmt.Sprintf("| %s | %s | %s |\n__________\n| %s | %s | %s |\n__________\n| %s | %s | %s |", gameDashboardArray[0], gameDashboardArray[1], gameDashboardArray[2], gameDashboardArray[3], gameDashboardArray[4], gameDashboardArray[5], gameDashboardArray[6], gameDashboardArray[7], gameDashboardArray[8])
	fmt.Println("######################################################################################################################")
	fmt.Println("#                                                  .SYS                                                              #")
	fmt.Println("######################################################################################################################")
	fmt.Println("                                               Tic Tac Toe                                                            ")
	fmt.Println("1.Play with computer\n2.Play with human\n3.Any other number for Exit\nEnter your Choice:")
	fmt.Scanln(&playmode)
	if playmode == 1 {
		fmt.Println("Choose difficulty level:\n1.Easy\n2.Medium\n3.Hard\nEnter your choice:")
		fmt.Scanln(&level)
		humanMode(gameDashboard, gameDashboardArray, false, level)
	} else if playmode == 2 {
		humanMode(gameDashboard, gameDashboardArray, true, level)
	} else {
		os.Exit(0)
	}

}
