package utils

// Configuration hold the all configurations in a single struct
type Configuration struct {
	Database DBSettings
	Host     HostSettings
}

// DBSettings hold the Cassandra DB Settings
type DBSettings struct {
	Url      string
	Keyspace string
}

// HostSettings hold the GIN Web Server Settings
type HostSettings struct {
	Port string
}
