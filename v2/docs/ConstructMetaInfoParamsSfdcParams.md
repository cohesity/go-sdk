# ConstructMetaInfoParamsSfdcParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetaInfoType** | **NullableString** | Specifies the type of meta info to fetch for salesforce object. | 
**ObjectName** | Pointer to **NullableString** | Specifies the name of the object. | [optional] 

## Methods

### NewConstructMetaInfoParamsSfdcParams

`func NewConstructMetaInfoParamsSfdcParams(metaInfoType NullableString, ) *ConstructMetaInfoParamsSfdcParams`

NewConstructMetaInfoParamsSfdcParams instantiates a new ConstructMetaInfoParamsSfdcParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstructMetaInfoParamsSfdcParamsWithDefaults

`func NewConstructMetaInfoParamsSfdcParamsWithDefaults() *ConstructMetaInfoParamsSfdcParams`

NewConstructMetaInfoParamsSfdcParamsWithDefaults instantiates a new ConstructMetaInfoParamsSfdcParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetaInfoType

`func (o *ConstructMetaInfoParamsSfdcParams) GetMetaInfoType() string`

GetMetaInfoType returns the MetaInfoType field if non-nil, zero value otherwise.

### GetMetaInfoTypeOk

`func (o *ConstructMetaInfoParamsSfdcParams) GetMetaInfoTypeOk() (*string, bool)`

GetMetaInfoTypeOk returns a tuple with the MetaInfoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaInfoType

`func (o *ConstructMetaInfoParamsSfdcParams) SetMetaInfoType(v string)`

SetMetaInfoType sets MetaInfoType field to given value.


### SetMetaInfoTypeNil

`func (o *ConstructMetaInfoParamsSfdcParams) SetMetaInfoTypeNil(b bool)`

 SetMetaInfoTypeNil sets the value for MetaInfoType to be an explicit nil

### UnsetMetaInfoType
`func (o *ConstructMetaInfoParamsSfdcParams) UnsetMetaInfoType()`

UnsetMetaInfoType ensures that no value is present for MetaInfoType, not even an explicit nil
### GetObjectName

`func (o *ConstructMetaInfoParamsSfdcParams) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ConstructMetaInfoParamsSfdcParams) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ConstructMetaInfoParamsSfdcParams) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ConstructMetaInfoParamsSfdcParams) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### SetObjectNameNil

`func (o *ConstructMetaInfoParamsSfdcParams) SetObjectNameNil(b bool)`

 SetObjectNameNil sets the value for ObjectName to be an explicit nil

### UnsetObjectName
`func (o *ConstructMetaInfoParamsSfdcParams) UnsetObjectName()`

UnsetObjectName ensures that no value is present for ObjectName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


