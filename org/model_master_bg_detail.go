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

// checks if the MasterBGDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MasterBGDetail{}

// MasterBGDetail struct for MasterBGDetail
type MasterBGDetail struct {
	// The client id of the organization
	ClientId *string `json:"clientId,omitempty"`
	// The creation date of the organization
	CreatedAt *string `json:"createdAt,omitempty"`
	// The domain of the organization
	Domain *string `json:"domain,omitempty"`
	Entitlements *Entitlements `json:"entitlements,omitempty"`
	// List of environments.
	Environments []Environment `json:"environments,omitempty"`
	// The id of the organization
	Id *string `json:"id,omitempty"`
	// The id of the idprovider of the organization.
	IdproviderId *string `json:"idprovider_id,omitempty"`
	// Whether the organization is exempt from automatic admin promotion.
	IsAutomaticAdminPromotionExempt *bool `json:"isAutomaticAdminPromotionExempt,omitempty"`
	// Whether the organization is federated.
	IsFederated *bool `json:"isFederated,omitempty"`
	// Whether the organization is a root organization.
	IsRoot *bool `json:"isRoot,omitempty"`
	// Whether the organization is a master organization.
	IsMaster *bool `json:"isMaster,omitempty"`
	// Whether multi-factor authentication is required for the organization.
	MfaRequired *string `json:"mfaRequired,omitempty"`
	// The name of the organization.
	Name *string `json:"name,omitempty"`
	// The id of the owner of the organization.
	OwnerId *string `json:"ownerId,omitempty"`
	// The ids of the parent organizations.
	ParentOrganizationIds []string `json:"parentOrganizationIds,omitempty"`
	// The properties of the organization.
	Properties map[string]interface{} `json:"properties,omitempty"`
	// The ids of the sub organizations.
	SubOrganizationIds []string `json:"subOrganizationIds,omitempty"`
	// The ids of the tenant organizations.
	TenantOrganizationIds []string `json:"tenantOrganizationIds,omitempty"`
	// The last update date of the organization.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	Owner *User `json:"owner,omitempty"`
	// The session timeout in minutes.
	SessionTimeout *int32 `json:"sessionTimeout,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
	// The gdot id of the organization
	GdotId *string `json:"gdotId,omitempty"`
	// The deleted date of the organization
	DeletedAt NullableString `json:"deletedAt,omitempty"`
	// The type of the organization
	OrgType *string `json:"orgType,omitempty"`
}

// NewMasterBGDetail instantiates a new MasterBGDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMasterBGDetail() *MasterBGDetail {
	this := MasterBGDetail{}
	var clientId string = ""
	this.ClientId = &clientId
	var createdAt string = ""
	this.CreatedAt = &createdAt
	var domain string = ""
	this.Domain = &domain
	var id string = ""
	this.Id = &id
	var idproviderId string = ""
	this.IdproviderId = &idproviderId
	var isAutomaticAdminPromotionExempt bool = false
	this.IsAutomaticAdminPromotionExempt = &isAutomaticAdminPromotionExempt
	var isFederated bool = false
	this.IsFederated = &isFederated
	var isRoot bool = false
	this.IsRoot = &isRoot
	var isMaster bool = false
	this.IsMaster = &isMaster
	var mfaRequired string = ""
	this.MfaRequired = &mfaRequired
	var name string = ""
	this.Name = &name
	var ownerId string = ""
	this.OwnerId = &ownerId
	var updatedAt string = ""
	this.UpdatedAt = &updatedAt
	var sessionTimeout int32 = 0
	this.SessionTimeout = &sessionTimeout
	var gdotId string = ""
	this.GdotId = &gdotId
	var orgType string = ""
	this.OrgType = &orgType
	return &this
}

// NewMasterBGDetailWithDefaults instantiates a new MasterBGDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMasterBGDetailWithDefaults() *MasterBGDetail {
	this := MasterBGDetail{}
	var clientId string = ""
	this.ClientId = &clientId
	var createdAt string = ""
	this.CreatedAt = &createdAt
	var domain string = ""
	this.Domain = &domain
	var id string = ""
	this.Id = &id
	var idproviderId string = ""
	this.IdproviderId = &idproviderId
	var isAutomaticAdminPromotionExempt bool = false
	this.IsAutomaticAdminPromotionExempt = &isAutomaticAdminPromotionExempt
	var isFederated bool = false
	this.IsFederated = &isFederated
	var isRoot bool = false
	this.IsRoot = &isRoot
	var isMaster bool = false
	this.IsMaster = &isMaster
	var mfaRequired string = ""
	this.MfaRequired = &mfaRequired
	var name string = ""
	this.Name = &name
	var ownerId string = ""
	this.OwnerId = &ownerId
	var updatedAt string = ""
	this.UpdatedAt = &updatedAt
	var sessionTimeout int32 = 0
	this.SessionTimeout = &sessionTimeout
	var gdotId string = ""
	this.GdotId = &gdotId
	var orgType string = ""
	this.OrgType = &orgType
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *MasterBGDetail) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *MasterBGDetail) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *MasterBGDetail) SetClientId(v string) {
	o.ClientId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MasterBGDetail) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MasterBGDetail) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *MasterBGDetail) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *MasterBGDetail) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *MasterBGDetail) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *MasterBGDetail) SetDomain(v string) {
	o.Domain = &v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *MasterBGDetail) GetEntitlements() Entitlements {
	if o == nil || IsNil(o.Entitlements) {
		var ret Entitlements
		return ret
	}
	return *o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetEntitlementsOk() (*Entitlements, bool) {
	if o == nil || IsNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *MasterBGDetail) HasEntitlements() bool {
	if o != nil && !IsNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given Entitlements and assigns it to the Entitlements field.
func (o *MasterBGDetail) SetEntitlements(v Entitlements) {
	o.Entitlements = &v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *MasterBGDetail) GetEnvironments() []Environment {
	if o == nil || IsNil(o.Environments) {
		var ret []Environment
		return ret
	}
	return o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetEnvironmentsOk() ([]Environment, bool) {
	if o == nil || IsNil(o.Environments) {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *MasterBGDetail) HasEnvironments() bool {
	if o != nil && !IsNil(o.Environments) {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given []Environment and assigns it to the Environments field.
func (o *MasterBGDetail) SetEnvironments(v []Environment) {
	o.Environments = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MasterBGDetail) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MasterBGDetail) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MasterBGDetail) SetId(v string) {
	o.Id = &v
}

// GetIdproviderId returns the IdproviderId field value if set, zero value otherwise.
func (o *MasterBGDetail) GetIdproviderId() string {
	if o == nil || IsNil(o.IdproviderId) {
		var ret string
		return ret
	}
	return *o.IdproviderId
}

// GetIdproviderIdOk returns a tuple with the IdproviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetIdproviderIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdproviderId) {
		return nil, false
	}
	return o.IdproviderId, true
}

// HasIdproviderId returns a boolean if a field has been set.
func (o *MasterBGDetail) HasIdproviderId() bool {
	if o != nil && !IsNil(o.IdproviderId) {
		return true
	}

	return false
}

// SetIdproviderId gets a reference to the given string and assigns it to the IdproviderId field.
func (o *MasterBGDetail) SetIdproviderId(v string) {
	o.IdproviderId = &v
}

// GetIsAutomaticAdminPromotionExempt returns the IsAutomaticAdminPromotionExempt field value if set, zero value otherwise.
func (o *MasterBGDetail) GetIsAutomaticAdminPromotionExempt() bool {
	if o == nil || IsNil(o.IsAutomaticAdminPromotionExempt) {
		var ret bool
		return ret
	}
	return *o.IsAutomaticAdminPromotionExempt
}

// GetIsAutomaticAdminPromotionExemptOk returns a tuple with the IsAutomaticAdminPromotionExempt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetIsAutomaticAdminPromotionExemptOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAutomaticAdminPromotionExempt) {
		return nil, false
	}
	return o.IsAutomaticAdminPromotionExempt, true
}

// HasIsAutomaticAdminPromotionExempt returns a boolean if a field has been set.
func (o *MasterBGDetail) HasIsAutomaticAdminPromotionExempt() bool {
	if o != nil && !IsNil(o.IsAutomaticAdminPromotionExempt) {
		return true
	}

	return false
}

// SetIsAutomaticAdminPromotionExempt gets a reference to the given bool and assigns it to the IsAutomaticAdminPromotionExempt field.
func (o *MasterBGDetail) SetIsAutomaticAdminPromotionExempt(v bool) {
	o.IsAutomaticAdminPromotionExempt = &v
}

// GetIsFederated returns the IsFederated field value if set, zero value otherwise.
func (o *MasterBGDetail) GetIsFederated() bool {
	if o == nil || IsNil(o.IsFederated) {
		var ret bool
		return ret
	}
	return *o.IsFederated
}

// GetIsFederatedOk returns a tuple with the IsFederated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetIsFederatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFederated) {
		return nil, false
	}
	return o.IsFederated, true
}

// HasIsFederated returns a boolean if a field has been set.
func (o *MasterBGDetail) HasIsFederated() bool {
	if o != nil && !IsNil(o.IsFederated) {
		return true
	}

	return false
}

// SetIsFederated gets a reference to the given bool and assigns it to the IsFederated field.
func (o *MasterBGDetail) SetIsFederated(v bool) {
	o.IsFederated = &v
}

// GetIsRoot returns the IsRoot field value if set, zero value otherwise.
func (o *MasterBGDetail) GetIsRoot() bool {
	if o == nil || IsNil(o.IsRoot) {
		var ret bool
		return ret
	}
	return *o.IsRoot
}

// GetIsRootOk returns a tuple with the IsRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetIsRootOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRoot) {
		return nil, false
	}
	return o.IsRoot, true
}

// HasIsRoot returns a boolean if a field has been set.
func (o *MasterBGDetail) HasIsRoot() bool {
	if o != nil && !IsNil(o.IsRoot) {
		return true
	}

	return false
}

// SetIsRoot gets a reference to the given bool and assigns it to the IsRoot field.
func (o *MasterBGDetail) SetIsRoot(v bool) {
	o.IsRoot = &v
}

// GetIsMaster returns the IsMaster field value if set, zero value otherwise.
func (o *MasterBGDetail) GetIsMaster() bool {
	if o == nil || IsNil(o.IsMaster) {
		var ret bool
		return ret
	}
	return *o.IsMaster
}

// GetIsMasterOk returns a tuple with the IsMaster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetIsMasterOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMaster) {
		return nil, false
	}
	return o.IsMaster, true
}

// HasIsMaster returns a boolean if a field has been set.
func (o *MasterBGDetail) HasIsMaster() bool {
	if o != nil && !IsNil(o.IsMaster) {
		return true
	}

	return false
}

// SetIsMaster gets a reference to the given bool and assigns it to the IsMaster field.
func (o *MasterBGDetail) SetIsMaster(v bool) {
	o.IsMaster = &v
}

// GetMfaRequired returns the MfaRequired field value if set, zero value otherwise.
func (o *MasterBGDetail) GetMfaRequired() string {
	if o == nil || IsNil(o.MfaRequired) {
		var ret string
		return ret
	}
	return *o.MfaRequired
}

// GetMfaRequiredOk returns a tuple with the MfaRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetMfaRequiredOk() (*string, bool) {
	if o == nil || IsNil(o.MfaRequired) {
		return nil, false
	}
	return o.MfaRequired, true
}

// HasMfaRequired returns a boolean if a field has been set.
func (o *MasterBGDetail) HasMfaRequired() bool {
	if o != nil && !IsNil(o.MfaRequired) {
		return true
	}

	return false
}

// SetMfaRequired gets a reference to the given string and assigns it to the MfaRequired field.
func (o *MasterBGDetail) SetMfaRequired(v string) {
	o.MfaRequired = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MasterBGDetail) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MasterBGDetail) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MasterBGDetail) SetName(v string) {
	o.Name = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *MasterBGDetail) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *MasterBGDetail) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *MasterBGDetail) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetParentOrganizationIds returns the ParentOrganizationIds field value if set, zero value otherwise.
func (o *MasterBGDetail) GetParentOrganizationIds() []string {
	if o == nil || IsNil(o.ParentOrganizationIds) {
		var ret []string
		return ret
	}
	return o.ParentOrganizationIds
}

// GetParentOrganizationIdsOk returns a tuple with the ParentOrganizationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetParentOrganizationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ParentOrganizationIds) {
		return nil, false
	}
	return o.ParentOrganizationIds, true
}

// HasParentOrganizationIds returns a boolean if a field has been set.
func (o *MasterBGDetail) HasParentOrganizationIds() bool {
	if o != nil && !IsNil(o.ParentOrganizationIds) {
		return true
	}

	return false
}

// SetParentOrganizationIds gets a reference to the given []string and assigns it to the ParentOrganizationIds field.
func (o *MasterBGDetail) SetParentOrganizationIds(v []string) {
	o.ParentOrganizationIds = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *MasterBGDetail) GetProperties() map[string]interface{} {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Properties) {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *MasterBGDetail) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *MasterBGDetail) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

// GetSubOrganizationIds returns the SubOrganizationIds field value if set, zero value otherwise.
func (o *MasterBGDetail) GetSubOrganizationIds() []string {
	if o == nil || IsNil(o.SubOrganizationIds) {
		var ret []string
		return ret
	}
	return o.SubOrganizationIds
}

// GetSubOrganizationIdsOk returns a tuple with the SubOrganizationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetSubOrganizationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SubOrganizationIds) {
		return nil, false
	}
	return o.SubOrganizationIds, true
}

// HasSubOrganizationIds returns a boolean if a field has been set.
func (o *MasterBGDetail) HasSubOrganizationIds() bool {
	if o != nil && !IsNil(o.SubOrganizationIds) {
		return true
	}

	return false
}

// SetSubOrganizationIds gets a reference to the given []string and assigns it to the SubOrganizationIds field.
func (o *MasterBGDetail) SetSubOrganizationIds(v []string) {
	o.SubOrganizationIds = v
}

// GetTenantOrganizationIds returns the TenantOrganizationIds field value if set, zero value otherwise.
func (o *MasterBGDetail) GetTenantOrganizationIds() []string {
	if o == nil || IsNil(o.TenantOrganizationIds) {
		var ret []string
		return ret
	}
	return o.TenantOrganizationIds
}

// GetTenantOrganizationIdsOk returns a tuple with the TenantOrganizationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetTenantOrganizationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TenantOrganizationIds) {
		return nil, false
	}
	return o.TenantOrganizationIds, true
}

// HasTenantOrganizationIds returns a boolean if a field has been set.
func (o *MasterBGDetail) HasTenantOrganizationIds() bool {
	if o != nil && !IsNil(o.TenantOrganizationIds) {
		return true
	}

	return false
}

// SetTenantOrganizationIds gets a reference to the given []string and assigns it to the TenantOrganizationIds field.
func (o *MasterBGDetail) SetTenantOrganizationIds(v []string) {
	o.TenantOrganizationIds = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *MasterBGDetail) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *MasterBGDetail) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *MasterBGDetail) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *MasterBGDetail) GetOwner() User {
	if o == nil || IsNil(o.Owner) {
		var ret User
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetOwnerOk() (*User, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *MasterBGDetail) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given User and assigns it to the Owner field.
func (o *MasterBGDetail) SetOwner(v User) {
	o.Owner = &v
}

// GetSessionTimeout returns the SessionTimeout field value if set, zero value otherwise.
func (o *MasterBGDetail) GetSessionTimeout() int32 {
	if o == nil || IsNil(o.SessionTimeout) {
		var ret int32
		return ret
	}
	return *o.SessionTimeout
}

// GetSessionTimeoutOk returns a tuple with the SessionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetSessionTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.SessionTimeout) {
		return nil, false
	}
	return o.SessionTimeout, true
}

// HasSessionTimeout returns a boolean if a field has been set.
func (o *MasterBGDetail) HasSessionTimeout() bool {
	if o != nil && !IsNil(o.SessionTimeout) {
		return true
	}

	return false
}

// SetSessionTimeout gets a reference to the given int32 and assigns it to the SessionTimeout field.
func (o *MasterBGDetail) SetSessionTimeout(v int32) {
	o.SessionTimeout = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *MasterBGDetail) GetSubscription() Subscription {
	if o == nil || IsNil(o.Subscription) {
		var ret Subscription
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetSubscriptionOk() (*Subscription, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *MasterBGDetail) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given Subscription and assigns it to the Subscription field.
func (o *MasterBGDetail) SetSubscription(v Subscription) {
	o.Subscription = &v
}

// GetGdotId returns the GdotId field value if set, zero value otherwise.
func (o *MasterBGDetail) GetGdotId() string {
	if o == nil || IsNil(o.GdotId) {
		var ret string
		return ret
	}
	return *o.GdotId
}

// GetGdotIdOk returns a tuple with the GdotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetGdotIdOk() (*string, bool) {
	if o == nil || IsNil(o.GdotId) {
		return nil, false
	}
	return o.GdotId, true
}

// HasGdotId returns a boolean if a field has been set.
func (o *MasterBGDetail) HasGdotId() bool {
	if o != nil && !IsNil(o.GdotId) {
		return true
	}

	return false
}

// SetGdotId gets a reference to the given string and assigns it to the GdotId field.
func (o *MasterBGDetail) SetGdotId(v string) {
	o.GdotId = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MasterBGDetail) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt.Get()) {
		var ret string
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MasterBGDetail) GetDeletedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *MasterBGDetail) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableString and assigns it to the DeletedAt field.
func (o *MasterBGDetail) SetDeletedAt(v string) {
	o.DeletedAt.Set(&v)
}
// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *MasterBGDetail) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *MasterBGDetail) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetOrgType returns the OrgType field value if set, zero value otherwise.
func (o *MasterBGDetail) GetOrgType() string {
	if o == nil || IsNil(o.OrgType) {
		var ret string
		return ret
	}
	return *o.OrgType
}

// GetOrgTypeOk returns a tuple with the OrgType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetOrgTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OrgType) {
		return nil, false
	}
	return o.OrgType, true
}

// HasOrgType returns a boolean if a field has been set.
func (o *MasterBGDetail) HasOrgType() bool {
	if o != nil && !IsNil(o.OrgType) {
		return true
	}

	return false
}

// SetOrgType gets a reference to the given string and assigns it to the OrgType field.
func (o *MasterBGDetail) SetOrgType(v string) {
	o.OrgType = &v
}

func (o MasterBGDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MasterBGDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}
	if !IsNil(o.Environments) {
		toSerialize["environments"] = o.Environments
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IdproviderId) {
		toSerialize["idprovider_id"] = o.IdproviderId
	}
	if !IsNil(o.IsAutomaticAdminPromotionExempt) {
		toSerialize["isAutomaticAdminPromotionExempt"] = o.IsAutomaticAdminPromotionExempt
	}
	if !IsNil(o.IsFederated) {
		toSerialize["isFederated"] = o.IsFederated
	}
	if !IsNil(o.IsRoot) {
		toSerialize["isRoot"] = o.IsRoot
	}
	if !IsNil(o.IsMaster) {
		toSerialize["isMaster"] = o.IsMaster
	}
	if !IsNil(o.MfaRequired) {
		toSerialize["mfaRequired"] = o.MfaRequired
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if !IsNil(o.ParentOrganizationIds) {
		toSerialize["parentOrganizationIds"] = o.ParentOrganizationIds
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.SubOrganizationIds) {
		toSerialize["subOrganizationIds"] = o.SubOrganizationIds
	}
	if !IsNil(o.TenantOrganizationIds) {
		toSerialize["tenantOrganizationIds"] = o.TenantOrganizationIds
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.SessionTimeout) {
		toSerialize["sessionTimeout"] = o.SessionTimeout
	}
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	if !IsNil(o.GdotId) {
		toSerialize["gdotId"] = o.GdotId
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deletedAt"] = o.DeletedAt.Get()
	}
	if !IsNil(o.OrgType) {
		toSerialize["orgType"] = o.OrgType
	}
	return toSerialize, nil
}

type NullableMasterBGDetail struct {
	value *MasterBGDetail
	isSet bool
}

func (v NullableMasterBGDetail) Get() *MasterBGDetail {
	return v.value
}

func (v *NullableMasterBGDetail) Set(val *MasterBGDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableMasterBGDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableMasterBGDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMasterBGDetail(val *MasterBGDetail) *NullableMasterBGDetail {
	return &NullableMasterBGDetail{value: val, isSet: true}
}

func (v NullableMasterBGDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMasterBGDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


