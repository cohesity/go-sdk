# AzureCoolBlobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **NullableString** | Specifies the category of the external target. | 
**FunctionAppDeploymentKey** | Pointer to **NullableString** | Specifies the access key to deploy Azure function to function app | [optional] 
**FunctionAppName** | Pointer to **NullableString** | Specifies the name of the Azure function app, which is the host of Azure functions. | [optional] 

## Methods

### NewAzureCoolBlobParams

`func NewAzureCoolBlobParams(category NullableString, ) *AzureCoolBlobParams`

NewAzureCoolBlobParams instantiates a new AzureCoolBlobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCoolBlobParamsWithDefaults

`func NewAzureCoolBlobParamsWithDefaults() *AzureCoolBlobParams`

NewAzureCoolBlobParamsWithDefaults instantiates a new AzureCoolBlobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *AzureCoolBlobParams) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AzureCoolBlobParams) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AzureCoolBlobParams) SetCategory(v string)`

SetCategory sets Category field to given value.


### SetCategoryNil

`func (o *AzureCoolBlobParams) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *AzureCoolBlobParams) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetFunctionAppDeploymentKey

`func (o *AzureCoolBlobParams) GetFunctionAppDeploymentKey() string`

GetFunctionAppDeploymentKey returns the FunctionAppDeploymentKey field if non-nil, zero value otherwise.

### GetFunctionAppDeploymentKeyOk

`func (o *AzureCoolBlobParams) GetFunctionAppDeploymentKeyOk() (*string, bool)`

GetFunctionAppDeploymentKeyOk returns a tuple with the FunctionAppDeploymentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionAppDeploymentKey

`func (o *AzureCoolBlobParams) SetFunctionAppDeploymentKey(v string)`

SetFunctionAppDeploymentKey sets FunctionAppDeploymentKey field to given value.

### HasFunctionAppDeploymentKey

`func (o *AzureCoolBlobParams) HasFunctionAppDeploymentKey() bool`

HasFunctionAppDeploymentKey returns a boolean if a field has been set.

### SetFunctionAppDeploymentKeyNil

`func (o *AzureCoolBlobParams) SetFunctionAppDeploymentKeyNil(b bool)`

 SetFunctionAppDeploymentKeyNil sets the value for FunctionAppDeploymentKey to be an explicit nil

### UnsetFunctionAppDeploymentKey
`func (o *AzureCoolBlobParams) UnsetFunctionAppDeploymentKey()`

UnsetFunctionAppDeploymentKey ensures that no value is present for FunctionAppDeploymentKey, not even an explicit nil
### GetFunctionAppName

`func (o *AzureCoolBlobParams) GetFunctionAppName() string`

GetFunctionAppName returns the FunctionAppName field if non-nil, zero value otherwise.

### GetFunctionAppNameOk

`func (o *AzureCoolBlobParams) GetFunctionAppNameOk() (*string, bool)`

GetFunctionAppNameOk returns a tuple with the FunctionAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionAppName

`func (o *AzureCoolBlobParams) SetFunctionAppName(v string)`

SetFunctionAppName sets FunctionAppName field to given value.

### HasFunctionAppName

`func (o *AzureCoolBlobParams) HasFunctionAppName() bool`

HasFunctionAppName returns a boolean if a field has been set.

### SetFunctionAppNameNil

`func (o *AzureCoolBlobParams) SetFunctionAppNameNil(b bool)`

 SetFunctionAppNameNil sets the value for FunctionAppName to be an explicit nil

### UnsetFunctionAppName
`func (o *AzureCoolBlobParams) UnsetFunctionAppName()`

UnsetFunctionAppName ensures that no value is present for FunctionAppName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


