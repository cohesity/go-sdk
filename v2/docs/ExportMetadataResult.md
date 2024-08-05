# ExportMetadataResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportOperationId** | Pointer to **string** | The exporting is a time taking process. This denotes the id for this export task which will be used to poll the status of exporting metadata through QueryExportMetadataStatus. This will not be filled if there is an error. | [optional] 

## Methods

### NewExportMetadataResult

`func NewExportMetadataResult() *ExportMetadataResult`

NewExportMetadataResult instantiates a new ExportMetadataResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportMetadataResultWithDefaults

`func NewExportMetadataResultWithDefaults() *ExportMetadataResult`

NewExportMetadataResultWithDefaults instantiates a new ExportMetadataResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportOperationId

`func (o *ExportMetadataResult) GetExportOperationId() string`

GetExportOperationId returns the ExportOperationId field if non-nil, zero value otherwise.

### GetExportOperationIdOk

`func (o *ExportMetadataResult) GetExportOperationIdOk() (*string, bool)`

GetExportOperationIdOk returns a tuple with the ExportOperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportOperationId

`func (o *ExportMetadataResult) SetExportOperationId(v string)`

SetExportOperationId sets ExportOperationId field to given value.

### HasExportOperationId

`func (o *ExportMetadataResult) HasExportOperationId() bool`

HasExportOperationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


