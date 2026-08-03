package enigma

import (
	"fmt"
	"strings"
)

type Key struct {
	ReflId   string
	RotorIds []string
	RotorCfg string
	PbSpec   []string
}

// B:I@A,II@A,III@A:AZ,BJ

func ParseKeySpec(keySpec string) (*Key, error) {
	key := strings.ToUpper(keySpec)
	keyParts := strings.SplitN(key, ":", 3)
	keyLen := len(keyParts)

	// Reflector : B or C
	var reflector string
	if keyLen > 0 && keyParts[0] != "" {
		reflector = keyParts[0]
	} else {
		return nil, fmt.Errorf("Invalid Key Specification: [%s]", keySpec)
	}

	// Rotors
	// I@A,II@A,II@A
	var rotors []string = make([]string, 0, 3)
	var rotorCfg string
	if keyLen > 1 {
		rotorSpec := strings.SplitN(keyParts[1], ",", 3)
		if len(rotorSpec) != 3 {
			return nil, fmt.Errorf("Invalid Rotor Specs: [%s]", keyParts[1])
		}

		for _, rSpec := range rotorSpec {
			rId, rCfg, _ := strings.Cut(rSpec, "@")

			if rId != "" {
				rotors = append(rotors, rId)
			} else {
				return nil, fmt.Errorf("Invalid Rotor: [%s]", rSpec)
			}

			// Default to "A"
			if rCfg == "" {
				rCfg = "A"
			}
			rotorCfg += rCfg

		}
	} else {
		return nil, fmt.Errorf("Invalid Key Specification. Missing Rotor Specs: [%s] ", keySpec)
	}

	// Plugboard
	// AZ,BJ
	var pbSpec []string
	if keyLen > 2 {
		pbSpec = strings.Split(keyParts[2], ",")
	}

	return &Key{
		ReflId:   reflector,
		RotorIds: rotors,
		RotorCfg: rotorCfg,
		PbSpec:   pbSpec,
	}, nil
}
