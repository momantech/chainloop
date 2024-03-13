// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/organization"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/workflow"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/workflowcontract"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/workflowcontractversion"
	"github.com/google/uuid"
)

// WorkflowContractCreate is the builder for creating a WorkflowContract entity.
type WorkflowContractCreate struct {
	config
	mutation *WorkflowContractMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (wcc *WorkflowContractCreate) SetName(s string) *WorkflowContractCreate {
	wcc.mutation.SetName(s)
	return wcc
}

// SetCreatedAt sets the "created_at" field.
func (wcc *WorkflowContractCreate) SetCreatedAt(t time.Time) *WorkflowContractCreate {
	wcc.mutation.SetCreatedAt(t)
	return wcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wcc *WorkflowContractCreate) SetNillableCreatedAt(t *time.Time) *WorkflowContractCreate {
	if t != nil {
		wcc.SetCreatedAt(*t)
	}
	return wcc
}

// SetDeletedAt sets the "deleted_at" field.
func (wcc *WorkflowContractCreate) SetDeletedAt(t time.Time) *WorkflowContractCreate {
	wcc.mutation.SetDeletedAt(t)
	return wcc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (wcc *WorkflowContractCreate) SetNillableDeletedAt(t *time.Time) *WorkflowContractCreate {
	if t != nil {
		wcc.SetDeletedAt(*t)
	}
	return wcc
}

// SetDescription sets the "description" field.
func (wcc *WorkflowContractCreate) SetDescription(s string) *WorkflowContractCreate {
	wcc.mutation.SetDescription(s)
	return wcc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wcc *WorkflowContractCreate) SetNillableDescription(s *string) *WorkflowContractCreate {
	if s != nil {
		wcc.SetDescription(*s)
	}
	return wcc
}

// SetID sets the "id" field.
func (wcc *WorkflowContractCreate) SetID(u uuid.UUID) *WorkflowContractCreate {
	wcc.mutation.SetID(u)
	return wcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wcc *WorkflowContractCreate) SetNillableID(u *uuid.UUID) *WorkflowContractCreate {
	if u != nil {
		wcc.SetID(*u)
	}
	return wcc
}

// AddVersionIDs adds the "versions" edge to the WorkflowContractVersion entity by IDs.
func (wcc *WorkflowContractCreate) AddVersionIDs(ids ...uuid.UUID) *WorkflowContractCreate {
	wcc.mutation.AddVersionIDs(ids...)
	return wcc
}

// AddVersions adds the "versions" edges to the WorkflowContractVersion entity.
func (wcc *WorkflowContractCreate) AddVersions(w ...*WorkflowContractVersion) *WorkflowContractCreate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wcc.AddVersionIDs(ids...)
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (wcc *WorkflowContractCreate) SetOrganizationID(id uuid.UUID) *WorkflowContractCreate {
	wcc.mutation.SetOrganizationID(id)
	return wcc
}

// SetNillableOrganizationID sets the "organization" edge to the Organization entity by ID if the given value is not nil.
func (wcc *WorkflowContractCreate) SetNillableOrganizationID(id *uuid.UUID) *WorkflowContractCreate {
	if id != nil {
		wcc = wcc.SetOrganizationID(*id)
	}
	return wcc
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (wcc *WorkflowContractCreate) SetOrganization(o *Organization) *WorkflowContractCreate {
	return wcc.SetOrganizationID(o.ID)
}

// AddWorkflowIDs adds the "workflows" edge to the Workflow entity by IDs.
func (wcc *WorkflowContractCreate) AddWorkflowIDs(ids ...uuid.UUID) *WorkflowContractCreate {
	wcc.mutation.AddWorkflowIDs(ids...)
	return wcc
}

// AddWorkflows adds the "workflows" edges to the Workflow entity.
func (wcc *WorkflowContractCreate) AddWorkflows(w ...*Workflow) *WorkflowContractCreate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wcc.AddWorkflowIDs(ids...)
}

// Mutation returns the WorkflowContractMutation object of the builder.
func (wcc *WorkflowContractCreate) Mutation() *WorkflowContractMutation {
	return wcc.mutation
}

// Save creates the WorkflowContract in the database.
func (wcc *WorkflowContractCreate) Save(ctx context.Context) (*WorkflowContract, error) {
	wcc.defaults()
	return withHooks(ctx, wcc.sqlSave, wcc.mutation, wcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wcc *WorkflowContractCreate) SaveX(ctx context.Context) *WorkflowContract {
	v, err := wcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcc *WorkflowContractCreate) Exec(ctx context.Context) error {
	_, err := wcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcc *WorkflowContractCreate) ExecX(ctx context.Context) {
	if err := wcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wcc *WorkflowContractCreate) defaults() {
	if _, ok := wcc.mutation.CreatedAt(); !ok {
		v := workflowcontract.DefaultCreatedAt()
		wcc.mutation.SetCreatedAt(v)
	}
	if _, ok := wcc.mutation.ID(); !ok {
		v := workflowcontract.DefaultID()
		wcc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wcc *WorkflowContractCreate) check() error {
	if _, ok := wcc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "WorkflowContract.name"`)}
	}
	if _, ok := wcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "WorkflowContract.created_at"`)}
	}
	return nil
}

func (wcc *WorkflowContractCreate) sqlSave(ctx context.Context) (*WorkflowContract, error) {
	if err := wcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	wcc.mutation.id = &_node.ID
	wcc.mutation.done = true
	return _node, nil
}

func (wcc *WorkflowContractCreate) createSpec() (*WorkflowContract, *sqlgraph.CreateSpec) {
	var (
		_node = &WorkflowContract{config: wcc.config}
		_spec = sqlgraph.NewCreateSpec(workflowcontract.Table, sqlgraph.NewFieldSpec(workflowcontract.FieldID, field.TypeUUID))
	)
	if id, ok := wcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := wcc.mutation.Name(); ok {
		_spec.SetField(workflowcontract.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := wcc.mutation.CreatedAt(); ok {
		_spec.SetField(workflowcontract.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := wcc.mutation.DeletedAt(); ok {
		_spec.SetField(workflowcontract.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := wcc.mutation.Description(); ok {
		_spec.SetField(workflowcontract.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := wcc.mutation.VersionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowcontract.VersionsTable,
			Columns: []string{workflowcontract.VersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowcontractversion.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wcc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflowcontract.OrganizationTable,
			Columns: []string{workflowcontract.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.organization_workflow_contracts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wcc.mutation.WorkflowsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   workflowcontract.WorkflowsTable,
			Columns: []string{workflowcontract.WorkflowsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WorkflowContractCreateBulk is the builder for creating many WorkflowContract entities in bulk.
type WorkflowContractCreateBulk struct {
	config
	builders []*WorkflowContractCreate
}

// Save creates the WorkflowContract entities in the database.
func (wccb *WorkflowContractCreateBulk) Save(ctx context.Context) ([]*WorkflowContract, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wccb.builders))
	nodes := make([]*WorkflowContract, len(wccb.builders))
	mutators := make([]Mutator, len(wccb.builders))
	for i := range wccb.builders {
		func(i int, root context.Context) {
			builder := wccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkflowContractMutation)
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
					_, err = mutators[i+1].Mutate(root, wccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wccb *WorkflowContractCreateBulk) SaveX(ctx context.Context) []*WorkflowContract {
	v, err := wccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wccb *WorkflowContractCreateBulk) Exec(ctx context.Context) error {
	_, err := wccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wccb *WorkflowContractCreateBulk) ExecX(ctx context.Context) {
	if err := wccb.Exec(ctx); err != nil {
		panic(err)
	}
}
