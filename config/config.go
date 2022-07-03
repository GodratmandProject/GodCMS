package config

type (
	Config struct {
		Database Database `yaml:"database" json:"database"`
	}

	Database struct {
		User     string `yaml:"user" json:"user,omitempty"`
		Password string `yaml:"password" json:"password,omitempty"`
		DBName   string `yaml:"db_name" json:"db_name,omitempty"`
		Host     string `yaml:"host" json:"host,omitempty"`
		IP       string `yaml:"ip" json:"ip,omitempty"`
	}
)
