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

// checks if the ApiGovernanceEntitlement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiGovernanceEntitlement{}

// ApiGovernanceEntitlement struct for ApiGovernanceEntitlement
type ApiGovernanceEntitlement struct {
	// API Governance enabled
	Enabled *bool `json:"enabled,omitempty"`
	// APIs per month
	ApisPerMonth *int32 `json:"apisPerMonth,omitempty"`
}

// NewApiGovernanceEntitlement instantiates a new ApiGovernanceEntitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiGovernanceEntitlement() *ApiGovernanceEntitlement {
	this := ApiGovernanceEntitlement{}
	return &this
}

// NewApiGovernanceEntitlementWithDefaults instantiates a new ApiGovernanceEntitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiGovernanceEntitlementWithDefaults() *ApiGovernanceEntitlement {
	this := ApiGovernanceEntitlement{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApiGovernanceEntitlement) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiGovernanceEntitlement) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApiGovernanceEntitlement) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApiGovernanceEntitlement) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetApisPerMonth returns the ApisPerMonth field value if set, zero value otherwise.
func (o *ApiGovernanceEntitlement) GetApisPerMonth() int32 {
	if o == nil || IsNil(o.ApisPerMonth) {
		var ret int32
		return ret
	}
	return *o.ApisPerMonth
}

// GetApisPerMonthOk returns a tuple with the ApisPerMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiGovernanceEntitlement) GetApisPerMonthOk() (*int32, bool) {
	if o == nil || IsNil(o.ApisPerMonth) {
		return nil, false
	}
	return o.ApisPerMonth, true
}

// HasApisPerMonth returns a boolean if a field has been set.
func (o *ApiGovernanceEntitlement) HasApisPerMonth() bool {
	if o != nil && !IsNil(o.ApisPerMonth) {
		return true
	}

	return false
}

// SetApisPerMonth gets a reference to the given int32 and assigns it to the ApisPerMonth field.
func (o *ApiGovernanceEntitlement) SetApisPerMonth(v int32) {
	o.ApisPerMonth = &v
}

func (o ApiGovernanceEntitlement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiGovernanceEntitlement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.ApisPerMonth) {
		toSerialize["apisPerMonth"] = o.ApisPerMonth
	}
	return toSerialize, nil
}

type NullableApiGovernanceEntitlement struct {
	value *ApiGovernanceEntitlement
	isSet bool
}

func (v NullableApiGovernanceEntitlement) Get() *ApiGovernanceEntitlement {
	return v.value
}

func (v *NullableApiGovernanceEntitlement) Set(val *ApiGovernanceEntitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableApiGovernanceEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableApiGovernanceEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiGovernanceEntitlement(val *ApiGovernanceEntitlement) *NullableApiGovernanceEntitlement {
	return &NullableApiGovernanceEntitlement{value: val, isSet: true}
}

func (v NullableApiGovernanceEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiGovernanceEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


