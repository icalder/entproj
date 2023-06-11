package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Artifact holds the schema definition for the Artifact entity.
type Artifact struct {
	ent.Schema
}

// Mixin of the Artifact.
func (Artifact) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// Embed the UUIDMixin
		UUIDMixin{},
	}
}

// Fields of the Artifact.
func (Artifact) Fields() []ent.Field {
	return []ent.Field{
		field.Text("digest").NotEmpty(),
		field.Text("mediaType").NotEmpty(),
		field.JSON("tags", []string{}).Optional(),
		field.Text("artifactType").Optional(),
		field.Time("lastPush").Optional(),
		field.Time("lastPull").Optional(),
	}
}

// Edges of the Artifact.
func (Artifact) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("repository", Repository.Type).
			Ref("artifacts").
			Required().
			Unique(),
		edge.To("children", Artifact.Type).
			From("parent").
			Unique(),
	}
}

// Indexes of the Artifact.
func (Artifact) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("digest").
			Edges("repository").
			Unique(),
	}
}
