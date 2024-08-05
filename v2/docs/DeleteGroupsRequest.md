# DeleteGroupsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sids** | **[]string** | Specifies a list of group sids to delete. | 

## Methods

### NewDeleteGroupsRequest

`func NewDeleteGroupsRequest(sids []string, ) *DeleteGroupsRequest`

NewDeleteGroupsRequest instantiates a new DeleteGroupsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteGroupsRequestWithDefaults

`func NewDeleteGroupsRequestWithDefaults() *DeleteGroupsRequest`

NewDeleteGroupsRequestWithDefaults instantiates a new DeleteGroupsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSids

`func (o *DeleteGroupsRequest) GetSids() []string`

GetSids returns the Sids field if non-nil, zero value otherwise.

### GetSidsOk

`func (o *DeleteGroupsRequest) GetSidsOk() (*[]string, bool)`

GetSidsOk returns a tuple with the Sids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSids

`func (o *DeleteGroupsRequest) SetSids(v []string)`

SetSids sets Sids field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


