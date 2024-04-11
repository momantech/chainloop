// Code generated by ent, DO NOT EDIT.

package casbackend

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/biz"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the casbackend type in the database.
	Label = "cas_backend"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldProvider holds the string denoting the provider field in the database.
	FieldProvider = "provider"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldSecretName holds the string denoting the secret_name field in the database.
	FieldSecretName = "secret_name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldValidationStatus holds the string denoting the validation_status field in the database.
	FieldValidationStatus = "validation_status"
	// FieldValidatedAt holds the string denoting the validated_at field in the database.
	FieldValidatedAt = "validated_at"
	// FieldDefault holds the string denoting the default field in the database.
	FieldDefault = "default"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldFallback holds the string denoting the fallback field in the database.
	FieldFallback = "fallback"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// EdgeWorkflowRun holds the string denoting the workflow_run edge name in mutations.
	EdgeWorkflowRun = "workflow_run"
	// Table holds the table name of the casbackend in the database.
	Table = "cas_backends"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "cas_backends"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organizations"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "organization_cas_backends"
	// WorkflowRunTable is the table that holds the workflow_run relation/edge. The primary key declared below.
	WorkflowRunTable = "workflow_run_cas_backends"
	// WorkflowRunInverseTable is the table name for the WorkflowRun entity.
	// It exists in this package in order to avoid circular dependency with the "workflowrun" package.
	WorkflowRunInverseTable = "workflow_runs"
)

// Columns holds all SQL columns for casbackend fields.
var Columns = []string{
	FieldID,
	FieldLocation,
	FieldName,
	FieldProvider,
	FieldDescription,
	FieldSecretName,
	FieldCreatedAt,
	FieldValidationStatus,
	FieldValidatedAt,
	FieldDefault,
	FieldDeletedAt,
	FieldFallback,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "cas_backends"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"organization_cas_backends",
}

var (
	// WorkflowRunPrimaryKey and WorkflowRunColumn2 are the table columns denoting the
	// primary key for the workflow_run relation (M2M).
	WorkflowRunPrimaryKey = []string{"workflow_run_id", "cas_backend_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultValidatedAt holds the default value on creation for the "validated_at" field.
	DefaultValidatedAt func() time.Time
	// DefaultDefault holds the default value on creation for the "default" field.
	DefaultDefault bool
	// DefaultFallback holds the default value on creation for the "fallback" field.
	DefaultFallback bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// ProviderValidator is a validator for the "provider" field enum values. It is called by the builders before save.
func ProviderValidator(pr biz.CASBackendProvider) error {
	switch pr {
	case "AzureBlob", "OCI", "INLINE", "AWS-S3":
		return nil
	default:
		return fmt.Errorf("casbackend: invalid enum value for provider field: %q", pr)
	}
}

const DefaultValidationStatus biz.CASBackendValidationStatus = "OK"

// ValidationStatusValidator is a validator for the "validation_status" field enum values. It is called by the builders before save.
func ValidationStatusValidator(vs biz.CASBackendValidationStatus) error {
	switch vs {
	case "OK", "Invalid":
		return nil
	default:
		return fmt.Errorf("casbackend: invalid enum value for validation_status field: %q", vs)
	}
}

// OrderOption defines the ordering options for the CASBackend queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByLocation orders the results by the location field.
func ByLocation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocation, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByProvider orders the results by the provider field.
func ByProvider(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProvider, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// BySecretName orders the results by the secret_name field.
func BySecretName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSecretName, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByValidationStatus orders the results by the validation_status field.
func ByValidationStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValidationStatus, opts...).ToFunc()
}

// ByValidatedAt orders the results by the validated_at field.
func ByValidatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValidatedAt, opts...).ToFunc()
}

// ByDefault orders the results by the default field.
func ByDefault(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDefault, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByFallback orders the results by the fallback field.
func ByFallback(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFallback, opts...).ToFunc()
}

// ByOrganizationField orders the results by organization field.
func ByOrganizationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrganizationStep(), sql.OrderByField(field, opts...))
	}
}

// ByWorkflowRunCount orders the results by workflow_run count.
func ByWorkflowRunCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkflowRunStep(), opts...)
	}
}

// ByWorkflowRun orders the results by workflow_run terms.
func ByWorkflowRun(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkflowRunStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOrganizationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrganizationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OrganizationTable, OrganizationColumn),
	)
}
func newWorkflowRunStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkflowRunInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, WorkflowRunTable, WorkflowRunPrimaryKey...),
	)
}
