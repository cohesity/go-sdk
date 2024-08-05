# GetGaiaIndexingStatusParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobHandle** | Pointer to **NullableString** | Job handle for this request. | [optional] 

## Methods

### NewGetGaiaIndexingStatusParams

`func NewGetGaiaIndexingStatusParams() *GetGaiaIndexingStatusParams`

NewGetGaiaIndexingStatusParams instantiates a new GetGaiaIndexingStatusParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGaiaIndexingStatusParamsWithDefaults

`func NewGetGaiaIndexingStatusParamsWithDefaults() *GetGaiaIndexingStatusParams`

NewGetGaiaIndexingStatusParamsWithDefaults instantiates a new GetGaiaIndexingStatusParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobHandle

`func (o *GetGaiaIndexingStatusParams) GetJobHandle() string`

GetJobHandle returns the JobHandle field if non-nil, zero value otherwise.

### GetJobHandleOk

`func (o *GetGaiaIndexingStatusParams) GetJobHandleOk() (*string, bool)`

GetJobHandleOk returns a tuple with the JobHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobHandle

`func (o *GetGaiaIndexingStatusParams) SetJobHandle(v string)`

SetJobHandle sets JobHandle field to given value.

### HasJobHandle

`func (o *GetGaiaIndexingStatusParams) HasJobHandle() bool`

HasJobHandle returns a boolean if a field has been set.

### SetJobHandleNil

`func (o *GetGaiaIndexingStatusParams) SetJobHandleNil(b bool)`

 SetJobHandleNil sets the value for JobHandle to be an explicit nil

### UnsetJobHandle
`func (o *GetGaiaIndexingStatusParams) UnsetJobHandle()`

UnsetJobHandle ensures that no value is present for JobHandle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


