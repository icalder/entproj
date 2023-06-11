// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// RegistriesColumns holds the columns for the "registries" table.
	RegistriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Size: 2147483647},
	}
	// RegistriesTable holds the schema information for the "registries" table.
	RegistriesTable = &schema.Table{
		Name:       "registries",
		Columns:    RegistriesColumns,
		PrimaryKey: []*schema.Column{RegistriesColumns[0]},
	}
	// RepositoriesColumns holds the columns for the "repositories" table.
	RepositoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Size: 2147483647},
		{Name: "registry_repositories", Type: field.TypeUUID},
	}
	// RepositoriesTable holds the schema information for the "repositories" table.
	RepositoriesTable = &schema.Table{
		Name:       "repositories",
		Columns:    RepositoriesColumns,
		PrimaryKey: []*schema.Column{RepositoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "repositories_registries_repositories",
				Columns:    []*schema.Column{RepositoriesColumns[2]},
				RefColumns: []*schema.Column{RegistriesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "repository_name_registry_repositories",
				Unique:  true,
				Columns: []*schema.Column{RepositoriesColumns[1], RepositoriesColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		RegistriesTable,
		RepositoriesTable,
	}
)

func init() {
	RepositoriesTable.ForeignKeys[0].RefTable = RegistriesTable
}
