package keysheet

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"

	"github.com/ccaroon/enigmachine/enigma"
)

const (
	rotorCount  = 3
	pbSwapCount = 10
	dayKeyCount = 4
)

type Entry struct {
	Day        int `yaml:"day"`
	enigma.Key `yaml:"key"`
	DayKeys    []string `yaml:"day_keys"`
}

// Generate an Entry
func GenerateEntry() *Entry {

	// Rotors
	var rotors []string = make([]string, rotorCount)
	rotorChoices := []string{"I", "II", "III", "IV", "V"}

	for i := range rotorCount {
		idx := rand.IntN(len(rotorChoices))
		rotors[i] = rotorChoices[idx]

		rotorChoices = slices.Delete(rotorChoices, idx, idx+1)
	}

	// Ring Settings -- Choose from (1-26)
	var rings []int = make([]int, rotorCount)
	for i := range rotorCount {
		num := rand.IntN(25) + 1
		for slices.Index(rings[:], num) != -1 {
			num = rand.IntN(25) + 1
		}
		rings[i] = num
	}

	// reflChoices := [3]rune{'A', 'B', 'C'}
	// reflIdx := rand.IntN(len(reflChoices))
	// reflector := reflChoices[reflIdx]
	// TODO: Hard-coded...better way? What's the procedure?
	reflector := 'B'

	// Plugboard -- Choose Letter Pairs; no letter used more than once
	var swaps []string = make([]string, pbSwapCount)
	letters := []rune(enigma.ALPHABET)

	for i := range pbSwapCount {
		idx1 := rand.IntN(len(letters))
		ltr1 := letters[idx1]
		letters = slices.Delete(letters, idx1, idx1+1)

		idx2 := rand.IntN(len(letters))
		ltr2 := letters[idx2]
		letters = slices.Delete(letters, idx2, idx2+1)

		swaps[i] = fmt.Sprintf("%c%c", ltr1, ltr2)
	}

	// Rotor Top Letter Groups // Day Keys -- Groups of 3 letters
	// Indivdual groups can contain the same letter
	// Different groups can contain the same letters
	// TODO: No two groups should be the same
	var dayKeys []string = make([]string, dayKeyCount)
	letters = []rune(enigma.ALPHABET)
	for i := range dayKeyCount {
		idx1 := rand.IntN(len(letters))
		ltr1 := letters[idx1]

		idx2 := rand.IntN(len(letters))
		ltr2 := letters[idx2]

		idx3 := rand.IntN(len(letters))
		ltr3 := letters[idx3]

		dayKeys[i] = fmt.Sprintf("%c%c%c", ltr1, ltr2, ltr3)
	}

	entry := Entry{
		Key: enigma.Key{
			Rotors:         rotors,
			RingSettings:   rings,
			Reflector:      reflector,
			PlugboardSwaps: swaps,
		},
		DayKeys: dayKeys,
	}

	return &entry
}

func (entry *Entry) RandomDayKey() string {
	idx := rand.IntN(len(entry.DayKeys))

	return entry.DayKeys[idx]
}

func (entry *Entry) Format() string {
	rotors := fmt.Sprintf(
		"%c %3s %3s %3s",
		entry.Reflector,
		entry.Rotors[0],
		entry.Rotors[1],
		entry.Rotors[2],
	)

	rings := fmt.Sprintf(
		"%02d %02d %02d",
		entry.RingSettings[0],
		entry.RingSettings[1],
		entry.RingSettings[2],
	)
	swaps := strings.Join(entry.PlugboardSwaps[:], " ")
	dayKeys := strings.Join(entry.DayKeys[:], " ")

	return fmt.Sprintf("| %2d. | %s | %v | %v | %v |", entry.Day, rotors, rings, swaps, dayKeys)
}

// Print an Entry
func (entry *Entry) Print() {
	fmt.Println(entry.Format())
}
