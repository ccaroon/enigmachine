package enigma

import (
	"fmt"
	"strings"
)

type Key struct {
	Reflector      rune     `yaml:"reflector"`
	Rotors         []string `yaml:"rotors"`
	RingSettings   []int    `yaml:"rings"`
	PlugboardSwaps []string `yaml:"swaps"`
}

// B:I,II,III:AZ,BJ

// TODO: does not handled ring yet
func ParseKeySpec(keySpec string) (*Key, error) {
	key := strings.ToUpper(keySpec)
	keyParts := strings.SplitN(key, ":", 3)
	keyLen := len(keyParts)

	// Reflector : B or C
	var reflector rune
	if keyLen > 0 && keyParts[0] != "" {
		reflector = rune(keyParts[0][0])
	} else {
		return nil, fmt.Errorf("Invalid Key Specification: [%s]", keySpec)
	}

	// Rotors
	// I,II,III
	var rotors []string
	if keyLen > 1 {
		rotors = strings.SplitN(keyParts[1], ",", 3)
		if len(rotors) != 3 {
			return nil, fmt.Errorf("Invalid Rotor Specs: [%s]", keyParts[1])
		}
	} else {
		return nil, fmt.Errorf("Invalid Key Specification. Missing Rotor Specs: [%s] ", keySpec)
	}

	// Plugboard
	// AZ,BJ
	var pbSpec []string = make([]string, 0, 10)
	if keyLen > 2 {
		pbSpec = strings.Split(keyParts[2], ",")
	}

	return &Key{
		Reflector:      reflector,
		Rotors:         rotors,
		PlugboardSwaps: pbSpec,
	}, nil
}
