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

// checks if the NewDeploymentStructMuleVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewDeploymentStructMuleVersion{}

// NewDeploymentStructMuleVersion struct for NewDeploymentStructMuleVersion
type NewDeploymentStructMuleVersion struct {
	Version *string `json:"version,omitempty"`
}

// NewNewDeploymentStructMuleVersion instantiates a new NewDeploymentStructMuleVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewDeploymentStructMuleVersion() *NewDeploymentStructMuleVersion {
	this := NewDeploymentStructMuleVersion{}
	return &this
}

// NewNewDeploymentStructMuleVersionWithDefaults instantiates a new NewDeploymentStructMuleVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewDeploymentStructMuleVersionWithDefaults() *NewDeploymentStructMuleVersion {
	this := NewDeploymentStructMuleVersion{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *NewDeploymentStructMuleVersion) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewDeploymentStructMuleVersion) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *NewDeploymentStructMuleVersion) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *NewDeploymentStructMuleVersion) SetVersion(v string) {
	o.Version = &v
}

func (o NewDeploymentStructMuleVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewDeploymentStructMuleVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableNewDeploymentStructMuleVersion struct {
	value *NewDeploymentStructMuleVersion
	isSet bool
}

func (v NullableNewDeploymentStructMuleVersion) Get() *NewDeploymentStructMuleVersion {
	return v.value
}

func (v *NullableNewDeploymentStructMuleVersion) Set(val *NewDeploymentStructMuleVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableNewDeploymentStructMuleVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableNewDeploymentStructMuleVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewDeploymentStructMuleVersion(val *NewDeploymentStructMuleVersion) *NullableNewDeploymentStructMuleVersion {
	return &NullableNewDeploymentStructMuleVersion{value: val, isSet: true}
}

func (v NullableNewDeploymentStructMuleVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewDeploymentStructMuleVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


