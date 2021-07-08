# MappersWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mappers** | Pointer to [**[]MapperInfo**](MapperInfo.md) | Mappers specifies the list of available mappers in analytics workbench. | [optional] 

## Methods

### NewMappersWrapper

`func NewMappersWrapper() *MappersWrapper`

NewMappersWrapper instantiates a new MappersWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMappersWrapperWithDefaults

`func NewMappersWrapperWithDefaults() *MappersWrapper`

NewMappersWrapperWithDefaults instantiates a new MappersWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMappers

`func (o *MappersWrapper) GetMappers() []MapperInfo`

GetMappers returns the Mappers field if non-nil, zero value otherwise.

### GetMappersOk

`func (o *MappersWrapper) GetMappersOk() (*[]MapperInfo, bool)`

GetMappersOk returns a tuple with the Mappers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappers

`func (o *MappersWrapper) SetMappers(v []MapperInfo)`

SetMappers sets Mappers field to given value.

### HasMappers

`func (o *MappersWrapper) HasMappers() bool`

HasMappers returns a boolean if a field has been set.

### SetMappersNil

`func (o *MappersWrapper) SetMappersNil(b bool)`

 SetMappersNil sets the value for Mappers to be an explicit nil

### UnsetMappers
`func (o *MappersWrapper) UnsetMappers()`

UnsetMappers ensures that no value is present for Mappers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


