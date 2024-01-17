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

// checks if the Routing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Routing{}

// Routing struct for Routing
type Routing struct {
	Label *string `json:"label,omitempty"`
	Rules *RoutingRules `json:"rules,omitempty"`
	Upstreams []RoutingUpstreamsInner `json:"upstreams,omitempty"`
}

// NewRouting instantiates a new Routing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouting() *Routing {
	this := Routing{}
	return &this
}

// NewRoutingWithDefaults instantiates a new Routing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingWithDefaults() *Routing {
	this := Routing{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *Routing) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Routing) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *Routing) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *Routing) SetLabel(v string) {
	o.Label = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *Routing) GetRules() RoutingRules {
	if o == nil || IsNil(o.Rules) {
		var ret RoutingRules
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Routing) GetRulesOk() (*RoutingRules, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *Routing) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given RoutingRules and assigns it to the Rules field.
func (o *Routing) SetRules(v RoutingRules) {
	o.Rules = &v
}

// GetUpstreams returns the Upstreams field value if set, zero value otherwise.
func (o *Routing) GetUpstreams() []RoutingUpstreamsInner {
	if o == nil || IsNil(o.Upstreams) {
		var ret []RoutingUpstreamsInner
		return ret
	}
	return o.Upstreams
}

// GetUpstreamsOk returns a tuple with the Upstreams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Routing) GetUpstreamsOk() ([]RoutingUpstreamsInner, bool) {
	if o == nil || IsNil(o.Upstreams) {
		return nil, false
	}
	return o.Upstreams, true
}

// HasUpstreams returns a boolean if a field has been set.
func (o *Routing) HasUpstreams() bool {
	if o != nil && !IsNil(o.Upstreams) {
		return true
	}

	return false
}

// SetUpstreams gets a reference to the given []RoutingUpstreamsInner and assigns it to the Upstreams field.
func (o *Routing) SetUpstreams(v []RoutingUpstreamsInner) {
	o.Upstreams = v
}

func (o Routing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Routing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !IsNil(o.Upstreams) {
		toSerialize["upstreams"] = o.Upstreams
	}
	return toSerialize, nil
}

type NullableRouting struct {
	value *Routing
	isSet bool
}

func (v NullableRouting) Get() *Routing {
	return v.value
}

func (v *NullableRouting) Set(val *Routing) {
	v.value = val
	v.isSet = true
}

func (v NullableRouting) IsSet() bool {
	return v.isSet
}

func (v *NullableRouting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouting(val *Routing) *NullableRouting {
	return &NullableRouting{value: val, isSet: true}
}

func (v NullableRouting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


