# OnpremDeployParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies the unique id of the onprem entity. | [optional] 
**RestoreVMwareParams** | Pointer to [**RestoreVMwareVMParams**](RestoreVMwareVMParams.md) |  | [optional] 

## Methods

### NewOnpremDeployParams

`func NewOnpremDeployParams() *OnpremDeployParams`

NewOnpremDeployParams instantiates a new OnpremDeployParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnpremDeployParamsWithDefaults

`func NewOnpremDeployParamsWithDefaults() *OnpremDeployParams`

NewOnpremDeployParamsWithDefaults instantiates a new OnpremDeployParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OnpremDeployParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OnpremDeployParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OnpremDeployParams) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OnpremDeployParams) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *OnpremDeployParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *OnpremDeployParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetRestoreVMwareParams

`func (o *OnpremDeployParams) GetRestoreVMwareParams() RestoreVMwareVMParams`

GetRestoreVMwareParams returns the RestoreVMwareParams field if non-nil, zero value otherwise.

### GetRestoreVMwareParamsOk

`func (o *OnpremDeployParams) GetRestoreVMwareParamsOk() (*RestoreVMwareVMParams, bool)`

GetRestoreVMwareParamsOk returns a tuple with the RestoreVMwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreVMwareParams

`func (o *OnpremDeployParams) SetRestoreVMwareParams(v RestoreVMwareVMParams)`

SetRestoreVMwareParams sets RestoreVMwareParams field to given value.

### HasRestoreVMwareParams

`func (o *OnpremDeployParams) HasRestoreVMwareParams() bool`

HasRestoreVMwareParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


