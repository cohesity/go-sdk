# FailedRunDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessage** | Pointer to **NullableString** | Specifies the error mesage for failed run. | [optional] 
**RunId** | Pointer to **NullableString** | Specifies the id of the failed run. | [optional] 

## Methods

### NewFailedRunDetails

`func NewFailedRunDetails() *FailedRunDetails`

NewFailedRunDetails instantiates a new FailedRunDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailedRunDetailsWithDefaults

`func NewFailedRunDetailsWithDefaults() *FailedRunDetails`

NewFailedRunDetailsWithDefaults instantiates a new FailedRunDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessage

`func (o *FailedRunDetails) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *FailedRunDetails) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *FailedRunDetails) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *FailedRunDetails) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *FailedRunDetails) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *FailedRunDetails) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetRunId

`func (o *FailedRunDetails) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *FailedRunDetails) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *FailedRunDetails) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *FailedRunDetails) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### SetRunIdNil

`func (o *FailedRunDetails) SetRunIdNil(b bool)`

 SetRunIdNil sets the value for RunId to be an explicit nil

### UnsetRunId
`func (o *FailedRunDetails) UnsetRunId()`

UnsetRunId ensures that no value is present for RunId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


