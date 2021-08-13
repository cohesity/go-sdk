# RecoverNetappToNetappFilesTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or the original Netapp target. | 
**NewSourceConfig** | Pointer to [**NullableRecoverOtherNasToNetappFilesTargetParams**](RecoverOtherNasToNetappFilesTargetParams.md) | Specifies the new destination Source configuration parameters where the files will be recovered. This is mandatory if recoverToNewSource is set to true. | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableOriginalNetappFilesTargetParams**](OriginalNetappFilesTargetParams.md) | Specifies the Source configuration if files are being recovered to original Source. If not specified, all the configuration parameters will be retained. | [optional] 

## Methods

### NewRecoverNetappToNetappFilesTargetParams

`func NewRecoverNetappToNetappFilesTargetParams(recoverToNewSource bool, ) *RecoverNetappToNetappFilesTargetParams`

NewRecoverNetappToNetappFilesTargetParams instantiates a new RecoverNetappToNetappFilesTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverNetappToNetappFilesTargetParamsWithDefaults

`func NewRecoverNetappToNetappFilesTargetParamsWithDefaults() *RecoverNetappToNetappFilesTargetParams`

NewRecoverNetappToNetappFilesTargetParamsWithDefaults instantiates a new RecoverNetappToNetappFilesTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoverToNewSource

`func (o *RecoverNetappToNetappFilesTargetParams) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *RecoverNetappToNetappFilesTargetParams) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *RecoverNetappToNetappFilesTargetParams) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.


### GetNewSourceConfig

`func (o *RecoverNetappToNetappFilesTargetParams) GetNewSourceConfig() RecoverOtherNasToNetappFilesTargetParams`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *RecoverNetappToNetappFilesTargetParams) GetNewSourceConfigOk() (*RecoverOtherNasToNetappFilesTargetParams, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *RecoverNetappToNetappFilesTargetParams) SetNewSourceConfig(v RecoverOtherNasToNetappFilesTargetParams)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *RecoverNetappToNetappFilesTargetParams) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *RecoverNetappToNetappFilesTargetParams) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *RecoverNetappToNetappFilesTargetParams) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *RecoverNetappToNetappFilesTargetParams) GetOriginalSourceConfig() OriginalNetappFilesTargetParams`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *RecoverNetappToNetappFilesTargetParams) GetOriginalSourceConfigOk() (*OriginalNetappFilesTargetParams, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *RecoverNetappToNetappFilesTargetParams) SetOriginalSourceConfig(v OriginalNetappFilesTargetParams)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *RecoverNetappToNetappFilesTargetParams) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *RecoverNetappToNetappFilesTargetParams) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *RecoverNetappToNetappFilesTargetParams) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


