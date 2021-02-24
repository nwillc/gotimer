package setup

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommandLines(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		parseError bool
		version    bool
		font       string
		colorName  string
	}{
		{
			name:       "UnknownOption",
			args:       []string{"-foo"},
			parseError: true,
		},
		{
			name:      "Version",
			args:      []string{"-version"},
			version:   true,
			font:      defaultFont,
			colorName: defaultColor,
		},
		{
			name:      "NoVersion",
			args:      []string{},
			font:      defaultFont,
			colorName: defaultColor,
		},
		{
			name:      "MediumFont",
			args:      []string{"-font", "medium"},
			font:      "medium",
			colorName: defaultColor,
		},
		{
			name:       "BadFont",
			args:       []string{"-font", "foo"},
			parseError: true,
		},
		{
			name:       "BadColor",
			args:       []string{"-color", "polkadot"},
			parseError: true,
		},
		{
			name:      "AnotherColor",
			args:      []string{"-color", "green"},
			font:      defaultFont,
			colorName: "green",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cli := newFlagSetWithValues(test.name)
			err := cli.Parse(test.args)
			assert.Equal(t, test.parseError, err != nil)
			if !test.parseError {
				assert.Equal(t, test.version, *cli.Values.Version)
				assert.Equal(t, test.font, *cli.Values.FontName)
				assert.Equal(t, test.colorName, *cli.Values.ColorName)
			}
		})
	}
}
