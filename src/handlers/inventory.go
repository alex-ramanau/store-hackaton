package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"inventory-service/models"
)

func GetInventory(c *gin.Context) {
	var request struct {
		PlayerID int `json:"player_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error_code": "400", "error_message": "Validation Error", "context": err.Error()})
		return
	}

	db := models.GetDB()
	var inventory []models.Inventory
	if err := db.Select(&inventory, "SELECT id, player_id, item_code, inventory_type, amount FROM player_inventory WHERE player_id=$1", request.PlayerID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error_code": "500", "error_message": "Internal Server Error", "context": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "data": gin.H{"player_id": request.PlayerID, "inventory": inventory}})
}

func GrantItem(c *gin.Context) {
	var request struct {
		PlayerID      int    `json:"player_id" binding:"required"`
		ItemCode      string `json:"item_code" binding:"required"`
		Amount        int    `json:"amount" binding:"required"`
		ExtTrxID      string `json:"ext_trx_id" binding:"required"`
		InventoryType string `json:"inventory_type" binding:"omitempty,oneof=consumable weapon jewelry other"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error_code": "400", "error_message": "Validation Error", "context": err.Error()})
		return
	}

	db := models.GetDB()
	_, err := db.Exec("INSERT INTO player_inventory (player_id, item_code, inventory_type, amount) VALUES ($1, $2, $3, $4)",
		request.PlayerID, request.ItemCode, request.InventoryType, request.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error_code": "500", "error_message": "Internal Server Error", "context": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "data": gin.H{}})
}

func ConsumeItem(c *gin.Context) {
	var request struct {
		PlayerID int    `json:"player_id" binding:"required"`
		ItemCode string `json:"item_code" binding:"required"`
		Amount   int    `json:"amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error_code": "400", "error_message": "Validation Error", "context": err.Error()})
		return
	}

	db := models.GetDB()
	_, err := db.Exec("UPDATE inventory SET amount = amount - $1 WHERE player_id = $2 AND item_code = $3 AND amount >= $1",
		request.Amount, request.PlayerID, request.ItemCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error_code": "500", "error_message": "Internal Server Error", "context": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "data": gin.H{}})
}

func UpdateInventory(c *gin.Context) {
	var requests []struct {
		Operation string `json:"operation" binding:"required"`
		PlayerID  int    `json:"player_id" binding:"required"`
		ItemCode  string `json:"item_code" binding:"required"`
		Amount    int    `json:"amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requests); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error_code": "400", "error_message": "Validation Error", "context": err.Error()})
		return
	}

	db := models.GetDB()
	tx := db.MustBegin()
	for _, request := range requests {
		if request.Operation == "add" {
			_, err := tx.Exec("UPDATE inventory SET amount = amount + $1 WHERE player_id = $2 AND item_code = $3",
				request.Amount, request.PlayerID, request.ItemCode)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error_code": "500", "error_message": "Internal Server Error", "context": err.Error()})
				return
			}
		} else if request.Operation == "remove" {
			_, err := tx.Exec("UPDATE inventory SET amount = amount - $1 WHERE player_id = $2 AND item_code = $3 AND amount >= $1",
				request.Amount, request.PlayerID, request.ItemCode)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error_code": "500", "error_message": "Internal Server Error", "context": err.Error()})
				return
			}
		}
	}
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"status": "OK", "data": gin.H{}})
}

