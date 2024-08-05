# CapacityTrendDatapointStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumBytes** | Pointer to **NullableInt64** | Specifies the size of data corresponding to above tag. | [optional] 
**TagInfo** | Pointer to [**DataTieringTag**](DataTieringTag.md) |  | [optional] 

## Methods

### NewCapacityTrendDatapointStats

`func NewCapacityTrendDatapointStats() *CapacityTrendDatapointStats`

NewCapacityTrendDatapointStats instantiates a new CapacityTrendDatapointStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapacityTrendDatapointStatsWithDefaults

`func NewCapacityTrendDatapointStatsWithDefaults() *CapacityTrendDatapointStats`

NewCapacityTrendDatapointStatsWithDefaults instantiates a new CapacityTrendDatapointStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumBytes

`func (o *CapacityTrendDatapointStats) GetNumBytes() int64`

GetNumBytes returns the NumBytes field if non-nil, zero value otherwise.

### GetNumBytesOk

`func (o *CapacityTrendDatapointStats) GetNumBytesOk() (*int64, bool)`

GetNumBytesOk returns a tuple with the NumBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumBytes

`func (o *CapacityTrendDatapointStats) SetNumBytes(v int64)`

SetNumBytes sets NumBytes field to given value.

### HasNumBytes

`func (o *CapacityTrendDatapointStats) HasNumBytes() bool`

HasNumBytes returns a boolean if a field has been set.

### SetNumBytesNil

`func (o *CapacityTrendDatapointStats) SetNumBytesNil(b bool)`

 SetNumBytesNil sets the value for NumBytes to be an explicit nil

### UnsetNumBytes
`func (o *CapacityTrendDatapointStats) UnsetNumBytes()`

UnsetNumBytes ensures that no value is present for NumBytes, not even an explicit nil
### GetTagInfo

`func (o *CapacityTrendDatapointStats) GetTagInfo() DataTieringTag`

GetTagInfo returns the TagInfo field if non-nil, zero value otherwise.

### GetTagInfoOk

`func (o *CapacityTrendDatapointStats) GetTagInfoOk() (*DataTieringTag, bool)`

GetTagInfoOk returns a tuple with the TagInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagInfo

`func (o *CapacityTrendDatapointStats) SetTagInfo(v DataTieringTag)`

SetTagInfo sets TagInfo field to given value.

### HasTagInfo

`func (o *CapacityTrendDatapointStats) HasTagInfo() bool`

HasTagInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


