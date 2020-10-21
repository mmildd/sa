// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/mmildd_s/app/ent/medicine"
)

// Medicine is the model entity for the Medicine schema.
type Medicine struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// MedicineName holds the value of the "Medicine_Name" field.
	MedicineName string `json:"Medicine_Name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MedicineQuery when eager-loading is set.
	Edges MedicineEdges `json:"edges"`
}

// MedicineEdges holds the relations/edges for other nodes in the graph.
type MedicineEdges struct {
	// MedicineDrugAllergy holds the value of the Medicine_DrugAllergy edge.
	MedicineDrugAllergy []*DrugAllergy
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MedicineDrugAllergyOrErr returns the MedicineDrugAllergy value or an error if the edge
// was not loaded in eager-loading.
func (e MedicineEdges) MedicineDrugAllergyOrErr() ([]*DrugAllergy, error) {
	if e.loadedTypes[0] {
		return e.MedicineDrugAllergy, nil
	}
	return nil, &NotLoadedError{edge: "Medicine_DrugAllergy"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Medicine) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Medicine_Name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Medicine fields.
func (m *Medicine) assignValues(values ...interface{}) error {
	if m, n := len(values), len(medicine.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	m.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Medicine_Name", values[0])
	} else if value.Valid {
		m.MedicineName = value.String
	}
	return nil
}

// QueryMedicineDrugAllergy queries the Medicine_DrugAllergy edge of the Medicine.
func (m *Medicine) QueryMedicineDrugAllergy() *DrugAllergyQuery {
	return (&MedicineClient{config: m.config}).QueryMedicineDrugAllergy(m)
}

// Update returns a builder for updating this Medicine.
// Note that, you need to call Medicine.Unwrap() before calling this method, if this Medicine
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Medicine) Update() *MedicineUpdateOne {
	return (&MedicineClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (m *Medicine) Unwrap() *Medicine {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Medicine is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Medicine) String() string {
	var builder strings.Builder
	builder.WriteString("Medicine(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", Medicine_Name=")
	builder.WriteString(m.MedicineName)
	builder.WriteByte(')')
	return builder.String()
}

// Medicines is a parsable slice of Medicine.
type Medicines []*Medicine

func (m Medicines) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}