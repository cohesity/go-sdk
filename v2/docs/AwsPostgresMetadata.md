# AwsPostgresMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataType** | **NullableString** | Specifies the type of metadata being sent in the request. | 
**StandardCredentials** | Pointer to [**AwsCredentials**](AwsCredentials.md) |  | [optional] 

## Methods

### NewAwsPostgresMetadata

`func NewAwsPostgresMetadata(metadataType NullableString, ) *AwsPostgresMetadata`

NewAwsPostgresMetadata instantiates a new AwsPostgresMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsPostgresMetadataWithDefaults

`func NewAwsPostgresMetadataWithDefaults() *AwsPostgresMetadata`

NewAwsPostgresMetadataWithDefaults instantiates a new AwsPostgresMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataType

`func (o *AwsPostgresMetadata) GetMetadataType() string`

GetMetadataType returns the MetadataType field if non-nil, zero value otherwise.

### GetMetadataTypeOk

`func (o *AwsPostgresMetadata) GetMetadataTypeOk() (*string, bool)`

GetMetadataTypeOk returns a tuple with the MetadataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataType

`func (o *AwsPostgresMetadata) SetMetadataType(v string)`

SetMetadataType sets MetadataType field to given value.


### SetMetadataTypeNil

`func (o *AwsPostgresMetadata) SetMetadataTypeNil(b bool)`

 SetMetadataTypeNil sets the value for MetadataType to be an explicit nil

### UnsetMetadataType
`func (o *AwsPostgresMetadata) UnsetMetadataType()`

UnsetMetadataType ensures that no value is present for MetadataType, not even an explicit nil
### GetStandardCredentials

`func (o *AwsPostgresMetadata) GetStandardCredentials() AwsCredentials`

GetStandardCredentials returns the StandardCredentials field if non-nil, zero value otherwise.

### GetStandardCredentialsOk

`func (o *AwsPostgresMetadata) GetStandardCredentialsOk() (*AwsCredentials, bool)`

GetStandardCredentialsOk returns a tuple with the StandardCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardCredentials

`func (o *AwsPostgresMetadata) SetStandardCredentials(v AwsCredentials)`

SetStandardCredentials sets StandardCredentials field to given value.

### HasStandardCredentials

`func (o *AwsPostgresMetadata) HasStandardCredentials() bool`

HasStandardCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


