// Code generated by ent, DO NOT EDIT.

package model

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Southclaws/storyden/internal/infrastructure/db/model/notification"
	"github.com/Southclaws/storyden/internal/infrastructure/db/model/subscription"
	"github.com/google/uuid"
)

// Notification is the model entity for the Notification schema.
type Notification struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Link holds the value of the "link" field.
	Link string `json:"link,omitempty"`
	// Read holds the value of the "read" field.
	Read bool `json:"read,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NotificationQuery when eager-loading is set.
	Edges                      NotificationEdges `json:"edges"`
	notification_subscription  *uuid.UUID
	subscription_notifications *uuid.UUID
}

// NotificationEdges holds the relations/edges for other nodes in the graph.
type NotificationEdges struct {
	// Subscription holds the value of the subscription edge.
	Subscription *Subscription `json:"subscription,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SubscriptionOrErr returns the Subscription value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NotificationEdges) SubscriptionOrErr() (*Subscription, error) {
	if e.loadedTypes[0] {
		if e.Subscription == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: subscription.Label}
		}
		return e.Subscription, nil
	}
	return nil, &NotLoadedError{edge: "subscription"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Notification) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case notification.FieldRead:
			values[i] = new(sql.NullBool)
		case notification.FieldTitle, notification.FieldDescription, notification.FieldLink:
			values[i] = new(sql.NullString)
		case notification.FieldCreateTime:
			values[i] = new(sql.NullTime)
		case notification.FieldID:
			values[i] = new(uuid.UUID)
		case notification.ForeignKeys[0]: // notification_subscription
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case notification.ForeignKeys[1]: // subscription_notifications
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Notification", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Notification fields.
func (n *Notification) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case notification.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				n.ID = *value
			}
		case notification.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				n.Title = value.String
			}
		case notification.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				n.Description = value.String
			}
		case notification.FieldLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field link", values[i])
			} else if value.Valid {
				n.Link = value.String
			}
		case notification.FieldRead:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field read", values[i])
			} else if value.Valid {
				n.Read = value.Bool
			}
		case notification.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				n.CreateTime = value.Time
			}
		case notification.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field notification_subscription", values[i])
			} else if value.Valid {
				n.notification_subscription = new(uuid.UUID)
				*n.notification_subscription = *value.S.(*uuid.UUID)
			}
		case notification.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field subscription_notifications", values[i])
			} else if value.Valid {
				n.subscription_notifications = new(uuid.UUID)
				*n.subscription_notifications = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QuerySubscription queries the "subscription" edge of the Notification entity.
func (n *Notification) QuerySubscription() *SubscriptionQuery {
	return (&NotificationClient{config: n.config}).QuerySubscription(n)
}

// Update returns a builder for updating this Notification.
// Note that you need to call Notification.Unwrap() before calling this method if this Notification
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Notification) Update() *NotificationUpdateOne {
	return (&NotificationClient{config: n.config}).UpdateOne(n)
}

// Unwrap unwraps the Notification entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Notification) Unwrap() *Notification {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("model: Notification is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Notification) String() string {
	var builder strings.Builder
	builder.WriteString("Notification(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	builder.WriteString("title=")
	builder.WriteString(n.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(n.Description)
	builder.WriteString(", ")
	builder.WriteString("link=")
	builder.WriteString(n.Link)
	builder.WriteString(", ")
	builder.WriteString("read=")
	builder.WriteString(fmt.Sprintf("%v", n.Read))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(n.CreateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Notifications is a parsable slice of Notification.
type Notifications []*Notification

func (n Notifications) config(cfg config) {
	for _i := range n {
		n[_i].config = cfg
	}
}
