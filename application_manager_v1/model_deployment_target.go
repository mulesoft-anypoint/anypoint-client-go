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

// checks if the DeploymentTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentTarget{}

// DeploymentTarget Information about the target environment where the deployment is hosted.
type DeploymentTarget struct {
	// The target type (e.g., CLOUDHUB, MC).
	Type *string `json:"type,omitempty"`
	// The target subtype (if applicable).
	Subtype *string `json:"subtype,omitempty"`
	// The provider of the target platform.
	Provider *string `json:"provider,omitempty"`
	// The identifier of the target.
	Id *string `json:"id,omitempty"`
	// The name of the target.
	Name *string `json:"name,omitempty"`
	// The current status of the target.
	Status *string `json:"status,omitempty"`
}

// NewDeploymentTarget instantiates a new DeploymentTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentTarget() *DeploymentTarget {
	this := DeploymentTarget{}
	return &this
}

// NewDeploymentTargetWithDefaults instantiates a new DeploymentTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentTargetWithDefaults() *DeploymentTarget {
	this := DeploymentTarget{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DeploymentTarget) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentTarget) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeploymentTarget) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DeploymentTarget) SetType(v string) {
	o.Type = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *DeploymentTarget) GetSubtype() string {
	if o == nil || IsNil(o.Subtype) {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentTarget) GetSubtypeOk() (*string, bool) {
	if o == nil || IsNil(o.Subtype) {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *DeploymentTarget) HasSubtype() bool {
	if o != nil && !IsNil(o.Subtype) {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *DeploymentTarget) SetSubtype(v string) {
	o.Subtype = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *DeploymentTarget) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentTarget) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *DeploymentTarget) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *DeploymentTarget) SetProvider(v string) {
	o.Provider = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeploymentTarget) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentTarget) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeploymentTarget) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeploymentTarget) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentTarget) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentTarget) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentTarget) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentTarget) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeploymentTarget) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentTarget) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeploymentTarget) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DeploymentTarget) SetStatus(v string) {
	o.Status = &v
}

func (o DeploymentTarget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Subtype) {
		toSerialize["subtype"] = o.Subtype
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableDeploymentTarget struct {
	value *DeploymentTarget
	isSet bool
}

func (v NullableDeploymentTarget) Get() *DeploymentTarget {
	return v.value
}

func (v *NullableDeploymentTarget) Set(val *DeploymentTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentTarget(val *DeploymentTarget) *NullableDeploymentTarget {
	return &NullableDeploymentTarget{value: val, isSet: true}
}

func (v NullableDeploymentTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


