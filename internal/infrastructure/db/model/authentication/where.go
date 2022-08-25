// Code generated by ent, DO NOT EDIT.

package authentication

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Southclaws/storyden/internal/infrastructure/db/model/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// Service applies equality check predicate on the "service" field. It's identical to ServiceEQ.
func Service(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldService), v))
	})
}

// Identifier applies equality check predicate on the "identifier" field. It's identical to IdentifierEQ.
func Identifier(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIdentifier), v))
	})
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToken), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Authentication {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Authentication {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// ServiceEQ applies the EQ predicate on the "service" field.
func ServiceEQ(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldService), v))
	})
}

// ServiceNEQ applies the NEQ predicate on the "service" field.
func ServiceNEQ(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldService), v))
	})
}

// ServiceIn applies the In predicate on the "service" field.
func ServiceIn(vs ...string) predicate.Authentication {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldService), v...))
	})
}

// ServiceNotIn applies the NotIn predicate on the "service" field.
func ServiceNotIn(vs ...string) predicate.Authentication {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldService), v...))
	})
}

// ServiceGT applies the GT predicate on the "service" field.
func ServiceGT(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldService), v))
	})
}

// ServiceGTE applies the GTE predicate on the "service" field.
func ServiceGTE(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldService), v))
	})
}

// ServiceLT applies the LT predicate on the "service" field.
func ServiceLT(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldService), v))
	})
}

// ServiceLTE applies the LTE predicate on the "service" field.
func ServiceLTE(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldService), v))
	})
}

// ServiceContains applies the Contains predicate on the "service" field.
func ServiceContains(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldService), v))
	})
}

// ServiceHasPrefix applies the HasPrefix predicate on the "service" field.
func ServiceHasPrefix(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldService), v))
	})
}

// ServiceHasSuffix applies the HasSuffix predicate on the "service" field.
func ServiceHasSuffix(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldService), v))
	})
}

// ServiceEqualFold applies the EqualFold predicate on the "service" field.
func ServiceEqualFold(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldService), v))
	})
}

// ServiceContainsFold applies the ContainsFold predicate on the "service" field.
func ServiceContainsFold(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldService), v))
	})
}

// IdentifierEQ applies the EQ predicate on the "identifier" field.
func IdentifierEQ(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIdentifier), v))
	})
}

// IdentifierNEQ applies the NEQ predicate on the "identifier" field.
func IdentifierNEQ(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIdentifier), v))
	})
}

// IdentifierIn applies the In predicate on the "identifier" field.
func IdentifierIn(vs ...string) predicate.Authentication {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldIdentifier), v...))
	})
}

// IdentifierNotIn applies the NotIn predicate on the "identifier" field.
func IdentifierNotIn(vs ...string) predicate.Authentication {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldIdentifier), v...))
	})
}

// IdentifierGT applies the GT predicate on the "identifier" field.
func IdentifierGT(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIdentifier), v))
	})
}

// IdentifierGTE applies the GTE predicate on the "identifier" field.
func IdentifierGTE(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIdentifier), v))
	})
}

// IdentifierLT applies the LT predicate on the "identifier" field.
func IdentifierLT(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIdentifier), v))
	})
}

// IdentifierLTE applies the LTE predicate on the "identifier" field.
func IdentifierLTE(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIdentifier), v))
	})
}

// IdentifierContains applies the Contains predicate on the "identifier" field.
func IdentifierContains(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIdentifier), v))
	})
}

// IdentifierHasPrefix applies the HasPrefix predicate on the "identifier" field.
func IdentifierHasPrefix(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIdentifier), v))
	})
}

// IdentifierHasSuffix applies the HasSuffix predicate on the "identifier" field.
func IdentifierHasSuffix(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIdentifier), v))
	})
}

// IdentifierEqualFold applies the EqualFold predicate on the "identifier" field.
func IdentifierEqualFold(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIdentifier), v))
	})
}

// IdentifierContainsFold applies the ContainsFold predicate on the "identifier" field.
func IdentifierContainsFold(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIdentifier), v))
	})
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToken), v))
	})
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldToken), v))
	})
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...string) predicate.Authentication {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldToken), v...))
	})
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...string) predicate.Authentication {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldToken), v...))
	})
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldToken), v))
	})
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldToken), v))
	})
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldToken), v))
	})
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldToken), v))
	})
}

// TokenContains applies the Contains predicate on the "token" field.
func TokenContains(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldToken), v))
	})
}

// TokenHasPrefix applies the HasPrefix predicate on the "token" field.
func TokenHasPrefix(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldToken), v))
	})
}

// TokenHasSuffix applies the HasSuffix predicate on the "token" field.
func TokenHasSuffix(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldToken), v))
	})
}

// TokenEqualFold applies the EqualFold predicate on the "token" field.
func TokenEqualFold(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldToken), v))
	})
}

// TokenContainsFold applies the ContainsFold predicate on the "token" field.
func TokenContainsFold(v string) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldToken), v))
	})
}

// MetadataIsNil applies the IsNil predicate on the "metadata" field.
func MetadataIsNil() predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMetadata)))
	})
}

// MetadataNotNil applies the NotNil predicate on the "metadata" field.
func MetadataNotNil() predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMetadata)))
	})
}

// HasAccount applies the HasEdge predicate on the "account" edge.
func HasAccount() predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AccountTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AccountTable, AccountColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccountWith applies the HasEdge predicate on the "account" edge with a given conditions (other predicates).
func HasAccountWith(preds ...predicate.Account) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AccountInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AccountTable, AccountColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Authentication) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Authentication) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
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
func Not(p predicate.Authentication) predicate.Authentication {
	return predicate.Authentication(func(s *sql.Selector) {
		p(s.Not())
	})
}
