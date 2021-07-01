# AnalyseJarResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**Mappers** | Pointer to **[]string** | Name of all mapper classes found in jar file. | [optional] 
**Reducers** | Pointer to **[]string** | Name of all reducers classes found in jar file. | [optional] 

## Methods

### NewAnalyseJarResult

`func NewAnalyseJarResult() *AnalyseJarResult`

NewAnalyseJarResult instantiates a new AnalyseJarResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyseJarResultWithDefaults

`func NewAnalyseJarResultWithDefaults() *AnalyseJarResult`

NewAnalyseJarResultWithDefaults instantiates a new AnalyseJarResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *AnalyseJarResult) GetError() ErrorProto`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AnalyseJarResult) GetErrorOk() (*ErrorProto, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AnalyseJarResult) SetError(v ErrorProto)`

SetError sets Error field to given value.

### HasError

`func (o *AnalyseJarResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMappers

`func (o *AnalyseJarResult) GetMappers() []string`

GetMappers returns the Mappers field if non-nil, zero value otherwise.

### GetMappersOk

`func (o *AnalyseJarResult) GetMappersOk() (*[]string, bool)`

GetMappersOk returns a tuple with the Mappers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappers

`func (o *AnalyseJarResult) SetMappers(v []string)`

SetMappers sets Mappers field to given value.

### HasMappers

`func (o *AnalyseJarResult) HasMappers() bool`

HasMappers returns a boolean if a field has been set.

### SetMappersNil

`func (o *AnalyseJarResult) SetMappersNil(b bool)`

 SetMappersNil sets the value for Mappers to be an explicit nil

### UnsetMappers
`func (o *AnalyseJarResult) UnsetMappers()`

UnsetMappers ensures that no value is present for Mappers, not even an explicit nil
### GetReducers

`func (o *AnalyseJarResult) GetReducers() []string`

GetReducers returns the Reducers field if non-nil, zero value otherwise.

### GetReducersOk

`func (o *AnalyseJarResult) GetReducersOk() (*[]string, bool)`

GetReducersOk returns a tuple with the Reducers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReducers

`func (o *AnalyseJarResult) SetReducers(v []string)`

SetReducers sets Reducers field to given value.

### HasReducers

`func (o *AnalyseJarResult) HasReducers() bool`

HasReducers returns a boolean if a field has been set.

### SetReducersNil

`func (o *AnalyseJarResult) SetReducersNil(b bool)`

 SetReducersNil sets the value for Reducers to be an explicit nil

### UnsetReducers
`func (o *AnalyseJarResult) UnsetReducers()`

UnsetReducers ensures that no value is present for Reducers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


