package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"sgi-back/database"
)

type loginRequest struct {
	Login string `json:"login" binding:"required"`
	Senha string `json:"senha" binding:"required"`
}

type usuarioPublico struct {
	UsuarioID int    `json:"usuario_id"`
	Nome      string `json:"nome"`
	Login     string `json:"login"`
}

func Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Login e senha são obrigatórios"})
		return
	}

	var u usuarioPublico
	row := database.DB.QueryRow(
		"SELECT usuario_id, nome, login FROM tbUsuarios WHERE login = ? AND senha = SHA2(?, 256)",
		req.Login, req.Senha,
	)
	if err := row.Scan(&u.UsuarioID, &u.Nome, &u.Login); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha inválidos"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "usuario": u})
}
