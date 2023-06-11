package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Repository holds the schema definition for the Repository entity.
type Repository struct {
	ent.Schema
}

// Mixin of the Repository.
func (Repository) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// Embed the UUIDMixin
		UUIDMixin{},
	}
}

// Fields of the Repository.
func (Repository) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").NotEmpty(),
	}
}

// Edges of the Repository.
func (Repository) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("registry", Registry.Type).
			Ref("repositories").
			Required().
			Unique(),
		// edge.To("artifacts", Artifact.Type).
		// Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

// Indexes of the Repository.
func (Repository) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").
			Edges("registry").
			Unique(),
	}
}
