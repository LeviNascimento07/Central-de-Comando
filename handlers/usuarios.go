package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"sgi-back/database"
	"sgi-back/models"
)

func ListUsuarios(c *gin.Context) {
	rows, err := database.DB.Query("SELECT usuario_id, nome, login, atualizado_em, atualizado_por FROM tbUsuarios")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var usuarios []models.Usuario
	for rows.Next() {
		var u models.Usuario
		if err := rows.Scan(&u.UsuarioID, &u.Nome, &u.Login, &u.AtualizadoEm, &u.AtualizadoPor); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		usuarios = append(usuarios, u)
	}

	c.JSON(http.StatusOK, usuarios)
}

func GetUsuario(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
		return
	}

	var u models.Usuario
	row := database.DB.QueryRow(
		"SELECT usuario_id, nome, login, atualizado_em, atualizado_por FROM tbUsuarios WHERE usuario_id = ?", id,
	)
	if err := row.Scan(&u.UsuarioID, &u.Nome, &u.Login, &u.AtualizadoEm, &u.AtualizadoPor); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario nao encontrado"})
		return
	}

	c.JSON(http.StatusOK, u)
}

func CreateUsuario(c *gin.Context) {
	var u models.Usuario
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u.AtualizadoEm = time.Now()

	result, err := database.DB.Exec(
		"INSERT INTO tbUsuarios (nome, login, senha, atualizado_em, atualizado_por) VALUES (?, ?, SHA2(?, 256), ?, ?)",
		u.Nome, u.Login, u.Senha, u.AtualizadoEm, u.AtualizadoPor,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	u.UsuarioID = int(id)
	u.Senha = ""

	c.JSON(http.StatusCreated, u)
}

func UpdateUsuario(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
		return
	}

	var u models.Usuario
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u.AtualizadoEm = time.Now()

	var execErr error
	if u.Senha != "" {
		_, execErr = database.DB.Exec(
			"UPDATE tbUsuarios SET nome = ?, login = ?, senha = SHA2(?, 256), atualizado_em = ?, atualizado_por = ? WHERE usuario_id = ?",
			u.Nome, u.Login, u.Senha, u.AtualizadoEm, u.AtualizadoPor, id,
		)
	} else {
		_, execErr = database.DB.Exec(
			"UPDATE tbUsuarios SET nome = ?, login = ?, atualizado_em = ?, atualizado_por = ? WHERE usuario_id = ?",
			u.Nome, u.Login, u.AtualizadoEm, u.AtualizadoPor, id,
		)
	}
	if execErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": execErr.Error()})
		return
	}

	u.UsuarioID = id
	u.Senha = ""
	c.JSON(http.StatusOK, u)
}

func DeleteUsuario(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
		return
	}

	_, err = database.DB.Exec("DELETE FROM tbUsuarios WHERE usuario_id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario deletado com sucesso"})
}
