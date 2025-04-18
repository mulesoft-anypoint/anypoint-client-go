/*
Organization API

Description of the Organization API

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package org

import (
	"encoding/json"
)

// checks if the Usage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Usage{}

// Usage struct for Usage
type Usage struct {
	// Production VCores consumed
	ProductionVCoresConsumed *float32 `json:"productionVCoresConsumed,omitempty"`
	// Sandbox VCores consumed
	SandboxVCoresConsumed *float32 `json:"sandboxVCoresConsumed,omitempty"`
	// Design VCores consumed
	DesignVCoresConsumed *float32 `json:"designVCoresConsumed,omitempty"`
	// Static IPs consumed
	StaticIpsConsumed *float32 `json:"staticIpsConsumed,omitempty"`
	// VPCs consumed
	VpcsConsumed *int32 `json:"vpcsConsumed,omitempty"`
	// VPNs consumed
	VpnsConsumed *int32 `json:"vpnsConsumed,omitempty"`
	// Network connections consumed
	NetworkConnectionsConsumed *int32 `json:"networkConnectionsConsumed,omitempty"`
	// Load balancers consumed
	LoadBalancersConsumed *int32 `json:"loadBalancersConsumed,omitempty"`
	// Load balancer workers consumed
	LoadbalancerWorkersConsumed *int32 `json:"loadbalancerWorkersConsumed,omitempty"`
}

// NewUsage instantiates a new Usage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsage() *Usage {
	this := Usage{}
	var productionVCoresConsumed float32 = 0
	this.ProductionVCoresConsumed = &productionVCoresConsumed
	var sandboxVCoresConsumed float32 = 0
	this.SandboxVCoresConsumed = &sandboxVCoresConsumed
	var designVCoresConsumed float32 = 0
	this.DesignVCoresConsumed = &designVCoresConsumed
	var staticIpsConsumed float32 = 0
	this.StaticIpsConsumed = &staticIpsConsumed
	var vpcsConsumed int32 = 0
	this.VpcsConsumed = &vpcsConsumed
	var vpnsConsumed int32 = 0
	this.VpnsConsumed = &vpnsConsumed
	var networkConnectionsConsumed int32 = 0
	this.NetworkConnectionsConsumed = &networkConnectionsConsumed
	var loadBalancersConsumed int32 = 0
	this.LoadBalancersConsumed = &loadBalancersConsumed
	var loadbalancerWorkersConsumed int32 = 0
	this.LoadbalancerWorkersConsumed = &loadbalancerWorkersConsumed
	return &this
}

// NewUsageWithDefaults instantiates a new Usage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageWithDefaults() *Usage {
	this := Usage{}
	var productionVCoresConsumed float32 = 0
	this.ProductionVCoresConsumed = &productionVCoresConsumed
	var sandboxVCoresConsumed float32 = 0
	this.SandboxVCoresConsumed = &sandboxVCoresConsumed
	var designVCoresConsumed float32 = 0
	this.DesignVCoresConsumed = &designVCoresConsumed
	var staticIpsConsumed float32 = 0
	this.StaticIpsConsumed = &staticIpsConsumed
	var vpcsConsumed int32 = 0
	this.VpcsConsumed = &vpcsConsumed
	var vpnsConsumed int32 = 0
	this.VpnsConsumed = &vpnsConsumed
	var networkConnectionsConsumed int32 = 0
	this.NetworkConnectionsConsumed = &networkConnectionsConsumed
	var loadBalancersConsumed int32 = 0
	this.LoadBalancersConsumed = &loadBalancersConsumed
	var loadbalancerWorkersConsumed int32 = 0
	this.LoadbalancerWorkersConsumed = &loadbalancerWorkersConsumed
	return &this
}

// GetProductionVCoresConsumed returns the ProductionVCoresConsumed field value if set, zero value otherwise.
func (o *Usage) GetProductionVCoresConsumed() float32 {
	if o == nil || IsNil(o.ProductionVCoresConsumed) {
		var ret float32
		return ret
	}
	return *o.ProductionVCoresConsumed
}

// GetProductionVCoresConsumedOk returns a tuple with the ProductionVCoresConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Usage) GetProductionVCoresConsumedOk() (*float32, bool) {
	if o == nil || IsNil(o.ProductionVCoresConsumed) {
		return nil, false
	}
	return o.ProductionVCoresConsumed, true
}

// HasProductionVCoresConsumed returns a boolean if a field has been set.
func (o *Usage) HasProductionVCoresConsumed() bool {
	if o != nil && !IsNil(o.ProductionVCoresConsumed) {
		return true
	}

	return false
}

// SetProductionVCoresConsumed gets a reference to the given float32 and assigns it to the ProductionVCoresConsumed field.
func (o *Usage) SetProductionVCoresConsumed(v float32) {
	o.ProductionVCoresConsumed = &v
}

// GetSandboxVCoresConsumed returns the SandboxVCoresConsumed field value if set, zero value otherwise.
func (o *Usage) GetSandboxVCoresConsumed() float32 {
	if o == nil || IsNil(o.SandboxVCoresConsumed) {
		var ret float32
		return ret
	}
	return *o.SandboxVCoresConsumed
}

// GetSandboxVCoresConsumedOk returns a tuple with the SandboxVCoresConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Usage) GetSandboxVCoresConsumedOk() (*float32, bool) {
	if o == nil || IsNil(o.SandboxVCoresConsumed) {
		return nil, false
	}
	return o.SandboxVCoresConsumed, true
}

// HasSandboxVCoresConsumed returns a boolean if a field has been set.
func (o *Usage) HasSandboxVCoresConsumed() bool {
	if o != nil && !IsNil(o.SandboxVCoresConsumed) {
		return true
	}

	return false
}

// SetSandboxVCoresConsumed gets a reference to the given float32 and assigns it to the SandboxVCoresConsumed field.
func (o *Usage) SetSandboxVCoresConsumed(v float32) {
	o.SandboxVCoresConsumed = &v
}

// GetDesignVCoresConsumed returns the DesignVCoresConsumed field value if set, zero value otherwise.
func (o *Usage) GetDesignVCoresConsumed() float32 {
	if o == nil || IsNil(o.DesignVCoresConsumed) {
		var ret float32
		return ret
	}
	return *o.DesignVCoresConsumed
}

// GetDesignVCoresConsumedOk returns a tuple with the DesignVCoresConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Usage) GetDesignVCoresConsumedOk() (*float32, bool) {
	if o == nil || IsNil(o.DesignVCoresConsumed) {
		return nil, false
	}
	return o.DesignVCoresConsumed, true
}

// HasDesignVCoresConsumed returns a boolean if a field has been set.
func (o *Usage) HasDesignVCoresConsumed() bool {
	if o != nil && !IsNil(o.DesignVCoresConsumed) {
		return true
	}

	return false
}

// SetDesignVCoresConsumed gets a reference to the given float32 and assigns it to the DesignVCoresConsumed field.
func (o *Usage) SetDesignVCoresConsumed(v float32) {
	o.DesignVCoresConsumed = &v
}

// GetStaticIpsConsumed returns the StaticIpsConsumed field value if set, zero value otherwise.
func (o *Usage) GetStaticIpsConsumed() float32 {
	if o == nil || IsNil(o.StaticIpsConsumed) {
		var ret float32
		return ret
	}
	return *o.StaticIpsConsumed
}

// GetStaticIpsConsumedOk returns a tuple with the StaticIpsConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Usage) GetStaticIpsConsumedOk() (*float32, bool) {
	if o == nil || IsNil(o.StaticIpsConsumed) {
		return nil, false
	}
	return o.StaticIpsConsumed, true
}

// HasStaticIpsConsumed returns a boolean if a field has been set.
func (o *Usage) HasStaticIpsConsumed() bool {
	if o != nil && !IsNil(o.StaticIpsConsumed) {
		return true
	}

	return false
}

// SetStaticIpsConsumed gets a reference to the given float32 and assigns it to the StaticIpsConsumed field.
func (o *Usage) SetStaticIpsConsumed(v float32) {
	o.StaticIpsConsumed = &v
}

// GetVpcsConsumed returns the VpcsConsumed field value if set, zero value otherwise.
func (o *Usage) GetVpcsConsumed() int32 {
	if o == nil || IsNil(o.VpcsConsumed) {
		var ret int32
		return ret
	}
	return *o.VpcsConsumed
}

// GetVpcsConsumedOk returns a tuple with the VpcsConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Usage) GetVpcsConsumedOk() (*int32, bool) {
	if o == nil || IsNil(o.VpcsConsumed) {
		return nil, false
	}
	return o.VpcsConsumed, true
}

// HasVpcsConsumed returns a boolean if a field has been set.
func (o *Usage) HasVpcsConsumed() bool {
	if o != nil && !IsNil(o.VpcsConsumed) {
		return true
	}

	return false
}

// SetVpcsConsumed gets a reference to the given int32 and assigns it to the VpcsConsumed field.
func (o *Usage) SetVpcsConsumed(v int32) {
	o.VpcsConsumed = &v
}

// GetVpnsConsumed returns the VpnsConsumed field value if set, zero value otherwise.
func (o *Usage) GetVpnsConsumed() int32 {
	if o == nil || IsNil(o.VpnsConsumed) {
		var ret int32
		return ret
	}
	return *o.VpnsConsumed
}

// GetVpnsConsumedOk returns a tuple with the VpnsConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Usage) GetVpnsConsumedOk() (*int32, bool) {
	if o == nil || IsNil(o.VpnsConsumed) {
		return nil, false
	}
	return o.VpnsConsumed, true
}

// HasVpnsConsumed returns a boolean if a field has been set.
func (o *Usage) HasVpnsConsumed() bool {
	if o != nil && !IsNil(o.VpnsConsumed) {
		return true
	}

	return false
}

// SetVpnsConsumed gets a reference to the given int32 and assigns it to the VpnsConsumed field.
func (o *Usage) SetVpnsConsumed(v int32) {
	o.VpnsConsumed = &v
}

// GetNetworkConnectionsConsumed returns the NetworkConnectionsConsumed field value if set, zero value otherwise.
func (o *Usage) GetNetworkConnectionsConsumed() int32 {
	if o == nil || IsNil(o.NetworkConnectionsConsumed) {
		var ret int32
		return ret
	}
	return *o.NetworkConnectionsConsumed
}

// GetNetworkConnectionsConsumedOk returns a tuple with the NetworkConnectionsConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Usage) GetNetworkConnectionsConsumedOk() (*int32, bool) {
	if o == nil || IsNil(o.NetworkConnectionsConsumed) {
		return nil, false
	}
	return o.NetworkConnectionsConsumed, true
}

// HasNetworkConnectionsConsumed returns a boolean if a field has been set.
func (o *Usage) HasNetworkConnectionsConsumed() bool {
	if o != nil && !IsNil(o.NetworkConnectionsConsumed) {
		return true
	}

	return false
}

// SetNetworkConnectionsConsumed gets a reference to the given int32 and assigns it to the NetworkConnectionsConsumed field.
func (o *Usage) SetNetworkConnectionsConsumed(v int32) {
	o.NetworkConnectionsConsumed = &v
}

// GetLoadBalancersConsumed returns the LoadBalancersConsumed field value if set, zero value otherwise.
func (o *Usage) GetLoadBalancersConsumed() int32 {
	if o == nil || IsNil(o.LoadBalancersConsumed) {
		var ret int32
		return ret
	}
	return *o.LoadBalancersConsumed
}

// GetLoadBalancersConsumedOk returns a tuple with the LoadBalancersConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Usage) GetLoadBalancersConsumedOk() (*int32, bool) {
	if o == nil || IsNil(o.LoadBalancersConsumed) {
		return nil, false
	}
	return o.LoadBalancersConsumed, true
}

// HasLoadBalancersConsumed returns a boolean if a field has been set.
func (o *Usage) HasLoadBalancersConsumed() bool {
	if o != nil && !IsNil(o.LoadBalancersConsumed) {
		return true
	}

	return false
}

// SetLoadBalancersConsumed gets a reference to the given int32 and assigns it to the LoadBalancersConsumed field.
func (o *Usage) SetLoadBalancersConsumed(v int32) {
	o.LoadBalancersConsumed = &v
}

// GetLoadbalancerWorkersConsumed returns the LoadbalancerWorkersConsumed field value if set, zero value otherwise.
func (o *Usage) GetLoadbalancerWorkersConsumed() int32 {
	if o == nil || IsNil(o.LoadbalancerWorkersConsumed) {
		var ret int32
		return ret
	}
	return *o.LoadbalancerWorkersConsumed
}

// GetLoadbalancerWorkersConsumedOk returns a tuple with the LoadbalancerWorkersConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Usage) GetLoadbalancerWorkersConsumedOk() (*int32, bool) {
	if o == nil || IsNil(o.LoadbalancerWorkersConsumed) {
		return nil, false
	}
	return o.LoadbalancerWorkersConsumed, true
}

// HasLoadbalancerWorkersConsumed returns a boolean if a field has been set.
func (o *Usage) HasLoadbalancerWorkersConsumed() bool {
	if o != nil && !IsNil(o.LoadbalancerWorkersConsumed) {
		return true
	}

	return false
}

// SetLoadbalancerWorkersConsumed gets a reference to the given int32 and assigns it to the LoadbalancerWorkersConsumed field.
func (o *Usage) SetLoadbalancerWorkersConsumed(v int32) {
	o.LoadbalancerWorkersConsumed = &v
}

func (o Usage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Usage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductionVCoresConsumed) {
		toSerialize["productionVCoresConsumed"] = o.ProductionVCoresConsumed
	}
	if !IsNil(o.SandboxVCoresConsumed) {
		toSerialize["sandboxVCoresConsumed"] = o.SandboxVCoresConsumed
	}
	if !IsNil(o.DesignVCoresConsumed) {
		toSerialize["designVCoresConsumed"] = o.DesignVCoresConsumed
	}
	if !IsNil(o.StaticIpsConsumed) {
		toSerialize["staticIpsConsumed"] = o.StaticIpsConsumed
	}
	if !IsNil(o.VpcsConsumed) {
		toSerialize["vpcsConsumed"] = o.VpcsConsumed
	}
	if !IsNil(o.VpnsConsumed) {
		toSerialize["vpnsConsumed"] = o.VpnsConsumed
	}
	if !IsNil(o.NetworkConnectionsConsumed) {
		toSerialize["networkConnectionsConsumed"] = o.NetworkConnectionsConsumed
	}
	if !IsNil(o.LoadBalancersConsumed) {
		toSerialize["loadBalancersConsumed"] = o.LoadBalancersConsumed
	}
	if !IsNil(o.LoadbalancerWorkersConsumed) {
		toSerialize["loadbalancerWorkersConsumed"] = o.LoadbalancerWorkersConsumed
	}
	return toSerialize, nil
}

type NullableUsage struct {
	value *Usage
	isSet bool
}

func (v NullableUsage) Get() *Usage {
	return v.value
}

func (v *NullableUsage) Set(val *Usage) {
	v.value = val
	v.isSet = true
}

func (v NullableUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsage(val *Usage) *NullableUsage {
	return &NullableUsage{value: val, isSet: true}
}

func (v NullableUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


