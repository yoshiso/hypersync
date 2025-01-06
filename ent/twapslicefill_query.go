// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yoshiso/hypersync/ent/predicate"
	"github.com/yoshiso/hypersync/ent/twapslicefill"
)

// TwapSliceFillQuery is the builder for querying TwapSliceFill entities.
type TwapSliceFillQuery struct {
	config
	ctx        *QueryContext
	order      []twapslicefill.OrderOption
	inters     []Interceptor
	predicates []predicate.TwapSliceFill
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TwapSliceFillQuery builder.
func (tsfq *TwapSliceFillQuery) Where(ps ...predicate.TwapSliceFill) *TwapSliceFillQuery {
	tsfq.predicates = append(tsfq.predicates, ps...)
	return tsfq
}

// Limit the number of records to be returned by this query.
func (tsfq *TwapSliceFillQuery) Limit(limit int) *TwapSliceFillQuery {
	tsfq.ctx.Limit = &limit
	return tsfq
}

// Offset to start from.
func (tsfq *TwapSliceFillQuery) Offset(offset int) *TwapSliceFillQuery {
	tsfq.ctx.Offset = &offset
	return tsfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tsfq *TwapSliceFillQuery) Unique(unique bool) *TwapSliceFillQuery {
	tsfq.ctx.Unique = &unique
	return tsfq
}

// Order specifies how the records should be ordered.
func (tsfq *TwapSliceFillQuery) Order(o ...twapslicefill.OrderOption) *TwapSliceFillQuery {
	tsfq.order = append(tsfq.order, o...)
	return tsfq
}

// First returns the first TwapSliceFill entity from the query.
// Returns a *NotFoundError when no TwapSliceFill was found.
func (tsfq *TwapSliceFillQuery) First(ctx context.Context) (*TwapSliceFill, error) {
	nodes, err := tsfq.Limit(1).All(setContextOp(ctx, tsfq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{twapslicefill.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tsfq *TwapSliceFillQuery) FirstX(ctx context.Context) *TwapSliceFill {
	node, err := tsfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TwapSliceFill ID from the query.
// Returns a *NotFoundError when no TwapSliceFill ID was found.
func (tsfq *TwapSliceFillQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tsfq.Limit(1).IDs(setContextOp(ctx, tsfq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{twapslicefill.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tsfq *TwapSliceFillQuery) FirstIDX(ctx context.Context) int {
	id, err := tsfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TwapSliceFill entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TwapSliceFill entity is found.
// Returns a *NotFoundError when no TwapSliceFill entities are found.
func (tsfq *TwapSliceFillQuery) Only(ctx context.Context) (*TwapSliceFill, error) {
	nodes, err := tsfq.Limit(2).All(setContextOp(ctx, tsfq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{twapslicefill.Label}
	default:
		return nil, &NotSingularError{twapslicefill.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tsfq *TwapSliceFillQuery) OnlyX(ctx context.Context) *TwapSliceFill {
	node, err := tsfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TwapSliceFill ID in the query.
// Returns a *NotSingularError when more than one TwapSliceFill ID is found.
// Returns a *NotFoundError when no entities are found.
func (tsfq *TwapSliceFillQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tsfq.Limit(2).IDs(setContextOp(ctx, tsfq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{twapslicefill.Label}
	default:
		err = &NotSingularError{twapslicefill.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tsfq *TwapSliceFillQuery) OnlyIDX(ctx context.Context) int {
	id, err := tsfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TwapSliceFills.
func (tsfq *TwapSliceFillQuery) All(ctx context.Context) ([]*TwapSliceFill, error) {
	ctx = setContextOp(ctx, tsfq.ctx, ent.OpQueryAll)
	if err := tsfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TwapSliceFill, *TwapSliceFillQuery]()
	return withInterceptors[[]*TwapSliceFill](ctx, tsfq, qr, tsfq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tsfq *TwapSliceFillQuery) AllX(ctx context.Context) []*TwapSliceFill {
	nodes, err := tsfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TwapSliceFill IDs.
func (tsfq *TwapSliceFillQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tsfq.ctx.Unique == nil && tsfq.path != nil {
		tsfq.Unique(true)
	}
	ctx = setContextOp(ctx, tsfq.ctx, ent.OpQueryIDs)
	if err = tsfq.Select(twapslicefill.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tsfq *TwapSliceFillQuery) IDsX(ctx context.Context) []int {
	ids, err := tsfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tsfq *TwapSliceFillQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tsfq.ctx, ent.OpQueryCount)
	if err := tsfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tsfq, querierCount[*TwapSliceFillQuery](), tsfq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tsfq *TwapSliceFillQuery) CountX(ctx context.Context) int {
	count, err := tsfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tsfq *TwapSliceFillQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tsfq.ctx, ent.OpQueryExist)
	switch _, err := tsfq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tsfq *TwapSliceFillQuery) ExistX(ctx context.Context) bool {
	exist, err := tsfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TwapSliceFillQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tsfq *TwapSliceFillQuery) Clone() *TwapSliceFillQuery {
	if tsfq == nil {
		return nil
	}
	return &TwapSliceFillQuery{
		config:     tsfq.config,
		ctx:        tsfq.ctx.Clone(),
		order:      append([]twapslicefill.OrderOption{}, tsfq.order...),
		inters:     append([]Interceptor{}, tsfq.inters...),
		predicates: append([]predicate.TwapSliceFill{}, tsfq.predicates...),
		// clone intermediate query.
		sql:  tsfq.sql.Clone(),
		path: tsfq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Coin string `json:"coin,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TwapSliceFill.Query().
//		GroupBy(twapslicefill.FieldCoin).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tsfq *TwapSliceFillQuery) GroupBy(field string, fields ...string) *TwapSliceFillGroupBy {
	tsfq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TwapSliceFillGroupBy{build: tsfq}
	grbuild.flds = &tsfq.ctx.Fields
	grbuild.label = twapslicefill.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Coin string `json:"coin,omitempty"`
//	}
//
//	client.TwapSliceFill.Query().
//		Select(twapslicefill.FieldCoin).
//		Scan(ctx, &v)
func (tsfq *TwapSliceFillQuery) Select(fields ...string) *TwapSliceFillSelect {
	tsfq.ctx.Fields = append(tsfq.ctx.Fields, fields...)
	sbuild := &TwapSliceFillSelect{TwapSliceFillQuery: tsfq}
	sbuild.label = twapslicefill.Label
	sbuild.flds, sbuild.scan = &tsfq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TwapSliceFillSelect configured with the given aggregations.
func (tsfq *TwapSliceFillQuery) Aggregate(fns ...AggregateFunc) *TwapSliceFillSelect {
	return tsfq.Select().Aggregate(fns...)
}

func (tsfq *TwapSliceFillQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tsfq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tsfq); err != nil {
				return err
			}
		}
	}
	for _, f := range tsfq.ctx.Fields {
		if !twapslicefill.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tsfq.path != nil {
		prev, err := tsfq.path(ctx)
		if err != nil {
			return err
		}
		tsfq.sql = prev
	}
	return nil
}

func (tsfq *TwapSliceFillQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TwapSliceFill, error) {
	var (
		nodes = []*TwapSliceFill{}
		_spec = tsfq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TwapSliceFill).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TwapSliceFill{config: tsfq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tsfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tsfq *TwapSliceFillQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tsfq.querySpec()
	_spec.Node.Columns = tsfq.ctx.Fields
	if len(tsfq.ctx.Fields) > 0 {
		_spec.Unique = tsfq.ctx.Unique != nil && *tsfq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tsfq.driver, _spec)
}

func (tsfq *TwapSliceFillQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(twapslicefill.Table, twapslicefill.Columns, sqlgraph.NewFieldSpec(twapslicefill.FieldID, field.TypeInt))
	_spec.From = tsfq.sql
	if unique := tsfq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tsfq.path != nil {
		_spec.Unique = true
	}
	if fields := tsfq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, twapslicefill.FieldID)
		for i := range fields {
			if fields[i] != twapslicefill.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tsfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tsfq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tsfq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tsfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tsfq *TwapSliceFillQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tsfq.driver.Dialect())
	t1 := builder.Table(twapslicefill.Table)
	columns := tsfq.ctx.Fields
	if len(columns) == 0 {
		columns = twapslicefill.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tsfq.sql != nil {
		selector = tsfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tsfq.ctx.Unique != nil && *tsfq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tsfq.predicates {
		p(selector)
	}
	for _, p := range tsfq.order {
		p(selector)
	}
	if offset := tsfq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tsfq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TwapSliceFillGroupBy is the group-by builder for TwapSliceFill entities.
type TwapSliceFillGroupBy struct {
	selector
	build *TwapSliceFillQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tsfgb *TwapSliceFillGroupBy) Aggregate(fns ...AggregateFunc) *TwapSliceFillGroupBy {
	tsfgb.fns = append(tsfgb.fns, fns...)
	return tsfgb
}

// Scan applies the selector query and scans the result into the given value.
func (tsfgb *TwapSliceFillGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tsfgb.build.ctx, ent.OpQueryGroupBy)
	if err := tsfgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TwapSliceFillQuery, *TwapSliceFillGroupBy](ctx, tsfgb.build, tsfgb, tsfgb.build.inters, v)
}

func (tsfgb *TwapSliceFillGroupBy) sqlScan(ctx context.Context, root *TwapSliceFillQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tsfgb.fns))
	for _, fn := range tsfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tsfgb.flds)+len(tsfgb.fns))
		for _, f := range *tsfgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tsfgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tsfgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TwapSliceFillSelect is the builder for selecting fields of TwapSliceFill entities.
type TwapSliceFillSelect struct {
	*TwapSliceFillQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tsfs *TwapSliceFillSelect) Aggregate(fns ...AggregateFunc) *TwapSliceFillSelect {
	tsfs.fns = append(tsfs.fns, fns...)
	return tsfs
}

// Scan applies the selector query and scans the result into the given value.
func (tsfs *TwapSliceFillSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tsfs.ctx, ent.OpQuerySelect)
	if err := tsfs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TwapSliceFillQuery, *TwapSliceFillSelect](ctx, tsfs.TwapSliceFillQuery, tsfs, tsfs.inters, v)
}

func (tsfs *TwapSliceFillSelect) sqlScan(ctx context.Context, root *TwapSliceFillQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tsfs.fns))
	for _, fn := range tsfs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tsfs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tsfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
