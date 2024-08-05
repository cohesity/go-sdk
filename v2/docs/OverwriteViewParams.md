# OverwriteViewParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceViewId** | **NullableInt64** | Specifies the source View id. Target View will be overwritten by the source View. | 

## Methods

### NewOverwriteViewParams

`func NewOverwriteViewParams(sourceViewId NullableInt64, ) *OverwriteViewParams`

NewOverwriteViewParams instantiates a new OverwriteViewParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverwriteViewParamsWithDefaults

`func NewOverwriteViewParamsWithDefaults() *OverwriteViewParams`

NewOverwriteViewParamsWithDefaults instantiates a new OverwriteViewParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceViewId

`func (o *OverwriteViewParams) GetSourceViewId() int64`

GetSourceViewId returns the SourceViewId field if non-nil, zero value otherwise.

### GetSourceViewIdOk

`func (o *OverwriteViewParams) GetSourceViewIdOk() (*int64, bool)`

GetSourceViewIdOk returns a tuple with the SourceViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceViewId

`func (o *OverwriteViewParams) SetSourceViewId(v int64)`

SetSourceViewId sets SourceViewId field to given value.


### SetSourceViewIdNil

`func (o *OverwriteViewParams) SetSourceViewIdNil(b bool)`

 SetSourceViewIdNil sets the value for SourceViewId to be an explicit nil

### UnsetSourceViewId
`func (o *OverwriteViewParams) UnsetSourceViewId()`

UnsetSourceViewId ensures that no value is present for SourceViewId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


