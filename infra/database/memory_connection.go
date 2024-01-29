package database

type MemoryConnection struct {
	Connection string
}

func NewMemoryConnection() *MemoryConnection {
	return &MemoryConnection{
		Connection: "connected",
	}
}

func (m *MemoryConnection) GetConnection() string {
	return m.Connection
}
