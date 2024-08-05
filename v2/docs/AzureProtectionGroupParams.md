# AzureProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentProtectionTypeParams** | Pointer to [**AzureAgentProtectionGroupParams**](AzureAgentProtectionGroupParams.md) |  | [optional] 
**NativeProtectionTypeParams** | Pointer to [**AzureNativeProtectionGroupParams**](AzureNativeProtectionGroupParams.md) |  | [optional] 
**ProtectionType** | **string** | Specifies the Azure Protection Group type. | 
**SnapshotManagerProtectionTypeParams** | Pointer to [**AzureSnapshotManagerProtectionGroupParams**](AzureSnapshotManagerProtectionGroupParams.md) |  | [optional] 

## Methods

### NewAzureProtectionGroupParams

`func NewAzureProtectionGroupParams(protectionType string, ) *AzureProtectionGroupParams`

NewAzureProtectionGroupParams instantiates a new AzureProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureProtectionGroupParamsWithDefaults

`func NewAzureProtectionGroupParamsWithDefaults() *AzureProtectionGroupParams`

NewAzureProtectionGroupParamsWithDefaults instantiates a new AzureProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentProtectionTypeParams

`func (o *AzureProtectionGroupParams) GetAgentProtectionTypeParams() AzureAgentProtectionGroupParams`

GetAgentProtectionTypeParams returns the AgentProtectionTypeParams field if non-nil, zero value otherwise.

### GetAgentProtectionTypeParamsOk

`func (o *AzureProtectionGroupParams) GetAgentProtectionTypeParamsOk() (*AzureAgentProtectionGroupParams, bool)`

GetAgentProtectionTypeParamsOk returns a tuple with the AgentProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentProtectionTypeParams

`func (o *AzureProtectionGroupParams) SetAgentProtectionTypeParams(v AzureAgentProtectionGroupParams)`

SetAgentProtectionTypeParams sets AgentProtectionTypeParams field to given value.

### HasAgentProtectionTypeParams

`func (o *AzureProtectionGroupParams) HasAgentProtectionTypeParams() bool`

HasAgentProtectionTypeParams returns a boolean if a field has been set.

### GetNativeProtectionTypeParams

`func (o *AzureProtectionGroupParams) GetNativeProtectionTypeParams() AzureNativeProtectionGroupParams`

GetNativeProtectionTypeParams returns the NativeProtectionTypeParams field if non-nil, zero value otherwise.

### GetNativeProtectionTypeParamsOk

`func (o *AzureProtectionGroupParams) GetNativeProtectionTypeParamsOk() (*AzureNativeProtectionGroupParams, bool)`

GetNativeProtectionTypeParamsOk returns a tuple with the NativeProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeProtectionTypeParams

`func (o *AzureProtectionGroupParams) SetNativeProtectionTypeParams(v AzureNativeProtectionGroupParams)`

SetNativeProtectionTypeParams sets NativeProtectionTypeParams field to given value.

### HasNativeProtectionTypeParams

`func (o *AzureProtectionGroupParams) HasNativeProtectionTypeParams() bool`

HasNativeProtectionTypeParams returns a boolean if a field has been set.

### GetProtectionType

`func (o *AzureProtectionGroupParams) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *AzureProtectionGroupParams) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *AzureProtectionGroupParams) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.


### GetSnapshotManagerProtectionTypeParams

`func (o *AzureProtectionGroupParams) GetSnapshotManagerProtectionTypeParams() AzureSnapshotManagerProtectionGroupParams`

GetSnapshotManagerProtectionTypeParams returns the SnapshotManagerProtectionTypeParams field if non-nil, zero value otherwise.

### GetSnapshotManagerProtectionTypeParamsOk

`func (o *AzureProtectionGroupParams) GetSnapshotManagerProtectionTypeParamsOk() (*AzureSnapshotManagerProtectionGroupParams, bool)`

GetSnapshotManagerProtectionTypeParamsOk returns a tuple with the SnapshotManagerProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotManagerProtectionTypeParams

`func (o *AzureProtectionGroupParams) SetSnapshotManagerProtectionTypeParams(v AzureSnapshotManagerProtectionGroupParams)`

SetSnapshotManagerProtectionTypeParams sets SnapshotManagerProtectionTypeParams field to given value.

### HasSnapshotManagerProtectionTypeParams

`func (o *AzureProtectionGroupParams) HasSnapshotManagerProtectionTypeParams() bool`

HasSnapshotManagerProtectionTypeParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


