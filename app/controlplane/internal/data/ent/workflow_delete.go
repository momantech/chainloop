// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/predicate"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/workflow"
)

// WorkflowDelete is the builder for deleting a Workflow entity.
type WorkflowDelete struct {
	config
	hooks    []Hook
	mutation *WorkflowMutation
}

// Where appends a list predicates to the WorkflowDelete builder.
func (wd *WorkflowDelete) Where(ps ...predicate.Workflow) *WorkflowDelete {
	wd.mutation.Where(ps...)
	return wd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wd *WorkflowDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wd.sqlExec, wd.mutation, wd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wd *WorkflowDelete) ExecX(ctx context.Context) int {
	n, err := wd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wd *WorkflowDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(workflow.Table, sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeUUID))
	if ps := wd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wd.mutation.done = true
	return affected, err
}

// WorkflowDeleteOne is the builder for deleting a single Workflow entity.
type WorkflowDeleteOne struct {
	wd *WorkflowDelete
}

// Where appends a list predicates to the WorkflowDelete builder.
func (wdo *WorkflowDeleteOne) Where(ps ...predicate.Workflow) *WorkflowDeleteOne {
	wdo.wd.mutation.Where(ps...)
	return wdo
}

// Exec executes the deletion query.
func (wdo *WorkflowDeleteOne) Exec(ctx context.Context) error {
	n, err := wdo.wd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{workflow.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wdo *WorkflowDeleteOne) ExecX(ctx context.Context) {
	if err := wdo.Exec(ctx); err != nil {
		panic(err)
	}
}
