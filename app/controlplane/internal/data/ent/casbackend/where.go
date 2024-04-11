// Code generated by ent, DO NOT EDIT.

package casbackend

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/biz"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldLTE(FieldID, id))
}

// Location applies equality check predicate on the "location" field. It's identical to LocationEQ.
func Location(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldEQ(FieldLocation, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldEQ(FieldDescription, v))
}

// SecretName applies equality check predicate on the "secret_name" field. It's identical to SecretNameEQ.
func SecretName(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldEQ(FieldSecretName, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldEQ(FieldCreatedAt, v))
}

// ValidatedAt applies equality check predicate on the "validated_at" field. It's identical to ValidatedAtEQ.
func ValidatedAt(v time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldEQ(FieldValidatedAt, v))
}

// Default applies equality check predicate on the "default" field. It's identical to DefaultEQ.
func Default(v bool) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldEQ(FieldDefault, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldEQ(FieldDeletedAt, v))
}

// Fallback applies equality check predicate on the "fallback" field. It's identical to FallbackEQ.
func Fallback(v bool) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldEQ(FieldFallback, v))
}

// LocationEQ applies the EQ predicate on the "location" field.
func LocationEQ(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldEQ(FieldLocation, v))
}

// LocationNEQ applies the NEQ predicate on the "location" field.
func LocationNEQ(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldNEQ(FieldLocation, v))
}

// LocationIn applies the In predicate on the "location" field.
func LocationIn(vs ...string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldIn(FieldLocation, vs...))
}

// LocationNotIn applies the NotIn predicate on the "location" field.
func LocationNotIn(vs ...string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldNotIn(FieldLocation, vs...))
}

// LocationGT applies the GT predicate on the "location" field.
func LocationGT(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldGT(FieldLocation, v))
}

// LocationGTE applies the GTE predicate on the "location" field.
func LocationGTE(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldGTE(FieldLocation, v))
}

// LocationLT applies the LT predicate on the "location" field.
func LocationLT(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldLT(FieldLocation, v))
}

// LocationLTE applies the LTE predicate on the "location" field.
func LocationLTE(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldLTE(FieldLocation, v))
}

// LocationContains applies the Contains predicate on the "location" field.
func LocationContains(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldContains(FieldLocation, v))
}

// LocationHasPrefix applies the HasPrefix predicate on the "location" field.
func LocationHasPrefix(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldHasPrefix(FieldLocation, v))
}

// LocationHasSuffix applies the HasSuffix predicate on the "location" field.
func LocationHasSuffix(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldHasSuffix(FieldLocation, v))
}

// LocationEqualFold applies the EqualFold predicate on the "location" field.
func LocationEqualFold(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldEqualFold(FieldLocation, v))
}

// LocationContainsFold applies the ContainsFold predicate on the "location" field.
func LocationContainsFold(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldContainsFold(FieldLocation, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldContainsFold(FieldName, v))
}

// ProviderEQ applies the EQ predicate on the "provider" field.
func ProviderEQ(v biz.CASBackendProvider) predicate.CASBackend {
	vc := v
	return predicate.CASBackend(sql.FieldEQ(FieldProvider, vc))
}

// ProviderNEQ applies the NEQ predicate on the "provider" field.
func ProviderNEQ(v biz.CASBackendProvider) predicate.CASBackend {
	vc := v
	return predicate.CASBackend(sql.FieldNEQ(FieldProvider, vc))
}

// ProviderIn applies the In predicate on the "provider" field.
func ProviderIn(vs ...biz.CASBackendProvider) predicate.CASBackend {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CASBackend(sql.FieldIn(FieldProvider, v...))
}

// ProviderNotIn applies the NotIn predicate on the "provider" field.
func ProviderNotIn(vs ...biz.CASBackendProvider) predicate.CASBackend {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CASBackend(sql.FieldNotIn(FieldProvider, v...))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.CASBackend {
	return predicate.CASBackend(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.CASBackend {
	return predicate.CASBackend(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldContainsFold(FieldDescription, v))
}

// SecretNameEQ applies the EQ predicate on the "secret_name" field.
func SecretNameEQ(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldEQ(FieldSecretName, v))
}

// SecretNameNEQ applies the NEQ predicate on the "secret_name" field.
func SecretNameNEQ(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldNEQ(FieldSecretName, v))
}

// SecretNameIn applies the In predicate on the "secret_name" field.
func SecretNameIn(vs ...string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldIn(FieldSecretName, vs...))
}

// SecretNameNotIn applies the NotIn predicate on the "secret_name" field.
func SecretNameNotIn(vs ...string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldNotIn(FieldSecretName, vs...))
}

// SecretNameGT applies the GT predicate on the "secret_name" field.
func SecretNameGT(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldGT(FieldSecretName, v))
}

// SecretNameGTE applies the GTE predicate on the "secret_name" field.
func SecretNameGTE(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldGTE(FieldSecretName, v))
}

// SecretNameLT applies the LT predicate on the "secret_name" field.
func SecretNameLT(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldLT(FieldSecretName, v))
}

// SecretNameLTE applies the LTE predicate on the "secret_name" field.
func SecretNameLTE(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldLTE(FieldSecretName, v))
}

// SecretNameContains applies the Contains predicate on the "secret_name" field.
func SecretNameContains(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldContains(FieldSecretName, v))
}

// SecretNameHasPrefix applies the HasPrefix predicate on the "secret_name" field.
func SecretNameHasPrefix(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldHasPrefix(FieldSecretName, v))
}

// SecretNameHasSuffix applies the HasSuffix predicate on the "secret_name" field.
func SecretNameHasSuffix(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldHasSuffix(FieldSecretName, v))
}

// SecretNameEqualFold applies the EqualFold predicate on the "secret_name" field.
func SecretNameEqualFold(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldEqualFold(FieldSecretName, v))
}

// SecretNameContainsFold applies the ContainsFold predicate on the "secret_name" field.
func SecretNameContainsFold(v string) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldContainsFold(FieldSecretName, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldLTE(FieldCreatedAt, v))
}

// ValidationStatusEQ applies the EQ predicate on the "validation_status" field.
func ValidationStatusEQ(v biz.CASBackendValidationStatus) predicate.CASBackend {
	vc := v
	return predicate.CASBackend(sql.FieldEQ(FieldValidationStatus, vc))
}

// ValidationStatusNEQ applies the NEQ predicate on the "validation_status" field.
func ValidationStatusNEQ(v biz.CASBackendValidationStatus) predicate.CASBackend {
	vc := v
	return predicate.CASBackend(sql.FieldNEQ(FieldValidationStatus, vc))
}

// ValidationStatusIn applies the In predicate on the "validation_status" field.
func ValidationStatusIn(vs ...biz.CASBackendValidationStatus) predicate.CASBackend {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CASBackend(sql.FieldIn(FieldValidationStatus, v...))
}

// ValidationStatusNotIn applies the NotIn predicate on the "validation_status" field.
func ValidationStatusNotIn(vs ...biz.CASBackendValidationStatus) predicate.CASBackend {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CASBackend(sql.FieldNotIn(FieldValidationStatus, v...))
}

// ValidatedAtEQ applies the EQ predicate on the "validated_at" field.
func ValidatedAtEQ(v time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldEQ(FieldValidatedAt, v))
}

// ValidatedAtNEQ applies the NEQ predicate on the "validated_at" field.
func ValidatedAtNEQ(v time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldNEQ(FieldValidatedAt, v))
}

// ValidatedAtIn applies the In predicate on the "validated_at" field.
func ValidatedAtIn(vs ...time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldIn(FieldValidatedAt, vs...))
}

// ValidatedAtNotIn applies the NotIn predicate on the "validated_at" field.
func ValidatedAtNotIn(vs ...time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldNotIn(FieldValidatedAt, vs...))
}

// ValidatedAtGT applies the GT predicate on the "validated_at" field.
func ValidatedAtGT(v time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldGT(FieldValidatedAt, v))
}

// ValidatedAtGTE applies the GTE predicate on the "validated_at" field.
func ValidatedAtGTE(v time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldGTE(FieldValidatedAt, v))
}

// ValidatedAtLT applies the LT predicate on the "validated_at" field.
func ValidatedAtLT(v time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldLT(FieldValidatedAt, v))
}

// ValidatedAtLTE applies the LTE predicate on the "validated_at" field.
func ValidatedAtLTE(v time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldLTE(FieldValidatedAt, v))
}

// DefaultEQ applies the EQ predicate on the "default" field.
func DefaultEQ(v bool) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldEQ(FieldDefault, v))
}

// DefaultNEQ applies the NEQ predicate on the "default" field.
func DefaultNEQ(v bool) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldNEQ(FieldDefault, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.CASBackend {
	return predicate.CASBackend(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.CASBackend {
	return predicate.CASBackend(sql.FieldNotNull(FieldDeletedAt))
}

// FallbackEQ applies the EQ predicate on the "fallback" field.
func FallbackEQ(v bool) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldEQ(FieldFallback, v))
}

// FallbackNEQ applies the NEQ predicate on the "fallback" field.
func FallbackNEQ(v bool) predicate.CASBackend {
	return predicate.CASBackend(sql.FieldNEQ(FieldFallback, v))
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.CASBackend {
	return predicate.CASBackend(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.CASBackend {
	return predicate.CASBackend(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWorkflowRun applies the HasEdge predicate on the "workflow_run" edge.
func HasWorkflowRun() predicate.CASBackend {
	return predicate.CASBackend(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, WorkflowRunTable, WorkflowRunPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkflowRunWith applies the HasEdge predicate on the "workflow_run" edge with a given conditions (other predicates).
func HasWorkflowRunWith(preds ...predicate.WorkflowRun) predicate.CASBackend {
	return predicate.CASBackend(func(s *sql.Selector) {
		step := newWorkflowRunStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CASBackend) predicate.CASBackend {
	return predicate.CASBackend(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CASBackend) predicate.CASBackend {
	return predicate.CASBackend(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CASBackend) predicate.CASBackend {
	return predicate.CASBackend(func(s *sql.Selector) {
		p(s.Not())
	})
}
