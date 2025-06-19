# PrivateSpaceNetworkEditable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** | The region of the Private Space. Required when create a Private Space network. | [optional] 
**CidrBlock** | Pointer to **string** | The CIDR block of the VPC where the Private Space will be created. Required when create a Private Space in a new VPC. Default is 10.0.0.0/16. | [optional] 
**InternalDns** | Pointer to [**PrivateSpaceNetworkEditableInternalDns**](PrivateSpaceNetworkEditableInternalDns.md) |  | [optional] 
**ReservedCidrs** | Pointer to **[]string** | The list of reserved CIDR blocks for your private space to prevent IP address overlap. Required when you want to connect your private space to your corporate network (either on-premises or in the cloud). Use CIDR notation and commas.  | [optional] 

## Methods

### NewPrivateSpaceNetworkEditable

`func NewPrivateSpaceNetworkEditable() *PrivateSpaceNetworkEditable`

NewPrivateSpaceNetworkEditable instantiates a new PrivateSpaceNetworkEditable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpaceNetworkEditableWithDefaults

`func NewPrivateSpaceNetworkEditableWithDefaults() *PrivateSpaceNetworkEditable`

NewPrivateSpaceNetworkEditableWithDefaults instantiates a new PrivateSpaceNetworkEditable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *PrivateSpaceNetworkEditable) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PrivateSpaceNetworkEditable) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PrivateSpaceNetworkEditable) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PrivateSpaceNetworkEditable) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCidrBlock

`func (o *PrivateSpaceNetworkEditable) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *PrivateSpaceNetworkEditable) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *PrivateSpaceNetworkEditable) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *PrivateSpaceNetworkEditable) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### GetInternalDns

`func (o *PrivateSpaceNetworkEditable) GetInternalDns() PrivateSpaceNetworkEditableInternalDns`

GetInternalDns returns the InternalDns field if non-nil, zero value otherwise.

### GetInternalDnsOk

`func (o *PrivateSpaceNetworkEditable) GetInternalDnsOk() (*PrivateSpaceNetworkEditableInternalDns, bool)`

GetInternalDnsOk returns a tuple with the InternalDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDns

`func (o *PrivateSpaceNetworkEditable) SetInternalDns(v PrivateSpaceNetworkEditableInternalDns)`

SetInternalDns sets InternalDns field to given value.

### HasInternalDns

`func (o *PrivateSpaceNetworkEditable) HasInternalDns() bool`

HasInternalDns returns a boolean if a field has been set.

### GetReservedCidrs

`func (o *PrivateSpaceNetworkEditable) GetReservedCidrs() []string`

GetReservedCidrs returns the ReservedCidrs field if non-nil, zero value otherwise.

### GetReservedCidrsOk

`func (o *PrivateSpaceNetworkEditable) GetReservedCidrsOk() (*[]string, bool)`

GetReservedCidrsOk returns a tuple with the ReservedCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedCidrs

`func (o *PrivateSpaceNetworkEditable) SetReservedCidrs(v []string)`

SetReservedCidrs sets ReservedCidrs field to given value.

### HasReservedCidrs

`func (o *PrivateSpaceNetworkEditable) HasReservedCidrs() bool`

HasReservedCidrs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


