// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/entc/integration/customid/ent/link"
	"entgo.io/ent/entc/integration/customid/ent/schema"
	uuidc "entgo.io/ent/entc/integration/customid/uuidcompatible"
)

// Link is the model entity for the Link schema.
type Link struct {
	config `json:"-"`
	// ID of the ent.
	ID uuidc.UUIDC `json:"id,omitempty"`
	// LinkInformation holds the value of the "link_information" field.
	LinkInformation map[string]schema.LinkInformation `json:"link_information,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Link) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case link.FieldLinkInformation:
			values[i] = new([]byte)
		case link.FieldID:
			values[i] = new(uuidc.UUIDC)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Link", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Link fields.
func (l *Link) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case link.FieldID:
			if value, ok := values[i].(*uuidc.UUIDC); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				l.ID = *value
			}
		case link.FieldLinkInformation:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field link_information", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &l.LinkInformation); err != nil {
					return fmt.Errorf("unmarshal field link_information: %w", err)
				}
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Link.
// Note that you need to call Link.Unwrap() before calling this method if this Link
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Link) Update() *LinkUpdateOne {
	return NewLinkClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the Link entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Link) Unwrap() *Link {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Link is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Link) String() string {
	var builder strings.Builder
	builder.WriteString("Link(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("link_information=")
	builder.WriteString(fmt.Sprintf("%v", l.LinkInformation))
	builder.WriteByte(')')
	return builder.String()
}

// Links is a parsable slice of Link.
type Links []*Link
