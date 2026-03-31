package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func normalizeLang(lang string) string {
	// "en-US" → "en", "id-ID" → "id"
	parts := strings.Split(lang, "-")
	return strings.ToLower(parts[0])
}

func Middleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		lang := context.GetHeader("Accept-Language")
		if lang == "" {
			lang = "id"
		}

		traceId := context.GetHeader("Request-Id")
		if traceId == "" {
			traceId = uuid.NewString()
		}

		context.Set("lang", normalizeLang(lang))
		context.Set("traceId", traceId)
		context.Next()
	}
}
