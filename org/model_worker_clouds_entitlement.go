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

// checks if the WorkerCloudsEntitlement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkerCloudsEntitlement{}

// WorkerCloudsEntitlement struct for WorkerCloudsEntitlement
type WorkerCloudsEntitlement struct {
	// Worker Clouds assigned
	Assigned *int32 `json:"assigned,omitempty"`
	// Worker Clouds reassigned
	Reassigned *int32 `json:"reassigned,omitempty"`
}

// NewWorkerCloudsEntitlement instantiates a new WorkerCloudsEntitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkerCloudsEntitlement() *WorkerCloudsEntitlement {
	this := WorkerCloudsEntitlement{}
	var assigned int32 = 0
	this.Assigned = &assigned
	var reassigned int32 = 0
	this.Reassigned = &reassigned
	return &this
}

// NewWorkerCloudsEntitlementWithDefaults instantiates a new WorkerCloudsEntitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkerCloudsEntitlementWithDefaults() *WorkerCloudsEntitlement {
	this := WorkerCloudsEntitlement{}
	var assigned int32 = 0
	this.Assigned = &assigned
	var reassigned int32 = 0
	this.Reassigned = &reassigned
	return &this
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *WorkerCloudsEntitlement) GetAssigned() int32 {
	if o == nil || IsNil(o.Assigned) {
		var ret int32
		return ret
	}
	return *o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkerCloudsEntitlement) GetAssignedOk() (*int32, bool) {
	if o == nil || IsNil(o.Assigned) {
		return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *WorkerCloudsEntitlement) HasAssigned() bool {
	if o != nil && !IsNil(o.Assigned) {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given int32 and assigns it to the Assigned field.
func (o *WorkerCloudsEntitlement) SetAssigned(v int32) {
	o.Assigned = &v
}

// GetReassigned returns the Reassigned field value if set, zero value otherwise.
func (o *WorkerCloudsEntitlement) GetReassigned() int32 {
	if o == nil || IsNil(o.Reassigned) {
		var ret int32
		return ret
	}
	return *o.Reassigned
}

// GetReassignedOk returns a tuple with the Reassigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkerCloudsEntitlement) GetReassignedOk() (*int32, bool) {
	if o == nil || IsNil(o.Reassigned) {
		return nil, false
	}
	return o.Reassigned, true
}

// HasReassigned returns a boolean if a field has been set.
func (o *WorkerCloudsEntitlement) HasReassigned() bool {
	if o != nil && !IsNil(o.Reassigned) {
		return true
	}

	return false
}

// SetReassigned gets a reference to the given int32 and assigns it to the Reassigned field.
func (o *WorkerCloudsEntitlement) SetReassigned(v int32) {
	o.Reassigned = &v
}

func (o WorkerCloudsEntitlement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkerCloudsEntitlement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assigned) {
		toSerialize["assigned"] = o.Assigned
	}
	if !IsNil(o.Reassigned) {
		toSerialize["reassigned"] = o.Reassigned
	}
	return toSerialize, nil
}

type NullableWorkerCloudsEntitlement struct {
	value *WorkerCloudsEntitlement
	isSet bool
}

func (v NullableWorkerCloudsEntitlement) Get() *WorkerCloudsEntitlement {
	return v.value
}

func (v *NullableWorkerCloudsEntitlement) Set(val *WorkerCloudsEntitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkerCloudsEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkerCloudsEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkerCloudsEntitlement(val *WorkerCloudsEntitlement) *NullableWorkerCloudsEntitlement {
	return &NullableWorkerCloudsEntitlement{value: val, isSet: true}
}

func (v NullableWorkerCloudsEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkerCloudsEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


