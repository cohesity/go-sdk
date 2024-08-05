# OAuthAudience

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudienceId** | **NullableString** | Specifies the ID of this audience. This must match the &#39;aud&#39; field in the token at login time. | 
**ClientIds** | **[]string** | Specifies the list of client IDs which should be allowed to log in via this audience. The &#39;appid&#39; in the token must match one of the values specified here. | 

## Methods

### NewOAuthAudience

`func NewOAuthAudience(audienceId NullableString, clientIds []string, ) *OAuthAudience`

NewOAuthAudience instantiates a new OAuthAudience object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthAudienceWithDefaults

`func NewOAuthAudienceWithDefaults() *OAuthAudience`

NewOAuthAudienceWithDefaults instantiates a new OAuthAudience object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudienceId

`func (o *OAuthAudience) GetAudienceId() string`

GetAudienceId returns the AudienceId field if non-nil, zero value otherwise.

### GetAudienceIdOk

`func (o *OAuthAudience) GetAudienceIdOk() (*string, bool)`

GetAudienceIdOk returns a tuple with the AudienceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceId

`func (o *OAuthAudience) SetAudienceId(v string)`

SetAudienceId sets AudienceId field to given value.


### SetAudienceIdNil

`func (o *OAuthAudience) SetAudienceIdNil(b bool)`

 SetAudienceIdNil sets the value for AudienceId to be an explicit nil

### UnsetAudienceId
`func (o *OAuthAudience) UnsetAudienceId()`

UnsetAudienceId ensures that no value is present for AudienceId, not even an explicit nil
### GetClientIds

`func (o *OAuthAudience) GetClientIds() []string`

GetClientIds returns the ClientIds field if non-nil, zero value otherwise.

### GetClientIdsOk

`func (o *OAuthAudience) GetClientIdsOk() (*[]string, bool)`

GetClientIdsOk returns a tuple with the ClientIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIds

`func (o *OAuthAudience) SetClientIds(v []string)`

SetClientIds sets ClientIds field to given value.


### SetClientIdsNil

`func (o *OAuthAudience) SetClientIdsNil(b bool)`

 SetClientIdsNil sets the value for ClientIds to be an explicit nil

### UnsetClientIds
`func (o *OAuthAudience) UnsetClientIds()`

UnsetClientIds ensures that no value is present for ClientIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


