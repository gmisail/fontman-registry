// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/gmisail/fontman/registry/ent/fontstyle"
)

// FontStyle is the model entity for the FontStyle schema.
type FontStyle struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FontStyle) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case fontstyle.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type FontStyle", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FontStyle fields.
func (fs *FontStyle) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case fontstyle.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fs.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this FontStyle.
// Note that you need to call FontStyle.Unwrap() before calling this method if this FontStyle
// was returned from a transaction, and the transaction was committed or rolled back.
func (fs *FontStyle) Update() *FontStyleUpdateOne {
	return (&FontStyleClient{config: fs.config}).UpdateOne(fs)
}

// Unwrap unwraps the FontStyle entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fs *FontStyle) Unwrap() *FontStyle {
	_tx, ok := fs.config.driver.(*txDriver)
	if !ok {
		panic("ent: FontStyle is not a transactional entity")
	}
	fs.config.driver = _tx.drv
	return fs
}

// String implements the fmt.Stringer.
func (fs *FontStyle) String() string {
	var builder strings.Builder
	builder.WriteString("FontStyle(")
	builder.WriteString(fmt.Sprintf("id=%v", fs.ID))
	builder.WriteByte(')')
	return builder.String()
}

// FontStyles is a parsable slice of FontStyle.
type FontStyles []*FontStyle

func (fs FontStyles) config(cfg config) {
	for _i := range fs {
		fs[_i].config = cfg
	}
}
