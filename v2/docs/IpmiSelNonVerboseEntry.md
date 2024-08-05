# IpmiSelNonVerboseEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordAssertionType** | Pointer to **NullableString** | Specifies whether the event is asserted or not. This is only returned in case of verbose &#x3D; false. | [optional] 
**RecordDate** | Pointer to **NullableString** | Specifies the date on which the record is added to SEL. This is only returned in case of verbose &#x3D; false. | [optional] 
**RecordDescription** | Pointer to **NullableString** | Specifies a short description corresponding to the sensor event for which record is added to SEL. | [optional] 
**RecordEvent** | Pointer to **NullableString** | Provides a short description related to sensor action. This is only returned in case of verbose &#x3D; false. | [optional] 
**RecordId** | Pointer to **NullableString** | Specifies the ID corresponding to record in SEL(System Event Log) for given node. | [optional] 
**RecordTime** | Pointer to **NullableString** | Specifies the time at which the record is added to SEL. This is only returned in case of verbose &#x3D; false. | [optional] 

## Methods

### NewIpmiSelNonVerboseEntry

`func NewIpmiSelNonVerboseEntry() *IpmiSelNonVerboseEntry`

NewIpmiSelNonVerboseEntry instantiates a new IpmiSelNonVerboseEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpmiSelNonVerboseEntryWithDefaults

`func NewIpmiSelNonVerboseEntryWithDefaults() *IpmiSelNonVerboseEntry`

NewIpmiSelNonVerboseEntryWithDefaults instantiates a new IpmiSelNonVerboseEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordAssertionType

`func (o *IpmiSelNonVerboseEntry) GetRecordAssertionType() string`

GetRecordAssertionType returns the RecordAssertionType field if non-nil, zero value otherwise.

### GetRecordAssertionTypeOk

`func (o *IpmiSelNonVerboseEntry) GetRecordAssertionTypeOk() (*string, bool)`

GetRecordAssertionTypeOk returns a tuple with the RecordAssertionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordAssertionType

`func (o *IpmiSelNonVerboseEntry) SetRecordAssertionType(v string)`

SetRecordAssertionType sets RecordAssertionType field to given value.

### HasRecordAssertionType

`func (o *IpmiSelNonVerboseEntry) HasRecordAssertionType() bool`

HasRecordAssertionType returns a boolean if a field has been set.

### SetRecordAssertionTypeNil

`func (o *IpmiSelNonVerboseEntry) SetRecordAssertionTypeNil(b bool)`

 SetRecordAssertionTypeNil sets the value for RecordAssertionType to be an explicit nil

### UnsetRecordAssertionType
`func (o *IpmiSelNonVerboseEntry) UnsetRecordAssertionType()`

UnsetRecordAssertionType ensures that no value is present for RecordAssertionType, not even an explicit nil
### GetRecordDate

`func (o *IpmiSelNonVerboseEntry) GetRecordDate() string`

GetRecordDate returns the RecordDate field if non-nil, zero value otherwise.

### GetRecordDateOk

`func (o *IpmiSelNonVerboseEntry) GetRecordDateOk() (*string, bool)`

GetRecordDateOk returns a tuple with the RecordDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordDate

`func (o *IpmiSelNonVerboseEntry) SetRecordDate(v string)`

SetRecordDate sets RecordDate field to given value.

### HasRecordDate

`func (o *IpmiSelNonVerboseEntry) HasRecordDate() bool`

HasRecordDate returns a boolean if a field has been set.

### SetRecordDateNil

`func (o *IpmiSelNonVerboseEntry) SetRecordDateNil(b bool)`

 SetRecordDateNil sets the value for RecordDate to be an explicit nil

### UnsetRecordDate
`func (o *IpmiSelNonVerboseEntry) UnsetRecordDate()`

UnsetRecordDate ensures that no value is present for RecordDate, not even an explicit nil
### GetRecordDescription

`func (o *IpmiSelNonVerboseEntry) GetRecordDescription() string`

GetRecordDescription returns the RecordDescription field if non-nil, zero value otherwise.

### GetRecordDescriptionOk

`func (o *IpmiSelNonVerboseEntry) GetRecordDescriptionOk() (*string, bool)`

GetRecordDescriptionOk returns a tuple with the RecordDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordDescription

`func (o *IpmiSelNonVerboseEntry) SetRecordDescription(v string)`

SetRecordDescription sets RecordDescription field to given value.

### HasRecordDescription

`func (o *IpmiSelNonVerboseEntry) HasRecordDescription() bool`

HasRecordDescription returns a boolean if a field has been set.

### SetRecordDescriptionNil

`func (o *IpmiSelNonVerboseEntry) SetRecordDescriptionNil(b bool)`

 SetRecordDescriptionNil sets the value for RecordDescription to be an explicit nil

### UnsetRecordDescription
`func (o *IpmiSelNonVerboseEntry) UnsetRecordDescription()`

UnsetRecordDescription ensures that no value is present for RecordDescription, not even an explicit nil
### GetRecordEvent

`func (o *IpmiSelNonVerboseEntry) GetRecordEvent() string`

GetRecordEvent returns the RecordEvent field if non-nil, zero value otherwise.

### GetRecordEventOk

`func (o *IpmiSelNonVerboseEntry) GetRecordEventOk() (*string, bool)`

GetRecordEventOk returns a tuple with the RecordEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordEvent

`func (o *IpmiSelNonVerboseEntry) SetRecordEvent(v string)`

SetRecordEvent sets RecordEvent field to given value.

### HasRecordEvent

`func (o *IpmiSelNonVerboseEntry) HasRecordEvent() bool`

HasRecordEvent returns a boolean if a field has been set.

### SetRecordEventNil

`func (o *IpmiSelNonVerboseEntry) SetRecordEventNil(b bool)`

 SetRecordEventNil sets the value for RecordEvent to be an explicit nil

### UnsetRecordEvent
`func (o *IpmiSelNonVerboseEntry) UnsetRecordEvent()`

UnsetRecordEvent ensures that no value is present for RecordEvent, not even an explicit nil
### GetRecordId

`func (o *IpmiSelNonVerboseEntry) GetRecordId() string`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *IpmiSelNonVerboseEntry) GetRecordIdOk() (*string, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *IpmiSelNonVerboseEntry) SetRecordId(v string)`

SetRecordId sets RecordId field to given value.

### HasRecordId

`func (o *IpmiSelNonVerboseEntry) HasRecordId() bool`

HasRecordId returns a boolean if a field has been set.

### SetRecordIdNil

`func (o *IpmiSelNonVerboseEntry) SetRecordIdNil(b bool)`

 SetRecordIdNil sets the value for RecordId to be an explicit nil

### UnsetRecordId
`func (o *IpmiSelNonVerboseEntry) UnsetRecordId()`

UnsetRecordId ensures that no value is present for RecordId, not even an explicit nil
### GetRecordTime

`func (o *IpmiSelNonVerboseEntry) GetRecordTime() string`

GetRecordTime returns the RecordTime field if non-nil, zero value otherwise.

### GetRecordTimeOk

`func (o *IpmiSelNonVerboseEntry) GetRecordTimeOk() (*string, bool)`

GetRecordTimeOk returns a tuple with the RecordTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTime

`func (o *IpmiSelNonVerboseEntry) SetRecordTime(v string)`

SetRecordTime sets RecordTime field to given value.

### HasRecordTime

`func (o *IpmiSelNonVerboseEntry) HasRecordTime() bool`

HasRecordTime returns a boolean if a field has been set.

### SetRecordTimeNil

`func (o *IpmiSelNonVerboseEntry) SetRecordTimeNil(b bool)`

 SetRecordTimeNil sets the value for RecordTime to be an explicit nil

### UnsetRecordTime
`func (o *IpmiSelNonVerboseEntry) UnsetRecordTime()`

UnsetRecordTime ensures that no value is present for RecordTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


