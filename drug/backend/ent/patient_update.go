// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/mmildd_s/app/ent/drugallergy"
	"github.com/mmildd_s/app/ent/patient"
	"github.com/mmildd_s/app/ent/predicate"
)

// PatientUpdate is the builder for updating Patient entities.
type PatientUpdate struct {
	config
	hooks      []Hook
	mutation   *PatientMutation
	predicates []predicate.Patient
}

// Where adds a new predicate for the builder.
func (pu *PatientUpdate) Where(ps ...predicate.Patient) *PatientUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetPatientName sets the Patient_Name field.
func (pu *PatientUpdate) SetPatientName(s string) *PatientUpdate {
	pu.mutation.SetPatientName(s)
	return pu
}

// AddPatientDrugAllergyIDs adds the Patient_DrugAllergy edge to DrugAllergy by ids.
func (pu *PatientUpdate) AddPatientDrugAllergyIDs(ids ...int) *PatientUpdate {
	pu.mutation.AddPatientDrugAllergyIDs(ids...)
	return pu
}

// AddPatientDrugAllergy adds the Patient_DrugAllergy edges to DrugAllergy.
func (pu *PatientUpdate) AddPatientDrugAllergy(d ...*DrugAllergy) *PatientUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pu.AddPatientDrugAllergyIDs(ids...)
}

// Mutation returns the PatientMutation object of the builder.
func (pu *PatientUpdate) Mutation() *PatientMutation {
	return pu.mutation
}

// RemovePatientDrugAllergyIDs removes the Patient_DrugAllergy edge to DrugAllergy by ids.
func (pu *PatientUpdate) RemovePatientDrugAllergyIDs(ids ...int) *PatientUpdate {
	pu.mutation.RemovePatientDrugAllergyIDs(ids...)
	return pu
}

// RemovePatientDrugAllergy removes Patient_DrugAllergy edges to DrugAllergy.
func (pu *PatientUpdate) RemovePatientDrugAllergy(d ...*DrugAllergy) *PatientUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pu.RemovePatientDrugAllergyIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PatientUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := pu.mutation.PatientName(); ok {
		if err := patient.PatientNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "Patient_Name", err: fmt.Errorf("ent: validator failed for field \"Patient_Name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PatientUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PatientUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PatientUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PatientUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patient.Table,
			Columns: patient.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patient.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.PatientName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldPatientName,
		})
	}
	if nodes := pu.mutation.RemovedPatientDrugAllergyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.PatientDrugAllergyTable,
			Columns: []string{patient.PatientDrugAllergyColumn},
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
	if nodes := pu.mutation.PatientDrugAllergyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.PatientDrugAllergyTable,
			Columns: []string{patient.PatientDrugAllergyColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{patient.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PatientUpdateOne is the builder for updating a single Patient entity.
type PatientUpdateOne struct {
	config
	hooks    []Hook
	mutation *PatientMutation
}

// SetPatientName sets the Patient_Name field.
func (puo *PatientUpdateOne) SetPatientName(s string) *PatientUpdateOne {
	puo.mutation.SetPatientName(s)
	return puo
}

// AddPatientDrugAllergyIDs adds the Patient_DrugAllergy edge to DrugAllergy by ids.
func (puo *PatientUpdateOne) AddPatientDrugAllergyIDs(ids ...int) *PatientUpdateOne {
	puo.mutation.AddPatientDrugAllergyIDs(ids...)
	return puo
}

// AddPatientDrugAllergy adds the Patient_DrugAllergy edges to DrugAllergy.
func (puo *PatientUpdateOne) AddPatientDrugAllergy(d ...*DrugAllergy) *PatientUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return puo.AddPatientDrugAllergyIDs(ids...)
}

// Mutation returns the PatientMutation object of the builder.
func (puo *PatientUpdateOne) Mutation() *PatientMutation {
	return puo.mutation
}

// RemovePatientDrugAllergyIDs removes the Patient_DrugAllergy edge to DrugAllergy by ids.
func (puo *PatientUpdateOne) RemovePatientDrugAllergyIDs(ids ...int) *PatientUpdateOne {
	puo.mutation.RemovePatientDrugAllergyIDs(ids...)
	return puo
}

// RemovePatientDrugAllergy removes Patient_DrugAllergy edges to DrugAllergy.
func (puo *PatientUpdateOne) RemovePatientDrugAllergy(d ...*DrugAllergy) *PatientUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return puo.RemovePatientDrugAllergyIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *PatientUpdateOne) Save(ctx context.Context) (*Patient, error) {
	if v, ok := puo.mutation.PatientName(); ok {
		if err := patient.PatientNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Patient_Name", err: fmt.Errorf("ent: validator failed for field \"Patient_Name\": %w", err)}
		}
	}

	var (
		err  error
		node *Patient
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PatientUpdateOne) SaveX(ctx context.Context) *Patient {
	pa, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pa
}

// Exec executes the query on the entity.
func (puo *PatientUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PatientUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PatientUpdateOne) sqlSave(ctx context.Context) (pa *Patient, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patient.Table,
			Columns: patient.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patient.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Patient.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.PatientName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldPatientName,
		})
	}
	if nodes := puo.mutation.RemovedPatientDrugAllergyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.PatientDrugAllergyTable,
			Columns: []string{patient.PatientDrugAllergyColumn},
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
	if nodes := puo.mutation.PatientDrugAllergyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.PatientDrugAllergyTable,
			Columns: []string{patient.PatientDrugAllergyColumn},
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
	pa = &Patient{config: puo.config}
	_spec.Assign = pa.assignValues
	_spec.ScanValues = pa.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{patient.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pa, nil
}