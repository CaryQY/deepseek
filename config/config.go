package config

// Config for deepseek client.
// FullURL overrides BaseURL if specified. Use for custom or proxy endpoints.
// BaseURL - base URL for deepseek API, default is "https://api.deepseek.com".
// ApiKey - deepseek API key.
// TimeoutSeconds - http client timeout used by deepseek client.
// DisableRequestValidation - disable request validation by deepseek client.
type Config struct {
	FullURL                  string
	BaseURL                  string
	ApiKey                   string
	TimeoutSeconds           int
	DisableRequestValidation bool
}
