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
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/integrationattachment"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/organization"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/predicate"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/robotaccount"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/workflow"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/workflowcontract"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/workflowrun"
	"github.com/google/uuid"
)

// WorkflowUpdate is the builder for updating Workflow entities.
type WorkflowUpdate struct {
	config
	hooks    []Hook
	mutation *WorkflowMutation
}

// Where appends a list predicates to the WorkflowUpdate builder.
func (wu *WorkflowUpdate) Where(ps ...predicate.Workflow) *WorkflowUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetName sets the "name" field.
func (wu *WorkflowUpdate) SetName(s string) *WorkflowUpdate {
	wu.mutation.SetName(s)
	return wu
}

// SetProject sets the "project" field.
func (wu *WorkflowUpdate) SetProject(s string) *WorkflowUpdate {
	wu.mutation.SetProject(s)
	return wu
}

// SetNillableProject sets the "project" field if the given value is not nil.
func (wu *WorkflowUpdate) SetNillableProject(s *string) *WorkflowUpdate {
	if s != nil {
		wu.SetProject(*s)
	}
	return wu
}

// ClearProject clears the value of the "project" field.
func (wu *WorkflowUpdate) ClearProject() *WorkflowUpdate {
	wu.mutation.ClearProject()
	return wu
}

// SetTeam sets the "team" field.
func (wu *WorkflowUpdate) SetTeam(s string) *WorkflowUpdate {
	wu.mutation.SetTeam(s)
	return wu
}

// SetNillableTeam sets the "team" field if the given value is not nil.
func (wu *WorkflowUpdate) SetNillableTeam(s *string) *WorkflowUpdate {
	if s != nil {
		wu.SetTeam(*s)
	}
	return wu
}

// ClearTeam clears the value of the "team" field.
func (wu *WorkflowUpdate) ClearTeam() *WorkflowUpdate {
	wu.mutation.ClearTeam()
	return wu
}

// SetRunsCount sets the "runs_count" field.
func (wu *WorkflowUpdate) SetRunsCount(i int) *WorkflowUpdate {
	wu.mutation.ResetRunsCount()
	wu.mutation.SetRunsCount(i)
	return wu
}

// SetNillableRunsCount sets the "runs_count" field if the given value is not nil.
func (wu *WorkflowUpdate) SetNillableRunsCount(i *int) *WorkflowUpdate {
	if i != nil {
		wu.SetRunsCount(*i)
	}
	return wu
}

// AddRunsCount adds i to the "runs_count" field.
func (wu *WorkflowUpdate) AddRunsCount(i int) *WorkflowUpdate {
	wu.mutation.AddRunsCount(i)
	return wu
}

// SetDeletedAt sets the "deleted_at" field.
func (wu *WorkflowUpdate) SetDeletedAt(t time.Time) *WorkflowUpdate {
	wu.mutation.SetDeletedAt(t)
	return wu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (wu *WorkflowUpdate) SetNillableDeletedAt(t *time.Time) *WorkflowUpdate {
	if t != nil {
		wu.SetDeletedAt(*t)
	}
	return wu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (wu *WorkflowUpdate) ClearDeletedAt() *WorkflowUpdate {
	wu.mutation.ClearDeletedAt()
	return wu
}

// AddRobotaccountIDs adds the "robotaccounts" edge to the RobotAccount entity by IDs.
func (wu *WorkflowUpdate) AddRobotaccountIDs(ids ...uuid.UUID) *WorkflowUpdate {
	wu.mutation.AddRobotaccountIDs(ids...)
	return wu
}

// AddRobotaccounts adds the "robotaccounts" edges to the RobotAccount entity.
func (wu *WorkflowUpdate) AddRobotaccounts(r ...*RobotAccount) *WorkflowUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return wu.AddRobotaccountIDs(ids...)
}

// AddWorkflowrunIDs adds the "workflowruns" edge to the WorkflowRun entity by IDs.
func (wu *WorkflowUpdate) AddWorkflowrunIDs(ids ...uuid.UUID) *WorkflowUpdate {
	wu.mutation.AddWorkflowrunIDs(ids...)
	return wu
}

// AddWorkflowruns adds the "workflowruns" edges to the WorkflowRun entity.
func (wu *WorkflowUpdate) AddWorkflowruns(w ...*WorkflowRun) *WorkflowUpdate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.AddWorkflowrunIDs(ids...)
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (wu *WorkflowUpdate) SetOrganizationID(id uuid.UUID) *WorkflowUpdate {
	wu.mutation.SetOrganizationID(id)
	return wu
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (wu *WorkflowUpdate) SetOrganization(o *Organization) *WorkflowUpdate {
	return wu.SetOrganizationID(o.ID)
}

// SetContractID sets the "contract" edge to the WorkflowContract entity by ID.
func (wu *WorkflowUpdate) SetContractID(id uuid.UUID) *WorkflowUpdate {
	wu.mutation.SetContractID(id)
	return wu
}

// SetContract sets the "contract" edge to the WorkflowContract entity.
func (wu *WorkflowUpdate) SetContract(w *WorkflowContract) *WorkflowUpdate {
	return wu.SetContractID(w.ID)
}

// AddIntegrationAttachmentIDs adds the "integration_attachments" edge to the IntegrationAttachment entity by IDs.
func (wu *WorkflowUpdate) AddIntegrationAttachmentIDs(ids ...uuid.UUID) *WorkflowUpdate {
	wu.mutation.AddIntegrationAttachmentIDs(ids...)
	return wu
}

// AddIntegrationAttachments adds the "integration_attachments" edges to the IntegrationAttachment entity.
func (wu *WorkflowUpdate) AddIntegrationAttachments(i ...*IntegrationAttachment) *WorkflowUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return wu.AddIntegrationAttachmentIDs(ids...)
}

// Mutation returns the WorkflowMutation object of the builder.
func (wu *WorkflowUpdate) Mutation() *WorkflowMutation {
	return wu.mutation
}

// ClearRobotaccounts clears all "robotaccounts" edges to the RobotAccount entity.
func (wu *WorkflowUpdate) ClearRobotaccounts() *WorkflowUpdate {
	wu.mutation.ClearRobotaccounts()
	return wu
}

// RemoveRobotaccountIDs removes the "robotaccounts" edge to RobotAccount entities by IDs.
func (wu *WorkflowUpdate) RemoveRobotaccountIDs(ids ...uuid.UUID) *WorkflowUpdate {
	wu.mutation.RemoveRobotaccountIDs(ids...)
	return wu
}

// RemoveRobotaccounts removes "robotaccounts" edges to RobotAccount entities.
func (wu *WorkflowUpdate) RemoveRobotaccounts(r ...*RobotAccount) *WorkflowUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return wu.RemoveRobotaccountIDs(ids...)
}

// ClearWorkflowruns clears all "workflowruns" edges to the WorkflowRun entity.
func (wu *WorkflowUpdate) ClearWorkflowruns() *WorkflowUpdate {
	wu.mutation.ClearWorkflowruns()
	return wu
}

// RemoveWorkflowrunIDs removes the "workflowruns" edge to WorkflowRun entities by IDs.
func (wu *WorkflowUpdate) RemoveWorkflowrunIDs(ids ...uuid.UUID) *WorkflowUpdate {
	wu.mutation.RemoveWorkflowrunIDs(ids...)
	return wu
}

// RemoveWorkflowruns removes "workflowruns" edges to WorkflowRun entities.
func (wu *WorkflowUpdate) RemoveWorkflowruns(w ...*WorkflowRun) *WorkflowUpdate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.RemoveWorkflowrunIDs(ids...)
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (wu *WorkflowUpdate) ClearOrganization() *WorkflowUpdate {
	wu.mutation.ClearOrganization()
	return wu
}

// ClearContract clears the "contract" edge to the WorkflowContract entity.
func (wu *WorkflowUpdate) ClearContract() *WorkflowUpdate {
	wu.mutation.ClearContract()
	return wu
}

// ClearIntegrationAttachments clears all "integration_attachments" edges to the IntegrationAttachment entity.
func (wu *WorkflowUpdate) ClearIntegrationAttachments() *WorkflowUpdate {
	wu.mutation.ClearIntegrationAttachments()
	return wu
}

// RemoveIntegrationAttachmentIDs removes the "integration_attachments" edge to IntegrationAttachment entities by IDs.
func (wu *WorkflowUpdate) RemoveIntegrationAttachmentIDs(ids ...uuid.UUID) *WorkflowUpdate {
	wu.mutation.RemoveIntegrationAttachmentIDs(ids...)
	return wu
}

// RemoveIntegrationAttachments removes "integration_attachments" edges to IntegrationAttachment entities.
func (wu *WorkflowUpdate) RemoveIntegrationAttachments(i ...*IntegrationAttachment) *WorkflowUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return wu.RemoveIntegrationAttachmentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WorkflowUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, wu.sqlSave, wu.mutation, wu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WorkflowUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WorkflowUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WorkflowUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WorkflowUpdate) check() error {
	if _, ok := wu.mutation.OrganizationID(); wu.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Workflow.organization"`)
	}
	if _, ok := wu.mutation.ContractID(); wu.mutation.ContractCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Workflow.contract"`)
	}
	return nil
}

func (wu *WorkflowUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(workflow.Table, workflow.Columns, sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeUUID))
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.Name(); ok {
		_spec.SetField(workflow.FieldName, field.TypeString, value)
	}
	if value, ok := wu.mutation.Project(); ok {
		_spec.SetField(workflow.FieldProject, field.TypeString, value)
	}
	if wu.mutation.ProjectCleared() {
		_spec.ClearField(workflow.FieldProject, field.TypeString)
	}
	if value, ok := wu.mutation.Team(); ok {
		_spec.SetField(workflow.FieldTeam, field.TypeString, value)
	}
	if wu.mutation.TeamCleared() {
		_spec.ClearField(workflow.FieldTeam, field.TypeString)
	}
	if value, ok := wu.mutation.RunsCount(); ok {
		_spec.SetField(workflow.FieldRunsCount, field.TypeInt, value)
	}
	if value, ok := wu.mutation.AddedRunsCount(); ok {
		_spec.AddField(workflow.FieldRunsCount, field.TypeInt, value)
	}
	if value, ok := wu.mutation.DeletedAt(); ok {
		_spec.SetField(workflow.FieldDeletedAt, field.TypeTime, value)
	}
	if wu.mutation.DeletedAtCleared() {
		_spec.ClearField(workflow.FieldDeletedAt, field.TypeTime)
	}
	if wu.mutation.RobotaccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.RobotaccountsTable,
			Columns: []string{workflow.RobotaccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(robotaccount.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedRobotaccountsIDs(); len(nodes) > 0 && !wu.mutation.RobotaccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.RobotaccountsTable,
			Columns: []string{workflow.RobotaccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(robotaccount.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RobotaccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.RobotaccountsTable,
			Columns: []string{workflow.RobotaccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(robotaccount.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wu.mutation.WorkflowrunsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.WorkflowrunsTable,
			Columns: []string{workflow.WorkflowrunsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowrun.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedWorkflowrunsIDs(); len(nodes) > 0 && !wu.mutation.WorkflowrunsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.WorkflowrunsTable,
			Columns: []string{workflow.WorkflowrunsColumn},
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
	if nodes := wu.mutation.WorkflowrunsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.WorkflowrunsTable,
			Columns: []string{workflow.WorkflowrunsColumn},
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
	if wu.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflow.OrganizationTable,
			Columns: []string{workflow.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflow.OrganizationTable,
			Columns: []string{workflow.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wu.mutation.ContractCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workflow.ContractTable,
			Columns: []string{workflow.ContractColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowcontract.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.ContractIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workflow.ContractTable,
			Columns: []string{workflow.ContractColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowcontract.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wu.mutation.IntegrationAttachmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   workflow.IntegrationAttachmentsTable,
			Columns: []string{workflow.IntegrationAttachmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integrationattachment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedIntegrationAttachmentsIDs(); len(nodes) > 0 && !wu.mutation.IntegrationAttachmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   workflow.IntegrationAttachmentsTable,
			Columns: []string{workflow.IntegrationAttachmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integrationattachment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.IntegrationAttachmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   workflow.IntegrationAttachmentsTable,
			Columns: []string{workflow.IntegrationAttachmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integrationattachment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workflow.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wu.mutation.done = true
	return n, nil
}

// WorkflowUpdateOne is the builder for updating a single Workflow entity.
type WorkflowUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WorkflowMutation
}

// SetName sets the "name" field.
func (wuo *WorkflowUpdateOne) SetName(s string) *WorkflowUpdateOne {
	wuo.mutation.SetName(s)
	return wuo
}

// SetProject sets the "project" field.
func (wuo *WorkflowUpdateOne) SetProject(s string) *WorkflowUpdateOne {
	wuo.mutation.SetProject(s)
	return wuo
}

// SetNillableProject sets the "project" field if the given value is not nil.
func (wuo *WorkflowUpdateOne) SetNillableProject(s *string) *WorkflowUpdateOne {
	if s != nil {
		wuo.SetProject(*s)
	}
	return wuo
}

// ClearProject clears the value of the "project" field.
func (wuo *WorkflowUpdateOne) ClearProject() *WorkflowUpdateOne {
	wuo.mutation.ClearProject()
	return wuo
}

// SetTeam sets the "team" field.
func (wuo *WorkflowUpdateOne) SetTeam(s string) *WorkflowUpdateOne {
	wuo.mutation.SetTeam(s)
	return wuo
}

// SetNillableTeam sets the "team" field if the given value is not nil.
func (wuo *WorkflowUpdateOne) SetNillableTeam(s *string) *WorkflowUpdateOne {
	if s != nil {
		wuo.SetTeam(*s)
	}
	return wuo
}

// ClearTeam clears the value of the "team" field.
func (wuo *WorkflowUpdateOne) ClearTeam() *WorkflowUpdateOne {
	wuo.mutation.ClearTeam()
	return wuo
}

// SetRunsCount sets the "runs_count" field.
func (wuo *WorkflowUpdateOne) SetRunsCount(i int) *WorkflowUpdateOne {
	wuo.mutation.ResetRunsCount()
	wuo.mutation.SetRunsCount(i)
	return wuo
}

// SetNillableRunsCount sets the "runs_count" field if the given value is not nil.
func (wuo *WorkflowUpdateOne) SetNillableRunsCount(i *int) *WorkflowUpdateOne {
	if i != nil {
		wuo.SetRunsCount(*i)
	}
	return wuo
}

// AddRunsCount adds i to the "runs_count" field.
func (wuo *WorkflowUpdateOne) AddRunsCount(i int) *WorkflowUpdateOne {
	wuo.mutation.AddRunsCount(i)
	return wuo
}

// SetDeletedAt sets the "deleted_at" field.
func (wuo *WorkflowUpdateOne) SetDeletedAt(t time.Time) *WorkflowUpdateOne {
	wuo.mutation.SetDeletedAt(t)
	return wuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (wuo *WorkflowUpdateOne) SetNillableDeletedAt(t *time.Time) *WorkflowUpdateOne {
	if t != nil {
		wuo.SetDeletedAt(*t)
	}
	return wuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (wuo *WorkflowUpdateOne) ClearDeletedAt() *WorkflowUpdateOne {
	wuo.mutation.ClearDeletedAt()
	return wuo
}

// AddRobotaccountIDs adds the "robotaccounts" edge to the RobotAccount entity by IDs.
func (wuo *WorkflowUpdateOne) AddRobotaccountIDs(ids ...uuid.UUID) *WorkflowUpdateOne {
	wuo.mutation.AddRobotaccountIDs(ids...)
	return wuo
}

// AddRobotaccounts adds the "robotaccounts" edges to the RobotAccount entity.
func (wuo *WorkflowUpdateOne) AddRobotaccounts(r ...*RobotAccount) *WorkflowUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return wuo.AddRobotaccountIDs(ids...)
}

// AddWorkflowrunIDs adds the "workflowruns" edge to the WorkflowRun entity by IDs.
func (wuo *WorkflowUpdateOne) AddWorkflowrunIDs(ids ...uuid.UUID) *WorkflowUpdateOne {
	wuo.mutation.AddWorkflowrunIDs(ids...)
	return wuo
}

// AddWorkflowruns adds the "workflowruns" edges to the WorkflowRun entity.
func (wuo *WorkflowUpdateOne) AddWorkflowruns(w ...*WorkflowRun) *WorkflowUpdateOne {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.AddWorkflowrunIDs(ids...)
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (wuo *WorkflowUpdateOne) SetOrganizationID(id uuid.UUID) *WorkflowUpdateOne {
	wuo.mutation.SetOrganizationID(id)
	return wuo
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (wuo *WorkflowUpdateOne) SetOrganization(o *Organization) *WorkflowUpdateOne {
	return wuo.SetOrganizationID(o.ID)
}

// SetContractID sets the "contract" edge to the WorkflowContract entity by ID.
func (wuo *WorkflowUpdateOne) SetContractID(id uuid.UUID) *WorkflowUpdateOne {
	wuo.mutation.SetContractID(id)
	return wuo
}

// SetContract sets the "contract" edge to the WorkflowContract entity.
func (wuo *WorkflowUpdateOne) SetContract(w *WorkflowContract) *WorkflowUpdateOne {
	return wuo.SetContractID(w.ID)
}

// AddIntegrationAttachmentIDs adds the "integration_attachments" edge to the IntegrationAttachment entity by IDs.
func (wuo *WorkflowUpdateOne) AddIntegrationAttachmentIDs(ids ...uuid.UUID) *WorkflowUpdateOne {
	wuo.mutation.AddIntegrationAttachmentIDs(ids...)
	return wuo
}

// AddIntegrationAttachments adds the "integration_attachments" edges to the IntegrationAttachment entity.
func (wuo *WorkflowUpdateOne) AddIntegrationAttachments(i ...*IntegrationAttachment) *WorkflowUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return wuo.AddIntegrationAttachmentIDs(ids...)
}

// Mutation returns the WorkflowMutation object of the builder.
func (wuo *WorkflowUpdateOne) Mutation() *WorkflowMutation {
	return wuo.mutation
}

// ClearRobotaccounts clears all "robotaccounts" edges to the RobotAccount entity.
func (wuo *WorkflowUpdateOne) ClearRobotaccounts() *WorkflowUpdateOne {
	wuo.mutation.ClearRobotaccounts()
	return wuo
}

// RemoveRobotaccountIDs removes the "robotaccounts" edge to RobotAccount entities by IDs.
func (wuo *WorkflowUpdateOne) RemoveRobotaccountIDs(ids ...uuid.UUID) *WorkflowUpdateOne {
	wuo.mutation.RemoveRobotaccountIDs(ids...)
	return wuo
}

// RemoveRobotaccounts removes "robotaccounts" edges to RobotAccount entities.
func (wuo *WorkflowUpdateOne) RemoveRobotaccounts(r ...*RobotAccount) *WorkflowUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return wuo.RemoveRobotaccountIDs(ids...)
}

// ClearWorkflowruns clears all "workflowruns" edges to the WorkflowRun entity.
func (wuo *WorkflowUpdateOne) ClearWorkflowruns() *WorkflowUpdateOne {
	wuo.mutation.ClearWorkflowruns()
	return wuo
}

// RemoveWorkflowrunIDs removes the "workflowruns" edge to WorkflowRun entities by IDs.
func (wuo *WorkflowUpdateOne) RemoveWorkflowrunIDs(ids ...uuid.UUID) *WorkflowUpdateOne {
	wuo.mutation.RemoveWorkflowrunIDs(ids...)
	return wuo
}

// RemoveWorkflowruns removes "workflowruns" edges to WorkflowRun entities.
func (wuo *WorkflowUpdateOne) RemoveWorkflowruns(w ...*WorkflowRun) *WorkflowUpdateOne {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.RemoveWorkflowrunIDs(ids...)
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (wuo *WorkflowUpdateOne) ClearOrganization() *WorkflowUpdateOne {
	wuo.mutation.ClearOrganization()
	return wuo
}

// ClearContract clears the "contract" edge to the WorkflowContract entity.
func (wuo *WorkflowUpdateOne) ClearContract() *WorkflowUpdateOne {
	wuo.mutation.ClearContract()
	return wuo
}

// ClearIntegrationAttachments clears all "integration_attachments" edges to the IntegrationAttachment entity.
func (wuo *WorkflowUpdateOne) ClearIntegrationAttachments() *WorkflowUpdateOne {
	wuo.mutation.ClearIntegrationAttachments()
	return wuo
}

// RemoveIntegrationAttachmentIDs removes the "integration_attachments" edge to IntegrationAttachment entities by IDs.
func (wuo *WorkflowUpdateOne) RemoveIntegrationAttachmentIDs(ids ...uuid.UUID) *WorkflowUpdateOne {
	wuo.mutation.RemoveIntegrationAttachmentIDs(ids...)
	return wuo
}

// RemoveIntegrationAttachments removes "integration_attachments" edges to IntegrationAttachment entities.
func (wuo *WorkflowUpdateOne) RemoveIntegrationAttachments(i ...*IntegrationAttachment) *WorkflowUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return wuo.RemoveIntegrationAttachmentIDs(ids...)
}

// Where appends a list predicates to the WorkflowUpdate builder.
func (wuo *WorkflowUpdateOne) Where(ps ...predicate.Workflow) *WorkflowUpdateOne {
	wuo.mutation.Where(ps...)
	return wuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WorkflowUpdateOne) Select(field string, fields ...string) *WorkflowUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Workflow entity.
func (wuo *WorkflowUpdateOne) Save(ctx context.Context) (*Workflow, error) {
	return withHooks(ctx, wuo.sqlSave, wuo.mutation, wuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WorkflowUpdateOne) SaveX(ctx context.Context) *Workflow {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WorkflowUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WorkflowUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WorkflowUpdateOne) check() error {
	if _, ok := wuo.mutation.OrganizationID(); wuo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Workflow.organization"`)
	}
	if _, ok := wuo.mutation.ContractID(); wuo.mutation.ContractCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Workflow.contract"`)
	}
	return nil
}

func (wuo *WorkflowUpdateOne) sqlSave(ctx context.Context) (_node *Workflow, err error) {
	if err := wuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(workflow.Table, workflow.Columns, sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeUUID))
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Workflow.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workflow.FieldID)
		for _, f := range fields {
			if !workflow.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != workflow.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.Name(); ok {
		_spec.SetField(workflow.FieldName, field.TypeString, value)
	}
	if value, ok := wuo.mutation.Project(); ok {
		_spec.SetField(workflow.FieldProject, field.TypeString, value)
	}
	if wuo.mutation.ProjectCleared() {
		_spec.ClearField(workflow.FieldProject, field.TypeString)
	}
	if value, ok := wuo.mutation.Team(); ok {
		_spec.SetField(workflow.FieldTeam, field.TypeString, value)
	}
	if wuo.mutation.TeamCleared() {
		_spec.ClearField(workflow.FieldTeam, field.TypeString)
	}
	if value, ok := wuo.mutation.RunsCount(); ok {
		_spec.SetField(workflow.FieldRunsCount, field.TypeInt, value)
	}
	if value, ok := wuo.mutation.AddedRunsCount(); ok {
		_spec.AddField(workflow.FieldRunsCount, field.TypeInt, value)
	}
	if value, ok := wuo.mutation.DeletedAt(); ok {
		_spec.SetField(workflow.FieldDeletedAt, field.TypeTime, value)
	}
	if wuo.mutation.DeletedAtCleared() {
		_spec.ClearField(workflow.FieldDeletedAt, field.TypeTime)
	}
	if wuo.mutation.RobotaccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.RobotaccountsTable,
			Columns: []string{workflow.RobotaccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(robotaccount.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedRobotaccountsIDs(); len(nodes) > 0 && !wuo.mutation.RobotaccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.RobotaccountsTable,
			Columns: []string{workflow.RobotaccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(robotaccount.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RobotaccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.RobotaccountsTable,
			Columns: []string{workflow.RobotaccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(robotaccount.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wuo.mutation.WorkflowrunsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.WorkflowrunsTable,
			Columns: []string{workflow.WorkflowrunsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowrun.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedWorkflowrunsIDs(); len(nodes) > 0 && !wuo.mutation.WorkflowrunsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.WorkflowrunsTable,
			Columns: []string{workflow.WorkflowrunsColumn},
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
	if nodes := wuo.mutation.WorkflowrunsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.WorkflowrunsTable,
			Columns: []string{workflow.WorkflowrunsColumn},
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
	if wuo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflow.OrganizationTable,
			Columns: []string{workflow.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflow.OrganizationTable,
			Columns: []string{workflow.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wuo.mutation.ContractCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workflow.ContractTable,
			Columns: []string{workflow.ContractColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowcontract.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.ContractIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workflow.ContractTable,
			Columns: []string{workflow.ContractColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowcontract.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wuo.mutation.IntegrationAttachmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   workflow.IntegrationAttachmentsTable,
			Columns: []string{workflow.IntegrationAttachmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integrationattachment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedIntegrationAttachmentsIDs(); len(nodes) > 0 && !wuo.mutation.IntegrationAttachmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   workflow.IntegrationAttachmentsTable,
			Columns: []string{workflow.IntegrationAttachmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integrationattachment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.IntegrationAttachmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   workflow.IntegrationAttachmentsTable,
			Columns: []string{workflow.IntegrationAttachmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integrationattachment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Workflow{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workflow.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wuo.mutation.done = true
	return _node, nil
}
