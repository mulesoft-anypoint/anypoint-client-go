/*
 * API Manager API
 *
 * API Manager API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim

import (
	"encoding/json"
)

// EndpointPostBody struct for EndpointPostBody
type EndpointPostBody struct {
	DeploymentType *string `json:"deploymentType,omitempty"`
	MuleVersion4OrAbove *bool `json:"muleVersion4OrAbove,omitempty"`
	Uri *string `json:"uri,omitempty"`
	Type *string `json:"type,omitempty"`
	IsCloudHub NullableString `json:"isCloudHub,omitempty"`
	ProxyUri NullableString `json:"proxyUri,omitempty"`
	ProxyRegistrationUri NullableString `json:"proxyRegistrationUri,omitempty"`
	ReferencesUserDomain NullableString `json:"referencesUserDomain,omitempty"`
	ResponseTimeout NullableString `json:"responseTimeout,omitempty"`
	TlsContexts NullableEndpointPostBodyTlsContexts `json:"tlsContexts,omitempty"`
}

// NewEndpointPostBody instantiates a new EndpointPostBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointPostBody() *EndpointPostBody {
	this := EndpointPostBody{}
	return &this
}

// NewEndpointPostBodyWithDefaults instantiates a new EndpointPostBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointPostBodyWithDefaults() *EndpointPostBody {
	this := EndpointPostBody{}
	return &this
}

// GetDeploymentType returns the DeploymentType field value if set, zero value otherwise.
func (o *EndpointPostBody) GetDeploymentType() string {
	if o == nil || o.DeploymentType == nil {
		var ret string
		return ret
	}
	return *o.DeploymentType
}

// GetDeploymentTypeOk returns a tuple with the DeploymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointPostBody) GetDeploymentTypeOk() (*string, bool) {
	if o == nil || o.DeploymentType == nil {
		return nil, false
	}
	return o.DeploymentType, true
}

// HasDeploymentType returns a boolean if a field has been set.
func (o *EndpointPostBody) HasDeploymentType() bool {
	if o != nil && o.DeploymentType != nil {
		return true
	}

	return false
}

// SetDeploymentType gets a reference to the given string and assigns it to the DeploymentType field.
func (o *EndpointPostBody) SetDeploymentType(v string) {
	o.DeploymentType = &v
}

// GetMuleVersion4OrAbove returns the MuleVersion4OrAbove field value if set, zero value otherwise.
func (o *EndpointPostBody) GetMuleVersion4OrAbove() bool {
	if o == nil || o.MuleVersion4OrAbove == nil {
		var ret bool
		return ret
	}
	return *o.MuleVersion4OrAbove
}

// GetMuleVersion4OrAboveOk returns a tuple with the MuleVersion4OrAbove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointPostBody) GetMuleVersion4OrAboveOk() (*bool, bool) {
	if o == nil || o.MuleVersion4OrAbove == nil {
		return nil, false
	}
	return o.MuleVersion4OrAbove, true
}

// HasMuleVersion4OrAbove returns a boolean if a field has been set.
func (o *EndpointPostBody) HasMuleVersion4OrAbove() bool {
	if o != nil && o.MuleVersion4OrAbove != nil {
		return true
	}

	return false
}

// SetMuleVersion4OrAbove gets a reference to the given bool and assigns it to the MuleVersion4OrAbove field.
func (o *EndpointPostBody) SetMuleVersion4OrAbove(v bool) {
	o.MuleVersion4OrAbove = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *EndpointPostBody) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointPostBody) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *EndpointPostBody) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *EndpointPostBody) SetUri(v string) {
	o.Uri = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EndpointPostBody) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointPostBody) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EndpointPostBody) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EndpointPostBody) SetType(v string) {
	o.Type = &v
}

// GetIsCloudHub returns the IsCloudHub field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointPostBody) GetIsCloudHub() string {
	if o == nil || o.IsCloudHub.Get() == nil {
		var ret string
		return ret
	}
	return *o.IsCloudHub.Get()
}

// GetIsCloudHubOk returns a tuple with the IsCloudHub field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointPostBody) GetIsCloudHubOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsCloudHub.Get(), o.IsCloudHub.IsSet()
}

// HasIsCloudHub returns a boolean if a field has been set.
func (o *EndpointPostBody) HasIsCloudHub() bool {
	if o != nil && o.IsCloudHub.IsSet() {
		return true
	}

	return false
}

// SetIsCloudHub gets a reference to the given NullableString and assigns it to the IsCloudHub field.
func (o *EndpointPostBody) SetIsCloudHub(v string) {
	o.IsCloudHub.Set(&v)
}
// SetIsCloudHubNil sets the value for IsCloudHub to be an explicit nil
func (o *EndpointPostBody) SetIsCloudHubNil() {
	o.IsCloudHub.Set(nil)
}

// UnsetIsCloudHub ensures that no value is present for IsCloudHub, not even an explicit nil
func (o *EndpointPostBody) UnsetIsCloudHub() {
	o.IsCloudHub.Unset()
}

// GetProxyUri returns the ProxyUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointPostBody) GetProxyUri() string {
	if o == nil || o.ProxyUri.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProxyUri.Get()
}

// GetProxyUriOk returns a tuple with the ProxyUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointPostBody) GetProxyUriOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProxyUri.Get(), o.ProxyUri.IsSet()
}

// HasProxyUri returns a boolean if a field has been set.
func (o *EndpointPostBody) HasProxyUri() bool {
	if o != nil && o.ProxyUri.IsSet() {
		return true
	}

	return false
}

// SetProxyUri gets a reference to the given NullableString and assigns it to the ProxyUri field.
func (o *EndpointPostBody) SetProxyUri(v string) {
	o.ProxyUri.Set(&v)
}
// SetProxyUriNil sets the value for ProxyUri to be an explicit nil
func (o *EndpointPostBody) SetProxyUriNil() {
	o.ProxyUri.Set(nil)
}

// UnsetProxyUri ensures that no value is present for ProxyUri, not even an explicit nil
func (o *EndpointPostBody) UnsetProxyUri() {
	o.ProxyUri.Unset()
}

// GetProxyRegistrationUri returns the ProxyRegistrationUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointPostBody) GetProxyRegistrationUri() string {
	if o == nil || o.ProxyRegistrationUri.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProxyRegistrationUri.Get()
}

// GetProxyRegistrationUriOk returns a tuple with the ProxyRegistrationUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointPostBody) GetProxyRegistrationUriOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProxyRegistrationUri.Get(), o.ProxyRegistrationUri.IsSet()
}

// HasProxyRegistrationUri returns a boolean if a field has been set.
func (o *EndpointPostBody) HasProxyRegistrationUri() bool {
	if o != nil && o.ProxyRegistrationUri.IsSet() {
		return true
	}

	return false
}

// SetProxyRegistrationUri gets a reference to the given NullableString and assigns it to the ProxyRegistrationUri field.
func (o *EndpointPostBody) SetProxyRegistrationUri(v string) {
	o.ProxyRegistrationUri.Set(&v)
}
// SetProxyRegistrationUriNil sets the value for ProxyRegistrationUri to be an explicit nil
func (o *EndpointPostBody) SetProxyRegistrationUriNil() {
	o.ProxyRegistrationUri.Set(nil)
}

// UnsetProxyRegistrationUri ensures that no value is present for ProxyRegistrationUri, not even an explicit nil
func (o *EndpointPostBody) UnsetProxyRegistrationUri() {
	o.ProxyRegistrationUri.Unset()
}

// GetReferencesUserDomain returns the ReferencesUserDomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointPostBody) GetReferencesUserDomain() string {
	if o == nil || o.ReferencesUserDomain.Get() == nil {
		var ret string
		return ret
	}
	return *o.ReferencesUserDomain.Get()
}

// GetReferencesUserDomainOk returns a tuple with the ReferencesUserDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointPostBody) GetReferencesUserDomainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReferencesUserDomain.Get(), o.ReferencesUserDomain.IsSet()
}

// HasReferencesUserDomain returns a boolean if a field has been set.
func (o *EndpointPostBody) HasReferencesUserDomain() bool {
	if o != nil && o.ReferencesUserDomain.IsSet() {
		return true
	}

	return false
}

// SetReferencesUserDomain gets a reference to the given NullableString and assigns it to the ReferencesUserDomain field.
func (o *EndpointPostBody) SetReferencesUserDomain(v string) {
	o.ReferencesUserDomain.Set(&v)
}
// SetReferencesUserDomainNil sets the value for ReferencesUserDomain to be an explicit nil
func (o *EndpointPostBody) SetReferencesUserDomainNil() {
	o.ReferencesUserDomain.Set(nil)
}

// UnsetReferencesUserDomain ensures that no value is present for ReferencesUserDomain, not even an explicit nil
func (o *EndpointPostBody) UnsetReferencesUserDomain() {
	o.ReferencesUserDomain.Unset()
}

// GetResponseTimeout returns the ResponseTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointPostBody) GetResponseTimeout() string {
	if o == nil || o.ResponseTimeout.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResponseTimeout.Get()
}

// GetResponseTimeoutOk returns a tuple with the ResponseTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointPostBody) GetResponseTimeoutOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResponseTimeout.Get(), o.ResponseTimeout.IsSet()
}

// HasResponseTimeout returns a boolean if a field has been set.
func (o *EndpointPostBody) HasResponseTimeout() bool {
	if o != nil && o.ResponseTimeout.IsSet() {
		return true
	}

	return false
}

// SetResponseTimeout gets a reference to the given NullableString and assigns it to the ResponseTimeout field.
func (o *EndpointPostBody) SetResponseTimeout(v string) {
	o.ResponseTimeout.Set(&v)
}
// SetResponseTimeoutNil sets the value for ResponseTimeout to be an explicit nil
func (o *EndpointPostBody) SetResponseTimeoutNil() {
	o.ResponseTimeout.Set(nil)
}

// UnsetResponseTimeout ensures that no value is present for ResponseTimeout, not even an explicit nil
func (o *EndpointPostBody) UnsetResponseTimeout() {
	o.ResponseTimeout.Unset()
}

// GetTlsContexts returns the TlsContexts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointPostBody) GetTlsContexts() EndpointPostBodyTlsContexts {
	if o == nil || o.TlsContexts.Get() == nil {
		var ret EndpointPostBodyTlsContexts
		return ret
	}
	return *o.TlsContexts.Get()
}

// GetTlsContextsOk returns a tuple with the TlsContexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointPostBody) GetTlsContextsOk() (*EndpointPostBodyTlsContexts, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TlsContexts.Get(), o.TlsContexts.IsSet()
}

// HasTlsContexts returns a boolean if a field has been set.
func (o *EndpointPostBody) HasTlsContexts() bool {
	if o != nil && o.TlsContexts.IsSet() {
		return true
	}

	return false
}

// SetTlsContexts gets a reference to the given NullableEndpointPostBodyTlsContexts and assigns it to the TlsContexts field.
func (o *EndpointPostBody) SetTlsContexts(v EndpointPostBodyTlsContexts) {
	o.TlsContexts.Set(&v)
}
// SetTlsContextsNil sets the value for TlsContexts to be an explicit nil
func (o *EndpointPostBody) SetTlsContextsNil() {
	o.TlsContexts.Set(nil)
}

// UnsetTlsContexts ensures that no value is present for TlsContexts, not even an explicit nil
func (o *EndpointPostBody) UnsetTlsContexts() {
	o.TlsContexts.Unset()
}

func (o EndpointPostBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeploymentType != nil {
		toSerialize["deploymentType"] = o.DeploymentType
	}
	if o.MuleVersion4OrAbove != nil {
		toSerialize["muleVersion4OrAbove"] = o.MuleVersion4OrAbove
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.IsCloudHub.IsSet() {
		toSerialize["isCloudHub"] = o.IsCloudHub.Get()
	}
	if o.ProxyUri.IsSet() {
		toSerialize["proxyUri"] = o.ProxyUri.Get()
	}
	if o.ProxyRegistrationUri.IsSet() {
		toSerialize["proxyRegistrationUri"] = o.ProxyRegistrationUri.Get()
	}
	if o.ReferencesUserDomain.IsSet() {
		toSerialize["referencesUserDomain"] = o.ReferencesUserDomain.Get()
	}
	if o.ResponseTimeout.IsSet() {
		toSerialize["responseTimeout"] = o.ResponseTimeout.Get()
	}
	if o.TlsContexts.IsSet() {
		toSerialize["tlsContexts"] = o.TlsContexts.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEndpointPostBody struct {
	value *EndpointPostBody
	isSet bool
}

func (v NullableEndpointPostBody) Get() *EndpointPostBody {
	return v.value
}

func (v *NullableEndpointPostBody) Set(val *EndpointPostBody) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointPostBody) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointPostBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointPostBody(val *EndpointPostBody) *NullableEndpointPostBody {
	return &NullableEndpointPostBody{value: val, isSet: true}
}

func (v NullableEndpointPostBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointPostBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


