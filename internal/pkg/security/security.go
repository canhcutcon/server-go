package security

import (
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
)

func Security() gin.HandlerFunc {
	return secure.New(secure.Config{
		FrameDeny:             true,                 // Deny clickjacking
		ContentTypeNosniff:    true,                 // Prevent MIME type sniffing
		BrowserXssFilter:      true,                 // XSS Protection
		ContentSecurityPolicy: "default-src 'self'", // Content Security Policy
		STSSeconds:            31536000,             // Strict Transport Security
		STSIncludeSubdomains:  true,                 // Include subdomains
		IsDevelopment:         false,
		IENoOpen:              true,                                            // X-Download-Options for IE8+
		ReferrerPolicy:        "same-origin",                                   // Referrer Policy
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"}, // SSL Proxy Headers
	})
}
