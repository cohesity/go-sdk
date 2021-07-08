# ClusterCreationProgressResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionPercentage** | Pointer to **NullableInt32** | Specifies an approximate completion percentage for the Cluster creation process. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Specifies a description of an error if any error was encountered during Cluster creation. | [optional] 
**Events** | Pointer to **[]string** | Specifies a list of events that took place during Cluster creation. | [optional] 
**InProgress** | Pointer to **NullableBool** | Specifies whether or not the Cluster is still in the process of being created. Once the creation process is complete, this will be set to false and then, shortly afterward, all Cluster services will restart. The Cluster will be unreachable for about a minute while the services are being restarted. | [optional] 
**Message** | Pointer to **NullableString** | Specifies an optional message describing the current state of the creation progress operation. | [optional] 
**SecondsRemaining** | Pointer to **NullableInt64** | Specifies an estimated number of seconds until the Cluster creation process is complete. | [optional] 
**WarningsFound** | Pointer to **NullableBool** | Specifies whether or not any warnings were encountered during Cluster creation. | [optional] 

## Methods

### NewClusterCreationProgressResult

`func NewClusterCreationProgressResult() *ClusterCreationProgressResult`

NewClusterCreationProgressResult instantiates a new ClusterCreationProgressResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCreationProgressResultWithDefaults

`func NewClusterCreationProgressResultWithDefaults() *ClusterCreationProgressResult`

NewClusterCreationProgressResultWithDefaults instantiates a new ClusterCreationProgressResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionPercentage

`func (o *ClusterCreationProgressResult) GetCompletionPercentage() int32`

GetCompletionPercentage returns the CompletionPercentage field if non-nil, zero value otherwise.

### GetCompletionPercentageOk

`func (o *ClusterCreationProgressResult) GetCompletionPercentageOk() (*int32, bool)`

GetCompletionPercentageOk returns a tuple with the CompletionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionPercentage

`func (o *ClusterCreationProgressResult) SetCompletionPercentage(v int32)`

SetCompletionPercentage sets CompletionPercentage field to given value.

### HasCompletionPercentage

`func (o *ClusterCreationProgressResult) HasCompletionPercentage() bool`

HasCompletionPercentage returns a boolean if a field has been set.

### SetCompletionPercentageNil

`func (o *ClusterCreationProgressResult) SetCompletionPercentageNil(b bool)`

 SetCompletionPercentageNil sets the value for CompletionPercentage to be an explicit nil

### UnsetCompletionPercentage
`func (o *ClusterCreationProgressResult) UnsetCompletionPercentage()`

UnsetCompletionPercentage ensures that no value is present for CompletionPercentage, not even an explicit nil
### GetErrorMessage

`func (o *ClusterCreationProgressResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ClusterCreationProgressResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ClusterCreationProgressResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ClusterCreationProgressResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ClusterCreationProgressResult) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ClusterCreationProgressResult) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetEvents

`func (o *ClusterCreationProgressResult) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ClusterCreationProgressResult) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ClusterCreationProgressResult) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ClusterCreationProgressResult) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *ClusterCreationProgressResult) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *ClusterCreationProgressResult) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetInProgress

`func (o *ClusterCreationProgressResult) GetInProgress() bool`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *ClusterCreationProgressResult) GetInProgressOk() (*bool, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *ClusterCreationProgressResult) SetInProgress(v bool)`

SetInProgress sets InProgress field to given value.

### HasInProgress

`func (o *ClusterCreationProgressResult) HasInProgress() bool`

HasInProgress returns a boolean if a field has been set.

### SetInProgressNil

`func (o *ClusterCreationProgressResult) SetInProgressNil(b bool)`

 SetInProgressNil sets the value for InProgress to be an explicit nil

### UnsetInProgress
`func (o *ClusterCreationProgressResult) UnsetInProgress()`

UnsetInProgress ensures that no value is present for InProgress, not even an explicit nil
### GetMessage

`func (o *ClusterCreationProgressResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterCreationProgressResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterCreationProgressResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ClusterCreationProgressResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ClusterCreationProgressResult) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ClusterCreationProgressResult) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetSecondsRemaining

`func (o *ClusterCreationProgressResult) GetSecondsRemaining() int64`

GetSecondsRemaining returns the SecondsRemaining field if non-nil, zero value otherwise.

### GetSecondsRemainingOk

`func (o *ClusterCreationProgressResult) GetSecondsRemainingOk() (*int64, bool)`

GetSecondsRemainingOk returns a tuple with the SecondsRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsRemaining

`func (o *ClusterCreationProgressResult) SetSecondsRemaining(v int64)`

SetSecondsRemaining sets SecondsRemaining field to given value.

### HasSecondsRemaining

`func (o *ClusterCreationProgressResult) HasSecondsRemaining() bool`

HasSecondsRemaining returns a boolean if a field has been set.

### SetSecondsRemainingNil

`func (o *ClusterCreationProgressResult) SetSecondsRemainingNil(b bool)`

 SetSecondsRemainingNil sets the value for SecondsRemaining to be an explicit nil

### UnsetSecondsRemaining
`func (o *ClusterCreationProgressResult) UnsetSecondsRemaining()`

UnsetSecondsRemaining ensures that no value is present for SecondsRemaining, not even an explicit nil
### GetWarningsFound

`func (o *ClusterCreationProgressResult) GetWarningsFound() bool`

GetWarningsFound returns the WarningsFound field if non-nil, zero value otherwise.

### GetWarningsFoundOk

`func (o *ClusterCreationProgressResult) GetWarningsFoundOk() (*bool, bool)`

GetWarningsFoundOk returns a tuple with the WarningsFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningsFound

`func (o *ClusterCreationProgressResult) SetWarningsFound(v bool)`

SetWarningsFound sets WarningsFound field to given value.

### HasWarningsFound

`func (o *ClusterCreationProgressResult) HasWarningsFound() bool`

HasWarningsFound returns a boolean if a field has been set.

### SetWarningsFoundNil

`func (o *ClusterCreationProgressResult) SetWarningsFoundNil(b bool)`

 SetWarningsFoundNil sets the value for WarningsFound to be an explicit nil

### UnsetWarningsFound
`func (o *ClusterCreationProgressResult) UnsetWarningsFound()`

UnsetWarningsFound ensures that no value is present for WarningsFound, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


