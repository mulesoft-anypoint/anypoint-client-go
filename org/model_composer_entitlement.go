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

// checks if the ComposerEntitlement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComposerEntitlement{}

// ComposerEntitlement struct for ComposerEntitlement
type ComposerEntitlement struct {
	// Composer enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Tasks per month
	TasksPerMonth *int32 `json:"tasksPerMonth,omitempty"`
	// Max connectors
	MaxConnectors *int32 `json:"maxConnectors,omitempty"`
	// Unlimited connectors
	UnlimitedConnectors *bool `json:"unlimitedConnectors,omitempty"`
	// Hyper Automation
	IsHyperAutomation *bool `json:"isHyperAutomation,omitempty"`
}

// NewComposerEntitlement instantiates a new ComposerEntitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComposerEntitlement() *ComposerEntitlement {
	this := ComposerEntitlement{}
	return &this
}

// NewComposerEntitlementWithDefaults instantiates a new ComposerEntitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComposerEntitlementWithDefaults() *ComposerEntitlement {
	this := ComposerEntitlement{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ComposerEntitlement) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComposerEntitlement) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ComposerEntitlement) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ComposerEntitlement) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTasksPerMonth returns the TasksPerMonth field value if set, zero value otherwise.
func (o *ComposerEntitlement) GetTasksPerMonth() int32 {
	if o == nil || IsNil(o.TasksPerMonth) {
		var ret int32
		return ret
	}
	return *o.TasksPerMonth
}

// GetTasksPerMonthOk returns a tuple with the TasksPerMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComposerEntitlement) GetTasksPerMonthOk() (*int32, bool) {
	if o == nil || IsNil(o.TasksPerMonth) {
		return nil, false
	}
	return o.TasksPerMonth, true
}

// HasTasksPerMonth returns a boolean if a field has been set.
func (o *ComposerEntitlement) HasTasksPerMonth() bool {
	if o != nil && !IsNil(o.TasksPerMonth) {
		return true
	}

	return false
}

// SetTasksPerMonth gets a reference to the given int32 and assigns it to the TasksPerMonth field.
func (o *ComposerEntitlement) SetTasksPerMonth(v int32) {
	o.TasksPerMonth = &v
}

// GetMaxConnectors returns the MaxConnectors field value if set, zero value otherwise.
func (o *ComposerEntitlement) GetMaxConnectors() int32 {
	if o == nil || IsNil(o.MaxConnectors) {
		var ret int32
		return ret
	}
	return *o.MaxConnectors
}

// GetMaxConnectorsOk returns a tuple with the MaxConnectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComposerEntitlement) GetMaxConnectorsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxConnectors) {
		return nil, false
	}
	return o.MaxConnectors, true
}

// HasMaxConnectors returns a boolean if a field has been set.
func (o *ComposerEntitlement) HasMaxConnectors() bool {
	if o != nil && !IsNil(o.MaxConnectors) {
		return true
	}

	return false
}

// SetMaxConnectors gets a reference to the given int32 and assigns it to the MaxConnectors field.
func (o *ComposerEntitlement) SetMaxConnectors(v int32) {
	o.MaxConnectors = &v
}

// GetUnlimitedConnectors returns the UnlimitedConnectors field value if set, zero value otherwise.
func (o *ComposerEntitlement) GetUnlimitedConnectors() bool {
	if o == nil || IsNil(o.UnlimitedConnectors) {
		var ret bool
		return ret
	}
	return *o.UnlimitedConnectors
}

// GetUnlimitedConnectorsOk returns a tuple with the UnlimitedConnectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComposerEntitlement) GetUnlimitedConnectorsOk() (*bool, bool) {
	if o == nil || IsNil(o.UnlimitedConnectors) {
		return nil, false
	}
	return o.UnlimitedConnectors, true
}

// HasUnlimitedConnectors returns a boolean if a field has been set.
func (o *ComposerEntitlement) HasUnlimitedConnectors() bool {
	if o != nil && !IsNil(o.UnlimitedConnectors) {
		return true
	}

	return false
}

// SetUnlimitedConnectors gets a reference to the given bool and assigns it to the UnlimitedConnectors field.
func (o *ComposerEntitlement) SetUnlimitedConnectors(v bool) {
	o.UnlimitedConnectors = &v
}

// GetIsHyperAutomation returns the IsHyperAutomation field value if set, zero value otherwise.
func (o *ComposerEntitlement) GetIsHyperAutomation() bool {
	if o == nil || IsNil(o.IsHyperAutomation) {
		var ret bool
		return ret
	}
	return *o.IsHyperAutomation
}

// GetIsHyperAutomationOk returns a tuple with the IsHyperAutomation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComposerEntitlement) GetIsHyperAutomationOk() (*bool, bool) {
	if o == nil || IsNil(o.IsHyperAutomation) {
		return nil, false
	}
	return o.IsHyperAutomation, true
}

// HasIsHyperAutomation returns a boolean if a field has been set.
func (o *ComposerEntitlement) HasIsHyperAutomation() bool {
	if o != nil && !IsNil(o.IsHyperAutomation) {
		return true
	}

	return false
}

// SetIsHyperAutomation gets a reference to the given bool and assigns it to the IsHyperAutomation field.
func (o *ComposerEntitlement) SetIsHyperAutomation(v bool) {
	o.IsHyperAutomation = &v
}

func (o ComposerEntitlement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComposerEntitlement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.TasksPerMonth) {
		toSerialize["tasksPerMonth"] = o.TasksPerMonth
	}
	if !IsNil(o.MaxConnectors) {
		toSerialize["maxConnectors"] = o.MaxConnectors
	}
	if !IsNil(o.UnlimitedConnectors) {
		toSerialize["unlimitedConnectors"] = o.UnlimitedConnectors
	}
	if !IsNil(o.IsHyperAutomation) {
		toSerialize["isHyperAutomation"] = o.IsHyperAutomation
	}
	return toSerialize, nil
}

type NullableComposerEntitlement struct {
	value *ComposerEntitlement
	isSet bool
}

func (v NullableComposerEntitlement) Get() *ComposerEntitlement {
	return v.value
}

func (v *NullableComposerEntitlement) Set(val *ComposerEntitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableComposerEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableComposerEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComposerEntitlement(val *ComposerEntitlement) *NullableComposerEntitlement {
	return &NullableComposerEntitlement{value: val, isSet: true}
}

func (v NullableComposerEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComposerEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


