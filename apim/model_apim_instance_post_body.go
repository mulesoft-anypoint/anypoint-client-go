/*
API Manager API

API Manager API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim

import (
	"encoding/json"
)

// checks if the ApimInstancePostBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApimInstancePostBody{}

// ApimInstancePostBody struct for ApimInstancePostBody
type ApimInstancePostBody struct {
	Technology *string `json:"technology,omitempty"`
	Endpoint NullableEndpointPostBody `json:"endpoint,omitempty"`
	Spec *Spec `json:"spec,omitempty"`
	Routing []RoutingPostBodyInner `json:"routing,omitempty"`
	Deployment NullableDeploymentPostBody `json:"deployment,omitempty"`
	InstanceLabel NullableString `json:"instanceLabel,omitempty"`
}

// NewApimInstancePostBody instantiates a new ApimInstancePostBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApimInstancePostBody() *ApimInstancePostBody {
	this := ApimInstancePostBody{}
	return &this
}

// NewApimInstancePostBodyWithDefaults instantiates a new ApimInstancePostBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApimInstancePostBodyWithDefaults() *ApimInstancePostBody {
	this := ApimInstancePostBody{}
	return &this
}

// GetTechnology returns the Technology field value if set, zero value otherwise.
func (o *ApimInstancePostBody) GetTechnology() string {
	if o == nil || IsNil(o.Technology) {
		var ret string
		return ret
	}
	return *o.Technology
}

// GetTechnologyOk returns a tuple with the Technology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstancePostBody) GetTechnologyOk() (*string, bool) {
	if o == nil || IsNil(o.Technology) {
		return nil, false
	}
	return o.Technology, true
}

// HasTechnology returns a boolean if a field has been set.
func (o *ApimInstancePostBody) HasTechnology() bool {
	if o != nil && !IsNil(o.Technology) {
		return true
	}

	return false
}

// SetTechnology gets a reference to the given string and assigns it to the Technology field.
func (o *ApimInstancePostBody) SetTechnology(v string) {
	o.Technology = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApimInstancePostBody) GetEndpoint() EndpointPostBody {
	if o == nil || IsNil(o.Endpoint.Get()) {
		var ret EndpointPostBody
		return ret
	}
	return *o.Endpoint.Get()
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApimInstancePostBody) GetEndpointOk() (*EndpointPostBody, bool) {
	if o == nil {
		return nil, false
	}
	return o.Endpoint.Get(), o.Endpoint.IsSet()
}

// HasEndpoint returns a boolean if a field has been set.
func (o *ApimInstancePostBody) HasEndpoint() bool {
	if o != nil && o.Endpoint.IsSet() {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given NullableEndpointPostBody and assigns it to the Endpoint field.
func (o *ApimInstancePostBody) SetEndpoint(v EndpointPostBody) {
	o.Endpoint.Set(&v)
}
// SetEndpointNil sets the value for Endpoint to be an explicit nil
func (o *ApimInstancePostBody) SetEndpointNil() {
	o.Endpoint.Set(nil)
}

// UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
func (o *ApimInstancePostBody) UnsetEndpoint() {
	o.Endpoint.Unset()
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *ApimInstancePostBody) GetSpec() Spec {
	if o == nil || IsNil(o.Spec) {
		var ret Spec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstancePostBody) GetSpecOk() (*Spec, bool) {
	if o == nil || IsNil(o.Spec) {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *ApimInstancePostBody) HasSpec() bool {
	if o != nil && !IsNil(o.Spec) {
		return true
	}

	return false
}

// SetSpec gets a reference to the given Spec and assigns it to the Spec field.
func (o *ApimInstancePostBody) SetSpec(v Spec) {
	o.Spec = &v
}

// GetRouting returns the Routing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApimInstancePostBody) GetRouting() []RoutingPostBodyInner {
	if o == nil {
		var ret []RoutingPostBodyInner
		return ret
	}
	return o.Routing
}

// GetRoutingOk returns a tuple with the Routing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApimInstancePostBody) GetRoutingOk() ([]RoutingPostBodyInner, bool) {
	if o == nil || IsNil(o.Routing) {
		return nil, false
	}
	return o.Routing, true
}

// HasRouting returns a boolean if a field has been set.
func (o *ApimInstancePostBody) HasRouting() bool {
	if o != nil && IsNil(o.Routing) {
		return true
	}

	return false
}

// SetRouting gets a reference to the given []RoutingPostBodyInner and assigns it to the Routing field.
func (o *ApimInstancePostBody) SetRouting(v []RoutingPostBodyInner) {
	o.Routing = v
}

// GetDeployment returns the Deployment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApimInstancePostBody) GetDeployment() DeploymentPostBody {
	if o == nil || IsNil(o.Deployment.Get()) {
		var ret DeploymentPostBody
		return ret
	}
	return *o.Deployment.Get()
}

// GetDeploymentOk returns a tuple with the Deployment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApimInstancePostBody) GetDeploymentOk() (*DeploymentPostBody, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deployment.Get(), o.Deployment.IsSet()
}

// HasDeployment returns a boolean if a field has been set.
func (o *ApimInstancePostBody) HasDeployment() bool {
	if o != nil && o.Deployment.IsSet() {
		return true
	}

	return false
}

// SetDeployment gets a reference to the given NullableDeploymentPostBody and assigns it to the Deployment field.
func (o *ApimInstancePostBody) SetDeployment(v DeploymentPostBody) {
	o.Deployment.Set(&v)
}
// SetDeploymentNil sets the value for Deployment to be an explicit nil
func (o *ApimInstancePostBody) SetDeploymentNil() {
	o.Deployment.Set(nil)
}

// UnsetDeployment ensures that no value is present for Deployment, not even an explicit nil
func (o *ApimInstancePostBody) UnsetDeployment() {
	o.Deployment.Unset()
}

// GetInstanceLabel returns the InstanceLabel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApimInstancePostBody) GetInstanceLabel() string {
	if o == nil || IsNil(o.InstanceLabel.Get()) {
		var ret string
		return ret
	}
	return *o.InstanceLabel.Get()
}

// GetInstanceLabelOk returns a tuple with the InstanceLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApimInstancePostBody) GetInstanceLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstanceLabel.Get(), o.InstanceLabel.IsSet()
}

// HasInstanceLabel returns a boolean if a field has been set.
func (o *ApimInstancePostBody) HasInstanceLabel() bool {
	if o != nil && o.InstanceLabel.IsSet() {
		return true
	}

	return false
}

// SetInstanceLabel gets a reference to the given NullableString and assigns it to the InstanceLabel field.
func (o *ApimInstancePostBody) SetInstanceLabel(v string) {
	o.InstanceLabel.Set(&v)
}
// SetInstanceLabelNil sets the value for InstanceLabel to be an explicit nil
func (o *ApimInstancePostBody) SetInstanceLabelNil() {
	o.InstanceLabel.Set(nil)
}

// UnsetInstanceLabel ensures that no value is present for InstanceLabel, not even an explicit nil
func (o *ApimInstancePostBody) UnsetInstanceLabel() {
	o.InstanceLabel.Unset()
}

func (o ApimInstancePostBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApimInstancePostBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Technology) {
		toSerialize["technology"] = o.Technology
	}
	if o.Endpoint.IsSet() {
		toSerialize["endpoint"] = o.Endpoint.Get()
	}
	if !IsNil(o.Spec) {
		toSerialize["spec"] = o.Spec
	}
	if o.Routing != nil {
		toSerialize["routing"] = o.Routing
	}
	if o.Deployment.IsSet() {
		toSerialize["deployment"] = o.Deployment.Get()
	}
	if o.InstanceLabel.IsSet() {
		toSerialize["instanceLabel"] = o.InstanceLabel.Get()
	}
	return toSerialize, nil
}

type NullableApimInstancePostBody struct {
	value *ApimInstancePostBody
	isSet bool
}

func (v NullableApimInstancePostBody) Get() *ApimInstancePostBody {
	return v.value
}

func (v *NullableApimInstancePostBody) Set(val *ApimInstancePostBody) {
	v.value = val
	v.isSet = true
}

func (v NullableApimInstancePostBody) IsSet() bool {
	return v.isSet
}

func (v *NullableApimInstancePostBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApimInstancePostBody(val *ApimInstancePostBody) *NullableApimInstancePostBody {
	return &NullableApimInstancePostBody{value: val, isSet: true}
}

func (v NullableApimInstancePostBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApimInstancePostBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


