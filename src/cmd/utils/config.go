package utils

type Configuration struct {
	Database DBSettings
}

type DBSettings struct {
	Url      string
	Keyspace string
}
