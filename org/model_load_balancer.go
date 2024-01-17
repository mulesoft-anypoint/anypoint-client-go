/*
Organization API

Description of the Organization API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package org

import (
	"encoding/json"
)

// checks if the LoadBalancer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancer{}

// LoadBalancer An explanation about the purpose of this instance.
type LoadBalancer struct {
	// An explanation about the purpose of this instance.
	Assigned int32 `json:"assigned"`
	// An explanation about the purpose of this instance.
	Reassigned *int32 `json:"reassigned,omitempty"`
}

// NewLoadBalancer instantiates a new LoadBalancer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancer(assigned int32) *LoadBalancer {
	this := LoadBalancer{}
	this.Assigned = assigned
	var reassigned int32 = 0
	this.Reassigned = &reassigned
	return &this
}

// NewLoadBalancerWithDefaults instantiates a new LoadBalancer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerWithDefaults() *LoadBalancer {
	this := LoadBalancer{}
	var assigned int32 = 0
	this.Assigned = assigned
	var reassigned int32 = 0
	this.Reassigned = &reassigned
	return &this
}

// GetAssigned returns the Assigned field value
func (o *LoadBalancer) GetAssigned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetAssignedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Assigned, true
}

// SetAssigned sets field value
func (o *LoadBalancer) SetAssigned(v int32) {
	o.Assigned = v
}

// GetReassigned returns the Reassigned field value if set, zero value otherwise.
func (o *LoadBalancer) GetReassigned() int32 {
	if o == nil || IsNil(o.Reassigned) {
		var ret int32
		return ret
	}
	return *o.Reassigned
}

// GetReassignedOk returns a tuple with the Reassigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetReassignedOk() (*int32, bool) {
	if o == nil || IsNil(o.Reassigned) {
		return nil, false
	}
	return o.Reassigned, true
}

// HasReassigned returns a boolean if a field has been set.
func (o *LoadBalancer) HasReassigned() bool {
	if o != nil && !IsNil(o.Reassigned) {
		return true
	}

	return false
}

// SetReassigned gets a reference to the given int32 and assigns it to the Reassigned field.
func (o *LoadBalancer) SetReassigned(v int32) {
	o.Reassigned = &v
}

func (o LoadBalancer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["assigned"] = o.Assigned
	if !IsNil(o.Reassigned) {
		toSerialize["reassigned"] = o.Reassigned
	}
	return toSerialize, nil
}

type NullableLoadBalancer struct {
	value *LoadBalancer
	isSet bool
}

func (v NullableLoadBalancer) Get() *LoadBalancer {
	return v.value
}

func (v *NullableLoadBalancer) Set(val *LoadBalancer) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancer) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancer(val *LoadBalancer) *NullableLoadBalancer {
	return &NullableLoadBalancer{value: val, isSet: true}
}

func (v NullableLoadBalancer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


