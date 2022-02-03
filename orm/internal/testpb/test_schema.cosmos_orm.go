// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package testpb

import (
	context "context"

	ormdb "github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	ormlist "github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	ormtable "github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	ormerrors "github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
)

type ExampleTableStore interface {
	Insert(ctx context.Context, exampleTable *ExampleTable) error
	Update(ctx context.Context, exampleTable *ExampleTable) error
	Save(ctx context.Context, exampleTable *ExampleTable) error
	Delete(ctx context.Context, exampleTable *ExampleTable) error
	Has(ctx context.Context, u32 uint32, i64 int64, str string) (found bool, err error)
	Get(ctx context.Context, u32 uint32, i64 int64, str string) (*ExampleTable, error)
	HasByU64Str(ctx context.Context, u64 uint64, str string) (found bool, err error)
	GetByU64Str(ctx context.Context, u64 uint64, str string) (*ExampleTable, error)
	List(ctx context.Context, prefixKey ExampleTableIndexKey, opts ...ormlist.Option) (ExampleTableIterator, error)
	ListRange(ctx context.Context, from, to ExampleTableIndexKey, opts ...ormlist.Option) (ExampleTableIterator, error)
	DeleteBy(ctx context.Context, prefixKey ExampleTableIndexKey) error
	DeleteRange(ctx context.Context, from, to ExampleTableIndexKey) error

	doNotImplement()
}

type ExampleTableIterator struct {
	ormtable.Iterator
}

func (i ExampleTableIterator) Value() (*ExampleTable, error) {
	var exampleTable ExampleTable
	err := i.UnmarshalMessage(&exampleTable)
	return &exampleTable, err
}

type ExampleTableIndexKey interface {
	id() uint32
	values() []interface{}
	exampleTableIndexKey()
}

// primary key starting index..
type ExampleTablePrimaryKey = ExampleTableU32I64StrIndexKey

type ExampleTableU32I64StrIndexKey struct {
	vs []interface{}
}

func (x ExampleTableU32I64StrIndexKey) id() uint32            { return 0 }
func (x ExampleTableU32I64StrIndexKey) values() []interface{} { return x.vs }
func (x ExampleTableU32I64StrIndexKey) exampleTableIndexKey() {}

func (this ExampleTableU32I64StrIndexKey) WithU32(u32 uint32) ExampleTableU32I64StrIndexKey {
	this.vs = []interface{}{u32}
	return this
}

func (this ExampleTableU32I64StrIndexKey) WithU32I64(u32 uint32, i64 int64) ExampleTableU32I64StrIndexKey {
	this.vs = []interface{}{u32, i64}
	return this
}

func (this ExampleTableU32I64StrIndexKey) WithU32I64Str(u32 uint32, i64 int64, str string) ExampleTableU32I64StrIndexKey {
	this.vs = []interface{}{u32, i64, str}
	return this
}

type ExampleTableU64StrIndexKey struct {
	vs []interface{}
}

func (x ExampleTableU64StrIndexKey) id() uint32            { return 1 }
func (x ExampleTableU64StrIndexKey) values() []interface{} { return x.vs }
func (x ExampleTableU64StrIndexKey) exampleTableIndexKey() {}

func (this ExampleTableU64StrIndexKey) WithU64(u64 uint64) ExampleTableU64StrIndexKey {
	this.vs = []interface{}{u64}
	return this
}

func (this ExampleTableU64StrIndexKey) WithU64Str(u64 uint64, str string) ExampleTableU64StrIndexKey {
	this.vs = []interface{}{u64, str}
	return this
}

type ExampleTableStrU32IndexKey struct {
	vs []interface{}
}

func (x ExampleTableStrU32IndexKey) id() uint32            { return 2 }
func (x ExampleTableStrU32IndexKey) values() []interface{} { return x.vs }
func (x ExampleTableStrU32IndexKey) exampleTableIndexKey() {}

func (this ExampleTableStrU32IndexKey) WithStr(str string) ExampleTableStrU32IndexKey {
	this.vs = []interface{}{str}
	return this
}

func (this ExampleTableStrU32IndexKey) WithStrU32(str string, u32 uint32) ExampleTableStrU32IndexKey {
	this.vs = []interface{}{str, u32}
	return this
}

type ExampleTableBzStrIndexKey struct {
	vs []interface{}
}

func (x ExampleTableBzStrIndexKey) id() uint32            { return 3 }
func (x ExampleTableBzStrIndexKey) values() []interface{} { return x.vs }
func (x ExampleTableBzStrIndexKey) exampleTableIndexKey() {}

func (this ExampleTableBzStrIndexKey) WithBz(bz []byte) ExampleTableBzStrIndexKey {
	this.vs = []interface{}{bz}
	return this
}

func (this ExampleTableBzStrIndexKey) WithBzStr(bz []byte, str string) ExampleTableBzStrIndexKey {
	this.vs = []interface{}{bz, str}
	return this
}

type exampleTableStore struct {
	table ormtable.Table
}

func (this exampleTableStore) Insert(ctx context.Context, exampleTable *ExampleTable) error {
	return this.table.Insert(ctx, exampleTable)
}

func (this exampleTableStore) Update(ctx context.Context, exampleTable *ExampleTable) error {
	return this.table.Update(ctx, exampleTable)
}

func (this exampleTableStore) Save(ctx context.Context, exampleTable *ExampleTable) error {
	return this.table.Save(ctx, exampleTable)
}

func (this exampleTableStore) Delete(ctx context.Context, exampleTable *ExampleTable) error {
	return this.table.Delete(ctx, exampleTable)
}

func (this exampleTableStore) Has(ctx context.Context, u32 uint32, i64 int64, str string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, u32, i64, str)
}

func (this exampleTableStore) Get(ctx context.Context, u32 uint32, i64 int64, str string) (*ExampleTable, error) {
	var exampleTable ExampleTable
	found, err := this.table.PrimaryKey().Get(ctx, &exampleTable, u32, i64, str)
	if !found {
		return nil, err
	}
	return &exampleTable, err
}

func (this exampleTableStore) HasByU64Str(ctx context.Context, u64 uint64, str string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		u64,
		str,
	)
}

func (this exampleTableStore) GetByU64Str(ctx context.Context, u64 uint64, str string) (*ExampleTable, error) {
	var exampleTable ExampleTable
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &exampleTable,
		u64,
		str,
	)
	if !found {
		return nil, err
	}
	return &exampleTable, nil
}

func (this exampleTableStore) List(ctx context.Context, prefixKey ExampleTableIndexKey, opts ...ormlist.Option) (ExampleTableIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return ExampleTableIterator{it}, err
}

func (this exampleTableStore) ListRange(ctx context.Context, from, to ExampleTableIndexKey, opts ...ormlist.Option) (ExampleTableIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return ExampleTableIterator{it}, err
}

func (this exampleTableStore) DeleteBy(ctx context.Context, prefixKey ExampleTableIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this exampleTableStore) DeleteRange(ctx context.Context, from, to ExampleTableIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this exampleTableStore) doNotImplement() {}

var _ ExampleTableStore = exampleTableStore{}

func NewExampleTableStore(db ormdb.ModuleDB) (ExampleTableStore, error) {
	table := db.GetTable(&ExampleTable{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&ExampleTable{}).ProtoReflect().Descriptor().FullName()))
	}
	return exampleTableStore{table}, nil
}

type ExampleAutoIncrementTableStore interface {
	Insert(ctx context.Context, exampleAutoIncrementTable *ExampleAutoIncrementTable) error
	InsertReturningID(ctx context.Context, exampleAutoIncrementTable *ExampleAutoIncrementTable) (uint64, error)
	Update(ctx context.Context, exampleAutoIncrementTable *ExampleAutoIncrementTable) error
	Save(ctx context.Context, exampleAutoIncrementTable *ExampleAutoIncrementTable) error
	Delete(ctx context.Context, exampleAutoIncrementTable *ExampleAutoIncrementTable) error
	Has(ctx context.Context, id uint64) (found bool, err error)
	Get(ctx context.Context, id uint64) (*ExampleAutoIncrementTable, error)
	HasByX(ctx context.Context, x string) (found bool, err error)
	GetByX(ctx context.Context, x string) (*ExampleAutoIncrementTable, error)
	List(ctx context.Context, prefixKey ExampleAutoIncrementTableIndexKey, opts ...ormlist.Option) (ExampleAutoIncrementTableIterator, error)
	ListRange(ctx context.Context, from, to ExampleAutoIncrementTableIndexKey, opts ...ormlist.Option) (ExampleAutoIncrementTableIterator, error)
	DeleteBy(ctx context.Context, prefixKey ExampleAutoIncrementTableIndexKey) error
	DeleteRange(ctx context.Context, from, to ExampleAutoIncrementTableIndexKey) error

	doNotImplement()
}

type ExampleAutoIncrementTableIterator struct {
	ormtable.Iterator
}

func (i ExampleAutoIncrementTableIterator) Value() (*ExampleAutoIncrementTable, error) {
	var exampleAutoIncrementTable ExampleAutoIncrementTable
	err := i.UnmarshalMessage(&exampleAutoIncrementTable)
	return &exampleAutoIncrementTable, err
}

type ExampleAutoIncrementTableIndexKey interface {
	id() uint32
	values() []interface{}
	exampleAutoIncrementTableIndexKey()
}

// primary key starting index..
type ExampleAutoIncrementTablePrimaryKey = ExampleAutoIncrementTableIdIndexKey

type ExampleAutoIncrementTableIdIndexKey struct {
	vs []interface{}
}

func (x ExampleAutoIncrementTableIdIndexKey) id() uint32                         { return 0 }
func (x ExampleAutoIncrementTableIdIndexKey) values() []interface{}              { return x.vs }
func (x ExampleAutoIncrementTableIdIndexKey) exampleAutoIncrementTableIndexKey() {}

func (this ExampleAutoIncrementTableIdIndexKey) WithId(id uint64) ExampleAutoIncrementTableIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type ExampleAutoIncrementTableXIndexKey struct {
	vs []interface{}
}

func (x ExampleAutoIncrementTableXIndexKey) id() uint32                         { return 1 }
func (x ExampleAutoIncrementTableXIndexKey) values() []interface{}              { return x.vs }
func (x ExampleAutoIncrementTableXIndexKey) exampleAutoIncrementTableIndexKey() {}

func (this ExampleAutoIncrementTableXIndexKey) WithX(x string) ExampleAutoIncrementTableXIndexKey {
	this.vs = []interface{}{x}
	return this
}

type exampleAutoIncrementTableStore struct {
	table ormtable.AutoIncrementTable
}

func (this exampleAutoIncrementTableStore) Insert(ctx context.Context, exampleAutoIncrementTable *ExampleAutoIncrementTable) error {
	return this.table.Insert(ctx, exampleAutoIncrementTable)
}

func (this exampleAutoIncrementTableStore) Update(ctx context.Context, exampleAutoIncrementTable *ExampleAutoIncrementTable) error {
	return this.table.Update(ctx, exampleAutoIncrementTable)
}

func (this exampleAutoIncrementTableStore) Save(ctx context.Context, exampleAutoIncrementTable *ExampleAutoIncrementTable) error {
	return this.table.Save(ctx, exampleAutoIncrementTable)
}

func (this exampleAutoIncrementTableStore) Delete(ctx context.Context, exampleAutoIncrementTable *ExampleAutoIncrementTable) error {
	return this.table.Delete(ctx, exampleAutoIncrementTable)
}

func (this exampleAutoIncrementTableStore) InsertReturningID(ctx context.Context, exampleAutoIncrementTable *ExampleAutoIncrementTable) (uint64, error) {
	return this.table.InsertReturningID(ctx, exampleAutoIncrementTable)
}

func (this exampleAutoIncrementTableStore) Has(ctx context.Context, id uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this exampleAutoIncrementTableStore) Get(ctx context.Context, id uint64) (*ExampleAutoIncrementTable, error) {
	var exampleAutoIncrementTable ExampleAutoIncrementTable
	found, err := this.table.PrimaryKey().Get(ctx, &exampleAutoIncrementTable, id)
	if !found {
		return nil, err
	}
	return &exampleAutoIncrementTable, err
}

func (this exampleAutoIncrementTableStore) HasByX(ctx context.Context, x string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		x,
	)
}

func (this exampleAutoIncrementTableStore) GetByX(ctx context.Context, x string) (*ExampleAutoIncrementTable, error) {
	var exampleAutoIncrementTable ExampleAutoIncrementTable
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &exampleAutoIncrementTable,
		x,
	)
	if !found {
		return nil, err
	}
	return &exampleAutoIncrementTable, nil
}

func (this exampleAutoIncrementTableStore) List(ctx context.Context, prefixKey ExampleAutoIncrementTableIndexKey, opts ...ormlist.Option) (ExampleAutoIncrementTableIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return ExampleAutoIncrementTableIterator{it}, err
}

func (this exampleAutoIncrementTableStore) ListRange(ctx context.Context, from, to ExampleAutoIncrementTableIndexKey, opts ...ormlist.Option) (ExampleAutoIncrementTableIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return ExampleAutoIncrementTableIterator{it}, err
}

func (this exampleAutoIncrementTableStore) DeleteBy(ctx context.Context, prefixKey ExampleAutoIncrementTableIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this exampleAutoIncrementTableStore) DeleteRange(ctx context.Context, from, to ExampleAutoIncrementTableIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this exampleAutoIncrementTableStore) doNotImplement() {}

var _ ExampleAutoIncrementTableStore = exampleAutoIncrementTableStore{}

func NewExampleAutoIncrementTableStore(db ormdb.ModuleDB) (ExampleAutoIncrementTableStore, error) {
	table := db.GetTable(&ExampleAutoIncrementTable{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&ExampleAutoIncrementTable{}).ProtoReflect().Descriptor().FullName()))
	}
	return exampleAutoIncrementTableStore{table.(ormtable.AutoIncrementTable)}, nil
}

// singleton store
type ExampleSingletonStore interface {
	Get(ctx context.Context) (*ExampleSingleton, error)
	Save(ctx context.Context, exampleSingleton *ExampleSingleton) error
}

type exampleSingletonStore struct {
	table ormtable.Table
}

var _ ExampleSingletonStore = exampleSingletonStore{}

func (x exampleSingletonStore) Get(ctx context.Context) (*ExampleSingleton, error) {
	exampleSingleton := &ExampleSingleton{}
	_, err := x.table.Get(ctx, exampleSingleton)
	return exampleSingleton, err
}

func (x exampleSingletonStore) Save(ctx context.Context, exampleSingleton *ExampleSingleton) error {
	return x.table.Save(ctx, exampleSingleton)
}

func NewExampleSingletonStore(db ormdb.ModuleDB) (ExampleSingletonStore, error) {
	table := db.GetTable(&ExampleSingleton{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&ExampleSingleton{}).ProtoReflect().Descriptor().FullName()))
	}
	return &exampleSingletonStore{table}, nil
}

type TestSchemaStore interface {
	ExampleTableStore() ExampleTableStore
	ExampleAutoIncrementTableStore() ExampleAutoIncrementTableStore
	ExampleSingletonStore() ExampleSingletonStore

	doNotImplement()
}

type testSchemaStore struct {
	exampleTable              ExampleTableStore
	exampleAutoIncrementTable ExampleAutoIncrementTableStore
	exampleSingleton          ExampleSingletonStore
}

func (x testSchemaStore) ExampleTableStore() ExampleTableStore {
	return x.exampleTable
}

func (x testSchemaStore) ExampleAutoIncrementTableStore() ExampleAutoIncrementTableStore {
	return x.exampleAutoIncrementTable
}

func (x testSchemaStore) ExampleSingletonStore() ExampleSingletonStore {
	return x.exampleSingleton
}

func (testSchemaStore) doNotImplement() {}

var _ TestSchemaStore = testSchemaStore{}

func NewTestSchemaStore(db ormdb.ModuleDB) (TestSchemaStore, error) {
	exampleTableStore, err := NewExampleTableStore(db)
	if err != nil {
		return nil, err
	}

	exampleAutoIncrementTableStore, err := NewExampleAutoIncrementTableStore(db)
	if err != nil {
		return nil, err
	}

	exampleSingletonStore, err := NewExampleSingletonStore(db)
	if err != nil {
		return nil, err
	}

	return testSchemaStore{
		exampleTableStore,
		exampleAutoIncrementTableStore,
		exampleSingletonStore,
	}, nil
}
