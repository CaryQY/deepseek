package config

// Config for deepseek client.
//
// BaseURL - base URL for deepseek API, default is "https://api.deepseek.com".
// ApiKey - deepseek API key.
// TimeoutSeconds - http client timeout used by deepseek client.
// DisableRequestValidation - disable request validation by deepseek client.
type Config struct {
	BaseURL                  string
	ApiKey                   string
	TimeoutSeconds           int
	DisableRequestValidation bool
}
