package loader

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoader(t *testing.T) {	
	r, err := usage(context.TODO(), "test", "test")	
	assert.Equal(t, errNotExisted, err)
	assert.Equal(t, nil,r )
}