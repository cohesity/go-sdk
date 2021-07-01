# RestoreKVMVMsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**DatacenterEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**PowerStateConfig** | Pointer to [**PowerStateConfigProto**](PowerStateConfigProto.md) |  | [optional] 
**RenameRestoredObjectParam** | Pointer to [**RenameObjectParamProto**](RenameObjectParamProto.md) |  | [optional] 
**RestoredObjectsNetworkConfig** | Pointer to [**RestoredObjectNetworkConfigProto**](RestoredObjectNetworkConfigProto.md) |  | [optional] 
**StoragedomainEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 

## Methods

### NewRestoreKVMVMsParams

`func NewRestoreKVMVMsParams() *RestoreKVMVMsParams`

NewRestoreKVMVMsParams instantiates a new RestoreKVMVMsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreKVMVMsParamsWithDefaults

`func NewRestoreKVMVMsParamsWithDefaults() *RestoreKVMVMsParams`

NewRestoreKVMVMsParamsWithDefaults instantiates a new RestoreKVMVMsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterEntity

`func (o *RestoreKVMVMsParams) GetClusterEntity() EntityProto`

GetClusterEntity returns the ClusterEntity field if non-nil, zero value otherwise.

### GetClusterEntityOk

`func (o *RestoreKVMVMsParams) GetClusterEntityOk() (*EntityProto, bool)`

GetClusterEntityOk returns a tuple with the ClusterEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterEntity

`func (o *RestoreKVMVMsParams) SetClusterEntity(v EntityProto)`

SetClusterEntity sets ClusterEntity field to given value.

### HasClusterEntity

`func (o *RestoreKVMVMsParams) HasClusterEntity() bool`

HasClusterEntity returns a boolean if a field has been set.

### GetDatacenterEntity

`func (o *RestoreKVMVMsParams) GetDatacenterEntity() EntityProto`

GetDatacenterEntity returns the DatacenterEntity field if non-nil, zero value otherwise.

### GetDatacenterEntityOk

`func (o *RestoreKVMVMsParams) GetDatacenterEntityOk() (*EntityProto, bool)`

GetDatacenterEntityOk returns a tuple with the DatacenterEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterEntity

`func (o *RestoreKVMVMsParams) SetDatacenterEntity(v EntityProto)`

SetDatacenterEntity sets DatacenterEntity field to given value.

### HasDatacenterEntity

`func (o *RestoreKVMVMsParams) HasDatacenterEntity() bool`

HasDatacenterEntity returns a boolean if a field has been set.

### GetPowerStateConfig

`func (o *RestoreKVMVMsParams) GetPowerStateConfig() PowerStateConfigProto`

GetPowerStateConfig returns the PowerStateConfig field if non-nil, zero value otherwise.

### GetPowerStateConfigOk

`func (o *RestoreKVMVMsParams) GetPowerStateConfigOk() (*PowerStateConfigProto, bool)`

GetPowerStateConfigOk returns a tuple with the PowerStateConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerStateConfig

`func (o *RestoreKVMVMsParams) SetPowerStateConfig(v PowerStateConfigProto)`

SetPowerStateConfig sets PowerStateConfig field to given value.

### HasPowerStateConfig

`func (o *RestoreKVMVMsParams) HasPowerStateConfig() bool`

HasPowerStateConfig returns a boolean if a field has been set.

### GetRenameRestoredObjectParam

`func (o *RestoreKVMVMsParams) GetRenameRestoredObjectParam() RenameObjectParamProto`

GetRenameRestoredObjectParam returns the RenameRestoredObjectParam field if non-nil, zero value otherwise.

### GetRenameRestoredObjectParamOk

`func (o *RestoreKVMVMsParams) GetRenameRestoredObjectParamOk() (*RenameObjectParamProto, bool)`

GetRenameRestoredObjectParamOk returns a tuple with the RenameRestoredObjectParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRestoredObjectParam

`func (o *RestoreKVMVMsParams) SetRenameRestoredObjectParam(v RenameObjectParamProto)`

SetRenameRestoredObjectParam sets RenameRestoredObjectParam field to given value.

### HasRenameRestoredObjectParam

`func (o *RestoreKVMVMsParams) HasRenameRestoredObjectParam() bool`

HasRenameRestoredObjectParam returns a boolean if a field has been set.

### GetRestoredObjectsNetworkConfig

`func (o *RestoreKVMVMsParams) GetRestoredObjectsNetworkConfig() RestoredObjectNetworkConfigProto`

GetRestoredObjectsNetworkConfig returns the RestoredObjectsNetworkConfig field if non-nil, zero value otherwise.

### GetRestoredObjectsNetworkConfigOk

`func (o *RestoreKVMVMsParams) GetRestoredObjectsNetworkConfigOk() (*RestoredObjectNetworkConfigProto, bool)`

GetRestoredObjectsNetworkConfigOk returns a tuple with the RestoredObjectsNetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredObjectsNetworkConfig

`func (o *RestoreKVMVMsParams) SetRestoredObjectsNetworkConfig(v RestoredObjectNetworkConfigProto)`

SetRestoredObjectsNetworkConfig sets RestoredObjectsNetworkConfig field to given value.

### HasRestoredObjectsNetworkConfig

`func (o *RestoreKVMVMsParams) HasRestoredObjectsNetworkConfig() bool`

HasRestoredObjectsNetworkConfig returns a boolean if a field has been set.

### GetStoragedomainEntity

`func (o *RestoreKVMVMsParams) GetStoragedomainEntity() EntityProto`

GetStoragedomainEntity returns the StoragedomainEntity field if non-nil, zero value otherwise.

### GetStoragedomainEntityOk

`func (o *RestoreKVMVMsParams) GetStoragedomainEntityOk() (*EntityProto, bool)`

GetStoragedomainEntityOk returns a tuple with the StoragedomainEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragedomainEntity

`func (o *RestoreKVMVMsParams) SetStoragedomainEntity(v EntityProto)`

SetStoragedomainEntity sets StoragedomainEntity field to given value.

### HasStoragedomainEntity

`func (o *RestoreKVMVMsParams) HasStoragedomainEntity() bool`

HasStoragedomainEntity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


