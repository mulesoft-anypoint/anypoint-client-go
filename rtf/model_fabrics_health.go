/*
Runtime Fabrics API

Runtime Fabrics API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rtf

import (
	"encoding/json"
)

// checks if the FabricsHealth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricsHealth{}

// FabricsHealth struct for FabricsHealth
type FabricsHealth struct {
	ClusterMonitoring *FabricsHealthStatus `json:"clusterMonitoring,omitempty"`
	ManageDeployments *FabricsHealthStatus `json:"manageDeployments,omitempty"`
	LoadBalancing *FabricsHealthStatus `json:"loadBalancing,omitempty"`
	AnypointMonitoring *FabricsHealthStatus `json:"anypointMonitoring,omitempty"`
	ExternalLogForwarding *FabricsHealthStatus `json:"externalLogForwarding,omitempty"`
	Appliance *FabricsHealthStatus `json:"appliance,omitempty"`
	Infrastructure *FabricsHealthStatus `json:"infrastructure,omitempty"`
	PersistentGateway *FabricsHealthStatus `json:"persistentGateway,omitempty"`
}

// NewFabricsHealth instantiates a new FabricsHealth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricsHealth() *FabricsHealth {
	this := FabricsHealth{}
	return &this
}

// NewFabricsHealthWithDefaults instantiates a new FabricsHealth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricsHealthWithDefaults() *FabricsHealth {
	this := FabricsHealth{}
	return &this
}

// GetClusterMonitoring returns the ClusterMonitoring field value if set, zero value otherwise.
func (o *FabricsHealth) GetClusterMonitoring() FabricsHealthStatus {
	if o == nil || IsNil(o.ClusterMonitoring) {
		var ret FabricsHealthStatus
		return ret
	}
	return *o.ClusterMonitoring
}

// GetClusterMonitoringOk returns a tuple with the ClusterMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsHealth) GetClusterMonitoringOk() (*FabricsHealthStatus, bool) {
	if o == nil || IsNil(o.ClusterMonitoring) {
		return nil, false
	}
	return o.ClusterMonitoring, true
}

// HasClusterMonitoring returns a boolean if a field has been set.
func (o *FabricsHealth) HasClusterMonitoring() bool {
	if o != nil && !IsNil(o.ClusterMonitoring) {
		return true
	}

	return false
}

// SetClusterMonitoring gets a reference to the given FabricsHealthStatus and assigns it to the ClusterMonitoring field.
func (o *FabricsHealth) SetClusterMonitoring(v FabricsHealthStatus) {
	o.ClusterMonitoring = &v
}

// GetManageDeployments returns the ManageDeployments field value if set, zero value otherwise.
func (o *FabricsHealth) GetManageDeployments() FabricsHealthStatus {
	if o == nil || IsNil(o.ManageDeployments) {
		var ret FabricsHealthStatus
		return ret
	}
	return *o.ManageDeployments
}

// GetManageDeploymentsOk returns a tuple with the ManageDeployments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsHealth) GetManageDeploymentsOk() (*FabricsHealthStatus, bool) {
	if o == nil || IsNil(o.ManageDeployments) {
		return nil, false
	}
	return o.ManageDeployments, true
}

// HasManageDeployments returns a boolean if a field has been set.
func (o *FabricsHealth) HasManageDeployments() bool {
	if o != nil && !IsNil(o.ManageDeployments) {
		return true
	}

	return false
}

// SetManageDeployments gets a reference to the given FabricsHealthStatus and assigns it to the ManageDeployments field.
func (o *FabricsHealth) SetManageDeployments(v FabricsHealthStatus) {
	o.ManageDeployments = &v
}

// GetLoadBalancing returns the LoadBalancing field value if set, zero value otherwise.
func (o *FabricsHealth) GetLoadBalancing() FabricsHealthStatus {
	if o == nil || IsNil(o.LoadBalancing) {
		var ret FabricsHealthStatus
		return ret
	}
	return *o.LoadBalancing
}

// GetLoadBalancingOk returns a tuple with the LoadBalancing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsHealth) GetLoadBalancingOk() (*FabricsHealthStatus, bool) {
	if o == nil || IsNil(o.LoadBalancing) {
		return nil, false
	}
	return o.LoadBalancing, true
}

// HasLoadBalancing returns a boolean if a field has been set.
func (o *FabricsHealth) HasLoadBalancing() bool {
	if o != nil && !IsNil(o.LoadBalancing) {
		return true
	}

	return false
}

// SetLoadBalancing gets a reference to the given FabricsHealthStatus and assigns it to the LoadBalancing field.
func (o *FabricsHealth) SetLoadBalancing(v FabricsHealthStatus) {
	o.LoadBalancing = &v
}

// GetAnypointMonitoring returns the AnypointMonitoring field value if set, zero value otherwise.
func (o *FabricsHealth) GetAnypointMonitoring() FabricsHealthStatus {
	if o == nil || IsNil(o.AnypointMonitoring) {
		var ret FabricsHealthStatus
		return ret
	}
	return *o.AnypointMonitoring
}

// GetAnypointMonitoringOk returns a tuple with the AnypointMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsHealth) GetAnypointMonitoringOk() (*FabricsHealthStatus, bool) {
	if o == nil || IsNil(o.AnypointMonitoring) {
		return nil, false
	}
	return o.AnypointMonitoring, true
}

// HasAnypointMonitoring returns a boolean if a field has been set.
func (o *FabricsHealth) HasAnypointMonitoring() bool {
	if o != nil && !IsNil(o.AnypointMonitoring) {
		return true
	}

	return false
}

// SetAnypointMonitoring gets a reference to the given FabricsHealthStatus and assigns it to the AnypointMonitoring field.
func (o *FabricsHealth) SetAnypointMonitoring(v FabricsHealthStatus) {
	o.AnypointMonitoring = &v
}

// GetExternalLogForwarding returns the ExternalLogForwarding field value if set, zero value otherwise.
func (o *FabricsHealth) GetExternalLogForwarding() FabricsHealthStatus {
	if o == nil || IsNil(o.ExternalLogForwarding) {
		var ret FabricsHealthStatus
		return ret
	}
	return *o.ExternalLogForwarding
}

// GetExternalLogForwardingOk returns a tuple with the ExternalLogForwarding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsHealth) GetExternalLogForwardingOk() (*FabricsHealthStatus, bool) {
	if o == nil || IsNil(o.ExternalLogForwarding) {
		return nil, false
	}
	return o.ExternalLogForwarding, true
}

// HasExternalLogForwarding returns a boolean if a field has been set.
func (o *FabricsHealth) HasExternalLogForwarding() bool {
	if o != nil && !IsNil(o.ExternalLogForwarding) {
		return true
	}

	return false
}

// SetExternalLogForwarding gets a reference to the given FabricsHealthStatus and assigns it to the ExternalLogForwarding field.
func (o *FabricsHealth) SetExternalLogForwarding(v FabricsHealthStatus) {
	o.ExternalLogForwarding = &v
}

// GetAppliance returns the Appliance field value if set, zero value otherwise.
func (o *FabricsHealth) GetAppliance() FabricsHealthStatus {
	if o == nil || IsNil(o.Appliance) {
		var ret FabricsHealthStatus
		return ret
	}
	return *o.Appliance
}

// GetApplianceOk returns a tuple with the Appliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsHealth) GetApplianceOk() (*FabricsHealthStatus, bool) {
	if o == nil || IsNil(o.Appliance) {
		return nil, false
	}
	return o.Appliance, true
}

// HasAppliance returns a boolean if a field has been set.
func (o *FabricsHealth) HasAppliance() bool {
	if o != nil && !IsNil(o.Appliance) {
		return true
	}

	return false
}

// SetAppliance gets a reference to the given FabricsHealthStatus and assigns it to the Appliance field.
func (o *FabricsHealth) SetAppliance(v FabricsHealthStatus) {
	o.Appliance = &v
}

// GetInfrastructure returns the Infrastructure field value if set, zero value otherwise.
func (o *FabricsHealth) GetInfrastructure() FabricsHealthStatus {
	if o == nil || IsNil(o.Infrastructure) {
		var ret FabricsHealthStatus
		return ret
	}
	return *o.Infrastructure
}

// GetInfrastructureOk returns a tuple with the Infrastructure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsHealth) GetInfrastructureOk() (*FabricsHealthStatus, bool) {
	if o == nil || IsNil(o.Infrastructure) {
		return nil, false
	}
	return o.Infrastructure, true
}

// HasInfrastructure returns a boolean if a field has been set.
func (o *FabricsHealth) HasInfrastructure() bool {
	if o != nil && !IsNil(o.Infrastructure) {
		return true
	}

	return false
}

// SetInfrastructure gets a reference to the given FabricsHealthStatus and assigns it to the Infrastructure field.
func (o *FabricsHealth) SetInfrastructure(v FabricsHealthStatus) {
	o.Infrastructure = &v
}

// GetPersistentGateway returns the PersistentGateway field value if set, zero value otherwise.
func (o *FabricsHealth) GetPersistentGateway() FabricsHealthStatus {
	if o == nil || IsNil(o.PersistentGateway) {
		var ret FabricsHealthStatus
		return ret
	}
	return *o.PersistentGateway
}

// GetPersistentGatewayOk returns a tuple with the PersistentGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsHealth) GetPersistentGatewayOk() (*FabricsHealthStatus, bool) {
	if o == nil || IsNil(o.PersistentGateway) {
		return nil, false
	}
	return o.PersistentGateway, true
}

// HasPersistentGateway returns a boolean if a field has been set.
func (o *FabricsHealth) HasPersistentGateway() bool {
	if o != nil && !IsNil(o.PersistentGateway) {
		return true
	}

	return false
}

// SetPersistentGateway gets a reference to the given FabricsHealthStatus and assigns it to the PersistentGateway field.
func (o *FabricsHealth) SetPersistentGateway(v FabricsHealthStatus) {
	o.PersistentGateway = &v
}

func (o FabricsHealth) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricsHealth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClusterMonitoring) {
		toSerialize["clusterMonitoring"] = o.ClusterMonitoring
	}
	if !IsNil(o.ManageDeployments) {
		toSerialize["manageDeployments"] = o.ManageDeployments
	}
	if !IsNil(o.LoadBalancing) {
		toSerialize["loadBalancing"] = o.LoadBalancing
	}
	if !IsNil(o.AnypointMonitoring) {
		toSerialize["anypointMonitoring"] = o.AnypointMonitoring
	}
	if !IsNil(o.ExternalLogForwarding) {
		toSerialize["externalLogForwarding"] = o.ExternalLogForwarding
	}
	if !IsNil(o.Appliance) {
		toSerialize["appliance"] = o.Appliance
	}
	if !IsNil(o.Infrastructure) {
		toSerialize["infrastructure"] = o.Infrastructure
	}
	if !IsNil(o.PersistentGateway) {
		toSerialize["persistentGateway"] = o.PersistentGateway
	}
	return toSerialize, nil
}

type NullableFabricsHealth struct {
	value *FabricsHealth
	isSet bool
}

func (v NullableFabricsHealth) Get() *FabricsHealth {
	return v.value
}

func (v *NullableFabricsHealth) Set(val *FabricsHealth) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricsHealth) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricsHealth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricsHealth(val *FabricsHealth) *NullableFabricsHealth {
	return &NullableFabricsHealth{value: val, isSet: true}
}

func (v NullableFabricsHealth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricsHealth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


