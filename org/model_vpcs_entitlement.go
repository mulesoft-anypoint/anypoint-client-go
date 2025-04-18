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

// checks if the VpcsEntitlement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VpcsEntitlement{}

// VpcsEntitlement struct for VpcsEntitlement
type VpcsEntitlement struct {
	// VPCs assigned
	Assigned *int32 `json:"assigned,omitempty"`
	// VPCs reassigned
	Reassigned *int32 `json:"reassigned,omitempty"`
}

// NewVpcsEntitlement instantiates a new VpcsEntitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVpcsEntitlement() *VpcsEntitlement {
	this := VpcsEntitlement{}
	var assigned int32 = 0
	this.Assigned = &assigned
	var reassigned int32 = 0
	this.Reassigned = &reassigned
	return &this
}

// NewVpcsEntitlementWithDefaults instantiates a new VpcsEntitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVpcsEntitlementWithDefaults() *VpcsEntitlement {
	this := VpcsEntitlement{}
	var assigned int32 = 0
	this.Assigned = &assigned
	var reassigned int32 = 0
	this.Reassigned = &reassigned
	return &this
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *VpcsEntitlement) GetAssigned() int32 {
	if o == nil || IsNil(o.Assigned) {
		var ret int32
		return ret
	}
	return *o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VpcsEntitlement) GetAssignedOk() (*int32, bool) {
	if o == nil || IsNil(o.Assigned) {
		return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *VpcsEntitlement) HasAssigned() bool {
	if o != nil && !IsNil(o.Assigned) {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given int32 and assigns it to the Assigned field.
func (o *VpcsEntitlement) SetAssigned(v int32) {
	o.Assigned = &v
}

// GetReassigned returns the Reassigned field value if set, zero value otherwise.
func (o *VpcsEntitlement) GetReassigned() int32 {
	if o == nil || IsNil(o.Reassigned) {
		var ret int32
		return ret
	}
	return *o.Reassigned
}

// GetReassignedOk returns a tuple with the Reassigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VpcsEntitlement) GetReassignedOk() (*int32, bool) {
	if o == nil || IsNil(o.Reassigned) {
		return nil, false
	}
	return o.Reassigned, true
}

// HasReassigned returns a boolean if a field has been set.
func (o *VpcsEntitlement) HasReassigned() bool {
	if o != nil && !IsNil(o.Reassigned) {
		return true
	}

	return false
}

// SetReassigned gets a reference to the given int32 and assigns it to the Reassigned field.
func (o *VpcsEntitlement) SetReassigned(v int32) {
	o.Reassigned = &v
}

func (o VpcsEntitlement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VpcsEntitlement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assigned) {
		toSerialize["assigned"] = o.Assigned
	}
	if !IsNil(o.Reassigned) {
		toSerialize["reassigned"] = o.Reassigned
	}
	return toSerialize, nil
}

type NullableVpcsEntitlement struct {
	value *VpcsEntitlement
	isSet bool
}

func (v NullableVpcsEntitlement) Get() *VpcsEntitlement {
	return v.value
}

func (v *NullableVpcsEntitlement) Set(val *VpcsEntitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableVpcsEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableVpcsEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVpcsEntitlement(val *VpcsEntitlement) *NullableVpcsEntitlement {
	return &NullableVpcsEntitlement{value: val, isSet: true}
}

func (v NullableVpcsEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVpcsEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


