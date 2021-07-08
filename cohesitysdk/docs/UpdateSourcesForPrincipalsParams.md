# UpdateSourcesForPrincipalsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourcesForPrincipals** | Pointer to [**[]SourceForPrincipalParam**](SourceForPrincipalParam.md) | Array of Principals, Sources and Views.  Specifies a list of principals. For each principal, specify the Protection Sources and Views that the principal has permissions to access. | [optional] 

## Methods

### NewUpdateSourcesForPrincipalsParams

`func NewUpdateSourcesForPrincipalsParams() *UpdateSourcesForPrincipalsParams`

NewUpdateSourcesForPrincipalsParams instantiates a new UpdateSourcesForPrincipalsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSourcesForPrincipalsParamsWithDefaults

`func NewUpdateSourcesForPrincipalsParamsWithDefaults() *UpdateSourcesForPrincipalsParams`

NewUpdateSourcesForPrincipalsParamsWithDefaults instantiates a new UpdateSourcesForPrincipalsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourcesForPrincipals

`func (o *UpdateSourcesForPrincipalsParams) GetSourcesForPrincipals() []SourceForPrincipalParam`

GetSourcesForPrincipals returns the SourcesForPrincipals field if non-nil, zero value otherwise.

### GetSourcesForPrincipalsOk

`func (o *UpdateSourcesForPrincipalsParams) GetSourcesForPrincipalsOk() (*[]SourceForPrincipalParam, bool)`

GetSourcesForPrincipalsOk returns a tuple with the SourcesForPrincipals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcesForPrincipals

`func (o *UpdateSourcesForPrincipalsParams) SetSourcesForPrincipals(v []SourceForPrincipalParam)`

SetSourcesForPrincipals sets SourcesForPrincipals field to given value.

### HasSourcesForPrincipals

`func (o *UpdateSourcesForPrincipalsParams) HasSourcesForPrincipals() bool`

HasSourcesForPrincipals returns a boolean if a field has been set.

### SetSourcesForPrincipalsNil

`func (o *UpdateSourcesForPrincipalsParams) SetSourcesForPrincipalsNil(b bool)`

 SetSourcesForPrincipalsNil sets the value for SourcesForPrincipals to be an explicit nil

### UnsetSourcesForPrincipals
`func (o *UpdateSourcesForPrincipalsParams) UnsetSourcesForPrincipals()`

UnsetSourcesForPrincipals ensures that no value is present for SourcesForPrincipals, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


