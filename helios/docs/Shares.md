# Shares

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shares** | Pointer to [**[]Share**](Share.md) | Specifies the list of shares. | [optional] 
**Cookie** | Pointer to **NullableString** | Specifies the pagination cookie. | [optional] 

## Methods

### NewShares

`func NewShares() *Shares`

NewShares instantiates a new Shares object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharesWithDefaults

`func NewSharesWithDefaults() *Shares`

NewSharesWithDefaults instantiates a new Shares object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShares

`func (o *Shares) GetShares() []Share`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *Shares) GetSharesOk() (*[]Share, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *Shares) SetShares(v []Share)`

SetShares sets Shares field to given value.

### HasShares

`func (o *Shares) HasShares() bool`

HasShares returns a boolean if a field has been set.

### SetSharesNil

`func (o *Shares) SetSharesNil(b bool)`

 SetSharesNil sets the value for Shares to be an explicit nil

### UnsetShares
`func (o *Shares) UnsetShares()`

UnsetShares ensures that no value is present for Shares, not even an explicit nil
### GetCookie

`func (o *Shares) GetCookie() string`

GetCookie returns the Cookie field if non-nil, zero value otherwise.

### GetCookieOk

`func (o *Shares) GetCookieOk() (*string, bool)`

GetCookieOk returns a tuple with the Cookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookie

`func (o *Shares) SetCookie(v string)`

SetCookie sets Cookie field to given value.

### HasCookie

`func (o *Shares) HasCookie() bool`

HasCookie returns a boolean if a field has been set.

### SetCookieNil

`func (o *Shares) SetCookieNil(b bool)`

 SetCookieNil sets the value for Cookie to be an explicit nil

### UnsetCookie
`func (o *Shares) UnsetCookie()`

UnsetCookie ensures that no value is present for Cookie, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


