package simple

import "fmt"

type Connection struct {
	*File
}

func (c *Connection) Close() {
	fmt.Println("Close Connection", c.File.Name)
}

func NewConnection(f *File) (*Connection, func()) {
	conn := &Connection{
		File: f,
	}

	return conn, func() {
		conn.Close()
	}
}
