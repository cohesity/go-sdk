# CommonGraphNodeFilterParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Filters the nodes based on provided current node display name. | [optional] 
**RootOnly** | Pointer to **NullableBool** | If set to true only root nodes would be returned. A root node refers to nodes in the graph with no incoming edges. Defaults to false. | [optional] 

## Methods

### NewCommonGraphNodeFilterParams

`func NewCommonGraphNodeFilterParams() *CommonGraphNodeFilterParams`

NewCommonGraphNodeFilterParams instantiates a new CommonGraphNodeFilterParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonGraphNodeFilterParamsWithDefaults

`func NewCommonGraphNodeFilterParamsWithDefaults() *CommonGraphNodeFilterParams`

NewCommonGraphNodeFilterParamsWithDefaults instantiates a new CommonGraphNodeFilterParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CommonGraphNodeFilterParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonGraphNodeFilterParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonGraphNodeFilterParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CommonGraphNodeFilterParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CommonGraphNodeFilterParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CommonGraphNodeFilterParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRootOnly

`func (o *CommonGraphNodeFilterParams) GetRootOnly() bool`

GetRootOnly returns the RootOnly field if non-nil, zero value otherwise.

### GetRootOnlyOk

`func (o *CommonGraphNodeFilterParams) GetRootOnlyOk() (*bool, bool)`

GetRootOnlyOk returns a tuple with the RootOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootOnly

`func (o *CommonGraphNodeFilterParams) SetRootOnly(v bool)`

SetRootOnly sets RootOnly field to given value.

### HasRootOnly

`func (o *CommonGraphNodeFilterParams) HasRootOnly() bool`

HasRootOnly returns a boolean if a field has been set.

### SetRootOnlyNil

`func (o *CommonGraphNodeFilterParams) SetRootOnlyNil(b bool)`

 SetRootOnlyNil sets the value for RootOnly to be an explicit nil

### UnsetRootOnly
`func (o *CommonGraphNodeFilterParams) UnsetRootOnly()`

UnsetRootOnly ensures that no value is present for RootOnly, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


