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
func updateDashBoard(dashboard string, gameDashboardArray []string, poition int, symbol string, humanMode bool, level int) (string, []string) {
	gameDashboardArray[poition-1] = symbol
	var randomPotion int
	if humanMode == false {
		dashboard = fmt.Sprintf("| %s | %s | %s |\n__________\n| %s | %s | %s |\n__________\n| %s | %s | %s |", gameDashboardArray[0], gameDashboardArray[1], gameDashboardArray[2], gameDashboardArray[3], gameDashboardArray[4], gameDashboardArray[5], gameDashboardArray[6], gameDashboardArray[7], gameDashboardArray[8])
		if CheckWin(gameDashboardArray, dashboard, symbol) {
			os.Exit(1)
		}
		if symbol == "O" {
			symbol = "*"
		} else {
			symbol = "O"
		}
		if level == 1 {
			randomPotion = getRandomPosition(gameDashboardArray, poition)

		} else if level == 2 {
			randomPotion = getRandomPosition(gameDashboardArray, poition) //Comming Soon

		} else if level == 3 {
			randomPotion = getRandomPosition(gameDashboardArray, poition) //Comming Soon
		}
		gameDashboardArray[randomPotion-1] = symbol
		if CheckWin(gameDashboardArray, dashboard, symbol) {
			//fmt.Println(fmt.Sprintf("| %s | %s | %s |\n__________\n| %s | %s | %s |\n__________\n| %s | %s | %s |", gameDashboardArray[0], gameDashboardArray[1], gameDashboardArray[2], gameDashboardArray[3], gameDashboardArray[4], gameDashboardArray[5], gameDashboardArray[6], gameDashboardArray[7], gameDashboardArray[8]))
			os.Exit(1)
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
	if gameDashboardArray[1-1] == gameDashboardArray[2-1] && gameDashboardArray[1-1] == gameDashboardArray[3-1] {

		fmt.Println("Game Over:" + symbolChoice + " Win")
		println(gameDashboard)
		return true
	} else if gameDashboardArray[1-1] == gameDashboardArray[4-1] && gameDashboardArray[1-1] == gameDashboardArray[7-1] {
		fmt.Println("Game Over:" + symbolChoice + " Win")
		println(gameDashboard)
		return true
	} else if gameDashboardArray[1-1] == gameDashboardArray[5-1] && gameDashboardArray[1-1] == gameDashboardArray[9-1] {
		fmt.Println("Game Over:" + symbolChoice + " Win")
		println(gameDashboard)
		return true
	} else if gameDashboardArray[3-1] == gameDashboardArray[6-1] && gameDashboardArray[3-1] == gameDashboardArray[9-1] {
		fmt.Println("Game Over:" + symbolChoice + " Win")
		println(gameDashboard)
		return true
	} else if gameDashboardArray[3-1] == gameDashboardArray[5-1] && gameDashboardArray[3-1] == gameDashboardArray[7-1] {
		fmt.Println("Game Over:" + symbolChoice + " Win")
		println(gameDashboard)
		return true
	} else if gameDashboardArray[2-1] == gameDashboardArray[5-1] && gameDashboardArray[2-1] == gameDashboardArray[8-1] {
		fmt.Println("Game Over:" + symbolChoice + " Win")
		println(gameDashboard)
		return true
	} else if gameDashboardArray[4-1] == gameDashboardArray[5-1] && gameDashboardArray[4-1] == gameDashboardArray[6-1] {
		fmt.Println("Game Over:" + symbolChoice + " Win")
		println(gameDashboard)
		return true
	} else if checkDraw(gameDashboardArray) {
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

			fmt.Println("Updated")
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
			os.Exit(1)
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
	} else {
		humanMode(gameDashboard, gameDashboardArray, true, level)
	}

}
