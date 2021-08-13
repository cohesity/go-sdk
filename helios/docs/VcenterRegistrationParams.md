# VcenterRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Specifies the username to access target entity. | 
**Password** | **string** | Specifies the password to access target entity. | 
**Endpoint** | **string** | Specifies the endpoint IPaddress, URL or hostname of the host. | 
**Description** | Pointer to **NullableString** | Specifies the description of the source being registered. | [optional] 
**CaCert** | Pointer to **NullableString** | Specifies the CA certificate to enable SSL communication between host and cluster. | [optional] 
**UseVmBiosUuid** | Pointer to **NullableBool** | Specifies to use VM BIOS UUID to track virtual machines in the host. | [optional] 
**MinFreeDatastoreSpaceForBackupGb** | Pointer to **NullableInt64** | Specifies the minimum free space (in GB) expected to be available in the datastore where the virtual disks of the VM being backed up reside. If the space available is lower than the specified value, backup will be aborted. | [optional] 
**ThrottlingParams** | Pointer to [**VmwareThrottlingParams**](VmwareThrottlingParams.md) |  | [optional] 
**DataStoreParams** | Pointer to [**[]DatastoreParams**](DatastoreParams.md) | Specifies datastore specific parameters. | [optional] 

## Methods

### NewVcenterRegistrationParams

`func NewVcenterRegistrationParams(username string, password string, endpoint string, ) *VcenterRegistrationParams`

NewVcenterRegistrationParams instantiates a new VcenterRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterRegistrationParamsWithDefaults

`func NewVcenterRegistrationParamsWithDefaults() *VcenterRegistrationParams`

NewVcenterRegistrationParamsWithDefaults instantiates a new VcenterRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *VcenterRegistrationParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VcenterRegistrationParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VcenterRegistrationParams) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *VcenterRegistrationParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VcenterRegistrationParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VcenterRegistrationParams) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetEndpoint

`func (o *VcenterRegistrationParams) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *VcenterRegistrationParams) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *VcenterRegistrationParams) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetDescription

`func (o *VcenterRegistrationParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterRegistrationParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterRegistrationParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VcenterRegistrationParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VcenterRegistrationParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VcenterRegistrationParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCaCert

`func (o *VcenterRegistrationParams) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *VcenterRegistrationParams) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *VcenterRegistrationParams) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *VcenterRegistrationParams) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### SetCaCertNil

`func (o *VcenterRegistrationParams) SetCaCertNil(b bool)`

 SetCaCertNil sets the value for CaCert to be an explicit nil

### UnsetCaCert
`func (o *VcenterRegistrationParams) UnsetCaCert()`

UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
### GetUseVmBiosUuid

`func (o *VcenterRegistrationParams) GetUseVmBiosUuid() bool`

GetUseVmBiosUuid returns the UseVmBiosUuid field if non-nil, zero value otherwise.

### GetUseVmBiosUuidOk

`func (o *VcenterRegistrationParams) GetUseVmBiosUuidOk() (*bool, bool)`

GetUseVmBiosUuidOk returns a tuple with the UseVmBiosUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVmBiosUuid

`func (o *VcenterRegistrationParams) SetUseVmBiosUuid(v bool)`

SetUseVmBiosUuid sets UseVmBiosUuid field to given value.

### HasUseVmBiosUuid

`func (o *VcenterRegistrationParams) HasUseVmBiosUuid() bool`

HasUseVmBiosUuid returns a boolean if a field has been set.

### SetUseVmBiosUuidNil

`func (o *VcenterRegistrationParams) SetUseVmBiosUuidNil(b bool)`

 SetUseVmBiosUuidNil sets the value for UseVmBiosUuid to be an explicit nil

### UnsetUseVmBiosUuid
`func (o *VcenterRegistrationParams) UnsetUseVmBiosUuid()`

UnsetUseVmBiosUuid ensures that no value is present for UseVmBiosUuid, not even an explicit nil
### GetMinFreeDatastoreSpaceForBackupGb

`func (o *VcenterRegistrationParams) GetMinFreeDatastoreSpaceForBackupGb() int64`

GetMinFreeDatastoreSpaceForBackupGb returns the MinFreeDatastoreSpaceForBackupGb field if non-nil, zero value otherwise.

### GetMinFreeDatastoreSpaceForBackupGbOk

`func (o *VcenterRegistrationParams) GetMinFreeDatastoreSpaceForBackupGbOk() (*int64, bool)`

GetMinFreeDatastoreSpaceForBackupGbOk returns a tuple with the MinFreeDatastoreSpaceForBackupGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFreeDatastoreSpaceForBackupGb

`func (o *VcenterRegistrationParams) SetMinFreeDatastoreSpaceForBackupGb(v int64)`

SetMinFreeDatastoreSpaceForBackupGb sets MinFreeDatastoreSpaceForBackupGb field to given value.

### HasMinFreeDatastoreSpaceForBackupGb

`func (o *VcenterRegistrationParams) HasMinFreeDatastoreSpaceForBackupGb() bool`

HasMinFreeDatastoreSpaceForBackupGb returns a boolean if a field has been set.

### SetMinFreeDatastoreSpaceForBackupGbNil

`func (o *VcenterRegistrationParams) SetMinFreeDatastoreSpaceForBackupGbNil(b bool)`

 SetMinFreeDatastoreSpaceForBackupGbNil sets the value for MinFreeDatastoreSpaceForBackupGb to be an explicit nil

### UnsetMinFreeDatastoreSpaceForBackupGb
`func (o *VcenterRegistrationParams) UnsetMinFreeDatastoreSpaceForBackupGb()`

UnsetMinFreeDatastoreSpaceForBackupGb ensures that no value is present for MinFreeDatastoreSpaceForBackupGb, not even an explicit nil
### GetThrottlingParams

`func (o *VcenterRegistrationParams) GetThrottlingParams() VmwareThrottlingParams`

GetThrottlingParams returns the ThrottlingParams field if non-nil, zero value otherwise.

### GetThrottlingParamsOk

`func (o *VcenterRegistrationParams) GetThrottlingParamsOk() (*VmwareThrottlingParams, bool)`

GetThrottlingParamsOk returns a tuple with the ThrottlingParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingParams

`func (o *VcenterRegistrationParams) SetThrottlingParams(v VmwareThrottlingParams)`

SetThrottlingParams sets ThrottlingParams field to given value.

### HasThrottlingParams

`func (o *VcenterRegistrationParams) HasThrottlingParams() bool`

HasThrottlingParams returns a boolean if a field has been set.

### GetDataStoreParams

`func (o *VcenterRegistrationParams) GetDataStoreParams() []DatastoreParams`

GetDataStoreParams returns the DataStoreParams field if non-nil, zero value otherwise.

### GetDataStoreParamsOk

`func (o *VcenterRegistrationParams) GetDataStoreParamsOk() (*[]DatastoreParams, bool)`

GetDataStoreParamsOk returns a tuple with the DataStoreParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoreParams

`func (o *VcenterRegistrationParams) SetDataStoreParams(v []DatastoreParams)`

SetDataStoreParams sets DataStoreParams field to given value.

### HasDataStoreParams

`func (o *VcenterRegistrationParams) HasDataStoreParams() bool`

HasDataStoreParams returns a boolean if a field has been set.

### SetDataStoreParamsNil

`func (o *VcenterRegistrationParams) SetDataStoreParamsNil(b bool)`

 SetDataStoreParamsNil sets the value for DataStoreParams to be an explicit nil

### UnsetDataStoreParams
`func (o *VcenterRegistrationParams) UnsetDataStoreParams()`

UnsetDataStoreParams ensures that no value is present for DataStoreParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


