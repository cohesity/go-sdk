# QueryTenantMigrationValidationStatusResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportLocation** | Pointer to [**DataLocator**](DataLocator.md) |  | [optional] 
**Status** | Pointer to **string** | The status of the Validate Metadata task. | [optional] 

## Methods

### NewQueryTenantMigrationValidationStatusResult

`func NewQueryTenantMigrationValidationStatusResult() *QueryTenantMigrationValidationStatusResult`

NewQueryTenantMigrationValidationStatusResult instantiates a new QueryTenantMigrationValidationStatusResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryTenantMigrationValidationStatusResultWithDefaults

`func NewQueryTenantMigrationValidationStatusResultWithDefaults() *QueryTenantMigrationValidationStatusResult`

NewQueryTenantMigrationValidationStatusResultWithDefaults instantiates a new QueryTenantMigrationValidationStatusResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportLocation

`func (o *QueryTenantMigrationValidationStatusResult) GetReportLocation() DataLocator`

GetReportLocation returns the ReportLocation field if non-nil, zero value otherwise.

### GetReportLocationOk

`func (o *QueryTenantMigrationValidationStatusResult) GetReportLocationOk() (*DataLocator, bool)`

GetReportLocationOk returns a tuple with the ReportLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportLocation

`func (o *QueryTenantMigrationValidationStatusResult) SetReportLocation(v DataLocator)`

SetReportLocation sets ReportLocation field to given value.

### HasReportLocation

`func (o *QueryTenantMigrationValidationStatusResult) HasReportLocation() bool`

HasReportLocation returns a boolean if a field has been set.

### GetStatus

`func (o *QueryTenantMigrationValidationStatusResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryTenantMigrationValidationStatusResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryTenantMigrationValidationStatusResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QueryTenantMigrationValidationStatusResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


