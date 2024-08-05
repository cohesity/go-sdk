# Failover

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the failover complete time in micro seconds. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Specifies the error details if failover status is &#39;Failed&#39;. | [optional] 
**Id** | Pointer to **NullableString** | Specifies the failover id. | [optional] 
**Replications** | Pointer to [**[]FailoverReplication**](FailoverReplication.md) | Specifies a list of replications in this failover. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the failover start time in micro seconds. | [optional] 
**Status** | Pointer to **NullableString** | Specifies the failover status. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the failover type. | [optional] 

## Methods

### NewFailover

`func NewFailover() *Failover`

NewFailover instantiates a new Failover object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailoverWithDefaults

`func NewFailoverWithDefaults() *Failover`

NewFailoverWithDefaults instantiates a new Failover object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTimeUsecs

`func (o *Failover) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *Failover) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *Failover) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *Failover) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *Failover) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *Failover) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetErrorMessage

`func (o *Failover) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Failover) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Failover) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Failover) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *Failover) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *Failover) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetId

`func (o *Failover) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Failover) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Failover) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Failover) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Failover) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Failover) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetReplications

`func (o *Failover) GetReplications() []FailoverReplication`

GetReplications returns the Replications field if non-nil, zero value otherwise.

### GetReplicationsOk

`func (o *Failover) GetReplicationsOk() (*[]FailoverReplication, bool)`

GetReplicationsOk returns a tuple with the Replications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplications

`func (o *Failover) SetReplications(v []FailoverReplication)`

SetReplications sets Replications field to given value.

### HasReplications

`func (o *Failover) HasReplications() bool`

HasReplications returns a boolean if a field has been set.

### SetReplicationsNil

`func (o *Failover) SetReplicationsNil(b bool)`

 SetReplicationsNil sets the value for Replications to be an explicit nil

### UnsetReplications
`func (o *Failover) UnsetReplications()`

UnsetReplications ensures that no value is present for Replications, not even an explicit nil
### GetStartTimeUsecs

`func (o *Failover) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *Failover) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *Failover) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *Failover) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *Failover) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *Failover) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStatus

`func (o *Failover) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Failover) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Failover) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Failover) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Failover) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Failover) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetType

`func (o *Failover) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Failover) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Failover) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Failover) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Failover) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Failover) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


