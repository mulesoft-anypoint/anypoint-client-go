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

// checks if the DeploymentDetailsDeploymentGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentDetailsDeploymentGroup{}

// DeploymentDetailsDeploymentGroup Information about the deployment group to which this app belongs.
type DeploymentDetailsDeploymentGroup struct {
	// Unique identifier for the deployment group.
	Id *string `json:"id,omitempty"`
	// The name of the deployment group.
	Name *string `json:"name,omitempty"`
}

// NewDeploymentDetailsDeploymentGroup instantiates a new DeploymentDetailsDeploymentGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentDetailsDeploymentGroup() *DeploymentDetailsDeploymentGroup {
	this := DeploymentDetailsDeploymentGroup{}
	return &this
}

// NewDeploymentDetailsDeploymentGroupWithDefaults instantiates a new DeploymentDetailsDeploymentGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentDetailsDeploymentGroupWithDefaults() *DeploymentDetailsDeploymentGroup {
	this := DeploymentDetailsDeploymentGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeploymentDetailsDeploymentGroup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDetailsDeploymentGroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeploymentDetailsDeploymentGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeploymentDetailsDeploymentGroup) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentDetailsDeploymentGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDetailsDeploymentGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentDetailsDeploymentGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentDetailsDeploymentGroup) SetName(v string) {
	o.Name = &v
}

func (o DeploymentDetailsDeploymentGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentDetailsDeploymentGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableDeploymentDetailsDeploymentGroup struct {
	value *DeploymentDetailsDeploymentGroup
	isSet bool
}

func (v NullableDeploymentDetailsDeploymentGroup) Get() *DeploymentDetailsDeploymentGroup {
	return v.value
}

func (v *NullableDeploymentDetailsDeploymentGroup) Set(val *DeploymentDetailsDeploymentGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentDetailsDeploymentGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentDetailsDeploymentGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentDetailsDeploymentGroup(val *DeploymentDetailsDeploymentGroup) *NullableDeploymentDetailsDeploymentGroup {
	return &NullableDeploymentDetailsDeploymentGroup{value: val, isSet: true}
}

func (v NullableDeploymentDetailsDeploymentGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentDetailsDeploymentGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


