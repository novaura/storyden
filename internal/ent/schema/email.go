package schema

import (
	"net/mail"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

type Email struct {
	ent.Schema
}

func (Email) Mixin() []ent.Mixin {
	return []ent.Mixin{Identifier{}, CreatedAt{}}
}

func (Email) Fields() []ent.Field {
	return []ent.Field{
		field.String("account_id").
			Optional().
			Nillable().
			GoType(xid.ID{}).
			Comment("If set, this email is associated with an account, otherwise can be used for newsletter subscriptions etc."),

		field.String("authentication_record_id").
			Optional().
			Nillable().
			GoType(xid.ID{}).
			Comment("If set, this this email is used for authentication"),

		field.String("email_address").
			NotEmpty().
			Immutable().
			Unique().
			MinLen(3).
			MaxLen(254).Validate(func(s string) error {
			_, err := mail.ParseAddress(s)
			return err
		}),

		field.String("verification_code").
			MaxLen(6).
			Comment("A six digit code that is sent to the email address to verify ownership"),

		field.Bool("verified").
			Default(false).
			Annotations(entsql.Default("false")).
			Comment("Whether this email has been verified to be owned by the account via a token send+verify process"),

		field.Bool("public").
			Default(false).
			Annotations(entsql.Default("false")).
			Comment("Whether this email is publicly available via the public profile schema."),
	}
}

func (Email) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
			Ref("emails").
			Field("account_id").
			Unique(),

		edge.From("authentication", Authentication.Type).
			Ref("email_address").
			Field("authentication_record_id").
			Unique(),
	}
}
