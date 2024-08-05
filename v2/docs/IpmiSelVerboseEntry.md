# IpmiSelVerboseEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventData** | Pointer to **NullableString** | Specifies additional info about the event. | [optional] 
**EventDirection** | Pointer to **NullableString** | Specifies whether the event occurred is assertion or deassertion. | [optional] 
**EventType** | Pointer to **NullableString** | Specifies the type of event occurred. | [optional] 
**EvmRevision** | Pointer to **NullableString** | Specifies the version of the Event Message Format for the record added to SEL. | [optional] 
**GeneratorId** | Pointer to **NullableString** | Corresponds to source of component that generated the record. | [optional] 
**RecordDescription** | Pointer to **NullableString** | Specifies a short description corresponding to the sensor event for which record is added to SEL. | [optional] 
**RecordId** | Pointer to **NullableString** | Specifies the ID corresponding to record in SEL(System Event Log) for given node. | [optional] 
**RecordTimestamp** | Pointer to **NullableString** | Specifies the time stamp at which the record is added to SEL. | [optional] 
**RecordType** | Pointer to **NullableString** | Specifies the type of SEL record corresponding to the entry. | [optional] 
**SensorNumber** | Pointer to **NullableString** | Specifies the sensor number corresponding to the current SEL record. | [optional] 
**SensorType** | Pointer to **NullableString** | Specifies the type of the sensor corresponding to the current SEL record. | [optional] 

## Methods

### NewIpmiSelVerboseEntry

`func NewIpmiSelVerboseEntry() *IpmiSelVerboseEntry`

NewIpmiSelVerboseEntry instantiates a new IpmiSelVerboseEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpmiSelVerboseEntryWithDefaults

`func NewIpmiSelVerboseEntryWithDefaults() *IpmiSelVerboseEntry`

NewIpmiSelVerboseEntryWithDefaults instantiates a new IpmiSelVerboseEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventData

`func (o *IpmiSelVerboseEntry) GetEventData() string`

GetEventData returns the EventData field if non-nil, zero value otherwise.

### GetEventDataOk

`func (o *IpmiSelVerboseEntry) GetEventDataOk() (*string, bool)`

GetEventDataOk returns a tuple with the EventData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventData

`func (o *IpmiSelVerboseEntry) SetEventData(v string)`

SetEventData sets EventData field to given value.

### HasEventData

`func (o *IpmiSelVerboseEntry) HasEventData() bool`

HasEventData returns a boolean if a field has been set.

### SetEventDataNil

`func (o *IpmiSelVerboseEntry) SetEventDataNil(b bool)`

 SetEventDataNil sets the value for EventData to be an explicit nil

### UnsetEventData
`func (o *IpmiSelVerboseEntry) UnsetEventData()`

UnsetEventData ensures that no value is present for EventData, not even an explicit nil
### GetEventDirection

`func (o *IpmiSelVerboseEntry) GetEventDirection() string`

GetEventDirection returns the EventDirection field if non-nil, zero value otherwise.

### GetEventDirectionOk

`func (o *IpmiSelVerboseEntry) GetEventDirectionOk() (*string, bool)`

GetEventDirectionOk returns a tuple with the EventDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDirection

`func (o *IpmiSelVerboseEntry) SetEventDirection(v string)`

SetEventDirection sets EventDirection field to given value.

### HasEventDirection

`func (o *IpmiSelVerboseEntry) HasEventDirection() bool`

HasEventDirection returns a boolean if a field has been set.

### SetEventDirectionNil

`func (o *IpmiSelVerboseEntry) SetEventDirectionNil(b bool)`

 SetEventDirectionNil sets the value for EventDirection to be an explicit nil

### UnsetEventDirection
`func (o *IpmiSelVerboseEntry) UnsetEventDirection()`

UnsetEventDirection ensures that no value is present for EventDirection, not even an explicit nil
### GetEventType

`func (o *IpmiSelVerboseEntry) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *IpmiSelVerboseEntry) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *IpmiSelVerboseEntry) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *IpmiSelVerboseEntry) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### SetEventTypeNil

`func (o *IpmiSelVerboseEntry) SetEventTypeNil(b bool)`

 SetEventTypeNil sets the value for EventType to be an explicit nil

### UnsetEventType
`func (o *IpmiSelVerboseEntry) UnsetEventType()`

UnsetEventType ensures that no value is present for EventType, not even an explicit nil
### GetEvmRevision

`func (o *IpmiSelVerboseEntry) GetEvmRevision() string`

GetEvmRevision returns the EvmRevision field if non-nil, zero value otherwise.

### GetEvmRevisionOk

`func (o *IpmiSelVerboseEntry) GetEvmRevisionOk() (*string, bool)`

GetEvmRevisionOk returns a tuple with the EvmRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvmRevision

`func (o *IpmiSelVerboseEntry) SetEvmRevision(v string)`

SetEvmRevision sets EvmRevision field to given value.

### HasEvmRevision

`func (o *IpmiSelVerboseEntry) HasEvmRevision() bool`

HasEvmRevision returns a boolean if a field has been set.

### SetEvmRevisionNil

`func (o *IpmiSelVerboseEntry) SetEvmRevisionNil(b bool)`

 SetEvmRevisionNil sets the value for EvmRevision to be an explicit nil

### UnsetEvmRevision
`func (o *IpmiSelVerboseEntry) UnsetEvmRevision()`

UnsetEvmRevision ensures that no value is present for EvmRevision, not even an explicit nil
### GetGeneratorId

`func (o *IpmiSelVerboseEntry) GetGeneratorId() string`

GetGeneratorId returns the GeneratorId field if non-nil, zero value otherwise.

### GetGeneratorIdOk

`func (o *IpmiSelVerboseEntry) GetGeneratorIdOk() (*string, bool)`

GetGeneratorIdOk returns a tuple with the GeneratorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratorId

`func (o *IpmiSelVerboseEntry) SetGeneratorId(v string)`

SetGeneratorId sets GeneratorId field to given value.

### HasGeneratorId

`func (o *IpmiSelVerboseEntry) HasGeneratorId() bool`

HasGeneratorId returns a boolean if a field has been set.

### SetGeneratorIdNil

`func (o *IpmiSelVerboseEntry) SetGeneratorIdNil(b bool)`

 SetGeneratorIdNil sets the value for GeneratorId to be an explicit nil

### UnsetGeneratorId
`func (o *IpmiSelVerboseEntry) UnsetGeneratorId()`

UnsetGeneratorId ensures that no value is present for GeneratorId, not even an explicit nil
### GetRecordDescription

`func (o *IpmiSelVerboseEntry) GetRecordDescription() string`

GetRecordDescription returns the RecordDescription field if non-nil, zero value otherwise.

### GetRecordDescriptionOk

`func (o *IpmiSelVerboseEntry) GetRecordDescriptionOk() (*string, bool)`

GetRecordDescriptionOk returns a tuple with the RecordDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordDescription

`func (o *IpmiSelVerboseEntry) SetRecordDescription(v string)`

SetRecordDescription sets RecordDescription field to given value.

### HasRecordDescription

`func (o *IpmiSelVerboseEntry) HasRecordDescription() bool`

HasRecordDescription returns a boolean if a field has been set.

### SetRecordDescriptionNil

`func (o *IpmiSelVerboseEntry) SetRecordDescriptionNil(b bool)`

 SetRecordDescriptionNil sets the value for RecordDescription to be an explicit nil

### UnsetRecordDescription
`func (o *IpmiSelVerboseEntry) UnsetRecordDescription()`

UnsetRecordDescription ensures that no value is present for RecordDescription, not even an explicit nil
### GetRecordId

`func (o *IpmiSelVerboseEntry) GetRecordId() string`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *IpmiSelVerboseEntry) GetRecordIdOk() (*string, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *IpmiSelVerboseEntry) SetRecordId(v string)`

SetRecordId sets RecordId field to given value.

### HasRecordId

`func (o *IpmiSelVerboseEntry) HasRecordId() bool`

HasRecordId returns a boolean if a field has been set.

### SetRecordIdNil

`func (o *IpmiSelVerboseEntry) SetRecordIdNil(b bool)`

 SetRecordIdNil sets the value for RecordId to be an explicit nil

### UnsetRecordId
`func (o *IpmiSelVerboseEntry) UnsetRecordId()`

UnsetRecordId ensures that no value is present for RecordId, not even an explicit nil
### GetRecordTimestamp

`func (o *IpmiSelVerboseEntry) GetRecordTimestamp() string`

GetRecordTimestamp returns the RecordTimestamp field if non-nil, zero value otherwise.

### GetRecordTimestampOk

`func (o *IpmiSelVerboseEntry) GetRecordTimestampOk() (*string, bool)`

GetRecordTimestampOk returns a tuple with the RecordTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTimestamp

`func (o *IpmiSelVerboseEntry) SetRecordTimestamp(v string)`

SetRecordTimestamp sets RecordTimestamp field to given value.

### HasRecordTimestamp

`func (o *IpmiSelVerboseEntry) HasRecordTimestamp() bool`

HasRecordTimestamp returns a boolean if a field has been set.

### SetRecordTimestampNil

`func (o *IpmiSelVerboseEntry) SetRecordTimestampNil(b bool)`

 SetRecordTimestampNil sets the value for RecordTimestamp to be an explicit nil

### UnsetRecordTimestamp
`func (o *IpmiSelVerboseEntry) UnsetRecordTimestamp()`

UnsetRecordTimestamp ensures that no value is present for RecordTimestamp, not even an explicit nil
### GetRecordType

`func (o *IpmiSelVerboseEntry) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *IpmiSelVerboseEntry) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *IpmiSelVerboseEntry) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *IpmiSelVerboseEntry) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### SetRecordTypeNil

`func (o *IpmiSelVerboseEntry) SetRecordTypeNil(b bool)`

 SetRecordTypeNil sets the value for RecordType to be an explicit nil

### UnsetRecordType
`func (o *IpmiSelVerboseEntry) UnsetRecordType()`

UnsetRecordType ensures that no value is present for RecordType, not even an explicit nil
### GetSensorNumber

`func (o *IpmiSelVerboseEntry) GetSensorNumber() string`

GetSensorNumber returns the SensorNumber field if non-nil, zero value otherwise.

### GetSensorNumberOk

`func (o *IpmiSelVerboseEntry) GetSensorNumberOk() (*string, bool)`

GetSensorNumberOk returns a tuple with the SensorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensorNumber

`func (o *IpmiSelVerboseEntry) SetSensorNumber(v string)`

SetSensorNumber sets SensorNumber field to given value.

### HasSensorNumber

`func (o *IpmiSelVerboseEntry) HasSensorNumber() bool`

HasSensorNumber returns a boolean if a field has been set.

### SetSensorNumberNil

`func (o *IpmiSelVerboseEntry) SetSensorNumberNil(b bool)`

 SetSensorNumberNil sets the value for SensorNumber to be an explicit nil

### UnsetSensorNumber
`func (o *IpmiSelVerboseEntry) UnsetSensorNumber()`

UnsetSensorNumber ensures that no value is present for SensorNumber, not even an explicit nil
### GetSensorType

`func (o *IpmiSelVerboseEntry) GetSensorType() string`

GetSensorType returns the SensorType field if non-nil, zero value otherwise.

### GetSensorTypeOk

`func (o *IpmiSelVerboseEntry) GetSensorTypeOk() (*string, bool)`

GetSensorTypeOk returns a tuple with the SensorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensorType

`func (o *IpmiSelVerboseEntry) SetSensorType(v string)`

SetSensorType sets SensorType field to given value.

### HasSensorType

`func (o *IpmiSelVerboseEntry) HasSensorType() bool`

HasSensorType returns a boolean if a field has been set.

### SetSensorTypeNil

`func (o *IpmiSelVerboseEntry) SetSensorTypeNil(b bool)`

 SetSensorTypeNil sets the value for SensorType to be an explicit nil

### UnsetSensorType
`func (o *IpmiSelVerboseEntry) UnsetSensorType()`

UnsetSensorType ensures that no value is present for SensorType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


