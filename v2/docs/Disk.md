# Disk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityInBytes** | Pointer to **NullableInt64** | Specifies capacity of disk in bytes. | [optional] 
**EncryptionStatus** | Pointer to **string** | Specifies disk encryption state. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies id to uniquely identify a disk. | [optional] 
**Location** | Pointer to **NullableString** | Specifies location of the disk in node. | [optional] 
**Model** | Pointer to **NullableString** | Specifies product model of disk. | [optional] 
**NodeId** | Pointer to **NullableInt64** | Specifies node id of the node that this disk belong to. | [optional] 
**PrecheckTimestampSecs** | Pointer to **NullableInt64** | Specifies the last run time of the pre-checks execution in Unix epoch timestamp (in seconds). | [optional] 
**ProgressPercentage** | Pointer to **NullableInt32** | Specifies the overall progress percentage in removing the Disk. | [optional] 
**RemovalProgressList** | Pointer to [**[]ComponentRemovalProgress**](ComponentRemovalProgress.md) | Specifies the removal progress details for services that are not acked yet. | [optional] 
**RemovalReason** | Pointer to **NullableString** | Specifies the removal reason of the disk. | [optional] 
**RemovalTimestampSecs** | Pointer to **NullableInt64** | Specifies the Unix epoch timestamp (in seconds) when the Disk was marked for removal. | [optional] 
**SerialNumber** | Pointer to **NullableString** | Specifies serial number of disk. | [optional] 
**ServicesAckedList** | Pointer to **[]string** | Specifies the services already ACKed for removal of this entity. | [optional] 
**ServicesNotAcked** | Pointer to **NullableString** | Specifies the services that are not ACKed after disk is marked for removal. | [optional] 
**ServicesNotAckedList** | Pointer to **[]string** | Specifies the services not ACKed yet for removal of this entity. | [optional] 
**SsdUsedPercentage** | Pointer to **NullableInt32** | Specifies SSD used percentage. | [optional] 
**Status** | Pointer to **string** | Specifies status of the disk. | [optional] 
**TimeRemaining** | Pointer to **NullableInt64** | Specifies the total duration in seconds left to remove the Disk. | [optional] 
**Type** | Pointer to **string** | Specifies type of the disk. | [optional] 
**ValidationChecks** | Pointer to [**[]PreCheckValidation**](PreCheckValidation.md) | Specifies the pre-check validations results. | [optional] 

## Methods

### NewDisk

`func NewDisk() *Disk`

NewDisk instantiates a new Disk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskWithDefaults

`func NewDiskWithDefaults() *Disk`

NewDiskWithDefaults instantiates a new Disk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityInBytes

`func (o *Disk) GetCapacityInBytes() int64`

GetCapacityInBytes returns the CapacityInBytes field if non-nil, zero value otherwise.

### GetCapacityInBytesOk

`func (o *Disk) GetCapacityInBytesOk() (*int64, bool)`

GetCapacityInBytesOk returns a tuple with the CapacityInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityInBytes

`func (o *Disk) SetCapacityInBytes(v int64)`

SetCapacityInBytes sets CapacityInBytes field to given value.

### HasCapacityInBytes

`func (o *Disk) HasCapacityInBytes() bool`

HasCapacityInBytes returns a boolean if a field has been set.

### SetCapacityInBytesNil

`func (o *Disk) SetCapacityInBytesNil(b bool)`

 SetCapacityInBytesNil sets the value for CapacityInBytes to be an explicit nil

### UnsetCapacityInBytes
`func (o *Disk) UnsetCapacityInBytes()`

UnsetCapacityInBytes ensures that no value is present for CapacityInBytes, not even an explicit nil
### GetEncryptionStatus

`func (o *Disk) GetEncryptionStatus() string`

GetEncryptionStatus returns the EncryptionStatus field if non-nil, zero value otherwise.

### GetEncryptionStatusOk

`func (o *Disk) GetEncryptionStatusOk() (*string, bool)`

GetEncryptionStatusOk returns a tuple with the EncryptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionStatus

`func (o *Disk) SetEncryptionStatus(v string)`

SetEncryptionStatus sets EncryptionStatus field to given value.

### HasEncryptionStatus

`func (o *Disk) HasEncryptionStatus() bool`

HasEncryptionStatus returns a boolean if a field has been set.

### GetId

`func (o *Disk) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Disk) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Disk) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Disk) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Disk) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Disk) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLocation

`func (o *Disk) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Disk) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Disk) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Disk) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *Disk) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *Disk) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetModel

`func (o *Disk) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Disk) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Disk) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Disk) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *Disk) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *Disk) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetNodeId

`func (o *Disk) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *Disk) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *Disk) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *Disk) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### SetNodeIdNil

`func (o *Disk) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *Disk) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetPrecheckTimestampSecs

`func (o *Disk) GetPrecheckTimestampSecs() int64`

GetPrecheckTimestampSecs returns the PrecheckTimestampSecs field if non-nil, zero value otherwise.

### GetPrecheckTimestampSecsOk

`func (o *Disk) GetPrecheckTimestampSecsOk() (*int64, bool)`

GetPrecheckTimestampSecsOk returns a tuple with the PrecheckTimestampSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecheckTimestampSecs

`func (o *Disk) SetPrecheckTimestampSecs(v int64)`

SetPrecheckTimestampSecs sets PrecheckTimestampSecs field to given value.

### HasPrecheckTimestampSecs

`func (o *Disk) HasPrecheckTimestampSecs() bool`

HasPrecheckTimestampSecs returns a boolean if a field has been set.

### SetPrecheckTimestampSecsNil

`func (o *Disk) SetPrecheckTimestampSecsNil(b bool)`

 SetPrecheckTimestampSecsNil sets the value for PrecheckTimestampSecs to be an explicit nil

### UnsetPrecheckTimestampSecs
`func (o *Disk) UnsetPrecheckTimestampSecs()`

UnsetPrecheckTimestampSecs ensures that no value is present for PrecheckTimestampSecs, not even an explicit nil
### GetProgressPercentage

`func (o *Disk) GetProgressPercentage() int32`

GetProgressPercentage returns the ProgressPercentage field if non-nil, zero value otherwise.

### GetProgressPercentageOk

`func (o *Disk) GetProgressPercentageOk() (*int32, bool)`

GetProgressPercentageOk returns a tuple with the ProgressPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressPercentage

`func (o *Disk) SetProgressPercentage(v int32)`

SetProgressPercentage sets ProgressPercentage field to given value.

### HasProgressPercentage

`func (o *Disk) HasProgressPercentage() bool`

HasProgressPercentage returns a boolean if a field has been set.

### SetProgressPercentageNil

`func (o *Disk) SetProgressPercentageNil(b bool)`

 SetProgressPercentageNil sets the value for ProgressPercentage to be an explicit nil

### UnsetProgressPercentage
`func (o *Disk) UnsetProgressPercentage()`

UnsetProgressPercentage ensures that no value is present for ProgressPercentage, not even an explicit nil
### GetRemovalProgressList

`func (o *Disk) GetRemovalProgressList() []ComponentRemovalProgress`

GetRemovalProgressList returns the RemovalProgressList field if non-nil, zero value otherwise.

### GetRemovalProgressListOk

`func (o *Disk) GetRemovalProgressListOk() (*[]ComponentRemovalProgress, bool)`

GetRemovalProgressListOk returns a tuple with the RemovalProgressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalProgressList

`func (o *Disk) SetRemovalProgressList(v []ComponentRemovalProgress)`

SetRemovalProgressList sets RemovalProgressList field to given value.

### HasRemovalProgressList

`func (o *Disk) HasRemovalProgressList() bool`

HasRemovalProgressList returns a boolean if a field has been set.

### SetRemovalProgressListNil

`func (o *Disk) SetRemovalProgressListNil(b bool)`

 SetRemovalProgressListNil sets the value for RemovalProgressList to be an explicit nil

### UnsetRemovalProgressList
`func (o *Disk) UnsetRemovalProgressList()`

UnsetRemovalProgressList ensures that no value is present for RemovalProgressList, not even an explicit nil
### GetRemovalReason

`func (o *Disk) GetRemovalReason() string`

GetRemovalReason returns the RemovalReason field if non-nil, zero value otherwise.

### GetRemovalReasonOk

`func (o *Disk) GetRemovalReasonOk() (*string, bool)`

GetRemovalReasonOk returns a tuple with the RemovalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalReason

`func (o *Disk) SetRemovalReason(v string)`

SetRemovalReason sets RemovalReason field to given value.

### HasRemovalReason

`func (o *Disk) HasRemovalReason() bool`

HasRemovalReason returns a boolean if a field has been set.

### SetRemovalReasonNil

`func (o *Disk) SetRemovalReasonNil(b bool)`

 SetRemovalReasonNil sets the value for RemovalReason to be an explicit nil

### UnsetRemovalReason
`func (o *Disk) UnsetRemovalReason()`

UnsetRemovalReason ensures that no value is present for RemovalReason, not even an explicit nil
### GetRemovalTimestampSecs

`func (o *Disk) GetRemovalTimestampSecs() int64`

GetRemovalTimestampSecs returns the RemovalTimestampSecs field if non-nil, zero value otherwise.

### GetRemovalTimestampSecsOk

`func (o *Disk) GetRemovalTimestampSecsOk() (*int64, bool)`

GetRemovalTimestampSecsOk returns a tuple with the RemovalTimestampSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalTimestampSecs

`func (o *Disk) SetRemovalTimestampSecs(v int64)`

SetRemovalTimestampSecs sets RemovalTimestampSecs field to given value.

### HasRemovalTimestampSecs

`func (o *Disk) HasRemovalTimestampSecs() bool`

HasRemovalTimestampSecs returns a boolean if a field has been set.

### SetRemovalTimestampSecsNil

`func (o *Disk) SetRemovalTimestampSecsNil(b bool)`

 SetRemovalTimestampSecsNil sets the value for RemovalTimestampSecs to be an explicit nil

### UnsetRemovalTimestampSecs
`func (o *Disk) UnsetRemovalTimestampSecs()`

UnsetRemovalTimestampSecs ensures that no value is present for RemovalTimestampSecs, not even an explicit nil
### GetSerialNumber

`func (o *Disk) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Disk) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Disk) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Disk) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *Disk) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *Disk) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetServicesAckedList

`func (o *Disk) GetServicesAckedList() []string`

GetServicesAckedList returns the ServicesAckedList field if non-nil, zero value otherwise.

### GetServicesAckedListOk

`func (o *Disk) GetServicesAckedListOk() (*[]string, bool)`

GetServicesAckedListOk returns a tuple with the ServicesAckedList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesAckedList

`func (o *Disk) SetServicesAckedList(v []string)`

SetServicesAckedList sets ServicesAckedList field to given value.

### HasServicesAckedList

`func (o *Disk) HasServicesAckedList() bool`

HasServicesAckedList returns a boolean if a field has been set.

### SetServicesAckedListNil

`func (o *Disk) SetServicesAckedListNil(b bool)`

 SetServicesAckedListNil sets the value for ServicesAckedList to be an explicit nil

### UnsetServicesAckedList
`func (o *Disk) UnsetServicesAckedList()`

UnsetServicesAckedList ensures that no value is present for ServicesAckedList, not even an explicit nil
### GetServicesNotAcked

`func (o *Disk) GetServicesNotAcked() string`

GetServicesNotAcked returns the ServicesNotAcked field if non-nil, zero value otherwise.

### GetServicesNotAckedOk

`func (o *Disk) GetServicesNotAckedOk() (*string, bool)`

GetServicesNotAckedOk returns a tuple with the ServicesNotAcked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesNotAcked

`func (o *Disk) SetServicesNotAcked(v string)`

SetServicesNotAcked sets ServicesNotAcked field to given value.

### HasServicesNotAcked

`func (o *Disk) HasServicesNotAcked() bool`

HasServicesNotAcked returns a boolean if a field has been set.

### SetServicesNotAckedNil

`func (o *Disk) SetServicesNotAckedNil(b bool)`

 SetServicesNotAckedNil sets the value for ServicesNotAcked to be an explicit nil

### UnsetServicesNotAcked
`func (o *Disk) UnsetServicesNotAcked()`

UnsetServicesNotAcked ensures that no value is present for ServicesNotAcked, not even an explicit nil
### GetServicesNotAckedList

`func (o *Disk) GetServicesNotAckedList() []string`

GetServicesNotAckedList returns the ServicesNotAckedList field if non-nil, zero value otherwise.

### GetServicesNotAckedListOk

`func (o *Disk) GetServicesNotAckedListOk() (*[]string, bool)`

GetServicesNotAckedListOk returns a tuple with the ServicesNotAckedList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesNotAckedList

`func (o *Disk) SetServicesNotAckedList(v []string)`

SetServicesNotAckedList sets ServicesNotAckedList field to given value.

### HasServicesNotAckedList

`func (o *Disk) HasServicesNotAckedList() bool`

HasServicesNotAckedList returns a boolean if a field has been set.

### SetServicesNotAckedListNil

`func (o *Disk) SetServicesNotAckedListNil(b bool)`

 SetServicesNotAckedListNil sets the value for ServicesNotAckedList to be an explicit nil

### UnsetServicesNotAckedList
`func (o *Disk) UnsetServicesNotAckedList()`

UnsetServicesNotAckedList ensures that no value is present for ServicesNotAckedList, not even an explicit nil
### GetSsdUsedPercentage

`func (o *Disk) GetSsdUsedPercentage() int32`

GetSsdUsedPercentage returns the SsdUsedPercentage field if non-nil, zero value otherwise.

### GetSsdUsedPercentageOk

`func (o *Disk) GetSsdUsedPercentageOk() (*int32, bool)`

GetSsdUsedPercentageOk returns a tuple with the SsdUsedPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsdUsedPercentage

`func (o *Disk) SetSsdUsedPercentage(v int32)`

SetSsdUsedPercentage sets SsdUsedPercentage field to given value.

### HasSsdUsedPercentage

`func (o *Disk) HasSsdUsedPercentage() bool`

HasSsdUsedPercentage returns a boolean if a field has been set.

### SetSsdUsedPercentageNil

`func (o *Disk) SetSsdUsedPercentageNil(b bool)`

 SetSsdUsedPercentageNil sets the value for SsdUsedPercentage to be an explicit nil

### UnsetSsdUsedPercentage
`func (o *Disk) UnsetSsdUsedPercentage()`

UnsetSsdUsedPercentage ensures that no value is present for SsdUsedPercentage, not even an explicit nil
### GetStatus

`func (o *Disk) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Disk) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Disk) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Disk) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeRemaining

`func (o *Disk) GetTimeRemaining() int64`

GetTimeRemaining returns the TimeRemaining field if non-nil, zero value otherwise.

### GetTimeRemainingOk

`func (o *Disk) GetTimeRemainingOk() (*int64, bool)`

GetTimeRemainingOk returns a tuple with the TimeRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRemaining

`func (o *Disk) SetTimeRemaining(v int64)`

SetTimeRemaining sets TimeRemaining field to given value.

### HasTimeRemaining

`func (o *Disk) HasTimeRemaining() bool`

HasTimeRemaining returns a boolean if a field has been set.

### SetTimeRemainingNil

`func (o *Disk) SetTimeRemainingNil(b bool)`

 SetTimeRemainingNil sets the value for TimeRemaining to be an explicit nil

### UnsetTimeRemaining
`func (o *Disk) UnsetTimeRemaining()`

UnsetTimeRemaining ensures that no value is present for TimeRemaining, not even an explicit nil
### GetType

`func (o *Disk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Disk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Disk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Disk) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValidationChecks

`func (o *Disk) GetValidationChecks() []PreCheckValidation`

GetValidationChecks returns the ValidationChecks field if non-nil, zero value otherwise.

### GetValidationChecksOk

`func (o *Disk) GetValidationChecksOk() (*[]PreCheckValidation, bool)`

GetValidationChecksOk returns a tuple with the ValidationChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationChecks

`func (o *Disk) SetValidationChecks(v []PreCheckValidation)`

SetValidationChecks sets ValidationChecks field to given value.

### HasValidationChecks

`func (o *Disk) HasValidationChecks() bool`

HasValidationChecks returns a boolean if a field has been set.

### SetValidationChecksNil

`func (o *Disk) SetValidationChecksNil(b bool)`

 SetValidationChecksNil sets the value for ValidationChecks to be an explicit nil

### UnsetValidationChecks
`func (o *Disk) UnsetValidationChecks()`

UnsetValidationChecks ensures that no value is present for ValidationChecks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


