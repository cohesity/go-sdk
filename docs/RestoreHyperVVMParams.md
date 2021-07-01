# RestoreHyperVVMParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CopyRecovery** | Pointer to **NullableBool** | Whether to perform copy recovery. | [optional] 
**DatastoreEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**PowerStateConfig** | Pointer to [**PowerStateConfigProto**](PowerStateConfigProto.md) |  | [optional] 
**RenameRestoredObjectParam** | Pointer to [**RenameObjectParamProto**](RenameObjectParamProto.md) |  | [optional] 
**ResourceEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**RestoredObjectsNetworkConfig** | Pointer to [**RestoredObjectNetworkConfigProto**](RestoredObjectNetworkConfigProto.md) |  | [optional] 
**UuidConfig** | Pointer to [**UuidConfigProto**](UuidConfigProto.md) |  | [optional] 

## Methods

### NewRestoreHyperVVMParams

`func NewRestoreHyperVVMParams() *RestoreHyperVVMParams`

NewRestoreHyperVVMParams instantiates a new RestoreHyperVVMParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreHyperVVMParamsWithDefaults

`func NewRestoreHyperVVMParamsWithDefaults() *RestoreHyperVVMParams`

NewRestoreHyperVVMParamsWithDefaults instantiates a new RestoreHyperVVMParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopyRecovery

`func (o *RestoreHyperVVMParams) GetCopyRecovery() bool`

GetCopyRecovery returns the CopyRecovery field if non-nil, zero value otherwise.

### GetCopyRecoveryOk

`func (o *RestoreHyperVVMParams) GetCopyRecoveryOk() (*bool, bool)`

GetCopyRecoveryOk returns a tuple with the CopyRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyRecovery

`func (o *RestoreHyperVVMParams) SetCopyRecovery(v bool)`

SetCopyRecovery sets CopyRecovery field to given value.

### HasCopyRecovery

`func (o *RestoreHyperVVMParams) HasCopyRecovery() bool`

HasCopyRecovery returns a boolean if a field has been set.

### SetCopyRecoveryNil

`func (o *RestoreHyperVVMParams) SetCopyRecoveryNil(b bool)`

 SetCopyRecoveryNil sets the value for CopyRecovery to be an explicit nil

### UnsetCopyRecovery
`func (o *RestoreHyperVVMParams) UnsetCopyRecovery()`

UnsetCopyRecovery ensures that no value is present for CopyRecovery, not even an explicit nil
### GetDatastoreEntity

`func (o *RestoreHyperVVMParams) GetDatastoreEntity() EntityProto`

GetDatastoreEntity returns the DatastoreEntity field if non-nil, zero value otherwise.

### GetDatastoreEntityOk

`func (o *RestoreHyperVVMParams) GetDatastoreEntityOk() (*EntityProto, bool)`

GetDatastoreEntityOk returns a tuple with the DatastoreEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreEntity

`func (o *RestoreHyperVVMParams) SetDatastoreEntity(v EntityProto)`

SetDatastoreEntity sets DatastoreEntity field to given value.

### HasDatastoreEntity

`func (o *RestoreHyperVVMParams) HasDatastoreEntity() bool`

HasDatastoreEntity returns a boolean if a field has been set.

### GetPowerStateConfig

`func (o *RestoreHyperVVMParams) GetPowerStateConfig() PowerStateConfigProto`

GetPowerStateConfig returns the PowerStateConfig field if non-nil, zero value otherwise.

### GetPowerStateConfigOk

`func (o *RestoreHyperVVMParams) GetPowerStateConfigOk() (*PowerStateConfigProto, bool)`

GetPowerStateConfigOk returns a tuple with the PowerStateConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerStateConfig

`func (o *RestoreHyperVVMParams) SetPowerStateConfig(v PowerStateConfigProto)`

SetPowerStateConfig sets PowerStateConfig field to given value.

### HasPowerStateConfig

`func (o *RestoreHyperVVMParams) HasPowerStateConfig() bool`

HasPowerStateConfig returns a boolean if a field has been set.

### GetRenameRestoredObjectParam

`func (o *RestoreHyperVVMParams) GetRenameRestoredObjectParam() RenameObjectParamProto`

GetRenameRestoredObjectParam returns the RenameRestoredObjectParam field if non-nil, zero value otherwise.

### GetRenameRestoredObjectParamOk

`func (o *RestoreHyperVVMParams) GetRenameRestoredObjectParamOk() (*RenameObjectParamProto, bool)`

GetRenameRestoredObjectParamOk returns a tuple with the RenameRestoredObjectParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRestoredObjectParam

`func (o *RestoreHyperVVMParams) SetRenameRestoredObjectParam(v RenameObjectParamProto)`

SetRenameRestoredObjectParam sets RenameRestoredObjectParam field to given value.

### HasRenameRestoredObjectParam

`func (o *RestoreHyperVVMParams) HasRenameRestoredObjectParam() bool`

HasRenameRestoredObjectParam returns a boolean if a field has been set.

### GetResourceEntity

`func (o *RestoreHyperVVMParams) GetResourceEntity() EntityProto`

GetResourceEntity returns the ResourceEntity field if non-nil, zero value otherwise.

### GetResourceEntityOk

`func (o *RestoreHyperVVMParams) GetResourceEntityOk() (*EntityProto, bool)`

GetResourceEntityOk returns a tuple with the ResourceEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceEntity

`func (o *RestoreHyperVVMParams) SetResourceEntity(v EntityProto)`

SetResourceEntity sets ResourceEntity field to given value.

### HasResourceEntity

`func (o *RestoreHyperVVMParams) HasResourceEntity() bool`

HasResourceEntity returns a boolean if a field has been set.

### GetRestoredObjectsNetworkConfig

`func (o *RestoreHyperVVMParams) GetRestoredObjectsNetworkConfig() RestoredObjectNetworkConfigProto`

GetRestoredObjectsNetworkConfig returns the RestoredObjectsNetworkConfig field if non-nil, zero value otherwise.

### GetRestoredObjectsNetworkConfigOk

`func (o *RestoreHyperVVMParams) GetRestoredObjectsNetworkConfigOk() (*RestoredObjectNetworkConfigProto, bool)`

GetRestoredObjectsNetworkConfigOk returns a tuple with the RestoredObjectsNetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredObjectsNetworkConfig

`func (o *RestoreHyperVVMParams) SetRestoredObjectsNetworkConfig(v RestoredObjectNetworkConfigProto)`

SetRestoredObjectsNetworkConfig sets RestoredObjectsNetworkConfig field to given value.

### HasRestoredObjectsNetworkConfig

`func (o *RestoreHyperVVMParams) HasRestoredObjectsNetworkConfig() bool`

HasRestoredObjectsNetworkConfig returns a boolean if a field has been set.

### GetUuidConfig

`func (o *RestoreHyperVVMParams) GetUuidConfig() UuidConfigProto`

GetUuidConfig returns the UuidConfig field if non-nil, zero value otherwise.

### GetUuidConfigOk

`func (o *RestoreHyperVVMParams) GetUuidConfigOk() (*UuidConfigProto, bool)`

GetUuidConfigOk returns a tuple with the UuidConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuidConfig

`func (o *RestoreHyperVVMParams) SetUuidConfig(v UuidConfigProto)`

SetUuidConfig sets UuidConfig field to given value.

### HasUuidConfig

`func (o *RestoreHyperVVMParams) HasUuidConfig() bool`

HasUuidConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


