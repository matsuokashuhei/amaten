// Code generated by entc, DO NOT EDIT.

package ent

import (
	"matsuokashuhei/grocery-store/services/customer-service/ent/customer"
	"matsuokashuhei/grocery-store/services/customer-service/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	customerMixin := schema.Customer{}.Mixin()
	customerMixinFields0 := customerMixin[0].Fields()
	_ = customerMixinFields0
	customerFields := schema.Customer{}.Fields()
	_ = customerFields
	// customerDescCreateTime is the schema descriptor for create_time field.
	customerDescCreateTime := customerMixinFields0[0].Descriptor()
	// customer.DefaultCreateTime holds the default value on creation for the create_time field.
	customer.DefaultCreateTime = customerDescCreateTime.Default.(func() time.Time)
	// customerDescUpdateTime is the schema descriptor for update_time field.
	customerDescUpdateTime := customerMixinFields0[1].Descriptor()
	// customer.DefaultUpdateTime holds the default value on creation for the update_time field.
	customer.DefaultUpdateTime = customerDescUpdateTime.Default.(func() time.Time)
	// customer.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	customer.UpdateDefaultUpdateTime = customerDescUpdateTime.UpdateDefault.(func() time.Time)
}
