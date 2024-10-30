package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"inventory-service/models"
)

func GetCatalog(c *gin.Context) {
	db := models.GetDB()
	var catalog []models.Catalog
	if err := db.Select(&catalog, "SELECT * FROM catalog"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error_code": "500", "error_message": "Internal Server Error", "context": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "data": catalog})
}

func CreateCatalogEntry(c *gin.Context) {
	var entry models.Catalog

	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error_code": "400", "error_message": "Validation Error", "context": err.Error()})
		return
	}

	db := models.GetDB()
	_, err := db.NamedExec(`INSERT INTO catalog (item_code, inventory_type, item_rarity, gd_description, base_param_array, base_param1_name, base_param1_type, base_param1_value, base_param2_name, base_param2_type, base_param2_value, base_param3_name, base_param3_type, base_param3_value, base_param4_name, base_param4_type, base_param4_value, base_param5_name, base_param5_type, base_param5_value, ext_params, i18n) 
	VALUES (:item_code, :inventory_type, :item_rarity, :gd_description, :base_param_array, :base_param1_name, :base_param1_type, :base_param1_value, :base_param2_name, :base_param2_type, :base_param2_value, :base_param3_name, :base_param3_type, :base_param3_value, :base_param4_name, :base_param4_type, :base_param4_value, :base_param5_name, :base_param5_type, :base_param5_value, :ext_params, :i18n)`, &entry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error_code": "500", "error_message": "Internal Server Error", "context": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "data": gin.H{}})
}

func UpdateCatalogEntry(c *gin.Context) {
	var entry models.Catalog

	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error_code": "400", "error_message": "Validation Error", "context": err.Error()})
		return
	}

	db := models.GetDB()
	_, err := db.NamedExec(`UPDATE catalog SET inventory_type=:inventory_type, item_rarity=:item_rarity, gd_description=:gd_description, base_param_array=:base_param_array, base_param1_name=:base_param1_name, base_param1_type=:base_param1_type, base_param1_value=:base_param1_value, base_param2_name=:base_param2_name, base_param2_type=:base_param2_type, base_param2_value=:base_param2_value, base_param3_name=:base_param3_name, base_param3_type=:base_param3_type, base_param3_value=:base_param3_value, base_param4_name=:base_param4_name, base_param4_type=:base_param4_type, base_param4_value=:base_param4_value, base_param5_name=:base_param5_name, base_param5_type=:base_param5_type, base_param5_value=:base_param5_value, ext_params=:ext_params, i18n=:i18n WHERE item_code=:item_code`, &entry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error_code": "500", "error_message": "Internal Server Error", "context": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "data": gin.H{}})
}

func DeleteCatalogEntry(c *gin.Context) {
	var request struct {
		ItemCode string `json:"item_code" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error_code": "400", "error_message": "Validation Error", "context": err.Error()})
		return
	}

	db := models.GetDB()
	_, err := db.Exec("DELETE FROM catalog WHERE item_code=$1", request.ItemCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error_code": "500", "error_message": "Internal Server Error", "context": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "data": gin.H{}})
}

