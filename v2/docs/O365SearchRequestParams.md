# O365SearchRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainIds** | Pointer to **[]int64** | Specifies the domain Ids in which indexed items are searched. | [optional] 
**GroupIds** | Pointer to **[]int64** | Specifies the Group ids across which the indexed items needs to be searched. | [optional] 
**SiteIds** | Pointer to **[]int64** | Specifies the Sharepoint site ids across which the indexed items needs to be searched. | [optional] 
**TeamsIds** | Pointer to **[]int64** | Specifies the Teams ids across which the indexed items needs to be searched. | [optional] 
**UserIds** | Pointer to **[]int64** | Specifies the user ids across which the indexed items needs to be searched. | [optional] 

## Methods

### NewO365SearchRequestParams

`func NewO365SearchRequestParams() *O365SearchRequestParams`

NewO365SearchRequestParams instantiates a new O365SearchRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO365SearchRequestParamsWithDefaults

`func NewO365SearchRequestParamsWithDefaults() *O365SearchRequestParams`

NewO365SearchRequestParamsWithDefaults instantiates a new O365SearchRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainIds

`func (o *O365SearchRequestParams) GetDomainIds() []int64`

GetDomainIds returns the DomainIds field if non-nil, zero value otherwise.

### GetDomainIdsOk

`func (o *O365SearchRequestParams) GetDomainIdsOk() (*[]int64, bool)`

GetDomainIdsOk returns a tuple with the DomainIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainIds

`func (o *O365SearchRequestParams) SetDomainIds(v []int64)`

SetDomainIds sets DomainIds field to given value.

### HasDomainIds

`func (o *O365SearchRequestParams) HasDomainIds() bool`

HasDomainIds returns a boolean if a field has been set.

### SetDomainIdsNil

`func (o *O365SearchRequestParams) SetDomainIdsNil(b bool)`

 SetDomainIdsNil sets the value for DomainIds to be an explicit nil

### UnsetDomainIds
`func (o *O365SearchRequestParams) UnsetDomainIds()`

UnsetDomainIds ensures that no value is present for DomainIds, not even an explicit nil
### GetGroupIds

`func (o *O365SearchRequestParams) GetGroupIds() []int64`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *O365SearchRequestParams) GetGroupIdsOk() (*[]int64, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *O365SearchRequestParams) SetGroupIds(v []int64)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *O365SearchRequestParams) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### SetGroupIdsNil

`func (o *O365SearchRequestParams) SetGroupIdsNil(b bool)`

 SetGroupIdsNil sets the value for GroupIds to be an explicit nil

### UnsetGroupIds
`func (o *O365SearchRequestParams) UnsetGroupIds()`

UnsetGroupIds ensures that no value is present for GroupIds, not even an explicit nil
### GetSiteIds

`func (o *O365SearchRequestParams) GetSiteIds() []int64`

GetSiteIds returns the SiteIds field if non-nil, zero value otherwise.

### GetSiteIdsOk

`func (o *O365SearchRequestParams) GetSiteIdsOk() (*[]int64, bool)`

GetSiteIdsOk returns a tuple with the SiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteIds

`func (o *O365SearchRequestParams) SetSiteIds(v []int64)`

SetSiteIds sets SiteIds field to given value.

### HasSiteIds

`func (o *O365SearchRequestParams) HasSiteIds() bool`

HasSiteIds returns a boolean if a field has been set.

### SetSiteIdsNil

`func (o *O365SearchRequestParams) SetSiteIdsNil(b bool)`

 SetSiteIdsNil sets the value for SiteIds to be an explicit nil

### UnsetSiteIds
`func (o *O365SearchRequestParams) UnsetSiteIds()`

UnsetSiteIds ensures that no value is present for SiteIds, not even an explicit nil
### GetTeamsIds

`func (o *O365SearchRequestParams) GetTeamsIds() []int64`

GetTeamsIds returns the TeamsIds field if non-nil, zero value otherwise.

### GetTeamsIdsOk

`func (o *O365SearchRequestParams) GetTeamsIdsOk() (*[]int64, bool)`

GetTeamsIdsOk returns a tuple with the TeamsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsIds

`func (o *O365SearchRequestParams) SetTeamsIds(v []int64)`

SetTeamsIds sets TeamsIds field to given value.

### HasTeamsIds

`func (o *O365SearchRequestParams) HasTeamsIds() bool`

HasTeamsIds returns a boolean if a field has been set.

### SetTeamsIdsNil

`func (o *O365SearchRequestParams) SetTeamsIdsNil(b bool)`

 SetTeamsIdsNil sets the value for TeamsIds to be an explicit nil

### UnsetTeamsIds
`func (o *O365SearchRequestParams) UnsetTeamsIds()`

UnsetTeamsIds ensures that no value is present for TeamsIds, not even an explicit nil
### GetUserIds

`func (o *O365SearchRequestParams) GetUserIds() []int64`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *O365SearchRequestParams) GetUserIdsOk() (*[]int64, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *O365SearchRequestParams) SetUserIds(v []int64)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *O365SearchRequestParams) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### SetUserIdsNil

`func (o *O365SearchRequestParams) SetUserIdsNil(b bool)`

 SetUserIdsNil sets the value for UserIds to be an explicit nil

### UnsetUserIds
`func (o *O365SearchRequestParams) UnsetUserIds()`

UnsetUserIds ensures that no value is present for UserIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


