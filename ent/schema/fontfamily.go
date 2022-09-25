package schema

import "entgo.io/ent"

// FontFamily holds the schema definition for the FontFamily entity.
type FontFamily struct {
	ent.Schema
}

// Fields of the FontFamily.
func (FontFamily) Fields() []ent.Field {
	return nil
}

// Edges of the FontFamily.
func (FontFamily) Edges() []ent.Edge {
	return nil
}
