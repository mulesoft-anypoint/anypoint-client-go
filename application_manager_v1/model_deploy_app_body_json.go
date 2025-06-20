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

// checks if the DeployAppBodyJSON type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeployAppBodyJSON{}

// DeployAppBodyJSON Application deployment body to deploy from exchange.
type DeployAppBodyJSON struct {
	ApplicationInfo *NewDeploymentStruct `json:"applicationInfo,omitempty"`
	ApplicationSource *ExchangeApplicationSource `json:"applicationSource,omitempty"`
	// Indicates whether the application should be automatically started after deployment.
	AutoStart *bool `json:"autoStart,omitempty"`
}

// NewDeployAppBodyJSON instantiates a new DeployAppBodyJSON object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeployAppBodyJSON() *DeployAppBodyJSON {
	this := DeployAppBodyJSON{}
	return &this
}

// NewDeployAppBodyJSONWithDefaults instantiates a new DeployAppBodyJSON object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeployAppBodyJSONWithDefaults() *DeployAppBodyJSON {
	this := DeployAppBodyJSON{}
	return &this
}

// GetApplicationInfo returns the ApplicationInfo field value if set, zero value otherwise.
func (o *DeployAppBodyJSON) GetApplicationInfo() NewDeploymentStruct {
	if o == nil || IsNil(o.ApplicationInfo) {
		var ret NewDeploymentStruct
		return ret
	}
	return *o.ApplicationInfo
}

// GetApplicationInfoOk returns a tuple with the ApplicationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployAppBodyJSON) GetApplicationInfoOk() (*NewDeploymentStruct, bool) {
	if o == nil || IsNil(o.ApplicationInfo) {
		return nil, false
	}
	return o.ApplicationInfo, true
}

// HasApplicationInfo returns a boolean if a field has been set.
func (o *DeployAppBodyJSON) HasApplicationInfo() bool {
	if o != nil && !IsNil(o.ApplicationInfo) {
		return true
	}

	return false
}

// SetApplicationInfo gets a reference to the given NewDeploymentStruct and assigns it to the ApplicationInfo field.
func (o *DeployAppBodyJSON) SetApplicationInfo(v NewDeploymentStruct) {
	o.ApplicationInfo = &v
}

// GetApplicationSource returns the ApplicationSource field value if set, zero value otherwise.
func (o *DeployAppBodyJSON) GetApplicationSource() ExchangeApplicationSource {
	if o == nil || IsNil(o.ApplicationSource) {
		var ret ExchangeApplicationSource
		return ret
	}
	return *o.ApplicationSource
}

// GetApplicationSourceOk returns a tuple with the ApplicationSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployAppBodyJSON) GetApplicationSourceOk() (*ExchangeApplicationSource, bool) {
	if o == nil || IsNil(o.ApplicationSource) {
		return nil, false
	}
	return o.ApplicationSource, true
}

// HasApplicationSource returns a boolean if a field has been set.
func (o *DeployAppBodyJSON) HasApplicationSource() bool {
	if o != nil && !IsNil(o.ApplicationSource) {
		return true
	}

	return false
}

// SetApplicationSource gets a reference to the given ExchangeApplicationSource and assigns it to the ApplicationSource field.
func (o *DeployAppBodyJSON) SetApplicationSource(v ExchangeApplicationSource) {
	o.ApplicationSource = &v
}

// GetAutoStart returns the AutoStart field value if set, zero value otherwise.
func (o *DeployAppBodyJSON) GetAutoStart() bool {
	if o == nil || IsNil(o.AutoStart) {
		var ret bool
		return ret
	}
	return *o.AutoStart
}

// GetAutoStartOk returns a tuple with the AutoStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployAppBodyJSON) GetAutoStartOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoStart) {
		return nil, false
	}
	return o.AutoStart, true
}

// HasAutoStart returns a boolean if a field has been set.
func (o *DeployAppBodyJSON) HasAutoStart() bool {
	if o != nil && !IsNil(o.AutoStart) {
		return true
	}

	return false
}

// SetAutoStart gets a reference to the given bool and assigns it to the AutoStart field.
func (o *DeployAppBodyJSON) SetAutoStart(v bool) {
	o.AutoStart = &v
}

func (o DeployAppBodyJSON) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeployAppBodyJSON) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationInfo) {
		toSerialize["applicationInfo"] = o.ApplicationInfo
	}
	if !IsNil(o.ApplicationSource) {
		toSerialize["applicationSource"] = o.ApplicationSource
	}
	if !IsNil(o.AutoStart) {
		toSerialize["autoStart"] = o.AutoStart
	}
	return toSerialize, nil
}

type NullableDeployAppBodyJSON struct {
	value *DeployAppBodyJSON
	isSet bool
}

func (v NullableDeployAppBodyJSON) Get() *DeployAppBodyJSON {
	return v.value
}

func (v *NullableDeployAppBodyJSON) Set(val *DeployAppBodyJSON) {
	v.value = val
	v.isSet = true
}

func (v NullableDeployAppBodyJSON) IsSet() bool {
	return v.isSet
}

func (v *NullableDeployAppBodyJSON) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeployAppBodyJSON(val *DeployAppBodyJSON) *NullableDeployAppBodyJSON {
	return &NullableDeployAppBodyJSON{value: val, isSet: true}
}

func (v NullableDeployAppBodyJSON) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeployAppBodyJSON) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


