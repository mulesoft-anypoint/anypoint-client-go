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

// checks if the VCoresDesignEntitlement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VCoresDesignEntitlement{}

// VCoresDesignEntitlement struct for VCoresDesignEntitlement
type VCoresDesignEntitlement struct {
	// VCores Design assigned
	Assigned *float32 `json:"assigned,omitempty"`
	// VCores Design reassigned
	Reassigned *float32 `json:"reassigned,omitempty"`
}

// NewVCoresDesignEntitlement instantiates a new VCoresDesignEntitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVCoresDesignEntitlement() *VCoresDesignEntitlement {
	this := VCoresDesignEntitlement{}
	var assigned float32 = 0
	this.Assigned = &assigned
	var reassigned float32 = 0
	this.Reassigned = &reassigned
	return &this
}

// NewVCoresDesignEntitlementWithDefaults instantiates a new VCoresDesignEntitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVCoresDesignEntitlementWithDefaults() *VCoresDesignEntitlement {
	this := VCoresDesignEntitlement{}
	var assigned float32 = 0
	this.Assigned = &assigned
	var reassigned float32 = 0
	this.Reassigned = &reassigned
	return &this
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *VCoresDesignEntitlement) GetAssigned() float32 {
	if o == nil || IsNil(o.Assigned) {
		var ret float32
		return ret
	}
	return *o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VCoresDesignEntitlement) GetAssignedOk() (*float32, bool) {
	if o == nil || IsNil(o.Assigned) {
		return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *VCoresDesignEntitlement) HasAssigned() bool {
	if o != nil && !IsNil(o.Assigned) {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given float32 and assigns it to the Assigned field.
func (o *VCoresDesignEntitlement) SetAssigned(v float32) {
	o.Assigned = &v
}

// GetReassigned returns the Reassigned field value if set, zero value otherwise.
func (o *VCoresDesignEntitlement) GetReassigned() float32 {
	if o == nil || IsNil(o.Reassigned) {
		var ret float32
		return ret
	}
	return *o.Reassigned
}

// GetReassignedOk returns a tuple with the Reassigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VCoresDesignEntitlement) GetReassignedOk() (*float32, bool) {
	if o == nil || IsNil(o.Reassigned) {
		return nil, false
	}
	return o.Reassigned, true
}

// HasReassigned returns a boolean if a field has been set.
func (o *VCoresDesignEntitlement) HasReassigned() bool {
	if o != nil && !IsNil(o.Reassigned) {
		return true
	}

	return false
}

// SetReassigned gets a reference to the given float32 and assigns it to the Reassigned field.
func (o *VCoresDesignEntitlement) SetReassigned(v float32) {
	o.Reassigned = &v
}

func (o VCoresDesignEntitlement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VCoresDesignEntitlement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assigned) {
		toSerialize["assigned"] = o.Assigned
	}
	if !IsNil(o.Reassigned) {
		toSerialize["reassigned"] = o.Reassigned
	}
	return toSerialize, nil
}

type NullableVCoresDesignEntitlement struct {
	value *VCoresDesignEntitlement
	isSet bool
}

func (v NullableVCoresDesignEntitlement) Get() *VCoresDesignEntitlement {
	return v.value
}

func (v *NullableVCoresDesignEntitlement) Set(val *VCoresDesignEntitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableVCoresDesignEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableVCoresDesignEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVCoresDesignEntitlement(val *VCoresDesignEntitlement) *NullableVCoresDesignEntitlement {
	return &NullableVCoresDesignEntitlement{value: val, isSet: true}
}

func (v NullableVCoresDesignEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVCoresDesignEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


