# DataTieringShareInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShareId** | **NullableInt64** | Specifies the id of the share. | 
**UptierPath** | Pointer to **NullableString** | Only applicable for uptiering tasks. Ignore the uptiering policy and uptier the directory pointed by the &#39;uptierPath&#39;. If path is &#39;/&#39;, then uptier everything  This will override the global uptier path. | [optional] 

## Methods

### NewDataTieringShareInfo

`func NewDataTieringShareInfo(shareId NullableInt64, ) *DataTieringShareInfo`

NewDataTieringShareInfo instantiates a new DataTieringShareInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTieringShareInfoWithDefaults

`func NewDataTieringShareInfoWithDefaults() *DataTieringShareInfo`

NewDataTieringShareInfoWithDefaults instantiates a new DataTieringShareInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShareId

`func (o *DataTieringShareInfo) GetShareId() int64`

GetShareId returns the ShareId field if non-nil, zero value otherwise.

### GetShareIdOk

`func (o *DataTieringShareInfo) GetShareIdOk() (*int64, bool)`

GetShareIdOk returns a tuple with the ShareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareId

`func (o *DataTieringShareInfo) SetShareId(v int64)`

SetShareId sets ShareId field to given value.


### SetShareIdNil

`func (o *DataTieringShareInfo) SetShareIdNil(b bool)`

 SetShareIdNil sets the value for ShareId to be an explicit nil

### UnsetShareId
`func (o *DataTieringShareInfo) UnsetShareId()`

UnsetShareId ensures that no value is present for ShareId, not even an explicit nil
### GetUptierPath

`func (o *DataTieringShareInfo) GetUptierPath() string`

GetUptierPath returns the UptierPath field if non-nil, zero value otherwise.

### GetUptierPathOk

`func (o *DataTieringShareInfo) GetUptierPathOk() (*string, bool)`

GetUptierPathOk returns a tuple with the UptierPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptierPath

`func (o *DataTieringShareInfo) SetUptierPath(v string)`

SetUptierPath sets UptierPath field to given value.

### HasUptierPath

`func (o *DataTieringShareInfo) HasUptierPath() bool`

HasUptierPath returns a boolean if a field has been set.

### SetUptierPathNil

`func (o *DataTieringShareInfo) SetUptierPathNil(b bool)`

 SetUptierPathNil sets the value for UptierPath to be an explicit nil

### UnsetUptierPath
`func (o *DataTieringShareInfo) UnsetUptierPath()`

UnsetUptierPath ensures that no value is present for UptierPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


