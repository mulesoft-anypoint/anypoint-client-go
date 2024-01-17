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

// checks if the ApimInstanceCollectionAssetsInnerApisInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApimInstanceCollectionAssetsInnerApisInner{}

// ApimInstanceCollectionAssetsInnerApisInner struct for ApimInstanceCollectionAssetsInnerApisInner
type ApimInstanceCollectionAssetsInnerApisInner struct {
	Audit *Audit `json:"audit,omitempty"`
	MasterOrganizationId *string `json:"masterOrganizationId,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty"`
	Id *int32 `json:"id,omitempty"`
	InstanceLabel *string `json:"instanceLabel,omitempty"`
	GroupId *string `json:"groupId,omitempty"`
	AssetId *string `json:"assetId,omitempty"`
	AssetVersion *string `json:"assetVersion,omitempty"`
	ProductVersion *string `json:"productVersion,omitempty"`
	Description *string `json:"description,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Order *int32 `json:"order,omitempty"`
	ProviderId *string `json:"providerId,omitempty"`
	Deprecated *bool `json:"deprecated,omitempty"`
	LastActiveDate *string `json:"lastActiveDate,omitempty"`
	EndpointUri *string `json:"endpointUri,omitempty"`
	EnvironmentId *string `json:"environmentId,omitempty"`
	IsPublic *bool `json:"isPublic,omitempty"`
	Stage *string `json:"stage,omitempty"`
	Technology *string `json:"technology,omitempty"`
	Status *string `json:"status,omitempty"`
	Deployment NullableDeployment `json:"deployment,omitempty"`
	Routing []Routing `json:"routing,omitempty"`
	Pinned *bool `json:"pinned,omitempty"`
	ActiveContractsCount *int32 `json:"activeContractsCount,omitempty"`
	AutodiscoveryInstanceName *string `json:"autodiscoveryInstanceName,omitempty"`
}

// NewApimInstanceCollectionAssetsInnerApisInner instantiates a new ApimInstanceCollectionAssetsInnerApisInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApimInstanceCollectionAssetsInnerApisInner() *ApimInstanceCollectionAssetsInnerApisInner {
	this := ApimInstanceCollectionAssetsInnerApisInner{}
	return &this
}

// NewApimInstanceCollectionAssetsInnerApisInnerWithDefaults instantiates a new ApimInstanceCollectionAssetsInnerApisInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApimInstanceCollectionAssetsInnerApisInnerWithDefaults() *ApimInstanceCollectionAssetsInnerApisInner {
	this := ApimInstanceCollectionAssetsInnerApisInner{}
	return &this
}

// GetAudit returns the Audit field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetAudit() Audit {
	if o == nil || IsNil(o.Audit) {
		var ret Audit
		return ret
	}
	return *o.Audit
}

// GetAuditOk returns a tuple with the Audit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetAuditOk() (*Audit, bool) {
	if o == nil || IsNil(o.Audit) {
		return nil, false
	}
	return o.Audit, true
}

// HasAudit returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasAudit() bool {
	if o != nil && !IsNil(o.Audit) {
		return true
	}

	return false
}

// SetAudit gets a reference to the given Audit and assigns it to the Audit field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetAudit(v Audit) {
	o.Audit = &v
}

// GetMasterOrganizationId returns the MasterOrganizationId field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetMasterOrganizationId() string {
	if o == nil || IsNil(o.MasterOrganizationId) {
		var ret string
		return ret
	}
	return *o.MasterOrganizationId
}

// GetMasterOrganizationIdOk returns a tuple with the MasterOrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetMasterOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.MasterOrganizationId) {
		return nil, false
	}
	return o.MasterOrganizationId, true
}

// HasMasterOrganizationId returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasMasterOrganizationId() bool {
	if o != nil && !IsNil(o.MasterOrganizationId) {
		return true
	}

	return false
}

// SetMasterOrganizationId gets a reference to the given string and assigns it to the MasterOrganizationId field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetMasterOrganizationId(v string) {
	o.MasterOrganizationId = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetId(v int32) {
	o.Id = &v
}

// GetInstanceLabel returns the InstanceLabel field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetInstanceLabel() string {
	if o == nil || IsNil(o.InstanceLabel) {
		var ret string
		return ret
	}
	return *o.InstanceLabel
}

// GetInstanceLabelOk returns a tuple with the InstanceLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetInstanceLabelOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceLabel) {
		return nil, false
	}
	return o.InstanceLabel, true
}

// HasInstanceLabel returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasInstanceLabel() bool {
	if o != nil && !IsNil(o.InstanceLabel) {
		return true
	}

	return false
}

// SetInstanceLabel gets a reference to the given string and assigns it to the InstanceLabel field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetInstanceLabel(v string) {
	o.InstanceLabel = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetGroupId(v string) {
	o.GroupId = &v
}

// GetAssetId returns the AssetId field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetAssetId() string {
	if o == nil || IsNil(o.AssetId) {
		var ret string
		return ret
	}
	return *o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetAssetIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssetId) {
		return nil, false
	}
	return o.AssetId, true
}

// HasAssetId returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasAssetId() bool {
	if o != nil && !IsNil(o.AssetId) {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given string and assigns it to the AssetId field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetAssetId(v string) {
	o.AssetId = &v
}

// GetAssetVersion returns the AssetVersion field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetAssetVersion() string {
	if o == nil || IsNil(o.AssetVersion) {
		var ret string
		return ret
	}
	return *o.AssetVersion
}

// GetAssetVersionOk returns a tuple with the AssetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetAssetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AssetVersion) {
		return nil, false
	}
	return o.AssetVersion, true
}

// HasAssetVersion returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasAssetVersion() bool {
	if o != nil && !IsNil(o.AssetVersion) {
		return true
	}

	return false
}

// SetAssetVersion gets a reference to the given string and assigns it to the AssetVersion field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetAssetVersion(v string) {
	o.AssetVersion = &v
}

// GetProductVersion returns the ProductVersion field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetProductVersion() string {
	if o == nil || IsNil(o.ProductVersion) {
		var ret string
		return ret
	}
	return *o.ProductVersion
}

// GetProductVersionOk returns a tuple with the ProductVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetProductVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ProductVersion) {
		return nil, false
	}
	return o.ProductVersion, true
}

// HasProductVersion returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasProductVersion() bool {
	if o != nil && !IsNil(o.ProductVersion) {
		return true
	}

	return false
}

// SetProductVersion gets a reference to the given string and assigns it to the ProductVersion field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetProductVersion(v string) {
	o.ProductVersion = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetTags(v []string) {
	o.Tags = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetOrder(v int32) {
	o.Order = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetProviderId() string {
	if o == nil || IsNil(o.ProviderId) {
		var ret string
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetProviderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderId) {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasProviderId() bool {
	if o != nil && !IsNil(o.ProviderId) {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given string and assigns it to the ProviderId field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetProviderId(v string) {
	o.ProviderId = &v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetDeprecated() bool {
	if o == nil || IsNil(o.Deprecated) {
		var ret bool
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetDeprecatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deprecated) {
		return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasDeprecated() bool {
	if o != nil && !IsNil(o.Deprecated) {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetDeprecated(v bool) {
	o.Deprecated = &v
}

// GetLastActiveDate returns the LastActiveDate field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetLastActiveDate() string {
	if o == nil || IsNil(o.LastActiveDate) {
		var ret string
		return ret
	}
	return *o.LastActiveDate
}

// GetLastActiveDateOk returns a tuple with the LastActiveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetLastActiveDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastActiveDate) {
		return nil, false
	}
	return o.LastActiveDate, true
}

// HasLastActiveDate returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasLastActiveDate() bool {
	if o != nil && !IsNil(o.LastActiveDate) {
		return true
	}

	return false
}

// SetLastActiveDate gets a reference to the given string and assigns it to the LastActiveDate field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetLastActiveDate(v string) {
	o.LastActiveDate = &v
}

// GetEndpointUri returns the EndpointUri field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetEndpointUri() string {
	if o == nil || IsNil(o.EndpointUri) {
		var ret string
		return ret
	}
	return *o.EndpointUri
}

// GetEndpointUriOk returns a tuple with the EndpointUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetEndpointUriOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointUri) {
		return nil, false
	}
	return o.EndpointUri, true
}

// HasEndpointUri returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasEndpointUri() bool {
	if o != nil && !IsNil(o.EndpointUri) {
		return true
	}

	return false
}

// SetEndpointUri gets a reference to the given string and assigns it to the EndpointUri field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetEndpointUri(v string) {
	o.EndpointUri = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId) {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentId) {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasEnvironmentId() bool {
	if o != nil && !IsNil(o.EnvironmentId) {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetStage returns the Stage field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetStage() string {
	if o == nil || IsNil(o.Stage) {
		var ret string
		return ret
	}
	return *o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetStageOk() (*string, bool) {
	if o == nil || IsNil(o.Stage) {
		return nil, false
	}
	return o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasStage() bool {
	if o != nil && !IsNil(o.Stage) {
		return true
	}

	return false
}

// SetStage gets a reference to the given string and assigns it to the Stage field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetStage(v string) {
	o.Stage = &v
}

// GetTechnology returns the Technology field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetTechnology() string {
	if o == nil || IsNil(o.Technology) {
		var ret string
		return ret
	}
	return *o.Technology
}

// GetTechnologyOk returns a tuple with the Technology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetTechnologyOk() (*string, bool) {
	if o == nil || IsNil(o.Technology) {
		return nil, false
	}
	return o.Technology, true
}

// HasTechnology returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasTechnology() bool {
	if o != nil && !IsNil(o.Technology) {
		return true
	}

	return false
}

// SetTechnology gets a reference to the given string and assigns it to the Technology field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetTechnology(v string) {
	o.Technology = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetStatus(v string) {
	o.Status = &v
}

// GetDeployment returns the Deployment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetDeployment() Deployment {
	if o == nil || IsNil(o.Deployment.Get()) {
		var ret Deployment
		return ret
	}
	return *o.Deployment.Get()
}

// GetDeploymentOk returns a tuple with the Deployment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetDeploymentOk() (*Deployment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deployment.Get(), o.Deployment.IsSet()
}

// HasDeployment returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasDeployment() bool {
	if o != nil && o.Deployment.IsSet() {
		return true
	}

	return false
}

// SetDeployment gets a reference to the given NullableDeployment and assigns it to the Deployment field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetDeployment(v Deployment) {
	o.Deployment.Set(&v)
}
// SetDeploymentNil sets the value for Deployment to be an explicit nil
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetDeploymentNil() {
	o.Deployment.Set(nil)
}

// UnsetDeployment ensures that no value is present for Deployment, not even an explicit nil
func (o *ApimInstanceCollectionAssetsInnerApisInner) UnsetDeployment() {
	o.Deployment.Unset()
}

// GetRouting returns the Routing field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetRouting() []Routing {
	if o == nil || IsNil(o.Routing) {
		var ret []Routing
		return ret
	}
	return o.Routing
}

// GetRoutingOk returns a tuple with the Routing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetRoutingOk() ([]Routing, bool) {
	if o == nil || IsNil(o.Routing) {
		return nil, false
	}
	return o.Routing, true
}

// HasRouting returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasRouting() bool {
	if o != nil && !IsNil(o.Routing) {
		return true
	}

	return false
}

// SetRouting gets a reference to the given []Routing and assigns it to the Routing field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetRouting(v []Routing) {
	o.Routing = v
}

// GetPinned returns the Pinned field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetPinned() bool {
	if o == nil || IsNil(o.Pinned) {
		var ret bool
		return ret
	}
	return *o.Pinned
}

// GetPinnedOk returns a tuple with the Pinned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetPinnedOk() (*bool, bool) {
	if o == nil || IsNil(o.Pinned) {
		return nil, false
	}
	return o.Pinned, true
}

// HasPinned returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasPinned() bool {
	if o != nil && !IsNil(o.Pinned) {
		return true
	}

	return false
}

// SetPinned gets a reference to the given bool and assigns it to the Pinned field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetPinned(v bool) {
	o.Pinned = &v
}

// GetActiveContractsCount returns the ActiveContractsCount field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetActiveContractsCount() int32 {
	if o == nil || IsNil(o.ActiveContractsCount) {
		var ret int32
		return ret
	}
	return *o.ActiveContractsCount
}

// GetActiveContractsCountOk returns a tuple with the ActiveContractsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetActiveContractsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ActiveContractsCount) {
		return nil, false
	}
	return o.ActiveContractsCount, true
}

// HasActiveContractsCount returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasActiveContractsCount() bool {
	if o != nil && !IsNil(o.ActiveContractsCount) {
		return true
	}

	return false
}

// SetActiveContractsCount gets a reference to the given int32 and assigns it to the ActiveContractsCount field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetActiveContractsCount(v int32) {
	o.ActiveContractsCount = &v
}

// GetAutodiscoveryInstanceName returns the AutodiscoveryInstanceName field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetAutodiscoveryInstanceName() string {
	if o == nil || IsNil(o.AutodiscoveryInstanceName) {
		var ret string
		return ret
	}
	return *o.AutodiscoveryInstanceName
}

// GetAutodiscoveryInstanceNameOk returns a tuple with the AutodiscoveryInstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) GetAutodiscoveryInstanceNameOk() (*string, bool) {
	if o == nil || IsNil(o.AutodiscoveryInstanceName) {
		return nil, false
	}
	return o.AutodiscoveryInstanceName, true
}

// HasAutodiscoveryInstanceName returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssetsInnerApisInner) HasAutodiscoveryInstanceName() bool {
	if o != nil && !IsNil(o.AutodiscoveryInstanceName) {
		return true
	}

	return false
}

// SetAutodiscoveryInstanceName gets a reference to the given string and assigns it to the AutodiscoveryInstanceName field.
func (o *ApimInstanceCollectionAssetsInnerApisInner) SetAutodiscoveryInstanceName(v string) {
	o.AutodiscoveryInstanceName = &v
}

func (o ApimInstanceCollectionAssetsInnerApisInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApimInstanceCollectionAssetsInnerApisInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !IsNil(o.InstanceLabel) {
		toSerialize["instanceLabel"] = o.InstanceLabel
	}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.AssetId) {
		toSerialize["assetId"] = o.AssetId
	}
	if !IsNil(o.AssetVersion) {
		toSerialize["assetVersion"] = o.AssetVersion
	}
	if !IsNil(o.ProductVersion) {
		toSerialize["productVersion"] = o.ProductVersion
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.ProviderId) {
		toSerialize["providerId"] = o.ProviderId
	}
	if !IsNil(o.Deprecated) {
		toSerialize["deprecated"] = o.Deprecated
	}
	if !IsNil(o.LastActiveDate) {
		toSerialize["lastActiveDate"] = o.LastActiveDate
	}
	if !IsNil(o.EndpointUri) {
		toSerialize["endpointUri"] = o.EndpointUri
	}
	if !IsNil(o.EnvironmentId) {
		toSerialize["environmentId"] = o.EnvironmentId
	}
	if !IsNil(o.IsPublic) {
		toSerialize["isPublic"] = o.IsPublic
	}
	if !IsNil(o.Stage) {
		toSerialize["stage"] = o.Stage
	}
	if !IsNil(o.Technology) {
		toSerialize["technology"] = o.Technology
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Deployment.IsSet() {
		toSerialize["deployment"] = o.Deployment.Get()
	}
	if !IsNil(o.Routing) {
		toSerialize["routing"] = o.Routing
	}
	if !IsNil(o.Pinned) {
		toSerialize["pinned"] = o.Pinned
	}
	if !IsNil(o.ActiveContractsCount) {
		toSerialize["activeContractsCount"] = o.ActiveContractsCount
	}
	if !IsNil(o.AutodiscoveryInstanceName) {
		toSerialize["autodiscoveryInstanceName"] = o.AutodiscoveryInstanceName
	}
	return toSerialize, nil
}

type NullableApimInstanceCollectionAssetsInnerApisInner struct {
	value *ApimInstanceCollectionAssetsInnerApisInner
	isSet bool
}

func (v NullableApimInstanceCollectionAssetsInnerApisInner) Get() *ApimInstanceCollectionAssetsInnerApisInner {
	return v.value
}

func (v *NullableApimInstanceCollectionAssetsInnerApisInner) Set(val *ApimInstanceCollectionAssetsInnerApisInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApimInstanceCollectionAssetsInnerApisInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApimInstanceCollectionAssetsInnerApisInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApimInstanceCollectionAssetsInnerApisInner(val *ApimInstanceCollectionAssetsInnerApisInner) *NullableApimInstanceCollectionAssetsInnerApisInner {
	return &NullableApimInstanceCollectionAssetsInnerApisInner{value: val, isSet: true}
}

func (v NullableApimInstanceCollectionAssetsInnerApisInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApimInstanceCollectionAssetsInnerApisInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


