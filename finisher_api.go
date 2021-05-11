package lgorm

import (
	"database/sql"
	"gorm.io/gorm"
)

type FinisherPool struct {
	IsCall            bool
	Params            []interface{}
	HandleType        string
	HandleParamsIndex []int
}

// Create insert the value into database
func (db *Db) Create(value interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{value}
	tx.Statement.Create = append(tx.Statement.Create, FinisherPool{
		Params:            data,
		IsCall:            true,
		HandleType:        "Set",
		HandleParamsIndex: []int{0},
	})
	tx.RunFinisher()
	return
}

// CreateInBatches insert the value in batches into database
func (db *Db) CreateInBatches(value interface{}, batchSize int) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{value, batchSize}
	tx.Statement.CreateInBatches = append(tx.Statement.CreateInBatches, FinisherPool{
		Params:            data,
		IsCall:            true,
		HandleType:        "Set",
		HandleParamsIndex: []int{0},
	})
	tx.RunFinisher()
	return
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (db *Db) Save(value interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{value}
	tx.Statement.Save = append(tx.Statement.Save, FinisherPool{
		Params:            data,
		IsCall:            true,
		HandleType:        "Set",
		HandleParamsIndex: []int{0},
	})
	tx.RunFinisher()
	return
}

// First find first record that match given conditions, order by primary key
func (db *Db) First(dest interface{}, conds ...interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{dest}
	for _, d := range conds {
		data = append(data, d)
	}
	tx.Statement.First = append(tx.Statement.First, FinisherPool{
		Params:            data,
		IsCall:            true,
		HandleType:        "Get",
		HandleParamsIndex: []int{0},
	})
	tx.RunFinisher()
	return
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (db *Db) Take(dest interface{}, conds ...interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{dest}
	for _, d := range conds {
		data = append(data, d)
	}
	tx.Statement.Take = append(tx.Statement.Take, FinisherPool{
		Params:            data,
		IsCall:            true,
		HandleType:        "Get",
		HandleParamsIndex: []int{0},
	})
	tx.RunFinisher()
	return
}

// Last find last record that match given conditions, order by primary key
func (db *Db) Last(dest interface{}, conds ...interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{dest}
	for _, d := range conds {
		data = append(data, d)
	}
	tx.Statement.Last = append(tx.Statement.Last, FinisherPool{
		Params:            data,
		IsCall:            true,
		HandleType:        "Get",
		HandleParamsIndex: []int{0},
	})
	tx.RunFinisher()
	return
}

// Find find records that match given conditions
func (db *Db) Find(dest interface{}, conds ...interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{dest}
	for _, d := range conds {
		data = append(data, d)
	}
	tx.Statement.Find = append(tx.Statement.Find, FinisherPool{
		Params:            data,
		IsCall:            true,
		HandleType:        "Get",
		HandleParamsIndex: []int{0},
	})
	tx.RunFinisher()
	return
}

// FindInBatches find records in batches
func (db *Db) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *Db {
	tx := db.getInstance()
	data := []interface{}{dest, batchSize, fc}
	tx.Statement.FindInBatches = append(tx.Statement.FindInBatches, FinisherPool{
		Params:            data,
		IsCall:            true,
		HandleType:        "Get",
		HandleParamsIndex: []int{0},
	})
	tx.RunFinisher()
	return tx
}

func (db *Db) FirstOrInit(dest interface{}, conds ...interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{dest}
	for _, d := range conds {
		data = append(data, d)
	}
	tx.Statement.FirstOrInit = append(tx.Statement.FirstOrInit, FinisherPool{
		Params:            data,
		IsCall:            true,
		HandleType:        "Get",
		HandleParamsIndex: []int{0},
	})
	tx.RunFinisher()
	return
}

func (db *Db) FirstOrCreate(dest interface{}, conds ...interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{dest}
	for _, d := range conds {
		data = append(data, d)
	}
	tx.Statement.FirstOrCreate = append(tx.Statement.FirstOrCreate, FinisherPool{
		Params:            data,
		IsCall:            true,
		HandleType:        "Get",
		HandleParamsIndex: []int{0},
	})
	tx.RunFinisher()
	return db
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (db *Db) Update(column string, value interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{column, value}
	tx.Statement.Update = append(tx.Statement.Update, FinisherPool{
		Params:            data,
		IsCall:            true,
		HandleType:        "Set",
		HandleParamsIndex: []int{0},
	})
	tx.RunFinisher()
	return
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (db *Db) Updates(values interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{values}
	tx.Statement.Updates = append(tx.Statement.Updates, FinisherPool{
		Params:            data,
		IsCall:            true,
		HandleType:        "Set",
		HandleParamsIndex: []int{0},
	})
	tx.RunFinisher()
	return
}

func (db *Db) UpdateColumn(column string, value interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{column, value}
	tx.Statement.UpdateColumn = append(tx.Statement.UpdateColumn, FinisherPool{
		Params:            data,
		IsCall:            true,
		HandleType:        "Set",
		HandleParamsIndex: []int{0},
	})
	tx.RunFinisher()
	return
}

func (db *Db) UpdateColumns(values interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{values}
	tx.Statement.UpdateColumns = append(tx.Statement.UpdateColumns, FinisherPool{
		Params:            data,
		IsCall:            true,
		HandleType:        "Set",
		HandleParamsIndex: []int{0},
	})
	tx.RunFinisher()
	return
}

// Delete delete value match given conditions, if the value has primary key, then will including the primary key as condition
func (db *Db) Delete(value interface{}, conds ...interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{value}
	for _, d := range conds {
		data = append(data, d)
	}
	tx.Statement.Delete = append(tx.Statement.Delete, FinisherPool{
		Params:            data,
		IsCall:            true,
		HandleType:        "Set",
		HandleParamsIndex: []int{0},
	})
	tx.RunFinisher()
	return
}

func (db *Db) Count(count *int64) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{count}
	tx.Statement.Count = append(tx.Statement.Count, FinisherPool{
		Params: data,
		IsCall: true,
	})
	tx.RunFinisher()
	return
}

func (db *Db) Row() *sql.Row {
	return db.getInstance().Row()
}

func (db *Db) Rows() (*sql.Rows, error) {
	return db.getInstance().Rows()
}

// Scan scan value to a struct
func (db *Db) Scan(dest interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{dest}
	tx.Statement.Scan = append(tx.Statement.Scan, FinisherPool{
		Params:            data,
		IsCall:            true,
		HandleType:        "Get",
		HandleParamsIndex: []int{0},
	})
	tx.RunFinisher()
	return
}

// Pluck used to query single column from a model as a map
//     var ages []int64
//     db.Find(&users).Pluck("age", &ages)
func (db *Db) Pluck(column string, dest interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{column, dest}
	tx.Statement.Pluck = append(tx.Statement.Pluck, FinisherPool{
		Params:            data,
		IsCall:            true,
		HandleType:        "Get",
		HandleParamsIndex: []int{1},
	})
	tx.RunFinisher()
	return
}

func (db *Db) ScanRows(rows *sql.Rows, dest interface{}) error {
	return db.DB.ScanRows(rows, dest)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (db *Db) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return db.DB.Transaction(fc, opts...)
}

// Begin begins a transaction
func (db *Db) Begin(opts ...*sql.TxOptions) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{}
	for _, d := range opts {
		data = append(data, d)
	}
	tx.Statement.Begin = append(tx.Statement.Begin, FinisherPool{
		Params: data,
		IsCall: true,
	})
	tx.RunFinisher()
	return
}

// Commit commit a transaction
func (db *Db) Commit() (tx *Db) {
	tx = db.getInstance()
	tx.Statement.Commit = append(tx.Statement.Commit, FinisherPool{
		IsCall: true,
	})
	tx.RunFinisher()
	return
}

// Rollback rollback a transaction
func (db *Db) Rollback() (tx *Db) {
	tx = db.getInstance()
	tx.Statement.Rollback = append(tx.Statement.Rollback, FinisherPool{
		IsCall: true,
	})
	tx.RunFinisher()
	return
}

func (db *Db) SavePoint(name string) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{name}
	tx.Statement.SavePoint = append(tx.Statement.SavePoint, FinisherPool{
		Params: data,
		IsCall: true,
	})
	tx.RunFinisher()
	return
}

func (db *Db) RollbackTo(name string) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{name}
	tx.Statement.RollbackTo = append(tx.Statement.RollbackTo, FinisherPool{
		Params: data,
		IsCall: true,
	})
	tx.RunFinisher()
	return
}

// Exec execute raw sql
func (db *Db) Exec(sql string, values ...interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{sql}
	for _, d := range values {
		data = append(data, d)
	}
	tx.Statement.Exec = append(tx.Statement.Exec, FinisherPool{
		Params: data,
		IsCall: true,
	})
	tx.RunFinisher()
	return
}
