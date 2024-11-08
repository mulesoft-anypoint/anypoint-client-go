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

// checks if the OpenIDProviderManualPostBodyType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenIDProviderManualPostBodyType{}

// OpenIDProviderManualPostBodyType struct for OpenIDProviderManualPostBodyType
type OpenIDProviderManualPostBodyType struct {
	Description *string `json:"description,omitempty"`
	Name string `json:"name"`
}

// NewOpenIDProviderManualPostBodyType instantiates a new OpenIDProviderManualPostBodyType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenIDProviderManualPostBodyType(name string) *OpenIDProviderManualPostBodyType {
	this := OpenIDProviderManualPostBodyType{}
	this.Name = name
	return &this
}

// NewOpenIDProviderManualPostBodyTypeWithDefaults instantiates a new OpenIDProviderManualPostBodyType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenIDProviderManualPostBodyTypeWithDefaults() *OpenIDProviderManualPostBodyType {
	this := OpenIDProviderManualPostBodyType{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OpenIDProviderManualPostBodyType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIDProviderManualPostBodyType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OpenIDProviderManualPostBodyType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OpenIDProviderManualPostBodyType) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *OpenIDProviderManualPostBodyType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OpenIDProviderManualPostBodyType) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OpenIDProviderManualPostBodyType) SetName(v string) {
	o.Name = v
}

func (o OpenIDProviderManualPostBodyType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenIDProviderManualPostBodyType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableOpenIDProviderManualPostBodyType struct {
	value *OpenIDProviderManualPostBodyType
	isSet bool
}

func (v NullableOpenIDProviderManualPostBodyType) Get() *OpenIDProviderManualPostBodyType {
	return v.value
}

func (v *NullableOpenIDProviderManualPostBodyType) Set(val *OpenIDProviderManualPostBodyType) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenIDProviderManualPostBodyType) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenIDProviderManualPostBodyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenIDProviderManualPostBodyType(val *OpenIDProviderManualPostBodyType) *NullableOpenIDProviderManualPostBodyType {
	return &NullableOpenIDProviderManualPostBodyType{value: val, isSet: true}
}

func (v NullableOpenIDProviderManualPostBodyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenIDProviderManualPostBodyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

