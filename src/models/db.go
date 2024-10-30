package models

import (
	"fmt"
	"log"
	"os"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

// InitDB initializes the database connection
func InitDB() {
	var err error
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	log.Printf("Connected to DB host %v, DB %v.", os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

    // NOTE: moved to init.sql
	//db.MustExec(schema)
}

// GetDB returns the database connection
func GetDB() *sqlx.DB {
	return db
}

// schema contains the SQL schema for the inventory and catalog tables
const schema = `
CREATE TYPE inventory_type AS ENUM ('weapon', 'consumable', 'jewelry', 'other');
CREATE TABLE player_inventory (
    id bigserial NOT NULL,
    player_id bigint NOT NULL,
    inventory_type inventory_type NOT NULL,
    item_code varchar(20) NOT NULL,
    amount bigint NOT NULL DEFAULT 0,
    created timestamp without time zone DEFAULT timezone('UTC'::text, now()),
    updated timestamp without time zone DEFAULT timezone('UTC'::text, now()),
    CONSTRAINT pk_player_inventory_id PRIMARY KEY (id)
);
CREATE UNIQUE INDEX uniq_idx_player_inventory_player_id_item_code ON player_inventory (player_id, item_code);

CREATE TYPE game_param_type AS ENUM ('int', 'str', 'bool', 'float');
CREATE TYPE game_inventory_item_rarity AS ENUM ('common', 'rare', 'exceptional', 'elite', 'unique', 'epic', 'legendary', 'divine');

CREATE TABLE game_inventory_dict (
    item_code varchar(20) NOT NULL,
    inventory_type inventory_type NOT NULL,
    item_rarity game_inventory_item_rarity NOT NULL,
    gd_description TEXT NOT NULL,
    base_param_array varchar(10),
    base_param1_name varchar(10),
    base_param1_type game_param_type,
    base_param1_value varchar(20),
    base_param2_name varchar(10),
    base_param2_type game_param_type,
    base_param2_value varchar(20),
    base_param3_name varchar(10),
    base_param3_type game_param_type,
    base_param3_value varchar(20),
    base_param4_name varchar(10),
    base_param4_type game_param_type,
    base_param4_value varchar(20),
    base_param5_name varchar(10),
    base_param5_type game_param_type,
    base_param5_value varchar(20),
    ext_params JSONB,
    i18n JSON,
        CONSTRAINT pk_item_code PRIMARY KEY (item_code)
);
CREATE TYPE game_event_type AS ENUM (
    'inventory_granted', 'inventory_consumed'
);
CREATE TABLE log_player_event (
    id bigserial NOT NULL,
    event_time timestamp without time zone DEFAULT timezone('UTC'::text, now()),
    player_id bigint NOT NULL,
    event_type game_event_type NOT NULL,
    event_value_type game_param_type NOT NULL DEFAULT 'int',
    event_value_int bigint,
    event_value_float float,
    event_value_str varchar(20),
    ext_trx_id varchar(40),
    meta_data JSONB,
    created timestamp without time zone DEFAULT timezone('UTC'::text, now()),
    updated timestamp without time zone DEFAULT timezone('UTC'::text, now())
) PARTITION BY RANGE (event_time);


CREATE TABLE player_inventory_trx (
    id bigserial NOT NULL,
    player_id bigint NOT NULL,
    ext_trx_id varchar(40) NOT NULL,
    created timestamp without time zone DEFAULT timezone('UTC'::text, now()),
    updated timestamp without time zone DEFAULT timezone('UTC'::text, now()),
    CONSTRAINT pk_player_inventory_trx_id PRIMARY KEY (id)
);
CREATE UNIQUE INDEX uniq_idx_player_inventory_trx_player_id_idempotency_key ON
    player_inventory_trx (player_id, ext_trx_id);

`

