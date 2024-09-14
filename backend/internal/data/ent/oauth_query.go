// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/oauth"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/predicate"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/user"
)

// OAuthQuery is the builder for querying OAuth entities.
type OAuthQuery struct {
	config
	ctx        *QueryContext
	order      []oauth.OrderOption
	inters     []Interceptor
	predicates []predicate.OAuth
	withUser   *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OAuthQuery builder.
func (oq *OAuthQuery) Where(ps ...predicate.OAuth) *OAuthQuery {
	oq.predicates = append(oq.predicates, ps...)
	return oq
}

// Limit the number of records to be returned by this query.
func (oq *OAuthQuery) Limit(limit int) *OAuthQuery {
	oq.ctx.Limit = &limit
	return oq
}

// Offset to start from.
func (oq *OAuthQuery) Offset(offset int) *OAuthQuery {
	oq.ctx.Offset = &offset
	return oq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oq *OAuthQuery) Unique(unique bool) *OAuthQuery {
	oq.ctx.Unique = &unique
	return oq
}

// Order specifies how the records should be ordered.
func (oq *OAuthQuery) Order(o ...oauth.OrderOption) *OAuthQuery {
	oq.order = append(oq.order, o...)
	return oq
}

// QueryUser chains the current query on the "user" edge.
func (oq *OAuthQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: oq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(oauth.Table, oauth.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, oauth.UserTable, oauth.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OAuth entity from the query.
// Returns a *NotFoundError when no OAuth was found.
func (oq *OAuthQuery) First(ctx context.Context) (*OAuth, error) {
	nodes, err := oq.Limit(1).All(setContextOp(ctx, oq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{oauth.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oq *OAuthQuery) FirstX(ctx context.Context) *OAuth {
	node, err := oq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OAuth ID from the query.
// Returns a *NotFoundError when no OAuth ID was found.
func (oq *OAuthQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = oq.Limit(1).IDs(setContextOp(ctx, oq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{oauth.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oq *OAuthQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := oq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OAuth entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OAuth entity is found.
// Returns a *NotFoundError when no OAuth entities are found.
func (oq *OAuthQuery) Only(ctx context.Context) (*OAuth, error) {
	nodes, err := oq.Limit(2).All(setContextOp(ctx, oq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{oauth.Label}
	default:
		return nil, &NotSingularError{oauth.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oq *OAuthQuery) OnlyX(ctx context.Context) *OAuth {
	node, err := oq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OAuth ID in the query.
// Returns a *NotSingularError when more than one OAuth ID is found.
// Returns a *NotFoundError when no entities are found.
func (oq *OAuthQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = oq.Limit(2).IDs(setContextOp(ctx, oq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{oauth.Label}
	default:
		err = &NotSingularError{oauth.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oq *OAuthQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := oq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OAuths.
func (oq *OAuthQuery) All(ctx context.Context) ([]*OAuth, error) {
	ctx = setContextOp(ctx, oq.ctx, "All")
	if err := oq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OAuth, *OAuthQuery]()
	return withInterceptors[[]*OAuth](ctx, oq, qr, oq.inters)
}

// AllX is like All, but panics if an error occurs.
func (oq *OAuthQuery) AllX(ctx context.Context) []*OAuth {
	nodes, err := oq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OAuth IDs.
func (oq *OAuthQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if oq.ctx.Unique == nil && oq.path != nil {
		oq.Unique(true)
	}
	ctx = setContextOp(ctx, oq.ctx, "IDs")
	if err = oq.Select(oauth.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oq *OAuthQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := oq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oq *OAuthQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, oq.ctx, "Count")
	if err := oq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, oq, querierCount[*OAuthQuery](), oq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (oq *OAuthQuery) CountX(ctx context.Context) int {
	count, err := oq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oq *OAuthQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, oq.ctx, "Exist")
	switch _, err := oq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (oq *OAuthQuery) ExistX(ctx context.Context) bool {
	exist, err := oq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OAuthQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oq *OAuthQuery) Clone() *OAuthQuery {
	if oq == nil {
		return nil
	}
	return &OAuthQuery{
		config:     oq.config,
		ctx:        oq.ctx.Clone(),
		order:      append([]oauth.OrderOption{}, oq.order...),
		inters:     append([]Interceptor{}, oq.inters...),
		predicates: append([]predicate.OAuth{}, oq.predicates...),
		withUser:   oq.withUser.Clone(),
		// clone intermediate query.
		sql:  oq.sql.Clone(),
		path: oq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OAuthQuery) WithUser(opts ...func(*UserQuery)) *OAuthQuery {
	query := (&UserClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oq.withUser = query
	return oq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OAuth.Query().
//		GroupBy(oauth.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (oq *OAuthQuery) GroupBy(field string, fields ...string) *OAuthGroupBy {
	oq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OAuthGroupBy{build: oq}
	grbuild.flds = &oq.ctx.Fields
	grbuild.label = oauth.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.OAuth.Query().
//		Select(oauth.FieldCreatedAt).
//		Scan(ctx, &v)
func (oq *OAuthQuery) Select(fields ...string) *OAuthSelect {
	oq.ctx.Fields = append(oq.ctx.Fields, fields...)
	sbuild := &OAuthSelect{OAuthQuery: oq}
	sbuild.label = oauth.Label
	sbuild.flds, sbuild.scan = &oq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OAuthSelect configured with the given aggregations.
func (oq *OAuthQuery) Aggregate(fns ...AggregateFunc) *OAuthSelect {
	return oq.Select().Aggregate(fns...)
}

func (oq *OAuthQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range oq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, oq); err != nil {
				return err
			}
		}
	}
	for _, f := range oq.ctx.Fields {
		if !oauth.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oq.path != nil {
		prev, err := oq.path(ctx)
		if err != nil {
			return err
		}
		oq.sql = prev
	}
	return nil
}

func (oq *OAuthQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OAuth, error) {
	var (
		nodes       = []*OAuth{}
		withFKs     = oq.withFKs
		_spec       = oq.querySpec()
		loadedTypes = [1]bool{
			oq.withUser != nil,
		}
	)
	if oq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, oauth.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OAuth).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OAuth{config: oq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := oq.withUser; query != nil {
		if err := oq.loadUser(ctx, query, nodes, nil,
			func(n *OAuth, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (oq *OAuthQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*OAuth, init func(*OAuth), assign func(*OAuth, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*OAuth)
	for i := range nodes {
		if nodes[i].user_oauth == nil {
			continue
		}
		fk := *nodes[i].user_oauth
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
			return fmt.Errorf(`unexpected foreign-key "user_oauth" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (oq *OAuthQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oq.querySpec()
	_spec.Node.Columns = oq.ctx.Fields
	if len(oq.ctx.Fields) > 0 {
		_spec.Unique = oq.ctx.Unique != nil && *oq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, oq.driver, _spec)
}

func (oq *OAuthQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(oauth.Table, oauth.Columns, sqlgraph.NewFieldSpec(oauth.FieldID, field.TypeUUID))
	_spec.From = oq.sql
	if unique := oq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if oq.path != nil {
		_spec.Unique = true
	}
	if fields := oq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oauth.FieldID)
		for i := range fields {
			if fields[i] != oauth.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oq *OAuthQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oq.driver.Dialect())
	t1 := builder.Table(oauth.Table)
	columns := oq.ctx.Fields
	if len(columns) == 0 {
		columns = oauth.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oq.sql != nil {
		selector = oq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oq.ctx.Unique != nil && *oq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range oq.predicates {
		p(selector)
	}
	for _, p := range oq.order {
		p(selector)
	}
	if offset := oq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OAuthGroupBy is the group-by builder for OAuth entities.
type OAuthGroupBy struct {
	selector
	build *OAuthQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ogb *OAuthGroupBy) Aggregate(fns ...AggregateFunc) *OAuthGroupBy {
	ogb.fns = append(ogb.fns, fns...)
	return ogb
}

// Scan applies the selector query and scans the result into the given value.
func (ogb *OAuthGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ogb.build.ctx, "GroupBy")
	if err := ogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OAuthQuery, *OAuthGroupBy](ctx, ogb.build, ogb, ogb.build.inters, v)
}

func (ogb *OAuthGroupBy) sqlScan(ctx context.Context, root *OAuthQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ogb.fns))
	for _, fn := range ogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ogb.flds)+len(ogb.fns))
		for _, f := range *ogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OAuthSelect is the builder for selecting fields of OAuth entities.
type OAuthSelect struct {
	*OAuthQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (os *OAuthSelect) Aggregate(fns ...AggregateFunc) *OAuthSelect {
	os.fns = append(os.fns, fns...)
	return os
}

// Scan applies the selector query and scans the result into the given value.
func (os *OAuthSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, os.ctx, "Select")
	if err := os.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OAuthQuery, *OAuthSelect](ctx, os.OAuthQuery, os, os.inters, v)
}

func (os *OAuthSelect) sqlScan(ctx context.Context, root *OAuthQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(os.fns))
	for _, fn := range os.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*os.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := os.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}