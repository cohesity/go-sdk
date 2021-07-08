# LastProtectionRunSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfCancelledProtectionRuns** | Pointer to **NullableInt64** | Specifies the number of cancelled Protection Runs the given Protection Policy has in the Last Run. | [optional] 
**NumberOfFailedProtectionRuns** | Pointer to **NullableInt64** | Specifies the number of failed Protection Runs the given Protection Policy has in the Last Run. | [optional] 
**NumberOfProtectedSources** | Pointer to **NullableInt64** | Specifies the number of Protection Sources protected by the given Protection Policy. | [optional] 
**NumberOfRunningProtectionRuns** | Pointer to **NullableInt64** | Specifies the number of running Protection Runs using the current Protection Policy. | [optional] 
**NumberOfSlaViolations** | Pointer to **NullableInt64** | Specifies the number of SLA violations the given Protection Policy has in the Last Run.. | [optional] 
**NumberOfSuccessfulProtectionRuns** | Pointer to **NullableInt64** | Specifies the number of successful Protection Runs the given Protection Policy has in the Last Run. | [optional] 
**TotalLogicalBackupSizeInBytes** | Pointer to **NullableInt64** | Specifies the aggregated total logical backup performed in all the Latest Protection Runs made for all the Protection Jobs which have the given Protection Policy Specified. | [optional] 

## Methods

### NewLastProtectionRunSummary

`func NewLastProtectionRunSummary() *LastProtectionRunSummary`

NewLastProtectionRunSummary instantiates a new LastProtectionRunSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastProtectionRunSummaryWithDefaults

`func NewLastProtectionRunSummaryWithDefaults() *LastProtectionRunSummary`

NewLastProtectionRunSummaryWithDefaults instantiates a new LastProtectionRunSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfCancelledProtectionRuns

`func (o *LastProtectionRunSummary) GetNumberOfCancelledProtectionRuns() int64`

GetNumberOfCancelledProtectionRuns returns the NumberOfCancelledProtectionRuns field if non-nil, zero value otherwise.

### GetNumberOfCancelledProtectionRunsOk

`func (o *LastProtectionRunSummary) GetNumberOfCancelledProtectionRunsOk() (*int64, bool)`

GetNumberOfCancelledProtectionRunsOk returns a tuple with the NumberOfCancelledProtectionRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCancelledProtectionRuns

`func (o *LastProtectionRunSummary) SetNumberOfCancelledProtectionRuns(v int64)`

SetNumberOfCancelledProtectionRuns sets NumberOfCancelledProtectionRuns field to given value.

### HasNumberOfCancelledProtectionRuns

`func (o *LastProtectionRunSummary) HasNumberOfCancelledProtectionRuns() bool`

HasNumberOfCancelledProtectionRuns returns a boolean if a field has been set.

### SetNumberOfCancelledProtectionRunsNil

`func (o *LastProtectionRunSummary) SetNumberOfCancelledProtectionRunsNil(b bool)`

 SetNumberOfCancelledProtectionRunsNil sets the value for NumberOfCancelledProtectionRuns to be an explicit nil

### UnsetNumberOfCancelledProtectionRuns
`func (o *LastProtectionRunSummary) UnsetNumberOfCancelledProtectionRuns()`

UnsetNumberOfCancelledProtectionRuns ensures that no value is present for NumberOfCancelledProtectionRuns, not even an explicit nil
### GetNumberOfFailedProtectionRuns

`func (o *LastProtectionRunSummary) GetNumberOfFailedProtectionRuns() int64`

GetNumberOfFailedProtectionRuns returns the NumberOfFailedProtectionRuns field if non-nil, zero value otherwise.

### GetNumberOfFailedProtectionRunsOk

`func (o *LastProtectionRunSummary) GetNumberOfFailedProtectionRunsOk() (*int64, bool)`

GetNumberOfFailedProtectionRunsOk returns a tuple with the NumberOfFailedProtectionRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfFailedProtectionRuns

`func (o *LastProtectionRunSummary) SetNumberOfFailedProtectionRuns(v int64)`

SetNumberOfFailedProtectionRuns sets NumberOfFailedProtectionRuns field to given value.

### HasNumberOfFailedProtectionRuns

`func (o *LastProtectionRunSummary) HasNumberOfFailedProtectionRuns() bool`

HasNumberOfFailedProtectionRuns returns a boolean if a field has been set.

### SetNumberOfFailedProtectionRunsNil

`func (o *LastProtectionRunSummary) SetNumberOfFailedProtectionRunsNil(b bool)`

 SetNumberOfFailedProtectionRunsNil sets the value for NumberOfFailedProtectionRuns to be an explicit nil

### UnsetNumberOfFailedProtectionRuns
`func (o *LastProtectionRunSummary) UnsetNumberOfFailedProtectionRuns()`

UnsetNumberOfFailedProtectionRuns ensures that no value is present for NumberOfFailedProtectionRuns, not even an explicit nil
### GetNumberOfProtectedSources

`func (o *LastProtectionRunSummary) GetNumberOfProtectedSources() int64`

GetNumberOfProtectedSources returns the NumberOfProtectedSources field if non-nil, zero value otherwise.

### GetNumberOfProtectedSourcesOk

`func (o *LastProtectionRunSummary) GetNumberOfProtectedSourcesOk() (*int64, bool)`

GetNumberOfProtectedSourcesOk returns a tuple with the NumberOfProtectedSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfProtectedSources

`func (o *LastProtectionRunSummary) SetNumberOfProtectedSources(v int64)`

SetNumberOfProtectedSources sets NumberOfProtectedSources field to given value.

### HasNumberOfProtectedSources

`func (o *LastProtectionRunSummary) HasNumberOfProtectedSources() bool`

HasNumberOfProtectedSources returns a boolean if a field has been set.

### SetNumberOfProtectedSourcesNil

`func (o *LastProtectionRunSummary) SetNumberOfProtectedSourcesNil(b bool)`

 SetNumberOfProtectedSourcesNil sets the value for NumberOfProtectedSources to be an explicit nil

### UnsetNumberOfProtectedSources
`func (o *LastProtectionRunSummary) UnsetNumberOfProtectedSources()`

UnsetNumberOfProtectedSources ensures that no value is present for NumberOfProtectedSources, not even an explicit nil
### GetNumberOfRunningProtectionRuns

`func (o *LastProtectionRunSummary) GetNumberOfRunningProtectionRuns() int64`

GetNumberOfRunningProtectionRuns returns the NumberOfRunningProtectionRuns field if non-nil, zero value otherwise.

### GetNumberOfRunningProtectionRunsOk

`func (o *LastProtectionRunSummary) GetNumberOfRunningProtectionRunsOk() (*int64, bool)`

GetNumberOfRunningProtectionRunsOk returns a tuple with the NumberOfRunningProtectionRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfRunningProtectionRuns

`func (o *LastProtectionRunSummary) SetNumberOfRunningProtectionRuns(v int64)`

SetNumberOfRunningProtectionRuns sets NumberOfRunningProtectionRuns field to given value.

### HasNumberOfRunningProtectionRuns

`func (o *LastProtectionRunSummary) HasNumberOfRunningProtectionRuns() bool`

HasNumberOfRunningProtectionRuns returns a boolean if a field has been set.

### SetNumberOfRunningProtectionRunsNil

`func (o *LastProtectionRunSummary) SetNumberOfRunningProtectionRunsNil(b bool)`

 SetNumberOfRunningProtectionRunsNil sets the value for NumberOfRunningProtectionRuns to be an explicit nil

### UnsetNumberOfRunningProtectionRuns
`func (o *LastProtectionRunSummary) UnsetNumberOfRunningProtectionRuns()`

UnsetNumberOfRunningProtectionRuns ensures that no value is present for NumberOfRunningProtectionRuns, not even an explicit nil
### GetNumberOfSlaViolations

`func (o *LastProtectionRunSummary) GetNumberOfSlaViolations() int64`

GetNumberOfSlaViolations returns the NumberOfSlaViolations field if non-nil, zero value otherwise.

### GetNumberOfSlaViolationsOk

`func (o *LastProtectionRunSummary) GetNumberOfSlaViolationsOk() (*int64, bool)`

GetNumberOfSlaViolationsOk returns a tuple with the NumberOfSlaViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSlaViolations

`func (o *LastProtectionRunSummary) SetNumberOfSlaViolations(v int64)`

SetNumberOfSlaViolations sets NumberOfSlaViolations field to given value.

### HasNumberOfSlaViolations

`func (o *LastProtectionRunSummary) HasNumberOfSlaViolations() bool`

HasNumberOfSlaViolations returns a boolean if a field has been set.

### SetNumberOfSlaViolationsNil

`func (o *LastProtectionRunSummary) SetNumberOfSlaViolationsNil(b bool)`

 SetNumberOfSlaViolationsNil sets the value for NumberOfSlaViolations to be an explicit nil

### UnsetNumberOfSlaViolations
`func (o *LastProtectionRunSummary) UnsetNumberOfSlaViolations()`

UnsetNumberOfSlaViolations ensures that no value is present for NumberOfSlaViolations, not even an explicit nil
### GetNumberOfSuccessfulProtectionRuns

`func (o *LastProtectionRunSummary) GetNumberOfSuccessfulProtectionRuns() int64`

GetNumberOfSuccessfulProtectionRuns returns the NumberOfSuccessfulProtectionRuns field if non-nil, zero value otherwise.

### GetNumberOfSuccessfulProtectionRunsOk

`func (o *LastProtectionRunSummary) GetNumberOfSuccessfulProtectionRunsOk() (*int64, bool)`

GetNumberOfSuccessfulProtectionRunsOk returns a tuple with the NumberOfSuccessfulProtectionRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSuccessfulProtectionRuns

`func (o *LastProtectionRunSummary) SetNumberOfSuccessfulProtectionRuns(v int64)`

SetNumberOfSuccessfulProtectionRuns sets NumberOfSuccessfulProtectionRuns field to given value.

### HasNumberOfSuccessfulProtectionRuns

`func (o *LastProtectionRunSummary) HasNumberOfSuccessfulProtectionRuns() bool`

HasNumberOfSuccessfulProtectionRuns returns a boolean if a field has been set.

### SetNumberOfSuccessfulProtectionRunsNil

`func (o *LastProtectionRunSummary) SetNumberOfSuccessfulProtectionRunsNil(b bool)`

 SetNumberOfSuccessfulProtectionRunsNil sets the value for NumberOfSuccessfulProtectionRuns to be an explicit nil

### UnsetNumberOfSuccessfulProtectionRuns
`func (o *LastProtectionRunSummary) UnsetNumberOfSuccessfulProtectionRuns()`

UnsetNumberOfSuccessfulProtectionRuns ensures that no value is present for NumberOfSuccessfulProtectionRuns, not even an explicit nil
### GetTotalLogicalBackupSizeInBytes

`func (o *LastProtectionRunSummary) GetTotalLogicalBackupSizeInBytes() int64`

GetTotalLogicalBackupSizeInBytes returns the TotalLogicalBackupSizeInBytes field if non-nil, zero value otherwise.

### GetTotalLogicalBackupSizeInBytesOk

`func (o *LastProtectionRunSummary) GetTotalLogicalBackupSizeInBytesOk() (*int64, bool)`

GetTotalLogicalBackupSizeInBytesOk returns a tuple with the TotalLogicalBackupSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLogicalBackupSizeInBytes

`func (o *LastProtectionRunSummary) SetTotalLogicalBackupSizeInBytes(v int64)`

SetTotalLogicalBackupSizeInBytes sets TotalLogicalBackupSizeInBytes field to given value.

### HasTotalLogicalBackupSizeInBytes

`func (o *LastProtectionRunSummary) HasTotalLogicalBackupSizeInBytes() bool`

HasTotalLogicalBackupSizeInBytes returns a boolean if a field has been set.

### SetTotalLogicalBackupSizeInBytesNil

`func (o *LastProtectionRunSummary) SetTotalLogicalBackupSizeInBytesNil(b bool)`

 SetTotalLogicalBackupSizeInBytesNil sets the value for TotalLogicalBackupSizeInBytes to be an explicit nil

### UnsetTotalLogicalBackupSizeInBytes
`func (o *LastProtectionRunSummary) UnsetTotalLogicalBackupSizeInBytes()`

UnsetTotalLogicalBackupSizeInBytes ensures that no value is present for TotalLogicalBackupSizeInBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


