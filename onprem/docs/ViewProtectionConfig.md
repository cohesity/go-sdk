# ViewProtectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionGroupType** | **NullableString** | Specifies the View protection group type. | 
**ExistingGroupParam** | Pointer to [**ExistingGroupParam**](ExistingGroupParam.md) | Specifies the parameters used for existing protection group. | [optional] 
**NewGroupParam** | Pointer to [**NewGroupParam**](NewGroupParam.md) | Specifies the parameters used for new protection group. | [optional] 

## Methods

### NewViewProtectionConfig

`func NewViewProtectionConfig(protectionGroupType NullableString, ) *ViewProtectionConfig`

NewViewProtectionConfig instantiates a new ViewProtectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewProtectionConfigWithDefaults

`func NewViewProtectionConfigWithDefaults() *ViewProtectionConfig`

NewViewProtectionConfigWithDefaults instantiates a new ViewProtectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectionGroupType

`func (o *ViewProtectionConfig) GetProtectionGroupType() string`

GetProtectionGroupType returns the ProtectionGroupType field if non-nil, zero value otherwise.

### GetProtectionGroupTypeOk

`func (o *ViewProtectionConfig) GetProtectionGroupTypeOk() (*string, bool)`

GetProtectionGroupTypeOk returns a tuple with the ProtectionGroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupType

`func (o *ViewProtectionConfig) SetProtectionGroupType(v string)`

SetProtectionGroupType sets ProtectionGroupType field to given value.


### SetProtectionGroupTypeNil

`func (o *ViewProtectionConfig) SetProtectionGroupTypeNil(b bool)`

 SetProtectionGroupTypeNil sets the value for ProtectionGroupType to be an explicit nil

### UnsetProtectionGroupType
`func (o *ViewProtectionConfig) UnsetProtectionGroupType()`

UnsetProtectionGroupType ensures that no value is present for ProtectionGroupType, not even an explicit nil
### GetExistingGroupParam

`func (o *ViewProtectionConfig) GetExistingGroupParam() ExistingGroupParam`

GetExistingGroupParam returns the ExistingGroupParam field if non-nil, zero value otherwise.

### GetExistingGroupParamOk

`func (o *ViewProtectionConfig) GetExistingGroupParamOk() (*ExistingGroupParam, bool)`

GetExistingGroupParamOk returns a tuple with the ExistingGroupParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingGroupParam

`func (o *ViewProtectionConfig) SetExistingGroupParam(v ExistingGroupParam)`

SetExistingGroupParam sets ExistingGroupParam field to given value.

### HasExistingGroupParam

`func (o *ViewProtectionConfig) HasExistingGroupParam() bool`

HasExistingGroupParam returns a boolean if a field has been set.

### GetNewGroupParam

`func (o *ViewProtectionConfig) GetNewGroupParam() NewGroupParam`

GetNewGroupParam returns the NewGroupParam field if non-nil, zero value otherwise.

### GetNewGroupParamOk

`func (o *ViewProtectionConfig) GetNewGroupParamOk() (*NewGroupParam, bool)`

GetNewGroupParamOk returns a tuple with the NewGroupParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewGroupParam

`func (o *ViewProtectionConfig) SetNewGroupParam(v NewGroupParam)`

SetNewGroupParam sets NewGroupParam field to given value.

### HasNewGroupParam

`func (o *ViewProtectionConfig) HasNewGroupParam() bool`

HasNewGroupParam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


