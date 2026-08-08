package keysheet

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"

	"github.com/ccaroon/enigmachine/enigma"
)

type Entry struct {
	Day            int        `yaml:"day"`
	Rotors         [3]string  `yaml:"rotors"`
	RingSettings   [3]int     `yaml:"rings"`
	PlugboardSwaps [10]string `yaml:"swaps"`
	DayKeys        [4]string  `yaml:"day_keys"`
}

// type KeySheet [31]Entry // On Entry per Month Day

// Generate an Entry
func GenerateEntry() *Entry {

	// Rotors -- Choose 3
	var rotors [3]string
	rotorChoices := []string{"I", "II", "III", "IV", "V"}

	for i := range 3 {
		idx := rand.IntN(len(rotorChoices))
		rotors[i] = rotorChoices[idx]

		rotorChoices = slices.Delete(rotorChoices, idx, idx+1)
	}

	// Ring Settings -- Choose 3 from (1-26)
	var rings [3]int
	for i := range 3 {
		num := rand.IntN(25) + 1
		for slices.Index(rings[:], num) != -1 {
			num = rand.IntN(25) + 1
		}
		rings[i] = num
	}

	// Plugboard -- Choose 10 Letter Pairs; no letter used more than once
	var swaps [10]string
	letters := []rune(enigma.ALPHABET)

	for i := range 10 {
		idx1 := rand.IntN(len(letters))
		ltr1 := letters[idx1]
		letters = slices.Delete(letters, idx1, idx1+1)

		idx2 := rand.IntN(len(letters))
		ltr2 := letters[idx2]
		letters = slices.Delete(letters, idx2, idx2+1)

		swaps[i] = fmt.Sprintf("%c%c", ltr1, ltr2)
	}

	// Rotor Top Letter Groups // Day Keys -- 4 Groups of 3 letters
	// Indivdual groups can contain the same letter
	// Different groups can contain the same letters
	// TODO: No two groups should be the same
	var dayKeys [4]string
	letters = []rune(enigma.ALPHABET)
	for i := range 4 {
		idx1 := rand.IntN(len(letters))
		ltr1 := letters[idx1]

		idx2 := rand.IntN(len(letters))
		ltr2 := letters[idx2]

		idx3 := rand.IntN(len(letters))
		ltr3 := letters[idx3]

		dayKeys[i] = fmt.Sprintf("%c%c%c", ltr1, ltr2, ltr3)
	}

	entry := Entry{
		Rotors:         rotors,
		RingSettings:   rings,
		PlugboardSwaps: swaps,
		DayKeys:        dayKeys,
	}

	return &entry
}

// Print an Entry
func (entry *Entry) Print() {
	rotors := strings.Join(entry.Rotors[:], " ")
	rings := fmt.Sprintf(
		"%02d %02d %02d",
		entry.RingSettings[0],
		entry.RingSettings[1],
		entry.RingSettings[2],
	)
	swaps := strings.Join(entry.PlugboardSwaps[:], " ")
	dayKeys := strings.Join(entry.DayKeys[:], " ")

	fmt.Printf("| %2d. | %v | %v | %v | %v |\n", entry.Day, rotors, rings, swaps, dayKeys)
}

// Generate a Monthly Key Sheet
// Read a Key Sheet
// Print a Key Sheet using ASCII with Headers, etc
