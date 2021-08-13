# HiveAdditionalParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetastoreAddress** | Pointer to **string** | The MetastoreAddress for this Hive. | [optional] [readonly] 
**MetastorePort** | Pointer to **int32** | The MetastorePort for this Hive. | [optional] [readonly] 
**AuthType** | Pointer to **NullableString** | Authentication type. | [optional] [readonly] 

## Methods

### NewHiveAdditionalParams

`func NewHiveAdditionalParams() *HiveAdditionalParams`

NewHiveAdditionalParams instantiates a new HiveAdditionalParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHiveAdditionalParamsWithDefaults

`func NewHiveAdditionalParamsWithDefaults() *HiveAdditionalParams`

NewHiveAdditionalParamsWithDefaults instantiates a new HiveAdditionalParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetastoreAddress

`func (o *HiveAdditionalParams) GetMetastoreAddress() string`

GetMetastoreAddress returns the MetastoreAddress field if non-nil, zero value otherwise.

### GetMetastoreAddressOk

`func (o *HiveAdditionalParams) GetMetastoreAddressOk() (*string, bool)`

GetMetastoreAddressOk returns a tuple with the MetastoreAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetastoreAddress

`func (o *HiveAdditionalParams) SetMetastoreAddress(v string)`

SetMetastoreAddress sets MetastoreAddress field to given value.

### HasMetastoreAddress

`func (o *HiveAdditionalParams) HasMetastoreAddress() bool`

HasMetastoreAddress returns a boolean if a field has been set.

### GetMetastorePort

`func (o *HiveAdditionalParams) GetMetastorePort() int32`

GetMetastorePort returns the MetastorePort field if non-nil, zero value otherwise.

### GetMetastorePortOk

`func (o *HiveAdditionalParams) GetMetastorePortOk() (*int32, bool)`

GetMetastorePortOk returns a tuple with the MetastorePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetastorePort

`func (o *HiveAdditionalParams) SetMetastorePort(v int32)`

SetMetastorePort sets MetastorePort field to given value.

### HasMetastorePort

`func (o *HiveAdditionalParams) HasMetastorePort() bool`

HasMetastorePort returns a boolean if a field has been set.

### GetAuthType

`func (o *HiveAdditionalParams) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *HiveAdditionalParams) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *HiveAdditionalParams) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *HiveAdditionalParams) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### SetAuthTypeNil

`func (o *HiveAdditionalParams) SetAuthTypeNil(b bool)`

 SetAuthTypeNil sets the value for AuthType to be an explicit nil

### UnsetAuthType
`func (o *HiveAdditionalParams) UnsetAuthType()`

UnsetAuthType ensures that no value is present for AuthType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


