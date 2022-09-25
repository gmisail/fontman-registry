package schema

import "entgo.io/ent"

// FontStyle holds the schema definition for the FontStyle entity.
type FontStyle struct {
	ent.Schema
}

// Fields of the FontStyle.
func (FontStyle) Fields() []ent.Field {
	return nil
}

// Edges of the FontStyle.
func (FontStyle) Edges() []ent.Edge {
	return nil
}
