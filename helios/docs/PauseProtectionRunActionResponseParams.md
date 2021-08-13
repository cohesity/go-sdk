# PauseProtectionRunActionResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | Pointer to **NullableString** | Specifies a unique run id of the Protection Group run. | [optional] 
**Error** | Pointer to **NullableString** | Specifies an error occured when perfroming pause of a protection run. | [optional] 

## Methods

### NewPauseProtectionRunActionResponseParams

`func NewPauseProtectionRunActionResponseParams() *PauseProtectionRunActionResponseParams`

NewPauseProtectionRunActionResponseParams instantiates a new PauseProtectionRunActionResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPauseProtectionRunActionResponseParamsWithDefaults

`func NewPauseProtectionRunActionResponseParamsWithDefaults() *PauseProtectionRunActionResponseParams`

NewPauseProtectionRunActionResponseParamsWithDefaults instantiates a new PauseProtectionRunActionResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *PauseProtectionRunActionResponseParams) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *PauseProtectionRunActionResponseParams) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *PauseProtectionRunActionResponseParams) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *PauseProtectionRunActionResponseParams) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### SetRunIdNil

`func (o *PauseProtectionRunActionResponseParams) SetRunIdNil(b bool)`

 SetRunIdNil sets the value for RunId to be an explicit nil

### UnsetRunId
`func (o *PauseProtectionRunActionResponseParams) UnsetRunId()`

UnsetRunId ensures that no value is present for RunId, not even an explicit nil
### GetError

`func (o *PauseProtectionRunActionResponseParams) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PauseProtectionRunActionResponseParams) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PauseProtectionRunActionResponseParams) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *PauseProtectionRunActionResponseParams) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *PauseProtectionRunActionResponseParams) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *PauseProtectionRunActionResponseParams) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


