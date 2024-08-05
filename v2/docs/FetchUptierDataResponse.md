# FetchUptierDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSize** | Pointer to **NullableInt64** | Specifies the amount of data in bytes estimated to be uptiered as part of the current restore job. | [optional] 

## Methods

### NewFetchUptierDataResponse

`func NewFetchUptierDataResponse() *FetchUptierDataResponse`

NewFetchUptierDataResponse instantiates a new FetchUptierDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchUptierDataResponseWithDefaults

`func NewFetchUptierDataResponseWithDefaults() *FetchUptierDataResponse`

NewFetchUptierDataResponseWithDefaults instantiates a new FetchUptierDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSize

`func (o *FetchUptierDataResponse) GetDataSize() int64`

GetDataSize returns the DataSize field if non-nil, zero value otherwise.

### GetDataSizeOk

`func (o *FetchUptierDataResponse) GetDataSizeOk() (*int64, bool)`

GetDataSizeOk returns a tuple with the DataSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSize

`func (o *FetchUptierDataResponse) SetDataSize(v int64)`

SetDataSize sets DataSize field to given value.

### HasDataSize

`func (o *FetchUptierDataResponse) HasDataSize() bool`

HasDataSize returns a boolean if a field has been set.

### SetDataSizeNil

`func (o *FetchUptierDataResponse) SetDataSizeNil(b bool)`

 SetDataSizeNil sets the value for DataSize to be an explicit nil

### UnsetDataSize
`func (o *FetchUptierDataResponse) UnsetDataSize()`

UnsetDataSize ensures that no value is present for DataSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


