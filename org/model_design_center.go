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

// checks if the DesignCenter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DesignCenter{}

// DesignCenter Design Center entitlements
type DesignCenter struct {
	// API entitlements
	Api *bool `json:"api,omitempty"`
	// Mozart entitlements
	Mozart *bool `json:"mozart,omitempty"`
}

// NewDesignCenter instantiates a new DesignCenter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesignCenter() *DesignCenter {
	this := DesignCenter{}
	var api bool = false
	this.Api = &api
	var mozart bool = false
	this.Mozart = &mozart
	return &this
}

// NewDesignCenterWithDefaults instantiates a new DesignCenter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesignCenterWithDefaults() *DesignCenter {
	this := DesignCenter{}
	var api bool = false
	this.Api = &api
	var mozart bool = false
	this.Mozart = &mozart
	return &this
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *DesignCenter) GetApi() bool {
	if o == nil || IsNil(o.Api) {
		var ret bool
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesignCenter) GetApiOk() (*bool, bool) {
	if o == nil || IsNil(o.Api) {
		return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *DesignCenter) HasApi() bool {
	if o != nil && !IsNil(o.Api) {
		return true
	}

	return false
}

// SetApi gets a reference to the given bool and assigns it to the Api field.
func (o *DesignCenter) SetApi(v bool) {
	o.Api = &v
}

// GetMozart returns the Mozart field value if set, zero value otherwise.
func (o *DesignCenter) GetMozart() bool {
	if o == nil || IsNil(o.Mozart) {
		var ret bool
		return ret
	}
	return *o.Mozart
}

// GetMozartOk returns a tuple with the Mozart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesignCenter) GetMozartOk() (*bool, bool) {
	if o == nil || IsNil(o.Mozart) {
		return nil, false
	}
	return o.Mozart, true
}

// HasMozart returns a boolean if a field has been set.
func (o *DesignCenter) HasMozart() bool {
	if o != nil && !IsNil(o.Mozart) {
		return true
	}

	return false
}

// SetMozart gets a reference to the given bool and assigns it to the Mozart field.
func (o *DesignCenter) SetMozart(v bool) {
	o.Mozart = &v
}

func (o DesignCenter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DesignCenter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Api) {
		toSerialize["api"] = o.Api
	}
	if !IsNil(o.Mozart) {
		toSerialize["mozart"] = o.Mozart
	}
	return toSerialize, nil
}

type NullableDesignCenter struct {
	value *DesignCenter
	isSet bool
}

func (v NullableDesignCenter) Get() *DesignCenter {
	return v.value
}

func (v *NullableDesignCenter) Set(val *DesignCenter) {
	v.value = val
	v.isSet = true
}

func (v NullableDesignCenter) IsSet() bool {
	return v.isSet
}

func (v *NullableDesignCenter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesignCenter(val *DesignCenter) *NullableDesignCenter {
	return &NullableDesignCenter{value: val, isSet: true}
}

func (v NullableDesignCenter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesignCenter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


