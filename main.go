package main

import (
	"fmt"
	"myProg/keyEvent"
	"os"
	"time"
)

type GameSetting struct {
	fps   int
	delta float64
}

func InitGame() GameSetting {
	return GameSetting{
		fps:   60,
		delta: 1.0 / 60.0,
	}
}

func readKeys(keyChan chan int) {
	for {
		key, err := keyEvent.KeyEvent()
		if err == nil {
			keyChan <- key
		}
	}
}

func GameLoop(a GameSetting, keyChan <-chan int) {
	timeout := time.NewTicker(time.Duration(a.delta * float64(time.Second)))

	for {
		select {
		case <-timeout.C:
			// fmt.Printf("NO\n")
		case key := <-keyChan:
			KeyEvent(key)
		}
	}
}

func main() {
	keyChan := make(chan int)
	go readKeys(keyChan)
	GameLoop(InitGame(), keyChan)
}

func KeyEvent(key int) {
	switch key {
	case 27:
		os.Exit(0)
	case 65:
		fmt.Print("вверх ")
	case 66:
		fmt.Print("вниз ")
	case 67:
		fmt.Print("вправо ")
	case 68:
		fmt.Print("влево ")
	default:
		fmt.Print("Клавиша ", key, " ")
	}
}
