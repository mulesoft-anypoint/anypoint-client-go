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

// checks if the DeploymentInfoRuntime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentInfoRuntime{}

// DeploymentInfoRuntime struct for DeploymentInfoRuntime
type DeploymentInfoRuntime struct {
	Version *string `json:"version,omitempty"`
	ReleaseChannel *string `json:"releaseChannel,omitempty"`
	Java *string `json:"java,omitempty"`
}

// NewDeploymentInfoRuntime instantiates a new DeploymentInfoRuntime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentInfoRuntime() *DeploymentInfoRuntime {
	this := DeploymentInfoRuntime{}
	return &this
}

// NewDeploymentInfoRuntimeWithDefaults instantiates a new DeploymentInfoRuntime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentInfoRuntimeWithDefaults() *DeploymentInfoRuntime {
	this := DeploymentInfoRuntime{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DeploymentInfoRuntime) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentInfoRuntime) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DeploymentInfoRuntime) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *DeploymentInfoRuntime) SetVersion(v string) {
	o.Version = &v
}

// GetReleaseChannel returns the ReleaseChannel field value if set, zero value otherwise.
func (o *DeploymentInfoRuntime) GetReleaseChannel() string {
	if o == nil || IsNil(o.ReleaseChannel) {
		var ret string
		return ret
	}
	return *o.ReleaseChannel
}

// GetReleaseChannelOk returns a tuple with the ReleaseChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentInfoRuntime) GetReleaseChannelOk() (*string, bool) {
	if o == nil || IsNil(o.ReleaseChannel) {
		return nil, false
	}
	return o.ReleaseChannel, true
}

// HasReleaseChannel returns a boolean if a field has been set.
func (o *DeploymentInfoRuntime) HasReleaseChannel() bool {
	if o != nil && !IsNil(o.ReleaseChannel) {
		return true
	}

	return false
}

// SetReleaseChannel gets a reference to the given string and assigns it to the ReleaseChannel field.
func (o *DeploymentInfoRuntime) SetReleaseChannel(v string) {
	o.ReleaseChannel = &v
}

// GetJava returns the Java field value if set, zero value otherwise.
func (o *DeploymentInfoRuntime) GetJava() string {
	if o == nil || IsNil(o.Java) {
		var ret string
		return ret
	}
	return *o.Java
}

// GetJavaOk returns a tuple with the Java field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentInfoRuntime) GetJavaOk() (*string, bool) {
	if o == nil || IsNil(o.Java) {
		return nil, false
	}
	return o.Java, true
}

// HasJava returns a boolean if a field has been set.
func (o *DeploymentInfoRuntime) HasJava() bool {
	if o != nil && !IsNil(o.Java) {
		return true
	}

	return false
}

// SetJava gets a reference to the given string and assigns it to the Java field.
func (o *DeploymentInfoRuntime) SetJava(v string) {
	o.Java = &v
}

func (o DeploymentInfoRuntime) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentInfoRuntime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.ReleaseChannel) {
		toSerialize["releaseChannel"] = o.ReleaseChannel
	}
	if !IsNil(o.Java) {
		toSerialize["java"] = o.Java
	}
	return toSerialize, nil
}

type NullableDeploymentInfoRuntime struct {
	value *DeploymentInfoRuntime
	isSet bool
}

func (v NullableDeploymentInfoRuntime) Get() *DeploymentInfoRuntime {
	return v.value
}

func (v *NullableDeploymentInfoRuntime) Set(val *DeploymentInfoRuntime) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentInfoRuntime) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentInfoRuntime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentInfoRuntime(val *DeploymentInfoRuntime) *NullableDeploymentInfoRuntime {
	return &NullableDeploymentInfoRuntime{value: val, isSet: true}
}

func (v NullableDeploymentInfoRuntime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentInfoRuntime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


