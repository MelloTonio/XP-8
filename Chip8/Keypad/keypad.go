package Keypad

type Keypad struct {
	key string
}

func New() *Keypad {
	return new(Keypad)
}
