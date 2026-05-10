package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"sgi-back/database"
	"sgi-back/models"
)

func ListIncidentes(c *gin.Context) {
	rows, err := database.DB.Query(
		"SELECT incidente_id, data, descricao, equipamento_id, pessoa_id, status_id FROM tbIncidente",
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var incidentes []models.Incidente
	for rows.Next() {
		var i models.Incidente
		if err := rows.Scan(&i.IncidenteID, &i.Data, &i.Descricao, &i.EquipamentoID, &i.PessoaID, &i.StatusID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		incidentes = append(incidentes, i)
	}

	c.JSON(http.StatusOK, incidentes)
}

func GetIncidente(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
		return
	}

	var i models.Incidente
	row := database.DB.QueryRow(
		"SELECT incidente_id, data, descricao, equipamento_id, pessoa_id, status_id FROM tbIncidente WHERE incidente_id = ?", id,
	)
	if err := row.Scan(&i.IncidenteID, &i.Data, &i.Descricao, &i.EquipamentoID, &i.PessoaID, &i.StatusID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Incidente nao encontrado"})
		return
	}

	c.JSON(http.StatusOK, i)
}

func CreateIncidente(c *gin.Context) {
	var i models.Incidente
	if err := c.ShouldBindJSON(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if i.Data.IsZero() {
		i.Data = time.Now()
	}

	result, err := database.DB.Exec(
		"INSERT INTO tbIncidente (data, descricao, equipamento_id, pessoa_id, status_id) VALUES (?, ?, ?, ?, ?)",
		i.Data, i.Descricao, i.EquipamentoID, i.PessoaID, i.StatusID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	i.IncidenteID = int(id)

	c.JSON(http.StatusCreated, i)
}

func UpdateIncidente(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
		return
	}

	var i models.Incidente
	if err := c.ShouldBindJSON(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = database.DB.Exec(
		"UPDATE tbIncidente SET data = ?, descricao = ?, equipamento_id = ?, pessoa_id = ?, status_id = ? WHERE incidente_id = ?",
		i.Data, i.Descricao, i.EquipamentoID, i.PessoaID, i.StatusID, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	i.IncidenteID = id
	c.JSON(http.StatusOK, i)
}

func DeleteIncidente(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
		return
	}

	_, err = database.DB.Exec("DELETE FROM tbIncidente WHERE incidente_id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Incidente deletado com sucesso"})
}
