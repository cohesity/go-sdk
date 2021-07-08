# RestoreObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **NullableInt32** | The action to perform. | [optional] 
**DatastoreEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**PowerStateConfig** | Pointer to [**PowerStateConfigProto**](PowerStateConfigProto.md) |  | [optional] 
**RenameRestoredObjectParam** | Pointer to [**RenameObjectParamProto**](RenameObjectParamProto.md) |  | [optional] 
**ResourcePoolEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**RestoreParentSource** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**RestoredObjectsNetworkConfig** | Pointer to [**RestoredObjectNetworkConfigProto**](RestoredObjectNetworkConfigProto.md) |  | [optional] 
**ViewName** | Pointer to **NullableString** | Target view into which the objects are to be cloned. | [optional] 

## Methods

### NewRestoreObjectParams

`func NewRestoreObjectParams() *RestoreObjectParams`

NewRestoreObjectParams instantiates a new RestoreObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreObjectParamsWithDefaults

`func NewRestoreObjectParamsWithDefaults() *RestoreObjectParams`

NewRestoreObjectParamsWithDefaults instantiates a new RestoreObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RestoreObjectParams) GetAction() int32`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RestoreObjectParams) GetActionOk() (*int32, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RestoreObjectParams) SetAction(v int32)`

SetAction sets Action field to given value.

### HasAction

`func (o *RestoreObjectParams) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *RestoreObjectParams) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *RestoreObjectParams) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetDatastoreEntity

`func (o *RestoreObjectParams) GetDatastoreEntity() EntityProto`

GetDatastoreEntity returns the DatastoreEntity field if non-nil, zero value otherwise.

### GetDatastoreEntityOk

`func (o *RestoreObjectParams) GetDatastoreEntityOk() (*EntityProto, bool)`

GetDatastoreEntityOk returns a tuple with the DatastoreEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreEntity

`func (o *RestoreObjectParams) SetDatastoreEntity(v EntityProto)`

SetDatastoreEntity sets DatastoreEntity field to given value.

### HasDatastoreEntity

`func (o *RestoreObjectParams) HasDatastoreEntity() bool`

HasDatastoreEntity returns a boolean if a field has been set.

### GetPowerStateConfig

`func (o *RestoreObjectParams) GetPowerStateConfig() PowerStateConfigProto`

GetPowerStateConfig returns the PowerStateConfig field if non-nil, zero value otherwise.

### GetPowerStateConfigOk

`func (o *RestoreObjectParams) GetPowerStateConfigOk() (*PowerStateConfigProto, bool)`

GetPowerStateConfigOk returns a tuple with the PowerStateConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerStateConfig

`func (o *RestoreObjectParams) SetPowerStateConfig(v PowerStateConfigProto)`

SetPowerStateConfig sets PowerStateConfig field to given value.

### HasPowerStateConfig

`func (o *RestoreObjectParams) HasPowerStateConfig() bool`

HasPowerStateConfig returns a boolean if a field has been set.

### GetRenameRestoredObjectParam

`func (o *RestoreObjectParams) GetRenameRestoredObjectParam() RenameObjectParamProto`

GetRenameRestoredObjectParam returns the RenameRestoredObjectParam field if non-nil, zero value otherwise.

### GetRenameRestoredObjectParamOk

`func (o *RestoreObjectParams) GetRenameRestoredObjectParamOk() (*RenameObjectParamProto, bool)`

GetRenameRestoredObjectParamOk returns a tuple with the RenameRestoredObjectParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRestoredObjectParam

`func (o *RestoreObjectParams) SetRenameRestoredObjectParam(v RenameObjectParamProto)`

SetRenameRestoredObjectParam sets RenameRestoredObjectParam field to given value.

### HasRenameRestoredObjectParam

`func (o *RestoreObjectParams) HasRenameRestoredObjectParam() bool`

HasRenameRestoredObjectParam returns a boolean if a field has been set.

### GetResourcePoolEntity

`func (o *RestoreObjectParams) GetResourcePoolEntity() EntityProto`

GetResourcePoolEntity returns the ResourcePoolEntity field if non-nil, zero value otherwise.

### GetResourcePoolEntityOk

`func (o *RestoreObjectParams) GetResourcePoolEntityOk() (*EntityProto, bool)`

GetResourcePoolEntityOk returns a tuple with the ResourcePoolEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolEntity

`func (o *RestoreObjectParams) SetResourcePoolEntity(v EntityProto)`

SetResourcePoolEntity sets ResourcePoolEntity field to given value.

### HasResourcePoolEntity

`func (o *RestoreObjectParams) HasResourcePoolEntity() bool`

HasResourcePoolEntity returns a boolean if a field has been set.

### GetRestoreParentSource

`func (o *RestoreObjectParams) GetRestoreParentSource() EntityProto`

GetRestoreParentSource returns the RestoreParentSource field if non-nil, zero value otherwise.

### GetRestoreParentSourceOk

`func (o *RestoreObjectParams) GetRestoreParentSourceOk() (*EntityProto, bool)`

GetRestoreParentSourceOk returns a tuple with the RestoreParentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreParentSource

`func (o *RestoreObjectParams) SetRestoreParentSource(v EntityProto)`

SetRestoreParentSource sets RestoreParentSource field to given value.

### HasRestoreParentSource

`func (o *RestoreObjectParams) HasRestoreParentSource() bool`

HasRestoreParentSource returns a boolean if a field has been set.

### GetRestoredObjectsNetworkConfig

`func (o *RestoreObjectParams) GetRestoredObjectsNetworkConfig() RestoredObjectNetworkConfigProto`

GetRestoredObjectsNetworkConfig returns the RestoredObjectsNetworkConfig field if non-nil, zero value otherwise.

### GetRestoredObjectsNetworkConfigOk

`func (o *RestoreObjectParams) GetRestoredObjectsNetworkConfigOk() (*RestoredObjectNetworkConfigProto, bool)`

GetRestoredObjectsNetworkConfigOk returns a tuple with the RestoredObjectsNetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredObjectsNetworkConfig

`func (o *RestoreObjectParams) SetRestoredObjectsNetworkConfig(v RestoredObjectNetworkConfigProto)`

SetRestoredObjectsNetworkConfig sets RestoredObjectsNetworkConfig field to given value.

### HasRestoredObjectsNetworkConfig

`func (o *RestoreObjectParams) HasRestoredObjectsNetworkConfig() bool`

HasRestoredObjectsNetworkConfig returns a boolean if a field has been set.

### GetViewName

`func (o *RestoreObjectParams) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *RestoreObjectParams) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *RestoreObjectParams) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *RestoreObjectParams) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *RestoreObjectParams) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *RestoreObjectParams) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


