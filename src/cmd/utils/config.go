package utils

type Configuration struct {
	Database DBSettings
	Host     HostSettings
}

type DBSettings struct {
	Url      string
	Keyspace string
}

type HostSettings struct {
	Port string
}
