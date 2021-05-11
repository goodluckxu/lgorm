package lgorm

import (
	"gorm.io/gorm"
)

func Open(dialector gorm.Dialector, opts ...gorm.Option) (*Db, error) {
	db := new(Db)
	db.ConnPool = ConnPool{
		Dialector: dialector,
		Opts:      opts,
	}
	db.conn()
	return db, db.ConnPool.Err
}

func (db *Db) conn() {
	dialector := db.ConnPool.Dialector
	opts := db.ConnPool.Opts
	gormDb, err := gorm.Open(dialector, opts...)
	db.DB = gormDb
	db.ConnPool.Err = err
}
