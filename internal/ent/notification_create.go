// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/internal/ent/account"
	"github.com/Southclaws/storyden/internal/ent/notification"
	"github.com/rs/xid"
)

// NotificationCreate is the builder for creating a Notification entity.
type NotificationCreate struct {
	config
	mutation *NotificationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (nc *NotificationCreate) SetCreatedAt(t time.Time) *NotificationCreate {
	nc.mutation.SetCreatedAt(t)
	return nc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableCreatedAt(t *time.Time) *NotificationCreate {
	if t != nil {
		nc.SetCreatedAt(*t)
	}
	return nc
}

// SetDeletedAt sets the "deleted_at" field.
func (nc *NotificationCreate) SetDeletedAt(t time.Time) *NotificationCreate {
	nc.mutation.SetDeletedAt(t)
	return nc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableDeletedAt(t *time.Time) *NotificationCreate {
	if t != nil {
		nc.SetDeletedAt(*t)
	}
	return nc
}

// SetEventType sets the "event_type" field.
func (nc *NotificationCreate) SetEventType(s string) *NotificationCreate {
	nc.mutation.SetEventType(s)
	return nc
}

// SetDatagraphKind sets the "datagraph_kind" field.
func (nc *NotificationCreate) SetDatagraphKind(s string) *NotificationCreate {
	nc.mutation.SetDatagraphKind(s)
	return nc
}

// SetNillableDatagraphKind sets the "datagraph_kind" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableDatagraphKind(s *string) *NotificationCreate {
	if s != nil {
		nc.SetDatagraphKind(*s)
	}
	return nc
}

// SetDatagraphID sets the "datagraph_id" field.
func (nc *NotificationCreate) SetDatagraphID(x xid.ID) *NotificationCreate {
	nc.mutation.SetDatagraphID(x)
	return nc
}

// SetNillableDatagraphID sets the "datagraph_id" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableDatagraphID(x *xid.ID) *NotificationCreate {
	if x != nil {
		nc.SetDatagraphID(*x)
	}
	return nc
}

// SetRead sets the "read" field.
func (nc *NotificationCreate) SetRead(b bool) *NotificationCreate {
	nc.mutation.SetRead(b)
	return nc
}

// SetOwnerAccountID sets the "owner_account_id" field.
func (nc *NotificationCreate) SetOwnerAccountID(x xid.ID) *NotificationCreate {
	nc.mutation.SetOwnerAccountID(x)
	return nc
}

// SetSourceAccountID sets the "source_account_id" field.
func (nc *NotificationCreate) SetSourceAccountID(x xid.ID) *NotificationCreate {
	nc.mutation.SetSourceAccountID(x)
	return nc
}

// SetNillableSourceAccountID sets the "source_account_id" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableSourceAccountID(x *xid.ID) *NotificationCreate {
	if x != nil {
		nc.SetSourceAccountID(*x)
	}
	return nc
}

// SetID sets the "id" field.
func (nc *NotificationCreate) SetID(x xid.ID) *NotificationCreate {
	nc.mutation.SetID(x)
	return nc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableID(x *xid.ID) *NotificationCreate {
	if x != nil {
		nc.SetID(*x)
	}
	return nc
}

// SetOwnerID sets the "owner" edge to the Account entity by ID.
func (nc *NotificationCreate) SetOwnerID(id xid.ID) *NotificationCreate {
	nc.mutation.SetOwnerID(id)
	return nc
}

// SetOwner sets the "owner" edge to the Account entity.
func (nc *NotificationCreate) SetOwner(a *Account) *NotificationCreate {
	return nc.SetOwnerID(a.ID)
}

// SetSourceID sets the "source" edge to the Account entity by ID.
func (nc *NotificationCreate) SetSourceID(id xid.ID) *NotificationCreate {
	nc.mutation.SetSourceID(id)
	return nc
}

// SetNillableSourceID sets the "source" edge to the Account entity by ID if the given value is not nil.
func (nc *NotificationCreate) SetNillableSourceID(id *xid.ID) *NotificationCreate {
	if id != nil {
		nc = nc.SetSourceID(*id)
	}
	return nc
}

// SetSource sets the "source" edge to the Account entity.
func (nc *NotificationCreate) SetSource(a *Account) *NotificationCreate {
	return nc.SetSourceID(a.ID)
}

// Mutation returns the NotificationMutation object of the builder.
func (nc *NotificationCreate) Mutation() *NotificationMutation {
	return nc.mutation
}

// Save creates the Notification in the database.
func (nc *NotificationCreate) Save(ctx context.Context) (*Notification, error) {
	nc.defaults()
	return withHooks(ctx, nc.sqlSave, nc.mutation, nc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NotificationCreate) SaveX(ctx context.Context) *Notification {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NotificationCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NotificationCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nc *NotificationCreate) defaults() {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		v := notification.DefaultCreatedAt()
		nc.mutation.SetCreatedAt(v)
	}
	if _, ok := nc.mutation.ID(); !ok {
		v := notification.DefaultID()
		nc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NotificationCreate) check() error {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Notification.created_at"`)}
	}
	if _, ok := nc.mutation.EventType(); !ok {
		return &ValidationError{Name: "event_type", err: errors.New(`ent: missing required field "Notification.event_type"`)}
	}
	if _, ok := nc.mutation.Read(); !ok {
		return &ValidationError{Name: "read", err: errors.New(`ent: missing required field "Notification.read"`)}
	}
	if _, ok := nc.mutation.OwnerAccountID(); !ok {
		return &ValidationError{Name: "owner_account_id", err: errors.New(`ent: missing required field "Notification.owner_account_id"`)}
	}
	if v, ok := nc.mutation.ID(); ok {
		if err := notification.IDValidator(v.String()); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Notification.id": %w`, err)}
		}
	}
	if len(nc.mutation.OwnerIDs()) == 0 {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Notification.owner"`)}
	}
	return nil
}

func (nc *NotificationCreate) sqlSave(ctx context.Context) (*Notification, error) {
	if err := nc.check(); err != nil {
		return nil, err
	}
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*xid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	nc.mutation.id = &_node.ID
	nc.mutation.done = true
	return _node, nil
}

func (nc *NotificationCreate) createSpec() (*Notification, *sqlgraph.CreateSpec) {
	var (
		_node = &Notification{config: nc.config}
		_spec = sqlgraph.NewCreateSpec(notification.Table, sqlgraph.NewFieldSpec(notification.FieldID, field.TypeString))
	)
	_spec.OnConflict = nc.conflict
	if id, ok := nc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := nc.mutation.CreatedAt(); ok {
		_spec.SetField(notification.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := nc.mutation.DeletedAt(); ok {
		_spec.SetField(notification.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := nc.mutation.EventType(); ok {
		_spec.SetField(notification.FieldEventType, field.TypeString, value)
		_node.EventType = value
	}
	if value, ok := nc.mutation.DatagraphKind(); ok {
		_spec.SetField(notification.FieldDatagraphKind, field.TypeString, value)
		_node.DatagraphKind = &value
	}
	if value, ok := nc.mutation.DatagraphID(); ok {
		_spec.SetField(notification.FieldDatagraphID, field.TypeString, value)
		_node.DatagraphID = &value
	}
	if value, ok := nc.mutation.Read(); ok {
		_spec.SetField(notification.FieldRead, field.TypeBool, value)
		_node.Read = value
	}
	if nodes := nc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notification.OwnerTable,
			Columns: []string{notification.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OwnerAccountID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.SourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notification.SourceTable,
			Columns: []string{notification.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SourceAccountID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Notification.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NotificationUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (nc *NotificationCreate) OnConflict(opts ...sql.ConflictOption) *NotificationUpsertOne {
	nc.conflict = opts
	return &NotificationUpsertOne{
		create: nc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Notification.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (nc *NotificationCreate) OnConflictColumns(columns ...string) *NotificationUpsertOne {
	nc.conflict = append(nc.conflict, sql.ConflictColumns(columns...))
	return &NotificationUpsertOne{
		create: nc,
	}
}

type (
	// NotificationUpsertOne is the builder for "upsert"-ing
	//  one Notification node.
	NotificationUpsertOne struct {
		create *NotificationCreate
	}

	// NotificationUpsert is the "OnConflict" setter.
	NotificationUpsert struct {
		*sql.UpdateSet
	}
)

// SetDeletedAt sets the "deleted_at" field.
func (u *NotificationUpsert) SetDeletedAt(v time.Time) *NotificationUpsert {
	u.Set(notification.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *NotificationUpsert) UpdateDeletedAt() *NotificationUpsert {
	u.SetExcluded(notification.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *NotificationUpsert) ClearDeletedAt() *NotificationUpsert {
	u.SetNull(notification.FieldDeletedAt)
	return u
}

// SetEventType sets the "event_type" field.
func (u *NotificationUpsert) SetEventType(v string) *NotificationUpsert {
	u.Set(notification.FieldEventType, v)
	return u
}

// UpdateEventType sets the "event_type" field to the value that was provided on create.
func (u *NotificationUpsert) UpdateEventType() *NotificationUpsert {
	u.SetExcluded(notification.FieldEventType)
	return u
}

// SetDatagraphKind sets the "datagraph_kind" field.
func (u *NotificationUpsert) SetDatagraphKind(v string) *NotificationUpsert {
	u.Set(notification.FieldDatagraphKind, v)
	return u
}

// UpdateDatagraphKind sets the "datagraph_kind" field to the value that was provided on create.
func (u *NotificationUpsert) UpdateDatagraphKind() *NotificationUpsert {
	u.SetExcluded(notification.FieldDatagraphKind)
	return u
}

// ClearDatagraphKind clears the value of the "datagraph_kind" field.
func (u *NotificationUpsert) ClearDatagraphKind() *NotificationUpsert {
	u.SetNull(notification.FieldDatagraphKind)
	return u
}

// SetDatagraphID sets the "datagraph_id" field.
func (u *NotificationUpsert) SetDatagraphID(v xid.ID) *NotificationUpsert {
	u.Set(notification.FieldDatagraphID, v)
	return u
}

// UpdateDatagraphID sets the "datagraph_id" field to the value that was provided on create.
func (u *NotificationUpsert) UpdateDatagraphID() *NotificationUpsert {
	u.SetExcluded(notification.FieldDatagraphID)
	return u
}

// ClearDatagraphID clears the value of the "datagraph_id" field.
func (u *NotificationUpsert) ClearDatagraphID() *NotificationUpsert {
	u.SetNull(notification.FieldDatagraphID)
	return u
}

// SetRead sets the "read" field.
func (u *NotificationUpsert) SetRead(v bool) *NotificationUpsert {
	u.Set(notification.FieldRead, v)
	return u
}

// UpdateRead sets the "read" field to the value that was provided on create.
func (u *NotificationUpsert) UpdateRead() *NotificationUpsert {
	u.SetExcluded(notification.FieldRead)
	return u
}

// SetOwnerAccountID sets the "owner_account_id" field.
func (u *NotificationUpsert) SetOwnerAccountID(v xid.ID) *NotificationUpsert {
	u.Set(notification.FieldOwnerAccountID, v)
	return u
}

// UpdateOwnerAccountID sets the "owner_account_id" field to the value that was provided on create.
func (u *NotificationUpsert) UpdateOwnerAccountID() *NotificationUpsert {
	u.SetExcluded(notification.FieldOwnerAccountID)
	return u
}

// SetSourceAccountID sets the "source_account_id" field.
func (u *NotificationUpsert) SetSourceAccountID(v xid.ID) *NotificationUpsert {
	u.Set(notification.FieldSourceAccountID, v)
	return u
}

// UpdateSourceAccountID sets the "source_account_id" field to the value that was provided on create.
func (u *NotificationUpsert) UpdateSourceAccountID() *NotificationUpsert {
	u.SetExcluded(notification.FieldSourceAccountID)
	return u
}

// ClearSourceAccountID clears the value of the "source_account_id" field.
func (u *NotificationUpsert) ClearSourceAccountID() *NotificationUpsert {
	u.SetNull(notification.FieldSourceAccountID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Notification.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(notification.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *NotificationUpsertOne) UpdateNewValues() *NotificationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(notification.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(notification.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Notification.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *NotificationUpsertOne) Ignore() *NotificationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NotificationUpsertOne) DoNothing() *NotificationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NotificationCreate.OnConflict
// documentation for more info.
func (u *NotificationUpsertOne) Update(set func(*NotificationUpsert)) *NotificationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NotificationUpsert{UpdateSet: update})
	}))
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *NotificationUpsertOne) SetDeletedAt(v time.Time) *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *NotificationUpsertOne) UpdateDeletedAt() *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *NotificationUpsertOne) ClearDeletedAt() *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.ClearDeletedAt()
	})
}

// SetEventType sets the "event_type" field.
func (u *NotificationUpsertOne) SetEventType(v string) *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.SetEventType(v)
	})
}

// UpdateEventType sets the "event_type" field to the value that was provided on create.
func (u *NotificationUpsertOne) UpdateEventType() *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateEventType()
	})
}

// SetDatagraphKind sets the "datagraph_kind" field.
func (u *NotificationUpsertOne) SetDatagraphKind(v string) *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.SetDatagraphKind(v)
	})
}

// UpdateDatagraphKind sets the "datagraph_kind" field to the value that was provided on create.
func (u *NotificationUpsertOne) UpdateDatagraphKind() *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateDatagraphKind()
	})
}

// ClearDatagraphKind clears the value of the "datagraph_kind" field.
func (u *NotificationUpsertOne) ClearDatagraphKind() *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.ClearDatagraphKind()
	})
}

// SetDatagraphID sets the "datagraph_id" field.
func (u *NotificationUpsertOne) SetDatagraphID(v xid.ID) *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.SetDatagraphID(v)
	})
}

// UpdateDatagraphID sets the "datagraph_id" field to the value that was provided on create.
func (u *NotificationUpsertOne) UpdateDatagraphID() *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateDatagraphID()
	})
}

// ClearDatagraphID clears the value of the "datagraph_id" field.
func (u *NotificationUpsertOne) ClearDatagraphID() *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.ClearDatagraphID()
	})
}

// SetRead sets the "read" field.
func (u *NotificationUpsertOne) SetRead(v bool) *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.SetRead(v)
	})
}

// UpdateRead sets the "read" field to the value that was provided on create.
func (u *NotificationUpsertOne) UpdateRead() *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateRead()
	})
}

// SetOwnerAccountID sets the "owner_account_id" field.
func (u *NotificationUpsertOne) SetOwnerAccountID(v xid.ID) *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.SetOwnerAccountID(v)
	})
}

// UpdateOwnerAccountID sets the "owner_account_id" field to the value that was provided on create.
func (u *NotificationUpsertOne) UpdateOwnerAccountID() *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateOwnerAccountID()
	})
}

// SetSourceAccountID sets the "source_account_id" field.
func (u *NotificationUpsertOne) SetSourceAccountID(v xid.ID) *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.SetSourceAccountID(v)
	})
}

// UpdateSourceAccountID sets the "source_account_id" field to the value that was provided on create.
func (u *NotificationUpsertOne) UpdateSourceAccountID() *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateSourceAccountID()
	})
}

// ClearSourceAccountID clears the value of the "source_account_id" field.
func (u *NotificationUpsertOne) ClearSourceAccountID() *NotificationUpsertOne {
	return u.Update(func(s *NotificationUpsert) {
		s.ClearSourceAccountID()
	})
}

// Exec executes the query.
func (u *NotificationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NotificationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NotificationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *NotificationUpsertOne) ID(ctx context.Context) (id xid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: NotificationUpsertOne.ID is not supported by MySQL driver. Use NotificationUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *NotificationUpsertOne) IDX(ctx context.Context) xid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// NotificationCreateBulk is the builder for creating many Notification entities in bulk.
type NotificationCreateBulk struct {
	config
	err      error
	builders []*NotificationCreate
	conflict []sql.ConflictOption
}

// Save creates the Notification entities in the database.
func (ncb *NotificationCreateBulk) Save(ctx context.Context) ([]*Notification, error) {
	if ncb.err != nil {
		return nil, ncb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Notification, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NotificationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ncb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NotificationCreateBulk) SaveX(ctx context.Context) []*Notification {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NotificationCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NotificationCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Notification.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NotificationUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ncb *NotificationCreateBulk) OnConflict(opts ...sql.ConflictOption) *NotificationUpsertBulk {
	ncb.conflict = opts
	return &NotificationUpsertBulk{
		create: ncb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Notification.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ncb *NotificationCreateBulk) OnConflictColumns(columns ...string) *NotificationUpsertBulk {
	ncb.conflict = append(ncb.conflict, sql.ConflictColumns(columns...))
	return &NotificationUpsertBulk{
		create: ncb,
	}
}

// NotificationUpsertBulk is the builder for "upsert"-ing
// a bulk of Notification nodes.
type NotificationUpsertBulk struct {
	create *NotificationCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Notification.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(notification.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *NotificationUpsertBulk) UpdateNewValues() *NotificationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(notification.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(notification.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Notification.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *NotificationUpsertBulk) Ignore() *NotificationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NotificationUpsertBulk) DoNothing() *NotificationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NotificationCreateBulk.OnConflict
// documentation for more info.
func (u *NotificationUpsertBulk) Update(set func(*NotificationUpsert)) *NotificationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NotificationUpsert{UpdateSet: update})
	}))
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *NotificationUpsertBulk) SetDeletedAt(v time.Time) *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *NotificationUpsertBulk) UpdateDeletedAt() *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *NotificationUpsertBulk) ClearDeletedAt() *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.ClearDeletedAt()
	})
}

// SetEventType sets the "event_type" field.
func (u *NotificationUpsertBulk) SetEventType(v string) *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.SetEventType(v)
	})
}

// UpdateEventType sets the "event_type" field to the value that was provided on create.
func (u *NotificationUpsertBulk) UpdateEventType() *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateEventType()
	})
}

// SetDatagraphKind sets the "datagraph_kind" field.
func (u *NotificationUpsertBulk) SetDatagraphKind(v string) *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.SetDatagraphKind(v)
	})
}

// UpdateDatagraphKind sets the "datagraph_kind" field to the value that was provided on create.
func (u *NotificationUpsertBulk) UpdateDatagraphKind() *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateDatagraphKind()
	})
}

// ClearDatagraphKind clears the value of the "datagraph_kind" field.
func (u *NotificationUpsertBulk) ClearDatagraphKind() *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.ClearDatagraphKind()
	})
}

// SetDatagraphID sets the "datagraph_id" field.
func (u *NotificationUpsertBulk) SetDatagraphID(v xid.ID) *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.SetDatagraphID(v)
	})
}

// UpdateDatagraphID sets the "datagraph_id" field to the value that was provided on create.
func (u *NotificationUpsertBulk) UpdateDatagraphID() *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateDatagraphID()
	})
}

// ClearDatagraphID clears the value of the "datagraph_id" field.
func (u *NotificationUpsertBulk) ClearDatagraphID() *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.ClearDatagraphID()
	})
}

// SetRead sets the "read" field.
func (u *NotificationUpsertBulk) SetRead(v bool) *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.SetRead(v)
	})
}

// UpdateRead sets the "read" field to the value that was provided on create.
func (u *NotificationUpsertBulk) UpdateRead() *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateRead()
	})
}

// SetOwnerAccountID sets the "owner_account_id" field.
func (u *NotificationUpsertBulk) SetOwnerAccountID(v xid.ID) *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.SetOwnerAccountID(v)
	})
}

// UpdateOwnerAccountID sets the "owner_account_id" field to the value that was provided on create.
func (u *NotificationUpsertBulk) UpdateOwnerAccountID() *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateOwnerAccountID()
	})
}

// SetSourceAccountID sets the "source_account_id" field.
func (u *NotificationUpsertBulk) SetSourceAccountID(v xid.ID) *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.SetSourceAccountID(v)
	})
}

// UpdateSourceAccountID sets the "source_account_id" field to the value that was provided on create.
func (u *NotificationUpsertBulk) UpdateSourceAccountID() *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.UpdateSourceAccountID()
	})
}

// ClearSourceAccountID clears the value of the "source_account_id" field.
func (u *NotificationUpsertBulk) ClearSourceAccountID() *NotificationUpsertBulk {
	return u.Update(func(s *NotificationUpsert) {
		s.ClearSourceAccountID()
	})
}

// Exec executes the query.
func (u *NotificationUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the NotificationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NotificationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NotificationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
