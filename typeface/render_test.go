package typeface

import (
	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRenderRuneBadRune(t *testing.T) {
	font, ok := AvailableFonts["medium"]
	assert.True(t, ok)
	_, err := RenderRune(nil, 'P', font, tcell.ColorYellow, 0, 0)
	assert.Error(t, err)
}
