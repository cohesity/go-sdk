# SearchEmailRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttendeesAddresses** | Pointer to **[]string** | Filters the calendar items which have specified email addresses as attendees. | [optional] 
**BccRecipientAddresses** | Pointer to **[]string** | Filters the emails which are sent to specified email addresses in BCC. | [optional] 
**CcRecipientAddresses** | Pointer to **[]string** | Filters the emails which are sent to specified email addresses in CC. | [optional] 
**CreatedEndTimeSecs** | Pointer to **NullableInt64** | Specifies the end time in Unix timestamp epoch in seconds where the created time of the email/item is less than specified value. | [optional] 
**CreatedStartTimeSecs** | Pointer to **NullableInt64** | Specifies the start time in Unix timestamp epoch in seconds where the created time of the email/item is more than specified value. | [optional] 
**DueDateEndTimeSecs** | Pointer to **NullableInt64** | Specifies the end time in Unix timestamp epoch in seconds where the last modification time of the email/item is less than specified value. | [optional] 
**DueDateStartTimeSecs** | Pointer to **NullableInt64** | Specifies the start time in Unix timestamp epoch in seconds where the last modification time of the email/item is more than specified value. | [optional] 
**EmailAddress** | Pointer to **NullableString** | Filters the contact items which have specified text in email address. | [optional] 
**EmailSubject** | Pointer to **NullableString** | Filters the emails which have the specified text in its subject. | [optional] 
**FirstName** | Pointer to **NullableString** | Filters the contacts with specified text in first name. | [optional] 
**FolderNames** | Pointer to **[]string** | Filters the emails which are categorized to specified folders. | [optional] 
**HasAttachment** | Pointer to **NullableBool** | Filters the emails which have attachment. | [optional] 
**LastModifiedEndTimeSecs** | Pointer to **NullableInt64** | Specifies the end time in Unix timestamp epoch in seconds where the last modification time of the email/item is less than specified value. | [optional] 
**LastModifiedStartTimeSecs** | Pointer to **NullableInt64** | Specifies the start time in Unix timestamp epoch in seconds where the last modification time of the email/item is more than specified value. | [optional] 
**LastName** | Pointer to **NullableString** | Filters the contacts with specified text in last name. | [optional] 
**MiddleName** | Pointer to **NullableString** | Filters the contacts with specified text in middle name. | [optional] 
**OrganizerAddress** | Pointer to **NullableString** | Filters the calendar items which are organized by specified User&#39;s email address. | [optional] 
**ReceivedEndTimeSecs** | Pointer to **NullableInt64** | Specifies the end time in Unix timestamp epoch in seconds where the received time of the email is less than specified value. | [optional] 
**ReceivedStartTimeSecs** | Pointer to **NullableInt64** | Specifies the start time in Unix timestamp epoch in seconds where the received time of the email is more than specified value. | [optional] 
**RecipientAddresses** | Pointer to **[]string** | Filters the emails which are sent to specified email addresses. | [optional] 
**SenderAddress** | Pointer to **NullableString** | Filters the emails which are received from specified User&#39;s email address. | [optional] 
**SourceEnvironment** | Pointer to **NullableString** | Specifies the source environment. | [optional] 
**TaskStatusTypes** | Pointer to **[]string** | Specifies a list of task item status types. Task items having status within the given types will be returned. | [optional] 
**Types** | Pointer to **[]string** | Specifies a list of mailbox item types. Only items within the given types will be returned. | [optional] 
**O365Params** | Pointer to [**O365SearchEmailsRequestParams**](O365SearchEmailsRequestParams.md) |  | [optional] 

## Methods

### NewSearchEmailRequestParams

`func NewSearchEmailRequestParams() *SearchEmailRequestParams`

NewSearchEmailRequestParams instantiates a new SearchEmailRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchEmailRequestParamsWithDefaults

`func NewSearchEmailRequestParamsWithDefaults() *SearchEmailRequestParams`

NewSearchEmailRequestParamsWithDefaults instantiates a new SearchEmailRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttendeesAddresses

`func (o *SearchEmailRequestParams) GetAttendeesAddresses() []string`

GetAttendeesAddresses returns the AttendeesAddresses field if non-nil, zero value otherwise.

### GetAttendeesAddressesOk

`func (o *SearchEmailRequestParams) GetAttendeesAddressesOk() (*[]string, bool)`

GetAttendeesAddressesOk returns a tuple with the AttendeesAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendeesAddresses

`func (o *SearchEmailRequestParams) SetAttendeesAddresses(v []string)`

SetAttendeesAddresses sets AttendeesAddresses field to given value.

### HasAttendeesAddresses

`func (o *SearchEmailRequestParams) HasAttendeesAddresses() bool`

HasAttendeesAddresses returns a boolean if a field has been set.

### SetAttendeesAddressesNil

`func (o *SearchEmailRequestParams) SetAttendeesAddressesNil(b bool)`

 SetAttendeesAddressesNil sets the value for AttendeesAddresses to be an explicit nil

### UnsetAttendeesAddresses
`func (o *SearchEmailRequestParams) UnsetAttendeesAddresses()`

UnsetAttendeesAddresses ensures that no value is present for AttendeesAddresses, not even an explicit nil
### GetBccRecipientAddresses

`func (o *SearchEmailRequestParams) GetBccRecipientAddresses() []string`

GetBccRecipientAddresses returns the BccRecipientAddresses field if non-nil, zero value otherwise.

### GetBccRecipientAddressesOk

`func (o *SearchEmailRequestParams) GetBccRecipientAddressesOk() (*[]string, bool)`

GetBccRecipientAddressesOk returns a tuple with the BccRecipientAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBccRecipientAddresses

`func (o *SearchEmailRequestParams) SetBccRecipientAddresses(v []string)`

SetBccRecipientAddresses sets BccRecipientAddresses field to given value.

### HasBccRecipientAddresses

`func (o *SearchEmailRequestParams) HasBccRecipientAddresses() bool`

HasBccRecipientAddresses returns a boolean if a field has been set.

### SetBccRecipientAddressesNil

`func (o *SearchEmailRequestParams) SetBccRecipientAddressesNil(b bool)`

 SetBccRecipientAddressesNil sets the value for BccRecipientAddresses to be an explicit nil

### UnsetBccRecipientAddresses
`func (o *SearchEmailRequestParams) UnsetBccRecipientAddresses()`

UnsetBccRecipientAddresses ensures that no value is present for BccRecipientAddresses, not even an explicit nil
### GetCcRecipientAddresses

`func (o *SearchEmailRequestParams) GetCcRecipientAddresses() []string`

GetCcRecipientAddresses returns the CcRecipientAddresses field if non-nil, zero value otherwise.

### GetCcRecipientAddressesOk

`func (o *SearchEmailRequestParams) GetCcRecipientAddressesOk() (*[]string, bool)`

GetCcRecipientAddressesOk returns a tuple with the CcRecipientAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcRecipientAddresses

`func (o *SearchEmailRequestParams) SetCcRecipientAddresses(v []string)`

SetCcRecipientAddresses sets CcRecipientAddresses field to given value.

### HasCcRecipientAddresses

`func (o *SearchEmailRequestParams) HasCcRecipientAddresses() bool`

HasCcRecipientAddresses returns a boolean if a field has been set.

### SetCcRecipientAddressesNil

`func (o *SearchEmailRequestParams) SetCcRecipientAddressesNil(b bool)`

 SetCcRecipientAddressesNil sets the value for CcRecipientAddresses to be an explicit nil

### UnsetCcRecipientAddresses
`func (o *SearchEmailRequestParams) UnsetCcRecipientAddresses()`

UnsetCcRecipientAddresses ensures that no value is present for CcRecipientAddresses, not even an explicit nil
### GetCreatedEndTimeSecs

`func (o *SearchEmailRequestParams) GetCreatedEndTimeSecs() int64`

GetCreatedEndTimeSecs returns the CreatedEndTimeSecs field if non-nil, zero value otherwise.

### GetCreatedEndTimeSecsOk

`func (o *SearchEmailRequestParams) GetCreatedEndTimeSecsOk() (*int64, bool)`

GetCreatedEndTimeSecsOk returns a tuple with the CreatedEndTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedEndTimeSecs

`func (o *SearchEmailRequestParams) SetCreatedEndTimeSecs(v int64)`

SetCreatedEndTimeSecs sets CreatedEndTimeSecs field to given value.

### HasCreatedEndTimeSecs

`func (o *SearchEmailRequestParams) HasCreatedEndTimeSecs() bool`

HasCreatedEndTimeSecs returns a boolean if a field has been set.

### SetCreatedEndTimeSecsNil

`func (o *SearchEmailRequestParams) SetCreatedEndTimeSecsNil(b bool)`

 SetCreatedEndTimeSecsNil sets the value for CreatedEndTimeSecs to be an explicit nil

### UnsetCreatedEndTimeSecs
`func (o *SearchEmailRequestParams) UnsetCreatedEndTimeSecs()`

UnsetCreatedEndTimeSecs ensures that no value is present for CreatedEndTimeSecs, not even an explicit nil
### GetCreatedStartTimeSecs

`func (o *SearchEmailRequestParams) GetCreatedStartTimeSecs() int64`

GetCreatedStartTimeSecs returns the CreatedStartTimeSecs field if non-nil, zero value otherwise.

### GetCreatedStartTimeSecsOk

`func (o *SearchEmailRequestParams) GetCreatedStartTimeSecsOk() (*int64, bool)`

GetCreatedStartTimeSecsOk returns a tuple with the CreatedStartTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedStartTimeSecs

`func (o *SearchEmailRequestParams) SetCreatedStartTimeSecs(v int64)`

SetCreatedStartTimeSecs sets CreatedStartTimeSecs field to given value.

### HasCreatedStartTimeSecs

`func (o *SearchEmailRequestParams) HasCreatedStartTimeSecs() bool`

HasCreatedStartTimeSecs returns a boolean if a field has been set.

### SetCreatedStartTimeSecsNil

`func (o *SearchEmailRequestParams) SetCreatedStartTimeSecsNil(b bool)`

 SetCreatedStartTimeSecsNil sets the value for CreatedStartTimeSecs to be an explicit nil

### UnsetCreatedStartTimeSecs
`func (o *SearchEmailRequestParams) UnsetCreatedStartTimeSecs()`

UnsetCreatedStartTimeSecs ensures that no value is present for CreatedStartTimeSecs, not even an explicit nil
### GetDueDateEndTimeSecs

`func (o *SearchEmailRequestParams) GetDueDateEndTimeSecs() int64`

GetDueDateEndTimeSecs returns the DueDateEndTimeSecs field if non-nil, zero value otherwise.

### GetDueDateEndTimeSecsOk

`func (o *SearchEmailRequestParams) GetDueDateEndTimeSecsOk() (*int64, bool)`

GetDueDateEndTimeSecsOk returns a tuple with the DueDateEndTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDateEndTimeSecs

`func (o *SearchEmailRequestParams) SetDueDateEndTimeSecs(v int64)`

SetDueDateEndTimeSecs sets DueDateEndTimeSecs field to given value.

### HasDueDateEndTimeSecs

`func (o *SearchEmailRequestParams) HasDueDateEndTimeSecs() bool`

HasDueDateEndTimeSecs returns a boolean if a field has been set.

### SetDueDateEndTimeSecsNil

`func (o *SearchEmailRequestParams) SetDueDateEndTimeSecsNil(b bool)`

 SetDueDateEndTimeSecsNil sets the value for DueDateEndTimeSecs to be an explicit nil

### UnsetDueDateEndTimeSecs
`func (o *SearchEmailRequestParams) UnsetDueDateEndTimeSecs()`

UnsetDueDateEndTimeSecs ensures that no value is present for DueDateEndTimeSecs, not even an explicit nil
### GetDueDateStartTimeSecs

`func (o *SearchEmailRequestParams) GetDueDateStartTimeSecs() int64`

GetDueDateStartTimeSecs returns the DueDateStartTimeSecs field if non-nil, zero value otherwise.

### GetDueDateStartTimeSecsOk

`func (o *SearchEmailRequestParams) GetDueDateStartTimeSecsOk() (*int64, bool)`

GetDueDateStartTimeSecsOk returns a tuple with the DueDateStartTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDateStartTimeSecs

`func (o *SearchEmailRequestParams) SetDueDateStartTimeSecs(v int64)`

SetDueDateStartTimeSecs sets DueDateStartTimeSecs field to given value.

### HasDueDateStartTimeSecs

`func (o *SearchEmailRequestParams) HasDueDateStartTimeSecs() bool`

HasDueDateStartTimeSecs returns a boolean if a field has been set.

### SetDueDateStartTimeSecsNil

`func (o *SearchEmailRequestParams) SetDueDateStartTimeSecsNil(b bool)`

 SetDueDateStartTimeSecsNil sets the value for DueDateStartTimeSecs to be an explicit nil

### UnsetDueDateStartTimeSecs
`func (o *SearchEmailRequestParams) UnsetDueDateStartTimeSecs()`

UnsetDueDateStartTimeSecs ensures that no value is present for DueDateStartTimeSecs, not even an explicit nil
### GetEmailAddress

`func (o *SearchEmailRequestParams) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *SearchEmailRequestParams) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *SearchEmailRequestParams) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *SearchEmailRequestParams) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *SearchEmailRequestParams) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *SearchEmailRequestParams) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetEmailSubject

`func (o *SearchEmailRequestParams) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *SearchEmailRequestParams) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *SearchEmailRequestParams) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.

### HasEmailSubject

`func (o *SearchEmailRequestParams) HasEmailSubject() bool`

HasEmailSubject returns a boolean if a field has been set.

### SetEmailSubjectNil

`func (o *SearchEmailRequestParams) SetEmailSubjectNil(b bool)`

 SetEmailSubjectNil sets the value for EmailSubject to be an explicit nil

### UnsetEmailSubject
`func (o *SearchEmailRequestParams) UnsetEmailSubject()`

UnsetEmailSubject ensures that no value is present for EmailSubject, not even an explicit nil
### GetFirstName

`func (o *SearchEmailRequestParams) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *SearchEmailRequestParams) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *SearchEmailRequestParams) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *SearchEmailRequestParams) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *SearchEmailRequestParams) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *SearchEmailRequestParams) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetFolderNames

`func (o *SearchEmailRequestParams) GetFolderNames() []string`

GetFolderNames returns the FolderNames field if non-nil, zero value otherwise.

### GetFolderNamesOk

`func (o *SearchEmailRequestParams) GetFolderNamesOk() (*[]string, bool)`

GetFolderNamesOk returns a tuple with the FolderNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderNames

`func (o *SearchEmailRequestParams) SetFolderNames(v []string)`

SetFolderNames sets FolderNames field to given value.

### HasFolderNames

`func (o *SearchEmailRequestParams) HasFolderNames() bool`

HasFolderNames returns a boolean if a field has been set.

### SetFolderNamesNil

`func (o *SearchEmailRequestParams) SetFolderNamesNil(b bool)`

 SetFolderNamesNil sets the value for FolderNames to be an explicit nil

### UnsetFolderNames
`func (o *SearchEmailRequestParams) UnsetFolderNames()`

UnsetFolderNames ensures that no value is present for FolderNames, not even an explicit nil
### GetHasAttachment

`func (o *SearchEmailRequestParams) GetHasAttachment() bool`

GetHasAttachment returns the HasAttachment field if non-nil, zero value otherwise.

### GetHasAttachmentOk

`func (o *SearchEmailRequestParams) GetHasAttachmentOk() (*bool, bool)`

GetHasAttachmentOk returns a tuple with the HasAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachment

`func (o *SearchEmailRequestParams) SetHasAttachment(v bool)`

SetHasAttachment sets HasAttachment field to given value.

### HasHasAttachment

`func (o *SearchEmailRequestParams) HasHasAttachment() bool`

HasHasAttachment returns a boolean if a field has been set.

### SetHasAttachmentNil

`func (o *SearchEmailRequestParams) SetHasAttachmentNil(b bool)`

 SetHasAttachmentNil sets the value for HasAttachment to be an explicit nil

### UnsetHasAttachment
`func (o *SearchEmailRequestParams) UnsetHasAttachment()`

UnsetHasAttachment ensures that no value is present for HasAttachment, not even an explicit nil
### GetLastModifiedEndTimeSecs

`func (o *SearchEmailRequestParams) GetLastModifiedEndTimeSecs() int64`

GetLastModifiedEndTimeSecs returns the LastModifiedEndTimeSecs field if non-nil, zero value otherwise.

### GetLastModifiedEndTimeSecsOk

`func (o *SearchEmailRequestParams) GetLastModifiedEndTimeSecsOk() (*int64, bool)`

GetLastModifiedEndTimeSecsOk returns a tuple with the LastModifiedEndTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedEndTimeSecs

`func (o *SearchEmailRequestParams) SetLastModifiedEndTimeSecs(v int64)`

SetLastModifiedEndTimeSecs sets LastModifiedEndTimeSecs field to given value.

### HasLastModifiedEndTimeSecs

`func (o *SearchEmailRequestParams) HasLastModifiedEndTimeSecs() bool`

HasLastModifiedEndTimeSecs returns a boolean if a field has been set.

### SetLastModifiedEndTimeSecsNil

`func (o *SearchEmailRequestParams) SetLastModifiedEndTimeSecsNil(b bool)`

 SetLastModifiedEndTimeSecsNil sets the value for LastModifiedEndTimeSecs to be an explicit nil

### UnsetLastModifiedEndTimeSecs
`func (o *SearchEmailRequestParams) UnsetLastModifiedEndTimeSecs()`

UnsetLastModifiedEndTimeSecs ensures that no value is present for LastModifiedEndTimeSecs, not even an explicit nil
### GetLastModifiedStartTimeSecs

`func (o *SearchEmailRequestParams) GetLastModifiedStartTimeSecs() int64`

GetLastModifiedStartTimeSecs returns the LastModifiedStartTimeSecs field if non-nil, zero value otherwise.

### GetLastModifiedStartTimeSecsOk

`func (o *SearchEmailRequestParams) GetLastModifiedStartTimeSecsOk() (*int64, bool)`

GetLastModifiedStartTimeSecsOk returns a tuple with the LastModifiedStartTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedStartTimeSecs

`func (o *SearchEmailRequestParams) SetLastModifiedStartTimeSecs(v int64)`

SetLastModifiedStartTimeSecs sets LastModifiedStartTimeSecs field to given value.

### HasLastModifiedStartTimeSecs

`func (o *SearchEmailRequestParams) HasLastModifiedStartTimeSecs() bool`

HasLastModifiedStartTimeSecs returns a boolean if a field has been set.

### SetLastModifiedStartTimeSecsNil

`func (o *SearchEmailRequestParams) SetLastModifiedStartTimeSecsNil(b bool)`

 SetLastModifiedStartTimeSecsNil sets the value for LastModifiedStartTimeSecs to be an explicit nil

### UnsetLastModifiedStartTimeSecs
`func (o *SearchEmailRequestParams) UnsetLastModifiedStartTimeSecs()`

UnsetLastModifiedStartTimeSecs ensures that no value is present for LastModifiedStartTimeSecs, not even an explicit nil
### GetLastName

`func (o *SearchEmailRequestParams) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *SearchEmailRequestParams) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *SearchEmailRequestParams) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *SearchEmailRequestParams) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *SearchEmailRequestParams) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *SearchEmailRequestParams) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetMiddleName

`func (o *SearchEmailRequestParams) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *SearchEmailRequestParams) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *SearchEmailRequestParams) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *SearchEmailRequestParams) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *SearchEmailRequestParams) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *SearchEmailRequestParams) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetOrganizerAddress

`func (o *SearchEmailRequestParams) GetOrganizerAddress() string`

GetOrganizerAddress returns the OrganizerAddress field if non-nil, zero value otherwise.

### GetOrganizerAddressOk

`func (o *SearchEmailRequestParams) GetOrganizerAddressOk() (*string, bool)`

GetOrganizerAddressOk returns a tuple with the OrganizerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizerAddress

`func (o *SearchEmailRequestParams) SetOrganizerAddress(v string)`

SetOrganizerAddress sets OrganizerAddress field to given value.

### HasOrganizerAddress

`func (o *SearchEmailRequestParams) HasOrganizerAddress() bool`

HasOrganizerAddress returns a boolean if a field has been set.

### SetOrganizerAddressNil

`func (o *SearchEmailRequestParams) SetOrganizerAddressNil(b bool)`

 SetOrganizerAddressNil sets the value for OrganizerAddress to be an explicit nil

### UnsetOrganizerAddress
`func (o *SearchEmailRequestParams) UnsetOrganizerAddress()`

UnsetOrganizerAddress ensures that no value is present for OrganizerAddress, not even an explicit nil
### GetReceivedEndTimeSecs

`func (o *SearchEmailRequestParams) GetReceivedEndTimeSecs() int64`

GetReceivedEndTimeSecs returns the ReceivedEndTimeSecs field if non-nil, zero value otherwise.

### GetReceivedEndTimeSecsOk

`func (o *SearchEmailRequestParams) GetReceivedEndTimeSecsOk() (*int64, bool)`

GetReceivedEndTimeSecsOk returns a tuple with the ReceivedEndTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedEndTimeSecs

`func (o *SearchEmailRequestParams) SetReceivedEndTimeSecs(v int64)`

SetReceivedEndTimeSecs sets ReceivedEndTimeSecs field to given value.

### HasReceivedEndTimeSecs

`func (o *SearchEmailRequestParams) HasReceivedEndTimeSecs() bool`

HasReceivedEndTimeSecs returns a boolean if a field has been set.

### SetReceivedEndTimeSecsNil

`func (o *SearchEmailRequestParams) SetReceivedEndTimeSecsNil(b bool)`

 SetReceivedEndTimeSecsNil sets the value for ReceivedEndTimeSecs to be an explicit nil

### UnsetReceivedEndTimeSecs
`func (o *SearchEmailRequestParams) UnsetReceivedEndTimeSecs()`

UnsetReceivedEndTimeSecs ensures that no value is present for ReceivedEndTimeSecs, not even an explicit nil
### GetReceivedStartTimeSecs

`func (o *SearchEmailRequestParams) GetReceivedStartTimeSecs() int64`

GetReceivedStartTimeSecs returns the ReceivedStartTimeSecs field if non-nil, zero value otherwise.

### GetReceivedStartTimeSecsOk

`func (o *SearchEmailRequestParams) GetReceivedStartTimeSecsOk() (*int64, bool)`

GetReceivedStartTimeSecsOk returns a tuple with the ReceivedStartTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedStartTimeSecs

`func (o *SearchEmailRequestParams) SetReceivedStartTimeSecs(v int64)`

SetReceivedStartTimeSecs sets ReceivedStartTimeSecs field to given value.

### HasReceivedStartTimeSecs

`func (o *SearchEmailRequestParams) HasReceivedStartTimeSecs() bool`

HasReceivedStartTimeSecs returns a boolean if a field has been set.

### SetReceivedStartTimeSecsNil

`func (o *SearchEmailRequestParams) SetReceivedStartTimeSecsNil(b bool)`

 SetReceivedStartTimeSecsNil sets the value for ReceivedStartTimeSecs to be an explicit nil

### UnsetReceivedStartTimeSecs
`func (o *SearchEmailRequestParams) UnsetReceivedStartTimeSecs()`

UnsetReceivedStartTimeSecs ensures that no value is present for ReceivedStartTimeSecs, not even an explicit nil
### GetRecipientAddresses

`func (o *SearchEmailRequestParams) GetRecipientAddresses() []string`

GetRecipientAddresses returns the RecipientAddresses field if non-nil, zero value otherwise.

### GetRecipientAddressesOk

`func (o *SearchEmailRequestParams) GetRecipientAddressesOk() (*[]string, bool)`

GetRecipientAddressesOk returns a tuple with the RecipientAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddresses

`func (o *SearchEmailRequestParams) SetRecipientAddresses(v []string)`

SetRecipientAddresses sets RecipientAddresses field to given value.

### HasRecipientAddresses

`func (o *SearchEmailRequestParams) HasRecipientAddresses() bool`

HasRecipientAddresses returns a boolean if a field has been set.

### SetRecipientAddressesNil

`func (o *SearchEmailRequestParams) SetRecipientAddressesNil(b bool)`

 SetRecipientAddressesNil sets the value for RecipientAddresses to be an explicit nil

### UnsetRecipientAddresses
`func (o *SearchEmailRequestParams) UnsetRecipientAddresses()`

UnsetRecipientAddresses ensures that no value is present for RecipientAddresses, not even an explicit nil
### GetSenderAddress

`func (o *SearchEmailRequestParams) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *SearchEmailRequestParams) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *SearchEmailRequestParams) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.

### HasSenderAddress

`func (o *SearchEmailRequestParams) HasSenderAddress() bool`

HasSenderAddress returns a boolean if a field has been set.

### SetSenderAddressNil

`func (o *SearchEmailRequestParams) SetSenderAddressNil(b bool)`

 SetSenderAddressNil sets the value for SenderAddress to be an explicit nil

### UnsetSenderAddress
`func (o *SearchEmailRequestParams) UnsetSenderAddress()`

UnsetSenderAddress ensures that no value is present for SenderAddress, not even an explicit nil
### GetSourceEnvironment

`func (o *SearchEmailRequestParams) GetSourceEnvironment() string`

GetSourceEnvironment returns the SourceEnvironment field if non-nil, zero value otherwise.

### GetSourceEnvironmentOk

`func (o *SearchEmailRequestParams) GetSourceEnvironmentOk() (*string, bool)`

GetSourceEnvironmentOk returns a tuple with the SourceEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEnvironment

`func (o *SearchEmailRequestParams) SetSourceEnvironment(v string)`

SetSourceEnvironment sets SourceEnvironment field to given value.

### HasSourceEnvironment

`func (o *SearchEmailRequestParams) HasSourceEnvironment() bool`

HasSourceEnvironment returns a boolean if a field has been set.

### SetSourceEnvironmentNil

`func (o *SearchEmailRequestParams) SetSourceEnvironmentNil(b bool)`

 SetSourceEnvironmentNil sets the value for SourceEnvironment to be an explicit nil

### UnsetSourceEnvironment
`func (o *SearchEmailRequestParams) UnsetSourceEnvironment()`

UnsetSourceEnvironment ensures that no value is present for SourceEnvironment, not even an explicit nil
### GetTaskStatusTypes

`func (o *SearchEmailRequestParams) GetTaskStatusTypes() []string`

GetTaskStatusTypes returns the TaskStatusTypes field if non-nil, zero value otherwise.

### GetTaskStatusTypesOk

`func (o *SearchEmailRequestParams) GetTaskStatusTypesOk() (*[]string, bool)`

GetTaskStatusTypesOk returns a tuple with the TaskStatusTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStatusTypes

`func (o *SearchEmailRequestParams) SetTaskStatusTypes(v []string)`

SetTaskStatusTypes sets TaskStatusTypes field to given value.

### HasTaskStatusTypes

`func (o *SearchEmailRequestParams) HasTaskStatusTypes() bool`

HasTaskStatusTypes returns a boolean if a field has been set.

### SetTaskStatusTypesNil

`func (o *SearchEmailRequestParams) SetTaskStatusTypesNil(b bool)`

 SetTaskStatusTypesNil sets the value for TaskStatusTypes to be an explicit nil

### UnsetTaskStatusTypes
`func (o *SearchEmailRequestParams) UnsetTaskStatusTypes()`

UnsetTaskStatusTypes ensures that no value is present for TaskStatusTypes, not even an explicit nil
### GetTypes

`func (o *SearchEmailRequestParams) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *SearchEmailRequestParams) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *SearchEmailRequestParams) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *SearchEmailRequestParams) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### SetTypesNil

`func (o *SearchEmailRequestParams) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *SearchEmailRequestParams) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil
### GetO365Params

`func (o *SearchEmailRequestParams) GetO365Params() O365SearchEmailsRequestParams`

GetO365Params returns the O365Params field if non-nil, zero value otherwise.

### GetO365ParamsOk

`func (o *SearchEmailRequestParams) GetO365ParamsOk() (*O365SearchEmailsRequestParams, bool)`

GetO365ParamsOk returns a tuple with the O365Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO365Params

`func (o *SearchEmailRequestParams) SetO365Params(v O365SearchEmailsRequestParams)`

SetO365Params sets O365Params field to given value.

### HasO365Params

`func (o *SearchEmailRequestParams) HasO365Params() bool`

HasO365Params returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


