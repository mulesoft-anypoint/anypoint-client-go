/*
Runtime Fabrics API

Runtime Fabrics API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rtf

import (
	"encoding/json"
)

// checks if the Features type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Features{}

// Features struct for Features
type Features struct {
	EnhancedSecurity *bool `json:"enhancedSecurity,omitempty"`
	PersistentStore *bool `json:"persistentStore,omitempty"`
}

// NewFeatures instantiates a new Features object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatures() *Features {
	this := Features{}
	return &this
}

// NewFeaturesWithDefaults instantiates a new Features object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeaturesWithDefaults() *Features {
	this := Features{}
	return &this
}

// GetEnhancedSecurity returns the EnhancedSecurity field value if set, zero value otherwise.
func (o *Features) GetEnhancedSecurity() bool {
	if o == nil || IsNil(o.EnhancedSecurity) {
		var ret bool
		return ret
	}
	return *o.EnhancedSecurity
}

// GetEnhancedSecurityOk returns a tuple with the EnhancedSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Features) GetEnhancedSecurityOk() (*bool, bool) {
	if o == nil || IsNil(o.EnhancedSecurity) {
		return nil, false
	}
	return o.EnhancedSecurity, true
}

// HasEnhancedSecurity returns a boolean if a field has been set.
func (o *Features) HasEnhancedSecurity() bool {
	if o != nil && !IsNil(o.EnhancedSecurity) {
		return true
	}

	return false
}

// SetEnhancedSecurity gets a reference to the given bool and assigns it to the EnhancedSecurity field.
func (o *Features) SetEnhancedSecurity(v bool) {
	o.EnhancedSecurity = &v
}

// GetPersistentStore returns the PersistentStore field value if set, zero value otherwise.
func (o *Features) GetPersistentStore() bool {
	if o == nil || IsNil(o.PersistentStore) {
		var ret bool
		return ret
	}
	return *o.PersistentStore
}

// GetPersistentStoreOk returns a tuple with the PersistentStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Features) GetPersistentStoreOk() (*bool, bool) {
	if o == nil || IsNil(o.PersistentStore) {
		return nil, false
	}
	return o.PersistentStore, true
}

// HasPersistentStore returns a boolean if a field has been set.
func (o *Features) HasPersistentStore() bool {
	if o != nil && !IsNil(o.PersistentStore) {
		return true
	}

	return false
}

// SetPersistentStore gets a reference to the given bool and assigns it to the PersistentStore field.
func (o *Features) SetPersistentStore(v bool) {
	o.PersistentStore = &v
}

func (o Features) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Features) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnhancedSecurity) {
		toSerialize["enhancedSecurity"] = o.EnhancedSecurity
	}
	if !IsNil(o.PersistentStore) {
		toSerialize["persistentStore"] = o.PersistentStore
	}
	return toSerialize, nil
}

type NullableFeatures struct {
	value *Features
	isSet bool
}

func (v NullableFeatures) Get() *Features {
	return v.value
}

func (v *NullableFeatures) Set(val *Features) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatures(val *Features) *NullableFeatures {
	return &NullableFeatures{value: val, isSet: true}
}

func (v NullableFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

