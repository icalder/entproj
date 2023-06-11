package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/rs/xid"
)

// UUIDMixin holds the schema definition for the UUIDMixin entity.
type UUIDMixin struct {
	mixin.Schema
}

// Fields of the UUIDMixin.
func (UUIDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(xid.ID{}).
			DefaultFunc(xid.New),
	}
}
