// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/mmildd_s/app/ent/medicine"
	"github.com/mmildd_s/app/ent/predicate"
)

// MedicineDelete is the builder for deleting a Medicine entity.
type MedicineDelete struct {
	config
	hooks      []Hook
	mutation   *MedicineMutation
	predicates []predicate.Medicine
}

// Where adds a new predicate to the delete builder.
func (md *MedicineDelete) Where(ps ...predicate.Medicine) *MedicineDelete {
	md.predicates = append(md.predicates, ps...)
	return md
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (md *MedicineDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(md.hooks) == 0 {
		affected, err = md.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MedicineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			md.mutation = mutation
			affected, err = md.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(md.hooks) - 1; i >= 0; i-- {
			mut = md.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, md.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (md *MedicineDelete) ExecX(ctx context.Context) int {
	n, err := md.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (md *MedicineDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: medicine.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: medicine.FieldID,
			},
		},
	}
	if ps := md.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, md.driver, _spec)
}

// MedicineDeleteOne is the builder for deleting a single Medicine entity.
type MedicineDeleteOne struct {
	md *MedicineDelete
}

// Exec executes the deletion query.
func (mdo *MedicineDeleteOne) Exec(ctx context.Context) error {
	n, err := mdo.md.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{medicine.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mdo *MedicineDeleteOne) ExecX(ctx context.Context) {
	mdo.md.ExecX(ctx)
}
