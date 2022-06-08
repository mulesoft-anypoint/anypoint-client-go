/*
 * Dedicated Load Balancer API
 *
 * Description of the DLB API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dlb

import (
	"encoding/json"
)

// DlbExtras struct for DlbExtras
type DlbExtras struct {
	// The dlb Id
	Id *string `json:"id,omitempty"`
	// dlb domain
	Domain *string `json:"domain,omitempty"`
	// the dlb deployment id
	DeploymentId *string `json:"deploymentId,omitempty"`
	InstanceConfig *InstanceConfig `json:"instanceConfig,omitempty"`
	IpAddresses *[]string `json:"ipAddresses,omitempty"`
	StaticIPsDisabled *bool `json:"staticIPsDisabled,omitempty"`
	VpcId *string `json:"vpcId,omitempty"`
	Workers *int32 `json:"workers,omitempty"`
	DefaultCipherSuite *string `json:"defaultCipherSuite,omitempty"`
	KeepUrlEncoding *bool `json:"keepUrlEncoding,omitempty"`
	UpstreamTlsv12 *bool `json:"upstreamTlsv12,omitempty"`
	ProxyReadTimeout *int32 `json:"proxyReadTimeout,omitempty"`
	IpAddressesInfo *[]DlbExtrasIpAddressesInfo `json:"ipAddressesInfo,omitempty"`
	DoubleStaticIps *bool `json:"doubleStaticIps,omitempty"`
	EnableStreaming *bool `json:"enableStreaming,omitempty"`
	ForwardClientCertificate *bool `json:"forwardClientCertificate,omitempty"`
}

// NewDlbExtras instantiates a new DlbExtras object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDlbExtras() *DlbExtras {
	this := DlbExtras{}
	return &this
}

// NewDlbExtrasWithDefaults instantiates a new DlbExtras object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDlbExtrasWithDefaults() *DlbExtras {
	this := DlbExtras{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DlbExtras) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlbExtras) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DlbExtras) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DlbExtras) SetId(v string) {
	o.Id = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *DlbExtras) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlbExtras) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *DlbExtras) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *DlbExtras) SetDomain(v string) {
	o.Domain = &v
}

// GetDeploymentId returns the DeploymentId field value if set, zero value otherwise.
func (o *DlbExtras) GetDeploymentId() string {
	if o == nil || o.DeploymentId == nil {
		var ret string
		return ret
	}
	return *o.DeploymentId
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlbExtras) GetDeploymentIdOk() (*string, bool) {
	if o == nil || o.DeploymentId == nil {
		return nil, false
	}
	return o.DeploymentId, true
}

// HasDeploymentId returns a boolean if a field has been set.
func (o *DlbExtras) HasDeploymentId() bool {
	if o != nil && o.DeploymentId != nil {
		return true
	}

	return false
}

// SetDeploymentId gets a reference to the given string and assigns it to the DeploymentId field.
func (o *DlbExtras) SetDeploymentId(v string) {
	o.DeploymentId = &v
}

// GetInstanceConfig returns the InstanceConfig field value if set, zero value otherwise.
func (o *DlbExtras) GetInstanceConfig() InstanceConfig {
	if o == nil || o.InstanceConfig == nil {
		var ret InstanceConfig
		return ret
	}
	return *o.InstanceConfig
}

// GetInstanceConfigOk returns a tuple with the InstanceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlbExtras) GetInstanceConfigOk() (*InstanceConfig, bool) {
	if o == nil || o.InstanceConfig == nil {
		return nil, false
	}
	return o.InstanceConfig, true
}

// HasInstanceConfig returns a boolean if a field has been set.
func (o *DlbExtras) HasInstanceConfig() bool {
	if o != nil && o.InstanceConfig != nil {
		return true
	}

	return false
}

// SetInstanceConfig gets a reference to the given InstanceConfig and assigns it to the InstanceConfig field.
func (o *DlbExtras) SetInstanceConfig(v InstanceConfig) {
	o.InstanceConfig = &v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *DlbExtras) GetIpAddresses() []string {
	if o == nil || o.IpAddresses == nil {
		var ret []string
		return ret
	}
	return *o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlbExtras) GetIpAddressesOk() (*[]string, bool) {
	if o == nil || o.IpAddresses == nil {
		return nil, false
	}
	return o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *DlbExtras) HasIpAddresses() bool {
	if o != nil && o.IpAddresses != nil {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *DlbExtras) SetIpAddresses(v []string) {
	o.IpAddresses = &v
}

// GetStaticIPsDisabled returns the StaticIPsDisabled field value if set, zero value otherwise.
func (o *DlbExtras) GetStaticIPsDisabled() bool {
	if o == nil || o.StaticIPsDisabled == nil {
		var ret bool
		return ret
	}
	return *o.StaticIPsDisabled
}

// GetStaticIPsDisabledOk returns a tuple with the StaticIPsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlbExtras) GetStaticIPsDisabledOk() (*bool, bool) {
	if o == nil || o.StaticIPsDisabled == nil {
		return nil, false
	}
	return o.StaticIPsDisabled, true
}

// HasStaticIPsDisabled returns a boolean if a field has been set.
func (o *DlbExtras) HasStaticIPsDisabled() bool {
	if o != nil && o.StaticIPsDisabled != nil {
		return true
	}

	return false
}

// SetStaticIPsDisabled gets a reference to the given bool and assigns it to the StaticIPsDisabled field.
func (o *DlbExtras) SetStaticIPsDisabled(v bool) {
	o.StaticIPsDisabled = &v
}

// GetVpcId returns the VpcId field value if set, zero value otherwise.
func (o *DlbExtras) GetVpcId() string {
	if o == nil || o.VpcId == nil {
		var ret string
		return ret
	}
	return *o.VpcId
}

// GetVpcIdOk returns a tuple with the VpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlbExtras) GetVpcIdOk() (*string, bool) {
	if o == nil || o.VpcId == nil {
		return nil, false
	}
	return o.VpcId, true
}

// HasVpcId returns a boolean if a field has been set.
func (o *DlbExtras) HasVpcId() bool {
	if o != nil && o.VpcId != nil {
		return true
	}

	return false
}

// SetVpcId gets a reference to the given string and assigns it to the VpcId field.
func (o *DlbExtras) SetVpcId(v string) {
	o.VpcId = &v
}

// GetWorkers returns the Workers field value if set, zero value otherwise.
func (o *DlbExtras) GetWorkers() int32 {
	if o == nil || o.Workers == nil {
		var ret int32
		return ret
	}
	return *o.Workers
}

// GetWorkersOk returns a tuple with the Workers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlbExtras) GetWorkersOk() (*int32, bool) {
	if o == nil || o.Workers == nil {
		return nil, false
	}
	return o.Workers, true
}

// HasWorkers returns a boolean if a field has been set.
func (o *DlbExtras) HasWorkers() bool {
	if o != nil && o.Workers != nil {
		return true
	}

	return false
}

// SetWorkers gets a reference to the given int32 and assigns it to the Workers field.
func (o *DlbExtras) SetWorkers(v int32) {
	o.Workers = &v
}

// GetDefaultCipherSuite returns the DefaultCipherSuite field value if set, zero value otherwise.
func (o *DlbExtras) GetDefaultCipherSuite() string {
	if o == nil || o.DefaultCipherSuite == nil {
		var ret string
		return ret
	}
	return *o.DefaultCipherSuite
}

// GetDefaultCipherSuiteOk returns a tuple with the DefaultCipherSuite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlbExtras) GetDefaultCipherSuiteOk() (*string, bool) {
	if o == nil || o.DefaultCipherSuite == nil {
		return nil, false
	}
	return o.DefaultCipherSuite, true
}

// HasDefaultCipherSuite returns a boolean if a field has been set.
func (o *DlbExtras) HasDefaultCipherSuite() bool {
	if o != nil && o.DefaultCipherSuite != nil {
		return true
	}

	return false
}

// SetDefaultCipherSuite gets a reference to the given string and assigns it to the DefaultCipherSuite field.
func (o *DlbExtras) SetDefaultCipherSuite(v string) {
	o.DefaultCipherSuite = &v
}

// GetKeepUrlEncoding returns the KeepUrlEncoding field value if set, zero value otherwise.
func (o *DlbExtras) GetKeepUrlEncoding() bool {
	if o == nil || o.KeepUrlEncoding == nil {
		var ret bool
		return ret
	}
	return *o.KeepUrlEncoding
}

// GetKeepUrlEncodingOk returns a tuple with the KeepUrlEncoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlbExtras) GetKeepUrlEncodingOk() (*bool, bool) {
	if o == nil || o.KeepUrlEncoding == nil {
		return nil, false
	}
	return o.KeepUrlEncoding, true
}

// HasKeepUrlEncoding returns a boolean if a field has been set.
func (o *DlbExtras) HasKeepUrlEncoding() bool {
	if o != nil && o.KeepUrlEncoding != nil {
		return true
	}

	return false
}

// SetKeepUrlEncoding gets a reference to the given bool and assigns it to the KeepUrlEncoding field.
func (o *DlbExtras) SetKeepUrlEncoding(v bool) {
	o.KeepUrlEncoding = &v
}

// GetUpstreamTlsv12 returns the UpstreamTlsv12 field value if set, zero value otherwise.
func (o *DlbExtras) GetUpstreamTlsv12() bool {
	if o == nil || o.UpstreamTlsv12 == nil {
		var ret bool
		return ret
	}
	return *o.UpstreamTlsv12
}

// GetUpstreamTlsv12Ok returns a tuple with the UpstreamTlsv12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlbExtras) GetUpstreamTlsv12Ok() (*bool, bool) {
	if o == nil || o.UpstreamTlsv12 == nil {
		return nil, false
	}
	return o.UpstreamTlsv12, true
}

// HasUpstreamTlsv12 returns a boolean if a field has been set.
func (o *DlbExtras) HasUpstreamTlsv12() bool {
	if o != nil && o.UpstreamTlsv12 != nil {
		return true
	}

	return false
}

// SetUpstreamTlsv12 gets a reference to the given bool and assigns it to the UpstreamTlsv12 field.
func (o *DlbExtras) SetUpstreamTlsv12(v bool) {
	o.UpstreamTlsv12 = &v
}

// GetProxyReadTimeout returns the ProxyReadTimeout field value if set, zero value otherwise.
func (o *DlbExtras) GetProxyReadTimeout() int32 {
	if o == nil || o.ProxyReadTimeout == nil {
		var ret int32
		return ret
	}
	return *o.ProxyReadTimeout
}

// GetProxyReadTimeoutOk returns a tuple with the ProxyReadTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlbExtras) GetProxyReadTimeoutOk() (*int32, bool) {
	if o == nil || o.ProxyReadTimeout == nil {
		return nil, false
	}
	return o.ProxyReadTimeout, true
}

// HasProxyReadTimeout returns a boolean if a field has been set.
func (o *DlbExtras) HasProxyReadTimeout() bool {
	if o != nil && o.ProxyReadTimeout != nil {
		return true
	}

	return false
}

// SetProxyReadTimeout gets a reference to the given int32 and assigns it to the ProxyReadTimeout field.
func (o *DlbExtras) SetProxyReadTimeout(v int32) {
	o.ProxyReadTimeout = &v
}

// GetIpAddressesInfo returns the IpAddressesInfo field value if set, zero value otherwise.
func (o *DlbExtras) GetIpAddressesInfo() []DlbExtrasIpAddressesInfo {
	if o == nil || o.IpAddressesInfo == nil {
		var ret []DlbExtrasIpAddressesInfo
		return ret
	}
	return *o.IpAddressesInfo
}

// GetIpAddressesInfoOk returns a tuple with the IpAddressesInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlbExtras) GetIpAddressesInfoOk() (*[]DlbExtrasIpAddressesInfo, bool) {
	if o == nil || o.IpAddressesInfo == nil {
		return nil, false
	}
	return o.IpAddressesInfo, true
}

// HasIpAddressesInfo returns a boolean if a field has been set.
func (o *DlbExtras) HasIpAddressesInfo() bool {
	if o != nil && o.IpAddressesInfo != nil {
		return true
	}

	return false
}

// SetIpAddressesInfo gets a reference to the given []DlbExtrasIpAddressesInfo and assigns it to the IpAddressesInfo field.
func (o *DlbExtras) SetIpAddressesInfo(v []DlbExtrasIpAddressesInfo) {
	o.IpAddressesInfo = &v
}

// GetDoubleStaticIps returns the DoubleStaticIps field value if set, zero value otherwise.
func (o *DlbExtras) GetDoubleStaticIps() bool {
	if o == nil || o.DoubleStaticIps == nil {
		var ret bool
		return ret
	}
	return *o.DoubleStaticIps
}

// GetDoubleStaticIpsOk returns a tuple with the DoubleStaticIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlbExtras) GetDoubleStaticIpsOk() (*bool, bool) {
	if o == nil || o.DoubleStaticIps == nil {
		return nil, false
	}
	return o.DoubleStaticIps, true
}

// HasDoubleStaticIps returns a boolean if a field has been set.
func (o *DlbExtras) HasDoubleStaticIps() bool {
	if o != nil && o.DoubleStaticIps != nil {
		return true
	}

	return false
}

// SetDoubleStaticIps gets a reference to the given bool and assigns it to the DoubleStaticIps field.
func (o *DlbExtras) SetDoubleStaticIps(v bool) {
	o.DoubleStaticIps = &v
}

// GetEnableStreaming returns the EnableStreaming field value if set, zero value otherwise.
func (o *DlbExtras) GetEnableStreaming() bool {
	if o == nil || o.EnableStreaming == nil {
		var ret bool
		return ret
	}
	return *o.EnableStreaming
}

// GetEnableStreamingOk returns a tuple with the EnableStreaming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlbExtras) GetEnableStreamingOk() (*bool, bool) {
	if o == nil || o.EnableStreaming == nil {
		return nil, false
	}
	return o.EnableStreaming, true
}

// HasEnableStreaming returns a boolean if a field has been set.
func (o *DlbExtras) HasEnableStreaming() bool {
	if o != nil && o.EnableStreaming != nil {
		return true
	}

	return false
}

// SetEnableStreaming gets a reference to the given bool and assigns it to the EnableStreaming field.
func (o *DlbExtras) SetEnableStreaming(v bool) {
	o.EnableStreaming = &v
}

// GetForwardClientCertificate returns the ForwardClientCertificate field value if set, zero value otherwise.
func (o *DlbExtras) GetForwardClientCertificate() bool {
	if o == nil || o.ForwardClientCertificate == nil {
		var ret bool
		return ret
	}
	return *o.ForwardClientCertificate
}

// GetForwardClientCertificateOk returns a tuple with the ForwardClientCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DlbExtras) GetForwardClientCertificateOk() (*bool, bool) {
	if o == nil || o.ForwardClientCertificate == nil {
		return nil, false
	}
	return o.ForwardClientCertificate, true
}

// HasForwardClientCertificate returns a boolean if a field has been set.
func (o *DlbExtras) HasForwardClientCertificate() bool {
	if o != nil && o.ForwardClientCertificate != nil {
		return true
	}

	return false
}

// SetForwardClientCertificate gets a reference to the given bool and assigns it to the ForwardClientCertificate field.
func (o *DlbExtras) SetForwardClientCertificate(v bool) {
	o.ForwardClientCertificate = &v
}

func (o DlbExtras) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.DeploymentId != nil {
		toSerialize["deploymentId"] = o.DeploymentId
	}
	if o.InstanceConfig != nil {
		toSerialize["instanceConfig"] = o.InstanceConfig
	}
	if o.IpAddresses != nil {
		toSerialize["ipAddresses"] = o.IpAddresses
	}
	if o.StaticIPsDisabled != nil {
		toSerialize["staticIPsDisabled"] = o.StaticIPsDisabled
	}
	if o.VpcId != nil {
		toSerialize["vpcId"] = o.VpcId
	}
	if o.Workers != nil {
		toSerialize["workers"] = o.Workers
	}
	if o.DefaultCipherSuite != nil {
		toSerialize["defaultCipherSuite"] = o.DefaultCipherSuite
	}
	if o.KeepUrlEncoding != nil {
		toSerialize["keepUrlEncoding"] = o.KeepUrlEncoding
	}
	if o.UpstreamTlsv12 != nil {
		toSerialize["upstreamTlsv12"] = o.UpstreamTlsv12
	}
	if o.ProxyReadTimeout != nil {
		toSerialize["proxyReadTimeout"] = o.ProxyReadTimeout
	}
	if o.IpAddressesInfo != nil {
		toSerialize["ipAddressesInfo"] = o.IpAddressesInfo
	}
	if o.DoubleStaticIps != nil {
		toSerialize["doubleStaticIps"] = o.DoubleStaticIps
	}
	if o.EnableStreaming != nil {
		toSerialize["enableStreaming"] = o.EnableStreaming
	}
	if o.ForwardClientCertificate != nil {
		toSerialize["forwardClientCertificate"] = o.ForwardClientCertificate
	}
	return json.Marshal(toSerialize)
}

type NullableDlbExtras struct {
	value *DlbExtras
	isSet bool
}

func (v NullableDlbExtras) Get() *DlbExtras {
	return v.value
}

func (v *NullableDlbExtras) Set(val *DlbExtras) {
	v.value = val
	v.isSet = true
}

func (v NullableDlbExtras) IsSet() bool {
	return v.isSet
}

func (v *NullableDlbExtras) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDlbExtras(val *DlbExtras) *NullableDlbExtras {
	return &NullableDlbExtras{value: val, isSet: true}
}

func (v NullableDlbExtras) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDlbExtras) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


