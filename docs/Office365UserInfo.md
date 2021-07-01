# Office365UserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **NullableString** | Specifies the city in which the Office365 user is located. | [optional] 
**Country** | Pointer to **NullableString** | Specifies the country/region in which the Office365 user is located. | [optional] 
**Department** | Pointer to **NullableString** | Specifies the department within the enterprise of the Office365 user. | [optional] 
**Designation** | Pointer to **NullableString** | Specifies the designation of the Office365 user. | [optional] 
**GraphUuid** | Pointer to **NullableString** | Specifies the MS Graph UUID for the given user entity. | [optional] 
**MailboxSize** | Pointer to **NullableInt64** | Specifies the size of the Outlook Mailbox associated with this Office365 entity. | [optional] 
**OneDriveId** | Pointer to **NullableString** | Specifies the Id of the OneDrive associated with the this Office 365 entity. | [optional] 
**OneDriveSize** | Pointer to **NullableInt64** | Specifies the size of the OneDrive associated with this Office365 entity. | [optional] 

## Methods

### NewOffice365UserInfo

`func NewOffice365UserInfo() *Office365UserInfo`

NewOffice365UserInfo instantiates a new Office365UserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365UserInfoWithDefaults

`func NewOffice365UserInfoWithDefaults() *Office365UserInfo`

NewOffice365UserInfoWithDefaults instantiates a new Office365UserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *Office365UserInfo) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Office365UserInfo) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Office365UserInfo) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Office365UserInfo) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *Office365UserInfo) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *Office365UserInfo) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountry

`func (o *Office365UserInfo) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Office365UserInfo) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Office365UserInfo) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Office365UserInfo) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *Office365UserInfo) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *Office365UserInfo) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetDepartment

`func (o *Office365UserInfo) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *Office365UserInfo) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *Office365UserInfo) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *Office365UserInfo) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *Office365UserInfo) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *Office365UserInfo) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetDesignation

`func (o *Office365UserInfo) GetDesignation() string`

GetDesignation returns the Designation field if non-nil, zero value otherwise.

### GetDesignationOk

`func (o *Office365UserInfo) GetDesignationOk() (*string, bool)`

GetDesignationOk returns a tuple with the Designation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignation

`func (o *Office365UserInfo) SetDesignation(v string)`

SetDesignation sets Designation field to given value.

### HasDesignation

`func (o *Office365UserInfo) HasDesignation() bool`

HasDesignation returns a boolean if a field has been set.

### SetDesignationNil

`func (o *Office365UserInfo) SetDesignationNil(b bool)`

 SetDesignationNil sets the value for Designation to be an explicit nil

### UnsetDesignation
`func (o *Office365UserInfo) UnsetDesignation()`

UnsetDesignation ensures that no value is present for Designation, not even an explicit nil
### GetGraphUuid

`func (o *Office365UserInfo) GetGraphUuid() string`

GetGraphUuid returns the GraphUuid field if non-nil, zero value otherwise.

### GetGraphUuidOk

`func (o *Office365UserInfo) GetGraphUuidOk() (*string, bool)`

GetGraphUuidOk returns a tuple with the GraphUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphUuid

`func (o *Office365UserInfo) SetGraphUuid(v string)`

SetGraphUuid sets GraphUuid field to given value.

### HasGraphUuid

`func (o *Office365UserInfo) HasGraphUuid() bool`

HasGraphUuid returns a boolean if a field has been set.

### SetGraphUuidNil

`func (o *Office365UserInfo) SetGraphUuidNil(b bool)`

 SetGraphUuidNil sets the value for GraphUuid to be an explicit nil

### UnsetGraphUuid
`func (o *Office365UserInfo) UnsetGraphUuid()`

UnsetGraphUuid ensures that no value is present for GraphUuid, not even an explicit nil
### GetMailboxSize

`func (o *Office365UserInfo) GetMailboxSize() int64`

GetMailboxSize returns the MailboxSize field if non-nil, zero value otherwise.

### GetMailboxSizeOk

`func (o *Office365UserInfo) GetMailboxSizeOk() (*int64, bool)`

GetMailboxSizeOk returns a tuple with the MailboxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxSize

`func (o *Office365UserInfo) SetMailboxSize(v int64)`

SetMailboxSize sets MailboxSize field to given value.

### HasMailboxSize

`func (o *Office365UserInfo) HasMailboxSize() bool`

HasMailboxSize returns a boolean if a field has been set.

### SetMailboxSizeNil

`func (o *Office365UserInfo) SetMailboxSizeNil(b bool)`

 SetMailboxSizeNil sets the value for MailboxSize to be an explicit nil

### UnsetMailboxSize
`func (o *Office365UserInfo) UnsetMailboxSize()`

UnsetMailboxSize ensures that no value is present for MailboxSize, not even an explicit nil
### GetOneDriveId

`func (o *Office365UserInfo) GetOneDriveId() string`

GetOneDriveId returns the OneDriveId field if non-nil, zero value otherwise.

### GetOneDriveIdOk

`func (o *Office365UserInfo) GetOneDriveIdOk() (*string, bool)`

GetOneDriveIdOk returns a tuple with the OneDriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneDriveId

`func (o *Office365UserInfo) SetOneDriveId(v string)`

SetOneDriveId sets OneDriveId field to given value.

### HasOneDriveId

`func (o *Office365UserInfo) HasOneDriveId() bool`

HasOneDriveId returns a boolean if a field has been set.

### SetOneDriveIdNil

`func (o *Office365UserInfo) SetOneDriveIdNil(b bool)`

 SetOneDriveIdNil sets the value for OneDriveId to be an explicit nil

### UnsetOneDriveId
`func (o *Office365UserInfo) UnsetOneDriveId()`

UnsetOneDriveId ensures that no value is present for OneDriveId, not even an explicit nil
### GetOneDriveSize

`func (o *Office365UserInfo) GetOneDriveSize() int64`

GetOneDriveSize returns the OneDriveSize field if non-nil, zero value otherwise.

### GetOneDriveSizeOk

`func (o *Office365UserInfo) GetOneDriveSizeOk() (*int64, bool)`

GetOneDriveSizeOk returns a tuple with the OneDriveSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneDriveSize

`func (o *Office365UserInfo) SetOneDriveSize(v int64)`

SetOneDriveSize sets OneDriveSize field to given value.

### HasOneDriveSize

`func (o *Office365UserInfo) HasOneDriveSize() bool`

HasOneDriveSize returns a boolean if a field has been set.

### SetOneDriveSizeNil

`func (o *Office365UserInfo) SetOneDriveSizeNil(b bool)`

 SetOneDriveSizeNil sets the value for OneDriveSize to be an explicit nil

### UnsetOneDriveSize
`func (o *Office365UserInfo) UnsetOneDriveSize()`

UnsetOneDriveSize ensures that no value is present for OneDriveSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


