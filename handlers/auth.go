package handlers

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"sgi-back/database"
)

// HashSenha retorna o hex SHA-256 da senha. Usado também em CreateUsuario/UpdateUsuario.
func HashSenha(senha string) string {
	h := sha256.Sum256([]byte(senha))
	return fmt.Sprintf("%x", h)
}

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

	senhaHash := HashSenha(req.Senha)

	var u usuarioPublico
	row := database.DB.QueryRow(
		"SELECT usuario_id, nome, login FROM tbUsuarios WHERE login = ? AND senha = ?",
		req.Login, senhaHash,
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
