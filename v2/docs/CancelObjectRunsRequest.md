# CancelObjectRunsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectRuns** | Pointer to [**[]CancelObjectRunsParams**](CancelObjectRunsParams.md) | Specifies a list of object runs to cancel. | [optional] 

## Methods

### NewCancelObjectRunsRequest

`func NewCancelObjectRunsRequest() *CancelObjectRunsRequest`

NewCancelObjectRunsRequest instantiates a new CancelObjectRunsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelObjectRunsRequestWithDefaults

`func NewCancelObjectRunsRequestWithDefaults() *CancelObjectRunsRequest`

NewCancelObjectRunsRequestWithDefaults instantiates a new CancelObjectRunsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectRuns

`func (o *CancelObjectRunsRequest) GetObjectRuns() []CancelObjectRunsParams`

GetObjectRuns returns the ObjectRuns field if non-nil, zero value otherwise.

### GetObjectRunsOk

`func (o *CancelObjectRunsRequest) GetObjectRunsOk() (*[]CancelObjectRunsParams, bool)`

GetObjectRunsOk returns a tuple with the ObjectRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectRuns

`func (o *CancelObjectRunsRequest) SetObjectRuns(v []CancelObjectRunsParams)`

SetObjectRuns sets ObjectRuns field to given value.

### HasObjectRuns

`func (o *CancelObjectRunsRequest) HasObjectRuns() bool`

HasObjectRuns returns a boolean if a field has been set.

### SetObjectRunsNil

`func (o *CancelObjectRunsRequest) SetObjectRunsNil(b bool)`

 SetObjectRunsNil sets the value for ObjectRuns to be an explicit nil

### UnsetObjectRuns
`func (o *CancelObjectRunsRequest) UnsetObjectRuns()`

UnsetObjectRuns ensures that no value is present for ObjectRuns, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


