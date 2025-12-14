package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestIDMiddleware adiciona um ID único a cada requisição
// para rastreamento e correlação de logs
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Tenta obter request ID do header (útil para correlação de sistemas)
		requestID := c.GetHeader("X-Request-ID")

		// Se não existir, gera um novo UUID
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// Armazena no contexto para uso posterior
		c.Set("request_id", requestID)

		// Adiciona ao response header para o cliente
		c.Header("X-Request-ID", requestID)

		c.Next()
	}
}
