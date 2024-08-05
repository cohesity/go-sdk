# Email

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotTags** | Pointer to [**[]SnapshotTagInfo**](SnapshotTagInfo.md) | Specifies snapshot tags applied to the object. | [optional] 
**Tags** | Pointer to [**[]TagInfo**](TagInfo.md) | Specifies tag applied to the object. | [optional] 
**BccRecipientAddresses** | Pointer to **[]string** | \&quot;Specifies the email addresses of all the BCC receipients of this email.\&quot; | [optional] 
**CcRecipientAddresses** | Pointer to **[]string** | \&quot;Specifies the email addresses of all the CC receipients of this email.\&quot; | [optional] 
**CreatedTimeSecs** | Pointer to **NullableInt64** | \&quot;Specifies the Unix timestamp epoch in seconds at which this item is created.\&quot; | [optional] 
**DirectoryPath** | Pointer to **NullableString** | Specifies the directory path to this mailbox item. | [optional] 
**EmailAddresses** | Pointer to **[]string** | Specifies the email addresses of a contact. | [optional] 
**EmailSubject** | Pointer to **NullableString** | Specifies the subject of this email. | [optional] 
**FirstName** | Pointer to **NullableString** | Specifies the contact&#39;s first name. | [optional] 
**FolderName** | Pointer to **NullableString** | Specify the name of the email folder. | [optional] 
**HasAttachment** | Pointer to **NullableBool** | Specifies whether email has an attachment. | [optional] 
**Id** | Pointer to **NullableString** | Specifies the id of the email object. | [optional] 
**LastModificationName** | Pointer to **NullableString** | \&quot;Specifies the name of the person who modified this item.\&quot; | [optional] 
**LastModificationTimeSecs** | Pointer to **NullableInt64** | \&quot;Specifies the Unix timestamp epoch in seconds at which this item was modified.\&quot; | [optional] 
**LastName** | Pointer to **NullableString** | Specifies the contact&#39;s last name. | [optional] 
**OptionalAttendeesAddresses** | Pointer to **[]string** | \&quot;Specifies the email addresses of all the optional attendees of this calendar item.\&quot; | [optional] 
**OrganizerAddress** | Pointer to **NullableString** | \&quot;Specifies the calendar item organizer&#39;s email address.\&quot; | [optional] 
**ParentFolderId** | Pointer to **NullableInt64** | Specifies the id of parent folder the mailbox item. | [optional] 
**Path** | Pointer to **NullableString** | Specifies the path to this mailbox item. | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | \&quot;Specifies the Protection Group id protecting the mailbox.\&quot; | [optional] 
**ProtectionGroupName** | Pointer to **NullableString** | \&quot;Specifies the Protection Group name protecting the mailbox item.\&quot; | [optional] 
**ReceivedTimeSecs** | Pointer to **NullableInt64** | \&quot;Specifies the Unix timestamp epoch in seconds at which this email is received.\&quot; | [optional] 
**RecipientAddresses** | Pointer to **[]string** | \&quot;Specifies the email addresses of all receipients of this email.\&quot; | [optional] 
**RequiredAttendeesAddresses** | Pointer to **[]string** | \&quot;Specifies the email addresses of all required attendees of this calendar item.\&quot; | [optional] 
**SenderAddress** | Pointer to **NullableString** | Specifies the sender&#39;s email address. | [optional] 
**SentTimeSecs** | Pointer to **NullableInt64** | \&quot;Specifies the Unix timestamp epoch in seconds at which this email is sent.\&quot; | [optional] 
**StorageDomainId** | Pointer to **NullableInt64** | \&quot;Specifies the Storage Domain id where the backup data of Object is present.\&quot; | [optional] 
**TaskCompletionDateTimeSecs** | Pointer to **NullableInt64** | \&quot;Specifies the Unix timestamp epoch in seconds at which this task item was completed.\&quot; | [optional] 
**TaskDueDateTimeSecs** | Pointer to **NullableInt64** | \&quot;Specifies the Unix timestamp epoch in seconds at which this task item is due.\&quot; | [optional] 
**TaskStatus** | Pointer to **NullableString** | Specifies the task item status type. | [optional] 
**TenantId** | Pointer to **NullableString** | \&quot;Specify the tenant id to which this email belongs to.\&quot; | [optional] 
**Type** | Pointer to **NullableString** | Specifies the Mailbox item type. | [optional] 
**UserObjectInfo** | Pointer to [**ObjectSummary**](ObjectSummary.md) |  | [optional] 

## Methods

### NewEmail

`func NewEmail() *Email`

NewEmail instantiates a new Email object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailWithDefaults

`func NewEmailWithDefaults() *Email`

NewEmailWithDefaults instantiates a new Email object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotTags

`func (o *Email) GetSnapshotTags() []SnapshotTagInfo`

GetSnapshotTags returns the SnapshotTags field if non-nil, zero value otherwise.

### GetSnapshotTagsOk

`func (o *Email) GetSnapshotTagsOk() (*[]SnapshotTagInfo, bool)`

GetSnapshotTagsOk returns a tuple with the SnapshotTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTags

`func (o *Email) SetSnapshotTags(v []SnapshotTagInfo)`

SetSnapshotTags sets SnapshotTags field to given value.

### HasSnapshotTags

`func (o *Email) HasSnapshotTags() bool`

HasSnapshotTags returns a boolean if a field has been set.

### SetSnapshotTagsNil

`func (o *Email) SetSnapshotTagsNil(b bool)`

 SetSnapshotTagsNil sets the value for SnapshotTags to be an explicit nil

### UnsetSnapshotTags
`func (o *Email) UnsetSnapshotTags()`

UnsetSnapshotTags ensures that no value is present for SnapshotTags, not even an explicit nil
### GetTags

`func (o *Email) GetTags() []TagInfo`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Email) GetTagsOk() (*[]TagInfo, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Email) SetTags(v []TagInfo)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Email) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *Email) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *Email) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetBccRecipientAddresses

`func (o *Email) GetBccRecipientAddresses() []string`

GetBccRecipientAddresses returns the BccRecipientAddresses field if non-nil, zero value otherwise.

### GetBccRecipientAddressesOk

`func (o *Email) GetBccRecipientAddressesOk() (*[]string, bool)`

GetBccRecipientAddressesOk returns a tuple with the BccRecipientAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBccRecipientAddresses

`func (o *Email) SetBccRecipientAddresses(v []string)`

SetBccRecipientAddresses sets BccRecipientAddresses field to given value.

### HasBccRecipientAddresses

`func (o *Email) HasBccRecipientAddresses() bool`

HasBccRecipientAddresses returns a boolean if a field has been set.

### SetBccRecipientAddressesNil

`func (o *Email) SetBccRecipientAddressesNil(b bool)`

 SetBccRecipientAddressesNil sets the value for BccRecipientAddresses to be an explicit nil

### UnsetBccRecipientAddresses
`func (o *Email) UnsetBccRecipientAddresses()`

UnsetBccRecipientAddresses ensures that no value is present for BccRecipientAddresses, not even an explicit nil
### GetCcRecipientAddresses

`func (o *Email) GetCcRecipientAddresses() []string`

GetCcRecipientAddresses returns the CcRecipientAddresses field if non-nil, zero value otherwise.

### GetCcRecipientAddressesOk

`func (o *Email) GetCcRecipientAddressesOk() (*[]string, bool)`

GetCcRecipientAddressesOk returns a tuple with the CcRecipientAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcRecipientAddresses

`func (o *Email) SetCcRecipientAddresses(v []string)`

SetCcRecipientAddresses sets CcRecipientAddresses field to given value.

### HasCcRecipientAddresses

`func (o *Email) HasCcRecipientAddresses() bool`

HasCcRecipientAddresses returns a boolean if a field has been set.

### SetCcRecipientAddressesNil

`func (o *Email) SetCcRecipientAddressesNil(b bool)`

 SetCcRecipientAddressesNil sets the value for CcRecipientAddresses to be an explicit nil

### UnsetCcRecipientAddresses
`func (o *Email) UnsetCcRecipientAddresses()`

UnsetCcRecipientAddresses ensures that no value is present for CcRecipientAddresses, not even an explicit nil
### GetCreatedTimeSecs

`func (o *Email) GetCreatedTimeSecs() int64`

GetCreatedTimeSecs returns the CreatedTimeSecs field if non-nil, zero value otherwise.

### GetCreatedTimeSecsOk

`func (o *Email) GetCreatedTimeSecsOk() (*int64, bool)`

GetCreatedTimeSecsOk returns a tuple with the CreatedTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimeSecs

`func (o *Email) SetCreatedTimeSecs(v int64)`

SetCreatedTimeSecs sets CreatedTimeSecs field to given value.

### HasCreatedTimeSecs

`func (o *Email) HasCreatedTimeSecs() bool`

HasCreatedTimeSecs returns a boolean if a field has been set.

### SetCreatedTimeSecsNil

`func (o *Email) SetCreatedTimeSecsNil(b bool)`

 SetCreatedTimeSecsNil sets the value for CreatedTimeSecs to be an explicit nil

### UnsetCreatedTimeSecs
`func (o *Email) UnsetCreatedTimeSecs()`

UnsetCreatedTimeSecs ensures that no value is present for CreatedTimeSecs, not even an explicit nil
### GetDirectoryPath

`func (o *Email) GetDirectoryPath() string`

GetDirectoryPath returns the DirectoryPath field if non-nil, zero value otherwise.

### GetDirectoryPathOk

`func (o *Email) GetDirectoryPathOk() (*string, bool)`

GetDirectoryPathOk returns a tuple with the DirectoryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryPath

`func (o *Email) SetDirectoryPath(v string)`

SetDirectoryPath sets DirectoryPath field to given value.

### HasDirectoryPath

`func (o *Email) HasDirectoryPath() bool`

HasDirectoryPath returns a boolean if a field has been set.

### SetDirectoryPathNil

`func (o *Email) SetDirectoryPathNil(b bool)`

 SetDirectoryPathNil sets the value for DirectoryPath to be an explicit nil

### UnsetDirectoryPath
`func (o *Email) UnsetDirectoryPath()`

UnsetDirectoryPath ensures that no value is present for DirectoryPath, not even an explicit nil
### GetEmailAddresses

`func (o *Email) GetEmailAddresses() []string`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *Email) GetEmailAddressesOk() (*[]string, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddresses

`func (o *Email) SetEmailAddresses(v []string)`

SetEmailAddresses sets EmailAddresses field to given value.

### HasEmailAddresses

`func (o *Email) HasEmailAddresses() bool`

HasEmailAddresses returns a boolean if a field has been set.

### SetEmailAddressesNil

`func (o *Email) SetEmailAddressesNil(b bool)`

 SetEmailAddressesNil sets the value for EmailAddresses to be an explicit nil

### UnsetEmailAddresses
`func (o *Email) UnsetEmailAddresses()`

UnsetEmailAddresses ensures that no value is present for EmailAddresses, not even an explicit nil
### GetEmailSubject

`func (o *Email) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *Email) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *Email) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.

### HasEmailSubject

`func (o *Email) HasEmailSubject() bool`

HasEmailSubject returns a boolean if a field has been set.

### SetEmailSubjectNil

`func (o *Email) SetEmailSubjectNil(b bool)`

 SetEmailSubjectNil sets the value for EmailSubject to be an explicit nil

### UnsetEmailSubject
`func (o *Email) UnsetEmailSubject()`

UnsetEmailSubject ensures that no value is present for EmailSubject, not even an explicit nil
### GetFirstName

`func (o *Email) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Email) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Email) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Email) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *Email) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *Email) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetFolderName

`func (o *Email) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *Email) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *Email) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.

### HasFolderName

`func (o *Email) HasFolderName() bool`

HasFolderName returns a boolean if a field has been set.

### SetFolderNameNil

`func (o *Email) SetFolderNameNil(b bool)`

 SetFolderNameNil sets the value for FolderName to be an explicit nil

### UnsetFolderName
`func (o *Email) UnsetFolderName()`

UnsetFolderName ensures that no value is present for FolderName, not even an explicit nil
### GetHasAttachment

`func (o *Email) GetHasAttachment() bool`

GetHasAttachment returns the HasAttachment field if non-nil, zero value otherwise.

### GetHasAttachmentOk

`func (o *Email) GetHasAttachmentOk() (*bool, bool)`

GetHasAttachmentOk returns a tuple with the HasAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachment

`func (o *Email) SetHasAttachment(v bool)`

SetHasAttachment sets HasAttachment field to given value.

### HasHasAttachment

`func (o *Email) HasHasAttachment() bool`

HasHasAttachment returns a boolean if a field has been set.

### SetHasAttachmentNil

`func (o *Email) SetHasAttachmentNil(b bool)`

 SetHasAttachmentNil sets the value for HasAttachment to be an explicit nil

### UnsetHasAttachment
`func (o *Email) UnsetHasAttachment()`

UnsetHasAttachment ensures that no value is present for HasAttachment, not even an explicit nil
### GetId

`func (o *Email) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Email) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Email) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Email) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Email) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Email) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLastModificationName

`func (o *Email) GetLastModificationName() string`

GetLastModificationName returns the LastModificationName field if non-nil, zero value otherwise.

### GetLastModificationNameOk

`func (o *Email) GetLastModificationNameOk() (*string, bool)`

GetLastModificationNameOk returns a tuple with the LastModificationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationName

`func (o *Email) SetLastModificationName(v string)`

SetLastModificationName sets LastModificationName field to given value.

### HasLastModificationName

`func (o *Email) HasLastModificationName() bool`

HasLastModificationName returns a boolean if a field has been set.

### SetLastModificationNameNil

`func (o *Email) SetLastModificationNameNil(b bool)`

 SetLastModificationNameNil sets the value for LastModificationName to be an explicit nil

### UnsetLastModificationName
`func (o *Email) UnsetLastModificationName()`

UnsetLastModificationName ensures that no value is present for LastModificationName, not even an explicit nil
### GetLastModificationTimeSecs

`func (o *Email) GetLastModificationTimeSecs() int64`

GetLastModificationTimeSecs returns the LastModificationTimeSecs field if non-nil, zero value otherwise.

### GetLastModificationTimeSecsOk

`func (o *Email) GetLastModificationTimeSecsOk() (*int64, bool)`

GetLastModificationTimeSecsOk returns a tuple with the LastModificationTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTimeSecs

`func (o *Email) SetLastModificationTimeSecs(v int64)`

SetLastModificationTimeSecs sets LastModificationTimeSecs field to given value.

### HasLastModificationTimeSecs

`func (o *Email) HasLastModificationTimeSecs() bool`

HasLastModificationTimeSecs returns a boolean if a field has been set.

### SetLastModificationTimeSecsNil

`func (o *Email) SetLastModificationTimeSecsNil(b bool)`

 SetLastModificationTimeSecsNil sets the value for LastModificationTimeSecs to be an explicit nil

### UnsetLastModificationTimeSecs
`func (o *Email) UnsetLastModificationTimeSecs()`

UnsetLastModificationTimeSecs ensures that no value is present for LastModificationTimeSecs, not even an explicit nil
### GetLastName

`func (o *Email) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Email) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Email) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Email) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *Email) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *Email) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetOptionalAttendeesAddresses

`func (o *Email) GetOptionalAttendeesAddresses() []string`

GetOptionalAttendeesAddresses returns the OptionalAttendeesAddresses field if non-nil, zero value otherwise.

### GetOptionalAttendeesAddressesOk

`func (o *Email) GetOptionalAttendeesAddressesOk() (*[]string, bool)`

GetOptionalAttendeesAddressesOk returns a tuple with the OptionalAttendeesAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalAttendeesAddresses

`func (o *Email) SetOptionalAttendeesAddresses(v []string)`

SetOptionalAttendeesAddresses sets OptionalAttendeesAddresses field to given value.

### HasOptionalAttendeesAddresses

`func (o *Email) HasOptionalAttendeesAddresses() bool`

HasOptionalAttendeesAddresses returns a boolean if a field has been set.

### SetOptionalAttendeesAddressesNil

`func (o *Email) SetOptionalAttendeesAddressesNil(b bool)`

 SetOptionalAttendeesAddressesNil sets the value for OptionalAttendeesAddresses to be an explicit nil

### UnsetOptionalAttendeesAddresses
`func (o *Email) UnsetOptionalAttendeesAddresses()`

UnsetOptionalAttendeesAddresses ensures that no value is present for OptionalAttendeesAddresses, not even an explicit nil
### GetOrganizerAddress

`func (o *Email) GetOrganizerAddress() string`

GetOrganizerAddress returns the OrganizerAddress field if non-nil, zero value otherwise.

### GetOrganizerAddressOk

`func (o *Email) GetOrganizerAddressOk() (*string, bool)`

GetOrganizerAddressOk returns a tuple with the OrganizerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizerAddress

`func (o *Email) SetOrganizerAddress(v string)`

SetOrganizerAddress sets OrganizerAddress field to given value.

### HasOrganizerAddress

`func (o *Email) HasOrganizerAddress() bool`

HasOrganizerAddress returns a boolean if a field has been set.

### SetOrganizerAddressNil

`func (o *Email) SetOrganizerAddressNil(b bool)`

 SetOrganizerAddressNil sets the value for OrganizerAddress to be an explicit nil

### UnsetOrganizerAddress
`func (o *Email) UnsetOrganizerAddress()`

UnsetOrganizerAddress ensures that no value is present for OrganizerAddress, not even an explicit nil
### GetParentFolderId

`func (o *Email) GetParentFolderId() int64`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *Email) GetParentFolderIdOk() (*int64, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *Email) SetParentFolderId(v int64)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *Email) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderIdNil

`func (o *Email) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *Email) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
### GetPath

`func (o *Email) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Email) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Email) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Email) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *Email) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *Email) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetProtectionGroupId

`func (o *Email) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *Email) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *Email) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *Email) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *Email) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *Email) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *Email) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *Email) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *Email) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *Email) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *Email) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *Email) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetReceivedTimeSecs

`func (o *Email) GetReceivedTimeSecs() int64`

GetReceivedTimeSecs returns the ReceivedTimeSecs field if non-nil, zero value otherwise.

### GetReceivedTimeSecsOk

`func (o *Email) GetReceivedTimeSecsOk() (*int64, bool)`

GetReceivedTimeSecsOk returns a tuple with the ReceivedTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedTimeSecs

`func (o *Email) SetReceivedTimeSecs(v int64)`

SetReceivedTimeSecs sets ReceivedTimeSecs field to given value.

### HasReceivedTimeSecs

`func (o *Email) HasReceivedTimeSecs() bool`

HasReceivedTimeSecs returns a boolean if a field has been set.

### SetReceivedTimeSecsNil

`func (o *Email) SetReceivedTimeSecsNil(b bool)`

 SetReceivedTimeSecsNil sets the value for ReceivedTimeSecs to be an explicit nil

### UnsetReceivedTimeSecs
`func (o *Email) UnsetReceivedTimeSecs()`

UnsetReceivedTimeSecs ensures that no value is present for ReceivedTimeSecs, not even an explicit nil
### GetRecipientAddresses

`func (o *Email) GetRecipientAddresses() []string`

GetRecipientAddresses returns the RecipientAddresses field if non-nil, zero value otherwise.

### GetRecipientAddressesOk

`func (o *Email) GetRecipientAddressesOk() (*[]string, bool)`

GetRecipientAddressesOk returns a tuple with the RecipientAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddresses

`func (o *Email) SetRecipientAddresses(v []string)`

SetRecipientAddresses sets RecipientAddresses field to given value.

### HasRecipientAddresses

`func (o *Email) HasRecipientAddresses() bool`

HasRecipientAddresses returns a boolean if a field has been set.

### SetRecipientAddressesNil

`func (o *Email) SetRecipientAddressesNil(b bool)`

 SetRecipientAddressesNil sets the value for RecipientAddresses to be an explicit nil

### UnsetRecipientAddresses
`func (o *Email) UnsetRecipientAddresses()`

UnsetRecipientAddresses ensures that no value is present for RecipientAddresses, not even an explicit nil
### GetRequiredAttendeesAddresses

`func (o *Email) GetRequiredAttendeesAddresses() []string`

GetRequiredAttendeesAddresses returns the RequiredAttendeesAddresses field if non-nil, zero value otherwise.

### GetRequiredAttendeesAddressesOk

`func (o *Email) GetRequiredAttendeesAddressesOk() (*[]string, bool)`

GetRequiredAttendeesAddressesOk returns a tuple with the RequiredAttendeesAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredAttendeesAddresses

`func (o *Email) SetRequiredAttendeesAddresses(v []string)`

SetRequiredAttendeesAddresses sets RequiredAttendeesAddresses field to given value.

### HasRequiredAttendeesAddresses

`func (o *Email) HasRequiredAttendeesAddresses() bool`

HasRequiredAttendeesAddresses returns a boolean if a field has been set.

### SetRequiredAttendeesAddressesNil

`func (o *Email) SetRequiredAttendeesAddressesNil(b bool)`

 SetRequiredAttendeesAddressesNil sets the value for RequiredAttendeesAddresses to be an explicit nil

### UnsetRequiredAttendeesAddresses
`func (o *Email) UnsetRequiredAttendeesAddresses()`

UnsetRequiredAttendeesAddresses ensures that no value is present for RequiredAttendeesAddresses, not even an explicit nil
### GetSenderAddress

`func (o *Email) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *Email) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *Email) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.

### HasSenderAddress

`func (o *Email) HasSenderAddress() bool`

HasSenderAddress returns a boolean if a field has been set.

### SetSenderAddressNil

`func (o *Email) SetSenderAddressNil(b bool)`

 SetSenderAddressNil sets the value for SenderAddress to be an explicit nil

### UnsetSenderAddress
`func (o *Email) UnsetSenderAddress()`

UnsetSenderAddress ensures that no value is present for SenderAddress, not even an explicit nil
### GetSentTimeSecs

`func (o *Email) GetSentTimeSecs() int64`

GetSentTimeSecs returns the SentTimeSecs field if non-nil, zero value otherwise.

### GetSentTimeSecsOk

`func (o *Email) GetSentTimeSecsOk() (*int64, bool)`

GetSentTimeSecsOk returns a tuple with the SentTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentTimeSecs

`func (o *Email) SetSentTimeSecs(v int64)`

SetSentTimeSecs sets SentTimeSecs field to given value.

### HasSentTimeSecs

`func (o *Email) HasSentTimeSecs() bool`

HasSentTimeSecs returns a boolean if a field has been set.

### SetSentTimeSecsNil

`func (o *Email) SetSentTimeSecsNil(b bool)`

 SetSentTimeSecsNil sets the value for SentTimeSecs to be an explicit nil

### UnsetSentTimeSecs
`func (o *Email) UnsetSentTimeSecs()`

UnsetSentTimeSecs ensures that no value is present for SentTimeSecs, not even an explicit nil
### GetStorageDomainId

`func (o *Email) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *Email) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *Email) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *Email) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *Email) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *Email) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetTaskCompletionDateTimeSecs

`func (o *Email) GetTaskCompletionDateTimeSecs() int64`

GetTaskCompletionDateTimeSecs returns the TaskCompletionDateTimeSecs field if non-nil, zero value otherwise.

### GetTaskCompletionDateTimeSecsOk

`func (o *Email) GetTaskCompletionDateTimeSecsOk() (*int64, bool)`

GetTaskCompletionDateTimeSecsOk returns a tuple with the TaskCompletionDateTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCompletionDateTimeSecs

`func (o *Email) SetTaskCompletionDateTimeSecs(v int64)`

SetTaskCompletionDateTimeSecs sets TaskCompletionDateTimeSecs field to given value.

### HasTaskCompletionDateTimeSecs

`func (o *Email) HasTaskCompletionDateTimeSecs() bool`

HasTaskCompletionDateTimeSecs returns a boolean if a field has been set.

### SetTaskCompletionDateTimeSecsNil

`func (o *Email) SetTaskCompletionDateTimeSecsNil(b bool)`

 SetTaskCompletionDateTimeSecsNil sets the value for TaskCompletionDateTimeSecs to be an explicit nil

### UnsetTaskCompletionDateTimeSecs
`func (o *Email) UnsetTaskCompletionDateTimeSecs()`

UnsetTaskCompletionDateTimeSecs ensures that no value is present for TaskCompletionDateTimeSecs, not even an explicit nil
### GetTaskDueDateTimeSecs

`func (o *Email) GetTaskDueDateTimeSecs() int64`

GetTaskDueDateTimeSecs returns the TaskDueDateTimeSecs field if non-nil, zero value otherwise.

### GetTaskDueDateTimeSecsOk

`func (o *Email) GetTaskDueDateTimeSecsOk() (*int64, bool)`

GetTaskDueDateTimeSecsOk returns a tuple with the TaskDueDateTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDueDateTimeSecs

`func (o *Email) SetTaskDueDateTimeSecs(v int64)`

SetTaskDueDateTimeSecs sets TaskDueDateTimeSecs field to given value.

### HasTaskDueDateTimeSecs

`func (o *Email) HasTaskDueDateTimeSecs() bool`

HasTaskDueDateTimeSecs returns a boolean if a field has been set.

### SetTaskDueDateTimeSecsNil

`func (o *Email) SetTaskDueDateTimeSecsNil(b bool)`

 SetTaskDueDateTimeSecsNil sets the value for TaskDueDateTimeSecs to be an explicit nil

### UnsetTaskDueDateTimeSecs
`func (o *Email) UnsetTaskDueDateTimeSecs()`

UnsetTaskDueDateTimeSecs ensures that no value is present for TaskDueDateTimeSecs, not even an explicit nil
### GetTaskStatus

`func (o *Email) GetTaskStatus() string`

GetTaskStatus returns the TaskStatus field if non-nil, zero value otherwise.

### GetTaskStatusOk

`func (o *Email) GetTaskStatusOk() (*string, bool)`

GetTaskStatusOk returns a tuple with the TaskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStatus

`func (o *Email) SetTaskStatus(v string)`

SetTaskStatus sets TaskStatus field to given value.

### HasTaskStatus

`func (o *Email) HasTaskStatus() bool`

HasTaskStatus returns a boolean if a field has been set.

### SetTaskStatusNil

`func (o *Email) SetTaskStatusNil(b bool)`

 SetTaskStatusNil sets the value for TaskStatus to be an explicit nil

### UnsetTaskStatus
`func (o *Email) UnsetTaskStatus()`

UnsetTaskStatus ensures that no value is present for TaskStatus, not even an explicit nil
### GetTenantId

`func (o *Email) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Email) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Email) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Email) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *Email) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *Email) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetType

`func (o *Email) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Email) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Email) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Email) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Email) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Email) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUserObjectInfo

`func (o *Email) GetUserObjectInfo() ObjectSummary`

GetUserObjectInfo returns the UserObjectInfo field if non-nil, zero value otherwise.

### GetUserObjectInfoOk

`func (o *Email) GetUserObjectInfoOk() (*ObjectSummary, bool)`

GetUserObjectInfoOk returns a tuple with the UserObjectInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObjectInfo

`func (o *Email) SetUserObjectInfo(v ObjectSummary)`

SetUserObjectInfo sets UserObjectInfo field to given value.

### HasUserObjectInfo

`func (o *Email) HasUserObjectInfo() bool`

HasUserObjectInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


