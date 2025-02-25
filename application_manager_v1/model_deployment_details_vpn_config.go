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

// checks if the DeploymentDetailsVpnConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentDetailsVpnConfig{}

// DeploymentDetailsVpnConfig VPN configuration details if applicable.
type DeploymentDetailsVpnConfig struct {
	// The VPN network address.
	Network *string `json:"network,omitempty"`
	// The subnet mask for the VPN configuration.
	Mask *string `json:"mask,omitempty"`
}

// NewDeploymentDetailsVpnConfig instantiates a new DeploymentDetailsVpnConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentDetailsVpnConfig() *DeploymentDetailsVpnConfig {
	this := DeploymentDetailsVpnConfig{}
	return &this
}

// NewDeploymentDetailsVpnConfigWithDefaults instantiates a new DeploymentDetailsVpnConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentDetailsVpnConfigWithDefaults() *DeploymentDetailsVpnConfig {
	this := DeploymentDetailsVpnConfig{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *DeploymentDetailsVpnConfig) GetNetwork() string {
	if o == nil || IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDetailsVpnConfig) GetNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *DeploymentDetailsVpnConfig) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *DeploymentDetailsVpnConfig) SetNetwork(v string) {
	o.Network = &v
}

// GetMask returns the Mask field value if set, zero value otherwise.
func (o *DeploymentDetailsVpnConfig) GetMask() string {
	if o == nil || IsNil(o.Mask) {
		var ret string
		return ret
	}
	return *o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDetailsVpnConfig) GetMaskOk() (*string, bool) {
	if o == nil || IsNil(o.Mask) {
		return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *DeploymentDetailsVpnConfig) HasMask() bool {
	if o != nil && !IsNil(o.Mask) {
		return true
	}

	return false
}

// SetMask gets a reference to the given string and assigns it to the Mask field.
func (o *DeploymentDetailsVpnConfig) SetMask(v string) {
	o.Mask = &v
}

func (o DeploymentDetailsVpnConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentDetailsVpnConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.Mask) {
		toSerialize["mask"] = o.Mask
	}
	return toSerialize, nil
}

type NullableDeploymentDetailsVpnConfig struct {
	value *DeploymentDetailsVpnConfig
	isSet bool
}

func (v NullableDeploymentDetailsVpnConfig) Get() *DeploymentDetailsVpnConfig {
	return v.value
}

func (v *NullableDeploymentDetailsVpnConfig) Set(val *DeploymentDetailsVpnConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentDetailsVpnConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentDetailsVpnConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentDetailsVpnConfig(val *DeploymentDetailsVpnConfig) *NullableDeploymentDetailsVpnConfig {
	return &NullableDeploymentDetailsVpnConfig{value: val, isSet: true}
}

func (v NullableDeploymentDetailsVpnConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentDetailsVpnConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


