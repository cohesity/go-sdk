# ProtectedObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to [**UniversalId**](UniversalId.md) |  | [optional] 
**ProtectionFauilureReason** | Pointer to **NullableString** | If protection fails then specifies why the protection failed on this object. | [optional] 
**ProtectionSourceId** | Pointer to **NullableInt64** | Specifies the id of the Protection Source. | [optional] 

## Methods

### NewProtectedObject

`func NewProtectedObject() *ProtectedObject`

NewProtectedObject instantiates a new ProtectedObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedObjectWithDefaults

`func NewProtectedObjectWithDefaults() *ProtectedObject`

NewProtectedObjectWithDefaults instantiates a new ProtectedObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *ProtectedObject) GetJobId() UniversalId`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ProtectedObject) GetJobIdOk() (*UniversalId, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ProtectedObject) SetJobId(v UniversalId)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *ProtectedObject) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetProtectionFauilureReason

`func (o *ProtectedObject) GetProtectionFauilureReason() string`

GetProtectionFauilureReason returns the ProtectionFauilureReason field if non-nil, zero value otherwise.

### GetProtectionFauilureReasonOk

`func (o *ProtectedObject) GetProtectionFauilureReasonOk() (*string, bool)`

GetProtectionFauilureReasonOk returns a tuple with the ProtectionFauilureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionFauilureReason

`func (o *ProtectedObject) SetProtectionFauilureReason(v string)`

SetProtectionFauilureReason sets ProtectionFauilureReason field to given value.

### HasProtectionFauilureReason

`func (o *ProtectedObject) HasProtectionFauilureReason() bool`

HasProtectionFauilureReason returns a boolean if a field has been set.

### SetProtectionFauilureReasonNil

`func (o *ProtectedObject) SetProtectionFauilureReasonNil(b bool)`

 SetProtectionFauilureReasonNil sets the value for ProtectionFauilureReason to be an explicit nil

### UnsetProtectionFauilureReason
`func (o *ProtectedObject) UnsetProtectionFauilureReason()`

UnsetProtectionFauilureReason ensures that no value is present for ProtectionFauilureReason, not even an explicit nil
### GetProtectionSourceId

`func (o *ProtectedObject) GetProtectionSourceId() int64`

GetProtectionSourceId returns the ProtectionSourceId field if non-nil, zero value otherwise.

### GetProtectionSourceIdOk

`func (o *ProtectedObject) GetProtectionSourceIdOk() (*int64, bool)`

GetProtectionSourceIdOk returns a tuple with the ProtectionSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSourceId

`func (o *ProtectedObject) SetProtectionSourceId(v int64)`

SetProtectionSourceId sets ProtectionSourceId field to given value.

### HasProtectionSourceId

`func (o *ProtectedObject) HasProtectionSourceId() bool`

HasProtectionSourceId returns a boolean if a field has been set.

### SetProtectionSourceIdNil

`func (o *ProtectedObject) SetProtectionSourceIdNil(b bool)`

 SetProtectionSourceIdNil sets the value for ProtectionSourceId to be an explicit nil

### UnsetProtectionSourceId
`func (o *ProtectedObject) UnsetProtectionSourceId()`

UnsetProtectionSourceId ensures that no value is present for ProtectionSourceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


