# RemoveDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies id of the disk. | [optional] 
**MarkedForRemoval** | Pointer to **NullableBool** | If true, Disk is marked for removal. | [optional] 
**TimestampSecs** | Pointer to **NullableInt64** | Specifies the last run time of the pre-checks execution in Unix epoch timestamp (in seconds). | [optional] 
**ValidationChecks** | Pointer to [**[]PreCheckValidation**](PreCheckValidation.md) | Specifies the pre-check validations results. | [optional] 

## Methods

### NewRemoveDisk

`func NewRemoveDisk() *RemoveDisk`

NewRemoveDisk instantiates a new RemoveDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveDiskWithDefaults

`func NewRemoveDiskWithDefaults() *RemoveDisk`

NewRemoveDiskWithDefaults instantiates a new RemoveDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RemoveDisk) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoveDisk) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoveDisk) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RemoveDisk) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *RemoveDisk) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RemoveDisk) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMarkedForRemoval

`func (o *RemoveDisk) GetMarkedForRemoval() bool`

GetMarkedForRemoval returns the MarkedForRemoval field if non-nil, zero value otherwise.

### GetMarkedForRemovalOk

`func (o *RemoveDisk) GetMarkedForRemovalOk() (*bool, bool)`

GetMarkedForRemovalOk returns a tuple with the MarkedForRemoval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedForRemoval

`func (o *RemoveDisk) SetMarkedForRemoval(v bool)`

SetMarkedForRemoval sets MarkedForRemoval field to given value.

### HasMarkedForRemoval

`func (o *RemoveDisk) HasMarkedForRemoval() bool`

HasMarkedForRemoval returns a boolean if a field has been set.

### SetMarkedForRemovalNil

`func (o *RemoveDisk) SetMarkedForRemovalNil(b bool)`

 SetMarkedForRemovalNil sets the value for MarkedForRemoval to be an explicit nil

### UnsetMarkedForRemoval
`func (o *RemoveDisk) UnsetMarkedForRemoval()`

UnsetMarkedForRemoval ensures that no value is present for MarkedForRemoval, not even an explicit nil
### GetTimestampSecs

`func (o *RemoveDisk) GetTimestampSecs() int64`

GetTimestampSecs returns the TimestampSecs field if non-nil, zero value otherwise.

### GetTimestampSecsOk

`func (o *RemoveDisk) GetTimestampSecsOk() (*int64, bool)`

GetTimestampSecsOk returns a tuple with the TimestampSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampSecs

`func (o *RemoveDisk) SetTimestampSecs(v int64)`

SetTimestampSecs sets TimestampSecs field to given value.

### HasTimestampSecs

`func (o *RemoveDisk) HasTimestampSecs() bool`

HasTimestampSecs returns a boolean if a field has been set.

### SetTimestampSecsNil

`func (o *RemoveDisk) SetTimestampSecsNil(b bool)`

 SetTimestampSecsNil sets the value for TimestampSecs to be an explicit nil

### UnsetTimestampSecs
`func (o *RemoveDisk) UnsetTimestampSecs()`

UnsetTimestampSecs ensures that no value is present for TimestampSecs, not even an explicit nil
### GetValidationChecks

`func (o *RemoveDisk) GetValidationChecks() []PreCheckValidation`

GetValidationChecks returns the ValidationChecks field if non-nil, zero value otherwise.

### GetValidationChecksOk

`func (o *RemoveDisk) GetValidationChecksOk() (*[]PreCheckValidation, bool)`

GetValidationChecksOk returns a tuple with the ValidationChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationChecks

`func (o *RemoveDisk) SetValidationChecks(v []PreCheckValidation)`

SetValidationChecks sets ValidationChecks field to given value.

### HasValidationChecks

`func (o *RemoveDisk) HasValidationChecks() bool`

HasValidationChecks returns a boolean if a field has been set.

### SetValidationChecksNil

`func (o *RemoveDisk) SetValidationChecksNil(b bool)`

 SetValidationChecksNil sets the value for ValidationChecks to be an explicit nil

### UnsetValidationChecks
`func (o *RemoveDisk) UnsetValidationChecks()`

UnsetValidationChecks ensures that no value is present for ValidationChecks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


