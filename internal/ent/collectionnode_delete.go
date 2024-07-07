// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Southclaws/storyden/internal/ent/collectionnode"
	"github.com/Southclaws/storyden/internal/ent/predicate"
)

// CollectionNodeDelete is the builder for deleting a CollectionNode entity.
type CollectionNodeDelete struct {
	config
	hooks    []Hook
	mutation *CollectionNodeMutation
}

// Where appends a list predicates to the CollectionNodeDelete builder.
func (cnd *CollectionNodeDelete) Where(ps ...predicate.CollectionNode) *CollectionNodeDelete {
	cnd.mutation.Where(ps...)
	return cnd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cnd *CollectionNodeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cnd.sqlExec, cnd.mutation, cnd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cnd *CollectionNodeDelete) ExecX(ctx context.Context) int {
	n, err := cnd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cnd *CollectionNodeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(collectionnode.Table, nil)
	if ps := cnd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cnd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cnd.mutation.done = true
	return affected, err
}

// CollectionNodeDeleteOne is the builder for deleting a single CollectionNode entity.
type CollectionNodeDeleteOne struct {
	cnd *CollectionNodeDelete
}

// Where appends a list predicates to the CollectionNodeDelete builder.
func (cndo *CollectionNodeDeleteOne) Where(ps ...predicate.CollectionNode) *CollectionNodeDeleteOne {
	cndo.cnd.mutation.Where(ps...)
	return cndo
}

// Exec executes the deletion query.
func (cndo *CollectionNodeDeleteOne) Exec(ctx context.Context) error {
	n, err := cndo.cnd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{collectionnode.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cndo *CollectionNodeDeleteOne) ExecX(ctx context.Context) {
	if err := cndo.Exec(ctx); err != nil {
		panic(err)
	}
}