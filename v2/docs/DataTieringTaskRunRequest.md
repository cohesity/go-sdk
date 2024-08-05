# DataTieringTaskRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shares** | Pointer to [**[]DataTieringShareInfo**](DataTieringShareInfo.md) | Specifies the list of shares to tier. | [optional] 
**UptierPath** | Pointer to **NullableString** | Only applicable for uptiering tasks. Ignore the uptiering policy and uptier the directory pointed by the &#39;uptierPath&#39;. If path is &#39;/&#39;, then uptier everything.  This is a global property which will be applied to all shares by default. This can be overriden by specifying uptierPath for each share. | [optional] 

## Methods

### NewDataTieringTaskRunRequest

`func NewDataTieringTaskRunRequest() *DataTieringTaskRunRequest`

NewDataTieringTaskRunRequest instantiates a new DataTieringTaskRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTieringTaskRunRequestWithDefaults

`func NewDataTieringTaskRunRequestWithDefaults() *DataTieringTaskRunRequest`

NewDataTieringTaskRunRequestWithDefaults instantiates a new DataTieringTaskRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShares

`func (o *DataTieringTaskRunRequest) GetShares() []DataTieringShareInfo`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *DataTieringTaskRunRequest) GetSharesOk() (*[]DataTieringShareInfo, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *DataTieringTaskRunRequest) SetShares(v []DataTieringShareInfo)`

SetShares sets Shares field to given value.

### HasShares

`func (o *DataTieringTaskRunRequest) HasShares() bool`

HasShares returns a boolean if a field has been set.

### SetSharesNil

`func (o *DataTieringTaskRunRequest) SetSharesNil(b bool)`

 SetSharesNil sets the value for Shares to be an explicit nil

### UnsetShares
`func (o *DataTieringTaskRunRequest) UnsetShares()`

UnsetShares ensures that no value is present for Shares, not even an explicit nil
### GetUptierPath

`func (o *DataTieringTaskRunRequest) GetUptierPath() string`

GetUptierPath returns the UptierPath field if non-nil, zero value otherwise.

### GetUptierPathOk

`func (o *DataTieringTaskRunRequest) GetUptierPathOk() (*string, bool)`

GetUptierPathOk returns a tuple with the UptierPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptierPath

`func (o *DataTieringTaskRunRequest) SetUptierPath(v string)`

SetUptierPath sets UptierPath field to given value.

### HasUptierPath

`func (o *DataTieringTaskRunRequest) HasUptierPath() bool`

HasUptierPath returns a boolean if a field has been set.

### SetUptierPathNil

`func (o *DataTieringTaskRunRequest) SetUptierPathNil(b bool)`

 SetUptierPathNil sets the value for UptierPath to be an explicit nil

### UnsetUptierPath
`func (o *DataTieringTaskRunRequest) UnsetUptierPath()`

UnsetUptierPath ensures that no value is present for UptierPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


