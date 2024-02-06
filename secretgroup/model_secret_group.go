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

// checks if the SecretGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretGroup{}

// SecretGroup struct for SecretGroup
type SecretGroup struct {
	Name *string `json:"name,omitempty"`
	Downloadable *bool `json:"downloadable,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
}

// NewSecretGroup instantiates a new SecretGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretGroup() *SecretGroup {
	this := SecretGroup{}
	return &this
}

// NewSecretGroupWithDefaults instantiates a new SecretGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretGroupWithDefaults() *SecretGroup {
	this := SecretGroup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecretGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecretGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecretGroup) SetName(v string) {
	o.Name = &v
}

// GetDownloadable returns the Downloadable field value if set, zero value otherwise.
func (o *SecretGroup) GetDownloadable() bool {
	if o == nil || IsNil(o.Downloadable) {
		var ret bool
		return ret
	}
	return *o.Downloadable
}

// GetDownloadableOk returns a tuple with the Downloadable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretGroup) GetDownloadableOk() (*bool, bool) {
	if o == nil || IsNil(o.Downloadable) {
		return nil, false
	}
	return o.Downloadable, true
}

// HasDownloadable returns a boolean if a field has been set.
func (o *SecretGroup) HasDownloadable() bool {
	if o != nil && !IsNil(o.Downloadable) {
		return true
	}

	return false
}

// SetDownloadable gets a reference to the given bool and assigns it to the Downloadable field.
func (o *SecretGroup) SetDownloadable(v bool) {
	o.Downloadable = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SecretGroup) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretGroup) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SecretGroup) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *SecretGroup) SetMeta(v Meta) {
	o.Meta = &v
}

func (o SecretGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Downloadable) {
		toSerialize["downloadable"] = o.Downloadable
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableSecretGroup struct {
	value *SecretGroup
	isSet bool
}

func (v NullableSecretGroup) Get() *SecretGroup {
	return v.value
}

func (v *NullableSecretGroup) Set(val *SecretGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretGroup(val *SecretGroup) *NullableSecretGroup {
	return &NullableSecretGroup{value: val, isSet: true}
}

func (v NullableSecretGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


