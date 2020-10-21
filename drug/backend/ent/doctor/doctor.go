// Code generated by entc, DO NOT EDIT.

package doctor

const (
	// Label holds the string label denoting the doctor type in the database.
	Label = "doctor"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDoctorEmail holds the string denoting the doctor_email field in the database.
	FieldDoctorEmail = "doctor_email"
	// FieldDoctorPassword holds the string denoting the doctor_password field in the database.
	FieldDoctorPassword = "doctor_password"
	// FieldDoctorName holds the string denoting the doctor_name field in the database.
	FieldDoctorName = "doctor_name"
	// FieldDoctorTel holds the string denoting the doctor_tel field in the database.
	FieldDoctorTel = "doctor_tel"

	// EdgeDoctorDrugAllergy holds the string denoting the doctor_drugallergy edge name in mutations.
	EdgeDoctorDrugAllergy = "Doctor_DrugAllergy"

	// Table holds the table name of the doctor in the database.
	Table = "doctors"
	// DoctorDrugAllergyTable is the table the holds the Doctor_DrugAllergy relation/edge.
	DoctorDrugAllergyTable = "drug_allergies"
	// DoctorDrugAllergyInverseTable is the table name for the DrugAllergy entity.
	// It exists in this package in order to avoid circular dependency with the "drugallergy" package.
	DoctorDrugAllergyInverseTable = "drug_allergies"
	// DoctorDrugAllergyColumn is the table column denoting the Doctor_DrugAllergy relation/edge.
	DoctorDrugAllergyColumn = "doctor_id"
)

// Columns holds all SQL columns for doctor fields.
var Columns = []string{
	FieldID,
	FieldDoctorEmail,
	FieldDoctorPassword,
	FieldDoctorName,
	FieldDoctorTel,
}

var (
	// DoctorEmailValidator is a validator for the "Doctor_Email" field. It is called by the builders before save.
	DoctorEmailValidator func(string) error
	// DoctorPasswordValidator is a validator for the "Doctor_Password" field. It is called by the builders before save.
	DoctorPasswordValidator func(string) error
	// DoctorNameValidator is a validator for the "Doctor_Name" field. It is called by the builders before save.
	DoctorNameValidator func(string) error
	// DoctorTelValidator is a validator for the "Doctor_Tel" field. It is called by the builders before save.
	DoctorTelValidator func(string) error
)