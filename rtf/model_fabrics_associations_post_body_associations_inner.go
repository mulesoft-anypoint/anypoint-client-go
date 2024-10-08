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

// checks if the FabricsAssociationsPostBodyAssociationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricsAssociationsPostBodyAssociationsInner{}

// FabricsAssociationsPostBodyAssociationsInner struct for FabricsAssociationsPostBodyAssociationsInner
type FabricsAssociationsPostBodyAssociationsInner struct {
	OrganizationId *string `json:"organizationId,omitempty"`
	Environment *string `json:"environment,omitempty"`
}

// NewFabricsAssociationsPostBodyAssociationsInner instantiates a new FabricsAssociationsPostBodyAssociationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricsAssociationsPostBodyAssociationsInner() *FabricsAssociationsPostBodyAssociationsInner {
	this := FabricsAssociationsPostBodyAssociationsInner{}
	return &this
}

// NewFabricsAssociationsPostBodyAssociationsInnerWithDefaults instantiates a new FabricsAssociationsPostBodyAssociationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricsAssociationsPostBodyAssociationsInnerWithDefaults() *FabricsAssociationsPostBodyAssociationsInner {
	this := FabricsAssociationsPostBodyAssociationsInner{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *FabricsAssociationsPostBodyAssociationsInner) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsAssociationsPostBodyAssociationsInner) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *FabricsAssociationsPostBodyAssociationsInner) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *FabricsAssociationsPostBodyAssociationsInner) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *FabricsAssociationsPostBodyAssociationsInner) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsAssociationsPostBodyAssociationsInner) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *FabricsAssociationsPostBodyAssociationsInner) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *FabricsAssociationsPostBodyAssociationsInner) SetEnvironment(v string) {
	o.Environment = &v
}

func (o FabricsAssociationsPostBodyAssociationsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricsAssociationsPostBodyAssociationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	return toSerialize, nil
}

type NullableFabricsAssociationsPostBodyAssociationsInner struct {
	value *FabricsAssociationsPostBodyAssociationsInner
	isSet bool
}

func (v NullableFabricsAssociationsPostBodyAssociationsInner) Get() *FabricsAssociationsPostBodyAssociationsInner {
	return v.value
}

func (v *NullableFabricsAssociationsPostBodyAssociationsInner) Set(val *FabricsAssociationsPostBodyAssociationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricsAssociationsPostBodyAssociationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricsAssociationsPostBodyAssociationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricsAssociationsPostBodyAssociationsInner(val *FabricsAssociationsPostBodyAssociationsInner) *NullableFabricsAssociationsPostBodyAssociationsInner {
	return &NullableFabricsAssociationsPostBodyAssociationsInner{value: val, isSet: true}
}

func (v NullableFabricsAssociationsPostBodyAssociationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricsAssociationsPostBodyAssociationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


