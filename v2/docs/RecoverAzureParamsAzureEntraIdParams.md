# RecoverAzureParamsAzureEntraIdParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureTargetParams** | Pointer to [**NullableRecoverAzureEntraIdParamsAzureTargetParams**](RecoverAzureEntraIdParamsAzureTargetParams.md) |  | [optional] 
**ObjectAttributes** | [**[]CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) | Specifies the details of the azure entra id objects attributes to be recovered. | 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverAzureParamsAzureEntraIdParams

`func NewRecoverAzureParamsAzureEntraIdParams(objectAttributes []CommonRecoverObjectSnapshotParams, targetEnvironment string, ) *RecoverAzureParamsAzureEntraIdParams`

NewRecoverAzureParamsAzureEntraIdParams instantiates a new RecoverAzureParamsAzureEntraIdParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAzureParamsAzureEntraIdParamsWithDefaults

`func NewRecoverAzureParamsAzureEntraIdParamsWithDefaults() *RecoverAzureParamsAzureEntraIdParams`

NewRecoverAzureParamsAzureEntraIdParamsWithDefaults instantiates a new RecoverAzureParamsAzureEntraIdParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureTargetParams

`func (o *RecoverAzureParamsAzureEntraIdParams) GetAzureTargetParams() RecoverAzureEntraIdParamsAzureTargetParams`

GetAzureTargetParams returns the AzureTargetParams field if non-nil, zero value otherwise.

### GetAzureTargetParamsOk

`func (o *RecoverAzureParamsAzureEntraIdParams) GetAzureTargetParamsOk() (*RecoverAzureEntraIdParamsAzureTargetParams, bool)`

GetAzureTargetParamsOk returns a tuple with the AzureTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTargetParams

`func (o *RecoverAzureParamsAzureEntraIdParams) SetAzureTargetParams(v RecoverAzureEntraIdParamsAzureTargetParams)`

SetAzureTargetParams sets AzureTargetParams field to given value.

### HasAzureTargetParams

`func (o *RecoverAzureParamsAzureEntraIdParams) HasAzureTargetParams() bool`

HasAzureTargetParams returns a boolean if a field has been set.

### SetAzureTargetParamsNil

`func (o *RecoverAzureParamsAzureEntraIdParams) SetAzureTargetParamsNil(b bool)`

 SetAzureTargetParamsNil sets the value for AzureTargetParams to be an explicit nil

### UnsetAzureTargetParams
`func (o *RecoverAzureParamsAzureEntraIdParams) UnsetAzureTargetParams()`

UnsetAzureTargetParams ensures that no value is present for AzureTargetParams, not even an explicit nil
### GetObjectAttributes

`func (o *RecoverAzureParamsAzureEntraIdParams) GetObjectAttributes() []CommonRecoverObjectSnapshotParams`

GetObjectAttributes returns the ObjectAttributes field if non-nil, zero value otherwise.

### GetObjectAttributesOk

`func (o *RecoverAzureParamsAzureEntraIdParams) GetObjectAttributesOk() (*[]CommonRecoverObjectSnapshotParams, bool)`

GetObjectAttributesOk returns a tuple with the ObjectAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectAttributes

`func (o *RecoverAzureParamsAzureEntraIdParams) SetObjectAttributes(v []CommonRecoverObjectSnapshotParams)`

SetObjectAttributes sets ObjectAttributes field to given value.


### SetObjectAttributesNil

`func (o *RecoverAzureParamsAzureEntraIdParams) SetObjectAttributesNil(b bool)`

 SetObjectAttributesNil sets the value for ObjectAttributes to be an explicit nil

### UnsetObjectAttributes
`func (o *RecoverAzureParamsAzureEntraIdParams) UnsetObjectAttributes()`

UnsetObjectAttributes ensures that no value is present for ObjectAttributes, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverAzureParamsAzureEntraIdParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverAzureParamsAzureEntraIdParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverAzureParamsAzureEntraIdParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


