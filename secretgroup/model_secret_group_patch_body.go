/*
Secret Group API

Secret Group API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package secretgroup

import (
	"encoding/json"
)

// checks if the SecretGroupPatchBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretGroupPatchBody{}

// SecretGroupPatchBody struct for SecretGroupPatchBody
type SecretGroupPatchBody struct {
	Name *string `json:"name,omitempty"`
}

// NewSecretGroupPatchBody instantiates a new SecretGroupPatchBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretGroupPatchBody() *SecretGroupPatchBody {
	this := SecretGroupPatchBody{}
	return &this
}

// NewSecretGroupPatchBodyWithDefaults instantiates a new SecretGroupPatchBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretGroupPatchBodyWithDefaults() *SecretGroupPatchBody {
	this := SecretGroupPatchBody{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecretGroupPatchBody) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretGroupPatchBody) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecretGroupPatchBody) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecretGroupPatchBody) SetName(v string) {
	o.Name = &v
}

func (o SecretGroupPatchBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretGroupPatchBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableSecretGroupPatchBody struct {
	value *SecretGroupPatchBody
	isSet bool
}

func (v NullableSecretGroupPatchBody) Get() *SecretGroupPatchBody {
	return v.value
}

func (v *NullableSecretGroupPatchBody) Set(val *SecretGroupPatchBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretGroupPatchBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretGroupPatchBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretGroupPatchBody(val *SecretGroupPatchBody) *NullableSecretGroupPatchBody {
	return &NullableSecretGroupPatchBody{value: val, isSet: true}
}

func (v NullableSecretGroupPatchBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretGroupPatchBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


