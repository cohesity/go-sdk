# AwsRdsEntityMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbEngineId** | Pointer to **string** | Specifies the engine id of the database which can be used to identify the type of database. | [optional] 
**MetadataList** | [**[]AwsPostgresMetadata**](AwsPostgresMetadata.md) | Specifies the metadata list. | 

## Methods

### NewAwsRdsEntityMetadata

`func NewAwsRdsEntityMetadata(metadataList []AwsPostgresMetadata, ) *AwsRdsEntityMetadata`

NewAwsRdsEntityMetadata instantiates a new AwsRdsEntityMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsRdsEntityMetadataWithDefaults

`func NewAwsRdsEntityMetadataWithDefaults() *AwsRdsEntityMetadata`

NewAwsRdsEntityMetadataWithDefaults instantiates a new AwsRdsEntityMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbEngineId

`func (o *AwsRdsEntityMetadata) GetDbEngineId() string`

GetDbEngineId returns the DbEngineId field if non-nil, zero value otherwise.

### GetDbEngineIdOk

`func (o *AwsRdsEntityMetadata) GetDbEngineIdOk() (*string, bool)`

GetDbEngineIdOk returns a tuple with the DbEngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbEngineId

`func (o *AwsRdsEntityMetadata) SetDbEngineId(v string)`

SetDbEngineId sets DbEngineId field to given value.

### HasDbEngineId

`func (o *AwsRdsEntityMetadata) HasDbEngineId() bool`

HasDbEngineId returns a boolean if a field has been set.

### GetMetadataList

`func (o *AwsRdsEntityMetadata) GetMetadataList() []AwsPostgresMetadata`

GetMetadataList returns the MetadataList field if non-nil, zero value otherwise.

### GetMetadataListOk

`func (o *AwsRdsEntityMetadata) GetMetadataListOk() (*[]AwsPostgresMetadata, bool)`

GetMetadataListOk returns a tuple with the MetadataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataList

`func (o *AwsRdsEntityMetadata) SetMetadataList(v []AwsPostgresMetadata)`

SetMetadataList sets MetadataList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


