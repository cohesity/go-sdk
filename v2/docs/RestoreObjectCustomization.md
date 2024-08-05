# RestoreObjectCustomization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkConfig** | Pointer to [**NullableRestoreObjectCustomizationNetworkConfig**](RestoreObjectCustomizationNetworkConfig.md) |  | [optional] 
**ObjectId** | Pointer to **NullableInt64** | Specifies the object id of the VM. | [optional] 

## Methods

### NewRestoreObjectCustomization

`func NewRestoreObjectCustomization() *RestoreObjectCustomization`

NewRestoreObjectCustomization instantiates a new RestoreObjectCustomization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreObjectCustomizationWithDefaults

`func NewRestoreObjectCustomizationWithDefaults() *RestoreObjectCustomization`

NewRestoreObjectCustomizationWithDefaults instantiates a new RestoreObjectCustomization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkConfig

`func (o *RestoreObjectCustomization) GetNetworkConfig() RestoreObjectCustomizationNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *RestoreObjectCustomization) GetNetworkConfigOk() (*RestoreObjectCustomizationNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *RestoreObjectCustomization) SetNetworkConfig(v RestoreObjectCustomizationNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.

### HasNetworkConfig

`func (o *RestoreObjectCustomization) HasNetworkConfig() bool`

HasNetworkConfig returns a boolean if a field has been set.

### SetNetworkConfigNil

`func (o *RestoreObjectCustomization) SetNetworkConfigNil(b bool)`

 SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil

### UnsetNetworkConfig
`func (o *RestoreObjectCustomization) UnsetNetworkConfig()`

UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil
### GetObjectId

`func (o *RestoreObjectCustomization) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *RestoreObjectCustomization) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *RestoreObjectCustomization) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *RestoreObjectCustomization) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### SetObjectIdNil

`func (o *RestoreObjectCustomization) SetObjectIdNil(b bool)`

 SetObjectIdNil sets the value for ObjectId to be an explicit nil

### UnsetObjectId
`func (o *RestoreObjectCustomization) UnsetObjectId()`

UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


