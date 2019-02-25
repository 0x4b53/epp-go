package eppserver

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadWriteMessage(t *testing.T) {
	conn1, conn2 := net.Pipe()

	go func() {
		for i := 0; i < 10; i++ {
			err := WriteMessage(conn1, []byte(fmt.Sprintf("ping %d", i)))
			require.Nil(t, err)

			message, err := ReadMessage(conn1)
			require.Nil(t, err)
			assert.Equal(t, fmt.Sprintf("pong %d", i), string(message))
		}
	}()

	for i := 0; i < 10; i++ {
		message, err := ReadMessage(conn2)
		require.Nil(t, err)
		assert.Equal(t, fmt.Sprintf("ping %d", i), string(message))

		err = WriteMessage(conn2, []byte(fmt.Sprintf("pong %d", i)))
		require.Nil(t, err)
	}
}
