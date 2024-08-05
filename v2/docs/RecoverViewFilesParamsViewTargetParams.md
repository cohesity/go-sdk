# RecoverViewFilesParamsViewTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewViewConfig** | Pointer to [**NullableRecoverViewToViewFilesTargetParamsNewViewConfig**](RecoverViewToViewFilesTargetParamsNewViewConfig.md) |  | [optional] 
**OriginalViewConfig** | Pointer to [**NullableRecoverViewToViewFilesTargetParamsOriginalViewConfig**](RecoverViewToViewFilesTargetParamsOriginalViewConfig.md) |  | [optional] 
**RecoverToNewView** | **bool** | Specifies the parameter whether the recovery should be performed to a new or the original View target. | 
**ViewId** | Pointer to **NullableInt64** | Specifies the ID of the view. | [optional] 
**ViewName** | Pointer to **NullableString** | Specifies the name of the new view that&#39;s the target for recovery. | [optional] 

## Methods

### NewRecoverViewFilesParamsViewTargetParams

`func NewRecoverViewFilesParamsViewTargetParams(recoverToNewView bool, ) *RecoverViewFilesParamsViewTargetParams`

NewRecoverViewFilesParamsViewTargetParams instantiates a new RecoverViewFilesParamsViewTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverViewFilesParamsViewTargetParamsWithDefaults

`func NewRecoverViewFilesParamsViewTargetParamsWithDefaults() *RecoverViewFilesParamsViewTargetParams`

NewRecoverViewFilesParamsViewTargetParamsWithDefaults instantiates a new RecoverViewFilesParamsViewTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewViewConfig

`func (o *RecoverViewFilesParamsViewTargetParams) GetNewViewConfig() RecoverViewToViewFilesTargetParamsNewViewConfig`

GetNewViewConfig returns the NewViewConfig field if non-nil, zero value otherwise.

### GetNewViewConfigOk

`func (o *RecoverViewFilesParamsViewTargetParams) GetNewViewConfigOk() (*RecoverViewToViewFilesTargetParamsNewViewConfig, bool)`

GetNewViewConfigOk returns a tuple with the NewViewConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewViewConfig

`func (o *RecoverViewFilesParamsViewTargetParams) SetNewViewConfig(v RecoverViewToViewFilesTargetParamsNewViewConfig)`

SetNewViewConfig sets NewViewConfig field to given value.

### HasNewViewConfig

`func (o *RecoverViewFilesParamsViewTargetParams) HasNewViewConfig() bool`

HasNewViewConfig returns a boolean if a field has been set.

### SetNewViewConfigNil

`func (o *RecoverViewFilesParamsViewTargetParams) SetNewViewConfigNil(b bool)`

 SetNewViewConfigNil sets the value for NewViewConfig to be an explicit nil

### UnsetNewViewConfig
`func (o *RecoverViewFilesParamsViewTargetParams) UnsetNewViewConfig()`

UnsetNewViewConfig ensures that no value is present for NewViewConfig, not even an explicit nil
### GetOriginalViewConfig

`func (o *RecoverViewFilesParamsViewTargetParams) GetOriginalViewConfig() RecoverViewToViewFilesTargetParamsOriginalViewConfig`

GetOriginalViewConfig returns the OriginalViewConfig field if non-nil, zero value otherwise.

### GetOriginalViewConfigOk

`func (o *RecoverViewFilesParamsViewTargetParams) GetOriginalViewConfigOk() (*RecoverViewToViewFilesTargetParamsOriginalViewConfig, bool)`

GetOriginalViewConfigOk returns a tuple with the OriginalViewConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalViewConfig

`func (o *RecoverViewFilesParamsViewTargetParams) SetOriginalViewConfig(v RecoverViewToViewFilesTargetParamsOriginalViewConfig)`

SetOriginalViewConfig sets OriginalViewConfig field to given value.

### HasOriginalViewConfig

`func (o *RecoverViewFilesParamsViewTargetParams) HasOriginalViewConfig() bool`

HasOriginalViewConfig returns a boolean if a field has been set.

### SetOriginalViewConfigNil

`func (o *RecoverViewFilesParamsViewTargetParams) SetOriginalViewConfigNil(b bool)`

 SetOriginalViewConfigNil sets the value for OriginalViewConfig to be an explicit nil

### UnsetOriginalViewConfig
`func (o *RecoverViewFilesParamsViewTargetParams) UnsetOriginalViewConfig()`

UnsetOriginalViewConfig ensures that no value is present for OriginalViewConfig, not even an explicit nil
### GetRecoverToNewView

`func (o *RecoverViewFilesParamsViewTargetParams) GetRecoverToNewView() bool`

GetRecoverToNewView returns the RecoverToNewView field if non-nil, zero value otherwise.

### GetRecoverToNewViewOk

`func (o *RecoverViewFilesParamsViewTargetParams) GetRecoverToNewViewOk() (*bool, bool)`

GetRecoverToNewViewOk returns a tuple with the RecoverToNewView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewView

`func (o *RecoverViewFilesParamsViewTargetParams) SetRecoverToNewView(v bool)`

SetRecoverToNewView sets RecoverToNewView field to given value.


### GetViewId

`func (o *RecoverViewFilesParamsViewTargetParams) GetViewId() int64`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *RecoverViewFilesParamsViewTargetParams) GetViewIdOk() (*int64, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *RecoverViewFilesParamsViewTargetParams) SetViewId(v int64)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *RecoverViewFilesParamsViewTargetParams) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### SetViewIdNil

`func (o *RecoverViewFilesParamsViewTargetParams) SetViewIdNil(b bool)`

 SetViewIdNil sets the value for ViewId to be an explicit nil

### UnsetViewId
`func (o *RecoverViewFilesParamsViewTargetParams) UnsetViewId()`

UnsetViewId ensures that no value is present for ViewId, not even an explicit nil
### GetViewName

`func (o *RecoverViewFilesParamsViewTargetParams) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *RecoverViewFilesParamsViewTargetParams) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *RecoverViewFilesParamsViewTargetParams) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *RecoverViewFilesParamsViewTargetParams) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *RecoverViewFilesParamsViewTargetParams) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *RecoverViewFilesParamsViewTargetParams) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


