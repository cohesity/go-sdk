# EmailMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllUnderHierarchy** | Pointer to **NullableBool** | AllUnderHierarchy specifies if logs of all the tenants under the hierarchy of tenant with id TenantId should be returned. | [optional] 
**BccRecipientAddresses** | Pointer to **[]string** | Specifies the email addresses of the bcc recipients. | [optional] 
**CcRecipientAddresses** | Pointer to **[]string** | Specifies the email addresses of the cc recipients. | [optional] 
**DomainIds** | Pointer to **[]int64** | Specifies the domain Ids in which mailboxes are registered. | [optional] 
**EmailSubject** | Pointer to **NullableString** | Specifies the subject of the email. | [optional] 
**FolderKey** | Pointer to **NullableInt64** | Specifes the Parent Folder key. | [optional] 
**FolderName** | Pointer to **NullableString** | Specifies the parent folder name of the email. | [optional] 
**HasAttachments** | Pointer to **NullableBool** | Specifies whether the emails have any attachments. | [optional] 
**ItemKey** | Pointer to **NullableString** | Specifies the Key(unique within mailbox) for Outlook item such as Email. This key is not indexed but used for identifying the item while restore. | [optional] 
**MailboxIds** | Pointer to **[]int64** | Specifies the mailbox Ids which contains the emails/folders. | [optional] 
**ProtectionJobIds** | Pointer to **[]int64** | Specifies the protection job Ids which have backed up mailbox(es) continaing emails/folders. | [optional] 
**ReceivedEndTime** | Pointer to **NullableInt64** | Specifies the unix end time for querying on email&#39;s received time. | [optional] 
**ReceivedStartTime** | Pointer to **NullableInt64** | Specifies the unix start time for querying on email&#39;s received time. | [optional] 
**ReceivedTimeSeconds** | Pointer to **NullableInt64** | Specifies the unix time when the email was received. | [optional] 
**RecipientAddresses** | Pointer to **[]string** | Specifies the email addresses of the recipients. | [optional] 
**SenderAddress** | Pointer to **NullableString** | Specifies the email address of the sender. | [optional] 
**SentTimeSeconds** | Pointer to **NullableInt64** | Specifies the unix time when the email was sent. | [optional] 
**ShowOnlyEmailFolders** | Pointer to **NullableBool** | Specifies whether the query result should include only Email folders. | [optional] 
**TenantId** | Pointer to **NullableString** | TenantId specifies the tenant whose action resulted in the audit log. | [optional] 

## Methods

### NewEmailMetaData

`func NewEmailMetaData() *EmailMetaData`

NewEmailMetaData instantiates a new EmailMetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailMetaDataWithDefaults

`func NewEmailMetaDataWithDefaults() *EmailMetaData`

NewEmailMetaDataWithDefaults instantiates a new EmailMetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllUnderHierarchy

`func (o *EmailMetaData) GetAllUnderHierarchy() bool`

GetAllUnderHierarchy returns the AllUnderHierarchy field if non-nil, zero value otherwise.

### GetAllUnderHierarchyOk

`func (o *EmailMetaData) GetAllUnderHierarchyOk() (*bool, bool)`

GetAllUnderHierarchyOk returns a tuple with the AllUnderHierarchy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllUnderHierarchy

`func (o *EmailMetaData) SetAllUnderHierarchy(v bool)`

SetAllUnderHierarchy sets AllUnderHierarchy field to given value.

### HasAllUnderHierarchy

`func (o *EmailMetaData) HasAllUnderHierarchy() bool`

HasAllUnderHierarchy returns a boolean if a field has been set.

### SetAllUnderHierarchyNil

`func (o *EmailMetaData) SetAllUnderHierarchyNil(b bool)`

 SetAllUnderHierarchyNil sets the value for AllUnderHierarchy to be an explicit nil

### UnsetAllUnderHierarchy
`func (o *EmailMetaData) UnsetAllUnderHierarchy()`

UnsetAllUnderHierarchy ensures that no value is present for AllUnderHierarchy, not even an explicit nil
### GetBccRecipientAddresses

`func (o *EmailMetaData) GetBccRecipientAddresses() []string`

GetBccRecipientAddresses returns the BccRecipientAddresses field if non-nil, zero value otherwise.

### GetBccRecipientAddressesOk

`func (o *EmailMetaData) GetBccRecipientAddressesOk() (*[]string, bool)`

GetBccRecipientAddressesOk returns a tuple with the BccRecipientAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBccRecipientAddresses

`func (o *EmailMetaData) SetBccRecipientAddresses(v []string)`

SetBccRecipientAddresses sets BccRecipientAddresses field to given value.

### HasBccRecipientAddresses

`func (o *EmailMetaData) HasBccRecipientAddresses() bool`

HasBccRecipientAddresses returns a boolean if a field has been set.

### SetBccRecipientAddressesNil

`func (o *EmailMetaData) SetBccRecipientAddressesNil(b bool)`

 SetBccRecipientAddressesNil sets the value for BccRecipientAddresses to be an explicit nil

### UnsetBccRecipientAddresses
`func (o *EmailMetaData) UnsetBccRecipientAddresses()`

UnsetBccRecipientAddresses ensures that no value is present for BccRecipientAddresses, not even an explicit nil
### GetCcRecipientAddresses

`func (o *EmailMetaData) GetCcRecipientAddresses() []string`

GetCcRecipientAddresses returns the CcRecipientAddresses field if non-nil, zero value otherwise.

### GetCcRecipientAddressesOk

`func (o *EmailMetaData) GetCcRecipientAddressesOk() (*[]string, bool)`

GetCcRecipientAddressesOk returns a tuple with the CcRecipientAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcRecipientAddresses

`func (o *EmailMetaData) SetCcRecipientAddresses(v []string)`

SetCcRecipientAddresses sets CcRecipientAddresses field to given value.

### HasCcRecipientAddresses

`func (o *EmailMetaData) HasCcRecipientAddresses() bool`

HasCcRecipientAddresses returns a boolean if a field has been set.

### SetCcRecipientAddressesNil

`func (o *EmailMetaData) SetCcRecipientAddressesNil(b bool)`

 SetCcRecipientAddressesNil sets the value for CcRecipientAddresses to be an explicit nil

### UnsetCcRecipientAddresses
`func (o *EmailMetaData) UnsetCcRecipientAddresses()`

UnsetCcRecipientAddresses ensures that no value is present for CcRecipientAddresses, not even an explicit nil
### GetDomainIds

`func (o *EmailMetaData) GetDomainIds() []int64`

GetDomainIds returns the DomainIds field if non-nil, zero value otherwise.

### GetDomainIdsOk

`func (o *EmailMetaData) GetDomainIdsOk() (*[]int64, bool)`

GetDomainIdsOk returns a tuple with the DomainIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainIds

`func (o *EmailMetaData) SetDomainIds(v []int64)`

SetDomainIds sets DomainIds field to given value.

### HasDomainIds

`func (o *EmailMetaData) HasDomainIds() bool`

HasDomainIds returns a boolean if a field has been set.

### SetDomainIdsNil

`func (o *EmailMetaData) SetDomainIdsNil(b bool)`

 SetDomainIdsNil sets the value for DomainIds to be an explicit nil

### UnsetDomainIds
`func (o *EmailMetaData) UnsetDomainIds()`

UnsetDomainIds ensures that no value is present for DomainIds, not even an explicit nil
### GetEmailSubject

`func (o *EmailMetaData) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *EmailMetaData) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *EmailMetaData) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.

### HasEmailSubject

`func (o *EmailMetaData) HasEmailSubject() bool`

HasEmailSubject returns a boolean if a field has been set.

### SetEmailSubjectNil

`func (o *EmailMetaData) SetEmailSubjectNil(b bool)`

 SetEmailSubjectNil sets the value for EmailSubject to be an explicit nil

### UnsetEmailSubject
`func (o *EmailMetaData) UnsetEmailSubject()`

UnsetEmailSubject ensures that no value is present for EmailSubject, not even an explicit nil
### GetFolderKey

`func (o *EmailMetaData) GetFolderKey() int64`

GetFolderKey returns the FolderKey field if non-nil, zero value otherwise.

### GetFolderKeyOk

`func (o *EmailMetaData) GetFolderKeyOk() (*int64, bool)`

GetFolderKeyOk returns a tuple with the FolderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderKey

`func (o *EmailMetaData) SetFolderKey(v int64)`

SetFolderKey sets FolderKey field to given value.

### HasFolderKey

`func (o *EmailMetaData) HasFolderKey() bool`

HasFolderKey returns a boolean if a field has been set.

### SetFolderKeyNil

`func (o *EmailMetaData) SetFolderKeyNil(b bool)`

 SetFolderKeyNil sets the value for FolderKey to be an explicit nil

### UnsetFolderKey
`func (o *EmailMetaData) UnsetFolderKey()`

UnsetFolderKey ensures that no value is present for FolderKey, not even an explicit nil
### GetFolderName

`func (o *EmailMetaData) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *EmailMetaData) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *EmailMetaData) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.

### HasFolderName

`func (o *EmailMetaData) HasFolderName() bool`

HasFolderName returns a boolean if a field has been set.

### SetFolderNameNil

`func (o *EmailMetaData) SetFolderNameNil(b bool)`

 SetFolderNameNil sets the value for FolderName to be an explicit nil

### UnsetFolderName
`func (o *EmailMetaData) UnsetFolderName()`

UnsetFolderName ensures that no value is present for FolderName, not even an explicit nil
### GetHasAttachments

`func (o *EmailMetaData) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *EmailMetaData) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *EmailMetaData) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachments

`func (o *EmailMetaData) HasHasAttachments() bool`

HasHasAttachments returns a boolean if a field has been set.

### SetHasAttachmentsNil

`func (o *EmailMetaData) SetHasAttachmentsNil(b bool)`

 SetHasAttachmentsNil sets the value for HasAttachments to be an explicit nil

### UnsetHasAttachments
`func (o *EmailMetaData) UnsetHasAttachments()`

UnsetHasAttachments ensures that no value is present for HasAttachments, not even an explicit nil
### GetItemKey

`func (o *EmailMetaData) GetItemKey() string`

GetItemKey returns the ItemKey field if non-nil, zero value otherwise.

### GetItemKeyOk

`func (o *EmailMetaData) GetItemKeyOk() (*string, bool)`

GetItemKeyOk returns a tuple with the ItemKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemKey

`func (o *EmailMetaData) SetItemKey(v string)`

SetItemKey sets ItemKey field to given value.

### HasItemKey

`func (o *EmailMetaData) HasItemKey() bool`

HasItemKey returns a boolean if a field has been set.

### SetItemKeyNil

`func (o *EmailMetaData) SetItemKeyNil(b bool)`

 SetItemKeyNil sets the value for ItemKey to be an explicit nil

### UnsetItemKey
`func (o *EmailMetaData) UnsetItemKey()`

UnsetItemKey ensures that no value is present for ItemKey, not even an explicit nil
### GetMailboxIds

`func (o *EmailMetaData) GetMailboxIds() []int64`

GetMailboxIds returns the MailboxIds field if non-nil, zero value otherwise.

### GetMailboxIdsOk

`func (o *EmailMetaData) GetMailboxIdsOk() (*[]int64, bool)`

GetMailboxIdsOk returns a tuple with the MailboxIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxIds

`func (o *EmailMetaData) SetMailboxIds(v []int64)`

SetMailboxIds sets MailboxIds field to given value.

### HasMailboxIds

`func (o *EmailMetaData) HasMailboxIds() bool`

HasMailboxIds returns a boolean if a field has been set.

### SetMailboxIdsNil

`func (o *EmailMetaData) SetMailboxIdsNil(b bool)`

 SetMailboxIdsNil sets the value for MailboxIds to be an explicit nil

### UnsetMailboxIds
`func (o *EmailMetaData) UnsetMailboxIds()`

UnsetMailboxIds ensures that no value is present for MailboxIds, not even an explicit nil
### GetProtectionJobIds

`func (o *EmailMetaData) GetProtectionJobIds() []int64`

GetProtectionJobIds returns the ProtectionJobIds field if non-nil, zero value otherwise.

### GetProtectionJobIdsOk

`func (o *EmailMetaData) GetProtectionJobIdsOk() (*[]int64, bool)`

GetProtectionJobIdsOk returns a tuple with the ProtectionJobIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionJobIds

`func (o *EmailMetaData) SetProtectionJobIds(v []int64)`

SetProtectionJobIds sets ProtectionJobIds field to given value.

### HasProtectionJobIds

`func (o *EmailMetaData) HasProtectionJobIds() bool`

HasProtectionJobIds returns a boolean if a field has been set.

### SetProtectionJobIdsNil

`func (o *EmailMetaData) SetProtectionJobIdsNil(b bool)`

 SetProtectionJobIdsNil sets the value for ProtectionJobIds to be an explicit nil

### UnsetProtectionJobIds
`func (o *EmailMetaData) UnsetProtectionJobIds()`

UnsetProtectionJobIds ensures that no value is present for ProtectionJobIds, not even an explicit nil
### GetReceivedEndTime

`func (o *EmailMetaData) GetReceivedEndTime() int64`

GetReceivedEndTime returns the ReceivedEndTime field if non-nil, zero value otherwise.

### GetReceivedEndTimeOk

`func (o *EmailMetaData) GetReceivedEndTimeOk() (*int64, bool)`

GetReceivedEndTimeOk returns a tuple with the ReceivedEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedEndTime

`func (o *EmailMetaData) SetReceivedEndTime(v int64)`

SetReceivedEndTime sets ReceivedEndTime field to given value.

### HasReceivedEndTime

`func (o *EmailMetaData) HasReceivedEndTime() bool`

HasReceivedEndTime returns a boolean if a field has been set.

### SetReceivedEndTimeNil

`func (o *EmailMetaData) SetReceivedEndTimeNil(b bool)`

 SetReceivedEndTimeNil sets the value for ReceivedEndTime to be an explicit nil

### UnsetReceivedEndTime
`func (o *EmailMetaData) UnsetReceivedEndTime()`

UnsetReceivedEndTime ensures that no value is present for ReceivedEndTime, not even an explicit nil
### GetReceivedStartTime

`func (o *EmailMetaData) GetReceivedStartTime() int64`

GetReceivedStartTime returns the ReceivedStartTime field if non-nil, zero value otherwise.

### GetReceivedStartTimeOk

`func (o *EmailMetaData) GetReceivedStartTimeOk() (*int64, bool)`

GetReceivedStartTimeOk returns a tuple with the ReceivedStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedStartTime

`func (o *EmailMetaData) SetReceivedStartTime(v int64)`

SetReceivedStartTime sets ReceivedStartTime field to given value.

### HasReceivedStartTime

`func (o *EmailMetaData) HasReceivedStartTime() bool`

HasReceivedStartTime returns a boolean if a field has been set.

### SetReceivedStartTimeNil

`func (o *EmailMetaData) SetReceivedStartTimeNil(b bool)`

 SetReceivedStartTimeNil sets the value for ReceivedStartTime to be an explicit nil

### UnsetReceivedStartTime
`func (o *EmailMetaData) UnsetReceivedStartTime()`

UnsetReceivedStartTime ensures that no value is present for ReceivedStartTime, not even an explicit nil
### GetReceivedTimeSeconds

`func (o *EmailMetaData) GetReceivedTimeSeconds() int64`

GetReceivedTimeSeconds returns the ReceivedTimeSeconds field if non-nil, zero value otherwise.

### GetReceivedTimeSecondsOk

`func (o *EmailMetaData) GetReceivedTimeSecondsOk() (*int64, bool)`

GetReceivedTimeSecondsOk returns a tuple with the ReceivedTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedTimeSeconds

`func (o *EmailMetaData) SetReceivedTimeSeconds(v int64)`

SetReceivedTimeSeconds sets ReceivedTimeSeconds field to given value.

### HasReceivedTimeSeconds

`func (o *EmailMetaData) HasReceivedTimeSeconds() bool`

HasReceivedTimeSeconds returns a boolean if a field has been set.

### SetReceivedTimeSecondsNil

`func (o *EmailMetaData) SetReceivedTimeSecondsNil(b bool)`

 SetReceivedTimeSecondsNil sets the value for ReceivedTimeSeconds to be an explicit nil

### UnsetReceivedTimeSeconds
`func (o *EmailMetaData) UnsetReceivedTimeSeconds()`

UnsetReceivedTimeSeconds ensures that no value is present for ReceivedTimeSeconds, not even an explicit nil
### GetRecipientAddresses

`func (o *EmailMetaData) GetRecipientAddresses() []string`

GetRecipientAddresses returns the RecipientAddresses field if non-nil, zero value otherwise.

### GetRecipientAddressesOk

`func (o *EmailMetaData) GetRecipientAddressesOk() (*[]string, bool)`

GetRecipientAddressesOk returns a tuple with the RecipientAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddresses

`func (o *EmailMetaData) SetRecipientAddresses(v []string)`

SetRecipientAddresses sets RecipientAddresses field to given value.

### HasRecipientAddresses

`func (o *EmailMetaData) HasRecipientAddresses() bool`

HasRecipientAddresses returns a boolean if a field has been set.

### SetRecipientAddressesNil

`func (o *EmailMetaData) SetRecipientAddressesNil(b bool)`

 SetRecipientAddressesNil sets the value for RecipientAddresses to be an explicit nil

### UnsetRecipientAddresses
`func (o *EmailMetaData) UnsetRecipientAddresses()`

UnsetRecipientAddresses ensures that no value is present for RecipientAddresses, not even an explicit nil
### GetSenderAddress

`func (o *EmailMetaData) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *EmailMetaData) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *EmailMetaData) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.

### HasSenderAddress

`func (o *EmailMetaData) HasSenderAddress() bool`

HasSenderAddress returns a boolean if a field has been set.

### SetSenderAddressNil

`func (o *EmailMetaData) SetSenderAddressNil(b bool)`

 SetSenderAddressNil sets the value for SenderAddress to be an explicit nil

### UnsetSenderAddress
`func (o *EmailMetaData) UnsetSenderAddress()`

UnsetSenderAddress ensures that no value is present for SenderAddress, not even an explicit nil
### GetSentTimeSeconds

`func (o *EmailMetaData) GetSentTimeSeconds() int64`

GetSentTimeSeconds returns the SentTimeSeconds field if non-nil, zero value otherwise.

### GetSentTimeSecondsOk

`func (o *EmailMetaData) GetSentTimeSecondsOk() (*int64, bool)`

GetSentTimeSecondsOk returns a tuple with the SentTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentTimeSeconds

`func (o *EmailMetaData) SetSentTimeSeconds(v int64)`

SetSentTimeSeconds sets SentTimeSeconds field to given value.

### HasSentTimeSeconds

`func (o *EmailMetaData) HasSentTimeSeconds() bool`

HasSentTimeSeconds returns a boolean if a field has been set.

### SetSentTimeSecondsNil

`func (o *EmailMetaData) SetSentTimeSecondsNil(b bool)`

 SetSentTimeSecondsNil sets the value for SentTimeSeconds to be an explicit nil

### UnsetSentTimeSeconds
`func (o *EmailMetaData) UnsetSentTimeSeconds()`

UnsetSentTimeSeconds ensures that no value is present for SentTimeSeconds, not even an explicit nil
### GetShowOnlyEmailFolders

`func (o *EmailMetaData) GetShowOnlyEmailFolders() bool`

GetShowOnlyEmailFolders returns the ShowOnlyEmailFolders field if non-nil, zero value otherwise.

### GetShowOnlyEmailFoldersOk

`func (o *EmailMetaData) GetShowOnlyEmailFoldersOk() (*bool, bool)`

GetShowOnlyEmailFoldersOk returns a tuple with the ShowOnlyEmailFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnlyEmailFolders

`func (o *EmailMetaData) SetShowOnlyEmailFolders(v bool)`

SetShowOnlyEmailFolders sets ShowOnlyEmailFolders field to given value.

### HasShowOnlyEmailFolders

`func (o *EmailMetaData) HasShowOnlyEmailFolders() bool`

HasShowOnlyEmailFolders returns a boolean if a field has been set.

### SetShowOnlyEmailFoldersNil

`func (o *EmailMetaData) SetShowOnlyEmailFoldersNil(b bool)`

 SetShowOnlyEmailFoldersNil sets the value for ShowOnlyEmailFolders to be an explicit nil

### UnsetShowOnlyEmailFolders
`func (o *EmailMetaData) UnsetShowOnlyEmailFolders()`

UnsetShowOnlyEmailFolders ensures that no value is present for ShowOnlyEmailFolders, not even an explicit nil
### GetTenantId

`func (o *EmailMetaData) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *EmailMetaData) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *EmailMetaData) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *EmailMetaData) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *EmailMetaData) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *EmailMetaData) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


