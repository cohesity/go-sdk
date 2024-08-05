# QueryExportMetadataStatusResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportLocation** | Pointer to [**DataLocator**](DataLocator.md) |  | [optional] 
**Status** | Pointer to **string** | The status of the export task. | [optional] 

## Methods

### NewQueryExportMetadataStatusResult

`func NewQueryExportMetadataStatusResult() *QueryExportMetadataStatusResult`

NewQueryExportMetadataStatusResult instantiates a new QueryExportMetadataStatusResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryExportMetadataStatusResultWithDefaults

`func NewQueryExportMetadataStatusResultWithDefaults() *QueryExportMetadataStatusResult`

NewQueryExportMetadataStatusResultWithDefaults instantiates a new QueryExportMetadataStatusResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportLocation

`func (o *QueryExportMetadataStatusResult) GetExportLocation() DataLocator`

GetExportLocation returns the ExportLocation field if non-nil, zero value otherwise.

### GetExportLocationOk

`func (o *QueryExportMetadataStatusResult) GetExportLocationOk() (*DataLocator, bool)`

GetExportLocationOk returns a tuple with the ExportLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportLocation

`func (o *QueryExportMetadataStatusResult) SetExportLocation(v DataLocator)`

SetExportLocation sets ExportLocation field to given value.

### HasExportLocation

`func (o *QueryExportMetadataStatusResult) HasExportLocation() bool`

HasExportLocation returns a boolean if a field has been set.

### GetStatus

`func (o *QueryExportMetadataStatusResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryExportMetadataStatusResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryExportMetadataStatusResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QueryExportMetadataStatusResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


