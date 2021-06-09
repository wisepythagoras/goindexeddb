package indexeddb

import (
	"time"
)

// https://gorm.io/docs/conventions.html

type IDBTableDataKey struct {
	TableID   uint64 `gorm:"index"`
	DataID    uint64 `gorm:"primaryKey"`
	Data      string `gorm:"index"`
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (IDBTableDataKey) TableName() string {
	return "idb_table_data_keys"
}

type IDBTableData struct {
	ID            uint64 `gorm:"index,autoIncrement"`
	TableID       uint64 `gorm:"primaryKey"`
	Data          string `gorm:"not null; type:text"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	TableDataKeys []IDBTableDataKey `gorm:"foreignKey:DataID;references:ID"`
}

func (IDBTableData) TableName() string {
	return "idb_table_data"
}

type IDBTableKey struct {
	TableID   uint64 `gorm:"primaryKey"`
	Name      string `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (IDBTableKey) TableName() string {
	return "idb_table_keys"
}

type IDBTable struct {
	ID        uint64 `gorm:"primaryKey; autoIncrement; not_null;"`
	DBID      uint64 `gorm:"index,column:db_id"`
	Name      string `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	TableData []IDBTableData `gorm:"foreignKey:TableID;references:ID"`
	TableKey  []IDBTableKey  `gorm:"foreignKey:TableID;references:ID"`
}

func (IDBTable) TableName() string {
	return "idb_tables"
}

type IDBDatabase struct {
	ID        uint64 `gorm:"primaryKey; autoIncrement; not_null;"`
	Name      string `gorm:"index"`
	Version   uint32
	CreatedAt time.Time
	UpdatedAt time.Time
	Tables    []IDBTable `gorm:"foreignKey:DBID;references:ID"`
}

func (IDBDatabase) TableName() string {
	return "idb_databases"
}
