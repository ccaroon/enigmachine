package enigma

import (
	"fmt"
	"strings"
)

type Key struct {
	ReflId   string
	RotorIds []string
	PbSpec   []string
}

func ParseKeySpec(keySpec string) *Key {
	key := strings.ToUpper(keyFlag)
	keyParts := strings.SplitN(key, ":", 3)
	reflector := keyParts[0]
	rotors := strings.SplitN(keyParts[1], ",", 3)
	pbSpec := strings.Split(keyParts[2], ",")
	fmt.Println(pbSpec)
}
