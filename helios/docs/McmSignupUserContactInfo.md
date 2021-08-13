# McmSignupUserContactInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | Specifies the first name of the user. | 
**LastName** | **string** | Specifies the last name of the user. | 
**Email** | **string** | Specifies the user&#39;s email address. | 
**PhoneNumber** | Pointer to **NullableString** | Specifies the user&#39;s phone number. | [optional] 
**CompanyName** | **string** | Specifies the name of the Company to which the user belongs. | 
**CompanySize** | Pointer to **NullableString** | Specifies the size of the company. For example, 1-500. | [optional] 
**IndustryType** | Pointer to **NullableString** | Specifies the industry type of the company. For example IT. | [optional] 

## Methods

### NewMcmSignupUserContactInfo

`func NewMcmSignupUserContactInfo(firstName string, lastName string, email string, companyName string, ) *McmSignupUserContactInfo`

NewMcmSignupUserContactInfo instantiates a new McmSignupUserContactInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmSignupUserContactInfoWithDefaults

`func NewMcmSignupUserContactInfoWithDefaults() *McmSignupUserContactInfo`

NewMcmSignupUserContactInfoWithDefaults instantiates a new McmSignupUserContactInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *McmSignupUserContactInfo) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *McmSignupUserContactInfo) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *McmSignupUserContactInfo) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *McmSignupUserContactInfo) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *McmSignupUserContactInfo) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *McmSignupUserContactInfo) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetEmail

`func (o *McmSignupUserContactInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *McmSignupUserContactInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *McmSignupUserContactInfo) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPhoneNumber

`func (o *McmSignupUserContactInfo) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *McmSignupUserContactInfo) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *McmSignupUserContactInfo) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *McmSignupUserContactInfo) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *McmSignupUserContactInfo) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *McmSignupUserContactInfo) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetCompanyName

`func (o *McmSignupUserContactInfo) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *McmSignupUserContactInfo) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *McmSignupUserContactInfo) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetCompanySize

`func (o *McmSignupUserContactInfo) GetCompanySize() string`

GetCompanySize returns the CompanySize field if non-nil, zero value otherwise.

### GetCompanySizeOk

`func (o *McmSignupUserContactInfo) GetCompanySizeOk() (*string, bool)`

GetCompanySizeOk returns a tuple with the CompanySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanySize

`func (o *McmSignupUserContactInfo) SetCompanySize(v string)`

SetCompanySize sets CompanySize field to given value.

### HasCompanySize

`func (o *McmSignupUserContactInfo) HasCompanySize() bool`

HasCompanySize returns a boolean if a field has been set.

### SetCompanySizeNil

`func (o *McmSignupUserContactInfo) SetCompanySizeNil(b bool)`

 SetCompanySizeNil sets the value for CompanySize to be an explicit nil

### UnsetCompanySize
`func (o *McmSignupUserContactInfo) UnsetCompanySize()`

UnsetCompanySize ensures that no value is present for CompanySize, not even an explicit nil
### GetIndustryType

`func (o *McmSignupUserContactInfo) GetIndustryType() string`

GetIndustryType returns the IndustryType field if non-nil, zero value otherwise.

### GetIndustryTypeOk

`func (o *McmSignupUserContactInfo) GetIndustryTypeOk() (*string, bool)`

GetIndustryTypeOk returns a tuple with the IndustryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryType

`func (o *McmSignupUserContactInfo) SetIndustryType(v string)`

SetIndustryType sets IndustryType field to given value.

### HasIndustryType

`func (o *McmSignupUserContactInfo) HasIndustryType() bool`

HasIndustryType returns a boolean if a field has been set.

### SetIndustryTypeNil

`func (o *McmSignupUserContactInfo) SetIndustryTypeNil(b bool)`

 SetIndustryTypeNil sets the value for IndustryType to be an explicit nil

### UnsetIndustryType
`func (o *McmSignupUserContactInfo) UnsetIndustryType()`

UnsetIndustryType ensures that no value is present for IndustryType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


