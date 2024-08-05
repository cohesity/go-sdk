# ObjectsDiscoveryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscoverableObjectTypeList** | Pointer to **[]string** | Specifies the list of object types that will be discovered as part of source registration or refresh. | [optional] 
**SitesDiscoveryParams** | Pointer to [**SitesDiscoveryParams**](SitesDiscoveryParams.md) |  | [optional] 
**TeamsAdditionalParams** | Pointer to [**TeamsAdditionalParams**](TeamsAdditionalParams.md) |  | [optional] 
**UsersDiscoveryParams** | Pointer to [**UsersDiscoveryParams**](UsersDiscoveryParams.md) |  | [optional] 

## Methods

### NewObjectsDiscoveryParams

`func NewObjectsDiscoveryParams() *ObjectsDiscoveryParams`

NewObjectsDiscoveryParams instantiates a new ObjectsDiscoveryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectsDiscoveryParamsWithDefaults

`func NewObjectsDiscoveryParamsWithDefaults() *ObjectsDiscoveryParams`

NewObjectsDiscoveryParamsWithDefaults instantiates a new ObjectsDiscoveryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscoverableObjectTypeList

`func (o *ObjectsDiscoveryParams) GetDiscoverableObjectTypeList() []string`

GetDiscoverableObjectTypeList returns the DiscoverableObjectTypeList field if non-nil, zero value otherwise.

### GetDiscoverableObjectTypeListOk

`func (o *ObjectsDiscoveryParams) GetDiscoverableObjectTypeListOk() (*[]string, bool)`

GetDiscoverableObjectTypeListOk returns a tuple with the DiscoverableObjectTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverableObjectTypeList

`func (o *ObjectsDiscoveryParams) SetDiscoverableObjectTypeList(v []string)`

SetDiscoverableObjectTypeList sets DiscoverableObjectTypeList field to given value.

### HasDiscoverableObjectTypeList

`func (o *ObjectsDiscoveryParams) HasDiscoverableObjectTypeList() bool`

HasDiscoverableObjectTypeList returns a boolean if a field has been set.

### SetDiscoverableObjectTypeListNil

`func (o *ObjectsDiscoveryParams) SetDiscoverableObjectTypeListNil(b bool)`

 SetDiscoverableObjectTypeListNil sets the value for DiscoverableObjectTypeList to be an explicit nil

### UnsetDiscoverableObjectTypeList
`func (o *ObjectsDiscoveryParams) UnsetDiscoverableObjectTypeList()`

UnsetDiscoverableObjectTypeList ensures that no value is present for DiscoverableObjectTypeList, not even an explicit nil
### GetSitesDiscoveryParams

`func (o *ObjectsDiscoveryParams) GetSitesDiscoveryParams() SitesDiscoveryParams`

GetSitesDiscoveryParams returns the SitesDiscoveryParams field if non-nil, zero value otherwise.

### GetSitesDiscoveryParamsOk

`func (o *ObjectsDiscoveryParams) GetSitesDiscoveryParamsOk() (*SitesDiscoveryParams, bool)`

GetSitesDiscoveryParamsOk returns a tuple with the SitesDiscoveryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitesDiscoveryParams

`func (o *ObjectsDiscoveryParams) SetSitesDiscoveryParams(v SitesDiscoveryParams)`

SetSitesDiscoveryParams sets SitesDiscoveryParams field to given value.

### HasSitesDiscoveryParams

`func (o *ObjectsDiscoveryParams) HasSitesDiscoveryParams() bool`

HasSitesDiscoveryParams returns a boolean if a field has been set.

### GetTeamsAdditionalParams

`func (o *ObjectsDiscoveryParams) GetTeamsAdditionalParams() TeamsAdditionalParams`

GetTeamsAdditionalParams returns the TeamsAdditionalParams field if non-nil, zero value otherwise.

### GetTeamsAdditionalParamsOk

`func (o *ObjectsDiscoveryParams) GetTeamsAdditionalParamsOk() (*TeamsAdditionalParams, bool)`

GetTeamsAdditionalParamsOk returns a tuple with the TeamsAdditionalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsAdditionalParams

`func (o *ObjectsDiscoveryParams) SetTeamsAdditionalParams(v TeamsAdditionalParams)`

SetTeamsAdditionalParams sets TeamsAdditionalParams field to given value.

### HasTeamsAdditionalParams

`func (o *ObjectsDiscoveryParams) HasTeamsAdditionalParams() bool`

HasTeamsAdditionalParams returns a boolean if a field has been set.

### GetUsersDiscoveryParams

`func (o *ObjectsDiscoveryParams) GetUsersDiscoveryParams() UsersDiscoveryParams`

GetUsersDiscoveryParams returns the UsersDiscoveryParams field if non-nil, zero value otherwise.

### GetUsersDiscoveryParamsOk

`func (o *ObjectsDiscoveryParams) GetUsersDiscoveryParamsOk() (*UsersDiscoveryParams, bool)`

GetUsersDiscoveryParamsOk returns a tuple with the UsersDiscoveryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersDiscoveryParams

`func (o *ObjectsDiscoveryParams) SetUsersDiscoveryParams(v UsersDiscoveryParams)`

SetUsersDiscoveryParams sets UsersDiscoveryParams field to given value.

### HasUsersDiscoveryParams

`func (o *ObjectsDiscoveryParams) HasUsersDiscoveryParams() bool`

HasUsersDiscoveryParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


