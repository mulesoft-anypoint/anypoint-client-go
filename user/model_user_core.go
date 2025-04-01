/*
User API

Description of the User API

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package user

import (
	"encoding/json"
)

// checks if the UserCore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserCore{}

// UserCore struct for UserCore
type UserCore struct {
	Username *string `json:"username,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName *string `json:"lastName,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Email *string `json:"email,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Deleted *bool `json:"deleted,omitempty"`
	LastLogin *string `json:"lastLogin,omitempty"`
	MfaVerificationExcluded *bool `json:"mfaVerificationExcluded,omitempty"`
	MfaVerifiersConfigured *string `json:"mfaVerifiersConfigured,omitempty"`
	IdproviderId *string `json:"idprovider_id,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	IsFederated *bool `json:"isFederated,omitempty"`
	Type *string `json:"type,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
	MemberOfOrganizations []Org `json:"memberOfOrganizations,omitempty"`
	ContributorOfOrganizations []Org `json:"contributorOfOrganizations,omitempty"`
	Properties *Properties `json:"properties,omitempty"`
	OrganizationPreferences map[string]interface{} `json:"organizationPreferences,omitempty"`
	PrimaryOrganization *PrimaryOrganization `json:"primaryOrganization,omitempty"`
}

// NewUserCore instantiates a new UserCore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCore() *UserCore {
	this := UserCore{}
	var lastName string = "Mule"
	this.LastName = &lastName
	return &this
}

// NewUserCoreWithDefaults instantiates a new UserCore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCoreWithDefaults() *UserCore {
	this := UserCore{}
	var lastName string = "Mule"
	this.LastName = &lastName
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UserCore) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCore) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UserCore) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UserCore) SetUsername(v string) {
	o.Username = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UserCore) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCore) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UserCore) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UserCore) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UserCore) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCore) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UserCore) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UserCore) SetLastName(v string) {
	o.LastName = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *UserCore) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCore) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *UserCore) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *UserCore) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserCore) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCore) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserCore) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserCore) SetEmail(v string) {
	o.Email = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *UserCore) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCore) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *UserCore) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *UserCore) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UserCore) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCore) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UserCore) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UserCore) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *UserCore) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCore) GetDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *UserCore) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *UserCore) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise.
func (o *UserCore) GetLastLogin() string {
	if o == nil || IsNil(o.LastLogin) {
		var ret string
		return ret
	}
	return *o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCore) GetLastLoginOk() (*string, bool) {
	if o == nil || IsNil(o.LastLogin) {
		return nil, false
	}
	return o.LastLogin, true
}

// HasLastLogin returns a boolean if a field has been set.
func (o *UserCore) HasLastLogin() bool {
	if o != nil && !IsNil(o.LastLogin) {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given string and assigns it to the LastLogin field.
func (o *UserCore) SetLastLogin(v string) {
	o.LastLogin = &v
}

// GetMfaVerificationExcluded returns the MfaVerificationExcluded field value if set, zero value otherwise.
func (o *UserCore) GetMfaVerificationExcluded() bool {
	if o == nil || IsNil(o.MfaVerificationExcluded) {
		var ret bool
		return ret
	}
	return *o.MfaVerificationExcluded
}

// GetMfaVerificationExcludedOk returns a tuple with the MfaVerificationExcluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCore) GetMfaVerificationExcludedOk() (*bool, bool) {
	if o == nil || IsNil(o.MfaVerificationExcluded) {
		return nil, false
	}
	return o.MfaVerificationExcluded, true
}

// HasMfaVerificationExcluded returns a boolean if a field has been set.
func (o *UserCore) HasMfaVerificationExcluded() bool {
	if o != nil && !IsNil(o.MfaVerificationExcluded) {
		return true
	}

	return false
}

// SetMfaVerificationExcluded gets a reference to the given bool and assigns it to the MfaVerificationExcluded field.
func (o *UserCore) SetMfaVerificationExcluded(v bool) {
	o.MfaVerificationExcluded = &v
}

// GetMfaVerifiersConfigured returns the MfaVerifiersConfigured field value if set, zero value otherwise.
func (o *UserCore) GetMfaVerifiersConfigured() string {
	if o == nil || IsNil(o.MfaVerifiersConfigured) {
		var ret string
		return ret
	}
	return *o.MfaVerifiersConfigured
}

// GetMfaVerifiersConfiguredOk returns a tuple with the MfaVerifiersConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCore) GetMfaVerifiersConfiguredOk() (*string, bool) {
	if o == nil || IsNil(o.MfaVerifiersConfigured) {
		return nil, false
	}
	return o.MfaVerifiersConfigured, true
}

// HasMfaVerifiersConfigured returns a boolean if a field has been set.
func (o *UserCore) HasMfaVerifiersConfigured() bool {
	if o != nil && !IsNil(o.MfaVerifiersConfigured) {
		return true
	}

	return false
}

// SetMfaVerifiersConfigured gets a reference to the given string and assigns it to the MfaVerifiersConfigured field.
func (o *UserCore) SetMfaVerifiersConfigured(v string) {
	o.MfaVerifiersConfigured = &v
}

// GetIdproviderId returns the IdproviderId field value if set, zero value otherwise.
func (o *UserCore) GetIdproviderId() string {
	if o == nil || IsNil(o.IdproviderId) {
		var ret string
		return ret
	}
	return *o.IdproviderId
}

// GetIdproviderIdOk returns a tuple with the IdproviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCore) GetIdproviderIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdproviderId) {
		return nil, false
	}
	return o.IdproviderId, true
}

// HasIdproviderId returns a boolean if a field has been set.
func (o *UserCore) HasIdproviderId() bool {
	if o != nil && !IsNil(o.IdproviderId) {
		return true
	}

	return false
}

// SetIdproviderId gets a reference to the given string and assigns it to the IdproviderId field.
func (o *UserCore) SetIdproviderId(v string) {
	o.IdproviderId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UserCore) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCore) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UserCore) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *UserCore) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *UserCore) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCore) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *UserCore) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *UserCore) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetIsFederated returns the IsFederated field value if set, zero value otherwise.
func (o *UserCore) GetIsFederated() bool {
	if o == nil || IsNil(o.IsFederated) {
		var ret bool
		return ret
	}
	return *o.IsFederated
}

// GetIsFederatedOk returns a tuple with the IsFederated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCore) GetIsFederatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFederated) {
		return nil, false
	}
	return o.IsFederated, true
}

// HasIsFederated returns a boolean if a field has been set.
func (o *UserCore) HasIsFederated() bool {
	if o != nil && !IsNil(o.IsFederated) {
		return true
	}

	return false
}

// SetIsFederated gets a reference to the given bool and assigns it to the IsFederated field.
func (o *UserCore) SetIsFederated(v bool) {
	o.IsFederated = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserCore) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCore) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserCore) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserCore) SetType(v string) {
	o.Type = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *UserCore) GetOrganization() Organization {
	if o == nil || IsNil(o.Organization) {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCore) GetOrganizationOk() (*Organization, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *UserCore) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *UserCore) SetOrganization(v Organization) {
	o.Organization = &v
}

// GetMemberOfOrganizations returns the MemberOfOrganizations field value if set, zero value otherwise.
func (o *UserCore) GetMemberOfOrganizations() []Org {
	if o == nil || IsNil(o.MemberOfOrganizations) {
		var ret []Org
		return ret
	}
	return o.MemberOfOrganizations
}

// GetMemberOfOrganizationsOk returns a tuple with the MemberOfOrganizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCore) GetMemberOfOrganizationsOk() ([]Org, bool) {
	if o == nil || IsNil(o.MemberOfOrganizations) {
		return nil, false
	}
	return o.MemberOfOrganizations, true
}

// HasMemberOfOrganizations returns a boolean if a field has been set.
func (o *UserCore) HasMemberOfOrganizations() bool {
	if o != nil && !IsNil(o.MemberOfOrganizations) {
		return true
	}

	return false
}

// SetMemberOfOrganizations gets a reference to the given []Org and assigns it to the MemberOfOrganizations field.
func (o *UserCore) SetMemberOfOrganizations(v []Org) {
	o.MemberOfOrganizations = v
}

// GetContributorOfOrganizations returns the ContributorOfOrganizations field value if set, zero value otherwise.
func (o *UserCore) GetContributorOfOrganizations() []Org {
	if o == nil || IsNil(o.ContributorOfOrganizations) {
		var ret []Org
		return ret
	}
	return o.ContributorOfOrganizations
}

// GetContributorOfOrganizationsOk returns a tuple with the ContributorOfOrganizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCore) GetContributorOfOrganizationsOk() ([]Org, bool) {
	if o == nil || IsNil(o.ContributorOfOrganizations) {
		return nil, false
	}
	return o.ContributorOfOrganizations, true
}

// HasContributorOfOrganizations returns a boolean if a field has been set.
func (o *UserCore) HasContributorOfOrganizations() bool {
	if o != nil && !IsNil(o.ContributorOfOrganizations) {
		return true
	}

	return false
}

// SetContributorOfOrganizations gets a reference to the given []Org and assigns it to the ContributorOfOrganizations field.
func (o *UserCore) SetContributorOfOrganizations(v []Org) {
	o.ContributorOfOrganizations = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *UserCore) GetProperties() Properties {
	if o == nil || IsNil(o.Properties) {
		var ret Properties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCore) GetPropertiesOk() (*Properties, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *UserCore) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given Properties and assigns it to the Properties field.
func (o *UserCore) SetProperties(v Properties) {
	o.Properties = &v
}

// GetOrganizationPreferences returns the OrganizationPreferences field value if set, zero value otherwise.
func (o *UserCore) GetOrganizationPreferences() map[string]interface{} {
	if o == nil || IsNil(o.OrganizationPreferences) {
		var ret map[string]interface{}
		return ret
	}
	return o.OrganizationPreferences
}

// GetOrganizationPreferencesOk returns a tuple with the OrganizationPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCore) GetOrganizationPreferencesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.OrganizationPreferences) {
		return map[string]interface{}{}, false
	}
	return o.OrganizationPreferences, true
}

// HasOrganizationPreferences returns a boolean if a field has been set.
func (o *UserCore) HasOrganizationPreferences() bool {
	if o != nil && !IsNil(o.OrganizationPreferences) {
		return true
	}

	return false
}

// SetOrganizationPreferences gets a reference to the given map[string]interface{} and assigns it to the OrganizationPreferences field.
func (o *UserCore) SetOrganizationPreferences(v map[string]interface{}) {
	o.OrganizationPreferences = v
}

// GetPrimaryOrganization returns the PrimaryOrganization field value if set, zero value otherwise.
func (o *UserCore) GetPrimaryOrganization() PrimaryOrganization {
	if o == nil || IsNil(o.PrimaryOrganization) {
		var ret PrimaryOrganization
		return ret
	}
	return *o.PrimaryOrganization
}

// GetPrimaryOrganizationOk returns a tuple with the PrimaryOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCore) GetPrimaryOrganizationOk() (*PrimaryOrganization, bool) {
	if o == nil || IsNil(o.PrimaryOrganization) {
		return nil, false
	}
	return o.PrimaryOrganization, true
}

// HasPrimaryOrganization returns a boolean if a field has been set.
func (o *UserCore) HasPrimaryOrganization() bool {
	if o != nil && !IsNil(o.PrimaryOrganization) {
		return true
	}

	return false
}

// SetPrimaryOrganization gets a reference to the given PrimaryOrganization and assigns it to the PrimaryOrganization field.
func (o *UserCore) SetPrimaryOrganization(v PrimaryOrganization) {
	o.PrimaryOrganization = &v
}

func (o UserCore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserCore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	if !IsNil(o.LastLogin) {
		toSerialize["lastLogin"] = o.LastLogin
	}
	if !IsNil(o.MfaVerificationExcluded) {
		toSerialize["mfaVerificationExcluded"] = o.MfaVerificationExcluded
	}
	if !IsNil(o.MfaVerifiersConfigured) {
		toSerialize["mfaVerifiersConfigured"] = o.MfaVerifiersConfigured
	}
	if !IsNil(o.IdproviderId) {
		toSerialize["idprovider_id"] = o.IdproviderId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.IsFederated) {
		toSerialize["isFederated"] = o.IsFederated
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.MemberOfOrganizations) {
		toSerialize["memberOfOrganizations"] = o.MemberOfOrganizations
	}
	if !IsNil(o.ContributorOfOrganizations) {
		toSerialize["contributorOfOrganizations"] = o.ContributorOfOrganizations
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.OrganizationPreferences) {
		toSerialize["organizationPreferences"] = o.OrganizationPreferences
	}
	if !IsNil(o.PrimaryOrganization) {
		toSerialize["primaryOrganization"] = o.PrimaryOrganization
	}
	return toSerialize, nil
}

type NullableUserCore struct {
	value *UserCore
	isSet bool
}

func (v NullableUserCore) Get() *UserCore {
	return v.value
}

func (v *NullableUserCore) Set(val *UserCore) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCore) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCore(val *UserCore) *NullableUserCore {
	return &NullableUserCore{value: val, isSet: true}
}

func (v NullableUserCore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


