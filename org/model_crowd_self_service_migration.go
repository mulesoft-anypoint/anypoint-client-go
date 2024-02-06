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

// checks if the CrowdSelfServiceMigration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CrowdSelfServiceMigration{}

// CrowdSelfServiceMigration An explanation about the purpose of this instance.
type CrowdSelfServiceMigration struct {
	// An explanation about the purpose of this instance.
	Enabled bool `json:"enabled"`
}

// NewCrowdSelfServiceMigration instantiates a new CrowdSelfServiceMigration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrowdSelfServiceMigration(enabled bool) *CrowdSelfServiceMigration {
	this := CrowdSelfServiceMigration{}
	this.Enabled = enabled
	return &this
}

// NewCrowdSelfServiceMigrationWithDefaults instantiates a new CrowdSelfServiceMigration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrowdSelfServiceMigrationWithDefaults() *CrowdSelfServiceMigration {
	this := CrowdSelfServiceMigration{}
	var enabled bool = false
	this.Enabled = enabled
	return &this
}

// GetEnabled returns the Enabled field value
func (o *CrowdSelfServiceMigration) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CrowdSelfServiceMigration) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CrowdSelfServiceMigration) SetEnabled(v bool) {
	o.Enabled = v
}

func (o CrowdSelfServiceMigration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CrowdSelfServiceMigration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableCrowdSelfServiceMigration struct {
	value *CrowdSelfServiceMigration
	isSet bool
}

func (v NullableCrowdSelfServiceMigration) Get() *CrowdSelfServiceMigration {
	return v.value
}

func (v *NullableCrowdSelfServiceMigration) Set(val *CrowdSelfServiceMigration) {
	v.value = val
	v.isSet = true
}

func (v NullableCrowdSelfServiceMigration) IsSet() bool {
	return v.isSet
}

func (v *NullableCrowdSelfServiceMigration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCrowdSelfServiceMigration(val *CrowdSelfServiceMigration) *NullableCrowdSelfServiceMigration {
	return &NullableCrowdSelfServiceMigration{value: val, isSet: true}
}

func (v NullableCrowdSelfServiceMigration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrowdSelfServiceMigration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


