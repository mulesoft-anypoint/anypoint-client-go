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

// checks if the MonitoringCenterEntitlement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitoringCenterEntitlement{}

// MonitoringCenterEntitlement Monitoring Center entitlements
type MonitoringCenterEntitlement struct {
	// Product SKU
	ProductSKU *int32 `json:"productSKU,omitempty"`
	// Data region
	DataRegion *string `json:"dataRegion,omitempty"`
	// Storage base
	StorageBase *float32 `json:"storageBase,omitempty"`
	// Storage add-on
	StorageAddOn *float32 `json:"storageAddOn,omitempty"`
	// Production units
	ProductionUnits *float32 `json:"productionUnits,omitempty"`
	// Pre-production units
	PreProductionUnits *float32 `json:"preProductionUnits,omitempty"`
}

// NewMonitoringCenterEntitlement instantiates a new MonitoringCenterEntitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringCenterEntitlement() *MonitoringCenterEntitlement {
	this := MonitoringCenterEntitlement{}
	return &this
}

// NewMonitoringCenterEntitlementWithDefaults instantiates a new MonitoringCenterEntitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringCenterEntitlementWithDefaults() *MonitoringCenterEntitlement {
	this := MonitoringCenterEntitlement{}
	return &this
}

// GetProductSKU returns the ProductSKU field value if set, zero value otherwise.
func (o *MonitoringCenterEntitlement) GetProductSKU() int32 {
	if o == nil || IsNil(o.ProductSKU) {
		var ret int32
		return ret
	}
	return *o.ProductSKU
}

// GetProductSKUOk returns a tuple with the ProductSKU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringCenterEntitlement) GetProductSKUOk() (*int32, bool) {
	if o == nil || IsNil(o.ProductSKU) {
		return nil, false
	}
	return o.ProductSKU, true
}

// HasProductSKU returns a boolean if a field has been set.
func (o *MonitoringCenterEntitlement) HasProductSKU() bool {
	if o != nil && !IsNil(o.ProductSKU) {
		return true
	}

	return false
}

// SetProductSKU gets a reference to the given int32 and assigns it to the ProductSKU field.
func (o *MonitoringCenterEntitlement) SetProductSKU(v int32) {
	o.ProductSKU = &v
}

// GetDataRegion returns the DataRegion field value if set, zero value otherwise.
func (o *MonitoringCenterEntitlement) GetDataRegion() string {
	if o == nil || IsNil(o.DataRegion) {
		var ret string
		return ret
	}
	return *o.DataRegion
}

// GetDataRegionOk returns a tuple with the DataRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringCenterEntitlement) GetDataRegionOk() (*string, bool) {
	if o == nil || IsNil(o.DataRegion) {
		return nil, false
	}
	return o.DataRegion, true
}

// HasDataRegion returns a boolean if a field has been set.
func (o *MonitoringCenterEntitlement) HasDataRegion() bool {
	if o != nil && !IsNil(o.DataRegion) {
		return true
	}

	return false
}

// SetDataRegion gets a reference to the given string and assigns it to the DataRegion field.
func (o *MonitoringCenterEntitlement) SetDataRegion(v string) {
	o.DataRegion = &v
}

// GetStorageBase returns the StorageBase field value if set, zero value otherwise.
func (o *MonitoringCenterEntitlement) GetStorageBase() float32 {
	if o == nil || IsNil(o.StorageBase) {
		var ret float32
		return ret
	}
	return *o.StorageBase
}

// GetStorageBaseOk returns a tuple with the StorageBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringCenterEntitlement) GetStorageBaseOk() (*float32, bool) {
	if o == nil || IsNil(o.StorageBase) {
		return nil, false
	}
	return o.StorageBase, true
}

// HasStorageBase returns a boolean if a field has been set.
func (o *MonitoringCenterEntitlement) HasStorageBase() bool {
	if o != nil && !IsNil(o.StorageBase) {
		return true
	}

	return false
}

// SetStorageBase gets a reference to the given float32 and assigns it to the StorageBase field.
func (o *MonitoringCenterEntitlement) SetStorageBase(v float32) {
	o.StorageBase = &v
}

// GetStorageAddOn returns the StorageAddOn field value if set, zero value otherwise.
func (o *MonitoringCenterEntitlement) GetStorageAddOn() float32 {
	if o == nil || IsNil(o.StorageAddOn) {
		var ret float32
		return ret
	}
	return *o.StorageAddOn
}

// GetStorageAddOnOk returns a tuple with the StorageAddOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringCenterEntitlement) GetStorageAddOnOk() (*float32, bool) {
	if o == nil || IsNil(o.StorageAddOn) {
		return nil, false
	}
	return o.StorageAddOn, true
}

// HasStorageAddOn returns a boolean if a field has been set.
func (o *MonitoringCenterEntitlement) HasStorageAddOn() bool {
	if o != nil && !IsNil(o.StorageAddOn) {
		return true
	}

	return false
}

// SetStorageAddOn gets a reference to the given float32 and assigns it to the StorageAddOn field.
func (o *MonitoringCenterEntitlement) SetStorageAddOn(v float32) {
	o.StorageAddOn = &v
}

// GetProductionUnits returns the ProductionUnits field value if set, zero value otherwise.
func (o *MonitoringCenterEntitlement) GetProductionUnits() float32 {
	if o == nil || IsNil(o.ProductionUnits) {
		var ret float32
		return ret
	}
	return *o.ProductionUnits
}

// GetProductionUnitsOk returns a tuple with the ProductionUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringCenterEntitlement) GetProductionUnitsOk() (*float32, bool) {
	if o == nil || IsNil(o.ProductionUnits) {
		return nil, false
	}
	return o.ProductionUnits, true
}

// HasProductionUnits returns a boolean if a field has been set.
func (o *MonitoringCenterEntitlement) HasProductionUnits() bool {
	if o != nil && !IsNil(o.ProductionUnits) {
		return true
	}

	return false
}

// SetProductionUnits gets a reference to the given float32 and assigns it to the ProductionUnits field.
func (o *MonitoringCenterEntitlement) SetProductionUnits(v float32) {
	o.ProductionUnits = &v
}

// GetPreProductionUnits returns the PreProductionUnits field value if set, zero value otherwise.
func (o *MonitoringCenterEntitlement) GetPreProductionUnits() float32 {
	if o == nil || IsNil(o.PreProductionUnits) {
		var ret float32
		return ret
	}
	return *o.PreProductionUnits
}

// GetPreProductionUnitsOk returns a tuple with the PreProductionUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringCenterEntitlement) GetPreProductionUnitsOk() (*float32, bool) {
	if o == nil || IsNil(o.PreProductionUnits) {
		return nil, false
	}
	return o.PreProductionUnits, true
}

// HasPreProductionUnits returns a boolean if a field has been set.
func (o *MonitoringCenterEntitlement) HasPreProductionUnits() bool {
	if o != nil && !IsNil(o.PreProductionUnits) {
		return true
	}

	return false
}

// SetPreProductionUnits gets a reference to the given float32 and assigns it to the PreProductionUnits field.
func (o *MonitoringCenterEntitlement) SetPreProductionUnits(v float32) {
	o.PreProductionUnits = &v
}

func (o MonitoringCenterEntitlement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitoringCenterEntitlement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductSKU) {
		toSerialize["productSKU"] = o.ProductSKU
	}
	if !IsNil(o.DataRegion) {
		toSerialize["dataRegion"] = o.DataRegion
	}
	if !IsNil(o.StorageBase) {
		toSerialize["storageBase"] = o.StorageBase
	}
	if !IsNil(o.StorageAddOn) {
		toSerialize["storageAddOn"] = o.StorageAddOn
	}
	if !IsNil(o.ProductionUnits) {
		toSerialize["productionUnits"] = o.ProductionUnits
	}
	if !IsNil(o.PreProductionUnits) {
		toSerialize["preProductionUnits"] = o.PreProductionUnits
	}
	return toSerialize, nil
}

type NullableMonitoringCenterEntitlement struct {
	value *MonitoringCenterEntitlement
	isSet bool
}

func (v NullableMonitoringCenterEntitlement) Get() *MonitoringCenterEntitlement {
	return v.value
}

func (v *NullableMonitoringCenterEntitlement) Set(val *MonitoringCenterEntitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringCenterEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringCenterEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringCenterEntitlement(val *MonitoringCenterEntitlement) *NullableMonitoringCenterEntitlement {
	return &NullableMonitoringCenterEntitlement{value: val, isSet: true}
}

func (v NullableMonitoringCenterEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringCenterEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


