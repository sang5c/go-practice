package main

import (
	"fmt"
	"numbering-game/game"
)

func main() {
	for {
		fmt.Print("숫자 값을 입력하세요: ")
		input := 0
		_, err := fmt.Scanln(&input)
		if err != nil {
			return
		}

		numberingGame := game.New()
		compare, count := numberingGame.Compare(input)
		if compare == 0 {
			fmt.Println("숫자를 맞췄습니다. 축하합니다. 시도한 횟수: ", count)
			break
		} else if compare == -1 {
			fmt.Println("입력하신 숫자가 더 작습니다.")
		} else {
			fmt.Println("입력하신 숫자가 더 큽니다.")
		}
	}

}
