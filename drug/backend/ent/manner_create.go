// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/mmildd_s/app/ent/drugallergy"
	"github.com/mmildd_s/app/ent/manner"
)

// MannerCreate is the builder for creating a Manner entity.
type MannerCreate struct {
	config
	mutation *MannerMutation
	hooks    []Hook
}

// SetMannerName sets the Manner_Name field.
func (mc *MannerCreate) SetMannerName(s string) *MannerCreate {
	mc.mutation.SetMannerName(s)
	return mc
}

// AddMannerDrugAllergyIDs adds the Manner_DrugAllergy edge to DrugAllergy by ids.
func (mc *MannerCreate) AddMannerDrugAllergyIDs(ids ...int) *MannerCreate {
	mc.mutation.AddMannerDrugAllergyIDs(ids...)
	return mc
}

// AddMannerDrugAllergy adds the Manner_DrugAllergy edges to DrugAllergy.
func (mc *MannerCreate) AddMannerDrugAllergy(d ...*DrugAllergy) *MannerCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return mc.AddMannerDrugAllergyIDs(ids...)
}

// Mutation returns the MannerMutation object of the builder.
func (mc *MannerCreate) Mutation() *MannerMutation {
	return mc.mutation
}

// Save creates the Manner in the database.
func (mc *MannerCreate) Save(ctx context.Context) (*Manner, error) {
	if _, ok := mc.mutation.MannerName(); !ok {
		return nil, &ValidationError{Name: "Manner_Name", err: errors.New("ent: missing required field \"Manner_Name\"")}
	}
	if v, ok := mc.mutation.MannerName(); ok {
		if err := manner.MannerNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Manner_Name", err: fmt.Errorf("ent: validator failed for field \"Manner_Name\": %w", err)}
		}
	}
	var (
		err  error
		node *Manner
	)
	if len(mc.hooks) == 0 {
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MannerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mc.mutation = mutation
			node, err = mc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MannerCreate) SaveX(ctx context.Context) *Manner {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mc *MannerCreate) sqlSave(ctx context.Context) (*Manner, error) {
	m, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	m.ID = int(id)
	return m, nil
}

func (mc *MannerCreate) createSpec() (*Manner, *sqlgraph.CreateSpec) {
	var (
		m     = &Manner{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: manner.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: manner.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.MannerName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: manner.FieldMannerName,
		})
		m.MannerName = value
	}
	if nodes := mc.mutation.MannerDrugAllergyIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return m, _spec
}