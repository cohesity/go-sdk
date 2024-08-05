# ImportMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImportLocation** | Pointer to [**DataLocator**](DataLocator.md) |  | [optional] 

## Methods

### NewImportMetadataRequest

`func NewImportMetadataRequest() *ImportMetadataRequest`

NewImportMetadataRequest instantiates a new ImportMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportMetadataRequestWithDefaults

`func NewImportMetadataRequestWithDefaults() *ImportMetadataRequest`

NewImportMetadataRequestWithDefaults instantiates a new ImportMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImportLocation

`func (o *ImportMetadataRequest) GetImportLocation() DataLocator`

GetImportLocation returns the ImportLocation field if non-nil, zero value otherwise.

### GetImportLocationOk

`func (o *ImportMetadataRequest) GetImportLocationOk() (*DataLocator, bool)`

GetImportLocationOk returns a tuple with the ImportLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportLocation

`func (o *ImportMetadataRequest) SetImportLocation(v DataLocator)`

SetImportLocation sets ImportLocation field to given value.

### HasImportLocation

`func (o *ImportMetadataRequest) HasImportLocation() bool`

HasImportLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


