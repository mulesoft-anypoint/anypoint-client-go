/*
API Manager Contract API

API Manager Contract API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim_contract

import (
	"encoding/json"
)

// checks if the ContractDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractDetails{}

// ContractDetails struct for ContractDetails
type ContractDetails struct {
	Audit *Audit `json:"audit,omitempty"`
	MasterOrganizationId *string `json:"masterOrganizationId,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
	ApprovedDate NullableString `json:"approvedDate,omitempty"`
	RejectedDate NullableString `json:"rejectedDate,omitempty"`
	RevokedDate NullableString `json:"revokedDate,omitempty"`
	ApplicationId *int32 `json:"applicationId,omitempty"`
	Application *ApplicationContractResponse `json:"application,omitempty"`
	TierId *int32 `json:"tierId,omitempty"`
	Tier *TierContractResponse `json:"tier,omitempty"`
	RequestedTierId NullableInt32 `json:"requestedTierId,omitempty"`
	RequestedTier *TierContractResponse `json:"requestedTier,omitempty"`
	Condition *string `json:"condition,omitempty"`
	ApiId *int32 `json:"apiId,omitempty"`
	Api *ContractDetailsApi `json:"api,omitempty"`
}

// NewContractDetails instantiates a new ContractDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractDetails() *ContractDetails {
	this := ContractDetails{}
	var condition string = "APPLIED"
	this.Condition = &condition
	return &this
}

// NewContractDetailsWithDefaults instantiates a new ContractDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractDetailsWithDefaults() *ContractDetails {
	this := ContractDetails{}
	var condition string = "APPLIED"
	this.Condition = &condition
	return &this
}

// GetAudit returns the Audit field value if set, zero value otherwise.
func (o *ContractDetails) GetAudit() Audit {
	if o == nil || IsNil(o.Audit) {
		var ret Audit
		return ret
	}
	return *o.Audit
}

// GetAuditOk returns a tuple with the Audit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractDetails) GetAuditOk() (*Audit, bool) {
	if o == nil || IsNil(o.Audit) {
		return nil, false
	}
	return o.Audit, true
}

// HasAudit returns a boolean if a field has been set.
func (o *ContractDetails) HasAudit() bool {
	if o != nil && !IsNil(o.Audit) {
		return true
	}

	return false
}

// SetAudit gets a reference to the given Audit and assigns it to the Audit field.
func (o *ContractDetails) SetAudit(v Audit) {
	o.Audit = &v
}

// GetMasterOrganizationId returns the MasterOrganizationId field value if set, zero value otherwise.
func (o *ContractDetails) GetMasterOrganizationId() string {
	if o == nil || IsNil(o.MasterOrganizationId) {
		var ret string
		return ret
	}
	return *o.MasterOrganizationId
}

// GetMasterOrganizationIdOk returns a tuple with the MasterOrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractDetails) GetMasterOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.MasterOrganizationId) {
		return nil, false
	}
	return o.MasterOrganizationId, true
}

// HasMasterOrganizationId returns a boolean if a field has been set.
func (o *ContractDetails) HasMasterOrganizationId() bool {
	if o != nil && !IsNil(o.MasterOrganizationId) {
		return true
	}

	return false
}

// SetMasterOrganizationId gets a reference to the given string and assigns it to the MasterOrganizationId field.
func (o *ContractDetails) SetMasterOrganizationId(v string) {
	o.MasterOrganizationId = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ContractDetails) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractDetails) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ContractDetails) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *ContractDetails) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContractDetails) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractDetails) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContractDetails) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ContractDetails) SetId(v int32) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ContractDetails) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractDetails) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ContractDetails) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ContractDetails) SetStatus(v string) {
	o.Status = &v
}

// GetApprovedDate returns the ApprovedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContractDetails) GetApprovedDate() string {
	if o == nil || IsNil(o.ApprovedDate.Get()) {
		var ret string
		return ret
	}
	return *o.ApprovedDate.Get()
}

// GetApprovedDateOk returns a tuple with the ApprovedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContractDetails) GetApprovedDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApprovedDate.Get(), o.ApprovedDate.IsSet()
}

// HasApprovedDate returns a boolean if a field has been set.
func (o *ContractDetails) HasApprovedDate() bool {
	if o != nil && o.ApprovedDate.IsSet() {
		return true
	}

	return false
}

// SetApprovedDate gets a reference to the given NullableString and assigns it to the ApprovedDate field.
func (o *ContractDetails) SetApprovedDate(v string) {
	o.ApprovedDate.Set(&v)
}
// SetApprovedDateNil sets the value for ApprovedDate to be an explicit nil
func (o *ContractDetails) SetApprovedDateNil() {
	o.ApprovedDate.Set(nil)
}

// UnsetApprovedDate ensures that no value is present for ApprovedDate, not even an explicit nil
func (o *ContractDetails) UnsetApprovedDate() {
	o.ApprovedDate.Unset()
}

// GetRejectedDate returns the RejectedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContractDetails) GetRejectedDate() string {
	if o == nil || IsNil(o.RejectedDate.Get()) {
		var ret string
		return ret
	}
	return *o.RejectedDate.Get()
}

// GetRejectedDateOk returns a tuple with the RejectedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContractDetails) GetRejectedDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RejectedDate.Get(), o.RejectedDate.IsSet()
}

// HasRejectedDate returns a boolean if a field has been set.
func (o *ContractDetails) HasRejectedDate() bool {
	if o != nil && o.RejectedDate.IsSet() {
		return true
	}

	return false
}

// SetRejectedDate gets a reference to the given NullableString and assigns it to the RejectedDate field.
func (o *ContractDetails) SetRejectedDate(v string) {
	o.RejectedDate.Set(&v)
}
// SetRejectedDateNil sets the value for RejectedDate to be an explicit nil
func (o *ContractDetails) SetRejectedDateNil() {
	o.RejectedDate.Set(nil)
}

// UnsetRejectedDate ensures that no value is present for RejectedDate, not even an explicit nil
func (o *ContractDetails) UnsetRejectedDate() {
	o.RejectedDate.Unset()
}

// GetRevokedDate returns the RevokedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContractDetails) GetRevokedDate() string {
	if o == nil || IsNil(o.RevokedDate.Get()) {
		var ret string
		return ret
	}
	return *o.RevokedDate.Get()
}

// GetRevokedDateOk returns a tuple with the RevokedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContractDetails) GetRevokedDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RevokedDate.Get(), o.RevokedDate.IsSet()
}

// HasRevokedDate returns a boolean if a field has been set.
func (o *ContractDetails) HasRevokedDate() bool {
	if o != nil && o.RevokedDate.IsSet() {
		return true
	}

	return false
}

// SetRevokedDate gets a reference to the given NullableString and assigns it to the RevokedDate field.
func (o *ContractDetails) SetRevokedDate(v string) {
	o.RevokedDate.Set(&v)
}
// SetRevokedDateNil sets the value for RevokedDate to be an explicit nil
func (o *ContractDetails) SetRevokedDateNil() {
	o.RevokedDate.Set(nil)
}

// UnsetRevokedDate ensures that no value is present for RevokedDate, not even an explicit nil
func (o *ContractDetails) UnsetRevokedDate() {
	o.RevokedDate.Unset()
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *ContractDetails) GetApplicationId() int32 {
	if o == nil || IsNil(o.ApplicationId) {
		var ret int32
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractDetails) GetApplicationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *ContractDetails) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.
func (o *ContractDetails) SetApplicationId(v int32) {
	o.ApplicationId = &v
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *ContractDetails) GetApplication() ApplicationContractResponse {
	if o == nil || IsNil(o.Application) {
		var ret ApplicationContractResponse
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractDetails) GetApplicationOk() (*ApplicationContractResponse, bool) {
	if o == nil || IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *ContractDetails) HasApplication() bool {
	if o != nil && !IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given ApplicationContractResponse and assigns it to the Application field.
func (o *ContractDetails) SetApplication(v ApplicationContractResponse) {
	o.Application = &v
}

// GetTierId returns the TierId field value if set, zero value otherwise.
func (o *ContractDetails) GetTierId() int32 {
	if o == nil || IsNil(o.TierId) {
		var ret int32
		return ret
	}
	return *o.TierId
}

// GetTierIdOk returns a tuple with the TierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractDetails) GetTierIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TierId) {
		return nil, false
	}
	return o.TierId, true
}

// HasTierId returns a boolean if a field has been set.
func (o *ContractDetails) HasTierId() bool {
	if o != nil && !IsNil(o.TierId) {
		return true
	}

	return false
}

// SetTierId gets a reference to the given int32 and assigns it to the TierId field.
func (o *ContractDetails) SetTierId(v int32) {
	o.TierId = &v
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *ContractDetails) GetTier() TierContractResponse {
	if o == nil || IsNil(o.Tier) {
		var ret TierContractResponse
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractDetails) GetTierOk() (*TierContractResponse, bool) {
	if o == nil || IsNil(o.Tier) {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *ContractDetails) HasTier() bool {
	if o != nil && !IsNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given TierContractResponse and assigns it to the Tier field.
func (o *ContractDetails) SetTier(v TierContractResponse) {
	o.Tier = &v
}

// GetRequestedTierId returns the RequestedTierId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContractDetails) GetRequestedTierId() int32 {
	if o == nil || IsNil(o.RequestedTierId.Get()) {
		var ret int32
		return ret
	}
	return *o.RequestedTierId.Get()
}

// GetRequestedTierIdOk returns a tuple with the RequestedTierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContractDetails) GetRequestedTierIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestedTierId.Get(), o.RequestedTierId.IsSet()
}

// HasRequestedTierId returns a boolean if a field has been set.
func (o *ContractDetails) HasRequestedTierId() bool {
	if o != nil && o.RequestedTierId.IsSet() {
		return true
	}

	return false
}

// SetRequestedTierId gets a reference to the given NullableInt32 and assigns it to the RequestedTierId field.
func (o *ContractDetails) SetRequestedTierId(v int32) {
	o.RequestedTierId.Set(&v)
}
// SetRequestedTierIdNil sets the value for RequestedTierId to be an explicit nil
func (o *ContractDetails) SetRequestedTierIdNil() {
	o.RequestedTierId.Set(nil)
}

// UnsetRequestedTierId ensures that no value is present for RequestedTierId, not even an explicit nil
func (o *ContractDetails) UnsetRequestedTierId() {
	o.RequestedTierId.Unset()
}

// GetRequestedTier returns the RequestedTier field value if set, zero value otherwise.
func (o *ContractDetails) GetRequestedTier() TierContractResponse {
	if o == nil || IsNil(o.RequestedTier) {
		var ret TierContractResponse
		return ret
	}
	return *o.RequestedTier
}

// GetRequestedTierOk returns a tuple with the RequestedTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractDetails) GetRequestedTierOk() (*TierContractResponse, bool) {
	if o == nil || IsNil(o.RequestedTier) {
		return nil, false
	}
	return o.RequestedTier, true
}

// HasRequestedTier returns a boolean if a field has been set.
func (o *ContractDetails) HasRequestedTier() bool {
	if o != nil && !IsNil(o.RequestedTier) {
		return true
	}

	return false
}

// SetRequestedTier gets a reference to the given TierContractResponse and assigns it to the RequestedTier field.
func (o *ContractDetails) SetRequestedTier(v TierContractResponse) {
	o.RequestedTier = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *ContractDetails) GetCondition() string {
	if o == nil || IsNil(o.Condition) {
		var ret string
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractDetails) GetConditionOk() (*string, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *ContractDetails) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given string and assigns it to the Condition field.
func (o *ContractDetails) SetCondition(v string) {
	o.Condition = &v
}

// GetApiId returns the ApiId field value if set, zero value otherwise.
func (o *ContractDetails) GetApiId() int32 {
	if o == nil || IsNil(o.ApiId) {
		var ret int32
		return ret
	}
	return *o.ApiId
}

// GetApiIdOk returns a tuple with the ApiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractDetails) GetApiIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ApiId) {
		return nil, false
	}
	return o.ApiId, true
}

// HasApiId returns a boolean if a field has been set.
func (o *ContractDetails) HasApiId() bool {
	if o != nil && !IsNil(o.ApiId) {
		return true
	}

	return false
}

// SetApiId gets a reference to the given int32 and assigns it to the ApiId field.
func (o *ContractDetails) SetApiId(v int32) {
	o.ApiId = &v
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *ContractDetails) GetApi() ContractDetailsApi {
	if o == nil || IsNil(o.Api) {
		var ret ContractDetailsApi
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractDetails) GetApiOk() (*ContractDetailsApi, bool) {
	if o == nil || IsNil(o.Api) {
		return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *ContractDetails) HasApi() bool {
	if o != nil && !IsNil(o.Api) {
		return true
	}

	return false
}

// SetApi gets a reference to the given ContractDetailsApi and assigns it to the Api field.
func (o *ContractDetails) SetApi(v ContractDetailsApi) {
	o.Api = &v
}

func (o ContractDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractDetails) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.ApprovedDate.IsSet() {
		toSerialize["approvedDate"] = o.ApprovedDate.Get()
	}
	if o.RejectedDate.IsSet() {
		toSerialize["rejectedDate"] = o.RejectedDate.Get()
	}
	if o.RevokedDate.IsSet() {
		toSerialize["revokedDate"] = o.RevokedDate.Get()
	}
	if !IsNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !IsNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	if !IsNil(o.TierId) {
		toSerialize["tierId"] = o.TierId
	}
	if !IsNil(o.Tier) {
		toSerialize["tier"] = o.Tier
	}
	if o.RequestedTierId.IsSet() {
		toSerialize["requestedTierId"] = o.RequestedTierId.Get()
	}
	if !IsNil(o.RequestedTier) {
		toSerialize["requestedTier"] = o.RequestedTier
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	if !IsNil(o.ApiId) {
		toSerialize["apiId"] = o.ApiId
	}
	if !IsNil(o.Api) {
		toSerialize["api"] = o.Api
	}
	return toSerialize, nil
}

type NullableContractDetails struct {
	value *ContractDetails
	isSet bool
}

func (v NullableContractDetails) Get() *ContractDetails {
	return v.value
}

func (v *NullableContractDetails) Set(val *ContractDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableContractDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableContractDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractDetails(val *ContractDetails) *NullableContractDetails {
	return &NullableContractDetails{value: val, isSet: true}
}

func (v NullableContractDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


