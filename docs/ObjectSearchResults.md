# ObjectSearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectSnapshotInfo** | Pointer to [**[]ObjectSnapshotInfo**](ObjectSnapshotInfo.md) | Array of Snapshot Objects.  Specifies the list of backup objects returned by this request that match the specified search and filter criteria. The number of objects returned is limited by the pageCount field. | [optional] 
**TotalCount** | Pointer to **NullableInt64** | Specifies the total number of backup objects that match the filter and search criteria. Use this value to determine how many additional requests are required to get the full result. | [optional] 

## Methods

### NewObjectSearchResults

`func NewObjectSearchResults() *ObjectSearchResults`

NewObjectSearchResults instantiates a new ObjectSearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectSearchResultsWithDefaults

`func NewObjectSearchResultsWithDefaults() *ObjectSearchResults`

NewObjectSearchResultsWithDefaults instantiates a new ObjectSearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectSnapshotInfo

`func (o *ObjectSearchResults) GetObjectSnapshotInfo() []ObjectSnapshotInfo`

GetObjectSnapshotInfo returns the ObjectSnapshotInfo field if non-nil, zero value otherwise.

### GetObjectSnapshotInfoOk

`func (o *ObjectSearchResults) GetObjectSnapshotInfoOk() (*[]ObjectSnapshotInfo, bool)`

GetObjectSnapshotInfoOk returns a tuple with the ObjectSnapshotInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSnapshotInfo

`func (o *ObjectSearchResults) SetObjectSnapshotInfo(v []ObjectSnapshotInfo)`

SetObjectSnapshotInfo sets ObjectSnapshotInfo field to given value.

### HasObjectSnapshotInfo

`func (o *ObjectSearchResults) HasObjectSnapshotInfo() bool`

HasObjectSnapshotInfo returns a boolean if a field has been set.

### SetObjectSnapshotInfoNil

`func (o *ObjectSearchResults) SetObjectSnapshotInfoNil(b bool)`

 SetObjectSnapshotInfoNil sets the value for ObjectSnapshotInfo to be an explicit nil

### UnsetObjectSnapshotInfo
`func (o *ObjectSearchResults) UnsetObjectSnapshotInfo()`

UnsetObjectSnapshotInfo ensures that no value is present for ObjectSnapshotInfo, not even an explicit nil
### GetTotalCount

`func (o *ObjectSearchResults) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ObjectSearchResults) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ObjectSearchResults) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ObjectSearchResults) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### SetTotalCountNil

`func (o *ObjectSearchResults) SetTotalCountNil(b bool)`

 SetTotalCountNil sets the value for TotalCount to be an explicit nil

### UnsetTotalCount
`func (o *ObjectSearchResults) UnsetTotalCount()`

UnsetTotalCount ensures that no value is present for TotalCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


