# UnRegisterApplicationServersParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationEnvironments** | Pointer to **[]string** | Specifies the types of applications such as &#39;kSQL&#39;, &#39;kExchange&#39;, &#39;kAD&#39; etc. running on the Protection Source. | [optional] 

## Methods

### NewUnRegisterApplicationServersParams

`func NewUnRegisterApplicationServersParams() *UnRegisterApplicationServersParams`

NewUnRegisterApplicationServersParams instantiates a new UnRegisterApplicationServersParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnRegisterApplicationServersParamsWithDefaults

`func NewUnRegisterApplicationServersParamsWithDefaults() *UnRegisterApplicationServersParams`

NewUnRegisterApplicationServersParamsWithDefaults instantiates a new UnRegisterApplicationServersParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationEnvironments

`func (o *UnRegisterApplicationServersParams) GetApplicationEnvironments() []string`

GetApplicationEnvironments returns the ApplicationEnvironments field if non-nil, zero value otherwise.

### GetApplicationEnvironmentsOk

`func (o *UnRegisterApplicationServersParams) GetApplicationEnvironmentsOk() (*[]string, bool)`

GetApplicationEnvironmentsOk returns a tuple with the ApplicationEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationEnvironments

`func (o *UnRegisterApplicationServersParams) SetApplicationEnvironments(v []string)`

SetApplicationEnvironments sets ApplicationEnvironments field to given value.

### HasApplicationEnvironments

`func (o *UnRegisterApplicationServersParams) HasApplicationEnvironments() bool`

HasApplicationEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


