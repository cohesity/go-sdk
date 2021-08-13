# McmSourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **NullableInt64** | Specifies the cluster id. | [optional] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the cluster incarnation id. | [optional] 
**RegionId** | Pointer to **NullableString** | Specifies the region id. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the Protection Source. | [optional] 
**RegistrationId** | Pointer to **NullableString** | Specifies the registration id of the Protection Source. | [optional] 
**RegistrationDetails** | Pointer to [**McmSourceRegistrationInfo**](McmSourceRegistrationInfo.md) |  | [optional] 
**Applications** | Pointer to **[]string** | Specifies the list of applications registered with current Source. | [optional] 
**ProtectionStats** | Pointer to [**[]ObjectProtectionStatsSummary**](ObjectProtectionStatsSummary.md) | Specifies the protection statistics of the Source. | [optional] 
**PhysicalSourceInfo** | Pointer to [**McmPhysicalSourceInfo**](McmPhysicalSourceInfo.md) |  | [optional] 

## Methods

### NewMcmSourceInfo

`func NewMcmSourceInfo() *McmSourceInfo`

NewMcmSourceInfo instantiates a new McmSourceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmSourceInfoWithDefaults

`func NewMcmSourceInfoWithDefaults() *McmSourceInfo`

NewMcmSourceInfoWithDefaults instantiates a new McmSourceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *McmSourceInfo) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *McmSourceInfo) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *McmSourceInfo) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *McmSourceInfo) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *McmSourceInfo) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *McmSourceInfo) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *McmSourceInfo) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *McmSourceInfo) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *McmSourceInfo) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *McmSourceInfo) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *McmSourceInfo) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *McmSourceInfo) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetRegionId

`func (o *McmSourceInfo) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *McmSourceInfo) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *McmSourceInfo) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *McmSourceInfo) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *McmSourceInfo) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *McmSourceInfo) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetSourceId

`func (o *McmSourceInfo) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *McmSourceInfo) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *McmSourceInfo) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *McmSourceInfo) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *McmSourceInfo) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *McmSourceInfo) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetRegistrationId

`func (o *McmSourceInfo) GetRegistrationId() string`

GetRegistrationId returns the RegistrationId field if non-nil, zero value otherwise.

### GetRegistrationIdOk

`func (o *McmSourceInfo) GetRegistrationIdOk() (*string, bool)`

GetRegistrationIdOk returns a tuple with the RegistrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationId

`func (o *McmSourceInfo) SetRegistrationId(v string)`

SetRegistrationId sets RegistrationId field to given value.

### HasRegistrationId

`func (o *McmSourceInfo) HasRegistrationId() bool`

HasRegistrationId returns a boolean if a field has been set.

### SetRegistrationIdNil

`func (o *McmSourceInfo) SetRegistrationIdNil(b bool)`

 SetRegistrationIdNil sets the value for RegistrationId to be an explicit nil

### UnsetRegistrationId
`func (o *McmSourceInfo) UnsetRegistrationId()`

UnsetRegistrationId ensures that no value is present for RegistrationId, not even an explicit nil
### GetRegistrationDetails

`func (o *McmSourceInfo) GetRegistrationDetails() McmSourceRegistrationInfo`

GetRegistrationDetails returns the RegistrationDetails field if non-nil, zero value otherwise.

### GetRegistrationDetailsOk

`func (o *McmSourceInfo) GetRegistrationDetailsOk() (*McmSourceRegistrationInfo, bool)`

GetRegistrationDetailsOk returns a tuple with the RegistrationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationDetails

`func (o *McmSourceInfo) SetRegistrationDetails(v McmSourceRegistrationInfo)`

SetRegistrationDetails sets RegistrationDetails field to given value.

### HasRegistrationDetails

`func (o *McmSourceInfo) HasRegistrationDetails() bool`

HasRegistrationDetails returns a boolean if a field has been set.

### GetApplications

`func (o *McmSourceInfo) GetApplications() []string`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *McmSourceInfo) GetApplicationsOk() (*[]string, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *McmSourceInfo) SetApplications(v []string)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *McmSourceInfo) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### SetApplicationsNil

`func (o *McmSourceInfo) SetApplicationsNil(b bool)`

 SetApplicationsNil sets the value for Applications to be an explicit nil

### UnsetApplications
`func (o *McmSourceInfo) UnsetApplications()`

UnsetApplications ensures that no value is present for Applications, not even an explicit nil
### GetProtectionStats

`func (o *McmSourceInfo) GetProtectionStats() []ObjectProtectionStatsSummary`

GetProtectionStats returns the ProtectionStats field if non-nil, zero value otherwise.

### GetProtectionStatsOk

`func (o *McmSourceInfo) GetProtectionStatsOk() (*[]ObjectProtectionStatsSummary, bool)`

GetProtectionStatsOk returns a tuple with the ProtectionStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionStats

`func (o *McmSourceInfo) SetProtectionStats(v []ObjectProtectionStatsSummary)`

SetProtectionStats sets ProtectionStats field to given value.

### HasProtectionStats

`func (o *McmSourceInfo) HasProtectionStats() bool`

HasProtectionStats returns a boolean if a field has been set.

### SetProtectionStatsNil

`func (o *McmSourceInfo) SetProtectionStatsNil(b bool)`

 SetProtectionStatsNil sets the value for ProtectionStats to be an explicit nil

### UnsetProtectionStats
`func (o *McmSourceInfo) UnsetProtectionStats()`

UnsetProtectionStats ensures that no value is present for ProtectionStats, not even an explicit nil
### GetPhysicalSourceInfo

`func (o *McmSourceInfo) GetPhysicalSourceInfo() McmPhysicalSourceInfo`

GetPhysicalSourceInfo returns the PhysicalSourceInfo field if non-nil, zero value otherwise.

### GetPhysicalSourceInfoOk

`func (o *McmSourceInfo) GetPhysicalSourceInfoOk() (*McmPhysicalSourceInfo, bool)`

GetPhysicalSourceInfoOk returns a tuple with the PhysicalSourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalSourceInfo

`func (o *McmSourceInfo) SetPhysicalSourceInfo(v McmPhysicalSourceInfo)`

SetPhysicalSourceInfo sets PhysicalSourceInfo field to given value.

### HasPhysicalSourceInfo

`func (o *McmSourceInfo) HasPhysicalSourceInfo() bool`

HasPhysicalSourceInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


