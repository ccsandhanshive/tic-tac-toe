package main

import (
	"fmt"
	"os"
)

func updateDashBoard(dashboard string, gameDashboardArray []string, poition int, symbol string) (string, []string) {
	gameDashboardArray[poition-1] = symbol

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
func gameStatus(gameDashboardArray []string, gameDashboard string, symbolChoice string) bool {
	var currentPosition int
	fmt.Println("Enter your position")
	fmt.Scanln(&currentPosition)
	//fmt.Println("current:", gameDashboardArray[currentPosition-1], "O")
	if gameDashboardArray[currentPosition-1] != "O" {
		if gameDashboardArray[currentPosition-1] != "*" {
			gameDashboard, gameDashboardArray = updateDashBoard(gameDashboard, gameDashboardArray, currentPosition, symbolChoice)
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
			} else if checkDraw(gameDashboardArray) {
				println(gameDashboard)
				fmt.Println("Game Over:Match Draw")
				return true
			}

			if symbolChoice == "O" {
				symbolChoice = "*"
			} else {
				symbolChoice = "O"
			}
			fmt.Println("Updated")
		} else {
			fmt.Println("Invalid Choice")
		}
	} else {
		fmt.Println("Invalid Choice")
	}
	println(gameDashboard)
	gameStatus(gameDashboardArray, gameDashboard, symbolChoice)
	return true
}
func humanMode(gameDashboard string, gameDashboardArray []string) {
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
	gameStatus(gameDashboardArray, gameDashboard, symbolChoice)
}
func main() {

	var playmode int
	gameDashboardArray := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	gameDashboard := fmt.Sprintf("| %s | %s | %s |\n__________\n| %s | %s | %s |\n__________\n| %s | %s | %s |", gameDashboardArray[0], gameDashboardArray[1], gameDashboardArray[2], gameDashboardArray[3], gameDashboardArray[4], gameDashboardArray[5], gameDashboardArray[6], gameDashboardArray[7], gameDashboardArray[8])
	fmt.Println(".SYS")
	fmt.Println("Tic Toc Toe")
	fmt.Println("1.Play with computer\n2.Play with human\n3.Any other number for Exit\nEnter your Choice:")
	fmt.Scanln(&playmode)
	if playmode == 1 {
		fmt.Println("Comming Soon")
	} else {
		humanMode(gameDashboard, gameDashboardArray)
	}

}
