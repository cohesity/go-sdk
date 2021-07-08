# GetViewsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastResult** | Pointer to **NullableBool** | If false, more Views are available to return. If the number of Views to return exceeds the number of Views specified in maxCount (default of 1000) of the original GET /public/views request, the first set of Views are returned and this field returns false. To get the next set of Views, in the next GET /public/views request send the last id from the previous viewList. | [optional] 
**Views** | Pointer to [**[]View**](View.md) | Array of Views.  Specifies the list of Views returned in this response. The list is sorted by decreasing View ids. | [optional] 

## Methods

### NewGetViewsResult

`func NewGetViewsResult() *GetViewsResult`

NewGetViewsResult instantiates a new GetViewsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetViewsResultWithDefaults

`func NewGetViewsResultWithDefaults() *GetViewsResult`

NewGetViewsResultWithDefaults instantiates a new GetViewsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastResult

`func (o *GetViewsResult) GetLastResult() bool`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *GetViewsResult) GetLastResultOk() (*bool, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *GetViewsResult) SetLastResult(v bool)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *GetViewsResult) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### SetLastResultNil

`func (o *GetViewsResult) SetLastResultNil(b bool)`

 SetLastResultNil sets the value for LastResult to be an explicit nil

### UnsetLastResult
`func (o *GetViewsResult) UnsetLastResult()`

UnsetLastResult ensures that no value is present for LastResult, not even an explicit nil
### GetViews

`func (o *GetViewsResult) GetViews() []View`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *GetViewsResult) GetViewsOk() (*[]View, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *GetViewsResult) SetViews(v []View)`

SetViews sets Views field to given value.

### HasViews

`func (o *GetViewsResult) HasViews() bool`

HasViews returns a boolean if a field has been set.

### SetViewsNil

`func (o *GetViewsResult) SetViewsNil(b bool)`

 SetViewsNil sets the value for Views to be an explicit nil

### UnsetViews
`func (o *GetViewsResult) UnsetViews()`

UnsetViews ensures that no value is present for Views, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


