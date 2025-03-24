/*
API Manager Contract API

API Manager Contract API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim_contract

import (
	"encoding/json"
)

// checks if the ApiReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiReference{}

// ApiReference struct for ApiReference
type ApiReference struct {
	Audit *Audit `json:"audit,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty"`
	Id *int32 `json:"id,omitempty"`
}

// NewApiReference instantiates a new ApiReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiReference() *ApiReference {
	this := ApiReference{}
	return &this
}

// NewApiReferenceWithDefaults instantiates a new ApiReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiReferenceWithDefaults() *ApiReference {
	this := ApiReference{}
	return &this
}

// GetAudit returns the Audit field value if set, zero value otherwise.
func (o *ApiReference) GetAudit() Audit {
	if o == nil || IsNil(o.Audit) {
		var ret Audit
		return ret
	}
	return *o.Audit
}

// GetAuditOk returns a tuple with the Audit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReference) GetAuditOk() (*Audit, bool) {
	if o == nil || IsNil(o.Audit) {
		return nil, false
	}
	return o.Audit, true
}

// HasAudit returns a boolean if a field has been set.
func (o *ApiReference) HasAudit() bool {
	if o != nil && !IsNil(o.Audit) {
		return true
	}

	return false
}

// SetAudit gets a reference to the given Audit and assigns it to the Audit field.
func (o *ApiReference) SetAudit(v Audit) {
	o.Audit = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ApiReference) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReference) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ApiReference) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *ApiReference) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiReference) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReference) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiReference) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ApiReference) SetId(v int32) {
	o.Id = &v
}

func (o ApiReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Audit) {
		toSerialize["audit"] = o.Audit
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableApiReference struct {
	value *ApiReference
	isSet bool
}

func (v NullableApiReference) Get() *ApiReference {
	return v.value
}

func (v *NullableApiReference) Set(val *ApiReference) {
	v.value = val
	v.isSet = true
}

func (v NullableApiReference) IsSet() bool {
	return v.isSet
}

func (v *NullableApiReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiReference(val *ApiReference) *NullableApiReference {
	return &NullableApiReference{value: val, isSet: true}
}

func (v NullableApiReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


