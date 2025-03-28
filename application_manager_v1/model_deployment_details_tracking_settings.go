/*
Deployment v1

The Application Manager API exists to provide Mule Application management operations from Anypoint Control Planes to any Runtime Plane, with capabilities that include:   - Deploying a Mule Application or API to a Mule Runtime   - Scaling up or down a running application   - Managing application settings (ie: properties)   - Deleting a Mule Application Deployment   - Changing artifact version or configurations of a deployment   - Starting, Stopping or Restarting applications   - etc. This API currently supports deployments to Cloudhub 1.0 targets only. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package application_manager_v1

import (
	"encoding/json"
)

// checks if the DeploymentDetailsTrackingSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentDetailsTrackingSettings{}

// DeploymentDetailsTrackingSettings Settings for tracking application performance or usage.
type DeploymentDetailsTrackingSettings struct {
	// The level of tracking applied (e.g., DISABLED, BASIC, FULL).
	TrackingLevel *string `json:"trackingLevel,omitempty"`
}

// NewDeploymentDetailsTrackingSettings instantiates a new DeploymentDetailsTrackingSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentDetailsTrackingSettings() *DeploymentDetailsTrackingSettings {
	this := DeploymentDetailsTrackingSettings{}
	return &this
}

// NewDeploymentDetailsTrackingSettingsWithDefaults instantiates a new DeploymentDetailsTrackingSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentDetailsTrackingSettingsWithDefaults() *DeploymentDetailsTrackingSettings {
	this := DeploymentDetailsTrackingSettings{}
	return &this
}

// GetTrackingLevel returns the TrackingLevel field value if set, zero value otherwise.
func (o *DeploymentDetailsTrackingSettings) GetTrackingLevel() string {
	if o == nil || IsNil(o.TrackingLevel) {
		var ret string
		return ret
	}
	return *o.TrackingLevel
}

// GetTrackingLevelOk returns a tuple with the TrackingLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDetailsTrackingSettings) GetTrackingLevelOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingLevel) {
		return nil, false
	}
	return o.TrackingLevel, true
}

// HasTrackingLevel returns a boolean if a field has been set.
func (o *DeploymentDetailsTrackingSettings) HasTrackingLevel() bool {
	if o != nil && !IsNil(o.TrackingLevel) {
		return true
	}

	return false
}

// SetTrackingLevel gets a reference to the given string and assigns it to the TrackingLevel field.
func (o *DeploymentDetailsTrackingSettings) SetTrackingLevel(v string) {
	o.TrackingLevel = &v
}

func (o DeploymentDetailsTrackingSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentDetailsTrackingSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TrackingLevel) {
		toSerialize["trackingLevel"] = o.TrackingLevel
	}
	return toSerialize, nil
}

type NullableDeploymentDetailsTrackingSettings struct {
	value *DeploymentDetailsTrackingSettings
	isSet bool
}

func (v NullableDeploymentDetailsTrackingSettings) Get() *DeploymentDetailsTrackingSettings {
	return v.value
}

func (v *NullableDeploymentDetailsTrackingSettings) Set(val *DeploymentDetailsTrackingSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentDetailsTrackingSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentDetailsTrackingSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentDetailsTrackingSettings(val *DeploymentDetailsTrackingSettings) *NullableDeploymentDetailsTrackingSettings {
	return &NullableDeploymentDetailsTrackingSettings{value: val, isSet: true}
}

func (v NullableDeploymentDetailsTrackingSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentDetailsTrackingSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


