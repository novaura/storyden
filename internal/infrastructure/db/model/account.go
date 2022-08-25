// Code generated by ent, DO NOT EDIT.

package model

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Southclaws/storyden/internal/infrastructure/db/model/account"
	"github.com/google/uuid"
)

// Account is the model entity for the Account schema.
type Account struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Bio holds the value of the "bio" field.
	Bio string `json:"bio,omitempty"`
	// Admin holds the value of the "admin" field.
	Admin bool `json:"admin,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// DeletedAt holds the value of the "deletedAt" field.
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountQuery when eager-loading is set.
	Edges AccountEdges `json:"edges"`
}

// AccountEdges holds the relations/edges for other nodes in the graph.
type AccountEdges struct {
	// Posts holds the value of the posts edge.
	Posts []*Post `json:"posts,omitempty"`
	// Reacts holds the value of the reacts edge.
	Reacts []*React `json:"reacts,omitempty"`
	// Subscriptions holds the value of the subscriptions edge.
	Subscriptions []*Subscription `json:"subscriptions,omitempty"`
	// Authentication holds the value of the authentication edge.
	Authentication []*Authentication `json:"authentication,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// PostsOrErr returns the Posts value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) PostsOrErr() ([]*Post, error) {
	if e.loadedTypes[0] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "posts"}
}

// ReactsOrErr returns the Reacts value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) ReactsOrErr() ([]*React, error) {
	if e.loadedTypes[1] {
		return e.Reacts, nil
	}
	return nil, &NotLoadedError{edge: "reacts"}
}

// SubscriptionsOrErr returns the Subscriptions value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) SubscriptionsOrErr() ([]*Subscription, error) {
	if e.loadedTypes[2] {
		return e.Subscriptions, nil
	}
	return nil, &NotLoadedError{edge: "subscriptions"}
}

// AuthenticationOrErr returns the Authentication value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) AuthenticationOrErr() ([]*Authentication, error) {
	if e.loadedTypes[3] {
		return e.Authentication, nil
	}
	return nil, &NotLoadedError{edge: "authentication"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Account) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case account.FieldAdmin:
			values[i] = new(sql.NullBool)
		case account.FieldEmail, account.FieldName, account.FieldBio:
			values[i] = new(sql.NullString)
		case account.FieldCreatedAt, account.FieldUpdatedAt, account.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case account.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Account", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Account fields.
func (a *Account) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case account.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case account.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				a.Email = value.String
			}
		case account.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case account.FieldBio:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bio", values[i])
			} else if value.Valid {
				a.Bio = value.String
			}
		case account.FieldAdmin:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field admin", values[i])
			} else if value.Valid {
				a.Admin = value.Bool
			}
		case account.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case account.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case account.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deletedAt", values[i])
			} else if value.Valid {
				a.DeletedAt = value.Time
			}
		}
	}
	return nil
}

// QueryPosts queries the "posts" edge of the Account entity.
func (a *Account) QueryPosts() *PostQuery {
	return (&AccountClient{config: a.config}).QueryPosts(a)
}

// QueryReacts queries the "reacts" edge of the Account entity.
func (a *Account) QueryReacts() *ReactQuery {
	return (&AccountClient{config: a.config}).QueryReacts(a)
}

// QuerySubscriptions queries the "subscriptions" edge of the Account entity.
func (a *Account) QuerySubscriptions() *SubscriptionQuery {
	return (&AccountClient{config: a.config}).QuerySubscriptions(a)
}

// QueryAuthentication queries the "authentication" edge of the Account entity.
func (a *Account) QueryAuthentication() *AuthenticationQuery {
	return (&AccountClient{config: a.config}).QueryAuthentication(a)
}

// Update returns a builder for updating this Account.
// Note that you need to call Account.Unwrap() before calling this method if this Account
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Account) Update() *AccountUpdateOne {
	return (&AccountClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Account entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Account) Unwrap() *Account {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("model: Account is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Account) String() string {
	var builder strings.Builder
	builder.WriteString("Account(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("email=")
	builder.WriteString(a.Email)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("bio=")
	builder.WriteString(a.Bio)
	builder.WriteString(", ")
	builder.WriteString("admin=")
	builder.WriteString(fmt.Sprintf("%v", a.Admin))
	builder.WriteString(", ")
	builder.WriteString("createdAt=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deletedAt=")
	builder.WriteString(a.DeletedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Accounts is a parsable slice of Account.
type Accounts []*Account

func (a Accounts) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
