# Office365SourceRegistrationParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Office365AppCredentialsList** | Pointer to [**[]Office365AppCredentials**](Office365AppCredentials.md) | Specifies a list of office365 azure application credentials needed to authenticate &amp; authorize users for Office 365. | [optional] 
**UseOAuthForExchangeOnline** | Pointer to **NullableBool** | Specifies whether OAuth should be used for authentication in case of Exchange Online. | [optional] 
**ProxyHostSourceIdList** | Pointer to **[]int64** | Specifies the list of the protection source id of the windows physical host which will be used during the protection and recovery of the sites that belong to a office365 domain. | [optional] 
**Office365Region** | Pointer to **NullableString** | Specifies the region where Office 365 Exchange environment is. | [optional] 
**O365ObjectsDiscoveryParams** | Pointer to [**ObjectsDiscoveryParams**](ObjectsDiscoveryParams.md) |  | [optional] 

## Methods

### NewOffice365SourceRegistrationParamsAllOf

`func NewOffice365SourceRegistrationParamsAllOf() *Office365SourceRegistrationParamsAllOf`

NewOffice365SourceRegistrationParamsAllOf instantiates a new Office365SourceRegistrationParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365SourceRegistrationParamsAllOfWithDefaults

`func NewOffice365SourceRegistrationParamsAllOfWithDefaults() *Office365SourceRegistrationParamsAllOf`

NewOffice365SourceRegistrationParamsAllOfWithDefaults instantiates a new Office365SourceRegistrationParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffice365AppCredentialsList

`func (o *Office365SourceRegistrationParamsAllOf) GetOffice365AppCredentialsList() []Office365AppCredentials`

GetOffice365AppCredentialsList returns the Office365AppCredentialsList field if non-nil, zero value otherwise.

### GetOffice365AppCredentialsListOk

`func (o *Office365SourceRegistrationParamsAllOf) GetOffice365AppCredentialsListOk() (*[]Office365AppCredentials, bool)`

GetOffice365AppCredentialsListOk returns a tuple with the Office365AppCredentialsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365AppCredentialsList

`func (o *Office365SourceRegistrationParamsAllOf) SetOffice365AppCredentialsList(v []Office365AppCredentials)`

SetOffice365AppCredentialsList sets Office365AppCredentialsList field to given value.

### HasOffice365AppCredentialsList

`func (o *Office365SourceRegistrationParamsAllOf) HasOffice365AppCredentialsList() bool`

HasOffice365AppCredentialsList returns a boolean if a field has been set.

### GetUseOAuthForExchangeOnline

`func (o *Office365SourceRegistrationParamsAllOf) GetUseOAuthForExchangeOnline() bool`

GetUseOAuthForExchangeOnline returns the UseOAuthForExchangeOnline field if non-nil, zero value otherwise.

### GetUseOAuthForExchangeOnlineOk

`func (o *Office365SourceRegistrationParamsAllOf) GetUseOAuthForExchangeOnlineOk() (*bool, bool)`

GetUseOAuthForExchangeOnlineOk returns a tuple with the UseOAuthForExchangeOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOAuthForExchangeOnline

`func (o *Office365SourceRegistrationParamsAllOf) SetUseOAuthForExchangeOnline(v bool)`

SetUseOAuthForExchangeOnline sets UseOAuthForExchangeOnline field to given value.

### HasUseOAuthForExchangeOnline

`func (o *Office365SourceRegistrationParamsAllOf) HasUseOAuthForExchangeOnline() bool`

HasUseOAuthForExchangeOnline returns a boolean if a field has been set.

### SetUseOAuthForExchangeOnlineNil

`func (o *Office365SourceRegistrationParamsAllOf) SetUseOAuthForExchangeOnlineNil(b bool)`

 SetUseOAuthForExchangeOnlineNil sets the value for UseOAuthForExchangeOnline to be an explicit nil

### UnsetUseOAuthForExchangeOnline
`func (o *Office365SourceRegistrationParamsAllOf) UnsetUseOAuthForExchangeOnline()`

UnsetUseOAuthForExchangeOnline ensures that no value is present for UseOAuthForExchangeOnline, not even an explicit nil
### GetProxyHostSourceIdList

`func (o *Office365SourceRegistrationParamsAllOf) GetProxyHostSourceIdList() []int64`

GetProxyHostSourceIdList returns the ProxyHostSourceIdList field if non-nil, zero value otherwise.

### GetProxyHostSourceIdListOk

`func (o *Office365SourceRegistrationParamsAllOf) GetProxyHostSourceIdListOk() (*[]int64, bool)`

GetProxyHostSourceIdListOk returns a tuple with the ProxyHostSourceIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyHostSourceIdList

`func (o *Office365SourceRegistrationParamsAllOf) SetProxyHostSourceIdList(v []int64)`

SetProxyHostSourceIdList sets ProxyHostSourceIdList field to given value.

### HasProxyHostSourceIdList

`func (o *Office365SourceRegistrationParamsAllOf) HasProxyHostSourceIdList() bool`

HasProxyHostSourceIdList returns a boolean if a field has been set.

### SetProxyHostSourceIdListNil

`func (o *Office365SourceRegistrationParamsAllOf) SetProxyHostSourceIdListNil(b bool)`

 SetProxyHostSourceIdListNil sets the value for ProxyHostSourceIdList to be an explicit nil

### UnsetProxyHostSourceIdList
`func (o *Office365SourceRegistrationParamsAllOf) UnsetProxyHostSourceIdList()`

UnsetProxyHostSourceIdList ensures that no value is present for ProxyHostSourceIdList, not even an explicit nil
### GetOffice365Region

`func (o *Office365SourceRegistrationParamsAllOf) GetOffice365Region() string`

GetOffice365Region returns the Office365Region field if non-nil, zero value otherwise.

### GetOffice365RegionOk

`func (o *Office365SourceRegistrationParamsAllOf) GetOffice365RegionOk() (*string, bool)`

GetOffice365RegionOk returns a tuple with the Office365Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Region

`func (o *Office365SourceRegistrationParamsAllOf) SetOffice365Region(v string)`

SetOffice365Region sets Office365Region field to given value.

### HasOffice365Region

`func (o *Office365SourceRegistrationParamsAllOf) HasOffice365Region() bool`

HasOffice365Region returns a boolean if a field has been set.

### SetOffice365RegionNil

`func (o *Office365SourceRegistrationParamsAllOf) SetOffice365RegionNil(b bool)`

 SetOffice365RegionNil sets the value for Office365Region to be an explicit nil

### UnsetOffice365Region
`func (o *Office365SourceRegistrationParamsAllOf) UnsetOffice365Region()`

UnsetOffice365Region ensures that no value is present for Office365Region, not even an explicit nil
### GetO365ObjectsDiscoveryParams

`func (o *Office365SourceRegistrationParamsAllOf) GetO365ObjectsDiscoveryParams() ObjectsDiscoveryParams`

GetO365ObjectsDiscoveryParams returns the O365ObjectsDiscoveryParams field if non-nil, zero value otherwise.

### GetO365ObjectsDiscoveryParamsOk

`func (o *Office365SourceRegistrationParamsAllOf) GetO365ObjectsDiscoveryParamsOk() (*ObjectsDiscoveryParams, bool)`

GetO365ObjectsDiscoveryParamsOk returns a tuple with the O365ObjectsDiscoveryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO365ObjectsDiscoveryParams

`func (o *Office365SourceRegistrationParamsAllOf) SetO365ObjectsDiscoveryParams(v ObjectsDiscoveryParams)`

SetO365ObjectsDiscoveryParams sets O365ObjectsDiscoveryParams field to given value.

### HasO365ObjectsDiscoveryParams

`func (o *Office365SourceRegistrationParamsAllOf) HasO365ObjectsDiscoveryParams() bool`

HasO365ObjectsDiscoveryParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


