# AwsPostgresEntityMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbEngineId** | Pointer to **string** | Specifies the engine id of the database which can be used to identify the type of database. | [optional] 
**MetadataList** | [**[]AwsPostgresMetadata**](AwsPostgresMetadata.md) | Specifies the metadata list. | 

## Methods

### NewAwsPostgresEntityMetadata

`func NewAwsPostgresEntityMetadata(metadataList []AwsPostgresMetadata, ) *AwsPostgresEntityMetadata`

NewAwsPostgresEntityMetadata instantiates a new AwsPostgresEntityMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsPostgresEntityMetadataWithDefaults

`func NewAwsPostgresEntityMetadataWithDefaults() *AwsPostgresEntityMetadata`

NewAwsPostgresEntityMetadataWithDefaults instantiates a new AwsPostgresEntityMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbEngineId

`func (o *AwsPostgresEntityMetadata) GetDbEngineId() string`

GetDbEngineId returns the DbEngineId field if non-nil, zero value otherwise.

### GetDbEngineIdOk

`func (o *AwsPostgresEntityMetadata) GetDbEngineIdOk() (*string, bool)`

GetDbEngineIdOk returns a tuple with the DbEngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbEngineId

`func (o *AwsPostgresEntityMetadata) SetDbEngineId(v string)`

SetDbEngineId sets DbEngineId field to given value.

### HasDbEngineId

`func (o *AwsPostgresEntityMetadata) HasDbEngineId() bool`

HasDbEngineId returns a boolean if a field has been set.

### GetMetadataList

`func (o *AwsPostgresEntityMetadata) GetMetadataList() []AwsPostgresMetadata`

GetMetadataList returns the MetadataList field if non-nil, zero value otherwise.

### GetMetadataListOk

`func (o *AwsPostgresEntityMetadata) GetMetadataListOk() (*[]AwsPostgresMetadata, bool)`

GetMetadataListOk returns a tuple with the MetadataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataList

`func (o *AwsPostgresEntityMetadata) SetMetadataList(v []AwsPostgresMetadata)`

SetMetadataList sets MetadataList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


