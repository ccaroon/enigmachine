package enigma

import (
	"strings"
)

type Key struct {
	ReflId   string
	RotorIds []string
	RotorCfg string
	PbSpec   []string
}

// B:I@A,II@A,III@A:AZ,BJ

func ParseKeySpec(keySpec string) *Key {
	key := strings.ToUpper(keySpec)
	keyParts := strings.SplitN(key, ":", 3)

	// B
	reflector := keyParts[0]

	// I@A,II@A,II@A
	var rotors []string
	var rotorCfg string
	rotorSpec := strings.SplitN(keyParts[1], ",", 3)
	for _, rSpec := range rotorSpec {
		rId, rCfg, _ := strings.Cut(rSpec, "@")
		rotors = append(rotors, rId)

		// Default to "A"
		if rCfg == "" {
			rCfg = "A"
		}
		rotorCfg += rCfg
	}

	// AZ,BJ
	pbSpec := strings.Split(keyParts[2], ",")

	return &Key{
		ReflId:   reflector,
		RotorIds: rotors,
		RotorCfg: rotorCfg,
		PbSpec:   pbSpec,
	}
}
