package typeface

import (
	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRenderRuneBadRune(t *testing.T) {
	_, err := RenderRune(nil, 'P', Medium, tcell.ColorYellow, 0, 0)
	assert.Error(t, err)
}
