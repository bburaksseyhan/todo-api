package utils

// Hold the all configurations in a single struct
type Configuration struct {
	Database DBSettings
	Host     HostSettings
}

// Cassandra DB Settings
type DBSettings struct {
	Url      string
	Keyspace string
}

// GIN Web Server Settings
type HostSettings struct {
	Port string
}
