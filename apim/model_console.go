/*
API Manager API

API Manager API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim

import (
	"encoding/json"
)

// checks if the Console type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Console{}

// Console struct for Console
type Console struct {
	Enabled *bool `json:"enabled,omitempty"`
	Path *string `json:"path,omitempty"`
}

// NewConsole instantiates a new Console object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsole() *Console {
	this := Console{}
	return &this
}

// NewConsoleWithDefaults instantiates a new Console object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsoleWithDefaults() *Console {
	this := Console{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Console) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Console) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Console) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Console) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *Console) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Console) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *Console) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *Console) SetPath(v string) {
	o.Path = &v
}

func (o Console) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Console) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	return toSerialize, nil
}

type NullableConsole struct {
	value *Console
	isSet bool
}

func (v NullableConsole) Get() *Console {
	return v.value
}

func (v *NullableConsole) Set(val *Console) {
	v.value = val
	v.isSet = true
}

func (v NullableConsole) IsSet() bool {
	return v.isSet
}

func (v *NullableConsole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsole(val *Console) *NullableConsole {
	return &NullableConsole{value: val, isSet: true}
}

func (v NullableConsole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


