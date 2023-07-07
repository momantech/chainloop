// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/membership"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/organization"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/predicate"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/user"
	"github.com/google/uuid"
)

// MembershipQuery is the builder for querying Membership entities.
type MembershipQuery struct {
	config
	ctx              *QueryContext
	order            []membership.OrderOption
	inters           []Interceptor
	predicates       []predicate.Membership
	withOrganization *OrganizationQuery
	withUser         *UserQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MembershipQuery builder.
func (mq *MembershipQuery) Where(ps ...predicate.Membership) *MembershipQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit the number of records to be returned by this query.
func (mq *MembershipQuery) Limit(limit int) *MembershipQuery {
	mq.ctx.Limit = &limit
	return mq
}

// Offset to start from.
func (mq *MembershipQuery) Offset(offset int) *MembershipQuery {
	mq.ctx.Offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *MembershipQuery) Unique(unique bool) *MembershipQuery {
	mq.ctx.Unique = &unique
	return mq
}

// Order specifies how the records should be ordered.
func (mq *MembershipQuery) Order(o ...membership.OrderOption) *MembershipQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryOrganization chains the current query on the "organization" edge.
func (mq *MembershipQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(membership.Table, membership.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, membership.OrganizationTable, membership.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (mq *MembershipQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(membership.Table, membership.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, membership.UserTable, membership.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Membership entity from the query.
// Returns a *NotFoundError when no Membership was found.
func (mq *MembershipQuery) First(ctx context.Context) (*Membership, error) {
	nodes, err := mq.Limit(1).All(setContextOp(ctx, mq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{membership.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MembershipQuery) FirstX(ctx context.Context) *Membership {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Membership ID from the query.
// Returns a *NotFoundError when no Membership ID was found.
func (mq *MembershipQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mq.Limit(1).IDs(setContextOp(ctx, mq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{membership.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *MembershipQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Membership entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Membership entity is found.
// Returns a *NotFoundError when no Membership entities are found.
func (mq *MembershipQuery) Only(ctx context.Context) (*Membership, error) {
	nodes, err := mq.Limit(2).All(setContextOp(ctx, mq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{membership.Label}
	default:
		return nil, &NotSingularError{membership.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MembershipQuery) OnlyX(ctx context.Context) *Membership {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Membership ID in the query.
// Returns a *NotSingularError when more than one Membership ID is found.
// Returns a *NotFoundError when no entities are found.
func (mq *MembershipQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mq.Limit(2).IDs(setContextOp(ctx, mq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{membership.Label}
	default:
		err = &NotSingularError{membership.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MembershipQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Memberships.
func (mq *MembershipQuery) All(ctx context.Context) ([]*Membership, error) {
	ctx = setContextOp(ctx, mq.ctx, "All")
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Membership, *MembershipQuery]()
	return withInterceptors[[]*Membership](ctx, mq, qr, mq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mq *MembershipQuery) AllX(ctx context.Context) []*Membership {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Membership IDs.
func (mq *MembershipQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if mq.ctx.Unique == nil && mq.path != nil {
		mq.Unique(true)
	}
	ctx = setContextOp(ctx, mq.ctx, "IDs")
	if err = mq.Select(membership.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MembershipQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MembershipQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mq.ctx, "Count")
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mq, querierCount[*MembershipQuery](), mq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MembershipQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MembershipQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mq.ctx, "Exist")
	switch _, err := mq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MembershipQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MembershipQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MembershipQuery) Clone() *MembershipQuery {
	if mq == nil {
		return nil
	}
	return &MembershipQuery{
		config:           mq.config,
		ctx:              mq.ctx.Clone(),
		order:            append([]membership.OrderOption{}, mq.order...),
		inters:           append([]Interceptor{}, mq.inters...),
		predicates:       append([]predicate.Membership{}, mq.predicates...),
		withOrganization: mq.withOrganization.Clone(),
		withUser:         mq.withUser.Clone(),
		// clone intermediate query.
		sql:  mq.sql.Clone(),
		path: mq.path,
	}
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MembershipQuery) WithOrganization(opts ...func(*OrganizationQuery)) *MembershipQuery {
	query := (&OrganizationClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withOrganization = query
	return mq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MembershipQuery) WithUser(opts ...func(*UserQuery)) *MembershipQuery {
	query := (&UserClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withUser = query
	return mq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Current bool `json:"current,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Membership.Query().
//		GroupBy(membership.FieldCurrent).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mq *MembershipQuery) GroupBy(field string, fields ...string) *MembershipGroupBy {
	mq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MembershipGroupBy{build: mq}
	grbuild.flds = &mq.ctx.Fields
	grbuild.label = membership.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Current bool `json:"current,omitempty"`
//	}
//
//	client.Membership.Query().
//		Select(membership.FieldCurrent).
//		Scan(ctx, &v)
func (mq *MembershipQuery) Select(fields ...string) *MembershipSelect {
	mq.ctx.Fields = append(mq.ctx.Fields, fields...)
	sbuild := &MembershipSelect{MembershipQuery: mq}
	sbuild.label = membership.Label
	sbuild.flds, sbuild.scan = &mq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MembershipSelect configured with the given aggregations.
func (mq *MembershipQuery) Aggregate(fns ...AggregateFunc) *MembershipSelect {
	return mq.Select().Aggregate(fns...)
}

func (mq *MembershipQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mq); err != nil {
				return err
			}
		}
	}
	for _, f := range mq.ctx.Fields {
		if !membership.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MembershipQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Membership, error) {
	var (
		nodes       = []*Membership{}
		withFKs     = mq.withFKs
		_spec       = mq.querySpec()
		loadedTypes = [2]bool{
			mq.withOrganization != nil,
			mq.withUser != nil,
		}
	)
	if mq.withOrganization != nil || mq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, membership.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Membership).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Membership{config: mq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mq.withOrganization; query != nil {
		if err := mq.loadOrganization(ctx, query, nodes, nil,
			func(n *Membership, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := mq.withUser; query != nil {
		if err := mq.loadUser(ctx, query, nodes, nil,
			func(n *Membership, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mq *MembershipQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*Membership, init func(*Membership), assign func(*Membership, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Membership)
	for i := range nodes {
		if nodes[i].organization_memberships == nil {
			continue
		}
		fk := *nodes[i].organization_memberships
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(organization.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "organization_memberships" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mq *MembershipQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Membership, init func(*Membership), assign func(*Membership, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Membership)
	for i := range nodes {
		if nodes[i].user_memberships == nil {
			continue
		}
		fk := *nodes[i].user_memberships
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_memberships" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (mq *MembershipQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	_spec.Node.Columns = mq.ctx.Fields
	if len(mq.ctx.Fields) > 0 {
		_spec.Unique = mq.ctx.Unique != nil && *mq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MembershipQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(membership.Table, membership.Columns, sqlgraph.NewFieldSpec(membership.FieldID, field.TypeUUID))
	_spec.From = mq.sql
	if unique := mq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mq.path != nil {
		_spec.Unique = true
	}
	if fields := mq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, membership.FieldID)
		for i := range fields {
			if fields[i] != membership.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MembershipQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(membership.Table)
	columns := mq.ctx.Fields
	if len(columns) == 0 {
		columns = membership.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mq.ctx.Unique != nil && *mq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MembershipGroupBy is the group-by builder for Membership entities.
type MembershipGroupBy struct {
	selector
	build *MembershipQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MembershipGroupBy) Aggregate(fns ...AggregateFunc) *MembershipGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the selector query and scans the result into the given value.
func (mgb *MembershipGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mgb.build.ctx, "GroupBy")
	if err := mgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MembershipQuery, *MembershipGroupBy](ctx, mgb.build, mgb, mgb.build.inters, v)
}

func (mgb *MembershipGroupBy) sqlScan(ctx context.Context, root *MembershipQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mgb.fns))
	for _, fn := range mgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mgb.flds)+len(mgb.fns))
		for _, f := range *mgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MembershipSelect is the builder for selecting fields of Membership entities.
type MembershipSelect struct {
	*MembershipQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ms *MembershipSelect) Aggregate(fns ...AggregateFunc) *MembershipSelect {
	ms.fns = append(ms.fns, fns...)
	return ms
}

// Scan applies the selector query and scans the result into the given value.
func (ms *MembershipSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ms.ctx, "Select")
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MembershipQuery, *MembershipSelect](ctx, ms.MembershipQuery, ms, ms.inters, v)
}

func (ms *MembershipSelect) sqlScan(ctx context.Context, root *MembershipQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ms.fns))
	for _, fn := range ms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
