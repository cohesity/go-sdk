# ClusterSWUpdateHistoryEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTimestampSecs** | Pointer to **NullableInt64** | Latest Unix epoch timestamp (in seconds) when the upgrade/patch was done - will contain the value from the node that got updated last.  | [optional] 
**NodeHistory** | Pointer to [**[]ClusterSWUpdateNodeHistoryEvent**](ClusterSWUpdateNodeHistoryEvent.md) |  | [optional] 
**OperationType** | Pointer to **string** | Type of operation.  | [optional] 
**PackageSubType** | Pointer to **NullableString** | Sub-type of package - Security Patch or Product Patch | [optional] 
**PackageType** | Pointer to **string** | Type of the package | [optional] 
**ReleaseDate** | Pointer to **NullableTime** | Release date of the package. | [optional] 
**ReleaseVersion** | Pointer to **string** | Release version of the package. Examples: For upgrade package: &#39;6.6.0d_u6&#39;, &#39;7.0.&#39; For patch package - &#39;6.8.1-p1s1&#39;  | [optional] 
**VersionName** | Pointer to **NullableString** | Name of the package version. Example: &#39;6.6.0d_u6_release-20210714_0fad884e&#39;,   &#39;7.0.1_release-20230623_ddbb8c79&#39; for upgrade packages, &#39;6.8.1-p1s1-2023Jun26-221b8a5c&#39; for patch packages  | [optional] 

## Methods

### NewClusterSWUpdateHistoryEvent

`func NewClusterSWUpdateHistoryEvent() *ClusterSWUpdateHistoryEvent`

NewClusterSWUpdateHistoryEvent instantiates a new ClusterSWUpdateHistoryEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSWUpdateHistoryEventWithDefaults

`func NewClusterSWUpdateHistoryEventWithDefaults() *ClusterSWUpdateHistoryEvent`

NewClusterSWUpdateHistoryEventWithDefaults instantiates a new ClusterSWUpdateHistoryEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventTimestampSecs

`func (o *ClusterSWUpdateHistoryEvent) GetEventTimestampSecs() int64`

GetEventTimestampSecs returns the EventTimestampSecs field if non-nil, zero value otherwise.

### GetEventTimestampSecsOk

`func (o *ClusterSWUpdateHistoryEvent) GetEventTimestampSecsOk() (*int64, bool)`

GetEventTimestampSecsOk returns a tuple with the EventTimestampSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestampSecs

`func (o *ClusterSWUpdateHistoryEvent) SetEventTimestampSecs(v int64)`

SetEventTimestampSecs sets EventTimestampSecs field to given value.

### HasEventTimestampSecs

`func (o *ClusterSWUpdateHistoryEvent) HasEventTimestampSecs() bool`

HasEventTimestampSecs returns a boolean if a field has been set.

### SetEventTimestampSecsNil

`func (o *ClusterSWUpdateHistoryEvent) SetEventTimestampSecsNil(b bool)`

 SetEventTimestampSecsNil sets the value for EventTimestampSecs to be an explicit nil

### UnsetEventTimestampSecs
`func (o *ClusterSWUpdateHistoryEvent) UnsetEventTimestampSecs()`

UnsetEventTimestampSecs ensures that no value is present for EventTimestampSecs, not even an explicit nil
### GetNodeHistory

`func (o *ClusterSWUpdateHistoryEvent) GetNodeHistory() []ClusterSWUpdateNodeHistoryEvent`

GetNodeHistory returns the NodeHistory field if non-nil, zero value otherwise.

### GetNodeHistoryOk

`func (o *ClusterSWUpdateHistoryEvent) GetNodeHistoryOk() (*[]ClusterSWUpdateNodeHistoryEvent, bool)`

GetNodeHistoryOk returns a tuple with the NodeHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeHistory

`func (o *ClusterSWUpdateHistoryEvent) SetNodeHistory(v []ClusterSWUpdateNodeHistoryEvent)`

SetNodeHistory sets NodeHistory field to given value.

### HasNodeHistory

`func (o *ClusterSWUpdateHistoryEvent) HasNodeHistory() bool`

HasNodeHistory returns a boolean if a field has been set.

### GetOperationType

`func (o *ClusterSWUpdateHistoryEvent) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *ClusterSWUpdateHistoryEvent) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *ClusterSWUpdateHistoryEvent) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *ClusterSWUpdateHistoryEvent) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetPackageSubType

`func (o *ClusterSWUpdateHistoryEvent) GetPackageSubType() string`

GetPackageSubType returns the PackageSubType field if non-nil, zero value otherwise.

### GetPackageSubTypeOk

`func (o *ClusterSWUpdateHistoryEvent) GetPackageSubTypeOk() (*string, bool)`

GetPackageSubTypeOk returns a tuple with the PackageSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageSubType

`func (o *ClusterSWUpdateHistoryEvent) SetPackageSubType(v string)`

SetPackageSubType sets PackageSubType field to given value.

### HasPackageSubType

`func (o *ClusterSWUpdateHistoryEvent) HasPackageSubType() bool`

HasPackageSubType returns a boolean if a field has been set.

### SetPackageSubTypeNil

`func (o *ClusterSWUpdateHistoryEvent) SetPackageSubTypeNil(b bool)`

 SetPackageSubTypeNil sets the value for PackageSubType to be an explicit nil

### UnsetPackageSubType
`func (o *ClusterSWUpdateHistoryEvent) UnsetPackageSubType()`

UnsetPackageSubType ensures that no value is present for PackageSubType, not even an explicit nil
### GetPackageType

`func (o *ClusterSWUpdateHistoryEvent) GetPackageType() string`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *ClusterSWUpdateHistoryEvent) GetPackageTypeOk() (*string, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *ClusterSWUpdateHistoryEvent) SetPackageType(v string)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *ClusterSWUpdateHistoryEvent) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### GetReleaseDate

`func (o *ClusterSWUpdateHistoryEvent) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *ClusterSWUpdateHistoryEvent) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *ClusterSWUpdateHistoryEvent) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *ClusterSWUpdateHistoryEvent) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *ClusterSWUpdateHistoryEvent) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *ClusterSWUpdateHistoryEvent) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetReleaseVersion

`func (o *ClusterSWUpdateHistoryEvent) GetReleaseVersion() string`

GetReleaseVersion returns the ReleaseVersion field if non-nil, zero value otherwise.

### GetReleaseVersionOk

`func (o *ClusterSWUpdateHistoryEvent) GetReleaseVersionOk() (*string, bool)`

GetReleaseVersionOk returns a tuple with the ReleaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVersion

`func (o *ClusterSWUpdateHistoryEvent) SetReleaseVersion(v string)`

SetReleaseVersion sets ReleaseVersion field to given value.

### HasReleaseVersion

`func (o *ClusterSWUpdateHistoryEvent) HasReleaseVersion() bool`

HasReleaseVersion returns a boolean if a field has been set.

### GetVersionName

`func (o *ClusterSWUpdateHistoryEvent) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *ClusterSWUpdateHistoryEvent) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *ClusterSWUpdateHistoryEvent) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *ClusterSWUpdateHistoryEvent) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.

### SetVersionNameNil

`func (o *ClusterSWUpdateHistoryEvent) SetVersionNameNil(b bool)`

 SetVersionNameNil sets the value for VersionName to be an explicit nil

### UnsetVersionName
`func (o *ClusterSWUpdateHistoryEvent) UnsetVersionName()`

UnsetVersionName ensures that no value is present for VersionName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


