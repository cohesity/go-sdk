# HealthTile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityBytes** | Pointer to **NullableInt64** | Raw Cluster Capacity in Bytes. This is not usable capacity and does not take replication factor into account. | [optional] 
**ClusterCloudUsageBytes** | Pointer to **NullableInt64** | Usage in Bytes on the cloud. | [optional] 
**LastDayAlerts** | Pointer to [**[]Alert**](Alert.md) | Alerts in last 24 hours. | [optional] 
**LastDayNumCriticals** | Pointer to **NullableInt64** | Number of Critical Alerts. | [optional] 
**LastDayNumWarnings** | Pointer to **NullableInt64** | Number of Warning Alerts. | [optional] 
**NumNodes** | Pointer to **NullableInt32** | Number of nodes in the cluster. | [optional] 
**NumNodesWithIssues** | Pointer to **NullableInt32** | Number of nodes in the cluster that are unhealthy. | [optional] 
**PercentFull** | Pointer to **NullableFloat32** | Percent the cluster is full. | [optional] 
**RawUsedBytes** | Pointer to **NullableInt64** | Raw Bytes used in the cluster. | [optional] 

## Methods

### NewHealthTile

`func NewHealthTile() *HealthTile`

NewHealthTile instantiates a new HealthTile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthTileWithDefaults

`func NewHealthTileWithDefaults() *HealthTile`

NewHealthTileWithDefaults instantiates a new HealthTile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityBytes

`func (o *HealthTile) GetCapacityBytes() int64`

GetCapacityBytes returns the CapacityBytes field if non-nil, zero value otherwise.

### GetCapacityBytesOk

`func (o *HealthTile) GetCapacityBytesOk() (*int64, bool)`

GetCapacityBytesOk returns a tuple with the CapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityBytes

`func (o *HealthTile) SetCapacityBytes(v int64)`

SetCapacityBytes sets CapacityBytes field to given value.

### HasCapacityBytes

`func (o *HealthTile) HasCapacityBytes() bool`

HasCapacityBytes returns a boolean if a field has been set.

### SetCapacityBytesNil

`func (o *HealthTile) SetCapacityBytesNil(b bool)`

 SetCapacityBytesNil sets the value for CapacityBytes to be an explicit nil

### UnsetCapacityBytes
`func (o *HealthTile) UnsetCapacityBytes()`

UnsetCapacityBytes ensures that no value is present for CapacityBytes, not even an explicit nil
### GetClusterCloudUsageBytes

`func (o *HealthTile) GetClusterCloudUsageBytes() int64`

GetClusterCloudUsageBytes returns the ClusterCloudUsageBytes field if non-nil, zero value otherwise.

### GetClusterCloudUsageBytesOk

`func (o *HealthTile) GetClusterCloudUsageBytesOk() (*int64, bool)`

GetClusterCloudUsageBytesOk returns a tuple with the ClusterCloudUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCloudUsageBytes

`func (o *HealthTile) SetClusterCloudUsageBytes(v int64)`

SetClusterCloudUsageBytes sets ClusterCloudUsageBytes field to given value.

### HasClusterCloudUsageBytes

`func (o *HealthTile) HasClusterCloudUsageBytes() bool`

HasClusterCloudUsageBytes returns a boolean if a field has been set.

### SetClusterCloudUsageBytesNil

`func (o *HealthTile) SetClusterCloudUsageBytesNil(b bool)`

 SetClusterCloudUsageBytesNil sets the value for ClusterCloudUsageBytes to be an explicit nil

### UnsetClusterCloudUsageBytes
`func (o *HealthTile) UnsetClusterCloudUsageBytes()`

UnsetClusterCloudUsageBytes ensures that no value is present for ClusterCloudUsageBytes, not even an explicit nil
### GetLastDayAlerts

`func (o *HealthTile) GetLastDayAlerts() []Alert`

GetLastDayAlerts returns the LastDayAlerts field if non-nil, zero value otherwise.

### GetLastDayAlertsOk

`func (o *HealthTile) GetLastDayAlertsOk() (*[]Alert, bool)`

GetLastDayAlertsOk returns a tuple with the LastDayAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDayAlerts

`func (o *HealthTile) SetLastDayAlerts(v []Alert)`

SetLastDayAlerts sets LastDayAlerts field to given value.

### HasLastDayAlerts

`func (o *HealthTile) HasLastDayAlerts() bool`

HasLastDayAlerts returns a boolean if a field has been set.

### SetLastDayAlertsNil

`func (o *HealthTile) SetLastDayAlertsNil(b bool)`

 SetLastDayAlertsNil sets the value for LastDayAlerts to be an explicit nil

### UnsetLastDayAlerts
`func (o *HealthTile) UnsetLastDayAlerts()`

UnsetLastDayAlerts ensures that no value is present for LastDayAlerts, not even an explicit nil
### GetLastDayNumCriticals

`func (o *HealthTile) GetLastDayNumCriticals() int64`

GetLastDayNumCriticals returns the LastDayNumCriticals field if non-nil, zero value otherwise.

### GetLastDayNumCriticalsOk

`func (o *HealthTile) GetLastDayNumCriticalsOk() (*int64, bool)`

GetLastDayNumCriticalsOk returns a tuple with the LastDayNumCriticals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDayNumCriticals

`func (o *HealthTile) SetLastDayNumCriticals(v int64)`

SetLastDayNumCriticals sets LastDayNumCriticals field to given value.

### HasLastDayNumCriticals

`func (o *HealthTile) HasLastDayNumCriticals() bool`

HasLastDayNumCriticals returns a boolean if a field has been set.

### SetLastDayNumCriticalsNil

`func (o *HealthTile) SetLastDayNumCriticalsNil(b bool)`

 SetLastDayNumCriticalsNil sets the value for LastDayNumCriticals to be an explicit nil

### UnsetLastDayNumCriticals
`func (o *HealthTile) UnsetLastDayNumCriticals()`

UnsetLastDayNumCriticals ensures that no value is present for LastDayNumCriticals, not even an explicit nil
### GetLastDayNumWarnings

`func (o *HealthTile) GetLastDayNumWarnings() int64`

GetLastDayNumWarnings returns the LastDayNumWarnings field if non-nil, zero value otherwise.

### GetLastDayNumWarningsOk

`func (o *HealthTile) GetLastDayNumWarningsOk() (*int64, bool)`

GetLastDayNumWarningsOk returns a tuple with the LastDayNumWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDayNumWarnings

`func (o *HealthTile) SetLastDayNumWarnings(v int64)`

SetLastDayNumWarnings sets LastDayNumWarnings field to given value.

### HasLastDayNumWarnings

`func (o *HealthTile) HasLastDayNumWarnings() bool`

HasLastDayNumWarnings returns a boolean if a field has been set.

### SetLastDayNumWarningsNil

`func (o *HealthTile) SetLastDayNumWarningsNil(b bool)`

 SetLastDayNumWarningsNil sets the value for LastDayNumWarnings to be an explicit nil

### UnsetLastDayNumWarnings
`func (o *HealthTile) UnsetLastDayNumWarnings()`

UnsetLastDayNumWarnings ensures that no value is present for LastDayNumWarnings, not even an explicit nil
### GetNumNodes

`func (o *HealthTile) GetNumNodes() int32`

GetNumNodes returns the NumNodes field if non-nil, zero value otherwise.

### GetNumNodesOk

`func (o *HealthTile) GetNumNodesOk() (*int32, bool)`

GetNumNodesOk returns a tuple with the NumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNodes

`func (o *HealthTile) SetNumNodes(v int32)`

SetNumNodes sets NumNodes field to given value.

### HasNumNodes

`func (o *HealthTile) HasNumNodes() bool`

HasNumNodes returns a boolean if a field has been set.

### SetNumNodesNil

`func (o *HealthTile) SetNumNodesNil(b bool)`

 SetNumNodesNil sets the value for NumNodes to be an explicit nil

### UnsetNumNodes
`func (o *HealthTile) UnsetNumNodes()`

UnsetNumNodes ensures that no value is present for NumNodes, not even an explicit nil
### GetNumNodesWithIssues

`func (o *HealthTile) GetNumNodesWithIssues() int32`

GetNumNodesWithIssues returns the NumNodesWithIssues field if non-nil, zero value otherwise.

### GetNumNodesWithIssuesOk

`func (o *HealthTile) GetNumNodesWithIssuesOk() (*int32, bool)`

GetNumNodesWithIssuesOk returns a tuple with the NumNodesWithIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNodesWithIssues

`func (o *HealthTile) SetNumNodesWithIssues(v int32)`

SetNumNodesWithIssues sets NumNodesWithIssues field to given value.

### HasNumNodesWithIssues

`func (o *HealthTile) HasNumNodesWithIssues() bool`

HasNumNodesWithIssues returns a boolean if a field has been set.

### SetNumNodesWithIssuesNil

`func (o *HealthTile) SetNumNodesWithIssuesNil(b bool)`

 SetNumNodesWithIssuesNil sets the value for NumNodesWithIssues to be an explicit nil

### UnsetNumNodesWithIssues
`func (o *HealthTile) UnsetNumNodesWithIssues()`

UnsetNumNodesWithIssues ensures that no value is present for NumNodesWithIssues, not even an explicit nil
### GetPercentFull

`func (o *HealthTile) GetPercentFull() float32`

GetPercentFull returns the PercentFull field if non-nil, zero value otherwise.

### GetPercentFullOk

`func (o *HealthTile) GetPercentFullOk() (*float32, bool)`

GetPercentFullOk returns a tuple with the PercentFull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentFull

`func (o *HealthTile) SetPercentFull(v float32)`

SetPercentFull sets PercentFull field to given value.

### HasPercentFull

`func (o *HealthTile) HasPercentFull() bool`

HasPercentFull returns a boolean if a field has been set.

### SetPercentFullNil

`func (o *HealthTile) SetPercentFullNil(b bool)`

 SetPercentFullNil sets the value for PercentFull to be an explicit nil

### UnsetPercentFull
`func (o *HealthTile) UnsetPercentFull()`

UnsetPercentFull ensures that no value is present for PercentFull, not even an explicit nil
### GetRawUsedBytes

`func (o *HealthTile) GetRawUsedBytes() int64`

GetRawUsedBytes returns the RawUsedBytes field if non-nil, zero value otherwise.

### GetRawUsedBytesOk

`func (o *HealthTile) GetRawUsedBytesOk() (*int64, bool)`

GetRawUsedBytesOk returns a tuple with the RawUsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawUsedBytes

`func (o *HealthTile) SetRawUsedBytes(v int64)`

SetRawUsedBytes sets RawUsedBytes field to given value.

### HasRawUsedBytes

`func (o *HealthTile) HasRawUsedBytes() bool`

HasRawUsedBytes returns a boolean if a field has been set.

### SetRawUsedBytesNil

`func (o *HealthTile) SetRawUsedBytesNil(b bool)`

 SetRawUsedBytesNil sets the value for RawUsedBytes to be an explicit nil

### UnsetRawUsedBytes
`func (o *HealthTile) UnsetRawUsedBytes()`

UnsetRawUsedBytes ensures that no value is present for RawUsedBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


