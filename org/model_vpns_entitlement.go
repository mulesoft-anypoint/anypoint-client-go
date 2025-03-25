/*
Organization API

Description of the Organization API

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package org

import (
	"encoding/json"
)

// checks if the VpnsEntitlement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VpnsEntitlement{}

// VpnsEntitlement struct for VpnsEntitlement
type VpnsEntitlement struct {
	// VPNs assigned
	Assigned *int32 `json:"assigned,omitempty"`
	// VPNs reassigned
	Reassigned *int32 `json:"reassigned,omitempty"`
}

// NewVpnsEntitlement instantiates a new VpnsEntitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVpnsEntitlement() *VpnsEntitlement {
	this := VpnsEntitlement{}
	var assigned int32 = 0
	this.Assigned = &assigned
	var reassigned int32 = 0
	this.Reassigned = &reassigned
	return &this
}

// NewVpnsEntitlementWithDefaults instantiates a new VpnsEntitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVpnsEntitlementWithDefaults() *VpnsEntitlement {
	this := VpnsEntitlement{}
	var assigned int32 = 0
	this.Assigned = &assigned
	var reassigned int32 = 0
	this.Reassigned = &reassigned
	return &this
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *VpnsEntitlement) GetAssigned() int32 {
	if o == nil || IsNil(o.Assigned) {
		var ret int32
		return ret
	}
	return *o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VpnsEntitlement) GetAssignedOk() (*int32, bool) {
	if o == nil || IsNil(o.Assigned) {
		return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *VpnsEntitlement) HasAssigned() bool {
	if o != nil && !IsNil(o.Assigned) {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given int32 and assigns it to the Assigned field.
func (o *VpnsEntitlement) SetAssigned(v int32) {
	o.Assigned = &v
}

// GetReassigned returns the Reassigned field value if set, zero value otherwise.
func (o *VpnsEntitlement) GetReassigned() int32 {
	if o == nil || IsNil(o.Reassigned) {
		var ret int32
		return ret
	}
	return *o.Reassigned
}

// GetReassignedOk returns a tuple with the Reassigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VpnsEntitlement) GetReassignedOk() (*int32, bool) {
	if o == nil || IsNil(o.Reassigned) {
		return nil, false
	}
	return o.Reassigned, true
}

// HasReassigned returns a boolean if a field has been set.
func (o *VpnsEntitlement) HasReassigned() bool {
	if o != nil && !IsNil(o.Reassigned) {
		return true
	}

	return false
}

// SetReassigned gets a reference to the given int32 and assigns it to the Reassigned field.
func (o *VpnsEntitlement) SetReassigned(v int32) {
	o.Reassigned = &v
}

func (o VpnsEntitlement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VpnsEntitlement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assigned) {
		toSerialize["assigned"] = o.Assigned
	}
	if !IsNil(o.Reassigned) {
		toSerialize["reassigned"] = o.Reassigned
	}
	return toSerialize, nil
}

type NullableVpnsEntitlement struct {
	value *VpnsEntitlement
	isSet bool
}

func (v NullableVpnsEntitlement) Get() *VpnsEntitlement {
	return v.value
}

func (v *NullableVpnsEntitlement) Set(val *VpnsEntitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableVpnsEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableVpnsEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVpnsEntitlement(val *VpnsEntitlement) *NullableVpnsEntitlement {
	return &NullableVpnsEntitlement{value: val, isSet: true}
}

func (v NullableVpnsEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVpnsEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


