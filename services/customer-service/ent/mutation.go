// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"matsuokashuhei/grocery-store/services/customer-service/ent/customer"
	"matsuokashuhei/grocery-store/services/customer-service/ent/predicate"
	"sync"
	"time"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCustomer = "Customer"
)

// CustomerMutation represents an operation that mutates the Customer nodes in the graph.
type CustomerMutation struct {
	config
	op            Op
	typ           string
	id            *int
	create_time   *time.Time
	update_time   *time.Time
	name          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Customer, error)
	predicates    []predicate.Customer
}

var _ ent.Mutation = (*CustomerMutation)(nil)

// customerOption allows management of the mutation configuration using functional options.
type customerOption func(*CustomerMutation)

// newCustomerMutation creates new mutation for the Customer entity.
func newCustomerMutation(c config, op Op, opts ...customerOption) *CustomerMutation {
	m := &CustomerMutation{
		config:        c,
		op:            op,
		typ:           TypeCustomer,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCustomerID sets the ID field of the mutation.
func withCustomerID(id int) customerOption {
	return func(m *CustomerMutation) {
		var (
			err   error
			once  sync.Once
			value *Customer
		)
		m.oldValue = func(ctx context.Context) (*Customer, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Customer.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCustomer sets the old Customer of the mutation.
func withCustomer(node *Customer) customerOption {
	return func(m *CustomerMutation) {
		m.oldValue = func(context.Context) (*Customer, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CustomerMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CustomerMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CustomerMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CustomerMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Customer.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreateTime sets the "create_time" field.
func (m *CustomerMutation) SetCreateTime(t time.Time) {
	m.create_time = &t
}

// CreateTime returns the value of the "create_time" field in the mutation.
func (m *CustomerMutation) CreateTime() (r time.Time, exists bool) {
	v := m.create_time
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateTime returns the old "create_time" field's value of the Customer entity.
// If the Customer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CustomerMutation) OldCreateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateTime: %w", err)
	}
	return oldValue.CreateTime, nil
}

// ResetCreateTime resets all changes to the "create_time" field.
func (m *CustomerMutation) ResetCreateTime() {
	m.create_time = nil
}

// SetUpdateTime sets the "update_time" field.
func (m *CustomerMutation) SetUpdateTime(t time.Time) {
	m.update_time = &t
}

// UpdateTime returns the value of the "update_time" field in the mutation.
func (m *CustomerMutation) UpdateTime() (r time.Time, exists bool) {
	v := m.update_time
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdateTime returns the old "update_time" field's value of the Customer entity.
// If the Customer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CustomerMutation) OldUpdateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdateTime: %w", err)
	}
	return oldValue.UpdateTime, nil
}

// ResetUpdateTime resets all changes to the "update_time" field.
func (m *CustomerMutation) ResetUpdateTime() {
	m.update_time = nil
}

// SetName sets the "name" field.
func (m *CustomerMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *CustomerMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Customer entity.
// If the Customer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CustomerMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *CustomerMutation) ResetName() {
	m.name = nil
}

// Where appends a list predicates to the CustomerMutation builder.
func (m *CustomerMutation) Where(ps ...predicate.Customer) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *CustomerMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Customer).
func (m *CustomerMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CustomerMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.create_time != nil {
		fields = append(fields, customer.FieldCreateTime)
	}
	if m.update_time != nil {
		fields = append(fields, customer.FieldUpdateTime)
	}
	if m.name != nil {
		fields = append(fields, customer.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CustomerMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case customer.FieldCreateTime:
		return m.CreateTime()
	case customer.FieldUpdateTime:
		return m.UpdateTime()
	case customer.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CustomerMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case customer.FieldCreateTime:
		return m.OldCreateTime(ctx)
	case customer.FieldUpdateTime:
		return m.OldUpdateTime(ctx)
	case customer.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Customer field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CustomerMutation) SetField(name string, value ent.Value) error {
	switch name {
	case customer.FieldCreateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateTime(v)
		return nil
	case customer.FieldUpdateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdateTime(v)
		return nil
	case customer.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Customer field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CustomerMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CustomerMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CustomerMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Customer numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CustomerMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CustomerMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CustomerMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Customer nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CustomerMutation) ResetField(name string) error {
	switch name {
	case customer.FieldCreateTime:
		m.ResetCreateTime()
		return nil
	case customer.FieldUpdateTime:
		m.ResetUpdateTime()
		return nil
	case customer.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Customer field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CustomerMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CustomerMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CustomerMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CustomerMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CustomerMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CustomerMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CustomerMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Customer unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CustomerMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Customer edge %s", name)
}
