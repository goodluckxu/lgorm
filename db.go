package lgorm

import (
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
	Statement Statement
	ConnPool  ConnPool
}

type ConnPool struct {
	Dialector gorm.Dialector
	Opts      []gorm.Option
	Err       error
}

func (db *Db) getInstance() (tx *Db) {
	tx = new(Db)
	tx.Statement = db.Statement
	tx.DB = db.DB
	return
}
