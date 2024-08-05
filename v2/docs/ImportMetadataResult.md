# ImportMetadataResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImportOperationId** | Pointer to **string** | The importing is a time taking process. This denotes the id for this import task which will be used to poll the status of importing metadata through QueryImportMetadataStatus. This will not be filled if there is an error. | [optional] 

## Methods

### NewImportMetadataResult

`func NewImportMetadataResult() *ImportMetadataResult`

NewImportMetadataResult instantiates a new ImportMetadataResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportMetadataResultWithDefaults

`func NewImportMetadataResultWithDefaults() *ImportMetadataResult`

NewImportMetadataResultWithDefaults instantiates a new ImportMetadataResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImportOperationId

`func (o *ImportMetadataResult) GetImportOperationId() string`

GetImportOperationId returns the ImportOperationId field if non-nil, zero value otherwise.

### GetImportOperationIdOk

`func (o *ImportMetadataResult) GetImportOperationIdOk() (*string, bool)`

GetImportOperationIdOk returns a tuple with the ImportOperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportOperationId

`func (o *ImportMetadataResult) SetImportOperationId(v string)`

SetImportOperationId sets ImportOperationId field to given value.

### HasImportOperationId

`func (o *ImportMetadataResult) HasImportOperationId() bool`

HasImportOperationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


