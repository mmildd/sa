// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/mmildd_s/app/ent/drugallergy"
	"github.com/mmildd_s/app/ent/manner"
	"github.com/mmildd_s/app/ent/predicate"
)

// MannerUpdate is the builder for updating Manner entities.
type MannerUpdate struct {
	config
	hooks      []Hook
	mutation   *MannerMutation
	predicates []predicate.Manner
}

// Where adds a new predicate for the builder.
func (mu *MannerUpdate) Where(ps ...predicate.Manner) *MannerUpdate {
	mu.predicates = append(mu.predicates, ps...)
	return mu
}

// SetMannerName sets the Manner_Name field.
func (mu *MannerUpdate) SetMannerName(s string) *MannerUpdate {
	mu.mutation.SetMannerName(s)
	return mu
}

// AddMannerDrugAllergyIDs adds the Manner_DrugAllergy edge to DrugAllergy by ids.
func (mu *MannerUpdate) AddMannerDrugAllergyIDs(ids ...int) *MannerUpdate {
	mu.mutation.AddMannerDrugAllergyIDs(ids...)
	return mu
}

// AddMannerDrugAllergy adds the Manner_DrugAllergy edges to DrugAllergy.
func (mu *MannerUpdate) AddMannerDrugAllergy(d ...*DrugAllergy) *MannerUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return mu.AddMannerDrugAllergyIDs(ids...)
}

// Mutation returns the MannerMutation object of the builder.
func (mu *MannerUpdate) Mutation() *MannerMutation {
	return mu.mutation
}

// RemoveMannerDrugAllergyIDs removes the Manner_DrugAllergy edge to DrugAllergy by ids.
func (mu *MannerUpdate) RemoveMannerDrugAllergyIDs(ids ...int) *MannerUpdate {
	mu.mutation.RemoveMannerDrugAllergyIDs(ids...)
	return mu
}

// RemoveMannerDrugAllergy removes Manner_DrugAllergy edges to DrugAllergy.
func (mu *MannerUpdate) RemoveMannerDrugAllergy(d ...*DrugAllergy) *MannerUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return mu.RemoveMannerDrugAllergyIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (mu *MannerUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := mu.mutation.MannerName(); ok {
		if err := manner.MannerNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "Manner_Name", err: fmt.Errorf("ent: validator failed for field \"Manner_Name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MannerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MannerUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MannerUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MannerUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MannerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   manner.Table,
			Columns: manner.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: manner.FieldID,
			},
		},
	}
	if ps := mu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.MannerName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: manner.FieldMannerName,
		})
	}
	if nodes := mu.mutation.RemovedMannerDrugAllergyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   manner.MannerDrugAllergyTable,
			Columns: []string{manner.MannerDrugAllergyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: drugallergy.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MannerDrugAllergyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   manner.MannerDrugAllergyTable,
			Columns: []string{manner.MannerDrugAllergyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: drugallergy.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{manner.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// MannerUpdateOne is the builder for updating a single Manner entity.
type MannerUpdateOne struct {
	config
	hooks    []Hook
	mutation *MannerMutation
}

// SetMannerName sets the Manner_Name field.
func (muo *MannerUpdateOne) SetMannerName(s string) *MannerUpdateOne {
	muo.mutation.SetMannerName(s)
	return muo
}

// AddMannerDrugAllergyIDs adds the Manner_DrugAllergy edge to DrugAllergy by ids.
func (muo *MannerUpdateOne) AddMannerDrugAllergyIDs(ids ...int) *MannerUpdateOne {
	muo.mutation.AddMannerDrugAllergyIDs(ids...)
	return muo
}

// AddMannerDrugAllergy adds the Manner_DrugAllergy edges to DrugAllergy.
func (muo *MannerUpdateOne) AddMannerDrugAllergy(d ...*DrugAllergy) *MannerUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return muo.AddMannerDrugAllergyIDs(ids...)
}

// Mutation returns the MannerMutation object of the builder.
func (muo *MannerUpdateOne) Mutation() *MannerMutation {
	return muo.mutation
}

// RemoveMannerDrugAllergyIDs removes the Manner_DrugAllergy edge to DrugAllergy by ids.
func (muo *MannerUpdateOne) RemoveMannerDrugAllergyIDs(ids ...int) *MannerUpdateOne {
	muo.mutation.RemoveMannerDrugAllergyIDs(ids...)
	return muo
}

// RemoveMannerDrugAllergy removes Manner_DrugAllergy edges to DrugAllergy.
func (muo *MannerUpdateOne) RemoveMannerDrugAllergy(d ...*DrugAllergy) *MannerUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return muo.RemoveMannerDrugAllergyIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (muo *MannerUpdateOne) Save(ctx context.Context) (*Manner, error) {
	if v, ok := muo.mutation.MannerName(); ok {
		if err := manner.MannerNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Manner_Name", err: fmt.Errorf("ent: validator failed for field \"Manner_Name\": %w", err)}
		}
	}

	var (
		err  error
		node *Manner
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MannerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MannerUpdateOne) SaveX(ctx context.Context) *Manner {
	m, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return m
}

// Exec executes the query on the entity.
func (muo *MannerUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MannerUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MannerUpdateOne) sqlSave(ctx context.Context) (m *Manner, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   manner.Table,
			Columns: manner.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: manner.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Manner.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := muo.mutation.MannerName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: manner.FieldMannerName,
		})
	}
	if nodes := muo.mutation.RemovedMannerDrugAllergyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   manner.MannerDrugAllergyTable,
			Columns: []string{manner.MannerDrugAllergyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: drugallergy.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MannerDrugAllergyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   manner.MannerDrugAllergyTable,
			Columns: []string{manner.MannerDrugAllergyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: drugallergy.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	m = &Manner{config: muo.config}
	_spec.Assign = m.assignValues
	_spec.ScanValues = m.scanValues()
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{manner.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return m, nil
}
