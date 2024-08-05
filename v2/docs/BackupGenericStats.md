# BackupGenericStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataIngested** | Pointer to **NullableInt64** | Specifies the amount of data which has been ingested in bytes. | [optional] 
**DataIngestionRate** | Pointer to **NullableInt64** | Specifies the rate at which data is being ingested in bytes per minute. | [optional] 
**ErrorClasses** | Pointer to [**[]ErrorClass**](ErrorClass.md) | Divides the errors into classes for better understanding for the user. | [optional] 
**EstimatedBackupTime** | Pointer to **NullableInt64** | Specifies the time in which backup should finish in minutes. | [optional] 
**NumErrors** | Pointer to **NullableInt64** | Specifies the number of errors for this run. | [optional] 
**RemainingDataIngested** | Pointer to **NullableInt64** | Specifies the amount of data which has to be ingested in bytes. | [optional] 

## Methods

### NewBackupGenericStats

`func NewBackupGenericStats() *BackupGenericStats`

NewBackupGenericStats instantiates a new BackupGenericStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupGenericStatsWithDefaults

`func NewBackupGenericStatsWithDefaults() *BackupGenericStats`

NewBackupGenericStatsWithDefaults instantiates a new BackupGenericStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataIngested

`func (o *BackupGenericStats) GetDataIngested() int64`

GetDataIngested returns the DataIngested field if non-nil, zero value otherwise.

### GetDataIngestedOk

`func (o *BackupGenericStats) GetDataIngestedOk() (*int64, bool)`

GetDataIngestedOk returns a tuple with the DataIngested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataIngested

`func (o *BackupGenericStats) SetDataIngested(v int64)`

SetDataIngested sets DataIngested field to given value.

### HasDataIngested

`func (o *BackupGenericStats) HasDataIngested() bool`

HasDataIngested returns a boolean if a field has been set.

### SetDataIngestedNil

`func (o *BackupGenericStats) SetDataIngestedNil(b bool)`

 SetDataIngestedNil sets the value for DataIngested to be an explicit nil

### UnsetDataIngested
`func (o *BackupGenericStats) UnsetDataIngested()`

UnsetDataIngested ensures that no value is present for DataIngested, not even an explicit nil
### GetDataIngestionRate

`func (o *BackupGenericStats) GetDataIngestionRate() int64`

GetDataIngestionRate returns the DataIngestionRate field if non-nil, zero value otherwise.

### GetDataIngestionRateOk

`func (o *BackupGenericStats) GetDataIngestionRateOk() (*int64, bool)`

GetDataIngestionRateOk returns a tuple with the DataIngestionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataIngestionRate

`func (o *BackupGenericStats) SetDataIngestionRate(v int64)`

SetDataIngestionRate sets DataIngestionRate field to given value.

### HasDataIngestionRate

`func (o *BackupGenericStats) HasDataIngestionRate() bool`

HasDataIngestionRate returns a boolean if a field has been set.

### SetDataIngestionRateNil

`func (o *BackupGenericStats) SetDataIngestionRateNil(b bool)`

 SetDataIngestionRateNil sets the value for DataIngestionRate to be an explicit nil

### UnsetDataIngestionRate
`func (o *BackupGenericStats) UnsetDataIngestionRate()`

UnsetDataIngestionRate ensures that no value is present for DataIngestionRate, not even an explicit nil
### GetErrorClasses

`func (o *BackupGenericStats) GetErrorClasses() []ErrorClass`

GetErrorClasses returns the ErrorClasses field if non-nil, zero value otherwise.

### GetErrorClassesOk

`func (o *BackupGenericStats) GetErrorClassesOk() (*[]ErrorClass, bool)`

GetErrorClassesOk returns a tuple with the ErrorClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorClasses

`func (o *BackupGenericStats) SetErrorClasses(v []ErrorClass)`

SetErrorClasses sets ErrorClasses field to given value.

### HasErrorClasses

`func (o *BackupGenericStats) HasErrorClasses() bool`

HasErrorClasses returns a boolean if a field has been set.

### SetErrorClassesNil

`func (o *BackupGenericStats) SetErrorClassesNil(b bool)`

 SetErrorClassesNil sets the value for ErrorClasses to be an explicit nil

### UnsetErrorClasses
`func (o *BackupGenericStats) UnsetErrorClasses()`

UnsetErrorClasses ensures that no value is present for ErrorClasses, not even an explicit nil
### GetEstimatedBackupTime

`func (o *BackupGenericStats) GetEstimatedBackupTime() int64`

GetEstimatedBackupTime returns the EstimatedBackupTime field if non-nil, zero value otherwise.

### GetEstimatedBackupTimeOk

`func (o *BackupGenericStats) GetEstimatedBackupTimeOk() (*int64, bool)`

GetEstimatedBackupTimeOk returns a tuple with the EstimatedBackupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedBackupTime

`func (o *BackupGenericStats) SetEstimatedBackupTime(v int64)`

SetEstimatedBackupTime sets EstimatedBackupTime field to given value.

### HasEstimatedBackupTime

`func (o *BackupGenericStats) HasEstimatedBackupTime() bool`

HasEstimatedBackupTime returns a boolean if a field has been set.

### SetEstimatedBackupTimeNil

`func (o *BackupGenericStats) SetEstimatedBackupTimeNil(b bool)`

 SetEstimatedBackupTimeNil sets the value for EstimatedBackupTime to be an explicit nil

### UnsetEstimatedBackupTime
`func (o *BackupGenericStats) UnsetEstimatedBackupTime()`

UnsetEstimatedBackupTime ensures that no value is present for EstimatedBackupTime, not even an explicit nil
### GetNumErrors

`func (o *BackupGenericStats) GetNumErrors() int64`

GetNumErrors returns the NumErrors field if non-nil, zero value otherwise.

### GetNumErrorsOk

`func (o *BackupGenericStats) GetNumErrorsOk() (*int64, bool)`

GetNumErrorsOk returns a tuple with the NumErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumErrors

`func (o *BackupGenericStats) SetNumErrors(v int64)`

SetNumErrors sets NumErrors field to given value.

### HasNumErrors

`func (o *BackupGenericStats) HasNumErrors() bool`

HasNumErrors returns a boolean if a field has been set.

### SetNumErrorsNil

`func (o *BackupGenericStats) SetNumErrorsNil(b bool)`

 SetNumErrorsNil sets the value for NumErrors to be an explicit nil

### UnsetNumErrors
`func (o *BackupGenericStats) UnsetNumErrors()`

UnsetNumErrors ensures that no value is present for NumErrors, not even an explicit nil
### GetRemainingDataIngested

`func (o *BackupGenericStats) GetRemainingDataIngested() int64`

GetRemainingDataIngested returns the RemainingDataIngested field if non-nil, zero value otherwise.

### GetRemainingDataIngestedOk

`func (o *BackupGenericStats) GetRemainingDataIngestedOk() (*int64, bool)`

GetRemainingDataIngestedOk returns a tuple with the RemainingDataIngested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingDataIngested

`func (o *BackupGenericStats) SetRemainingDataIngested(v int64)`

SetRemainingDataIngested sets RemainingDataIngested field to given value.

### HasRemainingDataIngested

`func (o *BackupGenericStats) HasRemainingDataIngested() bool`

HasRemainingDataIngested returns a boolean if a field has been set.

### SetRemainingDataIngestedNil

`func (o *BackupGenericStats) SetRemainingDataIngestedNil(b bool)`

 SetRemainingDataIngestedNil sets the value for RemainingDataIngested to be an explicit nil

### UnsetRemainingDataIngested
`func (o *BackupGenericStats) UnsetRemainingDataIngested()`

UnsetRemainingDataIngested ensures that no value is present for RemainingDataIngested, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


