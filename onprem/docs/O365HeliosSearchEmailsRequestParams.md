# O365HeliosSearchEmailsRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterIdentifier** | Pointer to **NullableString** | List of Clusters Identifiers to filter from. The format is clusterId:clusterIncarnationId. | [optional] 
**RegionId** | Pointer to **NullableString** | Specifies the region id of the cluster. Only valid for DMaaS clusters. | [optional] 
**DomainIds** | Pointer to **[]int64** | Specifies the domain Ids in which mailboxes are registered. | [optional] 
**MailboxIds** | Pointer to **[]int64** | Specifies the mailbox Ids which contains the emails/folders. | [optional] 

## Methods

### NewO365HeliosSearchEmailsRequestParams

`func NewO365HeliosSearchEmailsRequestParams() *O365HeliosSearchEmailsRequestParams`

NewO365HeliosSearchEmailsRequestParams instantiates a new O365HeliosSearchEmailsRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO365HeliosSearchEmailsRequestParamsWithDefaults

`func NewO365HeliosSearchEmailsRequestParamsWithDefaults() *O365HeliosSearchEmailsRequestParams`

NewO365HeliosSearchEmailsRequestParamsWithDefaults instantiates a new O365HeliosSearchEmailsRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterIdentifier

`func (o *O365HeliosSearchEmailsRequestParams) GetClusterIdentifier() string`

GetClusterIdentifier returns the ClusterIdentifier field if non-nil, zero value otherwise.

### GetClusterIdentifierOk

`func (o *O365HeliosSearchEmailsRequestParams) GetClusterIdentifierOk() (*string, bool)`

GetClusterIdentifierOk returns a tuple with the ClusterIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdentifier

`func (o *O365HeliosSearchEmailsRequestParams) SetClusterIdentifier(v string)`

SetClusterIdentifier sets ClusterIdentifier field to given value.

### HasClusterIdentifier

`func (o *O365HeliosSearchEmailsRequestParams) HasClusterIdentifier() bool`

HasClusterIdentifier returns a boolean if a field has been set.

### SetClusterIdentifierNil

`func (o *O365HeliosSearchEmailsRequestParams) SetClusterIdentifierNil(b bool)`

 SetClusterIdentifierNil sets the value for ClusterIdentifier to be an explicit nil

### UnsetClusterIdentifier
`func (o *O365HeliosSearchEmailsRequestParams) UnsetClusterIdentifier()`

UnsetClusterIdentifier ensures that no value is present for ClusterIdentifier, not even an explicit nil
### GetRegionId

`func (o *O365HeliosSearchEmailsRequestParams) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *O365HeliosSearchEmailsRequestParams) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *O365HeliosSearchEmailsRequestParams) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *O365HeliosSearchEmailsRequestParams) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *O365HeliosSearchEmailsRequestParams) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *O365HeliosSearchEmailsRequestParams) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetDomainIds

`func (o *O365HeliosSearchEmailsRequestParams) GetDomainIds() []int64`

GetDomainIds returns the DomainIds field if non-nil, zero value otherwise.

### GetDomainIdsOk

`func (o *O365HeliosSearchEmailsRequestParams) GetDomainIdsOk() (*[]int64, bool)`

GetDomainIdsOk returns a tuple with the DomainIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainIds

`func (o *O365HeliosSearchEmailsRequestParams) SetDomainIds(v []int64)`

SetDomainIds sets DomainIds field to given value.

### HasDomainIds

`func (o *O365HeliosSearchEmailsRequestParams) HasDomainIds() bool`

HasDomainIds returns a boolean if a field has been set.

### SetDomainIdsNil

`func (o *O365HeliosSearchEmailsRequestParams) SetDomainIdsNil(b bool)`

 SetDomainIdsNil sets the value for DomainIds to be an explicit nil

### UnsetDomainIds
`func (o *O365HeliosSearchEmailsRequestParams) UnsetDomainIds()`

UnsetDomainIds ensures that no value is present for DomainIds, not even an explicit nil
### GetMailboxIds

`func (o *O365HeliosSearchEmailsRequestParams) GetMailboxIds() []int64`

GetMailboxIds returns the MailboxIds field if non-nil, zero value otherwise.

### GetMailboxIdsOk

`func (o *O365HeliosSearchEmailsRequestParams) GetMailboxIdsOk() (*[]int64, bool)`

GetMailboxIdsOk returns a tuple with the MailboxIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxIds

`func (o *O365HeliosSearchEmailsRequestParams) SetMailboxIds(v []int64)`

SetMailboxIds sets MailboxIds field to given value.

### HasMailboxIds

`func (o *O365HeliosSearchEmailsRequestParams) HasMailboxIds() bool`

HasMailboxIds returns a boolean if a field has been set.

### SetMailboxIdsNil

`func (o *O365HeliosSearchEmailsRequestParams) SetMailboxIdsNil(b bool)`

 SetMailboxIdsNil sets the value for MailboxIds to be an explicit nil

### UnsetMailboxIds
`func (o *O365HeliosSearchEmailsRequestParams) UnsetMailboxIds()`

UnsetMailboxIds ensures that no value is present for MailboxIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


