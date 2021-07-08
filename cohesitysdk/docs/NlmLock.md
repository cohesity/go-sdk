# NlmLock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **NullableString** | Specifies the client ID | [optional] 
**LockRanges** | Pointer to [**[]LockRange**](LockRange.md) |  | [optional] 

## Methods

### NewNlmLock

`func NewNlmLock() *NlmLock`

NewNlmLock instantiates a new NlmLock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNlmLockWithDefaults

`func NewNlmLockWithDefaults() *NlmLock`

NewNlmLockWithDefaults instantiates a new NlmLock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *NlmLock) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *NlmLock) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *NlmLock) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *NlmLock) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *NlmLock) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *NlmLock) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetLockRanges

`func (o *NlmLock) GetLockRanges() []LockRange`

GetLockRanges returns the LockRanges field if non-nil, zero value otherwise.

### GetLockRangesOk

`func (o *NlmLock) GetLockRangesOk() (*[]LockRange, bool)`

GetLockRangesOk returns a tuple with the LockRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockRanges

`func (o *NlmLock) SetLockRanges(v []LockRange)`

SetLockRanges sets LockRanges field to given value.

### HasLockRanges

`func (o *NlmLock) HasLockRanges() bool`

HasLockRanges returns a boolean if a field has been set.

### SetLockRangesNil

`func (o *NlmLock) SetLockRangesNil(b bool)`

 SetLockRangesNil sets the value for LockRanges to be an explicit nil

### UnsetLockRanges
`func (o *NlmLock) UnsetLockRanges()`

UnsetLockRanges ensures that no value is present for LockRanges, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


