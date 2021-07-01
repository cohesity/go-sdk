# MapReduceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppProperty** | Pointer to [**MapReduceInfoAppProperty**](MapReduceInfoAppProperty.md) |  | [optional] 
**AuxData** | Pointer to [**MapReduceAuxData**](MapReduceAuxData.md) |  | [optional] 
**Description** | Pointer to **NullableString** | Map reduce job description. | [optional] 
**ExcludedDataSourceVec** | Pointer to **[]int32** | List of all excluded data sources for this app. | [optional] 
**Id** | Pointer to **NullableInt64** | ID of map reduce job. | [optional] 
**IsSystemDefined** | Pointer to **NullableBool** | Flag to denote if this is system pre-defined app or user has written this app. | [optional] 
**MapperId** | Pointer to **NullableInt64** | ID of the mapper to process the input. | [optional] 
**Name** | Pointer to **NullableString** | Map reduce job name. | [optional] 
**ReducerId** | Pointer to **NullableInt64** | ID of the reducer. | [optional] 
**RequiredPropertyVec** | Pointer to [**[]MapReduceInfoRequiredProperty**](MapReduceInfoRequiredProperty.md) |  | [optional] 

## Methods

### NewMapReduceInfo

`func NewMapReduceInfo() *MapReduceInfo`

NewMapReduceInfo instantiates a new MapReduceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapReduceInfoWithDefaults

`func NewMapReduceInfoWithDefaults() *MapReduceInfo`

NewMapReduceInfoWithDefaults instantiates a new MapReduceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppProperty

`func (o *MapReduceInfo) GetAppProperty() MapReduceInfoAppProperty`

GetAppProperty returns the AppProperty field if non-nil, zero value otherwise.

### GetAppPropertyOk

`func (o *MapReduceInfo) GetAppPropertyOk() (*MapReduceInfoAppProperty, bool)`

GetAppPropertyOk returns a tuple with the AppProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProperty

`func (o *MapReduceInfo) SetAppProperty(v MapReduceInfoAppProperty)`

SetAppProperty sets AppProperty field to given value.

### HasAppProperty

`func (o *MapReduceInfo) HasAppProperty() bool`

HasAppProperty returns a boolean if a field has been set.

### GetAuxData

`func (o *MapReduceInfo) GetAuxData() MapReduceAuxData`

GetAuxData returns the AuxData field if non-nil, zero value otherwise.

### GetAuxDataOk

`func (o *MapReduceInfo) GetAuxDataOk() (*MapReduceAuxData, bool)`

GetAuxDataOk returns a tuple with the AuxData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxData

`func (o *MapReduceInfo) SetAuxData(v MapReduceAuxData)`

SetAuxData sets AuxData field to given value.

### HasAuxData

`func (o *MapReduceInfo) HasAuxData() bool`

HasAuxData returns a boolean if a field has been set.

### GetDescription

`func (o *MapReduceInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MapReduceInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MapReduceInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MapReduceInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MapReduceInfo) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MapReduceInfo) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExcludedDataSourceVec

`func (o *MapReduceInfo) GetExcludedDataSourceVec() []int32`

GetExcludedDataSourceVec returns the ExcludedDataSourceVec field if non-nil, zero value otherwise.

### GetExcludedDataSourceVecOk

`func (o *MapReduceInfo) GetExcludedDataSourceVecOk() (*[]int32, bool)`

GetExcludedDataSourceVecOk returns a tuple with the ExcludedDataSourceVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedDataSourceVec

`func (o *MapReduceInfo) SetExcludedDataSourceVec(v []int32)`

SetExcludedDataSourceVec sets ExcludedDataSourceVec field to given value.

### HasExcludedDataSourceVec

`func (o *MapReduceInfo) HasExcludedDataSourceVec() bool`

HasExcludedDataSourceVec returns a boolean if a field has been set.

### SetExcludedDataSourceVecNil

`func (o *MapReduceInfo) SetExcludedDataSourceVecNil(b bool)`

 SetExcludedDataSourceVecNil sets the value for ExcludedDataSourceVec to be an explicit nil

### UnsetExcludedDataSourceVec
`func (o *MapReduceInfo) UnsetExcludedDataSourceVec()`

UnsetExcludedDataSourceVec ensures that no value is present for ExcludedDataSourceVec, not even an explicit nil
### GetId

`func (o *MapReduceInfo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MapReduceInfo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MapReduceInfo) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MapReduceInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MapReduceInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MapReduceInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsSystemDefined

`func (o *MapReduceInfo) GetIsSystemDefined() bool`

GetIsSystemDefined returns the IsSystemDefined field if non-nil, zero value otherwise.

### GetIsSystemDefinedOk

`func (o *MapReduceInfo) GetIsSystemDefinedOk() (*bool, bool)`

GetIsSystemDefinedOk returns a tuple with the IsSystemDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemDefined

`func (o *MapReduceInfo) SetIsSystemDefined(v bool)`

SetIsSystemDefined sets IsSystemDefined field to given value.

### HasIsSystemDefined

`func (o *MapReduceInfo) HasIsSystemDefined() bool`

HasIsSystemDefined returns a boolean if a field has been set.

### SetIsSystemDefinedNil

`func (o *MapReduceInfo) SetIsSystemDefinedNil(b bool)`

 SetIsSystemDefinedNil sets the value for IsSystemDefined to be an explicit nil

### UnsetIsSystemDefined
`func (o *MapReduceInfo) UnsetIsSystemDefined()`

UnsetIsSystemDefined ensures that no value is present for IsSystemDefined, not even an explicit nil
### GetMapperId

`func (o *MapReduceInfo) GetMapperId() int64`

GetMapperId returns the MapperId field if non-nil, zero value otherwise.

### GetMapperIdOk

`func (o *MapReduceInfo) GetMapperIdOk() (*int64, bool)`

GetMapperIdOk returns a tuple with the MapperId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapperId

`func (o *MapReduceInfo) SetMapperId(v int64)`

SetMapperId sets MapperId field to given value.

### HasMapperId

`func (o *MapReduceInfo) HasMapperId() bool`

HasMapperId returns a boolean if a field has been set.

### SetMapperIdNil

`func (o *MapReduceInfo) SetMapperIdNil(b bool)`

 SetMapperIdNil sets the value for MapperId to be an explicit nil

### UnsetMapperId
`func (o *MapReduceInfo) UnsetMapperId()`

UnsetMapperId ensures that no value is present for MapperId, not even an explicit nil
### GetName

`func (o *MapReduceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MapReduceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MapReduceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MapReduceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MapReduceInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MapReduceInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReducerId

`func (o *MapReduceInfo) GetReducerId() int64`

GetReducerId returns the ReducerId field if non-nil, zero value otherwise.

### GetReducerIdOk

`func (o *MapReduceInfo) GetReducerIdOk() (*int64, bool)`

GetReducerIdOk returns a tuple with the ReducerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReducerId

`func (o *MapReduceInfo) SetReducerId(v int64)`

SetReducerId sets ReducerId field to given value.

### HasReducerId

`func (o *MapReduceInfo) HasReducerId() bool`

HasReducerId returns a boolean if a field has been set.

### SetReducerIdNil

`func (o *MapReduceInfo) SetReducerIdNil(b bool)`

 SetReducerIdNil sets the value for ReducerId to be an explicit nil

### UnsetReducerId
`func (o *MapReduceInfo) UnsetReducerId()`

UnsetReducerId ensures that no value is present for ReducerId, not even an explicit nil
### GetRequiredPropertyVec

`func (o *MapReduceInfo) GetRequiredPropertyVec() []MapReduceInfoRequiredProperty`

GetRequiredPropertyVec returns the RequiredPropertyVec field if non-nil, zero value otherwise.

### GetRequiredPropertyVecOk

`func (o *MapReduceInfo) GetRequiredPropertyVecOk() (*[]MapReduceInfoRequiredProperty, bool)`

GetRequiredPropertyVecOk returns a tuple with the RequiredPropertyVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredPropertyVec

`func (o *MapReduceInfo) SetRequiredPropertyVec(v []MapReduceInfoRequiredProperty)`

SetRequiredPropertyVec sets RequiredPropertyVec field to given value.

### HasRequiredPropertyVec

`func (o *MapReduceInfo) HasRequiredPropertyVec() bool`

HasRequiredPropertyVec returns a boolean if a field has been set.

### SetRequiredPropertyVecNil

`func (o *MapReduceInfo) SetRequiredPropertyVecNil(b bool)`

 SetRequiredPropertyVecNil sets the value for RequiredPropertyVec to be an explicit nil

### UnsetRequiredPropertyVec
`func (o *MapReduceInfo) UnsetRequiredPropertyVec()`

UnsetRequiredPropertyVec ensures that no value is present for RequiredPropertyVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


