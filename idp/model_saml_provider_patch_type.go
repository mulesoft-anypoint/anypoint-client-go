/*
Identity Provider Management API

Description of Identity Provider API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package idp

import (
	"encoding/json"
)

// checks if the SamlProviderPatchType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamlProviderPatchType{}

// SamlProviderPatchType struct for SamlProviderPatchType
type SamlProviderPatchType struct {
	Description *string `json:"description,omitempty"`
}

// NewSamlProviderPatchType instantiates a new SamlProviderPatchType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlProviderPatchType() *SamlProviderPatchType {
	this := SamlProviderPatchType{}
	return &this
}

// NewSamlProviderPatchTypeWithDefaults instantiates a new SamlProviderPatchType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlProviderPatchTypeWithDefaults() *SamlProviderPatchType {
	this := SamlProviderPatchType{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SamlProviderPatchType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlProviderPatchType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SamlProviderPatchType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SamlProviderPatchType) SetDescription(v string) {
	o.Description = &v
}

func (o SamlProviderPatchType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SamlProviderPatchType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableSamlProviderPatchType struct {
	value *SamlProviderPatchType
	isSet bool
}

func (v NullableSamlProviderPatchType) Get() *SamlProviderPatchType {
	return v.value
}

func (v *NullableSamlProviderPatchType) Set(val *SamlProviderPatchType) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlProviderPatchType) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlProviderPatchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlProviderPatchType(val *SamlProviderPatchType) *NullableSamlProviderPatchType {
	return &NullableSamlProviderPatchType{value: val, isSet: true}
}

func (v NullableSamlProviderPatchType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlProviderPatchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

