/*
Application Manager v2

The Application Manager API exists to provide Mule Application management operations from Anypoint Control Planes to any Runtime Plane, with capabilities that include:   - Deploying a Mule Application or API to a Mule Runtime   - Scaling up or down a running application   - Managing application settings (ie: properties)   - Deleting a Mule Application Deployment   - Changing artifact version or configurations of a deployment   - Starting, Stopping or Restarting applications   - etc. This API currently supports deployments to Runtime Fabric and CloudHub 2.0 targets only. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package application_manager_v2

import (
	"encoding/json"
)

// checks if the ObjectStoreV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectStoreV2{}

// ObjectStoreV2 struct for ObjectStoreV2
type ObjectStoreV2 struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// NewObjectStoreV2 instantiates a new ObjectStoreV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectStoreV2() *ObjectStoreV2 {
	this := ObjectStoreV2{}
	return &this
}

// NewObjectStoreV2WithDefaults instantiates a new ObjectStoreV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectStoreV2WithDefaults() *ObjectStoreV2 {
	this := ObjectStoreV2{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ObjectStoreV2) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoreV2) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ObjectStoreV2) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ObjectStoreV2) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o ObjectStoreV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectStoreV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableObjectStoreV2 struct {
	value *ObjectStoreV2
	isSet bool
}

func (v NullableObjectStoreV2) Get() *ObjectStoreV2 {
	return v.value
}

func (v *NullableObjectStoreV2) Set(val *ObjectStoreV2) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectStoreV2) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectStoreV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectStoreV2(val *ObjectStoreV2) *NullableObjectStoreV2 {
	return &NullableObjectStoreV2{value: val, isSet: true}
}

func (v NullableObjectStoreV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectStoreV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


