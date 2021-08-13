# O365SearchSharepointRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainIds** | Pointer to **[]int64** | Specifies the domain Ids in which Sharepoint sites are registered. | [optional] 
**SiteIds** | Pointer to **[]int64** | Specifies the Sharepoint site Ids across which the document library files/folders needs to be searched. | [optional] 

## Methods

### NewO365SearchSharepointRequestParams

`func NewO365SearchSharepointRequestParams() *O365SearchSharepointRequestParams`

NewO365SearchSharepointRequestParams instantiates a new O365SearchSharepointRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO365SearchSharepointRequestParamsWithDefaults

`func NewO365SearchSharepointRequestParamsWithDefaults() *O365SearchSharepointRequestParams`

NewO365SearchSharepointRequestParamsWithDefaults instantiates a new O365SearchSharepointRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainIds

`func (o *O365SearchSharepointRequestParams) GetDomainIds() []int64`

GetDomainIds returns the DomainIds field if non-nil, zero value otherwise.

### GetDomainIdsOk

`func (o *O365SearchSharepointRequestParams) GetDomainIdsOk() (*[]int64, bool)`

GetDomainIdsOk returns a tuple with the DomainIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainIds

`func (o *O365SearchSharepointRequestParams) SetDomainIds(v []int64)`

SetDomainIds sets DomainIds field to given value.

### HasDomainIds

`func (o *O365SearchSharepointRequestParams) HasDomainIds() bool`

HasDomainIds returns a boolean if a field has been set.

### SetDomainIdsNil

`func (o *O365SearchSharepointRequestParams) SetDomainIdsNil(b bool)`

 SetDomainIdsNil sets the value for DomainIds to be an explicit nil

### UnsetDomainIds
`func (o *O365SearchSharepointRequestParams) UnsetDomainIds()`

UnsetDomainIds ensures that no value is present for DomainIds, not even an explicit nil
### GetSiteIds

`func (o *O365SearchSharepointRequestParams) GetSiteIds() []int64`

GetSiteIds returns the SiteIds field if non-nil, zero value otherwise.

### GetSiteIdsOk

`func (o *O365SearchSharepointRequestParams) GetSiteIdsOk() (*[]int64, bool)`

GetSiteIdsOk returns a tuple with the SiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteIds

`func (o *O365SearchSharepointRequestParams) SetSiteIds(v []int64)`

SetSiteIds sets SiteIds field to given value.

### HasSiteIds

`func (o *O365SearchSharepointRequestParams) HasSiteIds() bool`

HasSiteIds returns a boolean if a field has been set.

### SetSiteIdsNil

`func (o *O365SearchSharepointRequestParams) SetSiteIdsNil(b bool)`

 SetSiteIdsNil sets the value for SiteIds to be an explicit nil

### UnsetSiteIds
`func (o *O365SearchSharepointRequestParams) UnsetSiteIds()`

UnsetSiteIds ensures that no value is present for SiteIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


