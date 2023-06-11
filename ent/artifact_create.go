// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/icalder/entproj/ent/artifact"
	"github.com/icalder/entproj/ent/repository"
	"github.com/rs/xid"
)

// ArtifactCreate is the builder for creating a Artifact entity.
type ArtifactCreate struct {
	config
	mutation *ArtifactMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetDigest sets the "digest" field.
func (ac *ArtifactCreate) SetDigest(s string) *ArtifactCreate {
	ac.mutation.SetDigest(s)
	return ac
}

// SetMediaType sets the "mediaType" field.
func (ac *ArtifactCreate) SetMediaType(s string) *ArtifactCreate {
	ac.mutation.SetMediaType(s)
	return ac
}

// SetTags sets the "tags" field.
func (ac *ArtifactCreate) SetTags(s []string) *ArtifactCreate {
	ac.mutation.SetTags(s)
	return ac
}

// SetArtifactType sets the "artifactType" field.
func (ac *ArtifactCreate) SetArtifactType(s string) *ArtifactCreate {
	ac.mutation.SetArtifactType(s)
	return ac
}

// SetNillableArtifactType sets the "artifactType" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillableArtifactType(s *string) *ArtifactCreate {
	if s != nil {
		ac.SetArtifactType(*s)
	}
	return ac
}

// SetLastPush sets the "lastPush" field.
func (ac *ArtifactCreate) SetLastPush(t time.Time) *ArtifactCreate {
	ac.mutation.SetLastPush(t)
	return ac
}

// SetNillableLastPush sets the "lastPush" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillableLastPush(t *time.Time) *ArtifactCreate {
	if t != nil {
		ac.SetLastPush(*t)
	}
	return ac
}

// SetLastPull sets the "lastPull" field.
func (ac *ArtifactCreate) SetLastPull(t time.Time) *ArtifactCreate {
	ac.mutation.SetLastPull(t)
	return ac
}

// SetNillableLastPull sets the "lastPull" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillableLastPull(t *time.Time) *ArtifactCreate {
	if t != nil {
		ac.SetLastPull(*t)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *ArtifactCreate) SetID(x xid.ID) *ArtifactCreate {
	ac.mutation.SetID(x)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillableID(x *xid.ID) *ArtifactCreate {
	if x != nil {
		ac.SetID(*x)
	}
	return ac
}

// SetRepositoryID sets the "repository" edge to the Repository entity by ID.
func (ac *ArtifactCreate) SetRepositoryID(id xid.ID) *ArtifactCreate {
	ac.mutation.SetRepositoryID(id)
	return ac
}

// SetRepository sets the "repository" edge to the Repository entity.
func (ac *ArtifactCreate) SetRepository(r *Repository) *ArtifactCreate {
	return ac.SetRepositoryID(r.ID)
}

// SetParentID sets the "parent" edge to the Artifact entity by ID.
func (ac *ArtifactCreate) SetParentID(id xid.ID) *ArtifactCreate {
	ac.mutation.SetParentID(id)
	return ac
}

// SetNillableParentID sets the "parent" edge to the Artifact entity by ID if the given value is not nil.
func (ac *ArtifactCreate) SetNillableParentID(id *xid.ID) *ArtifactCreate {
	if id != nil {
		ac = ac.SetParentID(*id)
	}
	return ac
}

// SetParent sets the "parent" edge to the Artifact entity.
func (ac *ArtifactCreate) SetParent(a *Artifact) *ArtifactCreate {
	return ac.SetParentID(a.ID)
}

// AddChildIDs adds the "children" edge to the Artifact entity by IDs.
func (ac *ArtifactCreate) AddChildIDs(ids ...xid.ID) *ArtifactCreate {
	ac.mutation.AddChildIDs(ids...)
	return ac
}

// AddChildren adds the "children" edges to the Artifact entity.
func (ac *ArtifactCreate) AddChildren(a ...*Artifact) *ArtifactCreate {
	ids := make([]xid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddChildIDs(ids...)
}

// Mutation returns the ArtifactMutation object of the builder.
func (ac *ArtifactCreate) Mutation() *ArtifactMutation {
	return ac.mutation
}

// Save creates the Artifact in the database.
func (ac *ArtifactCreate) Save(ctx context.Context) (*Artifact, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ArtifactCreate) SaveX(ctx context.Context) *Artifact {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ArtifactCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ArtifactCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *ArtifactCreate) defaults() {
	if _, ok := ac.mutation.ID(); !ok {
		v := artifact.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ArtifactCreate) check() error {
	if _, ok := ac.mutation.Digest(); !ok {
		return &ValidationError{Name: "digest", err: errors.New(`ent: missing required field "Artifact.digest"`)}
	}
	if v, ok := ac.mutation.Digest(); ok {
		if err := artifact.DigestValidator(v); err != nil {
			return &ValidationError{Name: "digest", err: fmt.Errorf(`ent: validator failed for field "Artifact.digest": %w`, err)}
		}
	}
	if _, ok := ac.mutation.MediaType(); !ok {
		return &ValidationError{Name: "mediaType", err: errors.New(`ent: missing required field "Artifact.mediaType"`)}
	}
	if v, ok := ac.mutation.MediaType(); ok {
		if err := artifact.MediaTypeValidator(v); err != nil {
			return &ValidationError{Name: "mediaType", err: fmt.Errorf(`ent: validator failed for field "Artifact.mediaType": %w`, err)}
		}
	}
	if _, ok := ac.mutation.RepositoryID(); !ok {
		return &ValidationError{Name: "repository", err: errors.New(`ent: missing required edge "Artifact.repository"`)}
	}
	return nil
}

func (ac *ArtifactCreate) sqlSave(ctx context.Context) (*Artifact, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*xid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *ArtifactCreate) createSpec() (*Artifact, *sqlgraph.CreateSpec) {
	var (
		_node = &Artifact{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(artifact.Table, sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeString))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.Digest(); ok {
		_spec.SetField(artifact.FieldDigest, field.TypeString, value)
		_node.Digest = value
	}
	if value, ok := ac.mutation.MediaType(); ok {
		_spec.SetField(artifact.FieldMediaType, field.TypeString, value)
		_node.MediaType = value
	}
	if value, ok := ac.mutation.Tags(); ok {
		_spec.SetField(artifact.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	if value, ok := ac.mutation.ArtifactType(); ok {
		_spec.SetField(artifact.FieldArtifactType, field.TypeString, value)
		_node.ArtifactType = value
	}
	if value, ok := ac.mutation.LastPush(); ok {
		_spec.SetField(artifact.FieldLastPush, field.TypeTime, value)
		_node.LastPush = value
	}
	if value, ok := ac.mutation.LastPull(); ok {
		_spec.SetField(artifact.FieldLastPull, field.TypeTime, value)
		_node.LastPull = value
	}
	if nodes := ac.mutation.RepositoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   artifact.RepositoryTable,
			Columns: []string{artifact.RepositoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(repository.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.repository_artifacts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   artifact.ParentTable,
			Columns: []string{artifact.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.artifact_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   artifact.ChildrenTable,
			Columns: []string{artifact.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Artifact.Create().
//		SetDigest(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ArtifactUpsert) {
//			SetDigest(v+v).
//		}).
//		Exec(ctx)
func (ac *ArtifactCreate) OnConflict(opts ...sql.ConflictOption) *ArtifactUpsertOne {
	ac.conflict = opts
	return &ArtifactUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Artifact.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *ArtifactCreate) OnConflictColumns(columns ...string) *ArtifactUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &ArtifactUpsertOne{
		create: ac,
	}
}

type (
	// ArtifactUpsertOne is the builder for "upsert"-ing
	//  one Artifact node.
	ArtifactUpsertOne struct {
		create *ArtifactCreate
	}

	// ArtifactUpsert is the "OnConflict" setter.
	ArtifactUpsert struct {
		*sql.UpdateSet
	}
)

// SetDigest sets the "digest" field.
func (u *ArtifactUpsert) SetDigest(v string) *ArtifactUpsert {
	u.Set(artifact.FieldDigest, v)
	return u
}

// UpdateDigest sets the "digest" field to the value that was provided on create.
func (u *ArtifactUpsert) UpdateDigest() *ArtifactUpsert {
	u.SetExcluded(artifact.FieldDigest)
	return u
}

// SetMediaType sets the "mediaType" field.
func (u *ArtifactUpsert) SetMediaType(v string) *ArtifactUpsert {
	u.Set(artifact.FieldMediaType, v)
	return u
}

// UpdateMediaType sets the "mediaType" field to the value that was provided on create.
func (u *ArtifactUpsert) UpdateMediaType() *ArtifactUpsert {
	u.SetExcluded(artifact.FieldMediaType)
	return u
}

// SetTags sets the "tags" field.
func (u *ArtifactUpsert) SetTags(v []string) *ArtifactUpsert {
	u.Set(artifact.FieldTags, v)
	return u
}

// UpdateTags sets the "tags" field to the value that was provided on create.
func (u *ArtifactUpsert) UpdateTags() *ArtifactUpsert {
	u.SetExcluded(artifact.FieldTags)
	return u
}

// ClearTags clears the value of the "tags" field.
func (u *ArtifactUpsert) ClearTags() *ArtifactUpsert {
	u.SetNull(artifact.FieldTags)
	return u
}

// SetArtifactType sets the "artifactType" field.
func (u *ArtifactUpsert) SetArtifactType(v string) *ArtifactUpsert {
	u.Set(artifact.FieldArtifactType, v)
	return u
}

// UpdateArtifactType sets the "artifactType" field to the value that was provided on create.
func (u *ArtifactUpsert) UpdateArtifactType() *ArtifactUpsert {
	u.SetExcluded(artifact.FieldArtifactType)
	return u
}

// ClearArtifactType clears the value of the "artifactType" field.
func (u *ArtifactUpsert) ClearArtifactType() *ArtifactUpsert {
	u.SetNull(artifact.FieldArtifactType)
	return u
}

// SetLastPush sets the "lastPush" field.
func (u *ArtifactUpsert) SetLastPush(v time.Time) *ArtifactUpsert {
	u.Set(artifact.FieldLastPush, v)
	return u
}

// UpdateLastPush sets the "lastPush" field to the value that was provided on create.
func (u *ArtifactUpsert) UpdateLastPush() *ArtifactUpsert {
	u.SetExcluded(artifact.FieldLastPush)
	return u
}

// ClearLastPush clears the value of the "lastPush" field.
func (u *ArtifactUpsert) ClearLastPush() *ArtifactUpsert {
	u.SetNull(artifact.FieldLastPush)
	return u
}

// SetLastPull sets the "lastPull" field.
func (u *ArtifactUpsert) SetLastPull(v time.Time) *ArtifactUpsert {
	u.Set(artifact.FieldLastPull, v)
	return u
}

// UpdateLastPull sets the "lastPull" field to the value that was provided on create.
func (u *ArtifactUpsert) UpdateLastPull() *ArtifactUpsert {
	u.SetExcluded(artifact.FieldLastPull)
	return u
}

// ClearLastPull clears the value of the "lastPull" field.
func (u *ArtifactUpsert) ClearLastPull() *ArtifactUpsert {
	u.SetNull(artifact.FieldLastPull)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Artifact.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(artifact.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ArtifactUpsertOne) UpdateNewValues() *ArtifactUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(artifact.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Artifact.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ArtifactUpsertOne) Ignore() *ArtifactUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ArtifactUpsertOne) DoNothing() *ArtifactUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ArtifactCreate.OnConflict
// documentation for more info.
func (u *ArtifactUpsertOne) Update(set func(*ArtifactUpsert)) *ArtifactUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ArtifactUpsert{UpdateSet: update})
	}))
	return u
}

// SetDigest sets the "digest" field.
func (u *ArtifactUpsertOne) SetDigest(v string) *ArtifactUpsertOne {
	return u.Update(func(s *ArtifactUpsert) {
		s.SetDigest(v)
	})
}

// UpdateDigest sets the "digest" field to the value that was provided on create.
func (u *ArtifactUpsertOne) UpdateDigest() *ArtifactUpsertOne {
	return u.Update(func(s *ArtifactUpsert) {
		s.UpdateDigest()
	})
}

// SetMediaType sets the "mediaType" field.
func (u *ArtifactUpsertOne) SetMediaType(v string) *ArtifactUpsertOne {
	return u.Update(func(s *ArtifactUpsert) {
		s.SetMediaType(v)
	})
}

// UpdateMediaType sets the "mediaType" field to the value that was provided on create.
func (u *ArtifactUpsertOne) UpdateMediaType() *ArtifactUpsertOne {
	return u.Update(func(s *ArtifactUpsert) {
		s.UpdateMediaType()
	})
}

// SetTags sets the "tags" field.
func (u *ArtifactUpsertOne) SetTags(v []string) *ArtifactUpsertOne {
	return u.Update(func(s *ArtifactUpsert) {
		s.SetTags(v)
	})
}

// UpdateTags sets the "tags" field to the value that was provided on create.
func (u *ArtifactUpsertOne) UpdateTags() *ArtifactUpsertOne {
	return u.Update(func(s *ArtifactUpsert) {
		s.UpdateTags()
	})
}

// ClearTags clears the value of the "tags" field.
func (u *ArtifactUpsertOne) ClearTags() *ArtifactUpsertOne {
	return u.Update(func(s *ArtifactUpsert) {
		s.ClearTags()
	})
}

// SetArtifactType sets the "artifactType" field.
func (u *ArtifactUpsertOne) SetArtifactType(v string) *ArtifactUpsertOne {
	return u.Update(func(s *ArtifactUpsert) {
		s.SetArtifactType(v)
	})
}

// UpdateArtifactType sets the "artifactType" field to the value that was provided on create.
func (u *ArtifactUpsertOne) UpdateArtifactType() *ArtifactUpsertOne {
	return u.Update(func(s *ArtifactUpsert) {
		s.UpdateArtifactType()
	})
}

// ClearArtifactType clears the value of the "artifactType" field.
func (u *ArtifactUpsertOne) ClearArtifactType() *ArtifactUpsertOne {
	return u.Update(func(s *ArtifactUpsert) {
		s.ClearArtifactType()
	})
}

// SetLastPush sets the "lastPush" field.
func (u *ArtifactUpsertOne) SetLastPush(v time.Time) *ArtifactUpsertOne {
	return u.Update(func(s *ArtifactUpsert) {
		s.SetLastPush(v)
	})
}

// UpdateLastPush sets the "lastPush" field to the value that was provided on create.
func (u *ArtifactUpsertOne) UpdateLastPush() *ArtifactUpsertOne {
	return u.Update(func(s *ArtifactUpsert) {
		s.UpdateLastPush()
	})
}

// ClearLastPush clears the value of the "lastPush" field.
func (u *ArtifactUpsertOne) ClearLastPush() *ArtifactUpsertOne {
	return u.Update(func(s *ArtifactUpsert) {
		s.ClearLastPush()
	})
}

// SetLastPull sets the "lastPull" field.
func (u *ArtifactUpsertOne) SetLastPull(v time.Time) *ArtifactUpsertOne {
	return u.Update(func(s *ArtifactUpsert) {
		s.SetLastPull(v)
	})
}

// UpdateLastPull sets the "lastPull" field to the value that was provided on create.
func (u *ArtifactUpsertOne) UpdateLastPull() *ArtifactUpsertOne {
	return u.Update(func(s *ArtifactUpsert) {
		s.UpdateLastPull()
	})
}

// ClearLastPull clears the value of the "lastPull" field.
func (u *ArtifactUpsertOne) ClearLastPull() *ArtifactUpsertOne {
	return u.Update(func(s *ArtifactUpsert) {
		s.ClearLastPull()
	})
}

// Exec executes the query.
func (u *ArtifactUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ArtifactCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ArtifactUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ArtifactUpsertOne) ID(ctx context.Context) (id xid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ArtifactUpsertOne.ID is not supported by MySQL driver. Use ArtifactUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ArtifactUpsertOne) IDX(ctx context.Context) xid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ArtifactCreateBulk is the builder for creating many Artifact entities in bulk.
type ArtifactCreateBulk struct {
	config
	builders []*ArtifactCreate
	conflict []sql.ConflictOption
}

// Save creates the Artifact entities in the database.
func (acb *ArtifactCreateBulk) Save(ctx context.Context) ([]*Artifact, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Artifact, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArtifactMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ArtifactCreateBulk) SaveX(ctx context.Context) []*Artifact {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ArtifactCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ArtifactCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Artifact.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ArtifactUpsert) {
//			SetDigest(v+v).
//		}).
//		Exec(ctx)
func (acb *ArtifactCreateBulk) OnConflict(opts ...sql.ConflictOption) *ArtifactUpsertBulk {
	acb.conflict = opts
	return &ArtifactUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Artifact.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *ArtifactCreateBulk) OnConflictColumns(columns ...string) *ArtifactUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &ArtifactUpsertBulk{
		create: acb,
	}
}

// ArtifactUpsertBulk is the builder for "upsert"-ing
// a bulk of Artifact nodes.
type ArtifactUpsertBulk struct {
	create *ArtifactCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Artifact.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(artifact.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ArtifactUpsertBulk) UpdateNewValues() *ArtifactUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(artifact.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Artifact.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ArtifactUpsertBulk) Ignore() *ArtifactUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ArtifactUpsertBulk) DoNothing() *ArtifactUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ArtifactCreateBulk.OnConflict
// documentation for more info.
func (u *ArtifactUpsertBulk) Update(set func(*ArtifactUpsert)) *ArtifactUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ArtifactUpsert{UpdateSet: update})
	}))
	return u
}

// SetDigest sets the "digest" field.
func (u *ArtifactUpsertBulk) SetDigest(v string) *ArtifactUpsertBulk {
	return u.Update(func(s *ArtifactUpsert) {
		s.SetDigest(v)
	})
}

// UpdateDigest sets the "digest" field to the value that was provided on create.
func (u *ArtifactUpsertBulk) UpdateDigest() *ArtifactUpsertBulk {
	return u.Update(func(s *ArtifactUpsert) {
		s.UpdateDigest()
	})
}

// SetMediaType sets the "mediaType" field.
func (u *ArtifactUpsertBulk) SetMediaType(v string) *ArtifactUpsertBulk {
	return u.Update(func(s *ArtifactUpsert) {
		s.SetMediaType(v)
	})
}

// UpdateMediaType sets the "mediaType" field to the value that was provided on create.
func (u *ArtifactUpsertBulk) UpdateMediaType() *ArtifactUpsertBulk {
	return u.Update(func(s *ArtifactUpsert) {
		s.UpdateMediaType()
	})
}

// SetTags sets the "tags" field.
func (u *ArtifactUpsertBulk) SetTags(v []string) *ArtifactUpsertBulk {
	return u.Update(func(s *ArtifactUpsert) {
		s.SetTags(v)
	})
}

// UpdateTags sets the "tags" field to the value that was provided on create.
func (u *ArtifactUpsertBulk) UpdateTags() *ArtifactUpsertBulk {
	return u.Update(func(s *ArtifactUpsert) {
		s.UpdateTags()
	})
}

// ClearTags clears the value of the "tags" field.
func (u *ArtifactUpsertBulk) ClearTags() *ArtifactUpsertBulk {
	return u.Update(func(s *ArtifactUpsert) {
		s.ClearTags()
	})
}

// SetArtifactType sets the "artifactType" field.
func (u *ArtifactUpsertBulk) SetArtifactType(v string) *ArtifactUpsertBulk {
	return u.Update(func(s *ArtifactUpsert) {
		s.SetArtifactType(v)
	})
}

// UpdateArtifactType sets the "artifactType" field to the value that was provided on create.
func (u *ArtifactUpsertBulk) UpdateArtifactType() *ArtifactUpsertBulk {
	return u.Update(func(s *ArtifactUpsert) {
		s.UpdateArtifactType()
	})
}

// ClearArtifactType clears the value of the "artifactType" field.
func (u *ArtifactUpsertBulk) ClearArtifactType() *ArtifactUpsertBulk {
	return u.Update(func(s *ArtifactUpsert) {
		s.ClearArtifactType()
	})
}

// SetLastPush sets the "lastPush" field.
func (u *ArtifactUpsertBulk) SetLastPush(v time.Time) *ArtifactUpsertBulk {
	return u.Update(func(s *ArtifactUpsert) {
		s.SetLastPush(v)
	})
}

// UpdateLastPush sets the "lastPush" field to the value that was provided on create.
func (u *ArtifactUpsertBulk) UpdateLastPush() *ArtifactUpsertBulk {
	return u.Update(func(s *ArtifactUpsert) {
		s.UpdateLastPush()
	})
}

// ClearLastPush clears the value of the "lastPush" field.
func (u *ArtifactUpsertBulk) ClearLastPush() *ArtifactUpsertBulk {
	return u.Update(func(s *ArtifactUpsert) {
		s.ClearLastPush()
	})
}

// SetLastPull sets the "lastPull" field.
func (u *ArtifactUpsertBulk) SetLastPull(v time.Time) *ArtifactUpsertBulk {
	return u.Update(func(s *ArtifactUpsert) {
		s.SetLastPull(v)
	})
}

// UpdateLastPull sets the "lastPull" field to the value that was provided on create.
func (u *ArtifactUpsertBulk) UpdateLastPull() *ArtifactUpsertBulk {
	return u.Update(func(s *ArtifactUpsert) {
		s.UpdateLastPull()
	})
}

// ClearLastPull clears the value of the "lastPull" field.
func (u *ArtifactUpsertBulk) ClearLastPull() *ArtifactUpsertBulk {
	return u.Update(func(s *ArtifactUpsert) {
		s.ClearLastPull()
	})
}

// Exec executes the query.
func (u *ArtifactUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ArtifactCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ArtifactCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ArtifactUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}