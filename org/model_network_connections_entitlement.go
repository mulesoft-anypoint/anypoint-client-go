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

// checks if the NetworkConnectionsEntitlement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkConnectionsEntitlement{}

// NetworkConnectionsEntitlement Network connections entitlements
type NetworkConnectionsEntitlement struct {
	// Network connections assigned
	Assigned *int32 `json:"assigned,omitempty"`
	// Network connections reassigned
	Reassigned *int32 `json:"reassigned,omitempty"`
}

// NewNetworkConnectionsEntitlement instantiates a new NetworkConnectionsEntitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkConnectionsEntitlement() *NetworkConnectionsEntitlement {
	this := NetworkConnectionsEntitlement{}
	return &this
}

// NewNetworkConnectionsEntitlementWithDefaults instantiates a new NetworkConnectionsEntitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkConnectionsEntitlementWithDefaults() *NetworkConnectionsEntitlement {
	this := NetworkConnectionsEntitlement{}
	return &this
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *NetworkConnectionsEntitlement) GetAssigned() int32 {
	if o == nil || IsNil(o.Assigned) {
		var ret int32
		return ret
	}
	return *o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConnectionsEntitlement) GetAssignedOk() (*int32, bool) {
	if o == nil || IsNil(o.Assigned) {
		return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *NetworkConnectionsEntitlement) HasAssigned() bool {
	if o != nil && !IsNil(o.Assigned) {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given int32 and assigns it to the Assigned field.
func (o *NetworkConnectionsEntitlement) SetAssigned(v int32) {
	o.Assigned = &v
}

// GetReassigned returns the Reassigned field value if set, zero value otherwise.
func (o *NetworkConnectionsEntitlement) GetReassigned() int32 {
	if o == nil || IsNil(o.Reassigned) {
		var ret int32
		return ret
	}
	return *o.Reassigned
}

// GetReassignedOk returns a tuple with the Reassigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConnectionsEntitlement) GetReassignedOk() (*int32, bool) {
	if o == nil || IsNil(o.Reassigned) {
		return nil, false
	}
	return o.Reassigned, true
}

// HasReassigned returns a boolean if a field has been set.
func (o *NetworkConnectionsEntitlement) HasReassigned() bool {
	if o != nil && !IsNil(o.Reassigned) {
		return true
	}

	return false
}

// SetReassigned gets a reference to the given int32 and assigns it to the Reassigned field.
func (o *NetworkConnectionsEntitlement) SetReassigned(v int32) {
	o.Reassigned = &v
}

func (o NetworkConnectionsEntitlement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkConnectionsEntitlement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assigned) {
		toSerialize["assigned"] = o.Assigned
	}
	if !IsNil(o.Reassigned) {
		toSerialize["reassigned"] = o.Reassigned
	}
	return toSerialize, nil
}

type NullableNetworkConnectionsEntitlement struct {
	value *NetworkConnectionsEntitlement
	isSet bool
}

func (v NullableNetworkConnectionsEntitlement) Get() *NetworkConnectionsEntitlement {
	return v.value
}

func (v *NullableNetworkConnectionsEntitlement) Set(val *NetworkConnectionsEntitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkConnectionsEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkConnectionsEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkConnectionsEntitlement(val *NetworkConnectionsEntitlement) *NullableNetworkConnectionsEntitlement {
	return &NullableNetworkConnectionsEntitlement{value: val, isSet: true}
}

func (v NullableNetworkConnectionsEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkConnectionsEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


