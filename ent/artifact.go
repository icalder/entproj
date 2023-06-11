// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/icalder/entproj/ent/artifact"
	"github.com/icalder/entproj/ent/repository"
	"github.com/rs/xid"
)

// Artifact is the model entity for the Artifact schema.
type Artifact struct {
	config `json:"-"`
	// ID of the ent.
	ID xid.ID `json:"id,omitempty"`
	// Digest holds the value of the "digest" field.
	Digest string `json:"digest,omitempty"`
	// MediaType holds the value of the "mediaType" field.
	MediaType string `json:"mediaType,omitempty"`
	// Tags holds the value of the "tags" field.
	Tags []string `json:"tags,omitempty"`
	// ArtifactType holds the value of the "artifactType" field.
	ArtifactType string `json:"artifactType,omitempty"`
	// LastPush holds the value of the "lastPush" field.
	LastPush time.Time `json:"lastPush,omitempty"`
	// LastPull holds the value of the "lastPull" field.
	LastPull time.Time `json:"lastPull,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ArtifactQuery when eager-loading is set.
	Edges                ArtifactEdges `json:"edges"`
	artifact_children    *xid.ID
	repository_artifacts *xid.ID
	selectValues         sql.SelectValues
}

// ArtifactEdges holds the relations/edges for other nodes in the graph.
type ArtifactEdges struct {
	// Repository holds the value of the repository edge.
	Repository *Repository `json:"repository,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *Artifact `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*Artifact `json:"children,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// RepositoryOrErr returns the Repository value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArtifactEdges) RepositoryOrErr() (*Repository, error) {
	if e.loadedTypes[0] {
		if e.Repository == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: repository.Label}
		}
		return e.Repository, nil
	}
	return nil, &NotLoadedError{edge: "repository"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArtifactEdges) ParentOrErr() (*Artifact, error) {
	if e.loadedTypes[1] {
		if e.Parent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: artifact.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e ArtifactEdges) ChildrenOrErr() ([]*Artifact, error) {
	if e.loadedTypes[2] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Artifact) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case artifact.FieldTags:
			values[i] = new([]byte)
		case artifact.FieldDigest, artifact.FieldMediaType, artifact.FieldArtifactType:
			values[i] = new(sql.NullString)
		case artifact.FieldLastPush, artifact.FieldLastPull:
			values[i] = new(sql.NullTime)
		case artifact.FieldID:
			values[i] = new(xid.ID)
		case artifact.ForeignKeys[0]: // artifact_children
			values[i] = &sql.NullScanner{S: new(xid.ID)}
		case artifact.ForeignKeys[1]: // repository_artifacts
			values[i] = &sql.NullScanner{S: new(xid.ID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Artifact fields.
func (a *Artifact) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case artifact.FieldID:
			if value, ok := values[i].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case artifact.FieldDigest:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field digest", values[i])
			} else if value.Valid {
				a.Digest = value.String
			}
		case artifact.FieldMediaType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mediaType", values[i])
			} else if value.Valid {
				a.MediaType = value.String
			}
		case artifact.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &a.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case artifact.FieldArtifactType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field artifactType", values[i])
			} else if value.Valid {
				a.ArtifactType = value.String
			}
		case artifact.FieldLastPush:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field lastPush", values[i])
			} else if value.Valid {
				a.LastPush = value.Time
			}
		case artifact.FieldLastPull:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field lastPull", values[i])
			} else if value.Valid {
				a.LastPull = value.Time
			}
		case artifact.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field artifact_children", values[i])
			} else if value.Valid {
				a.artifact_children = new(xid.ID)
				*a.artifact_children = *value.S.(*xid.ID)
			}
		case artifact.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field repository_artifacts", values[i])
			} else if value.Valid {
				a.repository_artifacts = new(xid.ID)
				*a.repository_artifacts = *value.S.(*xid.ID)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Artifact.
// This includes values selected through modifiers, order, etc.
func (a *Artifact) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryRepository queries the "repository" edge of the Artifact entity.
func (a *Artifact) QueryRepository() *RepositoryQuery {
	return NewArtifactClient(a.config).QueryRepository(a)
}

// QueryParent queries the "parent" edge of the Artifact entity.
func (a *Artifact) QueryParent() *ArtifactQuery {
	return NewArtifactClient(a.config).QueryParent(a)
}

// QueryChildren queries the "children" edge of the Artifact entity.
func (a *Artifact) QueryChildren() *ArtifactQuery {
	return NewArtifactClient(a.config).QueryChildren(a)
}

// Update returns a builder for updating this Artifact.
// Note that you need to call Artifact.Unwrap() before calling this method if this Artifact
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Artifact) Update() *ArtifactUpdateOne {
	return NewArtifactClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Artifact entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Artifact) Unwrap() *Artifact {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Artifact is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Artifact) String() string {
	var builder strings.Builder
	builder.WriteString("Artifact(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("digest=")
	builder.WriteString(a.Digest)
	builder.WriteString(", ")
	builder.WriteString("mediaType=")
	builder.WriteString(a.MediaType)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", a.Tags))
	builder.WriteString(", ")
	builder.WriteString("artifactType=")
	builder.WriteString(a.ArtifactType)
	builder.WriteString(", ")
	builder.WriteString("lastPush=")
	builder.WriteString(a.LastPush.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("lastPull=")
	builder.WriteString(a.LastPull.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Artifacts is a parsable slice of Artifact.
type Artifacts []*Artifact