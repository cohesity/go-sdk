# CreateAzureApplicationRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **NullableString** | Specifies the access token for Azure Application access. | 
**AppCount** | **int32** | Specifies the count of Azure application to be created. | 
**AzureTenantId** | Pointer to **NullableString** | Specifies the Azure Active Directory tenant ID or domain name. | [optional] 
**ExistingMicrosoft365AppCredentialsList** | Pointer to [**[]Office365AppCredentials**](Office365AppCredentials.md) | Specifies a list of Microsoft365 azure application credentials already added within the Microsoft365 source. | [optional] 
**Microsoft365Region** | Pointer to **NullableString** | Specifies the region where Office 365 Exchange environment is. | [optional] 
**UseCases** | Pointer to **[]string** | The usecases for which the application is to be created. | [optional] 
**Username** | Pointer to **NullableString** | Specifies the username to access Microsoft365 source. | [optional] 

## Methods

### NewCreateAzureApplicationRequestParams

`func NewCreateAzureApplicationRequestParams(accessToken NullableString, appCount int32, ) *CreateAzureApplicationRequestParams`

NewCreateAzureApplicationRequestParams instantiates a new CreateAzureApplicationRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAzureApplicationRequestParamsWithDefaults

`func NewCreateAzureApplicationRequestParamsWithDefaults() *CreateAzureApplicationRequestParams`

NewCreateAzureApplicationRequestParamsWithDefaults instantiates a new CreateAzureApplicationRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *CreateAzureApplicationRequestParams) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *CreateAzureApplicationRequestParams) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *CreateAzureApplicationRequestParams) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### SetAccessTokenNil

`func (o *CreateAzureApplicationRequestParams) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *CreateAzureApplicationRequestParams) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetAppCount

`func (o *CreateAzureApplicationRequestParams) GetAppCount() int32`

GetAppCount returns the AppCount field if non-nil, zero value otherwise.

### GetAppCountOk

`func (o *CreateAzureApplicationRequestParams) GetAppCountOk() (*int32, bool)`

GetAppCountOk returns a tuple with the AppCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCount

`func (o *CreateAzureApplicationRequestParams) SetAppCount(v int32)`

SetAppCount sets AppCount field to given value.


### GetAzureTenantId

`func (o *CreateAzureApplicationRequestParams) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *CreateAzureApplicationRequestParams) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *CreateAzureApplicationRequestParams) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *CreateAzureApplicationRequestParams) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### SetAzureTenantIdNil

`func (o *CreateAzureApplicationRequestParams) SetAzureTenantIdNil(b bool)`

 SetAzureTenantIdNil sets the value for AzureTenantId to be an explicit nil

### UnsetAzureTenantId
`func (o *CreateAzureApplicationRequestParams) UnsetAzureTenantId()`

UnsetAzureTenantId ensures that no value is present for AzureTenantId, not even an explicit nil
### GetExistingMicrosoft365AppCredentialsList

`func (o *CreateAzureApplicationRequestParams) GetExistingMicrosoft365AppCredentialsList() []Office365AppCredentials`

GetExistingMicrosoft365AppCredentialsList returns the ExistingMicrosoft365AppCredentialsList field if non-nil, zero value otherwise.

### GetExistingMicrosoft365AppCredentialsListOk

`func (o *CreateAzureApplicationRequestParams) GetExistingMicrosoft365AppCredentialsListOk() (*[]Office365AppCredentials, bool)`

GetExistingMicrosoft365AppCredentialsListOk returns a tuple with the ExistingMicrosoft365AppCredentialsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingMicrosoft365AppCredentialsList

`func (o *CreateAzureApplicationRequestParams) SetExistingMicrosoft365AppCredentialsList(v []Office365AppCredentials)`

SetExistingMicrosoft365AppCredentialsList sets ExistingMicrosoft365AppCredentialsList field to given value.

### HasExistingMicrosoft365AppCredentialsList

`func (o *CreateAzureApplicationRequestParams) HasExistingMicrosoft365AppCredentialsList() bool`

HasExistingMicrosoft365AppCredentialsList returns a boolean if a field has been set.

### GetMicrosoft365Region

`func (o *CreateAzureApplicationRequestParams) GetMicrosoft365Region() string`

GetMicrosoft365Region returns the Microsoft365Region field if non-nil, zero value otherwise.

### GetMicrosoft365RegionOk

`func (o *CreateAzureApplicationRequestParams) GetMicrosoft365RegionOk() (*string, bool)`

GetMicrosoft365RegionOk returns a tuple with the Microsoft365Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoft365Region

`func (o *CreateAzureApplicationRequestParams) SetMicrosoft365Region(v string)`

SetMicrosoft365Region sets Microsoft365Region field to given value.

### HasMicrosoft365Region

`func (o *CreateAzureApplicationRequestParams) HasMicrosoft365Region() bool`

HasMicrosoft365Region returns a boolean if a field has been set.

### SetMicrosoft365RegionNil

`func (o *CreateAzureApplicationRequestParams) SetMicrosoft365RegionNil(b bool)`

 SetMicrosoft365RegionNil sets the value for Microsoft365Region to be an explicit nil

### UnsetMicrosoft365Region
`func (o *CreateAzureApplicationRequestParams) UnsetMicrosoft365Region()`

UnsetMicrosoft365Region ensures that no value is present for Microsoft365Region, not even an explicit nil
### GetUseCases

`func (o *CreateAzureApplicationRequestParams) GetUseCases() []string`

GetUseCases returns the UseCases field if non-nil, zero value otherwise.

### GetUseCasesOk

`func (o *CreateAzureApplicationRequestParams) GetUseCasesOk() (*[]string, bool)`

GetUseCasesOk returns a tuple with the UseCases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCases

`func (o *CreateAzureApplicationRequestParams) SetUseCases(v []string)`

SetUseCases sets UseCases field to given value.

### HasUseCases

`func (o *CreateAzureApplicationRequestParams) HasUseCases() bool`

HasUseCases returns a boolean if a field has been set.

### SetUseCasesNil

`func (o *CreateAzureApplicationRequestParams) SetUseCasesNil(b bool)`

 SetUseCasesNil sets the value for UseCases to be an explicit nil

### UnsetUseCases
`func (o *CreateAzureApplicationRequestParams) UnsetUseCases()`

UnsetUseCases ensures that no value is present for UseCases, not even an explicit nil
### GetUsername

`func (o *CreateAzureApplicationRequestParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateAzureApplicationRequestParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateAzureApplicationRequestParams) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateAzureApplicationRequestParams) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *CreateAzureApplicationRequestParams) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *CreateAzureApplicationRequestParams) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


