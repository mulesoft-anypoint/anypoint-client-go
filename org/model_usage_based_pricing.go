/*
Organization API

Description of the Organization API

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package org

import (
	"encoding/json"
)

// checks if the UsageBasedPricing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageBasedPricing{}

// UsageBasedPricing struct for UsageBasedPricing
type UsageBasedPricing struct {
	Api *Api `json:"api,omitempty"`
}

// NewUsageBasedPricing instantiates a new UsageBasedPricing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageBasedPricing() *UsageBasedPricing {
	this := UsageBasedPricing{}
	return &this
}

// NewUsageBasedPricingWithDefaults instantiates a new UsageBasedPricing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageBasedPricingWithDefaults() *UsageBasedPricing {
	this := UsageBasedPricing{}
	return &this
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *UsageBasedPricing) GetApi() Api {
	if o == nil || IsNil(o.Api) {
		var ret Api
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBasedPricing) GetApiOk() (*Api, bool) {
	if o == nil || IsNil(o.Api) {
		return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *UsageBasedPricing) HasApi() bool {
	if o != nil && !IsNil(o.Api) {
		return true
	}

	return false
}

// SetApi gets a reference to the given Api and assigns it to the Api field.
func (o *UsageBasedPricing) SetApi(v Api) {
	o.Api = &v
}

func (o UsageBasedPricing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageBasedPricing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Api) {
		toSerialize["api"] = o.Api
	}
	return toSerialize, nil
}

type NullableUsageBasedPricing struct {
	value *UsageBasedPricing
	isSet bool
}

func (v NullableUsageBasedPricing) Get() *UsageBasedPricing {
	return v.value
}

func (v *NullableUsageBasedPricing) Set(val *UsageBasedPricing) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageBasedPricing) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageBasedPricing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageBasedPricing(val *UsageBasedPricing) *NullableUsageBasedPricing {
	return &NullableUsageBasedPricing{value: val, isSet: true}
}

func (v NullableUsageBasedPricing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageBasedPricing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


