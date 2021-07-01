# RestoreAppObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalParams** | Pointer to [**RestoreTaskAdditionalParams**](RestoreTaskAdditionalParams.md) |  | [optional] 
**AppEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**DisplayName** | Pointer to **NullableString** | The proper display name of this object in the UI, if app_entity is not empty. For example, for SQL databases the name should also include the instance name. | [optional] 
**RestoreParams** | Pointer to [**RestoreAppObjectParams**](RestoreAppObjectParams.md) |  | [optional] 

## Methods

### NewRestoreAppObject

`func NewRestoreAppObject() *RestoreAppObject`

NewRestoreAppObject instantiates a new RestoreAppObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreAppObjectWithDefaults

`func NewRestoreAppObjectWithDefaults() *RestoreAppObject`

NewRestoreAppObjectWithDefaults instantiates a new RestoreAppObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalParams

`func (o *RestoreAppObject) GetAdditionalParams() RestoreTaskAdditionalParams`

GetAdditionalParams returns the AdditionalParams field if non-nil, zero value otherwise.

### GetAdditionalParamsOk

`func (o *RestoreAppObject) GetAdditionalParamsOk() (*RestoreTaskAdditionalParams, bool)`

GetAdditionalParamsOk returns a tuple with the AdditionalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalParams

`func (o *RestoreAppObject) SetAdditionalParams(v RestoreTaskAdditionalParams)`

SetAdditionalParams sets AdditionalParams field to given value.

### HasAdditionalParams

`func (o *RestoreAppObject) HasAdditionalParams() bool`

HasAdditionalParams returns a boolean if a field has been set.

### GetAppEntity

`func (o *RestoreAppObject) GetAppEntity() EntityProto`

GetAppEntity returns the AppEntity field if non-nil, zero value otherwise.

### GetAppEntityOk

`func (o *RestoreAppObject) GetAppEntityOk() (*EntityProto, bool)`

GetAppEntityOk returns a tuple with the AppEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEntity

`func (o *RestoreAppObject) SetAppEntity(v EntityProto)`

SetAppEntity sets AppEntity field to given value.

### HasAppEntity

`func (o *RestoreAppObject) HasAppEntity() bool`

HasAppEntity returns a boolean if a field has been set.

### GetDisplayName

`func (o *RestoreAppObject) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RestoreAppObject) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RestoreAppObject) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RestoreAppObject) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *RestoreAppObject) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *RestoreAppObject) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetRestoreParams

`func (o *RestoreAppObject) GetRestoreParams() RestoreAppObjectParams`

GetRestoreParams returns the RestoreParams field if non-nil, zero value otherwise.

### GetRestoreParamsOk

`func (o *RestoreAppObject) GetRestoreParamsOk() (*RestoreAppObjectParams, bool)`

GetRestoreParamsOk returns a tuple with the RestoreParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreParams

`func (o *RestoreAppObject) SetRestoreParams(v RestoreAppObjectParams)`

SetRestoreParams sets RestoreParams field to given value.

### HasRestoreParams

`func (o *RestoreAppObject) HasRestoreParams() bool`

HasRestoreParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


