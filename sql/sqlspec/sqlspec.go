// Copyright 2021-present The Atlas Authors. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sqlspec

import (
	"ariga.io/atlas/schemahcl"

	"github.com/zclconf/go-cty/cty"
)

type (
	// Schema holds a specification for a Schema.
	Schema struct {
		Name      string      `spec:"name,name"`
		Functions []*Function `spec:"function"`
		schemahcl.DefaultExtension
	}

	// Table holds a specification for an SQL table.
	Table struct {
		Name        string         `spec:",name"`
		Qualifier   string         `spec:",qualifier"`
		Schema      *schemahcl.Ref `spec:"schema"`
		Columns     []*Column      `spec:"column"`
		PrimaryKey  *PrimaryKey    `spec:"primary_key"`
		ForeignKeys []*ForeignKey  `spec:"foreign_key"`
		Indexes     []*Index       `spec:"index"`
		Checks      []*Check       `spec:"check"`
		schemahcl.DefaultExtension
	}

	// View holds a specification for an SQL view.
	View struct {
		Name      string         `spec:",name"`
		Qualifier string         `spec:",qualifier"`
		Schema    *schemahcl.Ref `spec:"schema"`
		Columns   []*Column      `spec:"column"`
		// The definition is appended as additional attribute
		// by the spec creator to marshal it after the columns.
		schemahcl.DefaultExtension
	}

	// Column holds a specification for a column in an SQL table.
	Column struct {
		Name    string          `spec:",name"`
		Null    bool            `spec:"null"`
		Type    *schemahcl.Type `spec:"type"`
		Default cty.Value       `spec:"default"`
		schemahcl.DefaultExtension
	}

	// PrimaryKey holds a specification for the primary key of a table.
	PrimaryKey struct {
		Columns []*schemahcl.Ref `spec:"columns"`
		schemahcl.DefaultExtension
	}

	// Index holds a specification for the index key of a table.
	Index struct {
		Name             string           `spec:",name"`
		Unique           bool             `spec:"unique,omitempty"`
		Constrained      bool             `spec:"constrained,omitempty"`
		NullsNotDistinct bool             `spec:"nulls_not_distinct,omitempty"`
		Parts            []*IndexPart     `spec:"on"`
		Columns          []*schemahcl.Ref `spec:"columns"`
		schemahcl.DefaultExtension
	}

	// IndexPart holds a specification for the index key part.
	IndexPart struct {
		Desc   bool           `spec:"desc,omitempty"`
		Column *schemahcl.Ref `spec:"column"`
		Expr   string         `spec:"expr,omitempty"`
		schemahcl.DefaultExtension
	}

	// Check holds a specification for a check constraint on a table.
	Check struct {
		Name string `spec:",name"`
		Expr string `spec:"expr"`
		schemahcl.DefaultExtension
	}

	// ForeignKey holds a specification for the Foreign key of a table.
	ForeignKey struct {
		Symbol     string           `spec:",name"`
		Columns    []*schemahcl.Ref `spec:"columns"`
		RefColumns []*schemahcl.Ref `spec:"ref_columns"`
		OnUpdate   *schemahcl.Ref   `spec:"on_update"`
		OnDelete   *schemahcl.Ref   `spec:"on_delete"`
		schemahcl.DefaultExtension
	}

	Function struct {
		Name      string         `spec:",name"`
		Qualifier string         `spec:",qualifier"`
		Schema    *schemahcl.Ref `spec:"schema"`
		Args      string         `spec:"args"`
		Returns   string         `spec:"returns"`
		Language  string         `spec:"language"`
		// Definition string         `spec:"definition"`
		schemahcl.DefaultExtension
	}

	Trigger struct {
		On       *schemahcl.Ref `spec:"on"`
		Name     string         `spec:",name"`
		Type     string         `spec:"type"`
		Event    string         `spec:"event"`
		ForEach  string         `spec:"per"`
		OldTable string         `spec:"old_table"`
		NewTable string         `spec:"new_table"`
		Execute  *schemahcl.Ref `spec:"execute"`
		schemahcl.DefaultExtension
	}

	// Type represents a database agnostic column type.
	Type string
)

func init() {
	schemahcl.Register("view", &View{})
	schemahcl.Register("table", &Table{})
	schemahcl.Register("schema", &Schema{})
	schemahcl.Register("function", &Function{})
	schemahcl.Register("trigger", &Trigger{})
}
