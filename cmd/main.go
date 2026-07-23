package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type gameState struct {
	doorOpened   bool
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
		fmt.Println(handleCommand(command))
	}
}

func handleCommand(action string) string {

	switch {
	case strings.EqualFold(action, "осмотреться"):
		return lookAround()
	case strings.EqualFold(action, "надеть рюкзак"):
		if game.currentPlace == "комната" {
			game.backpackOn = true
			return "вы надели: рюкзак"
		} else {
			return "рюкзак лежит в комнате"
		}
	case strings.HasPrefix(action, "идти"):
		return move(action)
	case strings.HasPrefix(action, "взять"):
		if game.currentPlace == "комната" {
			return take(action)
		} else {
			return "нет такого"
		}
	case strings.HasPrefix(action, "применить"):
		return use(action)
	default:
		return "неизвестная команда"
	}

}

func initGame() {
	game.doorOpened = false
	game.backpackOn = false
	game.items = map[string]bool{"ключи": false, "конспекты": false}

	game.currentPlace = "кухня"
	//отправить приветствие
	fmt.Print("Добро пожаловать в игру! \n" +
		"Вам нужно покуинть дом взяв всё что нужно в универе \n" +
		"чтобы узнать команды введите Справка\n")
}

func lookAround() string {
	switch game.currentPlace {
	case "кухня":
		return "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"
	case "коридор":
		return "ты в коридоре. можно пройти - комната, кухня"
	case "комната":
		// 	TODO упростить
		switch {
		case game.items["ключ"] == false && game.items["конспекты"] == false:
			return "на столе: ключи, конспекты. можно пройти - коридор"
		case game.items["ключ"] == true && game.items["конспекты"] == false:
			return "на столе: конспекты. можно пройти - коридор"
		case game.items["ключ"] == false && game.items["конспекты"] == true:
			return "на столе: конспекты. можно пройти - коридор"
		case game.items["ключ"] == true && game.items["конспекты"] == true:
			return "пустая комната. можно пройти - коридор"
		default:
			return fmt.Sprint("что-то не так")
		}
	default:
		return fmt.Sprint("что-то не так")
	}
}

func move(action string) string {
	destination, _ := strings.CutPrefix(action, "идти ")
	if strings.EqualFold(destination, game.currentPlace) {
		return fmt.Sprintf("ты уже здесь")
	}

	switch game.currentPlace {
	case "кухня":
		switch destination {
		case "коридор":
			game.currentPlace = "коридор"
			return "ничего интересного. можно пройти - кухня, комната, улица"
		default:
			return fmt.Sprintf("нет пути в %s", destination)
		}

	case "коридор":
		switch destination {
		case "кухня":
			game.currentPlace = "кухня"
			return fmt.Sprintf("ты зашел %s", destination)
		case "комната":
			game.currentPlace = "комната"
			return "ты в своей комнате. можно пройти - коридор"
		case "улица":
			if game.doorOpened == true {
				return "на улице весна. можно пройти - домой"
			} else {
				return "дверь закрыта"
			}
		default:
			return fmt.Sprintf("нет пути в %s", destination)
		}

	case "комната":
		switch destination {
		case "коридор":
			game.currentPlace = "коридор"
			return fmt.Sprintf("ты зашел %s", destination)
		default:
			return fmt.Sprintf("нет пути в %s", destination)
		}

	case "улица":
		switch destination {
		case "домой":
			game.currentPlace = "коридор"
			return fmt.Sprintf("ты зашел %s", destination)
		default:
			return fmt.Sprintf("нет пути в %s", destination)
		}
	default:
		return fmt.Sprintf("что-то не так")
	}

}

func take(action string) string {
	item, _ := strings.CutPrefix(action, "взять ")
	if game.backpackOn == false {
		fmt.Println("некуда класть")
		return
	}
	if game.items[item] == false {
		game.items[item] = true
		return fmt.Sprintf("предмет добавлен в инвентарь: %s", item)
	} else {
		return fmt.Sprintf("нет такого")
	}
}

func use(action string) string {
	words := strings.Fields(action)

	if len(words) < 3 {
		return fmt.Sprintf("что-то не так")
	}

	if words[1] == "ключ" && game.items[words[1]] == false {
		return fmt.Sprintf("нет предмета в инвентаре - %s", words[1])
	} else if words[2] == "дверь" && game.currentPlace == "коридор" {
		game.doorOpened = true
		return fmt.Sprintf("дверь открыта")
	} else {
		return fmt.Sprintf("не к чему применить")
	}
}
