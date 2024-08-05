# AadGraphNodeFilterParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeTypes** | Pointer to **[]string** | Filters the nodes which matches with specified aad node types provided. Supported AAD node types - Users/Groups/Applications/ AdministrativeUnits/ServicePrincipals/DirectoryRoles/Contacts/ Devices | [optional] 

## Methods

### NewAadGraphNodeFilterParams

`func NewAadGraphNodeFilterParams() *AadGraphNodeFilterParams`

NewAadGraphNodeFilterParams instantiates a new AadGraphNodeFilterParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAadGraphNodeFilterParamsWithDefaults

`func NewAadGraphNodeFilterParamsWithDefaults() *AadGraphNodeFilterParams`

NewAadGraphNodeFilterParamsWithDefaults instantiates a new AadGraphNodeFilterParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeTypes

`func (o *AadGraphNodeFilterParams) GetNodeTypes() []string`

GetNodeTypes returns the NodeTypes field if non-nil, zero value otherwise.

### GetNodeTypesOk

`func (o *AadGraphNodeFilterParams) GetNodeTypesOk() (*[]string, bool)`

GetNodeTypesOk returns a tuple with the NodeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeTypes

`func (o *AadGraphNodeFilterParams) SetNodeTypes(v []string)`

SetNodeTypes sets NodeTypes field to given value.

### HasNodeTypes

`func (o *AadGraphNodeFilterParams) HasNodeTypes() bool`

HasNodeTypes returns a boolean if a field has been set.

### SetNodeTypesNil

`func (o *AadGraphNodeFilterParams) SetNodeTypesNil(b bool)`

 SetNodeTypesNil sets the value for NodeTypes to be an explicit nil

### UnsetNodeTypes
`func (o *AadGraphNodeFilterParams) UnsetNodeTypes()`

UnsetNodeTypes ensures that no value is present for NodeTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


