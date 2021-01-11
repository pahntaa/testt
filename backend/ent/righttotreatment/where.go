// Code generated by entc, DO NOT EDIT.

package righttotreatment

import (
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/team06/app/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Addedtime applies equality check predicate on the "Addedtime" field. It's identical to AddedtimeEQ.
func Addedtime(v time.Time) predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddedtime), v))
	})
}

// AddedtimeEQ applies the EQ predicate on the "Addedtime" field.
func AddedtimeEQ(v time.Time) predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddedtime), v))
	})
}

// AddedtimeNEQ applies the NEQ predicate on the "Addedtime" field.
func AddedtimeNEQ(v time.Time) predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddedtime), v))
	})
}

// AddedtimeIn applies the In predicate on the "Addedtime" field.
func AddedtimeIn(vs ...time.Time) predicate.RightToTreatment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RightToTreatment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAddedtime), v...))
	})
}

// AddedtimeNotIn applies the NotIn predicate on the "Addedtime" field.
func AddedtimeNotIn(vs ...time.Time) predicate.RightToTreatment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RightToTreatment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAddedtime), v...))
	})
}

// AddedtimeGT applies the GT predicate on the "Addedtime" field.
func AddedtimeGT(v time.Time) predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddedtime), v))
	})
}

// AddedtimeGTE applies the GTE predicate on the "Addedtime" field.
func AddedtimeGTE(v time.Time) predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddedtime), v))
	})
}

// AddedtimeLT applies the LT predicate on the "Addedtime" field.
func AddedtimeLT(v time.Time) predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddedtime), v))
	})
}

// AddedtimeLTE applies the LTE predicate on the "Addedtime" field.
func AddedtimeLTE(v time.Time) predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddedtime), v))
	})
}

// HasHospital applies the HasEdge predicate on the "Hospital" edge.
func HasHospital() predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HospitalTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HospitalTable, HospitalColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHospitalWith applies the HasEdge predicate on the "Hospital" edge with a given conditions (other predicates).
func HasHospitalWith(preds ...predicate.Hospital) predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HospitalInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HospitalTable, HospitalColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRightToTreatmentType applies the HasEdge predicate on the "RightToTreatmentType" edge.
func HasRightToTreatmentType() predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RightToTreatmentTypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RightToTreatmentTypeTable, RightToTreatmentTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRightToTreatmentTypeWith applies the HasEdge predicate on the "RightToTreatmentType" edge with a given conditions (other predicates).
func HasRightToTreatmentTypeWith(preds ...predicate.RightToTreatmentType) predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RightToTreatmentTypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RightToTreatmentTypeTable, RightToTreatmentTypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPatient applies the HasEdge predicate on the "Patient" edge.
func HasPatient() predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientWith applies the HasEdge predicate on the "Patient" edge with a given conditions (other predicates).
func HasPatientWith(preds ...predicate.Patient) predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RightToTreatment) predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RightToTreatment) predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RightToTreatment) predicate.RightToTreatment {
	return predicate.RightToTreatment(func(s *sql.Selector) {
		p(s.Not())
	})
}