package uuid

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUUID(t *testing.T) {
	id := New()
	_, err := uuid.Parse(id)
	assert.Nil(t, err)
}
