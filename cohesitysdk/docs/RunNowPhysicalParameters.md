# RunNowPhysicalParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataFilePath** | Pointer to **NullableString** | Specifies metadata file path during run-now requests for physical file based backups for some specific host entity. If specified, it will override any default metadata/directive file path set at the job level for the host. Also note that if the job default does not specify a metadata/directive file path for the host, then specifying this field for that host during run-now request will be rejected. | [optional] 

## Methods

### NewRunNowPhysicalParameters

`func NewRunNowPhysicalParameters() *RunNowPhysicalParameters`

NewRunNowPhysicalParameters instantiates a new RunNowPhysicalParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunNowPhysicalParametersWithDefaults

`func NewRunNowPhysicalParametersWithDefaults() *RunNowPhysicalParameters`

NewRunNowPhysicalParametersWithDefaults instantiates a new RunNowPhysicalParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataFilePath

`func (o *RunNowPhysicalParameters) GetMetadataFilePath() string`

GetMetadataFilePath returns the MetadataFilePath field if non-nil, zero value otherwise.

### GetMetadataFilePathOk

`func (o *RunNowPhysicalParameters) GetMetadataFilePathOk() (*string, bool)`

GetMetadataFilePathOk returns a tuple with the MetadataFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFilePath

`func (o *RunNowPhysicalParameters) SetMetadataFilePath(v string)`

SetMetadataFilePath sets MetadataFilePath field to given value.

### HasMetadataFilePath

`func (o *RunNowPhysicalParameters) HasMetadataFilePath() bool`

HasMetadataFilePath returns a boolean if a field has been set.

### SetMetadataFilePathNil

`func (o *RunNowPhysicalParameters) SetMetadataFilePathNil(b bool)`

 SetMetadataFilePathNil sets the value for MetadataFilePath to be an explicit nil

### UnsetMetadataFilePath
`func (o *RunNowPhysicalParameters) UnsetMetadataFilePath()`

UnsetMetadataFilePath ensures that no value is present for MetadataFilePath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


