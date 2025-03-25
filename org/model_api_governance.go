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

// checks if the ApiGovernance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiGovernance{}

// ApiGovernance struct for ApiGovernance
type ApiGovernance struct {
	// API Governance enabled
	Enabled *bool `json:"enabled,omitempty"`
	// APIs per month
	ApisPerMonth *int32 `json:"apisPerMonth,omitempty"`
}

// NewApiGovernance instantiates a new ApiGovernance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiGovernance() *ApiGovernance {
	this := ApiGovernance{}
	var enabled bool = false
	this.Enabled = &enabled
	var apisPerMonth int32 = 20
	this.ApisPerMonth = &apisPerMonth
	return &this
}

// NewApiGovernanceWithDefaults instantiates a new ApiGovernance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiGovernanceWithDefaults() *ApiGovernance {
	this := ApiGovernance{}
	var enabled bool = false
	this.Enabled = &enabled
	var apisPerMonth int32 = 20
	this.ApisPerMonth = &apisPerMonth
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApiGovernance) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiGovernance) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApiGovernance) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApiGovernance) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetApisPerMonth returns the ApisPerMonth field value if set, zero value otherwise.
func (o *ApiGovernance) GetApisPerMonth() int32 {
	if o == nil || IsNil(o.ApisPerMonth) {
		var ret int32
		return ret
	}
	return *o.ApisPerMonth
}

// GetApisPerMonthOk returns a tuple with the ApisPerMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiGovernance) GetApisPerMonthOk() (*int32, bool) {
	if o == nil || IsNil(o.ApisPerMonth) {
		return nil, false
	}
	return o.ApisPerMonth, true
}

// HasApisPerMonth returns a boolean if a field has been set.
func (o *ApiGovernance) HasApisPerMonth() bool {
	if o != nil && !IsNil(o.ApisPerMonth) {
		return true
	}

	return false
}

// SetApisPerMonth gets a reference to the given int32 and assigns it to the ApisPerMonth field.
func (o *ApiGovernance) SetApisPerMonth(v int32) {
	o.ApisPerMonth = &v
}

func (o ApiGovernance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiGovernance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.ApisPerMonth) {
		toSerialize["apisPerMonth"] = o.ApisPerMonth
	}
	return toSerialize, nil
}

type NullableApiGovernance struct {
	value *ApiGovernance
	isSet bool
}

func (v NullableApiGovernance) Get() *ApiGovernance {
	return v.value
}

func (v *NullableApiGovernance) Set(val *ApiGovernance) {
	v.value = val
	v.isSet = true
}

func (v NullableApiGovernance) IsSet() bool {
	return v.isSet
}

func (v *NullableApiGovernance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiGovernance(val *ApiGovernance) *NullableApiGovernance {
	return &NullableApiGovernance{value: val, isSet: true}
}

func (v NullableApiGovernance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiGovernance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


