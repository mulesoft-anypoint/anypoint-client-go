/*
Organization API

Description of the Organization API

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package org

import (
	"encoding/json"
)

// checks if the PartnersSandboxEntitlement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartnersSandboxEntitlement{}

// PartnersSandboxEntitlement Partners Sandbox entitlements
type PartnersSandboxEntitlement struct {
	// Partners Sandbox assigned
	Assigned *int32 `json:"assigned,omitempty"`
}

// NewPartnersSandboxEntitlement instantiates a new PartnersSandboxEntitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnersSandboxEntitlement() *PartnersSandboxEntitlement {
	this := PartnersSandboxEntitlement{}
	var assigned int32 = 0
	this.Assigned = &assigned
	return &this
}

// NewPartnersSandboxEntitlementWithDefaults instantiates a new PartnersSandboxEntitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnersSandboxEntitlementWithDefaults() *PartnersSandboxEntitlement {
	this := PartnersSandboxEntitlement{}
	var assigned int32 = 0
	this.Assigned = &assigned
	return &this
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *PartnersSandboxEntitlement) GetAssigned() int32 {
	if o == nil || IsNil(o.Assigned) {
		var ret int32
		return ret
	}
	return *o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnersSandboxEntitlement) GetAssignedOk() (*int32, bool) {
	if o == nil || IsNil(o.Assigned) {
		return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *PartnersSandboxEntitlement) HasAssigned() bool {
	if o != nil && !IsNil(o.Assigned) {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given int32 and assigns it to the Assigned field.
func (o *PartnersSandboxEntitlement) SetAssigned(v int32) {
	o.Assigned = &v
}

func (o PartnersSandboxEntitlement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartnersSandboxEntitlement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assigned) {
		toSerialize["assigned"] = o.Assigned
	}
	return toSerialize, nil
}

type NullablePartnersSandboxEntitlement struct {
	value *PartnersSandboxEntitlement
	isSet bool
}

func (v NullablePartnersSandboxEntitlement) Get() *PartnersSandboxEntitlement {
	return v.value
}

func (v *NullablePartnersSandboxEntitlement) Set(val *PartnersSandboxEntitlement) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnersSandboxEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnersSandboxEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnersSandboxEntitlement(val *PartnersSandboxEntitlement) *NullablePartnersSandboxEntitlement {
	return &NullablePartnersSandboxEntitlement{value: val, isSet: true}
}

func (v NullablePartnersSandboxEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnersSandboxEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


