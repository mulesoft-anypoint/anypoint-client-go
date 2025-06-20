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

// checks if the Services type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Services{}

// Services struct for Services
type Services struct {
	ObjectStoreV2 *ObjectStoreV2 `json:"objectStoreV2,omitempty"`
}

// NewServices instantiates a new Services object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServices() *Services {
	this := Services{}
	return &this
}

// NewServicesWithDefaults instantiates a new Services object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesWithDefaults() *Services {
	this := Services{}
	return &this
}

// GetObjectStoreV2 returns the ObjectStoreV2 field value if set, zero value otherwise.
func (o *Services) GetObjectStoreV2() ObjectStoreV2 {
	if o == nil || IsNil(o.ObjectStoreV2) {
		var ret ObjectStoreV2
		return ret
	}
	return *o.ObjectStoreV2
}

// GetObjectStoreV2Ok returns a tuple with the ObjectStoreV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Services) GetObjectStoreV2Ok() (*ObjectStoreV2, bool) {
	if o == nil || IsNil(o.ObjectStoreV2) {
		return nil, false
	}
	return o.ObjectStoreV2, true
}

// HasObjectStoreV2 returns a boolean if a field has been set.
func (o *Services) HasObjectStoreV2() bool {
	if o != nil && !IsNil(o.ObjectStoreV2) {
		return true
	}

	return false
}

// SetObjectStoreV2 gets a reference to the given ObjectStoreV2 and assigns it to the ObjectStoreV2 field.
func (o *Services) SetObjectStoreV2(v ObjectStoreV2) {
	o.ObjectStoreV2 = &v
}

func (o Services) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Services) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObjectStoreV2) {
		toSerialize["objectStoreV2"] = o.ObjectStoreV2
	}
	return toSerialize, nil
}

type NullableServices struct {
	value *Services
	isSet bool
}

func (v NullableServices) Get() *Services {
	return v.value
}

func (v *NullableServices) Set(val *Services) {
	v.value = val
	v.isSet = true
}

func (v NullableServices) IsSet() bool {
	return v.isSet
}

func (v *NullableServices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServices(val *Services) *NullableServices {
	return &NullableServices{value: val, isSet: true}
}

func (v NullableServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


