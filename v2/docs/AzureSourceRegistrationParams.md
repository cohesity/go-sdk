# AzureSourceRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationCredentials** | Pointer to [**[]AzureApplicationCredentials**](AzureApplicationCredentials.md) | Specifies the credentials for a list of applications from azure active directory. | [optional] 
**AzureTenantId** | Pointer to **NullableString** | Specifies Tenant Id of the active directory of Azure account. Accpets both Azure tanant Id and tenant domain name. | [optional] 
**GraphAccessToken** | Pointer to **NullableString** | Specifies the graph access token for using Azure graph API&#39;s. | [optional] 
**ManagementAccessToken** | Pointer to **NullableString** | Specifies the management access token for using Azure management API&#39;s. | [optional] 
**RegistrationLevel** | **NullableString** | Specifies whether the registration is at tenant level or subscription level. | 
**RegistrationWorkflow** | **NullableString** | Specifies whether the type of registration is express or manual. | 
**SubscriptionDetails** | Pointer to [**[]AzureSubscription**](AzureSubscription.md) | Specifies the list subscription ids to be registered. | [optional] 
**UseCases** | Pointer to **[]string** | The use cases for which the source is to be registered. | [optional] 

## Methods

### NewAzureSourceRegistrationParams

`func NewAzureSourceRegistrationParams(registrationLevel NullableString, registrationWorkflow NullableString, ) *AzureSourceRegistrationParams`

NewAzureSourceRegistrationParams instantiates a new AzureSourceRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureSourceRegistrationParamsWithDefaults

`func NewAzureSourceRegistrationParamsWithDefaults() *AzureSourceRegistrationParams`

NewAzureSourceRegistrationParamsWithDefaults instantiates a new AzureSourceRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationCredentials

`func (o *AzureSourceRegistrationParams) GetApplicationCredentials() []AzureApplicationCredentials`

GetApplicationCredentials returns the ApplicationCredentials field if non-nil, zero value otherwise.

### GetApplicationCredentialsOk

`func (o *AzureSourceRegistrationParams) GetApplicationCredentialsOk() (*[]AzureApplicationCredentials, bool)`

GetApplicationCredentialsOk returns a tuple with the ApplicationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCredentials

`func (o *AzureSourceRegistrationParams) SetApplicationCredentials(v []AzureApplicationCredentials)`

SetApplicationCredentials sets ApplicationCredentials field to given value.

### HasApplicationCredentials

`func (o *AzureSourceRegistrationParams) HasApplicationCredentials() bool`

HasApplicationCredentials returns a boolean if a field has been set.

### GetAzureTenantId

`func (o *AzureSourceRegistrationParams) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *AzureSourceRegistrationParams) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *AzureSourceRegistrationParams) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *AzureSourceRegistrationParams) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### SetAzureTenantIdNil

`func (o *AzureSourceRegistrationParams) SetAzureTenantIdNil(b bool)`

 SetAzureTenantIdNil sets the value for AzureTenantId to be an explicit nil

### UnsetAzureTenantId
`func (o *AzureSourceRegistrationParams) UnsetAzureTenantId()`

UnsetAzureTenantId ensures that no value is present for AzureTenantId, not even an explicit nil
### GetGraphAccessToken

`func (o *AzureSourceRegistrationParams) GetGraphAccessToken() string`

GetGraphAccessToken returns the GraphAccessToken field if non-nil, zero value otherwise.

### GetGraphAccessTokenOk

`func (o *AzureSourceRegistrationParams) GetGraphAccessTokenOk() (*string, bool)`

GetGraphAccessTokenOk returns a tuple with the GraphAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphAccessToken

`func (o *AzureSourceRegistrationParams) SetGraphAccessToken(v string)`

SetGraphAccessToken sets GraphAccessToken field to given value.

### HasGraphAccessToken

`func (o *AzureSourceRegistrationParams) HasGraphAccessToken() bool`

HasGraphAccessToken returns a boolean if a field has been set.

### SetGraphAccessTokenNil

`func (o *AzureSourceRegistrationParams) SetGraphAccessTokenNil(b bool)`

 SetGraphAccessTokenNil sets the value for GraphAccessToken to be an explicit nil

### UnsetGraphAccessToken
`func (o *AzureSourceRegistrationParams) UnsetGraphAccessToken()`

UnsetGraphAccessToken ensures that no value is present for GraphAccessToken, not even an explicit nil
### GetManagementAccessToken

`func (o *AzureSourceRegistrationParams) GetManagementAccessToken() string`

GetManagementAccessToken returns the ManagementAccessToken field if non-nil, zero value otherwise.

### GetManagementAccessTokenOk

`func (o *AzureSourceRegistrationParams) GetManagementAccessTokenOk() (*string, bool)`

GetManagementAccessTokenOk returns a tuple with the ManagementAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAccessToken

`func (o *AzureSourceRegistrationParams) SetManagementAccessToken(v string)`

SetManagementAccessToken sets ManagementAccessToken field to given value.

### HasManagementAccessToken

`func (o *AzureSourceRegistrationParams) HasManagementAccessToken() bool`

HasManagementAccessToken returns a boolean if a field has been set.

### SetManagementAccessTokenNil

`func (o *AzureSourceRegistrationParams) SetManagementAccessTokenNil(b bool)`

 SetManagementAccessTokenNil sets the value for ManagementAccessToken to be an explicit nil

### UnsetManagementAccessToken
`func (o *AzureSourceRegistrationParams) UnsetManagementAccessToken()`

UnsetManagementAccessToken ensures that no value is present for ManagementAccessToken, not even an explicit nil
### GetRegistrationLevel

`func (o *AzureSourceRegistrationParams) GetRegistrationLevel() string`

GetRegistrationLevel returns the RegistrationLevel field if non-nil, zero value otherwise.

### GetRegistrationLevelOk

`func (o *AzureSourceRegistrationParams) GetRegistrationLevelOk() (*string, bool)`

GetRegistrationLevelOk returns a tuple with the RegistrationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationLevel

`func (o *AzureSourceRegistrationParams) SetRegistrationLevel(v string)`

SetRegistrationLevel sets RegistrationLevel field to given value.


### SetRegistrationLevelNil

`func (o *AzureSourceRegistrationParams) SetRegistrationLevelNil(b bool)`

 SetRegistrationLevelNil sets the value for RegistrationLevel to be an explicit nil

### UnsetRegistrationLevel
`func (o *AzureSourceRegistrationParams) UnsetRegistrationLevel()`

UnsetRegistrationLevel ensures that no value is present for RegistrationLevel, not even an explicit nil
### GetRegistrationWorkflow

`func (o *AzureSourceRegistrationParams) GetRegistrationWorkflow() string`

GetRegistrationWorkflow returns the RegistrationWorkflow field if non-nil, zero value otherwise.

### GetRegistrationWorkflowOk

`func (o *AzureSourceRegistrationParams) GetRegistrationWorkflowOk() (*string, bool)`

GetRegistrationWorkflowOk returns a tuple with the RegistrationWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationWorkflow

`func (o *AzureSourceRegistrationParams) SetRegistrationWorkflow(v string)`

SetRegistrationWorkflow sets RegistrationWorkflow field to given value.


### SetRegistrationWorkflowNil

`func (o *AzureSourceRegistrationParams) SetRegistrationWorkflowNil(b bool)`

 SetRegistrationWorkflowNil sets the value for RegistrationWorkflow to be an explicit nil

### UnsetRegistrationWorkflow
`func (o *AzureSourceRegistrationParams) UnsetRegistrationWorkflow()`

UnsetRegistrationWorkflow ensures that no value is present for RegistrationWorkflow, not even an explicit nil
### GetSubscriptionDetails

`func (o *AzureSourceRegistrationParams) GetSubscriptionDetails() []AzureSubscription`

GetSubscriptionDetails returns the SubscriptionDetails field if non-nil, zero value otherwise.

### GetSubscriptionDetailsOk

`func (o *AzureSourceRegistrationParams) GetSubscriptionDetailsOk() (*[]AzureSubscription, bool)`

GetSubscriptionDetailsOk returns a tuple with the SubscriptionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionDetails

`func (o *AzureSourceRegistrationParams) SetSubscriptionDetails(v []AzureSubscription)`

SetSubscriptionDetails sets SubscriptionDetails field to given value.

### HasSubscriptionDetails

`func (o *AzureSourceRegistrationParams) HasSubscriptionDetails() bool`

HasSubscriptionDetails returns a boolean if a field has been set.

### GetUseCases

`func (o *AzureSourceRegistrationParams) GetUseCases() []string`

GetUseCases returns the UseCases field if non-nil, zero value otherwise.

### GetUseCasesOk

`func (o *AzureSourceRegistrationParams) GetUseCasesOk() (*[]string, bool)`

GetUseCasesOk returns a tuple with the UseCases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCases

`func (o *AzureSourceRegistrationParams) SetUseCases(v []string)`

SetUseCases sets UseCases field to given value.

### HasUseCases

`func (o *AzureSourceRegistrationParams) HasUseCases() bool`

HasUseCases returns a boolean if a field has been set.

### SetUseCasesNil

`func (o *AzureSourceRegistrationParams) SetUseCasesNil(b bool)`

 SetUseCasesNil sets the value for UseCases to be an explicit nil

### UnsetUseCases
`func (o *AzureSourceRegistrationParams) UnsetUseCases()`

UnsetUseCases ensures that no value is present for UseCases, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


