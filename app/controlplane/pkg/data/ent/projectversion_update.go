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
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/predicate"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/project"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/projectversion"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/workflowrun"
	"github.com/google/uuid"
)

// ProjectVersionUpdate is the builder for updating ProjectVersion entities.
type ProjectVersionUpdate struct {
	config
	hooks     []Hook
	mutation  *ProjectVersionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ProjectVersionUpdate builder.
func (pvu *ProjectVersionUpdate) Where(ps ...predicate.ProjectVersion) *ProjectVersionUpdate {
	pvu.mutation.Where(ps...)
	return pvu
}

// SetDeletedAt sets the "deleted_at" field.
func (pvu *ProjectVersionUpdate) SetDeletedAt(t time.Time) *ProjectVersionUpdate {
	pvu.mutation.SetDeletedAt(t)
	return pvu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pvu *ProjectVersionUpdate) SetNillableDeletedAt(t *time.Time) *ProjectVersionUpdate {
	if t != nil {
		pvu.SetDeletedAt(*t)
	}
	return pvu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (pvu *ProjectVersionUpdate) ClearDeletedAt() *ProjectVersionUpdate {
	pvu.mutation.ClearDeletedAt()
	return pvu
}

// SetProjectID sets the "project_id" field.
func (pvu *ProjectVersionUpdate) SetProjectID(u uuid.UUID) *ProjectVersionUpdate {
	pvu.mutation.SetProjectID(u)
	return pvu
}

// SetNillableProjectID sets the "project_id" field if the given value is not nil.
func (pvu *ProjectVersionUpdate) SetNillableProjectID(u *uuid.UUID) *ProjectVersionUpdate {
	if u != nil {
		pvu.SetProjectID(*u)
	}
	return pvu
}

// SetPrerelease sets the "prerelease" field.
func (pvu *ProjectVersionUpdate) SetPrerelease(b bool) *ProjectVersionUpdate {
	pvu.mutation.SetPrerelease(b)
	return pvu
}

// SetNillablePrerelease sets the "prerelease" field if the given value is not nil.
func (pvu *ProjectVersionUpdate) SetNillablePrerelease(b *bool) *ProjectVersionUpdate {
	if b != nil {
		pvu.SetPrerelease(*b)
	}
	return pvu
}

// SetWorkflowRunCount sets the "workflow_run_count" field.
func (pvu *ProjectVersionUpdate) SetWorkflowRunCount(i int) *ProjectVersionUpdate {
	pvu.mutation.ResetWorkflowRunCount()
	pvu.mutation.SetWorkflowRunCount(i)
	return pvu
}

// SetNillableWorkflowRunCount sets the "workflow_run_count" field if the given value is not nil.
func (pvu *ProjectVersionUpdate) SetNillableWorkflowRunCount(i *int) *ProjectVersionUpdate {
	if i != nil {
		pvu.SetWorkflowRunCount(*i)
	}
	return pvu
}

// AddWorkflowRunCount adds i to the "workflow_run_count" field.
func (pvu *ProjectVersionUpdate) AddWorkflowRunCount(i int) *ProjectVersionUpdate {
	pvu.mutation.AddWorkflowRunCount(i)
	return pvu
}

// SetReleasedAt sets the "released_at" field.
func (pvu *ProjectVersionUpdate) SetReleasedAt(t time.Time) *ProjectVersionUpdate {
	pvu.mutation.SetReleasedAt(t)
	return pvu
}

// SetNillableReleasedAt sets the "released_at" field if the given value is not nil.
func (pvu *ProjectVersionUpdate) SetNillableReleasedAt(t *time.Time) *ProjectVersionUpdate {
	if t != nil {
		pvu.SetReleasedAt(*t)
	}
	return pvu
}

// ClearReleasedAt clears the value of the "released_at" field.
func (pvu *ProjectVersionUpdate) ClearReleasedAt() *ProjectVersionUpdate {
	pvu.mutation.ClearReleasedAt()
	return pvu
}

// SetProject sets the "project" edge to the Project entity.
func (pvu *ProjectVersionUpdate) SetProject(p *Project) *ProjectVersionUpdate {
	return pvu.SetProjectID(p.ID)
}

// AddRunIDs adds the "runs" edge to the WorkflowRun entity by IDs.
func (pvu *ProjectVersionUpdate) AddRunIDs(ids ...uuid.UUID) *ProjectVersionUpdate {
	pvu.mutation.AddRunIDs(ids...)
	return pvu
}

// AddRuns adds the "runs" edges to the WorkflowRun entity.
func (pvu *ProjectVersionUpdate) AddRuns(w ...*WorkflowRun) *ProjectVersionUpdate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return pvu.AddRunIDs(ids...)
}

// Mutation returns the ProjectVersionMutation object of the builder.
func (pvu *ProjectVersionUpdate) Mutation() *ProjectVersionMutation {
	return pvu.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (pvu *ProjectVersionUpdate) ClearProject() *ProjectVersionUpdate {
	pvu.mutation.ClearProject()
	return pvu
}

// ClearRuns clears all "runs" edges to the WorkflowRun entity.
func (pvu *ProjectVersionUpdate) ClearRuns() *ProjectVersionUpdate {
	pvu.mutation.ClearRuns()
	return pvu
}

// RemoveRunIDs removes the "runs" edge to WorkflowRun entities by IDs.
func (pvu *ProjectVersionUpdate) RemoveRunIDs(ids ...uuid.UUID) *ProjectVersionUpdate {
	pvu.mutation.RemoveRunIDs(ids...)
	return pvu
}

// RemoveRuns removes "runs" edges to WorkflowRun entities.
func (pvu *ProjectVersionUpdate) RemoveRuns(w ...*WorkflowRun) *ProjectVersionUpdate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return pvu.RemoveRunIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pvu *ProjectVersionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pvu.sqlSave, pvu.mutation, pvu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pvu *ProjectVersionUpdate) SaveX(ctx context.Context) int {
	affected, err := pvu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pvu *ProjectVersionUpdate) Exec(ctx context.Context) error {
	_, err := pvu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pvu *ProjectVersionUpdate) ExecX(ctx context.Context) {
	if err := pvu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pvu *ProjectVersionUpdate) check() error {
	if pvu.mutation.ProjectCleared() && len(pvu.mutation.ProjectIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "ProjectVersion.project"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pvu *ProjectVersionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ProjectVersionUpdate {
	pvu.modifiers = append(pvu.modifiers, modifiers...)
	return pvu
}

func (pvu *ProjectVersionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pvu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(projectversion.Table, projectversion.Columns, sqlgraph.NewFieldSpec(projectversion.FieldID, field.TypeUUID))
	if ps := pvu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pvu.mutation.DeletedAt(); ok {
		_spec.SetField(projectversion.FieldDeletedAt, field.TypeTime, value)
	}
	if pvu.mutation.DeletedAtCleared() {
		_spec.ClearField(projectversion.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := pvu.mutation.Prerelease(); ok {
		_spec.SetField(projectversion.FieldPrerelease, field.TypeBool, value)
	}
	if value, ok := pvu.mutation.WorkflowRunCount(); ok {
		_spec.SetField(projectversion.FieldWorkflowRunCount, field.TypeInt, value)
	}
	if value, ok := pvu.mutation.AddedWorkflowRunCount(); ok {
		_spec.AddField(projectversion.FieldWorkflowRunCount, field.TypeInt, value)
	}
	if value, ok := pvu.mutation.ReleasedAt(); ok {
		_spec.SetField(projectversion.FieldReleasedAt, field.TypeTime, value)
	}
	if pvu.mutation.ReleasedAtCleared() {
		_spec.ClearField(projectversion.FieldReleasedAt, field.TypeTime)
	}
	if pvu.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectversion.ProjectTable,
			Columns: []string{projectversion.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvu.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectversion.ProjectTable,
			Columns: []string{projectversion.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pvu.mutation.RunsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectversion.RunsTable,
			Columns: []string{projectversion.RunsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowrun.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvu.mutation.RemovedRunsIDs(); len(nodes) > 0 && !pvu.mutation.RunsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectversion.RunsTable,
			Columns: []string{projectversion.RunsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowrun.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvu.mutation.RunsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectversion.RunsTable,
			Columns: []string{projectversion.RunsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowrun.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(pvu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, pvu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projectversion.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pvu.mutation.done = true
	return n, nil
}

// ProjectVersionUpdateOne is the builder for updating a single ProjectVersion entity.
type ProjectVersionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ProjectVersionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetDeletedAt sets the "deleted_at" field.
func (pvuo *ProjectVersionUpdateOne) SetDeletedAt(t time.Time) *ProjectVersionUpdateOne {
	pvuo.mutation.SetDeletedAt(t)
	return pvuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pvuo *ProjectVersionUpdateOne) SetNillableDeletedAt(t *time.Time) *ProjectVersionUpdateOne {
	if t != nil {
		pvuo.SetDeletedAt(*t)
	}
	return pvuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (pvuo *ProjectVersionUpdateOne) ClearDeletedAt() *ProjectVersionUpdateOne {
	pvuo.mutation.ClearDeletedAt()
	return pvuo
}

// SetProjectID sets the "project_id" field.
func (pvuo *ProjectVersionUpdateOne) SetProjectID(u uuid.UUID) *ProjectVersionUpdateOne {
	pvuo.mutation.SetProjectID(u)
	return pvuo
}

// SetNillableProjectID sets the "project_id" field if the given value is not nil.
func (pvuo *ProjectVersionUpdateOne) SetNillableProjectID(u *uuid.UUID) *ProjectVersionUpdateOne {
	if u != nil {
		pvuo.SetProjectID(*u)
	}
	return pvuo
}

// SetPrerelease sets the "prerelease" field.
func (pvuo *ProjectVersionUpdateOne) SetPrerelease(b bool) *ProjectVersionUpdateOne {
	pvuo.mutation.SetPrerelease(b)
	return pvuo
}

// SetNillablePrerelease sets the "prerelease" field if the given value is not nil.
func (pvuo *ProjectVersionUpdateOne) SetNillablePrerelease(b *bool) *ProjectVersionUpdateOne {
	if b != nil {
		pvuo.SetPrerelease(*b)
	}
	return pvuo
}

// SetWorkflowRunCount sets the "workflow_run_count" field.
func (pvuo *ProjectVersionUpdateOne) SetWorkflowRunCount(i int) *ProjectVersionUpdateOne {
	pvuo.mutation.ResetWorkflowRunCount()
	pvuo.mutation.SetWorkflowRunCount(i)
	return pvuo
}

// SetNillableWorkflowRunCount sets the "workflow_run_count" field if the given value is not nil.
func (pvuo *ProjectVersionUpdateOne) SetNillableWorkflowRunCount(i *int) *ProjectVersionUpdateOne {
	if i != nil {
		pvuo.SetWorkflowRunCount(*i)
	}
	return pvuo
}

// AddWorkflowRunCount adds i to the "workflow_run_count" field.
func (pvuo *ProjectVersionUpdateOne) AddWorkflowRunCount(i int) *ProjectVersionUpdateOne {
	pvuo.mutation.AddWorkflowRunCount(i)
	return pvuo
}

// SetReleasedAt sets the "released_at" field.
func (pvuo *ProjectVersionUpdateOne) SetReleasedAt(t time.Time) *ProjectVersionUpdateOne {
	pvuo.mutation.SetReleasedAt(t)
	return pvuo
}

// SetNillableReleasedAt sets the "released_at" field if the given value is not nil.
func (pvuo *ProjectVersionUpdateOne) SetNillableReleasedAt(t *time.Time) *ProjectVersionUpdateOne {
	if t != nil {
		pvuo.SetReleasedAt(*t)
	}
	return pvuo
}

// ClearReleasedAt clears the value of the "released_at" field.
func (pvuo *ProjectVersionUpdateOne) ClearReleasedAt() *ProjectVersionUpdateOne {
	pvuo.mutation.ClearReleasedAt()
	return pvuo
}

// SetProject sets the "project" edge to the Project entity.
func (pvuo *ProjectVersionUpdateOne) SetProject(p *Project) *ProjectVersionUpdateOne {
	return pvuo.SetProjectID(p.ID)
}

// AddRunIDs adds the "runs" edge to the WorkflowRun entity by IDs.
func (pvuo *ProjectVersionUpdateOne) AddRunIDs(ids ...uuid.UUID) *ProjectVersionUpdateOne {
	pvuo.mutation.AddRunIDs(ids...)
	return pvuo
}

// AddRuns adds the "runs" edges to the WorkflowRun entity.
func (pvuo *ProjectVersionUpdateOne) AddRuns(w ...*WorkflowRun) *ProjectVersionUpdateOne {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return pvuo.AddRunIDs(ids...)
}

// Mutation returns the ProjectVersionMutation object of the builder.
func (pvuo *ProjectVersionUpdateOne) Mutation() *ProjectVersionMutation {
	return pvuo.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (pvuo *ProjectVersionUpdateOne) ClearProject() *ProjectVersionUpdateOne {
	pvuo.mutation.ClearProject()
	return pvuo
}

// ClearRuns clears all "runs" edges to the WorkflowRun entity.
func (pvuo *ProjectVersionUpdateOne) ClearRuns() *ProjectVersionUpdateOne {
	pvuo.mutation.ClearRuns()
	return pvuo
}

// RemoveRunIDs removes the "runs" edge to WorkflowRun entities by IDs.
func (pvuo *ProjectVersionUpdateOne) RemoveRunIDs(ids ...uuid.UUID) *ProjectVersionUpdateOne {
	pvuo.mutation.RemoveRunIDs(ids...)
	return pvuo
}

// RemoveRuns removes "runs" edges to WorkflowRun entities.
func (pvuo *ProjectVersionUpdateOne) RemoveRuns(w ...*WorkflowRun) *ProjectVersionUpdateOne {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return pvuo.RemoveRunIDs(ids...)
}

// Where appends a list predicates to the ProjectVersionUpdate builder.
func (pvuo *ProjectVersionUpdateOne) Where(ps ...predicate.ProjectVersion) *ProjectVersionUpdateOne {
	pvuo.mutation.Where(ps...)
	return pvuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pvuo *ProjectVersionUpdateOne) Select(field string, fields ...string) *ProjectVersionUpdateOne {
	pvuo.fields = append([]string{field}, fields...)
	return pvuo
}

// Save executes the query and returns the updated ProjectVersion entity.
func (pvuo *ProjectVersionUpdateOne) Save(ctx context.Context) (*ProjectVersion, error) {
	return withHooks(ctx, pvuo.sqlSave, pvuo.mutation, pvuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pvuo *ProjectVersionUpdateOne) SaveX(ctx context.Context) *ProjectVersion {
	node, err := pvuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pvuo *ProjectVersionUpdateOne) Exec(ctx context.Context) error {
	_, err := pvuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pvuo *ProjectVersionUpdateOne) ExecX(ctx context.Context) {
	if err := pvuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pvuo *ProjectVersionUpdateOne) check() error {
	if pvuo.mutation.ProjectCleared() && len(pvuo.mutation.ProjectIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "ProjectVersion.project"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pvuo *ProjectVersionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ProjectVersionUpdateOne {
	pvuo.modifiers = append(pvuo.modifiers, modifiers...)
	return pvuo
}

func (pvuo *ProjectVersionUpdateOne) sqlSave(ctx context.Context) (_node *ProjectVersion, err error) {
	if err := pvuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(projectversion.Table, projectversion.Columns, sqlgraph.NewFieldSpec(projectversion.FieldID, field.TypeUUID))
	id, ok := pvuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProjectVersion.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pvuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, projectversion.FieldID)
		for _, f := range fields {
			if !projectversion.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != projectversion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pvuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pvuo.mutation.DeletedAt(); ok {
		_spec.SetField(projectversion.FieldDeletedAt, field.TypeTime, value)
	}
	if pvuo.mutation.DeletedAtCleared() {
		_spec.ClearField(projectversion.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := pvuo.mutation.Prerelease(); ok {
		_spec.SetField(projectversion.FieldPrerelease, field.TypeBool, value)
	}
	if value, ok := pvuo.mutation.WorkflowRunCount(); ok {
		_spec.SetField(projectversion.FieldWorkflowRunCount, field.TypeInt, value)
	}
	if value, ok := pvuo.mutation.AddedWorkflowRunCount(); ok {
		_spec.AddField(projectversion.FieldWorkflowRunCount, field.TypeInt, value)
	}
	if value, ok := pvuo.mutation.ReleasedAt(); ok {
		_spec.SetField(projectversion.FieldReleasedAt, field.TypeTime, value)
	}
	if pvuo.mutation.ReleasedAtCleared() {
		_spec.ClearField(projectversion.FieldReleasedAt, field.TypeTime)
	}
	if pvuo.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectversion.ProjectTable,
			Columns: []string{projectversion.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvuo.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectversion.ProjectTable,
			Columns: []string{projectversion.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pvuo.mutation.RunsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectversion.RunsTable,
			Columns: []string{projectversion.RunsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowrun.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvuo.mutation.RemovedRunsIDs(); len(nodes) > 0 && !pvuo.mutation.RunsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectversion.RunsTable,
			Columns: []string{projectversion.RunsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowrun.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvuo.mutation.RunsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectversion.RunsTable,
			Columns: []string{projectversion.RunsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowrun.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(pvuo.modifiers...)
	_node = &ProjectVersion{config: pvuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pvuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projectversion.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pvuo.mutation.done = true
	return _node, nil
}
