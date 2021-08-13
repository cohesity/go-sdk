# RecoverElastifileToElastifileVolumeTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or the original Elastifile target. | 
**NewSourceConfig** | Pointer to [**NullableRecoverOtherNasToElastifileVolumeTargetParams**](RecoverOtherNasToElastifileVolumeTargetParams.md) | Specifies the new destination Source configuration parameters where the volumes will be recovered. This is mandatory if recoverToNewSource is set to true. | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableOriginalElastifileTargetParams**](OriginalElastifileTargetParams.md) | Specifies the Source configuration if volumes are being recovered to original Source. If not specified, all the configuration parameters will be retained. | [optional] 

## Methods

### NewRecoverElastifileToElastifileVolumeTargetParams

`func NewRecoverElastifileToElastifileVolumeTargetParams(recoverToNewSource bool, ) *RecoverElastifileToElastifileVolumeTargetParams`

NewRecoverElastifileToElastifileVolumeTargetParams instantiates a new RecoverElastifileToElastifileVolumeTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverElastifileToElastifileVolumeTargetParamsWithDefaults

`func NewRecoverElastifileToElastifileVolumeTargetParamsWithDefaults() *RecoverElastifileToElastifileVolumeTargetParams`

NewRecoverElastifileToElastifileVolumeTargetParamsWithDefaults instantiates a new RecoverElastifileToElastifileVolumeTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoverToNewSource

`func (o *RecoverElastifileToElastifileVolumeTargetParams) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *RecoverElastifileToElastifileVolumeTargetParams) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *RecoverElastifileToElastifileVolumeTargetParams) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.


### GetNewSourceConfig

`func (o *RecoverElastifileToElastifileVolumeTargetParams) GetNewSourceConfig() RecoverOtherNasToElastifileVolumeTargetParams`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *RecoverElastifileToElastifileVolumeTargetParams) GetNewSourceConfigOk() (*RecoverOtherNasToElastifileVolumeTargetParams, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *RecoverElastifileToElastifileVolumeTargetParams) SetNewSourceConfig(v RecoverOtherNasToElastifileVolumeTargetParams)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *RecoverElastifileToElastifileVolumeTargetParams) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *RecoverElastifileToElastifileVolumeTargetParams) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *RecoverElastifileToElastifileVolumeTargetParams) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *RecoverElastifileToElastifileVolumeTargetParams) GetOriginalSourceConfig() OriginalElastifileTargetParams`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *RecoverElastifileToElastifileVolumeTargetParams) GetOriginalSourceConfigOk() (*OriginalElastifileTargetParams, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *RecoverElastifileToElastifileVolumeTargetParams) SetOriginalSourceConfig(v OriginalElastifileTargetParams)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *RecoverElastifileToElastifileVolumeTargetParams) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *RecoverElastifileToElastifileVolumeTargetParams) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *RecoverElastifileToElastifileVolumeTargetParams) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


