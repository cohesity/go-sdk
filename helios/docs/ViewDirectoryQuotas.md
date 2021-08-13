# ViewDirectoryQuotas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectoryQuotas** | Pointer to [**[]ViewDirectoryQuota**](ViewDirectoryQuota.md) | Specifies the list of View directory quotas. | [optional] 
**Cookie** | Pointer to **NullableInt64** | Specifies the pagination cookie. | [optional] 

## Methods

### NewViewDirectoryQuotas

`func NewViewDirectoryQuotas() *ViewDirectoryQuotas`

NewViewDirectoryQuotas instantiates a new ViewDirectoryQuotas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewDirectoryQuotasWithDefaults

`func NewViewDirectoryQuotasWithDefaults() *ViewDirectoryQuotas`

NewViewDirectoryQuotasWithDefaults instantiates a new ViewDirectoryQuotas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectoryQuotas

`func (o *ViewDirectoryQuotas) GetDirectoryQuotas() []ViewDirectoryQuota`

GetDirectoryQuotas returns the DirectoryQuotas field if non-nil, zero value otherwise.

### GetDirectoryQuotasOk

`func (o *ViewDirectoryQuotas) GetDirectoryQuotasOk() (*[]ViewDirectoryQuota, bool)`

GetDirectoryQuotasOk returns a tuple with the DirectoryQuotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryQuotas

`func (o *ViewDirectoryQuotas) SetDirectoryQuotas(v []ViewDirectoryQuota)`

SetDirectoryQuotas sets DirectoryQuotas field to given value.

### HasDirectoryQuotas

`func (o *ViewDirectoryQuotas) HasDirectoryQuotas() bool`

HasDirectoryQuotas returns a boolean if a field has been set.

### SetDirectoryQuotasNil

`func (o *ViewDirectoryQuotas) SetDirectoryQuotasNil(b bool)`

 SetDirectoryQuotasNil sets the value for DirectoryQuotas to be an explicit nil

### UnsetDirectoryQuotas
`func (o *ViewDirectoryQuotas) UnsetDirectoryQuotas()`

UnsetDirectoryQuotas ensures that no value is present for DirectoryQuotas, not even an explicit nil
### GetCookie

`func (o *ViewDirectoryQuotas) GetCookie() int64`

GetCookie returns the Cookie field if non-nil, zero value otherwise.

### GetCookieOk

`func (o *ViewDirectoryQuotas) GetCookieOk() (*int64, bool)`

GetCookieOk returns a tuple with the Cookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookie

`func (o *ViewDirectoryQuotas) SetCookie(v int64)`

SetCookie sets Cookie field to given value.

### HasCookie

`func (o *ViewDirectoryQuotas) HasCookie() bool`

HasCookie returns a boolean if a field has been set.

### SetCookieNil

`func (o *ViewDirectoryQuotas) SetCookieNil(b bool)`

 SetCookieNil sets the value for Cookie to be an explicit nil

### UnsetCookie
`func (o *ViewDirectoryQuotas) UnsetCookie()`

UnsetCookie ensures that no value is present for Cookie, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


