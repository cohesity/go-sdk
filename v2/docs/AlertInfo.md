# AlertInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertCategory** | Pointer to **NullableString** | Specifies the alert category. | [optional] 
**AlertCode** | Pointer to **NullableString** | Specifies a unique code that categorizes the Alert, for example: CE00200014, where CE stands for Cohesity Error, the alert state next 3 digits is the id of the Alert Category (e.g. 002 for &#39;kNode&#39;) and the last 5 digits is the id of the Alert Type (e.g. 00014 for &#39;kNodeHighCpuUsage&#39;). | [optional] 
**AlertDocument** | Pointer to [**AlertDocument**](AlertDocument.md) |  | [optional] 
**AlertState** | Pointer to **NullableString** | Specifies the alert state. | [optional] 
**AlertType** | Pointer to **NullableInt32** | Specifies the alert type. | [optional] 
**AlertTypeBucket** | Pointer to **NullableString** | Specifies the Alert type bucket. | [optional] 
**ClusterId** | Pointer to **NullableInt64** | Id of the cluster which the alert is associated | [optional] 
**ClusterName** | Pointer to **NullableString** | Specifies the name of cluster which alert is raised from. | [optional] 
**DedupCount** | Pointer to **NullableInt32** | Specifies the dedup count of alert. | [optional] 
**DedupTimestamps** | Pointer to **[]int64** | Specifies Unix epoch Timestamps (in microseconds) for the last 25 occurrences of duplicated Alerts that are stored with the original/primary Alert. Alerts are grouped into one Alert if the Alerts are the same type, are reporting on the same Object and occur within one hour. &#39;dedupCount&#39; always reports the total count of duplicated Alerts even if there are more than 25 occurrences. For example, if there are 100 occurrences of this Alert, dedupTimestamps stores the timestamps of the last 25 occurrences and dedupCount equals 100. | [optional] 
**EventSource** | Pointer to **NullableString** | Specifies source where the event occurred. | [optional] 
**FirstTimestampUsecs** | Pointer to **NullableInt64** | Specifies Unix epoch Timestamp (in microseconds) of the first occurrence of the Alert. | [optional] 
**Id** | Pointer to **NullableString** | Specifies unique id of the alert. | [optional] 
**LabelIds** | Pointer to **[]string** | Specifies the labels for which this alert has been raised. | [optional] 
**LatestTimestampUsecs** | Pointer to **NullableInt64** | Specifies Unix epoch Timestamp (in microseconds) of the most recent occurrence of the Alert. | [optional] 
**PropertyList** | Pointer to [**[]Label**](Label.md) | List of property key and values associated with alert | [optional] 
**RegionId** | Pointer to **NullableString** | Specifies the region id of the alert. | [optional] 
**ResolutionDetails** | Pointer to [**NullableAlertResolutionDetails**](AlertResolutionDetails.md) |  | [optional] 
**ResolutionIdString** | Pointer to **NullableString** | Resolution Id String. | [optional] 
**ResolvedTimestampUsecs** | Pointer to **NullableInt64** | Specifies Unix epoch Timestamps in microseconds when alert is resolved. | [optional] 
**Severity** | Pointer to **NullableString** | Specifies the alert severity. | [optional] 
**SuppressionId** | Pointer to **NullableInt64** | Specifies unique id generated when the Alert is suppressed by the admin. | [optional] 
**TenantIds** | Pointer to **[]string** | Specifies the tenants for which this alert has been raised. | [optional] 
**Vaults** | Pointer to [**[]Vault**](Vault.md) | Specifies information about vaults where source object associated with alert is vaulted. This could be empty if alert is not related to any source object or it is not vaulted. | [optional] 

## Methods

### NewAlertInfo

`func NewAlertInfo() *AlertInfo`

NewAlertInfo instantiates a new AlertInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertInfoWithDefaults

`func NewAlertInfoWithDefaults() *AlertInfo`

NewAlertInfoWithDefaults instantiates a new AlertInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertCategory

`func (o *AlertInfo) GetAlertCategory() string`

GetAlertCategory returns the AlertCategory field if non-nil, zero value otherwise.

### GetAlertCategoryOk

`func (o *AlertInfo) GetAlertCategoryOk() (*string, bool)`

GetAlertCategoryOk returns a tuple with the AlertCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCategory

`func (o *AlertInfo) SetAlertCategory(v string)`

SetAlertCategory sets AlertCategory field to given value.

### HasAlertCategory

`func (o *AlertInfo) HasAlertCategory() bool`

HasAlertCategory returns a boolean if a field has been set.

### SetAlertCategoryNil

`func (o *AlertInfo) SetAlertCategoryNil(b bool)`

 SetAlertCategoryNil sets the value for AlertCategory to be an explicit nil

### UnsetAlertCategory
`func (o *AlertInfo) UnsetAlertCategory()`

UnsetAlertCategory ensures that no value is present for AlertCategory, not even an explicit nil
### GetAlertCode

`func (o *AlertInfo) GetAlertCode() string`

GetAlertCode returns the AlertCode field if non-nil, zero value otherwise.

### GetAlertCodeOk

`func (o *AlertInfo) GetAlertCodeOk() (*string, bool)`

GetAlertCodeOk returns a tuple with the AlertCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCode

`func (o *AlertInfo) SetAlertCode(v string)`

SetAlertCode sets AlertCode field to given value.

### HasAlertCode

`func (o *AlertInfo) HasAlertCode() bool`

HasAlertCode returns a boolean if a field has been set.

### SetAlertCodeNil

`func (o *AlertInfo) SetAlertCodeNil(b bool)`

 SetAlertCodeNil sets the value for AlertCode to be an explicit nil

### UnsetAlertCode
`func (o *AlertInfo) UnsetAlertCode()`

UnsetAlertCode ensures that no value is present for AlertCode, not even an explicit nil
### GetAlertDocument

`func (o *AlertInfo) GetAlertDocument() AlertDocument`

GetAlertDocument returns the AlertDocument field if non-nil, zero value otherwise.

### GetAlertDocumentOk

`func (o *AlertInfo) GetAlertDocumentOk() (*AlertDocument, bool)`

GetAlertDocumentOk returns a tuple with the AlertDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertDocument

`func (o *AlertInfo) SetAlertDocument(v AlertDocument)`

SetAlertDocument sets AlertDocument field to given value.

### HasAlertDocument

`func (o *AlertInfo) HasAlertDocument() bool`

HasAlertDocument returns a boolean if a field has been set.

### GetAlertState

`func (o *AlertInfo) GetAlertState() string`

GetAlertState returns the AlertState field if non-nil, zero value otherwise.

### GetAlertStateOk

`func (o *AlertInfo) GetAlertStateOk() (*string, bool)`

GetAlertStateOk returns a tuple with the AlertState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertState

`func (o *AlertInfo) SetAlertState(v string)`

SetAlertState sets AlertState field to given value.

### HasAlertState

`func (o *AlertInfo) HasAlertState() bool`

HasAlertState returns a boolean if a field has been set.

### SetAlertStateNil

`func (o *AlertInfo) SetAlertStateNil(b bool)`

 SetAlertStateNil sets the value for AlertState to be an explicit nil

### UnsetAlertState
`func (o *AlertInfo) UnsetAlertState()`

UnsetAlertState ensures that no value is present for AlertState, not even an explicit nil
### GetAlertType

`func (o *AlertInfo) GetAlertType() int32`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *AlertInfo) GetAlertTypeOk() (*int32, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *AlertInfo) SetAlertType(v int32)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *AlertInfo) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### SetAlertTypeNil

`func (o *AlertInfo) SetAlertTypeNil(b bool)`

 SetAlertTypeNil sets the value for AlertType to be an explicit nil

### UnsetAlertType
`func (o *AlertInfo) UnsetAlertType()`

UnsetAlertType ensures that no value is present for AlertType, not even an explicit nil
### GetAlertTypeBucket

`func (o *AlertInfo) GetAlertTypeBucket() string`

GetAlertTypeBucket returns the AlertTypeBucket field if non-nil, zero value otherwise.

### GetAlertTypeBucketOk

`func (o *AlertInfo) GetAlertTypeBucketOk() (*string, bool)`

GetAlertTypeBucketOk returns a tuple with the AlertTypeBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTypeBucket

`func (o *AlertInfo) SetAlertTypeBucket(v string)`

SetAlertTypeBucket sets AlertTypeBucket field to given value.

### HasAlertTypeBucket

`func (o *AlertInfo) HasAlertTypeBucket() bool`

HasAlertTypeBucket returns a boolean if a field has been set.

### SetAlertTypeBucketNil

`func (o *AlertInfo) SetAlertTypeBucketNil(b bool)`

 SetAlertTypeBucketNil sets the value for AlertTypeBucket to be an explicit nil

### UnsetAlertTypeBucket
`func (o *AlertInfo) UnsetAlertTypeBucket()`

UnsetAlertTypeBucket ensures that no value is present for AlertTypeBucket, not even an explicit nil
### GetClusterId

`func (o *AlertInfo) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *AlertInfo) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *AlertInfo) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *AlertInfo) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *AlertInfo) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *AlertInfo) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterName

`func (o *AlertInfo) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *AlertInfo) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *AlertInfo) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *AlertInfo) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### SetClusterNameNil

`func (o *AlertInfo) SetClusterNameNil(b bool)`

 SetClusterNameNil sets the value for ClusterName to be an explicit nil

### UnsetClusterName
`func (o *AlertInfo) UnsetClusterName()`

UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil
### GetDedupCount

`func (o *AlertInfo) GetDedupCount() int32`

GetDedupCount returns the DedupCount field if non-nil, zero value otherwise.

### GetDedupCountOk

`func (o *AlertInfo) GetDedupCountOk() (*int32, bool)`

GetDedupCountOk returns a tuple with the DedupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupCount

`func (o *AlertInfo) SetDedupCount(v int32)`

SetDedupCount sets DedupCount field to given value.

### HasDedupCount

`func (o *AlertInfo) HasDedupCount() bool`

HasDedupCount returns a boolean if a field has been set.

### SetDedupCountNil

`func (o *AlertInfo) SetDedupCountNil(b bool)`

 SetDedupCountNil sets the value for DedupCount to be an explicit nil

### UnsetDedupCount
`func (o *AlertInfo) UnsetDedupCount()`

UnsetDedupCount ensures that no value is present for DedupCount, not even an explicit nil
### GetDedupTimestamps

`func (o *AlertInfo) GetDedupTimestamps() []int64`

GetDedupTimestamps returns the DedupTimestamps field if non-nil, zero value otherwise.

### GetDedupTimestampsOk

`func (o *AlertInfo) GetDedupTimestampsOk() (*[]int64, bool)`

GetDedupTimestampsOk returns a tuple with the DedupTimestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupTimestamps

`func (o *AlertInfo) SetDedupTimestamps(v []int64)`

SetDedupTimestamps sets DedupTimestamps field to given value.

### HasDedupTimestamps

`func (o *AlertInfo) HasDedupTimestamps() bool`

HasDedupTimestamps returns a boolean if a field has been set.

### GetEventSource

`func (o *AlertInfo) GetEventSource() string`

GetEventSource returns the EventSource field if non-nil, zero value otherwise.

### GetEventSourceOk

`func (o *AlertInfo) GetEventSourceOk() (*string, bool)`

GetEventSourceOk returns a tuple with the EventSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSource

`func (o *AlertInfo) SetEventSource(v string)`

SetEventSource sets EventSource field to given value.

### HasEventSource

`func (o *AlertInfo) HasEventSource() bool`

HasEventSource returns a boolean if a field has been set.

### SetEventSourceNil

`func (o *AlertInfo) SetEventSourceNil(b bool)`

 SetEventSourceNil sets the value for EventSource to be an explicit nil

### UnsetEventSource
`func (o *AlertInfo) UnsetEventSource()`

UnsetEventSource ensures that no value is present for EventSource, not even an explicit nil
### GetFirstTimestampUsecs

`func (o *AlertInfo) GetFirstTimestampUsecs() int64`

GetFirstTimestampUsecs returns the FirstTimestampUsecs field if non-nil, zero value otherwise.

### GetFirstTimestampUsecsOk

`func (o *AlertInfo) GetFirstTimestampUsecsOk() (*int64, bool)`

GetFirstTimestampUsecsOk returns a tuple with the FirstTimestampUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTimestampUsecs

`func (o *AlertInfo) SetFirstTimestampUsecs(v int64)`

SetFirstTimestampUsecs sets FirstTimestampUsecs field to given value.

### HasFirstTimestampUsecs

`func (o *AlertInfo) HasFirstTimestampUsecs() bool`

HasFirstTimestampUsecs returns a boolean if a field has been set.

### SetFirstTimestampUsecsNil

`func (o *AlertInfo) SetFirstTimestampUsecsNil(b bool)`

 SetFirstTimestampUsecsNil sets the value for FirstTimestampUsecs to be an explicit nil

### UnsetFirstTimestampUsecs
`func (o *AlertInfo) UnsetFirstTimestampUsecs()`

UnsetFirstTimestampUsecs ensures that no value is present for FirstTimestampUsecs, not even an explicit nil
### GetId

`func (o *AlertInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AlertInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AlertInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AlertInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLabelIds

`func (o *AlertInfo) GetLabelIds() []string`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *AlertInfo) GetLabelIdsOk() (*[]string, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *AlertInfo) SetLabelIds(v []string)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *AlertInfo) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.

### SetLabelIdsNil

`func (o *AlertInfo) SetLabelIdsNil(b bool)`

 SetLabelIdsNil sets the value for LabelIds to be an explicit nil

### UnsetLabelIds
`func (o *AlertInfo) UnsetLabelIds()`

UnsetLabelIds ensures that no value is present for LabelIds, not even an explicit nil
### GetLatestTimestampUsecs

`func (o *AlertInfo) GetLatestTimestampUsecs() int64`

GetLatestTimestampUsecs returns the LatestTimestampUsecs field if non-nil, zero value otherwise.

### GetLatestTimestampUsecsOk

`func (o *AlertInfo) GetLatestTimestampUsecsOk() (*int64, bool)`

GetLatestTimestampUsecsOk returns a tuple with the LatestTimestampUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestTimestampUsecs

`func (o *AlertInfo) SetLatestTimestampUsecs(v int64)`

SetLatestTimestampUsecs sets LatestTimestampUsecs field to given value.

### HasLatestTimestampUsecs

`func (o *AlertInfo) HasLatestTimestampUsecs() bool`

HasLatestTimestampUsecs returns a boolean if a field has been set.

### SetLatestTimestampUsecsNil

`func (o *AlertInfo) SetLatestTimestampUsecsNil(b bool)`

 SetLatestTimestampUsecsNil sets the value for LatestTimestampUsecs to be an explicit nil

### UnsetLatestTimestampUsecs
`func (o *AlertInfo) UnsetLatestTimestampUsecs()`

UnsetLatestTimestampUsecs ensures that no value is present for LatestTimestampUsecs, not even an explicit nil
### GetPropertyList

`func (o *AlertInfo) GetPropertyList() []Label`

GetPropertyList returns the PropertyList field if non-nil, zero value otherwise.

### GetPropertyListOk

`func (o *AlertInfo) GetPropertyListOk() (*[]Label, bool)`

GetPropertyListOk returns a tuple with the PropertyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyList

`func (o *AlertInfo) SetPropertyList(v []Label)`

SetPropertyList sets PropertyList field to given value.

### HasPropertyList

`func (o *AlertInfo) HasPropertyList() bool`

HasPropertyList returns a boolean if a field has been set.

### GetRegionId

`func (o *AlertInfo) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *AlertInfo) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *AlertInfo) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *AlertInfo) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *AlertInfo) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *AlertInfo) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetResolutionDetails

`func (o *AlertInfo) GetResolutionDetails() AlertResolutionDetails`

GetResolutionDetails returns the ResolutionDetails field if non-nil, zero value otherwise.

### GetResolutionDetailsOk

`func (o *AlertInfo) GetResolutionDetailsOk() (*AlertResolutionDetails, bool)`

GetResolutionDetailsOk returns a tuple with the ResolutionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionDetails

`func (o *AlertInfo) SetResolutionDetails(v AlertResolutionDetails)`

SetResolutionDetails sets ResolutionDetails field to given value.

### HasResolutionDetails

`func (o *AlertInfo) HasResolutionDetails() bool`

HasResolutionDetails returns a boolean if a field has been set.

### SetResolutionDetailsNil

`func (o *AlertInfo) SetResolutionDetailsNil(b bool)`

 SetResolutionDetailsNil sets the value for ResolutionDetails to be an explicit nil

### UnsetResolutionDetails
`func (o *AlertInfo) UnsetResolutionDetails()`

UnsetResolutionDetails ensures that no value is present for ResolutionDetails, not even an explicit nil
### GetResolutionIdString

`func (o *AlertInfo) GetResolutionIdString() string`

GetResolutionIdString returns the ResolutionIdString field if non-nil, zero value otherwise.

### GetResolutionIdStringOk

`func (o *AlertInfo) GetResolutionIdStringOk() (*string, bool)`

GetResolutionIdStringOk returns a tuple with the ResolutionIdString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionIdString

`func (o *AlertInfo) SetResolutionIdString(v string)`

SetResolutionIdString sets ResolutionIdString field to given value.

### HasResolutionIdString

`func (o *AlertInfo) HasResolutionIdString() bool`

HasResolutionIdString returns a boolean if a field has been set.

### SetResolutionIdStringNil

`func (o *AlertInfo) SetResolutionIdStringNil(b bool)`

 SetResolutionIdStringNil sets the value for ResolutionIdString to be an explicit nil

### UnsetResolutionIdString
`func (o *AlertInfo) UnsetResolutionIdString()`

UnsetResolutionIdString ensures that no value is present for ResolutionIdString, not even an explicit nil
### GetResolvedTimestampUsecs

`func (o *AlertInfo) GetResolvedTimestampUsecs() int64`

GetResolvedTimestampUsecs returns the ResolvedTimestampUsecs field if non-nil, zero value otherwise.

### GetResolvedTimestampUsecsOk

`func (o *AlertInfo) GetResolvedTimestampUsecsOk() (*int64, bool)`

GetResolvedTimestampUsecsOk returns a tuple with the ResolvedTimestampUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedTimestampUsecs

`func (o *AlertInfo) SetResolvedTimestampUsecs(v int64)`

SetResolvedTimestampUsecs sets ResolvedTimestampUsecs field to given value.

### HasResolvedTimestampUsecs

`func (o *AlertInfo) HasResolvedTimestampUsecs() bool`

HasResolvedTimestampUsecs returns a boolean if a field has been set.

### SetResolvedTimestampUsecsNil

`func (o *AlertInfo) SetResolvedTimestampUsecsNil(b bool)`

 SetResolvedTimestampUsecsNil sets the value for ResolvedTimestampUsecs to be an explicit nil

### UnsetResolvedTimestampUsecs
`func (o *AlertInfo) UnsetResolvedTimestampUsecs()`

UnsetResolvedTimestampUsecs ensures that no value is present for ResolvedTimestampUsecs, not even an explicit nil
### GetSeverity

`func (o *AlertInfo) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AlertInfo) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AlertInfo) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AlertInfo) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *AlertInfo) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *AlertInfo) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
### GetSuppressionId

`func (o *AlertInfo) GetSuppressionId() int64`

GetSuppressionId returns the SuppressionId field if non-nil, zero value otherwise.

### GetSuppressionIdOk

`func (o *AlertInfo) GetSuppressionIdOk() (*int64, bool)`

GetSuppressionIdOk returns a tuple with the SuppressionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionId

`func (o *AlertInfo) SetSuppressionId(v int64)`

SetSuppressionId sets SuppressionId field to given value.

### HasSuppressionId

`func (o *AlertInfo) HasSuppressionId() bool`

HasSuppressionId returns a boolean if a field has been set.

### SetSuppressionIdNil

`func (o *AlertInfo) SetSuppressionIdNil(b bool)`

 SetSuppressionIdNil sets the value for SuppressionId to be an explicit nil

### UnsetSuppressionId
`func (o *AlertInfo) UnsetSuppressionId()`

UnsetSuppressionId ensures that no value is present for SuppressionId, not even an explicit nil
### GetTenantIds

`func (o *AlertInfo) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *AlertInfo) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *AlertInfo) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *AlertInfo) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### GetVaults

`func (o *AlertInfo) GetVaults() []Vault`

GetVaults returns the Vaults field if non-nil, zero value otherwise.

### GetVaultsOk

`func (o *AlertInfo) GetVaultsOk() (*[]Vault, bool)`

GetVaultsOk returns a tuple with the Vaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaults

`func (o *AlertInfo) SetVaults(v []Vault)`

SetVaults sets Vaults field to given value.

### HasVaults

`func (o *AlertInfo) HasVaults() bool`

HasVaults returns a boolean if a field has been set.

### SetVaultsNil

`func (o *AlertInfo) SetVaultsNil(b bool)`

 SetVaultsNil sets the value for Vaults to be an explicit nil

### UnsetVaults
`func (o *AlertInfo) UnsetVaults()`

UnsetVaults ensures that no value is present for Vaults, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


