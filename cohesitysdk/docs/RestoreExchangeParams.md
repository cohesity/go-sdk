# RestoreExchangeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseOptions** | Pointer to [**RestoreExchangeParamsDatabaseOptions**](RestoreExchangeParamsDatabaseOptions.md) |  | [optional] 
**Type** | Pointer to **NullableInt32** | Restore type. | [optional] 
**ViewOptions** | Pointer to [**RestoreExchangeParamsViewOptions**](RestoreExchangeParamsViewOptions.md) |  | [optional] 

## Methods

### NewRestoreExchangeParams

`func NewRestoreExchangeParams() *RestoreExchangeParams`

NewRestoreExchangeParams instantiates a new RestoreExchangeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreExchangeParamsWithDefaults

`func NewRestoreExchangeParamsWithDefaults() *RestoreExchangeParams`

NewRestoreExchangeParamsWithDefaults instantiates a new RestoreExchangeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseOptions

`func (o *RestoreExchangeParams) GetDatabaseOptions() RestoreExchangeParamsDatabaseOptions`

GetDatabaseOptions returns the DatabaseOptions field if non-nil, zero value otherwise.

### GetDatabaseOptionsOk

`func (o *RestoreExchangeParams) GetDatabaseOptionsOk() (*RestoreExchangeParamsDatabaseOptions, bool)`

GetDatabaseOptionsOk returns a tuple with the DatabaseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseOptions

`func (o *RestoreExchangeParams) SetDatabaseOptions(v RestoreExchangeParamsDatabaseOptions)`

SetDatabaseOptions sets DatabaseOptions field to given value.

### HasDatabaseOptions

`func (o *RestoreExchangeParams) HasDatabaseOptions() bool`

HasDatabaseOptions returns a boolean if a field has been set.

### GetType

`func (o *RestoreExchangeParams) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestoreExchangeParams) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestoreExchangeParams) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *RestoreExchangeParams) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *RestoreExchangeParams) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *RestoreExchangeParams) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetViewOptions

`func (o *RestoreExchangeParams) GetViewOptions() RestoreExchangeParamsViewOptions`

GetViewOptions returns the ViewOptions field if non-nil, zero value otherwise.

### GetViewOptionsOk

`func (o *RestoreExchangeParams) GetViewOptionsOk() (*RestoreExchangeParamsViewOptions, bool)`

GetViewOptionsOk returns a tuple with the ViewOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewOptions

`func (o *RestoreExchangeParams) SetViewOptions(v RestoreExchangeParamsViewOptions)`

SetViewOptions sets ViewOptions field to given value.

### HasViewOptions

`func (o *RestoreExchangeParams) HasViewOptions() bool`

HasViewOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


