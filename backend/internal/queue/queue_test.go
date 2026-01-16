package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	qu := NewMemoryQueue()
	payload := Payload{
		PostId: "1",
	}

	err := qu.Enqueue(payload)
	if err != nil {
		assert.Error(t, err)
	}

	assert.NoError(t, err)
	assert.Equal(t, payload.PostId, "1")
}
