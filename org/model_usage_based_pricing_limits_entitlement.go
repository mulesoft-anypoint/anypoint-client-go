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

// checks if the UsageBasedPricingLimitsEntitlement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageBasedPricingLimitsEntitlement{}

// UsageBasedPricingLimitsEntitlement struct for UsageBasedPricingLimitsEntitlement
type UsageBasedPricingLimitsEntitlement struct {
	// CPU
	Cpu *int32 `json:"cpu,omitempty"`
	// Memory
	Memory *int32 `json:"memory,omitempty"`
}

// NewUsageBasedPricingLimitsEntitlement instantiates a new UsageBasedPricingLimitsEntitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageBasedPricingLimitsEntitlement() *UsageBasedPricingLimitsEntitlement {
	this := UsageBasedPricingLimitsEntitlement{}
	var cpu int32 = 0
	this.Cpu = &cpu
	var memory int32 = 0
	this.Memory = &memory
	return &this
}

// NewUsageBasedPricingLimitsEntitlementWithDefaults instantiates a new UsageBasedPricingLimitsEntitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageBasedPricingLimitsEntitlementWithDefaults() *UsageBasedPricingLimitsEntitlement {
	this := UsageBasedPricingLimitsEntitlement{}
	var cpu int32 = 0
	this.Cpu = &cpu
	var memory int32 = 0
	this.Memory = &memory
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *UsageBasedPricingLimitsEntitlement) GetCpu() int32 {
	if o == nil || IsNil(o.Cpu) {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBasedPricingLimitsEntitlement) GetCpuOk() (*int32, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *UsageBasedPricingLimitsEntitlement) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *UsageBasedPricingLimitsEntitlement) SetCpu(v int32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *UsageBasedPricingLimitsEntitlement) GetMemory() int32 {
	if o == nil || IsNil(o.Memory) {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBasedPricingLimitsEntitlement) GetMemoryOk() (*int32, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *UsageBasedPricingLimitsEntitlement) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *UsageBasedPricingLimitsEntitlement) SetMemory(v int32) {
	o.Memory = &v
}

func (o UsageBasedPricingLimitsEntitlement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageBasedPricingLimitsEntitlement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	return toSerialize, nil
}

type NullableUsageBasedPricingLimitsEntitlement struct {
	value *UsageBasedPricingLimitsEntitlement
	isSet bool
}

func (v NullableUsageBasedPricingLimitsEntitlement) Get() *UsageBasedPricingLimitsEntitlement {
	return v.value
}

func (v *NullableUsageBasedPricingLimitsEntitlement) Set(val *UsageBasedPricingLimitsEntitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageBasedPricingLimitsEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageBasedPricingLimitsEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageBasedPricingLimitsEntitlement(val *UsageBasedPricingLimitsEntitlement) *NullableUsageBasedPricingLimitsEntitlement {
	return &NullableUsageBasedPricingLimitsEntitlement{value: val, isSet: true}
}

func (v NullableUsageBasedPricingLimitsEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageBasedPricingLimitsEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


