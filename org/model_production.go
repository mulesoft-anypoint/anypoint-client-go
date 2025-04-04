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

// checks if the Production type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Production{}

// Production struct for Production
type Production struct {
	// Amount
	Amount *int32 `json:"amount,omitempty"`
}

// NewProduction instantiates a new Production object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduction() *Production {
	this := Production{}
	var amount int32 = 0
	this.Amount = &amount
	return &this
}

// NewProductionWithDefaults instantiates a new Production object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductionWithDefaults() *Production {
	this := Production{}
	var amount int32 = 0
	this.Amount = &amount
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Production) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Production) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Production) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *Production) SetAmount(v int32) {
	o.Amount = &v
}

func (o Production) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Production) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullableProduction struct {
	value *Production
	isSet bool
}

func (v NullableProduction) Get() *Production {
	return v.value
}

func (v *NullableProduction) Set(val *Production) {
	v.value = val
	v.isSet = true
}

func (v NullableProduction) IsSet() bool {
	return v.isSet
}

func (v *NullableProduction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduction(val *Production) *NullableProduction {
	return &NullableProduction{value: val, isSet: true}
}

func (v NullableProduction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


