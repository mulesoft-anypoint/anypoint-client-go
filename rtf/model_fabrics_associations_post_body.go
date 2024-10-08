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

// checks if the FabricsAssociationsPostBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricsAssociationsPostBody{}

// FabricsAssociationsPostBody struct for FabricsAssociationsPostBody
type FabricsAssociationsPostBody struct {
	Associations []FabricsAssociationsPostBodyAssociationsInner `json:"associations,omitempty"`
}

// NewFabricsAssociationsPostBody instantiates a new FabricsAssociationsPostBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricsAssociationsPostBody() *FabricsAssociationsPostBody {
	this := FabricsAssociationsPostBody{}
	return &this
}

// NewFabricsAssociationsPostBodyWithDefaults instantiates a new FabricsAssociationsPostBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricsAssociationsPostBodyWithDefaults() *FabricsAssociationsPostBody {
	this := FabricsAssociationsPostBody{}
	return &this
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *FabricsAssociationsPostBody) GetAssociations() []FabricsAssociationsPostBodyAssociationsInner {
	if o == nil || IsNil(o.Associations) {
		var ret []FabricsAssociationsPostBodyAssociationsInner
		return ret
	}
	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsAssociationsPostBody) GetAssociationsOk() ([]FabricsAssociationsPostBodyAssociationsInner, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *FabricsAssociationsPostBody) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given []FabricsAssociationsPostBodyAssociationsInner and assigns it to the Associations field.
func (o *FabricsAssociationsPostBody) SetAssociations(v []FabricsAssociationsPostBodyAssociationsInner) {
	o.Associations = v
}

func (o FabricsAssociationsPostBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricsAssociationsPostBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Associations) {
		toSerialize["associations"] = o.Associations
	}
	return toSerialize, nil
}

type NullableFabricsAssociationsPostBody struct {
	value *FabricsAssociationsPostBody
	isSet bool
}

func (v NullableFabricsAssociationsPostBody) Get() *FabricsAssociationsPostBody {
	return v.value
}

func (v *NullableFabricsAssociationsPostBody) Set(val *FabricsAssociationsPostBody) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricsAssociationsPostBody) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricsAssociationsPostBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricsAssociationsPostBody(val *FabricsAssociationsPostBody) *NullableFabricsAssociationsPostBody {
	return &NullableFabricsAssociationsPostBody{value: val, isSet: true}
}

func (v NullableFabricsAssociationsPostBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricsAssociationsPostBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


