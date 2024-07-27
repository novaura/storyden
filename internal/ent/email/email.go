// Code generated by ent, DO NOT EDIT.

package email

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/rs/xid"
)

const (
	// Label holds the string label denoting the email type in the database.
	Label = "email"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldAccountID holds the string denoting the account_id field in the database.
	FieldAccountID = "account_id"
	// FieldAuthenticationRecordID holds the string denoting the authentication_record_id field in the database.
	FieldAuthenticationRecordID = "authentication_record_id"
	// FieldEmailAddress holds the string denoting the email_address field in the database.
	FieldEmailAddress = "email_address"
	// FieldVerificationCode holds the string denoting the verification_code field in the database.
	FieldVerificationCode = "verification_code"
	// FieldVerified holds the string denoting the verified field in the database.
	FieldVerified = "verified"
	// EdgeAccount holds the string denoting the account edge name in mutations.
	EdgeAccount = "account"
	// EdgeAuthentication holds the string denoting the authentication edge name in mutations.
	EdgeAuthentication = "authentication"
	// Table holds the table name of the email in the database.
	Table = "emails"
	// AccountTable is the table that holds the account relation/edge.
	AccountTable = "emails"
	// AccountInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	AccountInverseTable = "accounts"
	// AccountColumn is the table column denoting the account relation/edge.
	AccountColumn = "account_id"
	// AuthenticationTable is the table that holds the authentication relation/edge.
	AuthenticationTable = "emails"
	// AuthenticationInverseTable is the table name for the Authentication entity.
	// It exists in this package in order to avoid circular dependency with the "authentication" package.
	AuthenticationInverseTable = "authentications"
	// AuthenticationColumn is the table column denoting the authentication relation/edge.
	AuthenticationColumn = "authentication_record_id"
)

// Columns holds all SQL columns for email fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldAccountID,
	FieldAuthenticationRecordID,
	FieldEmailAddress,
	FieldVerificationCode,
	FieldVerified,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// EmailAddressValidator is a validator for the "email_address" field. It is called by the builders before save.
	EmailAddressValidator func(string) error
	// VerificationCodeValidator is a validator for the "verification_code" field. It is called by the builders before save.
	VerificationCodeValidator func(string) error
	// DefaultVerified holds the default value on creation for the "verified" field.
	DefaultVerified bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() xid.ID
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Email queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByAccountID orders the results by the account_id field.
func ByAccountID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccountID, opts...).ToFunc()
}

// ByAuthenticationRecordID orders the results by the authentication_record_id field.
func ByAuthenticationRecordID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthenticationRecordID, opts...).ToFunc()
}

// ByEmailAddress orders the results by the email_address field.
func ByEmailAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmailAddress, opts...).ToFunc()
}

// ByVerificationCode orders the results by the verification_code field.
func ByVerificationCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVerificationCode, opts...).ToFunc()
}

// ByVerified orders the results by the verified field.
func ByVerified(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVerified, opts...).ToFunc()
}

// ByAccountField orders the results by account field.
func ByAccountField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAccountStep(), sql.OrderByField(field, opts...))
	}
}

// ByAuthenticationField orders the results by authentication field.
func ByAuthenticationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthenticationStep(), sql.OrderByField(field, opts...))
	}
}
func newAccountStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AccountInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AccountTable, AccountColumn),
	)
}
func newAuthenticationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthenticationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AuthenticationTable, AuthenticationColumn),
	)
}
