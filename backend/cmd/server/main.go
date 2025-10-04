package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Sa-Leonardo/WorkPoint/internal/config"
	"github.com/Sa-Leonardo/WorkPoint/internal/logger"
)

func main() {
	// Inicializa logger
	logger.Init()

	// Carrega config
	cfg := config.LoadConfig()

	// Cria engine do Gin
	r := gin.New()

	// Adiciona middlewares
	r.Use(gin.Recovery()) // evita que panics derrubem o servidor
	r.Use(gin.Logger())   // logger padrão do Gin (pode trocar pelo seu depois)

	// Rota de teste
	r.GET("/test", func(c *gin.Context) {
		logger.Info("Rota /test acessada", map[string]interface{}{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
		})
		c.JSON(200, gin.H{
			"message": "Funcionando com Gin",
		})
	})

	// Log de inicialização
	logger.Info("Servidor iniciado", map[string]interface{}{
		"port": cfg.Port,
	})

	// Sobe servidor
	if err := r.Run(":" + cfg.Port); err != nil {
		logger.Error("Erro ao iniciar servidor", map[string]interface{}{
			"error": err.Error(),
		})
	}
}
