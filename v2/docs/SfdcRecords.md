# SfdcRecords

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnNames** | Pointer to **[]string** | Specifies the column names for the records. | [optional] 
**Records** | Pointer to **[][]string** | Each record is represented by an array of strings having the same order as the &#39;columnNames&#39;. | [optional] 

## Methods

### NewSfdcRecords

`func NewSfdcRecords() *SfdcRecords`

NewSfdcRecords instantiates a new SfdcRecords object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSfdcRecordsWithDefaults

`func NewSfdcRecordsWithDefaults() *SfdcRecords`

NewSfdcRecordsWithDefaults instantiates a new SfdcRecords object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnNames

`func (o *SfdcRecords) GetColumnNames() []string`

GetColumnNames returns the ColumnNames field if non-nil, zero value otherwise.

### GetColumnNamesOk

`func (o *SfdcRecords) GetColumnNamesOk() (*[]string, bool)`

GetColumnNamesOk returns a tuple with the ColumnNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnNames

`func (o *SfdcRecords) SetColumnNames(v []string)`

SetColumnNames sets ColumnNames field to given value.

### HasColumnNames

`func (o *SfdcRecords) HasColumnNames() bool`

HasColumnNames returns a boolean if a field has been set.

### GetRecords

`func (o *SfdcRecords) GetRecords() [][]string`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *SfdcRecords) GetRecordsOk() (*[][]string, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *SfdcRecords) SetRecords(v [][]string)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *SfdcRecords) HasRecords() bool`

HasRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


