# GraphNodeFilterParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AadParams** | Pointer to [**NullableAadGraphNodeFilterParams**](AadGraphNodeFilterParams.md) | Specifies the aad adapter specific node filter params. | [optional] 
**Name** | Pointer to **NullableString** | Filters the nodes based on provided current node display name. | [optional] 
**RootOnly** | Pointer to **NullableBool** | If set to true only root nodes would be returned. A root node refers to nodes in the graph with no incoming edges. Defaults to false. | [optional] 

## Methods

### NewGraphNodeFilterParams

`func NewGraphNodeFilterParams() *GraphNodeFilterParams`

NewGraphNodeFilterParams instantiates a new GraphNodeFilterParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphNodeFilterParamsWithDefaults

`func NewGraphNodeFilterParamsWithDefaults() *GraphNodeFilterParams`

NewGraphNodeFilterParamsWithDefaults instantiates a new GraphNodeFilterParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAadParams

`func (o *GraphNodeFilterParams) GetAadParams() AadGraphNodeFilterParams`

GetAadParams returns the AadParams field if non-nil, zero value otherwise.

### GetAadParamsOk

`func (o *GraphNodeFilterParams) GetAadParamsOk() (*AadGraphNodeFilterParams, bool)`

GetAadParamsOk returns a tuple with the AadParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAadParams

`func (o *GraphNodeFilterParams) SetAadParams(v AadGraphNodeFilterParams)`

SetAadParams sets AadParams field to given value.

### HasAadParams

`func (o *GraphNodeFilterParams) HasAadParams() bool`

HasAadParams returns a boolean if a field has been set.

### SetAadParamsNil

`func (o *GraphNodeFilterParams) SetAadParamsNil(b bool)`

 SetAadParamsNil sets the value for AadParams to be an explicit nil

### UnsetAadParams
`func (o *GraphNodeFilterParams) UnsetAadParams()`

UnsetAadParams ensures that no value is present for AadParams, not even an explicit nil
### GetName

`func (o *GraphNodeFilterParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GraphNodeFilterParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GraphNodeFilterParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GraphNodeFilterParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GraphNodeFilterParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GraphNodeFilterParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRootOnly

`func (o *GraphNodeFilterParams) GetRootOnly() bool`

GetRootOnly returns the RootOnly field if non-nil, zero value otherwise.

### GetRootOnlyOk

`func (o *GraphNodeFilterParams) GetRootOnlyOk() (*bool, bool)`

GetRootOnlyOk returns a tuple with the RootOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootOnly

`func (o *GraphNodeFilterParams) SetRootOnly(v bool)`

SetRootOnly sets RootOnly field to given value.

### HasRootOnly

`func (o *GraphNodeFilterParams) HasRootOnly() bool`

HasRootOnly returns a boolean if a field has been set.

### SetRootOnlyNil

`func (o *GraphNodeFilterParams) SetRootOnlyNil(b bool)`

 SetRootOnlyNil sets the value for RootOnly to be an explicit nil

### UnsetRootOnly
`func (o *GraphNodeFilterParams) UnsetRootOnly()`

UnsetRootOnly ensures that no value is present for RootOnly, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


