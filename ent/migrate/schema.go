// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// FontFamiliesColumns holds the columns for the "font_families" table.
	FontFamiliesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// FontFamiliesTable holds the schema information for the "font_families" table.
	FontFamiliesTable = &schema.Table{
		Name:       "font_families",
		Columns:    FontFamiliesColumns,
		PrimaryKey: []*schema.Column{FontFamiliesColumns[0]},
	}
	// FontStylesColumns holds the columns for the "font_styles" table.
	FontStylesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// FontStylesTable holds the schema information for the "font_styles" table.
	FontStylesTable = &schema.Table{
		Name:       "font_styles",
		Columns:    FontStylesColumns,
		PrimaryKey: []*schema.Column{FontStylesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		FontFamiliesTable,
		FontStylesTable,
	}
)

func init() {
}