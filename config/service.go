package config

// Service details
type Service struct {
	Name    string `yaml:"service.name"`
	ID      string `yaml:"service.id"`
	BaseURL string `yaml:"service.baseURL"`
	GRPC    struct {
		Host     string `yaml:"grpc.host"`
		Port     string `yaml:"grpc.port"`
		TLS      bool   `yaml:"grpc.tls"`
		Protocol string `yaml:"protocol"`
	}
	HTTP struct {
		Host           string `yaml:"http.host"`
		Port           string `yaml:"http.port"`
		RequestTimeout string `yaml:"http.requestTimeout"`
	}
	Router []Router `yaml:"service.routers"`
}

type Router struct {
	Description       string   `yaml:"description"`
	Method            string   `yaml:"method"`
	MaxAllowedAnomaly float32  `yaml:"maxAllowedAnomaly"`
	Middlewares       []string `yaml:"middlewares"`
}
