package model

// Configuration model to read a config file to init our app
type Configuration struct {
	IsHttps         bool     `json:"is_https"`
	ServerPort      uint16   `json:"server_port"`
	LogFolder       string   `json:"log_folder"`
	LogPretty       bool     `json:"log_pretty"`
	CertPem         string   `json:"cert_pem"`
	KeyPem          string   `json:"key_pem"`
	PublicFileSign  string   `json:"public_file_sign"`
	PrivateFileSign string   `json:"private_file_sign"`
	AllowedOrigins  []string `json:"allowed_origins"`
	AllowedMethods  []string `json:"allowed_methods"`
	Database        Database `json:"database"`
}

// Database model to connect to a database
type Database struct {
	Engine   string `json:"engine"`
	User     string `json:"user"`
	Password string `json:"password"`
	Server   string `json:"server"`
	Port     uint   `json:"port"`
	Name     string `json:"name"`
	SSLMode  string `json:"ssl_mode"`
}
