package config

type Config struct {
	IP        string         `json:"ip"`
	Port      int            `json:"port"`
	Env       string         `json:"env"`
	DB        DatabaseConfig `json:"db"`
	Limiter   LimiterConfig  `json:"limiter"`
	Debug     bool           `json:"debug"`
	SMTP      SMTP           `json:"smtp"`
	CORS      CORS           `json:"cors"`
	LogConfig LogConfig      `json:"log_config"`
}

type DatabaseConfig struct {
	DSN          string `json:"-"` // Data Source Name (DSN)
	MaxOpenConns int    `json:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns"`
	MaxIdleTime  string `json:"max_idle_time"`
}

type LimiterConfig struct {
	RPS     float64 `json:"rps"`
	Burst   int     `json:"burst"`
	Enabled bool    `json:"enabled"`
}

type SMTP struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"-"`
	Password string `json:"-"`
	Sender   string `json:"sender"`
}

type CORS struct {
	TrustedOrigins map[string]struct{} `json:"trusted_origins"`
}

type LogConfig struct {
	Level string `json:"level"`
}
