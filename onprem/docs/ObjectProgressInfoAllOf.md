# ObjectProgressInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedAttempts** | Pointer to [**[]ProgressTaskInfo**](ProgressTaskInfo.md) | Specifies progress for failed attempts of this object. | [optional] 

## Methods

### NewObjectProgressInfoAllOf

`func NewObjectProgressInfoAllOf() *ObjectProgressInfoAllOf`

NewObjectProgressInfoAllOf instantiates a new ObjectProgressInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectProgressInfoAllOfWithDefaults

`func NewObjectProgressInfoAllOfWithDefaults() *ObjectProgressInfoAllOf`

NewObjectProgressInfoAllOfWithDefaults instantiates a new ObjectProgressInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedAttempts

`func (o *ObjectProgressInfoAllOf) GetFailedAttempts() []ProgressTaskInfo`

GetFailedAttempts returns the FailedAttempts field if non-nil, zero value otherwise.

### GetFailedAttemptsOk

`func (o *ObjectProgressInfoAllOf) GetFailedAttemptsOk() (*[]ProgressTaskInfo, bool)`

GetFailedAttemptsOk returns a tuple with the FailedAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAttempts

`func (o *ObjectProgressInfoAllOf) SetFailedAttempts(v []ProgressTaskInfo)`

SetFailedAttempts sets FailedAttempts field to given value.

### HasFailedAttempts

`func (o *ObjectProgressInfoAllOf) HasFailedAttempts() bool`

HasFailedAttempts returns a boolean if a field has been set.

### SetFailedAttemptsNil

`func (o *ObjectProgressInfoAllOf) SetFailedAttemptsNil(b bool)`

 SetFailedAttemptsNil sets the value for FailedAttempts to be an explicit nil

### UnsetFailedAttempts
`func (o *ObjectProgressInfoAllOf) UnsetFailedAttempts()`

UnsetFailedAttempts ensures that no value is present for FailedAttempts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


