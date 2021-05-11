package lgorm

import (
	"gorm.io/gorm/clause"
)

type ChainAblePool struct {
	IsCall bool
	Params []interface{}
}

// Model specify the model you would like to run db operations
//    // update all users's name to `hello`
//    db.Model(&User{}).Update("name", "hello")
//    // if user's primary key is non-blank, will use it as condition, then will only update the user's name to `hello`
//    db.Model(&user).Update("name", "hello")
func (db *Db) Model(value interface{}) (tx *Db) {
	tx = db.getInstance()
	tx.Statement.Model = ChainAblePool{
		Params: []interface{}{value},
		IsCall: true,
	}
	return
}

// Clauses Add clauses
func (db *Db) Clauses(conds ...clause.Expression) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{}
	for _, d := range conds {
		data = append(data, d)
	}
	tx.Statement.Clauses = append(tx.Statement.Clauses, ChainAblePool{
		Params: data,
		IsCall: true,
	})
	return
}

// Table specify the table you would like to run db operations
func (db *Db) Table(name string, args ...interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{name}
	for _, d := range args {
		data = append(data, d)
	}
	tx.Statement.Table = ChainAblePool{
		Params: data,
		IsCall: true,
	}
	return
}

// Distinct specify distinct fields that you want querying
func (db *Db) Distinct(args ...interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{}
	for _, d := range args {
		data = append(data, d)
	}
	tx.Statement.Distinct = append(tx.Statement.Distinct, ChainAblePool{
		Params: data,
		IsCall: true,
	})
	return
}

// Select specify fields that you want when querying, creating, updating
func (db *Db) Select(query interface{}, args ...interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{query}
	for _, d := range args {
		data = append(data, d)
	}
	tx.Statement.Select = append(tx.Statement.Select, ChainAblePool{
		Params: data,
		IsCall: true,
	})
	return
}

// Omit specify fields that you want to ignore when creating, updating and querying
func (db *Db) Omit(columns ...string) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{}
	for _, d := range columns {
		data = append(data, d)
	}
	tx.Statement.Omit = append(tx.Statement.Omit, ChainAblePool{
		Params: data,
		IsCall: true,
	})
	return
}

// Where add conditions
func (db *Db) Where(query interface{}, args ...interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{query}
	for _, d := range args {
		data = append(data, d)
	}
	tx.Statement.Where = append(tx.Statement.Where, ChainAblePool{
		Params: data,
		IsCall: true,
	})
	return
}

// Not add NOT conditions
func (db *Db) Not(query interface{}, args ...interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{query}
	for _, d := range args {
		data = append(data, d)
	}
	tx.Statement.Not = append(tx.Statement.Not, ChainAblePool{
		Params: data,
		IsCall: true,
	})
	return
}

// Or add OR conditions
func (db *Db) Or(query interface{}, args ...interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{query}
	for _, d := range args {
		data = append(data, d)
	}
	tx.Statement.Or = append(tx.Statement.Or, ChainAblePool{
		Params: data,
		IsCall: true,
	})
	return
}

// Joins specify Joins conditions
//     db.Joins("Account").Find(&user)
//     db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Find(&user)
func (db *Db) Joins(query string, args ...interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{query}
	for _, d := range args {
		data = append(data, d)
	}
	tx.Statement.Joins = append(tx.Statement.Joins, ChainAblePool{
		Params: data,
		IsCall: true,
	})
	return
}

// Group specify the group method on the find
func (db *Db) Group(name string) (tx *Db) {
	tx = db.getInstance()
	tx.Statement.Group = append(tx.Statement.Group, ChainAblePool{
		Params: []interface{}{name},
		IsCall: true,
	})
	return
}

// Having specify HAVING conditions for GROUP BY
func (db *Db) Having(query interface{}, args ...interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{query}
	for _, d := range args {
		data = append(data, d)
	}
	tx.Statement.Having = append(tx.Statement.Having, ChainAblePool{
		Params: data,
		IsCall: true,
	})
	return
}

// Order specify order when retrieve records from database
//     db.Order("name DESC")
//     db.Order(clause.OrderByColumn{Column: clause.Column{Name: "name"}, Desc: true})
func (db *Db) Order(value interface{}) (tx *Db) {
	tx = db.getInstance()
	tx.Statement.Order = append(tx.Statement.Order, ChainAblePool{
		Params: []interface{}{value},
		IsCall: true,
	})
	return
}

// Limit specify the number of records to be retrieved
func (db *Db) Limit(limit int) (tx *Db) {
	tx = db.getInstance()
	tx.Statement.Limit = ChainAblePool{
		Params: []interface{}{limit},
		IsCall: true,
	}
	return
}

// Offset specify the number of records to skip before starting to return the records
func (db *Db) Offset(offset int) (tx *Db) {
	tx = db.getInstance()
	tx.Statement.Offset = ChainAblePool{
		Params: []interface{}{offset},
		IsCall: true,
	}
	return
}

// Scopes pass current database connection to arguments `func(Db) Db`, which could be used to add conditions dynamically
//     func AmountGreaterThan1000(db *gorm.Db) *gorm.Db {
//         return db.Where("amount > ?", 1000)
//     }
//
//     func OrderStatus(status []string) func (db *gorm.Db) *gorm.Db {
//         return func (db *gorm.Db) *gorm.Db {
//             return db.Scopes(AmountGreaterThan1000).Where("status in (?)", status)
//         }
//     }
//
//     db.Scopes(AmountGreaterThan1000, OrderStatus([]string{"paid", "shipped"})).Find(&orders)
func (db *Db) Scopes(funcs ...func(*Db) *Db) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{}
	for _, d := range funcs {
		data = append(data, d)
	}
	tx.Statement.Scopes = append(tx.Statement.Scopes, ChainAblePool{
		Params: data,
		IsCall: true,
	})
	return
}

// Preload preload associations with given conditions
//    db.Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
func (db *Db) Preload(query string, args ...interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{query}
	for _, d := range args {
		data = append(data, d)
	}
	tx.Statement.Preload = append(tx.Statement.Preload, ChainAblePool{
		Params: data,
		IsCall: true,
	})
	return
}

func (db *Db) Attrs(attrs ...interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{}
	for _, d := range attrs {
		data = append(data, d)
	}
	tx.Statement.Attrs = append(tx.Statement.Attrs, ChainAblePool{
		Params: data,
		IsCall: true,
	})
	return
}

func (db *Db) Assign(attrs ...interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{}
	for _, d := range attrs {
		data = append(data, d)
	}
	tx.Statement.Assign = append(tx.Statement.Assign, ChainAblePool{
		Params: data,
		IsCall: true,
	})
	return
}

func (db *Db) Unscoped() (tx *Db) {
	tx = db.getInstance()
	tx.Statement.Unscoped = ChainAblePool{
		IsCall: true,
	}
	return
}

func (db *Db) Raw(sql string, values ...interface{}) (tx *Db) {
	tx = db.getInstance()
	data := []interface{}{sql}
	for _, d := range values {
		data = append(data, d)
	}
	tx.Statement.Raw = append(tx.Statement.Raw, ChainAblePool{
		Params: data,
		IsCall: true,
	})
	return
}
