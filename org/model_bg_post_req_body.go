/*
 * Organization API
 *
 * Description of the Organization API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package org

import (
	"encoding/json"
)

// BGPostReqBody struct for BGPostReqBody
type BGPostReqBody struct {
	Entitlements EntitlementsCore `json:"entitlements"`
	// An explanation about the purpose of this instance.
	Name string `json:"name"`
	// An explanation about the purpose of this instance.
	OwnerId string `json:"ownerId"`
	// An explanation about the purpose of this instance.
	ParentOrganizationId string `json:"parentOrganizationId"`
}

// NewBGPostReqBody instantiates a new BGPostReqBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBGPostReqBody(entitlements EntitlementsCore, name string, ownerId string, parentOrganizationId string) *BGPostReqBody {
	this := BGPostReqBody{}
	this.Entitlements = entitlements
	this.Name = name
	this.OwnerId = ownerId
	this.ParentOrganizationId = parentOrganizationId
	return &this
}

// NewBGPostReqBodyWithDefaults instantiates a new BGPostReqBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBGPostReqBodyWithDefaults() *BGPostReqBody {
	this := BGPostReqBody{}
	var name string = ""
	this.Name = name
	var ownerId string = ""
	this.OwnerId = ownerId
	var parentOrganizationId string = ""
	this.ParentOrganizationId = parentOrganizationId
	return &this
}

// GetEntitlements returns the Entitlements field value
func (o *BGPostReqBody) GetEntitlements() EntitlementsCore {
	if o == nil {
		var ret EntitlementsCore
		return ret
	}

	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value
// and a boolean to check if the value has been set.
func (o *BGPostReqBody) GetEntitlementsOk() (*EntitlementsCore, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Entitlements, true
}

// SetEntitlements sets field value
func (o *BGPostReqBody) SetEntitlements(v EntitlementsCore) {
	o.Entitlements = v
}

// GetName returns the Name field value
func (o *BGPostReqBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BGPostReqBody) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BGPostReqBody) SetName(v string) {
	o.Name = v
}

// GetOwnerId returns the OwnerId field value
func (o *BGPostReqBody) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *BGPostReqBody) GetOwnerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *BGPostReqBody) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetParentOrganizationId returns the ParentOrganizationId field value
func (o *BGPostReqBody) GetParentOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentOrganizationId
}

// GetParentOrganizationIdOk returns a tuple with the ParentOrganizationId field value
// and a boolean to check if the value has been set.
func (o *BGPostReqBody) GetParentOrganizationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ParentOrganizationId, true
}

// SetParentOrganizationId sets field value
func (o *BGPostReqBody) SetParentOrganizationId(v string) {
	o.ParentOrganizationId = v
}

func (o BGPostReqBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entitlements"] = o.Entitlements
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["ownerId"] = o.OwnerId
	}
	if true {
		toSerialize["parentOrganizationId"] = o.ParentOrganizationId
	}
	return json.Marshal(toSerialize)
}

type NullableBGPostReqBody struct {
	value *BGPostReqBody
	isSet bool
}

func (v NullableBGPostReqBody) Get() *BGPostReqBody {
	return v.value
}

func (v *NullableBGPostReqBody) Set(val *BGPostReqBody) {
	v.value = val
	v.isSet = true
}

func (v NullableBGPostReqBody) IsSet() bool {
	return v.isSet
}

func (v *NullableBGPostReqBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBGPostReqBody(val *BGPostReqBody) *NullableBGPostReqBody {
	return &NullableBGPostReqBody{value: val, isSet: true}
}

func (v NullableBGPostReqBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBGPostReqBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


