# AdObjectRestoreInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeRestoreInfo** | Pointer to [**[]AttributeRestoreInformation**](AttributeRestoreInformation.md) | Specifies the list of attributes of the AD object whose restore failed. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Specifies the error message while restoring the AD Object. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the AD object. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the restore of the AD object specified as a Unix epoch Timestamp(in microseconds). | [optional] 
**TimeTakenMsecs** | Pointer to **NullableInt32** | Specifies the time taken for restore of AD Object and its attributes in milliseconds. | [optional] 

## Methods

### NewAdObjectRestoreInformation

`func NewAdObjectRestoreInformation() *AdObjectRestoreInformation`

NewAdObjectRestoreInformation instantiates a new AdObjectRestoreInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdObjectRestoreInformationWithDefaults

`func NewAdObjectRestoreInformationWithDefaults() *AdObjectRestoreInformation`

NewAdObjectRestoreInformationWithDefaults instantiates a new AdObjectRestoreInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeRestoreInfo

`func (o *AdObjectRestoreInformation) GetAttributeRestoreInfo() []AttributeRestoreInformation`

GetAttributeRestoreInfo returns the AttributeRestoreInfo field if non-nil, zero value otherwise.

### GetAttributeRestoreInfoOk

`func (o *AdObjectRestoreInformation) GetAttributeRestoreInfoOk() (*[]AttributeRestoreInformation, bool)`

GetAttributeRestoreInfoOk returns a tuple with the AttributeRestoreInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeRestoreInfo

`func (o *AdObjectRestoreInformation) SetAttributeRestoreInfo(v []AttributeRestoreInformation)`

SetAttributeRestoreInfo sets AttributeRestoreInfo field to given value.

### HasAttributeRestoreInfo

`func (o *AdObjectRestoreInformation) HasAttributeRestoreInfo() bool`

HasAttributeRestoreInfo returns a boolean if a field has been set.

### SetAttributeRestoreInfoNil

`func (o *AdObjectRestoreInformation) SetAttributeRestoreInfoNil(b bool)`

 SetAttributeRestoreInfoNil sets the value for AttributeRestoreInfo to be an explicit nil

### UnsetAttributeRestoreInfo
`func (o *AdObjectRestoreInformation) UnsetAttributeRestoreInfo()`

UnsetAttributeRestoreInfo ensures that no value is present for AttributeRestoreInfo, not even an explicit nil
### GetErrorMessage

`func (o *AdObjectRestoreInformation) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AdObjectRestoreInformation) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AdObjectRestoreInformation) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AdObjectRestoreInformation) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *AdObjectRestoreInformation) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *AdObjectRestoreInformation) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetName

`func (o *AdObjectRestoreInformation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdObjectRestoreInformation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdObjectRestoreInformation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdObjectRestoreInformation) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AdObjectRestoreInformation) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AdObjectRestoreInformation) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStartTimeUsecs

`func (o *AdObjectRestoreInformation) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *AdObjectRestoreInformation) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *AdObjectRestoreInformation) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *AdObjectRestoreInformation) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *AdObjectRestoreInformation) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *AdObjectRestoreInformation) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetTimeTakenMsecs

`func (o *AdObjectRestoreInformation) GetTimeTakenMsecs() int32`

GetTimeTakenMsecs returns the TimeTakenMsecs field if non-nil, zero value otherwise.

### GetTimeTakenMsecsOk

`func (o *AdObjectRestoreInformation) GetTimeTakenMsecsOk() (*int32, bool)`

GetTimeTakenMsecsOk returns a tuple with the TimeTakenMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTakenMsecs

`func (o *AdObjectRestoreInformation) SetTimeTakenMsecs(v int32)`

SetTimeTakenMsecs sets TimeTakenMsecs field to given value.

### HasTimeTakenMsecs

`func (o *AdObjectRestoreInformation) HasTimeTakenMsecs() bool`

HasTimeTakenMsecs returns a boolean if a field has been set.

### SetTimeTakenMsecsNil

`func (o *AdObjectRestoreInformation) SetTimeTakenMsecsNil(b bool)`

 SetTimeTakenMsecsNil sets the value for TimeTakenMsecs to be an explicit nil

### UnsetTimeTakenMsecs
`func (o *AdObjectRestoreInformation) UnsetTimeTakenMsecs()`

UnsetTimeTakenMsecs ensures that no value is present for TimeTakenMsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


