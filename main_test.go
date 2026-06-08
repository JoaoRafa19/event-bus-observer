package eventbusobserver

import (
	"net"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTCPConn(t *testing.T) {
	listener, err := net.Listen("tcp", ":8080")
	require.NoError(t, err)

	t.Log("server started on 8080")

	for {
		conn, err := listener.Accept()
		require.NoError(t, err)

		go func(conn net.Conn) {
			buffer := make([]byte, 1024)
			for {
				n, err := conn.Read(buffer)
				require.NoError(t, err)

				t.Logf("%s", buffer[:n])
			}
		}(conn)
	}
}

func Test2(t *testing.T) {

}
