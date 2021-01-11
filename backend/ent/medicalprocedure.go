// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/team06/app/ent/doctor"
	"github.com/team06/app/ent/medicalprocedure"
	"github.com/team06/app/ent/patient"
	"github.com/team06/app/ent/proceduretype"
)

// MedicalProcedure is the model entity for the MedicalProcedure schema.
type MedicalProcedure struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Addtime holds the value of the "Addtime" field.
	Addtime time.Time `json:"Addtime,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MedicalProcedureQuery when eager-loading is set.
	Edges                                         MedicalProcedureEdges `json:"edges"`
	doctor_doctor_to_medical_procedure            *int
	patient_patient_to_medical_procedure          *int
	procedure_type_procedure_to_medical_procedure *int
}

// MedicalProcedureEdges holds the relations/edges for other nodes in the graph.
type MedicalProcedureEdges struct {
	// Patient holds the value of the Patient edge.
	Patient *Patient
	// ProcedureType holds the value of the ProcedureType edge.
	ProcedureType *ProcedureType
	// Doctor holds the value of the Doctor edge.
	Doctor *Doctor
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// PatientOrErr returns the Patient value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MedicalProcedureEdges) PatientOrErr() (*Patient, error) {
	if e.loadedTypes[0] {
		if e.Patient == nil {
			// The edge Patient was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: patient.Label}
		}
		return e.Patient, nil
	}
	return nil, &NotLoadedError{edge: "Patient"}
}

// ProcedureTypeOrErr returns the ProcedureType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MedicalProcedureEdges) ProcedureTypeOrErr() (*ProcedureType, error) {
	if e.loadedTypes[1] {
		if e.ProcedureType == nil {
			// The edge ProcedureType was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: proceduretype.Label}
		}
		return e.ProcedureType, nil
	}
	return nil, &NotLoadedError{edge: "ProcedureType"}
}

// DoctorOrErr returns the Doctor value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MedicalProcedureEdges) DoctorOrErr() (*Doctor, error) {
	if e.loadedTypes[2] {
		if e.Doctor == nil {
			// The edge Doctor was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: doctor.Label}
		}
		return e.Doctor, nil
	}
	return nil, &NotLoadedError{edge: "Doctor"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MedicalProcedure) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case medicalprocedure.FieldID:
			values[i] = &sql.NullInt64{}
		case medicalprocedure.FieldAddtime:
			values[i] = &sql.NullTime{}
		case medicalprocedure.ForeignKeys[0]: // doctor_doctor_to_medical_procedure
			values[i] = &sql.NullInt64{}
		case medicalprocedure.ForeignKeys[1]: // patient_patient_to_medical_procedure
			values[i] = &sql.NullInt64{}
		case medicalprocedure.ForeignKeys[2]: // procedure_type_procedure_to_medical_procedure
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type MedicalProcedure", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MedicalProcedure fields.
func (mp *MedicalProcedure) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case medicalprocedure.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mp.ID = int(value.Int64)
		case medicalprocedure.FieldAddtime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field Addtime", values[i])
			} else if value.Valid {
				mp.Addtime = value.Time
			}
		case medicalprocedure.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field doctor_doctor_to_medical_procedure", value)
			} else if value.Valid {
				mp.doctor_doctor_to_medical_procedure = new(int)
				*mp.doctor_doctor_to_medical_procedure = int(value.Int64)
			}
		case medicalprocedure.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field patient_patient_to_medical_procedure", value)
			} else if value.Valid {
				mp.patient_patient_to_medical_procedure = new(int)
				*mp.patient_patient_to_medical_procedure = int(value.Int64)
			}
		case medicalprocedure.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field procedure_type_procedure_to_medical_procedure", value)
			} else if value.Valid {
				mp.procedure_type_procedure_to_medical_procedure = new(int)
				*mp.procedure_type_procedure_to_medical_procedure = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryPatient queries the "Patient" edge of the MedicalProcedure entity.
func (mp *MedicalProcedure) QueryPatient() *PatientQuery {
	return (&MedicalProcedureClient{config: mp.config}).QueryPatient(mp)
}

// QueryProcedureType queries the "ProcedureType" edge of the MedicalProcedure entity.
func (mp *MedicalProcedure) QueryProcedureType() *ProcedureTypeQuery {
	return (&MedicalProcedureClient{config: mp.config}).QueryProcedureType(mp)
}

// QueryDoctor queries the "Doctor" edge of the MedicalProcedure entity.
func (mp *MedicalProcedure) QueryDoctor() *DoctorQuery {
	return (&MedicalProcedureClient{config: mp.config}).QueryDoctor(mp)
}

// Update returns a builder for updating this MedicalProcedure.
// Note that you need to call MedicalProcedure.Unwrap() before calling this method if this MedicalProcedure
// was returned from a transaction, and the transaction was committed or rolled back.
func (mp *MedicalProcedure) Update() *MedicalProcedureUpdateOne {
	return (&MedicalProcedureClient{config: mp.config}).UpdateOne(mp)
}

// Unwrap unwraps the MedicalProcedure entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mp *MedicalProcedure) Unwrap() *MedicalProcedure {
	tx, ok := mp.config.driver.(*txDriver)
	if !ok {
		panic("ent: MedicalProcedure is not a transactional entity")
	}
	mp.config.driver = tx.drv
	return mp
}

// String implements the fmt.Stringer.
func (mp *MedicalProcedure) String() string {
	var builder strings.Builder
	builder.WriteString("MedicalProcedure(")
	builder.WriteString(fmt.Sprintf("id=%v", mp.ID))
	builder.WriteString(", Addtime=")
	builder.WriteString(mp.Addtime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// MedicalProcedures is a parsable slice of MedicalProcedure.
type MedicalProcedures []*MedicalProcedure

func (mp MedicalProcedures) config(cfg config) {
	for _i := range mp {
		mp[_i].config = cfg
	}
}