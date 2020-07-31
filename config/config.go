package config

// AppConfig アプリケーション本体の設定
type AppConfig struct {
	LogLevel string
}

// Config 全ての設定内容
type Config struct {
	App        AppConfig
}

func NewConfig() (*Config, error) {
	appConf := AppConfig{LogLevel: "debug"}
	return &Config{
		App: appConf,
	}, nil
}