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
	"github.com/yoshiso/hypersync/ent/vaultdelta"
)

// VaultDeltaQuery is the builder for querying VaultDelta entities.
type VaultDeltaQuery struct {
	config
	ctx        *QueryContext
	order      []vaultdelta.OrderOption
	inters     []Interceptor
	predicates []predicate.VaultDelta
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VaultDeltaQuery builder.
func (vdq *VaultDeltaQuery) Where(ps ...predicate.VaultDelta) *VaultDeltaQuery {
	vdq.predicates = append(vdq.predicates, ps...)
	return vdq
}

// Limit the number of records to be returned by this query.
func (vdq *VaultDeltaQuery) Limit(limit int) *VaultDeltaQuery {
	vdq.ctx.Limit = &limit
	return vdq
}

// Offset to start from.
func (vdq *VaultDeltaQuery) Offset(offset int) *VaultDeltaQuery {
	vdq.ctx.Offset = &offset
	return vdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vdq *VaultDeltaQuery) Unique(unique bool) *VaultDeltaQuery {
	vdq.ctx.Unique = &unique
	return vdq
}

// Order specifies how the records should be ordered.
func (vdq *VaultDeltaQuery) Order(o ...vaultdelta.OrderOption) *VaultDeltaQuery {
	vdq.order = append(vdq.order, o...)
	return vdq
}

// First returns the first VaultDelta entity from the query.
// Returns a *NotFoundError when no VaultDelta was found.
func (vdq *VaultDeltaQuery) First(ctx context.Context) (*VaultDelta, error) {
	nodes, err := vdq.Limit(1).All(setContextOp(ctx, vdq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{vaultdelta.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vdq *VaultDeltaQuery) FirstX(ctx context.Context) *VaultDelta {
	node, err := vdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VaultDelta ID from the query.
// Returns a *NotFoundError when no VaultDelta ID was found.
func (vdq *VaultDeltaQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vdq.Limit(1).IDs(setContextOp(ctx, vdq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{vaultdelta.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vdq *VaultDeltaQuery) FirstIDX(ctx context.Context) int {
	id, err := vdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VaultDelta entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one VaultDelta entity is found.
// Returns a *NotFoundError when no VaultDelta entities are found.
func (vdq *VaultDeltaQuery) Only(ctx context.Context) (*VaultDelta, error) {
	nodes, err := vdq.Limit(2).All(setContextOp(ctx, vdq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{vaultdelta.Label}
	default:
		return nil, &NotSingularError{vaultdelta.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vdq *VaultDeltaQuery) OnlyX(ctx context.Context) *VaultDelta {
	node, err := vdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VaultDelta ID in the query.
// Returns a *NotSingularError when more than one VaultDelta ID is found.
// Returns a *NotFoundError when no entities are found.
func (vdq *VaultDeltaQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vdq.Limit(2).IDs(setContextOp(ctx, vdq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{vaultdelta.Label}
	default:
		err = &NotSingularError{vaultdelta.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vdq *VaultDeltaQuery) OnlyIDX(ctx context.Context) int {
	id, err := vdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VaultDeltaSlice.
func (vdq *VaultDeltaQuery) All(ctx context.Context) ([]*VaultDelta, error) {
	ctx = setContextOp(ctx, vdq.ctx, ent.OpQueryAll)
	if err := vdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*VaultDelta, *VaultDeltaQuery]()
	return withInterceptors[[]*VaultDelta](ctx, vdq, qr, vdq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vdq *VaultDeltaQuery) AllX(ctx context.Context) []*VaultDelta {
	nodes, err := vdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VaultDelta IDs.
func (vdq *VaultDeltaQuery) IDs(ctx context.Context) (ids []int, err error) {
	if vdq.ctx.Unique == nil && vdq.path != nil {
		vdq.Unique(true)
	}
	ctx = setContextOp(ctx, vdq.ctx, ent.OpQueryIDs)
	if err = vdq.Select(vaultdelta.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vdq *VaultDeltaQuery) IDsX(ctx context.Context) []int {
	ids, err := vdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vdq *VaultDeltaQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vdq.ctx, ent.OpQueryCount)
	if err := vdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vdq, querierCount[*VaultDeltaQuery](), vdq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vdq *VaultDeltaQuery) CountX(ctx context.Context) int {
	count, err := vdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vdq *VaultDeltaQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vdq.ctx, ent.OpQueryExist)
	switch _, err := vdq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vdq *VaultDeltaQuery) ExistX(ctx context.Context) bool {
	exist, err := vdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VaultDeltaQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vdq *VaultDeltaQuery) Clone() *VaultDeltaQuery {
	if vdq == nil {
		return nil
	}
	return &VaultDeltaQuery{
		config:     vdq.config,
		ctx:        vdq.ctx.Clone(),
		order:      append([]vaultdelta.OrderOption{}, vdq.order...),
		inters:     append([]Interceptor{}, vdq.inters...),
		predicates: append([]predicate.VaultDelta{}, vdq.predicates...),
		// clone intermediate query.
		sql:  vdq.sql.Clone(),
		path: vdq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Type string `json:"type,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.VaultDelta.Query().
//		GroupBy(vaultdelta.FieldType).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vdq *VaultDeltaQuery) GroupBy(field string, fields ...string) *VaultDeltaGroupBy {
	vdq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VaultDeltaGroupBy{build: vdq}
	grbuild.flds = &vdq.ctx.Fields
	grbuild.label = vaultdelta.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Type string `json:"type,omitempty"`
//	}
//
//	client.VaultDelta.Query().
//		Select(vaultdelta.FieldType).
//		Scan(ctx, &v)
func (vdq *VaultDeltaQuery) Select(fields ...string) *VaultDeltaSelect {
	vdq.ctx.Fields = append(vdq.ctx.Fields, fields...)
	sbuild := &VaultDeltaSelect{VaultDeltaQuery: vdq}
	sbuild.label = vaultdelta.Label
	sbuild.flds, sbuild.scan = &vdq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VaultDeltaSelect configured with the given aggregations.
func (vdq *VaultDeltaQuery) Aggregate(fns ...AggregateFunc) *VaultDeltaSelect {
	return vdq.Select().Aggregate(fns...)
}

func (vdq *VaultDeltaQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vdq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vdq); err != nil {
				return err
			}
		}
	}
	for _, f := range vdq.ctx.Fields {
		if !vaultdelta.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vdq.path != nil {
		prev, err := vdq.path(ctx)
		if err != nil {
			return err
		}
		vdq.sql = prev
	}
	return nil
}

func (vdq *VaultDeltaQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*VaultDelta, error) {
	var (
		nodes = []*VaultDelta{}
		_spec = vdq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*VaultDelta).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &VaultDelta{config: vdq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (vdq *VaultDeltaQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vdq.querySpec()
	_spec.Node.Columns = vdq.ctx.Fields
	if len(vdq.ctx.Fields) > 0 {
		_spec.Unique = vdq.ctx.Unique != nil && *vdq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vdq.driver, _spec)
}

func (vdq *VaultDeltaQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(vaultdelta.Table, vaultdelta.Columns, sqlgraph.NewFieldSpec(vaultdelta.FieldID, field.TypeInt))
	_spec.From = vdq.sql
	if unique := vdq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vdq.path != nil {
		_spec.Unique = true
	}
	if fields := vdq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vaultdelta.FieldID)
		for i := range fields {
			if fields[i] != vaultdelta.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vdq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vdq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vdq *VaultDeltaQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vdq.driver.Dialect())
	t1 := builder.Table(vaultdelta.Table)
	columns := vdq.ctx.Fields
	if len(columns) == 0 {
		columns = vaultdelta.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vdq.sql != nil {
		selector = vdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vdq.ctx.Unique != nil && *vdq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range vdq.predicates {
		p(selector)
	}
	for _, p := range vdq.order {
		p(selector)
	}
	if offset := vdq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vdq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VaultDeltaGroupBy is the group-by builder for VaultDelta entities.
type VaultDeltaGroupBy struct {
	selector
	build *VaultDeltaQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vdgb *VaultDeltaGroupBy) Aggregate(fns ...AggregateFunc) *VaultDeltaGroupBy {
	vdgb.fns = append(vdgb.fns, fns...)
	return vdgb
}

// Scan applies the selector query and scans the result into the given value.
func (vdgb *VaultDeltaGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vdgb.build.ctx, ent.OpQueryGroupBy)
	if err := vdgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VaultDeltaQuery, *VaultDeltaGroupBy](ctx, vdgb.build, vdgb, vdgb.build.inters, v)
}

func (vdgb *VaultDeltaGroupBy) sqlScan(ctx context.Context, root *VaultDeltaQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vdgb.fns))
	for _, fn := range vdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vdgb.flds)+len(vdgb.fns))
		for _, f := range *vdgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vdgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vdgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VaultDeltaSelect is the builder for selecting fields of VaultDelta entities.
type VaultDeltaSelect struct {
	*VaultDeltaQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vds *VaultDeltaSelect) Aggregate(fns ...AggregateFunc) *VaultDeltaSelect {
	vds.fns = append(vds.fns, fns...)
	return vds
}

// Scan applies the selector query and scans the result into the given value.
func (vds *VaultDeltaSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vds.ctx, ent.OpQuerySelect)
	if err := vds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VaultDeltaQuery, *VaultDeltaSelect](ctx, vds.VaultDeltaQuery, vds, vds.inters, v)
}

func (vds *VaultDeltaSelect) sqlScan(ctx context.Context, root *VaultDeltaQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vds.fns))
	for _, fn := range vds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
