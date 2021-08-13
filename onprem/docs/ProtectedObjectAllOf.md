# ProtectedObjectAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceInfo** | Pointer to [**NullableObjectSummary**](ObjectSummary.md) | Specifies the Source Object information. | [optional] 
**LatestSnapshotsInfo** | Pointer to [**[]ObjectSnapshotsInfo**](ObjectSnapshotsInfo.md) | Specifies the latest snapshot information for every Protection Group for a given object. | [optional] 

## Methods

### NewProtectedObjectAllOf

`func NewProtectedObjectAllOf() *ProtectedObjectAllOf`

NewProtectedObjectAllOf instantiates a new ProtectedObjectAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedObjectAllOfWithDefaults

`func NewProtectedObjectAllOfWithDefaults() *ProtectedObjectAllOf`

NewProtectedObjectAllOfWithDefaults instantiates a new ProtectedObjectAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceInfo

`func (o *ProtectedObjectAllOf) GetSourceInfo() ObjectSummary`

GetSourceInfo returns the SourceInfo field if non-nil, zero value otherwise.

### GetSourceInfoOk

`func (o *ProtectedObjectAllOf) GetSourceInfoOk() (*ObjectSummary, bool)`

GetSourceInfoOk returns a tuple with the SourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInfo

`func (o *ProtectedObjectAllOf) SetSourceInfo(v ObjectSummary)`

SetSourceInfo sets SourceInfo field to given value.

### HasSourceInfo

`func (o *ProtectedObjectAllOf) HasSourceInfo() bool`

HasSourceInfo returns a boolean if a field has been set.

### SetSourceInfoNil

`func (o *ProtectedObjectAllOf) SetSourceInfoNil(b bool)`

 SetSourceInfoNil sets the value for SourceInfo to be an explicit nil

### UnsetSourceInfo
`func (o *ProtectedObjectAllOf) UnsetSourceInfo()`

UnsetSourceInfo ensures that no value is present for SourceInfo, not even an explicit nil
### GetLatestSnapshotsInfo

`func (o *ProtectedObjectAllOf) GetLatestSnapshotsInfo() []ObjectSnapshotsInfo`

GetLatestSnapshotsInfo returns the LatestSnapshotsInfo field if non-nil, zero value otherwise.

### GetLatestSnapshotsInfoOk

`func (o *ProtectedObjectAllOf) GetLatestSnapshotsInfoOk() (*[]ObjectSnapshotsInfo, bool)`

GetLatestSnapshotsInfoOk returns a tuple with the LatestSnapshotsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSnapshotsInfo

`func (o *ProtectedObjectAllOf) SetLatestSnapshotsInfo(v []ObjectSnapshotsInfo)`

SetLatestSnapshotsInfo sets LatestSnapshotsInfo field to given value.

### HasLatestSnapshotsInfo

`func (o *ProtectedObjectAllOf) HasLatestSnapshotsInfo() bool`

HasLatestSnapshotsInfo returns a boolean if a field has been set.

### SetLatestSnapshotsInfoNil

`func (o *ProtectedObjectAllOf) SetLatestSnapshotsInfoNil(b bool)`

 SetLatestSnapshotsInfoNil sets the value for LatestSnapshotsInfo to be an explicit nil

### UnsetLatestSnapshotsInfo
`func (o *ProtectedObjectAllOf) UnsetLatestSnapshotsInfo()`

UnsetLatestSnapshotsInfo ensures that no value is present for LatestSnapshotsInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


