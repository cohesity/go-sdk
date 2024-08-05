# AzureSqlMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataType** | **NullableString** | Specifies the type of metadata being sent in the request. | 
**StandardCredentials** | Pointer to [**Credentials**](Credentials.md) |  | [optional] 

## Methods

### NewAzureSqlMetadata

`func NewAzureSqlMetadata(metadataType NullableString, ) *AzureSqlMetadata`

NewAzureSqlMetadata instantiates a new AzureSqlMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureSqlMetadataWithDefaults

`func NewAzureSqlMetadataWithDefaults() *AzureSqlMetadata`

NewAzureSqlMetadataWithDefaults instantiates a new AzureSqlMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataType

`func (o *AzureSqlMetadata) GetMetadataType() string`

GetMetadataType returns the MetadataType field if non-nil, zero value otherwise.

### GetMetadataTypeOk

`func (o *AzureSqlMetadata) GetMetadataTypeOk() (*string, bool)`

GetMetadataTypeOk returns a tuple with the MetadataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataType

`func (o *AzureSqlMetadata) SetMetadataType(v string)`

SetMetadataType sets MetadataType field to given value.


### SetMetadataTypeNil

`func (o *AzureSqlMetadata) SetMetadataTypeNil(b bool)`

 SetMetadataTypeNil sets the value for MetadataType to be an explicit nil

### UnsetMetadataType
`func (o *AzureSqlMetadata) UnsetMetadataType()`

UnsetMetadataType ensures that no value is present for MetadataType, not even an explicit nil
### GetStandardCredentials

`func (o *AzureSqlMetadata) GetStandardCredentials() Credentials`

GetStandardCredentials returns the StandardCredentials field if non-nil, zero value otherwise.

### GetStandardCredentialsOk

`func (o *AzureSqlMetadata) GetStandardCredentialsOk() (*Credentials, bool)`

GetStandardCredentialsOk returns a tuple with the StandardCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardCredentials

`func (o *AzureSqlMetadata) SetStandardCredentials(v Credentials)`

SetStandardCredentials sets StandardCredentials field to given value.

### HasStandardCredentials

`func (o *AzureSqlMetadata) HasStandardCredentials() bool`

HasStandardCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


