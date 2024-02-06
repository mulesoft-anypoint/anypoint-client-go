/*
Secret Group CRL Distributor Configs API

Secret Group CRL Distributor Configs API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package secretgroup_crl_distributor_configs

import (
	"encoding/json"
)

// checks if the SecretPath type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretPath{}

// SecretPath struct for SecretPath
type SecretPath struct {
	// Relative path of the secret to be referenced.
	Path *string `json:"path,omitempty"`
}

// NewSecretPath instantiates a new SecretPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretPath() *SecretPath {
	this := SecretPath{}
	return &this
}

// NewSecretPathWithDefaults instantiates a new SecretPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretPathWithDefaults() *SecretPath {
	this := SecretPath{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *SecretPath) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretPath) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *SecretPath) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *SecretPath) SetPath(v string) {
	o.Path = &v
}

func (o SecretPath) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretPath) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	return toSerialize, nil
}

type NullableSecretPath struct {
	value *SecretPath
	isSet bool
}

func (v NullableSecretPath) Get() *SecretPath {
	return v.value
}

func (v *NullableSecretPath) Set(val *SecretPath) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretPath) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretPath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretPath(val *SecretPath) *NullableSecretPath {
	return &NullableSecretPath{value: val, isSet: true}
}

func (v NullableSecretPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


