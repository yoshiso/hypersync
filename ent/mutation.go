// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/yoshiso/hypersync/ent/fill"
	"github.com/yoshiso/hypersync/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeFill = "Fill"
)

// FillMutation represents an operation that mutates the Fill nodes in the graph.
type FillMutation struct {
	config
	op             Op
	typ            string
	id             *int
	coin           *string
	address        *string
	px             *string
	sz             *string
	side           *string
	time           *int64
	addtime        *int64
	start_position *string
	dir            *string
	hash           *string
	crossed        *bool
	fee            *string
	oid            *int64
	addoid         *int64
	tid            *int64
	addtid         *int64
	fee_token      *string
	builder_fee    *string
	clearedFields  map[string]struct{}
	done           bool
	oldValue       func(context.Context) (*Fill, error)
	predicates     []predicate.Fill
}

var _ ent.Mutation = (*FillMutation)(nil)

// fillOption allows management of the mutation configuration using functional options.
type fillOption func(*FillMutation)

// newFillMutation creates new mutation for the Fill entity.
func newFillMutation(c config, op Op, opts ...fillOption) *FillMutation {
	m := &FillMutation{
		config:        c,
		op:            op,
		typ:           TypeFill,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withFillID sets the ID field of the mutation.
func withFillID(id int) fillOption {
	return func(m *FillMutation) {
		var (
			err   error
			once  sync.Once
			value *Fill
		)
		m.oldValue = func(ctx context.Context) (*Fill, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Fill.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withFill sets the old Fill of the mutation.
func withFill(node *Fill) fillOption {
	return func(m *FillMutation) {
		m.oldValue = func(context.Context) (*Fill, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m FillMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m FillMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *FillMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *FillMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Fill.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCoin sets the "coin" field.
func (m *FillMutation) SetCoin(s string) {
	m.coin = &s
}

// Coin returns the value of the "coin" field in the mutation.
func (m *FillMutation) Coin() (r string, exists bool) {
	v := m.coin
	if v == nil {
		return
	}
	return *v, true
}

// OldCoin returns the old "coin" field's value of the Fill entity.
// If the Fill object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FillMutation) OldCoin(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCoin is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCoin requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCoin: %w", err)
	}
	return oldValue.Coin, nil
}

// ResetCoin resets all changes to the "coin" field.
func (m *FillMutation) ResetCoin() {
	m.coin = nil
}

// SetAddress sets the "address" field.
func (m *FillMutation) SetAddress(s string) {
	m.address = &s
}

// Address returns the value of the "address" field in the mutation.
func (m *FillMutation) Address() (r string, exists bool) {
	v := m.address
	if v == nil {
		return
	}
	return *v, true
}

// OldAddress returns the old "address" field's value of the Fill entity.
// If the Fill object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FillMutation) OldAddress(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAddress is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAddress requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAddress: %w", err)
	}
	return oldValue.Address, nil
}

// ResetAddress resets all changes to the "address" field.
func (m *FillMutation) ResetAddress() {
	m.address = nil
}

// SetPx sets the "px" field.
func (m *FillMutation) SetPx(s string) {
	m.px = &s
}

// Px returns the value of the "px" field in the mutation.
func (m *FillMutation) Px() (r string, exists bool) {
	v := m.px
	if v == nil {
		return
	}
	return *v, true
}

// OldPx returns the old "px" field's value of the Fill entity.
// If the Fill object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FillMutation) OldPx(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPx is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPx requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPx: %w", err)
	}
	return oldValue.Px, nil
}

// ResetPx resets all changes to the "px" field.
func (m *FillMutation) ResetPx() {
	m.px = nil
}

// SetSz sets the "sz" field.
func (m *FillMutation) SetSz(s string) {
	m.sz = &s
}

// Sz returns the value of the "sz" field in the mutation.
func (m *FillMutation) Sz() (r string, exists bool) {
	v := m.sz
	if v == nil {
		return
	}
	return *v, true
}

// OldSz returns the old "sz" field's value of the Fill entity.
// If the Fill object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FillMutation) OldSz(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSz is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSz requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSz: %w", err)
	}
	return oldValue.Sz, nil
}

// ResetSz resets all changes to the "sz" field.
func (m *FillMutation) ResetSz() {
	m.sz = nil
}

// SetSide sets the "side" field.
func (m *FillMutation) SetSide(s string) {
	m.side = &s
}

// Side returns the value of the "side" field in the mutation.
func (m *FillMutation) Side() (r string, exists bool) {
	v := m.side
	if v == nil {
		return
	}
	return *v, true
}

// OldSide returns the old "side" field's value of the Fill entity.
// If the Fill object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FillMutation) OldSide(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSide is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSide requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSide: %w", err)
	}
	return oldValue.Side, nil
}

// ResetSide resets all changes to the "side" field.
func (m *FillMutation) ResetSide() {
	m.side = nil
}

// SetTime sets the "time" field.
func (m *FillMutation) SetTime(i int64) {
	m.time = &i
	m.addtime = nil
}

// Time returns the value of the "time" field in the mutation.
func (m *FillMutation) Time() (r int64, exists bool) {
	v := m.time
	if v == nil {
		return
	}
	return *v, true
}

// OldTime returns the old "time" field's value of the Fill entity.
// If the Fill object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FillMutation) OldTime(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTime: %w", err)
	}
	return oldValue.Time, nil
}

// AddTime adds i to the "time" field.
func (m *FillMutation) AddTime(i int64) {
	if m.addtime != nil {
		*m.addtime += i
	} else {
		m.addtime = &i
	}
}

// AddedTime returns the value that was added to the "time" field in this mutation.
func (m *FillMutation) AddedTime() (r int64, exists bool) {
	v := m.addtime
	if v == nil {
		return
	}
	return *v, true
}

// ResetTime resets all changes to the "time" field.
func (m *FillMutation) ResetTime() {
	m.time = nil
	m.addtime = nil
}

// SetStartPosition sets the "start_position" field.
func (m *FillMutation) SetStartPosition(s string) {
	m.start_position = &s
}

// StartPosition returns the value of the "start_position" field in the mutation.
func (m *FillMutation) StartPosition() (r string, exists bool) {
	v := m.start_position
	if v == nil {
		return
	}
	return *v, true
}

// OldStartPosition returns the old "start_position" field's value of the Fill entity.
// If the Fill object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FillMutation) OldStartPosition(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStartPosition is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStartPosition requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStartPosition: %w", err)
	}
	return oldValue.StartPosition, nil
}

// ResetStartPosition resets all changes to the "start_position" field.
func (m *FillMutation) ResetStartPosition() {
	m.start_position = nil
}

// SetDir sets the "dir" field.
func (m *FillMutation) SetDir(s string) {
	m.dir = &s
}

// Dir returns the value of the "dir" field in the mutation.
func (m *FillMutation) Dir() (r string, exists bool) {
	v := m.dir
	if v == nil {
		return
	}
	return *v, true
}

// OldDir returns the old "dir" field's value of the Fill entity.
// If the Fill object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FillMutation) OldDir(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDir is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDir requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDir: %w", err)
	}
	return oldValue.Dir, nil
}

// ResetDir resets all changes to the "dir" field.
func (m *FillMutation) ResetDir() {
	m.dir = nil
}

// SetHash sets the "hash" field.
func (m *FillMutation) SetHash(s string) {
	m.hash = &s
}

// Hash returns the value of the "hash" field in the mutation.
func (m *FillMutation) Hash() (r string, exists bool) {
	v := m.hash
	if v == nil {
		return
	}
	return *v, true
}

// OldHash returns the old "hash" field's value of the Fill entity.
// If the Fill object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FillMutation) OldHash(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldHash is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldHash requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldHash: %w", err)
	}
	return oldValue.Hash, nil
}

// ResetHash resets all changes to the "hash" field.
func (m *FillMutation) ResetHash() {
	m.hash = nil
}

// SetCrossed sets the "crossed" field.
func (m *FillMutation) SetCrossed(b bool) {
	m.crossed = &b
}

// Crossed returns the value of the "crossed" field in the mutation.
func (m *FillMutation) Crossed() (r bool, exists bool) {
	v := m.crossed
	if v == nil {
		return
	}
	return *v, true
}

// OldCrossed returns the old "crossed" field's value of the Fill entity.
// If the Fill object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FillMutation) OldCrossed(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCrossed is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCrossed requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCrossed: %w", err)
	}
	return oldValue.Crossed, nil
}

// ResetCrossed resets all changes to the "crossed" field.
func (m *FillMutation) ResetCrossed() {
	m.crossed = nil
}

// SetFee sets the "fee" field.
func (m *FillMutation) SetFee(s string) {
	m.fee = &s
}

// Fee returns the value of the "fee" field in the mutation.
func (m *FillMutation) Fee() (r string, exists bool) {
	v := m.fee
	if v == nil {
		return
	}
	return *v, true
}

// OldFee returns the old "fee" field's value of the Fill entity.
// If the Fill object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FillMutation) OldFee(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFee is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFee requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFee: %w", err)
	}
	return oldValue.Fee, nil
}

// ResetFee resets all changes to the "fee" field.
func (m *FillMutation) ResetFee() {
	m.fee = nil
}

// SetOid sets the "oid" field.
func (m *FillMutation) SetOid(i int64) {
	m.oid = &i
	m.addoid = nil
}

// Oid returns the value of the "oid" field in the mutation.
func (m *FillMutation) Oid() (r int64, exists bool) {
	v := m.oid
	if v == nil {
		return
	}
	return *v, true
}

// OldOid returns the old "oid" field's value of the Fill entity.
// If the Fill object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FillMutation) OldOid(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldOid is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldOid requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOid: %w", err)
	}
	return oldValue.Oid, nil
}

// AddOid adds i to the "oid" field.
func (m *FillMutation) AddOid(i int64) {
	if m.addoid != nil {
		*m.addoid += i
	} else {
		m.addoid = &i
	}
}

// AddedOid returns the value that was added to the "oid" field in this mutation.
func (m *FillMutation) AddedOid() (r int64, exists bool) {
	v := m.addoid
	if v == nil {
		return
	}
	return *v, true
}

// ResetOid resets all changes to the "oid" field.
func (m *FillMutation) ResetOid() {
	m.oid = nil
	m.addoid = nil
}

// SetTid sets the "tid" field.
func (m *FillMutation) SetTid(i int64) {
	m.tid = &i
	m.addtid = nil
}

// Tid returns the value of the "tid" field in the mutation.
func (m *FillMutation) Tid() (r int64, exists bool) {
	v := m.tid
	if v == nil {
		return
	}
	return *v, true
}

// OldTid returns the old "tid" field's value of the Fill entity.
// If the Fill object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FillMutation) OldTid(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTid is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTid requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTid: %w", err)
	}
	return oldValue.Tid, nil
}

// AddTid adds i to the "tid" field.
func (m *FillMutation) AddTid(i int64) {
	if m.addtid != nil {
		*m.addtid += i
	} else {
		m.addtid = &i
	}
}

// AddedTid returns the value that was added to the "tid" field in this mutation.
func (m *FillMutation) AddedTid() (r int64, exists bool) {
	v := m.addtid
	if v == nil {
		return
	}
	return *v, true
}

// ResetTid resets all changes to the "tid" field.
func (m *FillMutation) ResetTid() {
	m.tid = nil
	m.addtid = nil
}

// SetFeeToken sets the "fee_token" field.
func (m *FillMutation) SetFeeToken(s string) {
	m.fee_token = &s
}

// FeeToken returns the value of the "fee_token" field in the mutation.
func (m *FillMutation) FeeToken() (r string, exists bool) {
	v := m.fee_token
	if v == nil {
		return
	}
	return *v, true
}

// OldFeeToken returns the old "fee_token" field's value of the Fill entity.
// If the Fill object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FillMutation) OldFeeToken(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFeeToken is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFeeToken requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFeeToken: %w", err)
	}
	return oldValue.FeeToken, nil
}

// ResetFeeToken resets all changes to the "fee_token" field.
func (m *FillMutation) ResetFeeToken() {
	m.fee_token = nil
}

// SetBuilderFee sets the "builder_fee" field.
func (m *FillMutation) SetBuilderFee(s string) {
	m.builder_fee = &s
}

// BuilderFee returns the value of the "builder_fee" field in the mutation.
func (m *FillMutation) BuilderFee() (r string, exists bool) {
	v := m.builder_fee
	if v == nil {
		return
	}
	return *v, true
}

// OldBuilderFee returns the old "builder_fee" field's value of the Fill entity.
// If the Fill object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FillMutation) OldBuilderFee(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBuilderFee is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBuilderFee requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBuilderFee: %w", err)
	}
	return oldValue.BuilderFee, nil
}

// ClearBuilderFee clears the value of the "builder_fee" field.
func (m *FillMutation) ClearBuilderFee() {
	m.builder_fee = nil
	m.clearedFields[fill.FieldBuilderFee] = struct{}{}
}

// BuilderFeeCleared returns if the "builder_fee" field was cleared in this mutation.
func (m *FillMutation) BuilderFeeCleared() bool {
	_, ok := m.clearedFields[fill.FieldBuilderFee]
	return ok
}

// ResetBuilderFee resets all changes to the "builder_fee" field.
func (m *FillMutation) ResetBuilderFee() {
	m.builder_fee = nil
	delete(m.clearedFields, fill.FieldBuilderFee)
}

// Where appends a list predicates to the FillMutation builder.
func (m *FillMutation) Where(ps ...predicate.Fill) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the FillMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *FillMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Fill, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *FillMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *FillMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Fill).
func (m *FillMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *FillMutation) Fields() []string {
	fields := make([]string, 0, 15)
	if m.coin != nil {
		fields = append(fields, fill.FieldCoin)
	}
	if m.address != nil {
		fields = append(fields, fill.FieldAddress)
	}
	if m.px != nil {
		fields = append(fields, fill.FieldPx)
	}
	if m.sz != nil {
		fields = append(fields, fill.FieldSz)
	}
	if m.side != nil {
		fields = append(fields, fill.FieldSide)
	}
	if m.time != nil {
		fields = append(fields, fill.FieldTime)
	}
	if m.start_position != nil {
		fields = append(fields, fill.FieldStartPosition)
	}
	if m.dir != nil {
		fields = append(fields, fill.FieldDir)
	}
	if m.hash != nil {
		fields = append(fields, fill.FieldHash)
	}
	if m.crossed != nil {
		fields = append(fields, fill.FieldCrossed)
	}
	if m.fee != nil {
		fields = append(fields, fill.FieldFee)
	}
	if m.oid != nil {
		fields = append(fields, fill.FieldOid)
	}
	if m.tid != nil {
		fields = append(fields, fill.FieldTid)
	}
	if m.fee_token != nil {
		fields = append(fields, fill.FieldFeeToken)
	}
	if m.builder_fee != nil {
		fields = append(fields, fill.FieldBuilderFee)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *FillMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case fill.FieldCoin:
		return m.Coin()
	case fill.FieldAddress:
		return m.Address()
	case fill.FieldPx:
		return m.Px()
	case fill.FieldSz:
		return m.Sz()
	case fill.FieldSide:
		return m.Side()
	case fill.FieldTime:
		return m.Time()
	case fill.FieldStartPosition:
		return m.StartPosition()
	case fill.FieldDir:
		return m.Dir()
	case fill.FieldHash:
		return m.Hash()
	case fill.FieldCrossed:
		return m.Crossed()
	case fill.FieldFee:
		return m.Fee()
	case fill.FieldOid:
		return m.Oid()
	case fill.FieldTid:
		return m.Tid()
	case fill.FieldFeeToken:
		return m.FeeToken()
	case fill.FieldBuilderFee:
		return m.BuilderFee()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *FillMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case fill.FieldCoin:
		return m.OldCoin(ctx)
	case fill.FieldAddress:
		return m.OldAddress(ctx)
	case fill.FieldPx:
		return m.OldPx(ctx)
	case fill.FieldSz:
		return m.OldSz(ctx)
	case fill.FieldSide:
		return m.OldSide(ctx)
	case fill.FieldTime:
		return m.OldTime(ctx)
	case fill.FieldStartPosition:
		return m.OldStartPosition(ctx)
	case fill.FieldDir:
		return m.OldDir(ctx)
	case fill.FieldHash:
		return m.OldHash(ctx)
	case fill.FieldCrossed:
		return m.OldCrossed(ctx)
	case fill.FieldFee:
		return m.OldFee(ctx)
	case fill.FieldOid:
		return m.OldOid(ctx)
	case fill.FieldTid:
		return m.OldTid(ctx)
	case fill.FieldFeeToken:
		return m.OldFeeToken(ctx)
	case fill.FieldBuilderFee:
		return m.OldBuilderFee(ctx)
	}
	return nil, fmt.Errorf("unknown Fill field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *FillMutation) SetField(name string, value ent.Value) error {
	switch name {
	case fill.FieldCoin:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCoin(v)
		return nil
	case fill.FieldAddress:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAddress(v)
		return nil
	case fill.FieldPx:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPx(v)
		return nil
	case fill.FieldSz:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSz(v)
		return nil
	case fill.FieldSide:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSide(v)
		return nil
	case fill.FieldTime:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTime(v)
		return nil
	case fill.FieldStartPosition:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStartPosition(v)
		return nil
	case fill.FieldDir:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDir(v)
		return nil
	case fill.FieldHash:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetHash(v)
		return nil
	case fill.FieldCrossed:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCrossed(v)
		return nil
	case fill.FieldFee:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFee(v)
		return nil
	case fill.FieldOid:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOid(v)
		return nil
	case fill.FieldTid:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTid(v)
		return nil
	case fill.FieldFeeToken:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFeeToken(v)
		return nil
	case fill.FieldBuilderFee:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBuilderFee(v)
		return nil
	}
	return fmt.Errorf("unknown Fill field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *FillMutation) AddedFields() []string {
	var fields []string
	if m.addtime != nil {
		fields = append(fields, fill.FieldTime)
	}
	if m.addoid != nil {
		fields = append(fields, fill.FieldOid)
	}
	if m.addtid != nil {
		fields = append(fields, fill.FieldTid)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *FillMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case fill.FieldTime:
		return m.AddedTime()
	case fill.FieldOid:
		return m.AddedOid()
	case fill.FieldTid:
		return m.AddedTid()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *FillMutation) AddField(name string, value ent.Value) error {
	switch name {
	case fill.FieldTime:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddTime(v)
		return nil
	case fill.FieldOid:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddOid(v)
		return nil
	case fill.FieldTid:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddTid(v)
		return nil
	}
	return fmt.Errorf("unknown Fill numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *FillMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(fill.FieldBuilderFee) {
		fields = append(fields, fill.FieldBuilderFee)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *FillMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *FillMutation) ClearField(name string) error {
	switch name {
	case fill.FieldBuilderFee:
		m.ClearBuilderFee()
		return nil
	}
	return fmt.Errorf("unknown Fill nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *FillMutation) ResetField(name string) error {
	switch name {
	case fill.FieldCoin:
		m.ResetCoin()
		return nil
	case fill.FieldAddress:
		m.ResetAddress()
		return nil
	case fill.FieldPx:
		m.ResetPx()
		return nil
	case fill.FieldSz:
		m.ResetSz()
		return nil
	case fill.FieldSide:
		m.ResetSide()
		return nil
	case fill.FieldTime:
		m.ResetTime()
		return nil
	case fill.FieldStartPosition:
		m.ResetStartPosition()
		return nil
	case fill.FieldDir:
		m.ResetDir()
		return nil
	case fill.FieldHash:
		m.ResetHash()
		return nil
	case fill.FieldCrossed:
		m.ResetCrossed()
		return nil
	case fill.FieldFee:
		m.ResetFee()
		return nil
	case fill.FieldOid:
		m.ResetOid()
		return nil
	case fill.FieldTid:
		m.ResetTid()
		return nil
	case fill.FieldFeeToken:
		m.ResetFeeToken()
		return nil
	case fill.FieldBuilderFee:
		m.ResetBuilderFee()
		return nil
	}
	return fmt.Errorf("unknown Fill field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *FillMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *FillMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *FillMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *FillMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *FillMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *FillMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *FillMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Fill unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *FillMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Fill edge %s", name)
}
