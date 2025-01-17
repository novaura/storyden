// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/internal/ent/account"
	"github.com/Southclaws/storyden/internal/ent/predicate"
	"github.com/Southclaws/storyden/internal/ent/question"
	"github.com/rs/xid"
)

// QuestionUpdate is the builder for updating Question entities.
type QuestionUpdate struct {
	config
	hooks     []Hook
	mutation  *QuestionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the QuestionUpdate builder.
func (qu *QuestionUpdate) Where(ps ...predicate.Question) *QuestionUpdate {
	qu.mutation.Where(ps...)
	return qu
}

// SetIndexedAt sets the "indexed_at" field.
func (qu *QuestionUpdate) SetIndexedAt(t time.Time) *QuestionUpdate {
	qu.mutation.SetIndexedAt(t)
	return qu
}

// SetNillableIndexedAt sets the "indexed_at" field if the given value is not nil.
func (qu *QuestionUpdate) SetNillableIndexedAt(t *time.Time) *QuestionUpdate {
	if t != nil {
		qu.SetIndexedAt(*t)
	}
	return qu
}

// ClearIndexedAt clears the value of the "indexed_at" field.
func (qu *QuestionUpdate) ClearIndexedAt() *QuestionUpdate {
	qu.mutation.ClearIndexedAt()
	return qu
}

// SetSlug sets the "slug" field.
func (qu *QuestionUpdate) SetSlug(s string) *QuestionUpdate {
	qu.mutation.SetSlug(s)
	return qu
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (qu *QuestionUpdate) SetNillableSlug(s *string) *QuestionUpdate {
	if s != nil {
		qu.SetSlug(*s)
	}
	return qu
}

// SetQuery sets the "query" field.
func (qu *QuestionUpdate) SetQuery(s string) *QuestionUpdate {
	qu.mutation.SetQuery(s)
	return qu
}

// SetNillableQuery sets the "query" field if the given value is not nil.
func (qu *QuestionUpdate) SetNillableQuery(s *string) *QuestionUpdate {
	if s != nil {
		qu.SetQuery(*s)
	}
	return qu
}

// SetResult sets the "result" field.
func (qu *QuestionUpdate) SetResult(s string) *QuestionUpdate {
	qu.mutation.SetResult(s)
	return qu
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (qu *QuestionUpdate) SetNillableResult(s *string) *QuestionUpdate {
	if s != nil {
		qu.SetResult(*s)
	}
	return qu
}

// SetMetadata sets the "metadata" field.
func (qu *QuestionUpdate) SetMetadata(m map[string]interface{}) *QuestionUpdate {
	qu.mutation.SetMetadata(m)
	return qu
}

// ClearMetadata clears the value of the "metadata" field.
func (qu *QuestionUpdate) ClearMetadata() *QuestionUpdate {
	qu.mutation.ClearMetadata()
	return qu
}

// SetAccountID sets the "account_id" field.
func (qu *QuestionUpdate) SetAccountID(x xid.ID) *QuestionUpdate {
	qu.mutation.SetAccountID(x)
	return qu
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (qu *QuestionUpdate) SetNillableAccountID(x *xid.ID) *QuestionUpdate {
	if x != nil {
		qu.SetAccountID(*x)
	}
	return qu
}

// ClearAccountID clears the value of the "account_id" field.
func (qu *QuestionUpdate) ClearAccountID() *QuestionUpdate {
	qu.mutation.ClearAccountID()
	return qu
}

// SetAuthorID sets the "author" edge to the Account entity by ID.
func (qu *QuestionUpdate) SetAuthorID(id xid.ID) *QuestionUpdate {
	qu.mutation.SetAuthorID(id)
	return qu
}

// SetNillableAuthorID sets the "author" edge to the Account entity by ID if the given value is not nil.
func (qu *QuestionUpdate) SetNillableAuthorID(id *xid.ID) *QuestionUpdate {
	if id != nil {
		qu = qu.SetAuthorID(*id)
	}
	return qu
}

// SetAuthor sets the "author" edge to the Account entity.
func (qu *QuestionUpdate) SetAuthor(a *Account) *QuestionUpdate {
	return qu.SetAuthorID(a.ID)
}

// Mutation returns the QuestionMutation object of the builder.
func (qu *QuestionUpdate) Mutation() *QuestionMutation {
	return qu.mutation
}

// ClearAuthor clears the "author" edge to the Account entity.
func (qu *QuestionUpdate) ClearAuthor() *QuestionUpdate {
	qu.mutation.ClearAuthor()
	return qu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (qu *QuestionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, qu.sqlSave, qu.mutation, qu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (qu *QuestionUpdate) SaveX(ctx context.Context) int {
	affected, err := qu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (qu *QuestionUpdate) Exec(ctx context.Context) error {
	_, err := qu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qu *QuestionUpdate) ExecX(ctx context.Context) {
	if err := qu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (qu *QuestionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *QuestionUpdate {
	qu.modifiers = append(qu.modifiers, modifiers...)
	return qu
}

func (qu *QuestionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(question.Table, question.Columns, sqlgraph.NewFieldSpec(question.FieldID, field.TypeString))
	if ps := qu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := qu.mutation.IndexedAt(); ok {
		_spec.SetField(question.FieldIndexedAt, field.TypeTime, value)
	}
	if qu.mutation.IndexedAtCleared() {
		_spec.ClearField(question.FieldIndexedAt, field.TypeTime)
	}
	if value, ok := qu.mutation.Slug(); ok {
		_spec.SetField(question.FieldSlug, field.TypeString, value)
	}
	if value, ok := qu.mutation.Query(); ok {
		_spec.SetField(question.FieldQuery, field.TypeString, value)
	}
	if value, ok := qu.mutation.Result(); ok {
		_spec.SetField(question.FieldResult, field.TypeString, value)
	}
	if value, ok := qu.mutation.Metadata(); ok {
		_spec.SetField(question.FieldMetadata, field.TypeJSON, value)
	}
	if qu.mutation.MetadataCleared() {
		_spec.ClearField(question.FieldMetadata, field.TypeJSON)
	}
	if qu.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   question.AuthorTable,
			Columns: []string{question.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   question.AuthorTable,
			Columns: []string{question.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(qu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, qu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{question.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	qu.mutation.done = true
	return n, nil
}

// QuestionUpdateOne is the builder for updating a single Question entity.
type QuestionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *QuestionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetIndexedAt sets the "indexed_at" field.
func (quo *QuestionUpdateOne) SetIndexedAt(t time.Time) *QuestionUpdateOne {
	quo.mutation.SetIndexedAt(t)
	return quo
}

// SetNillableIndexedAt sets the "indexed_at" field if the given value is not nil.
func (quo *QuestionUpdateOne) SetNillableIndexedAt(t *time.Time) *QuestionUpdateOne {
	if t != nil {
		quo.SetIndexedAt(*t)
	}
	return quo
}

// ClearIndexedAt clears the value of the "indexed_at" field.
func (quo *QuestionUpdateOne) ClearIndexedAt() *QuestionUpdateOne {
	quo.mutation.ClearIndexedAt()
	return quo
}

// SetSlug sets the "slug" field.
func (quo *QuestionUpdateOne) SetSlug(s string) *QuestionUpdateOne {
	quo.mutation.SetSlug(s)
	return quo
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (quo *QuestionUpdateOne) SetNillableSlug(s *string) *QuestionUpdateOne {
	if s != nil {
		quo.SetSlug(*s)
	}
	return quo
}

// SetQuery sets the "query" field.
func (quo *QuestionUpdateOne) SetQuery(s string) *QuestionUpdateOne {
	quo.mutation.SetQuery(s)
	return quo
}

// SetNillableQuery sets the "query" field if the given value is not nil.
func (quo *QuestionUpdateOne) SetNillableQuery(s *string) *QuestionUpdateOne {
	if s != nil {
		quo.SetQuery(*s)
	}
	return quo
}

// SetResult sets the "result" field.
func (quo *QuestionUpdateOne) SetResult(s string) *QuestionUpdateOne {
	quo.mutation.SetResult(s)
	return quo
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (quo *QuestionUpdateOne) SetNillableResult(s *string) *QuestionUpdateOne {
	if s != nil {
		quo.SetResult(*s)
	}
	return quo
}

// SetMetadata sets the "metadata" field.
func (quo *QuestionUpdateOne) SetMetadata(m map[string]interface{}) *QuestionUpdateOne {
	quo.mutation.SetMetadata(m)
	return quo
}

// ClearMetadata clears the value of the "metadata" field.
func (quo *QuestionUpdateOne) ClearMetadata() *QuestionUpdateOne {
	quo.mutation.ClearMetadata()
	return quo
}

// SetAccountID sets the "account_id" field.
func (quo *QuestionUpdateOne) SetAccountID(x xid.ID) *QuestionUpdateOne {
	quo.mutation.SetAccountID(x)
	return quo
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (quo *QuestionUpdateOne) SetNillableAccountID(x *xid.ID) *QuestionUpdateOne {
	if x != nil {
		quo.SetAccountID(*x)
	}
	return quo
}

// ClearAccountID clears the value of the "account_id" field.
func (quo *QuestionUpdateOne) ClearAccountID() *QuestionUpdateOne {
	quo.mutation.ClearAccountID()
	return quo
}

// SetAuthorID sets the "author" edge to the Account entity by ID.
func (quo *QuestionUpdateOne) SetAuthorID(id xid.ID) *QuestionUpdateOne {
	quo.mutation.SetAuthorID(id)
	return quo
}

// SetNillableAuthorID sets the "author" edge to the Account entity by ID if the given value is not nil.
func (quo *QuestionUpdateOne) SetNillableAuthorID(id *xid.ID) *QuestionUpdateOne {
	if id != nil {
		quo = quo.SetAuthorID(*id)
	}
	return quo
}

// SetAuthor sets the "author" edge to the Account entity.
func (quo *QuestionUpdateOne) SetAuthor(a *Account) *QuestionUpdateOne {
	return quo.SetAuthorID(a.ID)
}

// Mutation returns the QuestionMutation object of the builder.
func (quo *QuestionUpdateOne) Mutation() *QuestionMutation {
	return quo.mutation
}

// ClearAuthor clears the "author" edge to the Account entity.
func (quo *QuestionUpdateOne) ClearAuthor() *QuestionUpdateOne {
	quo.mutation.ClearAuthor()
	return quo
}

// Where appends a list predicates to the QuestionUpdate builder.
func (quo *QuestionUpdateOne) Where(ps ...predicate.Question) *QuestionUpdateOne {
	quo.mutation.Where(ps...)
	return quo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (quo *QuestionUpdateOne) Select(field string, fields ...string) *QuestionUpdateOne {
	quo.fields = append([]string{field}, fields...)
	return quo
}

// Save executes the query and returns the updated Question entity.
func (quo *QuestionUpdateOne) Save(ctx context.Context) (*Question, error) {
	return withHooks(ctx, quo.sqlSave, quo.mutation, quo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (quo *QuestionUpdateOne) SaveX(ctx context.Context) *Question {
	node, err := quo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (quo *QuestionUpdateOne) Exec(ctx context.Context) error {
	_, err := quo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (quo *QuestionUpdateOne) ExecX(ctx context.Context) {
	if err := quo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (quo *QuestionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *QuestionUpdateOne {
	quo.modifiers = append(quo.modifiers, modifiers...)
	return quo
}

func (quo *QuestionUpdateOne) sqlSave(ctx context.Context) (_node *Question, err error) {
	_spec := sqlgraph.NewUpdateSpec(question.Table, question.Columns, sqlgraph.NewFieldSpec(question.FieldID, field.TypeString))
	id, ok := quo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Question.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := quo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, question.FieldID)
		for _, f := range fields {
			if !question.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != question.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := quo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := quo.mutation.IndexedAt(); ok {
		_spec.SetField(question.FieldIndexedAt, field.TypeTime, value)
	}
	if quo.mutation.IndexedAtCleared() {
		_spec.ClearField(question.FieldIndexedAt, field.TypeTime)
	}
	if value, ok := quo.mutation.Slug(); ok {
		_spec.SetField(question.FieldSlug, field.TypeString, value)
	}
	if value, ok := quo.mutation.Query(); ok {
		_spec.SetField(question.FieldQuery, field.TypeString, value)
	}
	if value, ok := quo.mutation.Result(); ok {
		_spec.SetField(question.FieldResult, field.TypeString, value)
	}
	if value, ok := quo.mutation.Metadata(); ok {
		_spec.SetField(question.FieldMetadata, field.TypeJSON, value)
	}
	if quo.mutation.MetadataCleared() {
		_spec.ClearField(question.FieldMetadata, field.TypeJSON)
	}
	if quo.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   question.AuthorTable,
			Columns: []string{question.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   question.AuthorTable,
			Columns: []string{question.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(quo.modifiers...)
	_node = &Question{config: quo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, quo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{question.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	quo.mutation.done = true
	return _node, nil
}
