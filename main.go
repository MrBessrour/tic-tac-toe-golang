package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	board := [9]string{}
	gameOver := false
	var err error
	player := "O"
	turnNumber := 0
	var winner string

	for gameOver != true {
		PrintBoard(board)
		move := askforplay()
		err, board = play(move, player, board)
		if err != nil {
			fmt.Println(err)
			continue
		}
		player = SwitchPlayers(player)
		turnNumber++
		gameOver, winner = CheckForWinner(board, turnNumber)
		fmt.Println(winner)
	}
	PrintBoard(board)
	if winner == "" {
		fmt.Println("it's a draw ")
	} else {
		fmt.Printf("YaaY %s is winner ", winner)
	}
}

func CheckForWinner(b [9]string, n int) (bool, string) {

	test := false
	i := 0

	//horizantel test
	for i < 9 {
		test = b[i] == b[i+1] && b[i+1] == b[i+2] && b[i] != ""
		if !test {
			i += 3
		} else {
			return true, b[i]
		}
	}
	i = 0
	//vertical test
	for i < 3 {
		test = b[i] == b[i+3] && b[i+3] == b[i+6] && b[i] != ""
		if !test {
			i += 1
		} else {
			return true, b[i]
		}
	}
	//diagonal 1 test
	if b[0] == b[4] && b[4] == b[8] && b[0] != "" {
		return true, b[i]
	}
	//diagonal 2 test
	if b[2] == b[4] && b[4] == b[6] && b[2] != "" {
		return true, b[i]
	}
	if n == 9 {
		return true, ""
	}
	return false, ""
}

func ClearScreen() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}

func SwitchPlayers(player string) string {
	if player == "O" {
		return "X"
	}
	return "O"
}

func play(pos int, player string, b [9]string) (error, [9]string) {
	if b[pos-1] == "" {
		b[pos-1] = player
		return nil, b
	}
	return errors.New("try another move"), b
}

func askforplay() int {
	var moveInt int
	//for moveInt < 0 || moveInt > 9 {
	fmt.Println("Enter Pos to play: ")
	fmt.Scan(&moveInt)
	//}
	return moveInt
}

func PrintBoard(b [9]string) {
	ClearScreen()
	for i, v := range b {
		if v == "" {
			fmt.Printf(" ")
		} else {
			fmt.Printf(v)
		}

		if i > 0 && (i+1)%3 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Print("|")
		}
	}

}
