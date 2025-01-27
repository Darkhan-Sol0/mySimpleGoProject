package keyEvent

import (
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

// KeyEvent считывает нажатия клавиш в терминале и возвращает их
func KeyEvent() ([]byte, error) {
	// Установка режима терминала
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return nil, err
	}
	defer term.Restore(fd, oldState)

	// Обработка сигналов для выхода
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		os.Exit(0)
	}()

	// Чтение символов из терминала
	var buf [3]byte
	_, err = os.Stdin.Read(buf[:])
	if err != nil {
		return nil, err
	}

	return buf[:], nil
}
