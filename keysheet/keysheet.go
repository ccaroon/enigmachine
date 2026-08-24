package keysheet

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/bykof/gostradamus"
	"github.com/ccaroon/enigmachine/util"
	"go.yaml.in/yaml/v3"
)

type KeySheet struct {
	Network string     `yaml:"network"`
	Month   int        `yaml:"month"`
	Year    int        `yaml:"year"`
	Entries [31]*Entry `yaml:"entries"`
}

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

func LoadKeySheet(ksPath string) (*KeySheet, error) {
	var keySheet KeySheet

	content, err := os.ReadFile(ksPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, &keySheet)

	return &keySheet, nil
}

func LoadActiveKeySheet(network string) (*KeySheet, error) {
	dt := gostradamus.Now()
	ksPath := BuildPath(network, dt.Month(), dt.Year())

	return LoadKeySheet(ksPath)
}

func BuildPath(network string, month, year int) string {
	dataDir := util.GetDataDir()

	re, err := regexp.Compile(`\W`)
	if err != nil {
		panic(err)
	}
	sanitizedNetwork := re.ReplaceAllString(network, "_")

	ksPath := fmt.Sprintf(
		"%s/enigmachine/keysheets/%s/%d-%02d.yml",
		dataDir, sanitizedNetwork,
		year,
		month,
	)

	return ksPath
}

func (ks *KeySheet) String() string {
	var output strings.Builder

	dt := gostradamus.NewLocalDateTime(ks.Year, ks.Month, 1, 0, 0, 0, 0)
	dtStamp := dt.Format("MMMM YYYY")

	// Length of a line == Length of one entry
	lineLen := len(ks.Entries[0].String())
	divider := fmt.Sprintf("+%s+", strings.Repeat("-", lineLen-2))

	// Header1
	header1 := "Nur für den Dienstgebrauch"
	netLen := len(ks.Network)
	h1Len := len(header1)
	filler := strings.Repeat(" ", int(lineLen-(h1Len*2)-netLen)/2)
	output.WriteString(
		fmt.Sprintf("%s%s%s%s%s\n", header1, filler, strings.ToUpper(ks.Network), filler, header1),
	)

	// Header2
	header2 := "Achtung! Streng Geheim!"
	h2Len := len(header2)
	dtLen := len(dtStamp)
	filler = strings.Repeat(" ", int(lineLen-(dtLen*2)-h2Len)/2)
	output.WriteString(
		fmt.Sprintf("%s%s%s%s%s\n", dtStamp, filler, header2, filler, dtStamp),
	)

	// Divider
	output.WriteString(divider)
	output.WriteString("\n")

	// Table Header
	output.WriteString(
		fmt.Sprintf("| %-3s | %-13s | %-8s | %-29s | %-15s |\n",
			"Day",
			"Rotors",
			"Rings",
			"Plugboard",
			"Day Keys",
		),
	)

	// Divider
	output.WriteString(divider)
	output.WriteString("\n")

	// Entries
	for i := len(ks.Entries) - 1; i >= 0; i-- {
		output.WriteString(
			ks.Entries[i].String(),
		)
		output.WriteString("\n")
	}

	// Footer / Final Divider
	output.WriteString(divider)
	output.WriteString("\n")

	return output.String()
}

func (ks *KeySheet) Print() {
	fmt.Print(ks.String())
}

func (ks *KeySheet) Save(path string) error {
	content, err := yaml.Marshal(ks)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, content, 0644)
	if err != nil {
		return err
	}

	return nil
}
