package models

type Inventory struct {
	ID            int    `db:"id" json:"id"`
	PlayerID      int    `db:"player_id" json:"player_id"`
	ItemCode      string `db:"item_code" json:"item_code"`
	InventoryType string `db:"inventory_type" json:"inventory_type"`
	Amount        int    `db:"amount" json:"amount"`
}

type Catalog struct {
	ID             int    `db:"id" json:"id"`
	ItemCode       string `db:"item_code" json:"item_code"`
	InventoryType  string `db:"inventory_type" json:"inventory_type"`
	ItemRarity     string `db:"item_rarity" json:"item_rarity"`
	GdDescription  string `db:"gd_description" json:"gd_description"`
	BaseParamArray string `db:"base_param_array" json:"base_param_array"`
	BaseParam1Name string `db:"base_param1_name" json:"base_param1_name"`
	BaseParam1Type string `db:"base_param1_type" json:"base_param1_type"`
	BaseParam1Value string `db:"base_param1_value" json:"base_param1_value"`
	BaseParam2Name string `db:"base_param2_name" json:"base_param2_name"`
	BaseParam2Type string `db:"base_param2_type" json:"base_param2_type"`
	BaseParam2Value string `db:"base_param2_value" json:"base_param2_value"`
	BaseParam3Name string `db:"base_param3_name" json:"base_param3_name"`
	BaseParam3Type string `db:"base_param3_type" json:"base_param3_type"`
	BaseParam3Value string `db:"base_param3_value" json:"base_param3_value"`
	BaseParam4Name string `db:"base_param4_name" json:"base_param4_name"`
	BaseParam4Type string `db:"base_param4_type" json:"base_param4_type"`
	BaseParam4Value string `db:"base_param4_value" json:"base_param4_value"`
	BaseParam5Name string `db:"base_param5_name" json:"base_param5_name"`
	BaseParam5Type string `db:"base_param5_type" json:"base_param5_type"`
	BaseParam5Value string `db:"base_param5_value" json:"base_param5_value"`
	ExtParams      string `db:"ext_params" json:"ext_params"`
	I18n           string `db:"i18n" json:"i18n"`
}

