# VcenterRegistrationParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaCert** | Pointer to **NullableString** | Specifies the CA certificate to enable SSL communication between host and cluster. | [optional] 
**UseVmBiosUuid** | Pointer to **NullableBool** | Specifies to use VM BIOS UUID to track virtual machines in the host. | [optional] 
**MinFreeDatastoreSpaceForBackupGb** | Pointer to **NullableInt64** | Specifies the minimum free space (in GB) expected to be available in the datastore where the virtual disks of the VM being backed up reside. If the space available is lower than the specified value, backup will be aborted. | [optional] 
**ThrottlingParams** | Pointer to [**VmwareThrottlingParams**](VmwareThrottlingParams.md) |  | [optional] 
**DataStoreParams** | Pointer to [**[]DatastoreParams**](DatastoreParams.md) | Specifies datastore specific parameters. | [optional] 

## Methods

### NewVcenterRegistrationParamsAllOf

`func NewVcenterRegistrationParamsAllOf() *VcenterRegistrationParamsAllOf`

NewVcenterRegistrationParamsAllOf instantiates a new VcenterRegistrationParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterRegistrationParamsAllOfWithDefaults

`func NewVcenterRegistrationParamsAllOfWithDefaults() *VcenterRegistrationParamsAllOf`

NewVcenterRegistrationParamsAllOfWithDefaults instantiates a new VcenterRegistrationParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaCert

`func (o *VcenterRegistrationParamsAllOf) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *VcenterRegistrationParamsAllOf) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *VcenterRegistrationParamsAllOf) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *VcenterRegistrationParamsAllOf) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### SetCaCertNil

`func (o *VcenterRegistrationParamsAllOf) SetCaCertNil(b bool)`

 SetCaCertNil sets the value for CaCert to be an explicit nil

### UnsetCaCert
`func (o *VcenterRegistrationParamsAllOf) UnsetCaCert()`

UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
### GetUseVmBiosUuid

`func (o *VcenterRegistrationParamsAllOf) GetUseVmBiosUuid() bool`

GetUseVmBiosUuid returns the UseVmBiosUuid field if non-nil, zero value otherwise.

### GetUseVmBiosUuidOk

`func (o *VcenterRegistrationParamsAllOf) GetUseVmBiosUuidOk() (*bool, bool)`

GetUseVmBiosUuidOk returns a tuple with the UseVmBiosUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVmBiosUuid

`func (o *VcenterRegistrationParamsAllOf) SetUseVmBiosUuid(v bool)`

SetUseVmBiosUuid sets UseVmBiosUuid field to given value.

### HasUseVmBiosUuid

`func (o *VcenterRegistrationParamsAllOf) HasUseVmBiosUuid() bool`

HasUseVmBiosUuid returns a boolean if a field has been set.

### SetUseVmBiosUuidNil

`func (o *VcenterRegistrationParamsAllOf) SetUseVmBiosUuidNil(b bool)`

 SetUseVmBiosUuidNil sets the value for UseVmBiosUuid to be an explicit nil

### UnsetUseVmBiosUuid
`func (o *VcenterRegistrationParamsAllOf) UnsetUseVmBiosUuid()`

UnsetUseVmBiosUuid ensures that no value is present for UseVmBiosUuid, not even an explicit nil
### GetMinFreeDatastoreSpaceForBackupGb

`func (o *VcenterRegistrationParamsAllOf) GetMinFreeDatastoreSpaceForBackupGb() int64`

GetMinFreeDatastoreSpaceForBackupGb returns the MinFreeDatastoreSpaceForBackupGb field if non-nil, zero value otherwise.

### GetMinFreeDatastoreSpaceForBackupGbOk

`func (o *VcenterRegistrationParamsAllOf) GetMinFreeDatastoreSpaceForBackupGbOk() (*int64, bool)`

GetMinFreeDatastoreSpaceForBackupGbOk returns a tuple with the MinFreeDatastoreSpaceForBackupGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFreeDatastoreSpaceForBackupGb

`func (o *VcenterRegistrationParamsAllOf) SetMinFreeDatastoreSpaceForBackupGb(v int64)`

SetMinFreeDatastoreSpaceForBackupGb sets MinFreeDatastoreSpaceForBackupGb field to given value.

### HasMinFreeDatastoreSpaceForBackupGb

`func (o *VcenterRegistrationParamsAllOf) HasMinFreeDatastoreSpaceForBackupGb() bool`

HasMinFreeDatastoreSpaceForBackupGb returns a boolean if a field has been set.

### SetMinFreeDatastoreSpaceForBackupGbNil

`func (o *VcenterRegistrationParamsAllOf) SetMinFreeDatastoreSpaceForBackupGbNil(b bool)`

 SetMinFreeDatastoreSpaceForBackupGbNil sets the value for MinFreeDatastoreSpaceForBackupGb to be an explicit nil

### UnsetMinFreeDatastoreSpaceForBackupGb
`func (o *VcenterRegistrationParamsAllOf) UnsetMinFreeDatastoreSpaceForBackupGb()`

UnsetMinFreeDatastoreSpaceForBackupGb ensures that no value is present for MinFreeDatastoreSpaceForBackupGb, not even an explicit nil
### GetThrottlingParams

`func (o *VcenterRegistrationParamsAllOf) GetThrottlingParams() VmwareThrottlingParams`

GetThrottlingParams returns the ThrottlingParams field if non-nil, zero value otherwise.

### GetThrottlingParamsOk

`func (o *VcenterRegistrationParamsAllOf) GetThrottlingParamsOk() (*VmwareThrottlingParams, bool)`

GetThrottlingParamsOk returns a tuple with the ThrottlingParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingParams

`func (o *VcenterRegistrationParamsAllOf) SetThrottlingParams(v VmwareThrottlingParams)`

SetThrottlingParams sets ThrottlingParams field to given value.

### HasThrottlingParams

`func (o *VcenterRegistrationParamsAllOf) HasThrottlingParams() bool`

HasThrottlingParams returns a boolean if a field has been set.

### GetDataStoreParams

`func (o *VcenterRegistrationParamsAllOf) GetDataStoreParams() []DatastoreParams`

GetDataStoreParams returns the DataStoreParams field if non-nil, zero value otherwise.

### GetDataStoreParamsOk

`func (o *VcenterRegistrationParamsAllOf) GetDataStoreParamsOk() (*[]DatastoreParams, bool)`

GetDataStoreParamsOk returns a tuple with the DataStoreParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoreParams

`func (o *VcenterRegistrationParamsAllOf) SetDataStoreParams(v []DatastoreParams)`

SetDataStoreParams sets DataStoreParams field to given value.

### HasDataStoreParams

`func (o *VcenterRegistrationParamsAllOf) HasDataStoreParams() bool`

HasDataStoreParams returns a boolean if a field has been set.

### SetDataStoreParamsNil

`func (o *VcenterRegistrationParamsAllOf) SetDataStoreParamsNil(b bool)`

 SetDataStoreParamsNil sets the value for DataStoreParams to be an explicit nil

### UnsetDataStoreParams
`func (o *VcenterRegistrationParamsAllOf) UnsetDataStoreParams()`

UnsetDataStoreParams ensures that no value is present for DataStoreParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


