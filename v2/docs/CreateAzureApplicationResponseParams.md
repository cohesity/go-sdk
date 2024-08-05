# CreateAzureApplicationResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Microsoft365AppCredentialsList** | Pointer to [**[]Office365AppCredentials**](Office365AppCredentials.md) | Specifies a list of Microsoft365 azure application credentials needed to authenticate &amp; authorize users for Office 365/Azure Workflows. | [optional] 

## Methods

### NewCreateAzureApplicationResponseParams

`func NewCreateAzureApplicationResponseParams() *CreateAzureApplicationResponseParams`

NewCreateAzureApplicationResponseParams instantiates a new CreateAzureApplicationResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAzureApplicationResponseParamsWithDefaults

`func NewCreateAzureApplicationResponseParamsWithDefaults() *CreateAzureApplicationResponseParams`

NewCreateAzureApplicationResponseParamsWithDefaults instantiates a new CreateAzureApplicationResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMicrosoft365AppCredentialsList

`func (o *CreateAzureApplicationResponseParams) GetMicrosoft365AppCredentialsList() []Office365AppCredentials`

GetMicrosoft365AppCredentialsList returns the Microsoft365AppCredentialsList field if non-nil, zero value otherwise.

### GetMicrosoft365AppCredentialsListOk

`func (o *CreateAzureApplicationResponseParams) GetMicrosoft365AppCredentialsListOk() (*[]Office365AppCredentials, bool)`

GetMicrosoft365AppCredentialsListOk returns a tuple with the Microsoft365AppCredentialsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoft365AppCredentialsList

`func (o *CreateAzureApplicationResponseParams) SetMicrosoft365AppCredentialsList(v []Office365AppCredentials)`

SetMicrosoft365AppCredentialsList sets Microsoft365AppCredentialsList field to given value.

### HasMicrosoft365AppCredentialsList

`func (o *CreateAzureApplicationResponseParams) HasMicrosoft365AppCredentialsList() bool`

HasMicrosoft365AppCredentialsList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


