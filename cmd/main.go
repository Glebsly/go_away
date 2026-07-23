package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type gameState struct {
	doorClosed   bool
	backpackOn   bool
	items        map[string]bool
	currentPlace string
}

var game gameState

func main() {

	initGame()

	scanner := bufio.NewScanner(os.Stdin)

	for {

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("ошибка чтения:", err)
			}
			break
		}

		command := scanner.Text()
		if command == "initGame()" {
			initGame()
			continue
		}
		doAction(command)
	}
}

func doAction(action string) {

	switch {
	case strings.EqualFold(action, "осмотреться"):
		lookAround()
	case strings.EqualFold(action, "надеть рюкзак"):
		if game.currentPlace == "комната" {
			game.backpackOn = true
			fmt.Print("вы надели: рюкзак")
		} else {
			fmt.Print("рюкзак лежит в комнате")
		}
	case strings.HasPrefix(action, "идти"):
		move(action)
	case strings.HasPrefix(action, "взять"):
		if game.currentPlace == "комната" {
			take(action)
		} else {
			fmt.Print("нет такого")
		}

	}

}

func initGame() {
	game.doorClosed = true
	game.backpackOn = false
	game.items = map[string]bool{"ключи": false, "конспекты": false}

	game.currentPlace = "кухня"
	//отправить приветствие
	fmt.Print("Добро пожаловать в игру! \n" +
		"Вам нужно покуинть дом взяв всё что нужно в универе \n" +
		"чтобы узнать команды введите Справка\n")
}

func lookAround() {
	switch game.currentPlace {
	case "кухня":
		fmt.Println("ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор")
	case "коридор":
		fmt.Println("ты в коридоре. можно пройти - комната, кухня")
	case "комната":
		// 	checkRoomItems() доработать проверку объектов в комнате
		fmt.Println("на столе: ключи, конспекты. можно пройти - коридор")
	}
}

func move(action string) {
	destination, _ := strings.CutPrefix(action, "идти ")
	if strings.EqualFold(destination, game.currentPlace) {
		fmt.Println("ты уже здесь")
	}

	switch game.currentPlace {
	case "кухня":
		switch destination {
		case "коридор":
			fmt.Println("ты зашел ", destination)
			game.currentPlace = "коридор"
		default:
			fmt.Println("нет пути в", destination)
		}

	case "коридор":
		switch destination {
		case "кухня":
			fmt.Println("ты зашел ", destination)
			game.currentPlace = "кухня"
		case "комната":
			fmt.Println("ты зашел ", destination)
			game.currentPlace = "комната"
		default:
			fmt.Println("нет пути в", destination)
		}

	case "комната":
		switch destination {
		case "коридор":
			fmt.Println("ты зашел ", destination)
			game.currentPlace = "коридор"
		default:
			fmt.Println("нет пути в", destination)
		}
	}
}

func take(action string) {
	item, _ := strings.CutPrefix(action, "взять ")

	if game.items[item] == false {
		game.items[item] = true
		fmt.Println("предмет добавлен в инвентарь: ", item)
	} else {
		fmt.Println("нет такого")
	}
}
