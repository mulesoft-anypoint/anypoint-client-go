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

// checks if the ApimInstancePostResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApimInstancePostResponse{}

// ApimInstancePostResponse struct for ApimInstancePostResponse
type ApimInstancePostResponse struct {
	EnvironmentId *string `json:"environmentId,omitempty"`
	InstanceLabel *string `json:"instanceLabel,omitempty"`
	ProviderId NullableString `json:"providerId,omitempty"`
	Technology *string `json:"technology,omitempty"`
	AssetVersion *string `json:"assetVersion,omitempty"`
	ProductVersion *string `json:"productVersion,omitempty"`
	Order *int32 `json:"order,omitempty"`
	Stage *string `json:"stage,omitempty"`
	Audit *Audit `json:"audit,omitempty"`
	MasterOrganizationId *string `json:"masterOrganizationId,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty"`
	Id *int32 `json:"id,omitempty"`
	GroupId *string `json:"groupId,omitempty"`
	AssetId *string `json:"assetId,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Endpoint *Endpoint `json:"endpoint,omitempty"`
	AutodiscoveryInstanceName *string `json:"autodiscoveryInstanceName,omitempty"`
}

// NewApimInstancePostResponse instantiates a new ApimInstancePostResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApimInstancePostResponse() *ApimInstancePostResponse {
	this := ApimInstancePostResponse{}
	return &this
}

// NewApimInstancePostResponseWithDefaults instantiates a new ApimInstancePostResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApimInstancePostResponseWithDefaults() *ApimInstancePostResponse {
	this := ApimInstancePostResponse{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *ApimInstancePostResponse) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId) {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstancePostResponse) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentId) {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *ApimInstancePostResponse) HasEnvironmentId() bool {
	if o != nil && !IsNil(o.EnvironmentId) {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *ApimInstancePostResponse) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetInstanceLabel returns the InstanceLabel field value if set, zero value otherwise.
func (o *ApimInstancePostResponse) GetInstanceLabel() string {
	if o == nil || IsNil(o.InstanceLabel) {
		var ret string
		return ret
	}
	return *o.InstanceLabel
}

// GetInstanceLabelOk returns a tuple with the InstanceLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstancePostResponse) GetInstanceLabelOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceLabel) {
		return nil, false
	}
	return o.InstanceLabel, true
}

// HasInstanceLabel returns a boolean if a field has been set.
func (o *ApimInstancePostResponse) HasInstanceLabel() bool {
	if o != nil && !IsNil(o.InstanceLabel) {
		return true
	}

	return false
}

// SetInstanceLabel gets a reference to the given string and assigns it to the InstanceLabel field.
func (o *ApimInstancePostResponse) SetInstanceLabel(v string) {
	o.InstanceLabel = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApimInstancePostResponse) GetProviderId() string {
	if o == nil || IsNil(o.ProviderId.Get()) {
		var ret string
		return ret
	}
	return *o.ProviderId.Get()
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApimInstancePostResponse) GetProviderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderId.Get(), o.ProviderId.IsSet()
}

// HasProviderId returns a boolean if a field has been set.
func (o *ApimInstancePostResponse) HasProviderId() bool {
	if o != nil && o.ProviderId.IsSet() {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given NullableString and assigns it to the ProviderId field.
func (o *ApimInstancePostResponse) SetProviderId(v string) {
	o.ProviderId.Set(&v)
}
// SetProviderIdNil sets the value for ProviderId to be an explicit nil
func (o *ApimInstancePostResponse) SetProviderIdNil() {
	o.ProviderId.Set(nil)
}

// UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
func (o *ApimInstancePostResponse) UnsetProviderId() {
	o.ProviderId.Unset()
}

// GetTechnology returns the Technology field value if set, zero value otherwise.
func (o *ApimInstancePostResponse) GetTechnology() string {
	if o == nil || IsNil(o.Technology) {
		var ret string
		return ret
	}
	return *o.Technology
}

// GetTechnologyOk returns a tuple with the Technology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstancePostResponse) GetTechnologyOk() (*string, bool) {
	if o == nil || IsNil(o.Technology) {
		return nil, false
	}
	return o.Technology, true
}

// HasTechnology returns a boolean if a field has been set.
func (o *ApimInstancePostResponse) HasTechnology() bool {
	if o != nil && !IsNil(o.Technology) {
		return true
	}

	return false
}

// SetTechnology gets a reference to the given string and assigns it to the Technology field.
func (o *ApimInstancePostResponse) SetTechnology(v string) {
	o.Technology = &v
}

// GetAssetVersion returns the AssetVersion field value if set, zero value otherwise.
func (o *ApimInstancePostResponse) GetAssetVersion() string {
	if o == nil || IsNil(o.AssetVersion) {
		var ret string
		return ret
	}
	return *o.AssetVersion
}

// GetAssetVersionOk returns a tuple with the AssetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstancePostResponse) GetAssetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AssetVersion) {
		return nil, false
	}
	return o.AssetVersion, true
}

// HasAssetVersion returns a boolean if a field has been set.
func (o *ApimInstancePostResponse) HasAssetVersion() bool {
	if o != nil && !IsNil(o.AssetVersion) {
		return true
	}

	return false
}

// SetAssetVersion gets a reference to the given string and assigns it to the AssetVersion field.
func (o *ApimInstancePostResponse) SetAssetVersion(v string) {
	o.AssetVersion = &v
}

// GetProductVersion returns the ProductVersion field value if set, zero value otherwise.
func (o *ApimInstancePostResponse) GetProductVersion() string {
	if o == nil || IsNil(o.ProductVersion) {
		var ret string
		return ret
	}
	return *o.ProductVersion
}

// GetProductVersionOk returns a tuple with the ProductVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstancePostResponse) GetProductVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ProductVersion) {
		return nil, false
	}
	return o.ProductVersion, true
}

// HasProductVersion returns a boolean if a field has been set.
func (o *ApimInstancePostResponse) HasProductVersion() bool {
	if o != nil && !IsNil(o.ProductVersion) {
		return true
	}

	return false
}

// SetProductVersion gets a reference to the given string and assigns it to the ProductVersion field.
func (o *ApimInstancePostResponse) SetProductVersion(v string) {
	o.ProductVersion = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ApimInstancePostResponse) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstancePostResponse) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ApimInstancePostResponse) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *ApimInstancePostResponse) SetOrder(v int32) {
	o.Order = &v
}

// GetStage returns the Stage field value if set, zero value otherwise.
func (o *ApimInstancePostResponse) GetStage() string {
	if o == nil || IsNil(o.Stage) {
		var ret string
		return ret
	}
	return *o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstancePostResponse) GetStageOk() (*string, bool) {
	if o == nil || IsNil(o.Stage) {
		return nil, false
	}
	return o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *ApimInstancePostResponse) HasStage() bool {
	if o != nil && !IsNil(o.Stage) {
		return true
	}

	return false
}

// SetStage gets a reference to the given string and assigns it to the Stage field.
func (o *ApimInstancePostResponse) SetStage(v string) {
	o.Stage = &v
}

// GetAudit returns the Audit field value if set, zero value otherwise.
func (o *ApimInstancePostResponse) GetAudit() Audit {
	if o == nil || IsNil(o.Audit) {
		var ret Audit
		return ret
	}
	return *o.Audit
}

// GetAuditOk returns a tuple with the Audit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstancePostResponse) GetAuditOk() (*Audit, bool) {
	if o == nil || IsNil(o.Audit) {
		return nil, false
	}
	return o.Audit, true
}

// HasAudit returns a boolean if a field has been set.
func (o *ApimInstancePostResponse) HasAudit() bool {
	if o != nil && !IsNil(o.Audit) {
		return true
	}

	return false
}

// SetAudit gets a reference to the given Audit and assigns it to the Audit field.
func (o *ApimInstancePostResponse) SetAudit(v Audit) {
	o.Audit = &v
}

// GetMasterOrganizationId returns the MasterOrganizationId field value if set, zero value otherwise.
func (o *ApimInstancePostResponse) GetMasterOrganizationId() string {
	if o == nil || IsNil(o.MasterOrganizationId) {
		var ret string
		return ret
	}
	return *o.MasterOrganizationId
}

// GetMasterOrganizationIdOk returns a tuple with the MasterOrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstancePostResponse) GetMasterOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.MasterOrganizationId) {
		return nil, false
	}
	return o.MasterOrganizationId, true
}

// HasMasterOrganizationId returns a boolean if a field has been set.
func (o *ApimInstancePostResponse) HasMasterOrganizationId() bool {
	if o != nil && !IsNil(o.MasterOrganizationId) {
		return true
	}

	return false
}

// SetMasterOrganizationId gets a reference to the given string and assigns it to the MasterOrganizationId field.
func (o *ApimInstancePostResponse) SetMasterOrganizationId(v string) {
	o.MasterOrganizationId = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ApimInstancePostResponse) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstancePostResponse) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ApimInstancePostResponse) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *ApimInstancePostResponse) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApimInstancePostResponse) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstancePostResponse) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApimInstancePostResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ApimInstancePostResponse) SetId(v int32) {
	o.Id = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ApimInstancePostResponse) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstancePostResponse) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ApimInstancePostResponse) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ApimInstancePostResponse) SetGroupId(v string) {
	o.GroupId = &v
}

// GetAssetId returns the AssetId field value if set, zero value otherwise.
func (o *ApimInstancePostResponse) GetAssetId() string {
	if o == nil || IsNil(o.AssetId) {
		var ret string
		return ret
	}
	return *o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstancePostResponse) GetAssetIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssetId) {
		return nil, false
	}
	return o.AssetId, true
}

// HasAssetId returns a boolean if a field has been set.
func (o *ApimInstancePostResponse) HasAssetId() bool {
	if o != nil && !IsNil(o.AssetId) {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given string and assigns it to the AssetId field.
func (o *ApimInstancePostResponse) SetAssetId(v string) {
	o.AssetId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ApimInstancePostResponse) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstancePostResponse) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ApimInstancePostResponse) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ApimInstancePostResponse) SetTags(v []string) {
	o.Tags = v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *ApimInstancePostResponse) GetEndpoint() Endpoint {
	if o == nil || IsNil(o.Endpoint) {
		var ret Endpoint
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstancePostResponse) GetEndpointOk() (*Endpoint, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *ApimInstancePostResponse) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given Endpoint and assigns it to the Endpoint field.
func (o *ApimInstancePostResponse) SetEndpoint(v Endpoint) {
	o.Endpoint = &v
}

// GetAutodiscoveryInstanceName returns the AutodiscoveryInstanceName field value if set, zero value otherwise.
func (o *ApimInstancePostResponse) GetAutodiscoveryInstanceName() string {
	if o == nil || IsNil(o.AutodiscoveryInstanceName) {
		var ret string
		return ret
	}
	return *o.AutodiscoveryInstanceName
}

// GetAutodiscoveryInstanceNameOk returns a tuple with the AutodiscoveryInstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstancePostResponse) GetAutodiscoveryInstanceNameOk() (*string, bool) {
	if o == nil || IsNil(o.AutodiscoveryInstanceName) {
		return nil, false
	}
	return o.AutodiscoveryInstanceName, true
}

// HasAutodiscoveryInstanceName returns a boolean if a field has been set.
func (o *ApimInstancePostResponse) HasAutodiscoveryInstanceName() bool {
	if o != nil && !IsNil(o.AutodiscoveryInstanceName) {
		return true
	}

	return false
}

// SetAutodiscoveryInstanceName gets a reference to the given string and assigns it to the AutodiscoveryInstanceName field.
func (o *ApimInstancePostResponse) SetAutodiscoveryInstanceName(v string) {
	o.AutodiscoveryInstanceName = &v
}

func (o ApimInstancePostResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApimInstancePostResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnvironmentId) {
		toSerialize["environmentId"] = o.EnvironmentId
	}
	if !IsNil(o.InstanceLabel) {
		toSerialize["instanceLabel"] = o.InstanceLabel
	}
	if o.ProviderId.IsSet() {
		toSerialize["providerId"] = o.ProviderId.Get()
	}
	if !IsNil(o.Technology) {
		toSerialize["technology"] = o.Technology
	}
	if !IsNil(o.AssetVersion) {
		toSerialize["assetVersion"] = o.AssetVersion
	}
	if !IsNil(o.ProductVersion) {
		toSerialize["productVersion"] = o.ProductVersion
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Stage) {
		toSerialize["stage"] = o.Stage
	}
	if !IsNil(o.Audit) {
		toSerialize["audit"] = o.Audit
	}
	if !IsNil(o.MasterOrganizationId) {
		toSerialize["masterOrganizationId"] = o.MasterOrganizationId
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.AssetId) {
		toSerialize["assetId"] = o.AssetId
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !IsNil(o.AutodiscoveryInstanceName) {
		toSerialize["autodiscoveryInstanceName"] = o.AutodiscoveryInstanceName
	}
	return toSerialize, nil
}

type NullableApimInstancePostResponse struct {
	value *ApimInstancePostResponse
	isSet bool
}

func (v NullableApimInstancePostResponse) Get() *ApimInstancePostResponse {
	return v.value
}

func (v *NullableApimInstancePostResponse) Set(val *ApimInstancePostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApimInstancePostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApimInstancePostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApimInstancePostResponse(val *ApimInstancePostResponse) *NullableApimInstancePostResponse {
	return &NullableApimInstancePostResponse{value: val, isSet: true}
}

func (v NullableApimInstancePostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApimInstancePostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


