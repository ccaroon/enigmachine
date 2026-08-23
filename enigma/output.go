package enigma

import (
	"fmt"
	"strings"
)

type EnigmaOutput string

type FormatOptions struct {
	KeepOriginalFormatting bool
	BlockSize              int
	LineSize               int
}

func NewFormatOptions(blockSize, lineSize int, keepOriginalFmt bool) FormatOptions {
	return FormatOptions{
		KeepOriginalFormatting: keepOriginalFmt,
		BlockSize:              blockSize,
		LineSize:               lineSize,
	}
}

func (output EnigmaOutput) Format(options FormatOptions) string {
	var outputLen int = len(output)
	var lineSize int
	var formattedOutput strings.Builder

	if options.KeepOriginalFormatting {
		formattedOutput.WriteString(string(output))
	} else {
		for idx := 0; idx < outputLen; idx += options.BlockSize {
			start := idx
			end := min(idx+options.BlockSize, outputLen)
			block := string(output[start:end])
			formattedOutput.WriteString(fmt.Sprintf("%s ", block))
			lineSize += 1
			if lineSize >= options.LineSize {
				lineSize = 0
				formattedOutput.WriteString("\n")
			}
		}
	}

	return formattedOutput.String()
}

func (output EnigmaOutput) String() string {
	return string(output)
}
