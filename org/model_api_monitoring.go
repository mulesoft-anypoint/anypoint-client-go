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

// checks if the ApiMonitoring type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMonitoring{}

// ApiMonitoring API Monitoring entitlements
type ApiMonitoring struct {
	// Schedules
	Schedules *int32 `json:"schedules,omitempty"`
}

// NewApiMonitoring instantiates a new ApiMonitoring object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMonitoring() *ApiMonitoring {
	this := ApiMonitoring{}
	var schedules int32 = 0
	this.Schedules = &schedules
	return &this
}

// NewApiMonitoringWithDefaults instantiates a new ApiMonitoring object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMonitoringWithDefaults() *ApiMonitoring {
	this := ApiMonitoring{}
	var schedules int32 = 0
	this.Schedules = &schedules
	return &this
}

// GetSchedules returns the Schedules field value if set, zero value otherwise.
func (o *ApiMonitoring) GetSchedules() int32 {
	if o == nil || IsNil(o.Schedules) {
		var ret int32
		return ret
	}
	return *o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMonitoring) GetSchedulesOk() (*int32, bool) {
	if o == nil || IsNil(o.Schedules) {
		return nil, false
	}
	return o.Schedules, true
}

// HasSchedules returns a boolean if a field has been set.
func (o *ApiMonitoring) HasSchedules() bool {
	if o != nil && !IsNil(o.Schedules) {
		return true
	}

	return false
}

// SetSchedules gets a reference to the given int32 and assigns it to the Schedules field.
func (o *ApiMonitoring) SetSchedules(v int32) {
	o.Schedules = &v
}

func (o ApiMonitoring) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMonitoring) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schedules) {
		toSerialize["schedules"] = o.Schedules
	}
	return toSerialize, nil
}

type NullableApiMonitoring struct {
	value *ApiMonitoring
	isSet bool
}

func (v NullableApiMonitoring) Get() *ApiMonitoring {
	return v.value
}

func (v *NullableApiMonitoring) Set(val *ApiMonitoring) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMonitoring) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMonitoring) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMonitoring(val *ApiMonitoring) *NullableApiMonitoring {
	return &NullableApiMonitoring{value: val, isSet: true}
}

func (v NullableApiMonitoring) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMonitoring) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


