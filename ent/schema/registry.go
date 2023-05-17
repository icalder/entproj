package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Registry holds the schema definition for the Registry entity.
type Registry struct {
	ent.Schema
}

// Fields of the Registry.
func (Registry) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.Text("name").NotEmpty(),
	}
}

// Edges of the Registry.
func (Registry) Edges() []ent.Edge {
	// https://entgo.io/docs/schema-edges/
	return []ent.Edge{
		edge.To("repositories", Repository.Type),
	}
}
