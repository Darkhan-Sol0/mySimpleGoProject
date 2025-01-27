package main

import (
	"bytes"
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

func readKeys(keyChan chan []byte) {
	for {
		key, err := keyEvent.KeyEvent()
		if err == nil {
			keyChan <- key
		}
	}
}

func GameLoop(a GameSetting, keyChan <-chan []byte) {
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
	keyChan := make(chan []byte)
	go readKeys(keyChan)
	GameLoop(InitGame(), keyChan)
}

func KeyEvent(key []byte) {
	// fmt.Println(key)
	switch {
	case bytes.Equal(key, []byte{27, 0, 0}):
		os.Exit(0)
	case bytes.Equal(key, []byte{27, 91, 65}):
		fmt.Println("вверх")
	case bytes.Equal(key, []byte{27, 91, 66}):
		fmt.Println("вниз")
	case bytes.Equal(key, []byte{27, 91, 67}):
		fmt.Println("вправо")
	case bytes.Equal(key, []byte{27, 91, 68}):
		fmt.Println("влево")
	default:
	}
}
