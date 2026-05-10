package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"sgi-back/database"
	"sgi-back/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env nao encontrado, usando variaveis de ambiente do sistema")
	}

	database.Connect()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))

	// Auth
	r.POST("/login", handlers.Login)

	// Usuarios
	r.GET("/usuarios", handlers.ListUsuarios)
	r.GET("/usuarios/:id", handlers.GetUsuario)
	r.POST("/usuarios", handlers.CreateUsuario)
	r.PUT("/usuarios/:id", handlers.UpdateUsuario)
	r.DELETE("/usuarios/:id", handlers.DeleteUsuario)

	// Equipamentos
	r.GET("/equipamentos", handlers.ListEquipamentos)
	r.GET("/equipamentos/:id", handlers.GetEquipamento)
	r.POST("/equipamentos", handlers.CreateEquipamento)
	r.PUT("/equipamentos/:id", handlers.UpdateEquipamento)
	r.DELETE("/equipamentos/:id", handlers.DeleteEquipamento)

	// Incidentes
	r.GET("/incidentes", handlers.ListIncidentes)
	r.GET("/incidentes/:id", handlers.GetIncidente)
	r.POST("/incidentes", handlers.CreateIncidente)
	r.PUT("/incidentes/:id", handlers.UpdateIncidente)
	r.DELETE("/incidentes/:id", handlers.DeleteIncidente)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor rodando na porta %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
