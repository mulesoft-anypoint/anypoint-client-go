/*
Application Manager v1

The Application Manager API exists to provide Mule Application management operations from Anypoint Control Planes to any Runtime Plane, with capabilities that include:   - Deploying a Mule Application or API to a Mule Runtime   - Scaling up or down a running application   - Managing application settings (ie: properties)   - Deleting a Mule Application Deployment   - Changing artifact version or configurations of a deployment   - Starting, Stopping or Restarting applications   - etc. This API currently supports deployments to Cloudhub 1.0 targets only. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package application_manager_v1

import (
	"encoding/json"
)

// checks if the NewDeploymentStructPropertiesOptionsValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewDeploymentStructPropertiesOptionsValue{}

// NewDeploymentStructPropertiesOptionsValue struct for NewDeploymentStructPropertiesOptionsValue
type NewDeploymentStructPropertiesOptionsValue struct {
	Secure *bool `json:"secure,omitempty"`
}

// NewNewDeploymentStructPropertiesOptionsValue instantiates a new NewDeploymentStructPropertiesOptionsValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewDeploymentStructPropertiesOptionsValue() *NewDeploymentStructPropertiesOptionsValue {
	this := NewDeploymentStructPropertiesOptionsValue{}
	return &this
}

// NewNewDeploymentStructPropertiesOptionsValueWithDefaults instantiates a new NewDeploymentStructPropertiesOptionsValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewDeploymentStructPropertiesOptionsValueWithDefaults() *NewDeploymentStructPropertiesOptionsValue {
	this := NewDeploymentStructPropertiesOptionsValue{}
	return &this
}

// GetSecure returns the Secure field value if set, zero value otherwise.
func (o *NewDeploymentStructPropertiesOptionsValue) GetSecure() bool {
	if o == nil || IsNil(o.Secure) {
		var ret bool
		return ret
	}
	return *o.Secure
}

// GetSecureOk returns a tuple with the Secure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewDeploymentStructPropertiesOptionsValue) GetSecureOk() (*bool, bool) {
	if o == nil || IsNil(o.Secure) {
		return nil, false
	}
	return o.Secure, true
}

// HasSecure returns a boolean if a field has been set.
func (o *NewDeploymentStructPropertiesOptionsValue) HasSecure() bool {
	if o != nil && !IsNil(o.Secure) {
		return true
	}

	return false
}

// SetSecure gets a reference to the given bool and assigns it to the Secure field.
func (o *NewDeploymentStructPropertiesOptionsValue) SetSecure(v bool) {
	o.Secure = &v
}

func (o NewDeploymentStructPropertiesOptionsValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewDeploymentStructPropertiesOptionsValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Secure) {
		toSerialize["secure"] = o.Secure
	}
	return toSerialize, nil
}

type NullableNewDeploymentStructPropertiesOptionsValue struct {
	value *NewDeploymentStructPropertiesOptionsValue
	isSet bool
}

func (v NullableNewDeploymentStructPropertiesOptionsValue) Get() *NewDeploymentStructPropertiesOptionsValue {
	return v.value
}

func (v *NullableNewDeploymentStructPropertiesOptionsValue) Set(val *NewDeploymentStructPropertiesOptionsValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNewDeploymentStructPropertiesOptionsValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNewDeploymentStructPropertiesOptionsValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewDeploymentStructPropertiesOptionsValue(val *NewDeploymentStructPropertiesOptionsValue) *NullableNewDeploymentStructPropertiesOptionsValue {
	return &NullableNewDeploymentStructPropertiesOptionsValue{value: val, isSet: true}
}

func (v NullableNewDeploymentStructPropertiesOptionsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewDeploymentStructPropertiesOptionsValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


