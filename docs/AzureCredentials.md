# AzureCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **NullableString** | Specifies Application Id of the active directory of Azure account. | [optional] 
**ApplicationKey** | Pointer to **NullableString** | Specifies Application key of the active directory of Azure account. | [optional] 
**AzureType** | Pointer to **NullableString** | Specifies the entity type such as &#39;kSubscription&#39; if the environment is kAzure. Specifies the type of an Azure source entity. &#39;kSubscription&#39; indicates a billing unit within Azure account. &#39;kResourceGroup&#39; indicates a container that holds related resources. &#39;kVirtualMachine&#39; indicates a Virtual Machine in Azure environment. &#39;kStorageAccount&#39; represents a collection of storage containers. &#39;kStorageKey&#39; indicates a key required to access the storage account. &#39;kStorageContainer&#39; represents a storage container within a storage account. &#39;kStorageBlob&#39; represents a storage blog within a storage container. &#39;kStorageResourceGroup&#39; indicates a container that holds related storage resources. &#39;kNetworkSecurityGroup&#39; represents a network security group. &#39;kVirtualNetwork&#39; represents a virtual network. &#39;kNetworkResourceGroup&#39; indicates a container that holds related network resources. &#39;kSubnet&#39; represents a subnet within the virtual network. &#39;kComputeOptions&#39; indicates the number of CPU cores and memory size available for a type of a Virtual Machine. | [optional] 
**SubscriptionId** | Pointer to **NullableString** | Specifies Subscription id inside a customer&#39;s Azure account. It represents sub-section within the Azure account where a customer allows us to create VMs, storage account etc. | [optional] 
**SubscriptionType** | Pointer to **NullableString** | Specifies the subscription type of Azure such as &#39;kAzureCommercial&#39; or &#39;kAzureGovCloud&#39;. Specifies the subscription type of an Azure source entity. &#39;kAzureCommercial&#39; indicates a standard Azure subscription. &#39;kAzureGovCloud&#39; indicates a govt Azure subscription. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies Tenant Id of the active directory of Azure account. | [optional] 

## Methods

### NewAzureCredentials

`func NewAzureCredentials() *AzureCredentials`

NewAzureCredentials instantiates a new AzureCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCredentialsWithDefaults

`func NewAzureCredentialsWithDefaults() *AzureCredentials`

NewAzureCredentialsWithDefaults instantiates a new AzureCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *AzureCredentials) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AzureCredentials) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AzureCredentials) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *AzureCredentials) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationIdNil

`func (o *AzureCredentials) SetApplicationIdNil(b bool)`

 SetApplicationIdNil sets the value for ApplicationId to be an explicit nil

### UnsetApplicationId
`func (o *AzureCredentials) UnsetApplicationId()`

UnsetApplicationId ensures that no value is present for ApplicationId, not even an explicit nil
### GetApplicationKey

`func (o *AzureCredentials) GetApplicationKey() string`

GetApplicationKey returns the ApplicationKey field if non-nil, zero value otherwise.

### GetApplicationKeyOk

`func (o *AzureCredentials) GetApplicationKeyOk() (*string, bool)`

GetApplicationKeyOk returns a tuple with the ApplicationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationKey

`func (o *AzureCredentials) SetApplicationKey(v string)`

SetApplicationKey sets ApplicationKey field to given value.

### HasApplicationKey

`func (o *AzureCredentials) HasApplicationKey() bool`

HasApplicationKey returns a boolean if a field has been set.

### SetApplicationKeyNil

`func (o *AzureCredentials) SetApplicationKeyNil(b bool)`

 SetApplicationKeyNil sets the value for ApplicationKey to be an explicit nil

### UnsetApplicationKey
`func (o *AzureCredentials) UnsetApplicationKey()`

UnsetApplicationKey ensures that no value is present for ApplicationKey, not even an explicit nil
### GetAzureType

`func (o *AzureCredentials) GetAzureType() string`

GetAzureType returns the AzureType field if non-nil, zero value otherwise.

### GetAzureTypeOk

`func (o *AzureCredentials) GetAzureTypeOk() (*string, bool)`

GetAzureTypeOk returns a tuple with the AzureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureType

`func (o *AzureCredentials) SetAzureType(v string)`

SetAzureType sets AzureType field to given value.

### HasAzureType

`func (o *AzureCredentials) HasAzureType() bool`

HasAzureType returns a boolean if a field has been set.

### SetAzureTypeNil

`func (o *AzureCredentials) SetAzureTypeNil(b bool)`

 SetAzureTypeNil sets the value for AzureType to be an explicit nil

### UnsetAzureType
`func (o *AzureCredentials) UnsetAzureType()`

UnsetAzureType ensures that no value is present for AzureType, not even an explicit nil
### GetSubscriptionId

`func (o *AzureCredentials) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureCredentials) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureCredentials) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AzureCredentials) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *AzureCredentials) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *AzureCredentials) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetSubscriptionType

`func (o *AzureCredentials) GetSubscriptionType() string`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *AzureCredentials) GetSubscriptionTypeOk() (*string, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *AzureCredentials) SetSubscriptionType(v string)`

SetSubscriptionType sets SubscriptionType field to given value.

### HasSubscriptionType

`func (o *AzureCredentials) HasSubscriptionType() bool`

HasSubscriptionType returns a boolean if a field has been set.

### SetSubscriptionTypeNil

`func (o *AzureCredentials) SetSubscriptionTypeNil(b bool)`

 SetSubscriptionTypeNil sets the value for SubscriptionType to be an explicit nil

### UnsetSubscriptionType
`func (o *AzureCredentials) UnsetSubscriptionType()`

UnsetSubscriptionType ensures that no value is present for SubscriptionType, not even an explicit nil
### GetTenantId

`func (o *AzureCredentials) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureCredentials) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureCredentials) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AzureCredentials) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *AzureCredentials) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *AzureCredentials) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


