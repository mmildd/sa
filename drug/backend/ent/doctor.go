// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/mmildd_s/app/ent/doctor"
)

// Doctor is the model entity for the Doctor schema.
type Doctor struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DoctorEmail holds the value of the "Doctor_Email" field.
	DoctorEmail string `json:"Doctor_Email,omitempty"`
	// DoctorPassword holds the value of the "Doctor_Password" field.
	DoctorPassword string `json:"Doctor_Password,omitempty"`
	// DoctorName holds the value of the "Doctor_Name" field.
	DoctorName string `json:"Doctor_Name,omitempty"`
	// DoctorTel holds the value of the "Doctor_Tel" field.
	DoctorTel string `json:"Doctor_Tel,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DoctorQuery when eager-loading is set.
	Edges DoctorEdges `json:"edges"`
}

// DoctorEdges holds the relations/edges for other nodes in the graph.
type DoctorEdges struct {
	// DoctorDrugAllergy holds the value of the Doctor_DrugAllergy edge.
	DoctorDrugAllergy []*DrugAllergy
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DoctorDrugAllergyOrErr returns the DoctorDrugAllergy value or an error if the edge
// was not loaded in eager-loading.
func (e DoctorEdges) DoctorDrugAllergyOrErr() ([]*DrugAllergy, error) {
	if e.loadedTypes[0] {
		return e.DoctorDrugAllergy, nil
	}
	return nil, &NotLoadedError{edge: "Doctor_DrugAllergy"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Doctor) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Doctor_Email
		&sql.NullString{}, // Doctor_Password
		&sql.NullString{}, // Doctor_Name
		&sql.NullString{}, // Doctor_Tel
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Doctor fields.
func (d *Doctor) assignValues(values ...interface{}) error {
	if m, n := len(values), len(doctor.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	d.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Doctor_Email", values[0])
	} else if value.Valid {
		d.DoctorEmail = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Doctor_Password", values[1])
	} else if value.Valid {
		d.DoctorPassword = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Doctor_Name", values[2])
	} else if value.Valid {
		d.DoctorName = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Doctor_Tel", values[3])
	} else if value.Valid {
		d.DoctorTel = value.String
	}
	return nil
}

// QueryDoctorDrugAllergy queries the Doctor_DrugAllergy edge of the Doctor.
func (d *Doctor) QueryDoctorDrugAllergy() *DrugAllergyQuery {
	return (&DoctorClient{config: d.config}).QueryDoctorDrugAllergy(d)
}

// Update returns a builder for updating this Doctor.
// Note that, you need to call Doctor.Unwrap() before calling this method, if this Doctor
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Doctor) Update() *DoctorUpdateOne {
	return (&DoctorClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (d *Doctor) Unwrap() *Doctor {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Doctor is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Doctor) String() string {
	var builder strings.Builder
	builder.WriteString("Doctor(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", Doctor_Email=")
	builder.WriteString(d.DoctorEmail)
	builder.WriteString(", Doctor_Password=")
	builder.WriteString(d.DoctorPassword)
	builder.WriteString(", Doctor_Name=")
	builder.WriteString(d.DoctorName)
	builder.WriteString(", Doctor_Tel=")
	builder.WriteString(d.DoctorTel)
	builder.WriteByte(')')
	return builder.String()
}

// Doctors is a parsable slice of Doctor.
type Doctors []*Doctor

func (d Doctors) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
