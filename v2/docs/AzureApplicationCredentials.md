# AzureApplicationCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **NullableString** | Specifies the application id of an Azure active directory application. | 
**ApplicationObjectId** | Pointer to **NullableString** | Specifies the application object id of an Azure active directory application. | [optional] 
**EncryptedApplicationKey** | Pointer to **NullableString** | Specifies the encrypted application key of an Azure active directory application. | [optional] 

## Methods

### NewAzureApplicationCredentials

`func NewAzureApplicationCredentials(applicationId NullableString, ) *AzureApplicationCredentials`

NewAzureApplicationCredentials instantiates a new AzureApplicationCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureApplicationCredentialsWithDefaults

`func NewAzureApplicationCredentialsWithDefaults() *AzureApplicationCredentials`

NewAzureApplicationCredentialsWithDefaults instantiates a new AzureApplicationCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *AzureApplicationCredentials) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AzureApplicationCredentials) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AzureApplicationCredentials) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### SetApplicationIdNil

`func (o *AzureApplicationCredentials) SetApplicationIdNil(b bool)`

 SetApplicationIdNil sets the value for ApplicationId to be an explicit nil

### UnsetApplicationId
`func (o *AzureApplicationCredentials) UnsetApplicationId()`

UnsetApplicationId ensures that no value is present for ApplicationId, not even an explicit nil
### GetApplicationObjectId

`func (o *AzureApplicationCredentials) GetApplicationObjectId() string`

GetApplicationObjectId returns the ApplicationObjectId field if non-nil, zero value otherwise.

### GetApplicationObjectIdOk

`func (o *AzureApplicationCredentials) GetApplicationObjectIdOk() (*string, bool)`

GetApplicationObjectIdOk returns a tuple with the ApplicationObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationObjectId

`func (o *AzureApplicationCredentials) SetApplicationObjectId(v string)`

SetApplicationObjectId sets ApplicationObjectId field to given value.

### HasApplicationObjectId

`func (o *AzureApplicationCredentials) HasApplicationObjectId() bool`

HasApplicationObjectId returns a boolean if a field has been set.

### SetApplicationObjectIdNil

`func (o *AzureApplicationCredentials) SetApplicationObjectIdNil(b bool)`

 SetApplicationObjectIdNil sets the value for ApplicationObjectId to be an explicit nil

### UnsetApplicationObjectId
`func (o *AzureApplicationCredentials) UnsetApplicationObjectId()`

UnsetApplicationObjectId ensures that no value is present for ApplicationObjectId, not even an explicit nil
### GetEncryptedApplicationKey

`func (o *AzureApplicationCredentials) GetEncryptedApplicationKey() string`

GetEncryptedApplicationKey returns the EncryptedApplicationKey field if non-nil, zero value otherwise.

### GetEncryptedApplicationKeyOk

`func (o *AzureApplicationCredentials) GetEncryptedApplicationKeyOk() (*string, bool)`

GetEncryptedApplicationKeyOk returns a tuple with the EncryptedApplicationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedApplicationKey

`func (o *AzureApplicationCredentials) SetEncryptedApplicationKey(v string)`

SetEncryptedApplicationKey sets EncryptedApplicationKey field to given value.

### HasEncryptedApplicationKey

`func (o *AzureApplicationCredentials) HasEncryptedApplicationKey() bool`

HasEncryptedApplicationKey returns a boolean if a field has been set.

### SetEncryptedApplicationKeyNil

`func (o *AzureApplicationCredentials) SetEncryptedApplicationKeyNil(b bool)`

 SetEncryptedApplicationKeyNil sets the value for EncryptedApplicationKey to be an explicit nil

### UnsetEncryptedApplicationKey
`func (o *AzureApplicationCredentials) UnsetEncryptedApplicationKey()`

UnsetEncryptedApplicationKey ensures that no value is present for EncryptedApplicationKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


