# CancelGaiaIndexingParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobHandle** | Pointer to **NullableString** | Job handle for this request. | [optional] 

## Methods

### NewCancelGaiaIndexingParams

`func NewCancelGaiaIndexingParams() *CancelGaiaIndexingParams`

NewCancelGaiaIndexingParams instantiates a new CancelGaiaIndexingParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelGaiaIndexingParamsWithDefaults

`func NewCancelGaiaIndexingParamsWithDefaults() *CancelGaiaIndexingParams`

NewCancelGaiaIndexingParamsWithDefaults instantiates a new CancelGaiaIndexingParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobHandle

`func (o *CancelGaiaIndexingParams) GetJobHandle() string`

GetJobHandle returns the JobHandle field if non-nil, zero value otherwise.

### GetJobHandleOk

`func (o *CancelGaiaIndexingParams) GetJobHandleOk() (*string, bool)`

GetJobHandleOk returns a tuple with the JobHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobHandle

`func (o *CancelGaiaIndexingParams) SetJobHandle(v string)`

SetJobHandle sets JobHandle field to given value.

### HasJobHandle

`func (o *CancelGaiaIndexingParams) HasJobHandle() bool`

HasJobHandle returns a boolean if a field has been set.

### SetJobHandleNil

`func (o *CancelGaiaIndexingParams) SetJobHandleNil(b bool)`

 SetJobHandleNil sets the value for JobHandle to be an explicit nil

### UnsetJobHandle
`func (o *CancelGaiaIndexingParams) UnsetJobHandle()`

UnsetJobHandle ensures that no value is present for JobHandle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


