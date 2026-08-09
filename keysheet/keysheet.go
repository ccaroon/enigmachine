package keysheet

import (
	"fmt"
	"strings"

	"github.com/bykof/gostradamus"
)

type KeySheet struct {
	Network string     `yaml:"network"`
	Month   int        `yaml:"month"`
	Year    int        `yaml:"year"`
	Entries [31]*Entry `yaml:"entries"`
}

// Generate a Monthly Key Sheet
func GenerateKeySheet(network string, month, year int) *KeySheet {
	targetDT := gostradamus.NewLocalDateTime(year, month, 1, 0, 0, 0, 0).CeilMonth()
	numDays := targetDT.Day()

	keySheet := KeySheet{
		Network: network,
		Month:   month,
		Year:    year,
	}

	for dayNum := range numDays {
		entry := GenerateEntry()
		entry.Day = dayNum + 1
		keySheet.Entries[dayNum] = entry
	}

	return &keySheet
}

func (ks *KeySheet) Print() {
	dt := gostradamus.NewLocalDateTime(ks.Year, ks.Month, 1, 0, 0, 0, 0)
	dtStamp := dt.Format("MMMM YYYY")

	// Length of a line == Length of one entry
	lineLen := len(ks.Entries[0].Format())
	divider := fmt.Sprintf("+%s+", strings.Repeat("-", lineLen-2))

	// Header1
	header1 := "Nur für den Dienstgebrauch"
	netLen := len(ks.Network)
	h1Len := len(header1)
	filler := strings.Repeat(" ", int(lineLen-(h1Len*2)-netLen)/2)
	fmt.Printf("%s%s%s%s%s\n", header1, filler, strings.ToUpper(ks.Network), filler, header1)

	// Header2
	header2 := "Achtung! Streng Geheim!"
	h2Len := len(header2)
	dtLen := len(dtStamp)
	filler = strings.Repeat(" ", int(lineLen-(dtLen*2)-h2Len)/2)
	fmt.Printf("%s%s%s%s%s\n", dtStamp, filler, header2, filler, dtStamp)

	// Divider
	fmt.Println(divider)

	// Table Header
	fmt.Printf("| %-3s | %-13s | %-8s | %-29s | %-15s |\n",
		"Day",
		"Rotors",
		"Rings",
		"Plugboard",
		"Day Keys",
	)

	// Divider
	fmt.Println(divider)

	// Entries
	for i := len(ks.Entries) - 1; i >= 0; i-- {
		ks.Entries[i].Print()
	}

	// Footer / Final Divider
	fmt.Println(divider)

}

// TODO: Read a Key Sheet

// TODO: Save a Key Sheet
func (ks *KeySheet) Save(path string) {

}
