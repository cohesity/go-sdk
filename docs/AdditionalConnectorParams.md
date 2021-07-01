# AdditionalConnectorParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**O365Region** | Pointer to [**O365RegionProto**](O365RegionProto.md) |  | [optional] 
**UseOutlookEwsOauth** | Pointer to **NullableBool** | Whether OAuth should be used for authentication with EWS API (outlook backup), applicable only for Exchange Online. | [optional] 

## Methods

### NewAdditionalConnectorParams

`func NewAdditionalConnectorParams() *AdditionalConnectorParams`

NewAdditionalConnectorParams instantiates a new AdditionalConnectorParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalConnectorParamsWithDefaults

`func NewAdditionalConnectorParamsWithDefaults() *AdditionalConnectorParams`

NewAdditionalConnectorParamsWithDefaults instantiates a new AdditionalConnectorParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetO365Region

`func (o *AdditionalConnectorParams) GetO365Region() O365RegionProto`

GetO365Region returns the O365Region field if non-nil, zero value otherwise.

### GetO365RegionOk

`func (o *AdditionalConnectorParams) GetO365RegionOk() (*O365RegionProto, bool)`

GetO365RegionOk returns a tuple with the O365Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO365Region

`func (o *AdditionalConnectorParams) SetO365Region(v O365RegionProto)`

SetO365Region sets O365Region field to given value.

### HasO365Region

`func (o *AdditionalConnectorParams) HasO365Region() bool`

HasO365Region returns a boolean if a field has been set.

### GetUseOutlookEwsOauth

`func (o *AdditionalConnectorParams) GetUseOutlookEwsOauth() bool`

GetUseOutlookEwsOauth returns the UseOutlookEwsOauth field if non-nil, zero value otherwise.

### GetUseOutlookEwsOauthOk

`func (o *AdditionalConnectorParams) GetUseOutlookEwsOauthOk() (*bool, bool)`

GetUseOutlookEwsOauthOk returns a tuple with the UseOutlookEwsOauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOutlookEwsOauth

`func (o *AdditionalConnectorParams) SetUseOutlookEwsOauth(v bool)`

SetUseOutlookEwsOauth sets UseOutlookEwsOauth field to given value.

### HasUseOutlookEwsOauth

`func (o *AdditionalConnectorParams) HasUseOutlookEwsOauth() bool`

HasUseOutlookEwsOauth returns a boolean if a field has been set.

### SetUseOutlookEwsOauthNil

`func (o *AdditionalConnectorParams) SetUseOutlookEwsOauthNil(b bool)`

 SetUseOutlookEwsOauthNil sets the value for UseOutlookEwsOauth to be an explicit nil

### UnsetUseOutlookEwsOauth
`func (o *AdditionalConnectorParams) UnsetUseOutlookEwsOauth()`

UnsetUseOutlookEwsOauth ensures that no value is present for UseOutlookEwsOauth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


