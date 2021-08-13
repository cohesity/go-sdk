# ReplicatedViewNameConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceViewId** | **NullableInt64** | Specifies the ID of the protected view. | 
**UseSameViewName** | Pointer to **NullableBool** | Specifies if the remote view name to be kept is same as the source view name. If this field is true, viewName field will be ignored. | [optional] 
**ViewName** | Pointer to **NullableString** | Specifies the name of the remote view. This field is only used when useSameViewName is false. If useSameViewName is true, this field is not used. | [optional] 

## Methods

### NewReplicatedViewNameConfig

`func NewReplicatedViewNameConfig(sourceViewId NullableInt64, ) *ReplicatedViewNameConfig`

NewReplicatedViewNameConfig instantiates a new ReplicatedViewNameConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicatedViewNameConfigWithDefaults

`func NewReplicatedViewNameConfigWithDefaults() *ReplicatedViewNameConfig`

NewReplicatedViewNameConfigWithDefaults instantiates a new ReplicatedViewNameConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceViewId

`func (o *ReplicatedViewNameConfig) GetSourceViewId() int64`

GetSourceViewId returns the SourceViewId field if non-nil, zero value otherwise.

### GetSourceViewIdOk

`func (o *ReplicatedViewNameConfig) GetSourceViewIdOk() (*int64, bool)`

GetSourceViewIdOk returns a tuple with the SourceViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceViewId

`func (o *ReplicatedViewNameConfig) SetSourceViewId(v int64)`

SetSourceViewId sets SourceViewId field to given value.


### SetSourceViewIdNil

`func (o *ReplicatedViewNameConfig) SetSourceViewIdNil(b bool)`

 SetSourceViewIdNil sets the value for SourceViewId to be an explicit nil

### UnsetSourceViewId
`func (o *ReplicatedViewNameConfig) UnsetSourceViewId()`

UnsetSourceViewId ensures that no value is present for SourceViewId, not even an explicit nil
### GetUseSameViewName

`func (o *ReplicatedViewNameConfig) GetUseSameViewName() bool`

GetUseSameViewName returns the UseSameViewName field if non-nil, zero value otherwise.

### GetUseSameViewNameOk

`func (o *ReplicatedViewNameConfig) GetUseSameViewNameOk() (*bool, bool)`

GetUseSameViewNameOk returns a tuple with the UseSameViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSameViewName

`func (o *ReplicatedViewNameConfig) SetUseSameViewName(v bool)`

SetUseSameViewName sets UseSameViewName field to given value.

### HasUseSameViewName

`func (o *ReplicatedViewNameConfig) HasUseSameViewName() bool`

HasUseSameViewName returns a boolean if a field has been set.

### SetUseSameViewNameNil

`func (o *ReplicatedViewNameConfig) SetUseSameViewNameNil(b bool)`

 SetUseSameViewNameNil sets the value for UseSameViewName to be an explicit nil

### UnsetUseSameViewName
`func (o *ReplicatedViewNameConfig) UnsetUseSameViewName()`

UnsetUseSameViewName ensures that no value is present for UseSameViewName, not even an explicit nil
### GetViewName

`func (o *ReplicatedViewNameConfig) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *ReplicatedViewNameConfig) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *ReplicatedViewNameConfig) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *ReplicatedViewNameConfig) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *ReplicatedViewNameConfig) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *ReplicatedViewNameConfig) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


