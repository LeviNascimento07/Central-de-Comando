package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"sgi-back/database"
	"sgi-back/models"
)

func ListEquipamentos(c *gin.Context) {
	rows, err := database.DB.Query("SELECT equipamento_id, descricao, valor_diaria FROM tbEquipamento")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var equipamentos []models.Equipamento
	for rows.Next() {
		var e models.Equipamento
		if err := rows.Scan(&e.EquipamentoID, &e.Descricao, &e.ValorDiaria); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		equipamentos = append(equipamentos, e)
	}

	c.JSON(http.StatusOK, equipamentos)
}

func GetEquipamento(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
		return
	}

	var e models.Equipamento
	row := database.DB.QueryRow(
		"SELECT equipamento_id, descricao, valor_diaria FROM tbEquipamento WHERE equipamento_id = ?", id,
	)
	if err := row.Scan(&e.EquipamentoID, &e.Descricao, &e.ValorDiaria); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Equipamento nao encontrado"})
		return
	}

	c.JSON(http.StatusOK, e)
}

func CreateEquipamento(c *gin.Context) {
	var e models.Equipamento
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := database.DB.Exec(
		"INSERT INTO tbEquipamento (descricao, valor_diaria) VALUES (?, ?)",
		e.Descricao, e.ValorDiaria,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	e.EquipamentoID = int(id)

	c.JSON(http.StatusCreated, e)
}

func UpdateEquipamento(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
		return
	}

	var e models.Equipamento
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = database.DB.Exec(
		"UPDATE tbEquipamento SET descricao = ?, valor_diaria = ? WHERE equipamento_id = ?",
		e.Descricao, e.ValorDiaria, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	e.EquipamentoID = id
	c.JSON(http.StatusOK, e)
}

func DeleteEquipamento(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
		return
	}

	_, err = database.DB.Exec("DELETE FROM tbEquipamento WHERE equipamento_id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Equipamento deletado com sucesso"})
}
