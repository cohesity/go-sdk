# AzureEntraIDObjectProtectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludedObjectTypes** | Pointer to **[]string** | Specifies the list of object types to be excluded from protection. | [optional] 
**Objects** | Pointer to [**[]AzureObjectLevelParams**](AzureObjectLevelParams.md) | Specifies the objects to be protected. | [optional] 

## Methods

### NewAzureEntraIDObjectProtectionParams

`func NewAzureEntraIDObjectProtectionParams() *AzureEntraIDObjectProtectionParams`

NewAzureEntraIDObjectProtectionParams instantiates a new AzureEntraIDObjectProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureEntraIDObjectProtectionParamsWithDefaults

`func NewAzureEntraIDObjectProtectionParamsWithDefaults() *AzureEntraIDObjectProtectionParams`

NewAzureEntraIDObjectProtectionParamsWithDefaults instantiates a new AzureEntraIDObjectProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludedObjectTypes

`func (o *AzureEntraIDObjectProtectionParams) GetExcludedObjectTypes() []string`

GetExcludedObjectTypes returns the ExcludedObjectTypes field if non-nil, zero value otherwise.

### GetExcludedObjectTypesOk

`func (o *AzureEntraIDObjectProtectionParams) GetExcludedObjectTypesOk() (*[]string, bool)`

GetExcludedObjectTypesOk returns a tuple with the ExcludedObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedObjectTypes

`func (o *AzureEntraIDObjectProtectionParams) SetExcludedObjectTypes(v []string)`

SetExcludedObjectTypes sets ExcludedObjectTypes field to given value.

### HasExcludedObjectTypes

`func (o *AzureEntraIDObjectProtectionParams) HasExcludedObjectTypes() bool`

HasExcludedObjectTypes returns a boolean if a field has been set.

### SetExcludedObjectTypesNil

`func (o *AzureEntraIDObjectProtectionParams) SetExcludedObjectTypesNil(b bool)`

 SetExcludedObjectTypesNil sets the value for ExcludedObjectTypes to be an explicit nil

### UnsetExcludedObjectTypes
`func (o *AzureEntraIDObjectProtectionParams) UnsetExcludedObjectTypes()`

UnsetExcludedObjectTypes ensures that no value is present for ExcludedObjectTypes, not even an explicit nil
### GetObjects

`func (o *AzureEntraIDObjectProtectionParams) GetObjects() []AzureObjectLevelParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *AzureEntraIDObjectProtectionParams) GetObjectsOk() (*[]AzureObjectLevelParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *AzureEntraIDObjectProtectionParams) SetObjects(v []AzureObjectLevelParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *AzureEntraIDObjectProtectionParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


