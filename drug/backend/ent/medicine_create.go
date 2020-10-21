// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/mmildd_s/app/ent/drugallergy"
	"github.com/mmildd_s/app/ent/medicine"
)

// MedicineCreate is the builder for creating a Medicine entity.
type MedicineCreate struct {
	config
	mutation *MedicineMutation
	hooks    []Hook
}

// SetMedicineName sets the Medicine_Name field.
func (mc *MedicineCreate) SetMedicineName(s string) *MedicineCreate {
	mc.mutation.SetMedicineName(s)
	return mc
}

// AddMedicineDrugAllergyIDs adds the Medicine_DrugAllergy edge to DrugAllergy by ids.
func (mc *MedicineCreate) AddMedicineDrugAllergyIDs(ids ...int) *MedicineCreate {
	mc.mutation.AddMedicineDrugAllergyIDs(ids...)
	return mc
}

// AddMedicineDrugAllergy adds the Medicine_DrugAllergy edges to DrugAllergy.
func (mc *MedicineCreate) AddMedicineDrugAllergy(d ...*DrugAllergy) *MedicineCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return mc.AddMedicineDrugAllergyIDs(ids...)
}

// Mutation returns the MedicineMutation object of the builder.
func (mc *MedicineCreate) Mutation() *MedicineMutation {
	return mc.mutation
}

// Save creates the Medicine in the database.
func (mc *MedicineCreate) Save(ctx context.Context) (*Medicine, error) {
	if _, ok := mc.mutation.MedicineName(); !ok {
		return nil, &ValidationError{Name: "Medicine_Name", err: errors.New("ent: missing required field \"Medicine_Name\"")}
	}
	if v, ok := mc.mutation.MedicineName(); ok {
		if err := medicine.MedicineNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Medicine_Name", err: fmt.Errorf("ent: validator failed for field \"Medicine_Name\": %w", err)}
		}
	}
	var (
		err  error
		node *Medicine
	)
	if len(mc.hooks) == 0 {
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MedicineMutation)
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
func (mc *MedicineCreate) SaveX(ctx context.Context) *Medicine {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mc *MedicineCreate) sqlSave(ctx context.Context) (*Medicine, error) {
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

func (mc *MedicineCreate) createSpec() (*Medicine, *sqlgraph.CreateSpec) {
	var (
		m     = &Medicine{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: medicine.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: medicine.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.MedicineName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: medicine.FieldMedicineName,
		})
		m.MedicineName = value
	}
	if nodes := mc.mutation.MedicineDrugAllergyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicine.MedicineDrugAllergyTable,
			Columns: []string{medicine.MedicineDrugAllergyColumn},
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