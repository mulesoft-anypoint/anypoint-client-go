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

// checks if the EndpointTlsContexts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointTlsContexts{}

// EndpointTlsContexts struct for EndpointTlsContexts
type EndpointTlsContexts struct {
	Audit *Audit `json:"audit,omitempty"`
	Inbound NullableEndpointTlsContextsInbound `json:"inbound,omitempty"`
	Outbound map[string]interface{} `json:"outbound,omitempty"`
}

// NewEndpointTlsContexts instantiates a new EndpointTlsContexts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointTlsContexts() *EndpointTlsContexts {
	this := EndpointTlsContexts{}
	return &this
}

// NewEndpointTlsContextsWithDefaults instantiates a new EndpointTlsContexts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointTlsContextsWithDefaults() *EndpointTlsContexts {
	this := EndpointTlsContexts{}
	return &this
}

// GetAudit returns the Audit field value if set, zero value otherwise.
func (o *EndpointTlsContexts) GetAudit() Audit {
	if o == nil || IsNil(o.Audit) {
		var ret Audit
		return ret
	}
	return *o.Audit
}

// GetAuditOk returns a tuple with the Audit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTlsContexts) GetAuditOk() (*Audit, bool) {
	if o == nil || IsNil(o.Audit) {
		return nil, false
	}
	return o.Audit, true
}

// HasAudit returns a boolean if a field has been set.
func (o *EndpointTlsContexts) HasAudit() bool {
	if o != nil && !IsNil(o.Audit) {
		return true
	}

	return false
}

// SetAudit gets a reference to the given Audit and assigns it to the Audit field.
func (o *EndpointTlsContexts) SetAudit(v Audit) {
	o.Audit = &v
}

// GetInbound returns the Inbound field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointTlsContexts) GetInbound() EndpointTlsContextsInbound {
	if o == nil || IsNil(o.Inbound.Get()) {
		var ret EndpointTlsContextsInbound
		return ret
	}
	return *o.Inbound.Get()
}

// GetInboundOk returns a tuple with the Inbound field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointTlsContexts) GetInboundOk() (*EndpointTlsContextsInbound, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inbound.Get(), o.Inbound.IsSet()
}

// HasInbound returns a boolean if a field has been set.
func (o *EndpointTlsContexts) HasInbound() bool {
	if o != nil && o.Inbound.IsSet() {
		return true
	}

	return false
}

// SetInbound gets a reference to the given NullableEndpointTlsContextsInbound and assigns it to the Inbound field.
func (o *EndpointTlsContexts) SetInbound(v EndpointTlsContextsInbound) {
	o.Inbound.Set(&v)
}
// SetInboundNil sets the value for Inbound to be an explicit nil
func (o *EndpointTlsContexts) SetInboundNil() {
	o.Inbound.Set(nil)
}

// UnsetInbound ensures that no value is present for Inbound, not even an explicit nil
func (o *EndpointTlsContexts) UnsetInbound() {
	o.Inbound.Unset()
}

// GetOutbound returns the Outbound field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointTlsContexts) GetOutbound() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Outbound
}

// GetOutboundOk returns a tuple with the Outbound field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointTlsContexts) GetOutboundOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Outbound) {
		return map[string]interface{}{}, false
	}
	return o.Outbound, true
}

// HasOutbound returns a boolean if a field has been set.
func (o *EndpointTlsContexts) HasOutbound() bool {
	if o != nil && IsNil(o.Outbound) {
		return true
	}

	return false
}

// SetOutbound gets a reference to the given map[string]interface{} and assigns it to the Outbound field.
func (o *EndpointTlsContexts) SetOutbound(v map[string]interface{}) {
	o.Outbound = v
}

func (o EndpointTlsContexts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointTlsContexts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Audit) {
		toSerialize["audit"] = o.Audit
	}
	if o.Inbound.IsSet() {
		toSerialize["inbound"] = o.Inbound.Get()
	}
	if o.Outbound != nil {
		toSerialize["outbound"] = o.Outbound
	}
	return toSerialize, nil
}

type NullableEndpointTlsContexts struct {
	value *EndpointTlsContexts
	isSet bool
}

func (v NullableEndpointTlsContexts) Get() *EndpointTlsContexts {
	return v.value
}

func (v *NullableEndpointTlsContexts) Set(val *EndpointTlsContexts) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointTlsContexts) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointTlsContexts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointTlsContexts(val *EndpointTlsContexts) *NullableEndpointTlsContexts {
	return &NullableEndpointTlsContexts{value: val, isSet: true}
}

func (v NullableEndpointTlsContexts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointTlsContexts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


