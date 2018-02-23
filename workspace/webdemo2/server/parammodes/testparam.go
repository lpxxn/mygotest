package parammodes

import (
	"time"
)

// NoParams is a struct that exists in a package
// but is not annotated with the swagger params annotations
// so it should now show up in a test
//
// swagger:parameters someOperation anotherOperation
type NoParams struct {
	// ID of this no model instance.
	// ids in this application start at 11 and are smaller than 1000
	//
	// required: true
	// minimum: > 10
	// maximum: < 1000
	// in: path
	// default: 1
	ID int64 `json:"id"`

	// The Score of this model
	//
	// required: true
	// minimum: 3
	// maximum: 45
	// multiple of: 3
	// in: query
	// default: 2
	// example: 27
	Score int32 `json:"score"`

	// Created holds the time when this entry was created
	//
	// required: false
	// in: query
	Created time.Time `json:"created"`

	// The Category of this model
	//
	// required: true
	// enum: foo,bar,none
	// default: bar
	// in: query
	Category string `json:"category"`

	// a FooSlice has foos which are strings
	//
	// min items: 3
	// max items: 10
	// unique: true
	// items.minLength: 3
	// items.maxLength: 10
	// items.pattern: \w+
	// collection format: pipe
	// items.default: bar
	// in: query
	FooSlice []string `json:"foo_slice"`

	// a BarSlice has bars which are strings
	//
	// min items: 3
	// max items: 10
	// unique: true
	// items.minItems: 4
	// items.maxItems: 9
	// items.enum: bar1,bar2,bar3
	// items.default: bar2
	// items.items.minItems: 5
	// items.items.maxItems: 8
	// items.items.items.minLength: 3
	// items.items.items.maxLength: 10
	// items.items.items.pattern: \w+
	// collection format: pipe
	// in: query
	BarSlice [][][]string `json:"bar_slice"`
}

type Pet struct {
	Name string `json:"name"`
}
