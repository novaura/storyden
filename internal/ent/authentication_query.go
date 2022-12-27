// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/internal/ent/account"
	"github.com/Southclaws/storyden/internal/ent/authentication"
	"github.com/Southclaws/storyden/internal/ent/predicate"
	"github.com/rs/xid"
)

// AuthenticationQuery is the builder for querying Authentication entities.
type AuthenticationQuery struct {
	config
	limit       *int
	offset      *int
	unique      *bool
	order       []OrderFunc
	fields      []string
	predicates  []predicate.Authentication
	withAccount *AccountQuery
	withFKs     bool
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AuthenticationQuery builder.
func (aq *AuthenticationQuery) Where(ps ...predicate.Authentication) *AuthenticationQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit adds a limit step to the query.
func (aq *AuthenticationQuery) Limit(limit int) *AuthenticationQuery {
	aq.limit = &limit
	return aq
}

// Offset adds an offset step to the query.
func (aq *AuthenticationQuery) Offset(offset int) *AuthenticationQuery {
	aq.offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AuthenticationQuery) Unique(unique bool) *AuthenticationQuery {
	aq.unique = &unique
	return aq
}

// Order adds an order step to the query.
func (aq *AuthenticationQuery) Order(o ...OrderFunc) *AuthenticationQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryAccount chains the current query on the "account" edge.
func (aq *AuthenticationQuery) QueryAccount() *AccountQuery {
	query := &AccountQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(authentication.Table, authentication.FieldID, selector),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, authentication.AccountTable, authentication.AccountColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Authentication entity from the query.
// Returns a *NotFoundError when no Authentication was found.
func (aq *AuthenticationQuery) First(ctx context.Context) (*Authentication, error) {
	nodes, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{authentication.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AuthenticationQuery) FirstX(ctx context.Context) *Authentication {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Authentication ID from the query.
// Returns a *NotFoundError when no Authentication ID was found.
func (aq *AuthenticationQuery) FirstID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = aq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{authentication.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AuthenticationQuery) FirstIDX(ctx context.Context) xid.ID {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Authentication entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Authentication entity is found.
// Returns a *NotFoundError when no Authentication entities are found.
func (aq *AuthenticationQuery) Only(ctx context.Context) (*Authentication, error) {
	nodes, err := aq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{authentication.Label}
	default:
		return nil, &NotSingularError{authentication.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AuthenticationQuery) OnlyX(ctx context.Context) *Authentication {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Authentication ID in the query.
// Returns a *NotSingularError when more than one Authentication ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AuthenticationQuery) OnlyID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = aq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{authentication.Label}
	default:
		err = &NotSingularError{authentication.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AuthenticationQuery) OnlyIDX(ctx context.Context) xid.ID {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Authentications.
func (aq *AuthenticationQuery) All(ctx context.Context) ([]*Authentication, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aq *AuthenticationQuery) AllX(ctx context.Context) []*Authentication {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Authentication IDs.
func (aq *AuthenticationQuery) IDs(ctx context.Context) ([]xid.ID, error) {
	var ids []xid.ID
	if err := aq.Select(authentication.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AuthenticationQuery) IDsX(ctx context.Context) []xid.ID {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AuthenticationQuery) Count(ctx context.Context) (int, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AuthenticationQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AuthenticationQuery) Exist(ctx context.Context) (bool, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AuthenticationQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AuthenticationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AuthenticationQuery) Clone() *AuthenticationQuery {
	if aq == nil {
		return nil
	}
	return &AuthenticationQuery{
		config:      aq.config,
		limit:       aq.limit,
		offset:      aq.offset,
		order:       append([]OrderFunc{}, aq.order...),
		predicates:  append([]predicate.Authentication{}, aq.predicates...),
		withAccount: aq.withAccount.Clone(),
		// clone intermediate query.
		sql:    aq.sql.Clone(),
		path:   aq.path,
		unique: aq.unique,
	}
}

// WithAccount tells the query-builder to eager-load the nodes that are connected to
// the "account" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AuthenticationQuery) WithAccount(opts ...func(*AccountQuery)) *AuthenticationQuery {
	query := &AccountQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withAccount = query
	return aq
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
//	client.Authentication.Query().
//		GroupBy(authentication.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AuthenticationQuery) GroupBy(field string, fields ...string) *AuthenticationGroupBy {
	grbuild := &AuthenticationGroupBy{config: aq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aq.sqlQuery(ctx), nil
	}
	grbuild.label = authentication.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
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
//	client.Authentication.Query().
//		Select(authentication.FieldCreatedAt).
//		Scan(ctx, &v)
func (aq *AuthenticationQuery) Select(fields ...string) *AuthenticationSelect {
	aq.fields = append(aq.fields, fields...)
	selbuild := &AuthenticationSelect{AuthenticationQuery: aq}
	selbuild.label = authentication.Label
	selbuild.flds, selbuild.scan = &aq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a AuthenticationSelect configured with the given aggregations.
func (aq *AuthenticationQuery) Aggregate(fns ...AggregateFunc) *AuthenticationSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AuthenticationQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aq.fields {
		if !authentication.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AuthenticationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Authentication, error) {
	var (
		nodes       = []*Authentication{}
		withFKs     = aq.withFKs
		_spec       = aq.querySpec()
		loadedTypes = [1]bool{
			aq.withAccount != nil,
		}
	)
	if aq.withAccount != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, authentication.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Authentication).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Authentication{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withAccount; query != nil {
		if err := aq.loadAccount(ctx, query, nodes, nil,
			func(n *Authentication, e *Account) { n.Edges.Account = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AuthenticationQuery) loadAccount(ctx context.Context, query *AccountQuery, nodes []*Authentication, init func(*Authentication), assign func(*Authentication, *Account)) error {
	ids := make([]xid.ID, 0, len(nodes))
	nodeids := make(map[xid.ID][]*Authentication)
	for i := range nodes {
		if nodes[i].account_authentication == nil {
			continue
		}
		fk := *nodes[i].account_authentication
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(account.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "account_authentication" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (aq *AuthenticationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	_spec.Node.Columns = aq.fields
	if len(aq.fields) > 0 {
		_spec.Unique = aq.unique != nil && *aq.unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AuthenticationQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (aq *AuthenticationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authentication.Table,
			Columns: authentication.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: authentication.FieldID,
			},
		},
		From:   aq.sql,
		Unique: true,
	}
	if unique := aq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authentication.FieldID)
		for i := range fields {
			if fields[i] != authentication.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AuthenticationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(authentication.Table)
	columns := aq.fields
	if len(columns) == 0 {
		columns = authentication.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.unique != nil && *aq.unique {
		selector.Distinct()
	}
	for _, m := range aq.modifiers {
		m(selector)
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (aq *AuthenticationQuery) Modify(modifiers ...func(s *sql.Selector)) *AuthenticationSelect {
	aq.modifiers = append(aq.modifiers, modifiers...)
	return aq.Select()
}

// AuthenticationGroupBy is the group-by builder for Authentication entities.
type AuthenticationGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AuthenticationGroupBy) Aggregate(fns ...AggregateFunc) *AuthenticationGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the group-by query and scans the result into the given value.
func (agb *AuthenticationGroupBy) Scan(ctx context.Context, v any) error {
	query, err := agb.path(ctx)
	if err != nil {
		return err
	}
	agb.sql = query
	return agb.sqlScan(ctx, v)
}

func (agb *AuthenticationGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range agb.fields {
		if !authentication.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := agb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (agb *AuthenticationGroupBy) sqlQuery() *sql.Selector {
	selector := agb.sql.Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(agb.fields)+len(agb.fns))
		for _, f := range agb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(agb.fields...)...)
}

// AuthenticationSelect is the builder for selecting fields of Authentication entities.
type AuthenticationSelect struct {
	*AuthenticationQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AuthenticationSelect) Aggregate(fns ...AggregateFunc) *AuthenticationSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AuthenticationSelect) Scan(ctx context.Context, v any) error {
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	as.sql = as.AuthenticationQuery.sqlQuery(ctx)
	return as.sqlScan(ctx, v)
}

func (as *AuthenticationSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(as.sql))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		as.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		as.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := as.sql.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (as *AuthenticationSelect) Modify(modifiers ...func(s *sql.Selector)) *AuthenticationSelect {
	as.modifiers = append(as.modifiers, modifiers...)
	return as
}
