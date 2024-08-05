# AzureObjectProtectionResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionType** | Pointer to **string** | Specifies the Azure Protection Job type. | [optional] 
**AzureEntraIdProtectionTypeParams** | Pointer to [**AzureEntraIDObjectProtectionParams**](AzureEntraIDObjectProtectionParams.md) |  | [optional] 
**AzureSqlProtectionTypeParams** | Pointer to [**AzureSqlObjectProtectionParams**](AzureSqlObjectProtectionParams.md) |  | [optional] 
**NativeProtectionTypeParams** | Pointer to [**AzureNativeObjectProtectionParams**](AzureNativeObjectProtectionParams.md) |  | [optional] 

## Methods

### NewAzureObjectProtectionResponseParams

`func NewAzureObjectProtectionResponseParams() *AzureObjectProtectionResponseParams`

NewAzureObjectProtectionResponseParams instantiates a new AzureObjectProtectionResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureObjectProtectionResponseParamsWithDefaults

`func NewAzureObjectProtectionResponseParamsWithDefaults() *AzureObjectProtectionResponseParams`

NewAzureObjectProtectionResponseParamsWithDefaults instantiates a new AzureObjectProtectionResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectionType

`func (o *AzureObjectProtectionResponseParams) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *AzureObjectProtectionResponseParams) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *AzureObjectProtectionResponseParams) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.

### HasProtectionType

`func (o *AzureObjectProtectionResponseParams) HasProtectionType() bool`

HasProtectionType returns a boolean if a field has been set.

### GetAzureEntraIdProtectionTypeParams

`func (o *AzureObjectProtectionResponseParams) GetAzureEntraIdProtectionTypeParams() AzureEntraIDObjectProtectionParams`

GetAzureEntraIdProtectionTypeParams returns the AzureEntraIdProtectionTypeParams field if non-nil, zero value otherwise.

### GetAzureEntraIdProtectionTypeParamsOk

`func (o *AzureObjectProtectionResponseParams) GetAzureEntraIdProtectionTypeParamsOk() (*AzureEntraIDObjectProtectionParams, bool)`

GetAzureEntraIdProtectionTypeParamsOk returns a tuple with the AzureEntraIdProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureEntraIdProtectionTypeParams

`func (o *AzureObjectProtectionResponseParams) SetAzureEntraIdProtectionTypeParams(v AzureEntraIDObjectProtectionParams)`

SetAzureEntraIdProtectionTypeParams sets AzureEntraIdProtectionTypeParams field to given value.

### HasAzureEntraIdProtectionTypeParams

`func (o *AzureObjectProtectionResponseParams) HasAzureEntraIdProtectionTypeParams() bool`

HasAzureEntraIdProtectionTypeParams returns a boolean if a field has been set.

### GetAzureSqlProtectionTypeParams

`func (o *AzureObjectProtectionResponseParams) GetAzureSqlProtectionTypeParams() AzureSqlObjectProtectionParams`

GetAzureSqlProtectionTypeParams returns the AzureSqlProtectionTypeParams field if non-nil, zero value otherwise.

### GetAzureSqlProtectionTypeParamsOk

`func (o *AzureObjectProtectionResponseParams) GetAzureSqlProtectionTypeParamsOk() (*AzureSqlObjectProtectionParams, bool)`

GetAzureSqlProtectionTypeParamsOk returns a tuple with the AzureSqlProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSqlProtectionTypeParams

`func (o *AzureObjectProtectionResponseParams) SetAzureSqlProtectionTypeParams(v AzureSqlObjectProtectionParams)`

SetAzureSqlProtectionTypeParams sets AzureSqlProtectionTypeParams field to given value.

### HasAzureSqlProtectionTypeParams

`func (o *AzureObjectProtectionResponseParams) HasAzureSqlProtectionTypeParams() bool`

HasAzureSqlProtectionTypeParams returns a boolean if a field has been set.

### GetNativeProtectionTypeParams

`func (o *AzureObjectProtectionResponseParams) GetNativeProtectionTypeParams() AzureNativeObjectProtectionParams`

GetNativeProtectionTypeParams returns the NativeProtectionTypeParams field if non-nil, zero value otherwise.

### GetNativeProtectionTypeParamsOk

`func (o *AzureObjectProtectionResponseParams) GetNativeProtectionTypeParamsOk() (*AzureNativeObjectProtectionParams, bool)`

GetNativeProtectionTypeParamsOk returns a tuple with the NativeProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeProtectionTypeParams

`func (o *AzureObjectProtectionResponseParams) SetNativeProtectionTypeParams(v AzureNativeObjectProtectionParams)`

SetNativeProtectionTypeParams sets NativeProtectionTypeParams field to given value.

### HasNativeProtectionTypeParams

`func (o *AzureObjectProtectionResponseParams) HasNativeProtectionTypeParams() bool`

HasNativeProtectionTypeParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


