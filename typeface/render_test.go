package typeface

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRenderRuneBadRune(t *testing.T) {
	_, err := RenderRune(nil, 'P', Medium, 0, 0)
	assert.Error(t, err)
}
