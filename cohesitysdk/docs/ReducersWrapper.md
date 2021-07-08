# ReducersWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reducers** | Pointer to [**[]ReducerInfo**](ReducerInfo.md) | Reducers specifies the list of available reducers in analytics workbench. | [optional] 

## Methods

### NewReducersWrapper

`func NewReducersWrapper() *ReducersWrapper`

NewReducersWrapper instantiates a new ReducersWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReducersWrapperWithDefaults

`func NewReducersWrapperWithDefaults() *ReducersWrapper`

NewReducersWrapperWithDefaults instantiates a new ReducersWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReducers

`func (o *ReducersWrapper) GetReducers() []ReducerInfo`

GetReducers returns the Reducers field if non-nil, zero value otherwise.

### GetReducersOk

`func (o *ReducersWrapper) GetReducersOk() (*[]ReducerInfo, bool)`

GetReducersOk returns a tuple with the Reducers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReducers

`func (o *ReducersWrapper) SetReducers(v []ReducerInfo)`

SetReducers sets Reducers field to given value.

### HasReducers

`func (o *ReducersWrapper) HasReducers() bool`

HasReducers returns a boolean if a field has been set.

### SetReducersNil

`func (o *ReducersWrapper) SetReducersNil(b bool)`

 SetReducersNil sets the value for Reducers to be an explicit nil

### UnsetReducers
`func (o *ReducersWrapper) UnsetReducers()`

UnsetReducers ensures that no value is present for Reducers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


