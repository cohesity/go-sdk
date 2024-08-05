# CredentialsMetadataEnvironmentType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **string** | Specifies the types of environment that supports credential metadata. The environment type can have the following possible values kAwsAuroraPostgres - for aurora db protection environment. kAwsRDSPostgres - for RDS db protection environment. | [optional] 

## Methods

### NewCredentialsMetadataEnvironmentType

`func NewCredentialsMetadataEnvironmentType() *CredentialsMetadataEnvironmentType`

NewCredentialsMetadataEnvironmentType instantiates a new CredentialsMetadataEnvironmentType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsMetadataEnvironmentTypeWithDefaults

`func NewCredentialsMetadataEnvironmentTypeWithDefaults() *CredentialsMetadataEnvironmentType`

NewCredentialsMetadataEnvironmentTypeWithDefaults instantiates a new CredentialsMetadataEnvironmentType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *CredentialsMetadataEnvironmentType) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CredentialsMetadataEnvironmentType) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CredentialsMetadataEnvironmentType) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CredentialsMetadataEnvironmentType) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


