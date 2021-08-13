# RecoverFlashbladeToFlashbladeFilesTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or the original Flashblade target. | 
**NewSourceConfig** | Pointer to [**NullableRecoverOtherNasToFlashbladeFilesTargetParams**](RecoverOtherNasToFlashbladeFilesTargetParams.md) | Specifies the new destination Source configuration parameters where the files will be recovered. This is mandatory if recoverToNewSource is set to true. | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableOriginalFlashbladeFilesTargetParams**](OriginalFlashbladeFilesTargetParams.md) | Specifies the Source configuration if files are being recovered to original Source. If not specified, all the configuration parameters will be retained. | [optional] 

## Methods

### NewRecoverFlashbladeToFlashbladeFilesTargetParams

`func NewRecoverFlashbladeToFlashbladeFilesTargetParams(recoverToNewSource bool, ) *RecoverFlashbladeToFlashbladeFilesTargetParams`

NewRecoverFlashbladeToFlashbladeFilesTargetParams instantiates a new RecoverFlashbladeToFlashbladeFilesTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverFlashbladeToFlashbladeFilesTargetParamsWithDefaults

`func NewRecoverFlashbladeToFlashbladeFilesTargetParamsWithDefaults() *RecoverFlashbladeToFlashbladeFilesTargetParams`

NewRecoverFlashbladeToFlashbladeFilesTargetParamsWithDefaults instantiates a new RecoverFlashbladeToFlashbladeFilesTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoverToNewSource

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParams) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParams) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParams) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.


### GetNewSourceConfig

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParams) GetNewSourceConfig() RecoverOtherNasToFlashbladeFilesTargetParams`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParams) GetNewSourceConfigOk() (*RecoverOtherNasToFlashbladeFilesTargetParams, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParams) SetNewSourceConfig(v RecoverOtherNasToFlashbladeFilesTargetParams)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParams) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParams) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *RecoverFlashbladeToFlashbladeFilesTargetParams) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParams) GetOriginalSourceConfig() OriginalFlashbladeFilesTargetParams`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParams) GetOriginalSourceConfigOk() (*OriginalFlashbladeFilesTargetParams, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParams) SetOriginalSourceConfig(v OriginalFlashbladeFilesTargetParams)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParams) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParams) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *RecoverFlashbladeToFlashbladeFilesTargetParams) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


