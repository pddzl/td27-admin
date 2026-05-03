package configs

// Logger configuration (used by slog initialization)
type Logger struct {
	Level        string `mapstructure:"level" json:"level" yaml:"level"`
	Service      string `mapstructure:"service" json:"service" yaml:"service"`          // injected as "service" attr on every log line
	Format       string `mapstructure:"format" json:"format" yaml:"format"`
	Director     string `mapstructure:"director" json:"director"  yaml:"director"`
	ShowLine     bool   `mapstructure:"show-line" json:"show-line" yaml:"show-line"`
	LogInConsole bool   `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"`
}
