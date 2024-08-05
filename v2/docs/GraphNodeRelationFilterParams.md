# GraphNodeRelationFilterParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AadParams** | Pointer to [**NullableAadRelationFilterParams**](AadRelationFilterParams.md) | Specifies the aad adapter specific relation filter params. | [optional] 

## Methods

### NewGraphNodeRelationFilterParams

`func NewGraphNodeRelationFilterParams() *GraphNodeRelationFilterParams`

NewGraphNodeRelationFilterParams instantiates a new GraphNodeRelationFilterParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphNodeRelationFilterParamsWithDefaults

`func NewGraphNodeRelationFilterParamsWithDefaults() *GraphNodeRelationFilterParams`

NewGraphNodeRelationFilterParamsWithDefaults instantiates a new GraphNodeRelationFilterParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAadParams

`func (o *GraphNodeRelationFilterParams) GetAadParams() AadRelationFilterParams`

GetAadParams returns the AadParams field if non-nil, zero value otherwise.

### GetAadParamsOk

`func (o *GraphNodeRelationFilterParams) GetAadParamsOk() (*AadRelationFilterParams, bool)`

GetAadParamsOk returns a tuple with the AadParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAadParams

`func (o *GraphNodeRelationFilterParams) SetAadParams(v AadRelationFilterParams)`

SetAadParams sets AadParams field to given value.

### HasAadParams

`func (o *GraphNodeRelationFilterParams) HasAadParams() bool`

HasAadParams returns a boolean if a field has been set.

### SetAadParamsNil

`func (o *GraphNodeRelationFilterParams) SetAadParamsNil(b bool)`

 SetAadParamsNil sets the value for AadParams to be an explicit nil

### UnsetAadParams
`func (o *GraphNodeRelationFilterParams) UnsetAadParams()`

UnsetAadParams ensures that no value is present for AadParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


