# BackupPolicyProtoMonthlySchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** | Count of the day on which to perform the backup (look above for a more detailed description). | [optional] 
**Day** | Pointer to **NullableInt32** | The day of the month the backup is to be performed. | [optional] 
**Time** | Pointer to [**Time**](Time.md) |  | [optional] 

## Methods

### NewBackupPolicyProtoMonthlySchedule

`func NewBackupPolicyProtoMonthlySchedule() *BackupPolicyProtoMonthlySchedule`

NewBackupPolicyProtoMonthlySchedule instantiates a new BackupPolicyProtoMonthlySchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupPolicyProtoMonthlyScheduleWithDefaults

`func NewBackupPolicyProtoMonthlyScheduleWithDefaults() *BackupPolicyProtoMonthlySchedule`

NewBackupPolicyProtoMonthlyScheduleWithDefaults instantiates a new BackupPolicyProtoMonthlySchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *BackupPolicyProtoMonthlySchedule) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BackupPolicyProtoMonthlySchedule) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BackupPolicyProtoMonthlySchedule) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *BackupPolicyProtoMonthlySchedule) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *BackupPolicyProtoMonthlySchedule) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *BackupPolicyProtoMonthlySchedule) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetDay

`func (o *BackupPolicyProtoMonthlySchedule) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *BackupPolicyProtoMonthlySchedule) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *BackupPolicyProtoMonthlySchedule) SetDay(v int32)`

SetDay sets Day field to given value.

### HasDay

`func (o *BackupPolicyProtoMonthlySchedule) HasDay() bool`

HasDay returns a boolean if a field has been set.

### SetDayNil

`func (o *BackupPolicyProtoMonthlySchedule) SetDayNil(b bool)`

 SetDayNil sets the value for Day to be an explicit nil

### UnsetDay
`func (o *BackupPolicyProtoMonthlySchedule) UnsetDay()`

UnsetDay ensures that no value is present for Day, not even an explicit nil
### GetTime

`func (o *BackupPolicyProtoMonthlySchedule) GetTime() Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *BackupPolicyProtoMonthlySchedule) GetTimeOk() (*Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *BackupPolicyProtoMonthlySchedule) SetTime(v Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *BackupPolicyProtoMonthlySchedule) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


