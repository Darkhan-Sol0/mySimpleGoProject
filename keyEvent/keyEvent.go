package keyEvent

import (
	"bytes"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

type KeyCode struct {
	Esc_key     int
	A_up_key    int
	A_down_key  int
	A_left_key  int
	A_right_key int
}

func (k *KeyCode) Init() KeyCode {
	return KeyCode{
		Esc_key:     27,
		A_up_key:    65,
		A_down_key:  66,
		A_left_key:  67,
		A_right_key: 68,
	}
}

func KeyEvent() (int, error) {
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return 0, err
	}
	defer term.Restore(fd, oldState)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		os.Exit(0)
	}()

	var buf [3]byte
	_, err = os.Stdin.Read(buf[:])
	if err != nil {
		return 0, err
	}

	var keyCode KeyCode
	return GetKey(buf[:], keyCode.Init()), nil
}

func GetKey(buf []byte, keyCode KeyCode) (Key int) {

	switch {
	case bytes.Equal(buf[:], []byte{27, 0, 0}):
		Key = keyCode.Esc_key
	case bytes.Equal(buf[:], []byte{27, 91, 65}):
		Key = keyCode.A_up_key
	case bytes.Equal(buf[:], []byte{27, 91, 66}):
		Key = keyCode.A_down_key
	case bytes.Equal(buf[:], []byte{27, 91, 67}):
		Key = keyCode.A_left_key
	case bytes.Equal(buf[:], []byte{27, 91, 68}):
		Key = keyCode.A_right_key
	default:
		Key = int(buf[0])
	}
	return
}
