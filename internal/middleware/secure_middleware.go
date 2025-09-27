package middleware

import (
	"os"

	"github.com/gin-contrib/secure"
)

func DefaultSecureConfig() secure.Config {
	environment, _ := os.LookupEnv("ENVIRONMENT")

	return secure.Config{
		SSLRedirect:           true,
		IsDevelopment:         environment == "development",
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
		IENoOpen:              true,
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
	}
}
