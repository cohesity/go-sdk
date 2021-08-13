# ExternallyTriggeredJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientType** | Pointer to **NullableString** | Specifies the client type of the externally triggered backup job. | [optional] 
**SbtParams** | Pointer to [**ExternallyTriggeredSbtParams**](ExternallyTriggeredSbtParams.md) | Specifies the SBT parameters for the externally triggered backup job. | [optional] 
**Tags** | Pointer to **[]string** | Specifies all the tags of the externally triggered job. | [optional] 

## Methods

### NewExternallyTriggeredJobParams

`func NewExternallyTriggeredJobParams() *ExternallyTriggeredJobParams`

NewExternallyTriggeredJobParams instantiates a new ExternallyTriggeredJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternallyTriggeredJobParamsWithDefaults

`func NewExternallyTriggeredJobParamsWithDefaults() *ExternallyTriggeredJobParams`

NewExternallyTriggeredJobParamsWithDefaults instantiates a new ExternallyTriggeredJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientType

`func (o *ExternallyTriggeredJobParams) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *ExternallyTriggeredJobParams) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *ExternallyTriggeredJobParams) SetClientType(v string)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *ExternallyTriggeredJobParams) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### SetClientTypeNil

`func (o *ExternallyTriggeredJobParams) SetClientTypeNil(b bool)`

 SetClientTypeNil sets the value for ClientType to be an explicit nil

### UnsetClientType
`func (o *ExternallyTriggeredJobParams) UnsetClientType()`

UnsetClientType ensures that no value is present for ClientType, not even an explicit nil
### GetSbtParams

`func (o *ExternallyTriggeredJobParams) GetSbtParams() ExternallyTriggeredSbtParams`

GetSbtParams returns the SbtParams field if non-nil, zero value otherwise.

### GetSbtParamsOk

`func (o *ExternallyTriggeredJobParams) GetSbtParamsOk() (*ExternallyTriggeredSbtParams, bool)`

GetSbtParamsOk returns a tuple with the SbtParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbtParams

`func (o *ExternallyTriggeredJobParams) SetSbtParams(v ExternallyTriggeredSbtParams)`

SetSbtParams sets SbtParams field to given value.

### HasSbtParams

`func (o *ExternallyTriggeredJobParams) HasSbtParams() bool`

HasSbtParams returns a boolean if a field has been set.

### GetTags

`func (o *ExternallyTriggeredJobParams) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExternallyTriggeredJobParams) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExternallyTriggeredJobParams) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExternallyTriggeredJobParams) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


