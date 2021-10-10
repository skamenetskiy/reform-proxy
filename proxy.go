package proxy

import (
	"database/sql"
	"sync/atomic"

	"gopkg.in/reform.v1"
)

// New returns new proxy.
func New(primary DB, secondaries ...DB) Proxy {
	return &proxy{primary: primary, secondaries: secondaries}
}

// Proxy interface.
type Proxy interface {
	DB

	// Primary returns primary database node.
	Primary() DB

	// Secondaries returns list of all secondary database nodes.
	Secondaries() []DB

	// Secondary returns the next secondary node.
	Secondary() DB
}

// DB interface.
type DB interface {
	DBReader
	DBWriter
}

// DBReader interface.
type DBReader interface {
	DBSelector
	DBFinder
	DBCounter
}

// DBSelector interface.
type DBSelector interface {
	SelectOneTo(str reform.Struct, tail string, args ...interface{}) error
	SelectOneFrom(view reform.View, tail string, args ...interface{}) (reform.Struct, error)
	SelectRows(view reform.View, tail string, args ...interface{}) (*sql.Rows, error)
	SelectAllFrom(view reform.View, tail string, args ...interface{}) ([]reform.Struct, error)
}

// DBFinder interface.
type DBFinder interface {
	FindOneTo(str reform.Struct, column string, arg interface{}) error
	FindOneFrom(view reform.View, column string, arg interface{}) (reform.Struct, error)
	FindRows(view reform.View, column string, arg interface{}) (*sql.Rows, error)
	FindAllFrom(view reform.View, column string, args ...interface{}) ([]reform.Struct, error)
	FindByPrimaryKeyTo(record reform.Record, pk interface{}) error
	FindByPrimaryKeyFrom(table reform.Table, pk interface{}) (reform.Record, error)
}

// DBCounter interface.
type DBCounter interface {
	Count(view reform.View, tail string, args ...interface{}) (int, error)
}

// DBWriter interface.
type DBWriter interface {
	DBInserter
	DBUpdater
	DBSaver
	DBDeleter
}

// DBInserter interface.
type DBInserter interface {
	Insert(str reform.Struct) error
	InsertColumns(str reform.Struct, columns ...string) error
	InsertMulti(structs ...reform.Struct) error
}

// DBUpdater interface.
type DBUpdater interface {
	Update(record reform.Record) error
	UpdateColumns(record reform.Record, columns ...string) error
	UpdateView(str reform.Struct, columns []string, tail string, args ...interface{}) (uint, error)
}

// DBSaver interface.
type DBSaver interface {
	Save(record reform.Record) error
}

// DBDeleter interface.
type DBDeleter interface {
	Delete(record reform.Record) error
	DeleteFrom(view reform.View, tail string, args ...interface{}) (uint, error)
}

type proxy struct {
	primary     DB
	secondaries []DB
	next        uint32
}

func (p *proxy) Primary() DB {
	return p.primary
}

func (p *proxy) Secondaries() []DB {
	return p.secondaries
}

func (p *proxy) Secondary() DB {
	if len(p.secondaries) == 0 {
		return p.primary
	}
	n := atomic.AddUint32(&p.next, 1)
	return p.secondaries[(int(n)-1)%len(p.secondaries)]
}

func (p *proxy) SelectOneTo(str reform.Struct, tail string, args ...interface{}) error {
	return p.Secondary().SelectOneTo(str, tail, args...)
}

func (p *proxy) SelectOneFrom(view reform.View, tail string, args ...interface{}) (reform.Struct, error) {
	return p.Secondary().SelectOneFrom(view, tail, args...)
}

func (p *proxy) SelectRows(view reform.View, tail string, args ...interface{}) (*sql.Rows, error) {
	return p.Secondary().SelectRows(view, tail, args...)
}

func (p *proxy) SelectAllFrom(view reform.View, tail string, args ...interface{}) ([]reform.Struct, error) {
	return p.Secondary().SelectAllFrom(view, tail, args...)
}

func (p *proxy) FindOneTo(str reform.Struct, column string, arg interface{}) error {
	return p.Secondary().FindOneTo(str, column, arg)
}

func (p *proxy) FindOneFrom(view reform.View, column string, arg interface{}) (reform.Struct, error) {
	return p.Secondary().FindOneFrom(view, column, arg)
}

func (p *proxy) FindRows(view reform.View, column string, arg interface{}) (*sql.Rows, error) {
	return p.Secondary().FindRows(view, column, arg)
}

func (p *proxy) FindAllFrom(view reform.View, column string, args ...interface{}) ([]reform.Struct, error) {
	return p.Secondary().FindAllFrom(view, column, args...)
}

func (p *proxy) FindByPrimaryKeyTo(record reform.Record, pk interface{}) error {
	return p.Secondary().FindByPrimaryKeyTo(record, pk)
}

func (p *proxy) FindByPrimaryKeyFrom(table reform.Table, pk interface{}) (reform.Record, error) {
	return p.Secondary().FindByPrimaryKeyFrom(table, pk)
}

func (p *proxy) Count(view reform.View, tail string, args ...interface{}) (int, error) {
	return p.Secondary().Count(view, tail, args)
}

func (p *proxy) Insert(str reform.Struct) error {
	return p.Primary().Insert(str)
}

func (p *proxy) InsertColumns(str reform.Struct, columns ...string) error {
	return p.Primary().InsertColumns(str, columns...)
}

func (p *proxy) InsertMulti(structs ...reform.Struct) error {
	return p.Primary().InsertMulti(structs...)
}

func (p *proxy) Update(record reform.Record) error {
	return p.Primary().Update(record)
}

func (p *proxy) UpdateColumns(record reform.Record, columns ...string) error {
	return p.Primary().UpdateColumns(record, columns...)
}

func (p *proxy) UpdateView(str reform.Struct, columns []string, tail string, args ...interface{}) (uint, error) {
	return p.Primary().UpdateView(str, columns, tail, args)
}

func (p *proxy) Save(record reform.Record) error {
	return p.Primary().Save(record)
}

func (p *proxy) Delete(record reform.Record) error {
	return p.Primary().Delete(record)
}

func (p *proxy) DeleteFrom(view reform.View, tail string, args ...interface{}) (uint, error) {
	return p.Primary().DeleteFrom(view, tail, args...)
}
