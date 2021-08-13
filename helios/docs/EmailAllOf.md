# EmailAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies the id of the email object. | [optional] 
**UserObjectInfo** | Pointer to [**ObjectSummary**](ObjectSummary.md) |  | [optional] 
**FolderName** | Pointer to **NullableString** | Specify the name of the email folder. | [optional] 
**ParentFolderId** | Pointer to **NullableInt64** | Specifies the id of parent folder the mailbox item. | [optional] 
**Path** | Pointer to **NullableString** | Specifies the path to this mailbox item. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the Mailbox item type. | [optional] 
**EmailSubject** | Pointer to **NullableString** | Specifies the subject of this email. | [optional] 
**HasAttachment** | Pointer to **NullableBool** | Specifies whether email has an attachment. | [optional] 
**SenderAddress** | Pointer to **NullableString** | Specifies the sender&#39;s email address. | [optional] 
**RecipientAddresses** | Pointer to **[]string** | \&quot;Specifies the email addresses of all receipients of this email.\&quot; | [optional] 
**CcRecipientAddresses** | Pointer to **[]string** | \&quot;Specifies the email addresses of all the CC receipients of this email.\&quot; | [optional] 
**BccRecipientAddresses** | Pointer to **[]string** | \&quot;Specifies the email addresses of all the BCC receipients of this email.\&quot; | [optional] 
**SentTimeSecs** | Pointer to **NullableInt64** | \&quot;Specifies the Unix timestamp epoch in seconds at which this email is sent.\&quot; | [optional] 
**ReceivedTimeSecs** | Pointer to **NullableInt64** | \&quot;Specifies the Unix timestamp epoch in seconds at which this email is received.\&quot; | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | \&quot;Specifies the Protection Group id protecting the mailbox.\&quot; | [optional] 
**StorageDomainId** | Pointer to **NullableInt64** | \&quot;Specifies the Storage Domain id where the backup data of Object is present.\&quot; | [optional] 
**TenantId** | Pointer to **NullableString** | \&quot;Specify the tenant id to which this email belongs to.\&quot; | [optional] 

## Methods

### NewEmailAllOf

`func NewEmailAllOf() *EmailAllOf`

NewEmailAllOf instantiates a new EmailAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailAllOfWithDefaults

`func NewEmailAllOfWithDefaults() *EmailAllOf`

NewEmailAllOfWithDefaults instantiates a new EmailAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmailAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *EmailAllOf) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *EmailAllOf) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUserObjectInfo

`func (o *EmailAllOf) GetUserObjectInfo() ObjectSummary`

GetUserObjectInfo returns the UserObjectInfo field if non-nil, zero value otherwise.

### GetUserObjectInfoOk

`func (o *EmailAllOf) GetUserObjectInfoOk() (*ObjectSummary, bool)`

GetUserObjectInfoOk returns a tuple with the UserObjectInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObjectInfo

`func (o *EmailAllOf) SetUserObjectInfo(v ObjectSummary)`

SetUserObjectInfo sets UserObjectInfo field to given value.

### HasUserObjectInfo

`func (o *EmailAllOf) HasUserObjectInfo() bool`

HasUserObjectInfo returns a boolean if a field has been set.

### GetFolderName

`func (o *EmailAllOf) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *EmailAllOf) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *EmailAllOf) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.

### HasFolderName

`func (o *EmailAllOf) HasFolderName() bool`

HasFolderName returns a boolean if a field has been set.

### SetFolderNameNil

`func (o *EmailAllOf) SetFolderNameNil(b bool)`

 SetFolderNameNil sets the value for FolderName to be an explicit nil

### UnsetFolderName
`func (o *EmailAllOf) UnsetFolderName()`

UnsetFolderName ensures that no value is present for FolderName, not even an explicit nil
### GetParentFolderId

`func (o *EmailAllOf) GetParentFolderId() int64`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *EmailAllOf) GetParentFolderIdOk() (*int64, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *EmailAllOf) SetParentFolderId(v int64)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *EmailAllOf) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderIdNil

`func (o *EmailAllOf) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *EmailAllOf) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
### GetPath

`func (o *EmailAllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *EmailAllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *EmailAllOf) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *EmailAllOf) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *EmailAllOf) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *EmailAllOf) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetType

`func (o *EmailAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EmailAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EmailAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EmailAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *EmailAllOf) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *EmailAllOf) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetEmailSubject

`func (o *EmailAllOf) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *EmailAllOf) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *EmailAllOf) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.

### HasEmailSubject

`func (o *EmailAllOf) HasEmailSubject() bool`

HasEmailSubject returns a boolean if a field has been set.

### SetEmailSubjectNil

`func (o *EmailAllOf) SetEmailSubjectNil(b bool)`

 SetEmailSubjectNil sets the value for EmailSubject to be an explicit nil

### UnsetEmailSubject
`func (o *EmailAllOf) UnsetEmailSubject()`

UnsetEmailSubject ensures that no value is present for EmailSubject, not even an explicit nil
### GetHasAttachment

`func (o *EmailAllOf) GetHasAttachment() bool`

GetHasAttachment returns the HasAttachment field if non-nil, zero value otherwise.

### GetHasAttachmentOk

`func (o *EmailAllOf) GetHasAttachmentOk() (*bool, bool)`

GetHasAttachmentOk returns a tuple with the HasAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachment

`func (o *EmailAllOf) SetHasAttachment(v bool)`

SetHasAttachment sets HasAttachment field to given value.

### HasHasAttachment

`func (o *EmailAllOf) HasHasAttachment() bool`

HasHasAttachment returns a boolean if a field has been set.

### SetHasAttachmentNil

`func (o *EmailAllOf) SetHasAttachmentNil(b bool)`

 SetHasAttachmentNil sets the value for HasAttachment to be an explicit nil

### UnsetHasAttachment
`func (o *EmailAllOf) UnsetHasAttachment()`

UnsetHasAttachment ensures that no value is present for HasAttachment, not even an explicit nil
### GetSenderAddress

`func (o *EmailAllOf) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *EmailAllOf) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *EmailAllOf) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.

### HasSenderAddress

`func (o *EmailAllOf) HasSenderAddress() bool`

HasSenderAddress returns a boolean if a field has been set.

### SetSenderAddressNil

`func (o *EmailAllOf) SetSenderAddressNil(b bool)`

 SetSenderAddressNil sets the value for SenderAddress to be an explicit nil

### UnsetSenderAddress
`func (o *EmailAllOf) UnsetSenderAddress()`

UnsetSenderAddress ensures that no value is present for SenderAddress, not even an explicit nil
### GetRecipientAddresses

`func (o *EmailAllOf) GetRecipientAddresses() []string`

GetRecipientAddresses returns the RecipientAddresses field if non-nil, zero value otherwise.

### GetRecipientAddressesOk

`func (o *EmailAllOf) GetRecipientAddressesOk() (*[]string, bool)`

GetRecipientAddressesOk returns a tuple with the RecipientAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddresses

`func (o *EmailAllOf) SetRecipientAddresses(v []string)`

SetRecipientAddresses sets RecipientAddresses field to given value.

### HasRecipientAddresses

`func (o *EmailAllOf) HasRecipientAddresses() bool`

HasRecipientAddresses returns a boolean if a field has been set.

### SetRecipientAddressesNil

`func (o *EmailAllOf) SetRecipientAddressesNil(b bool)`

 SetRecipientAddressesNil sets the value for RecipientAddresses to be an explicit nil

### UnsetRecipientAddresses
`func (o *EmailAllOf) UnsetRecipientAddresses()`

UnsetRecipientAddresses ensures that no value is present for RecipientAddresses, not even an explicit nil
### GetCcRecipientAddresses

`func (o *EmailAllOf) GetCcRecipientAddresses() []string`

GetCcRecipientAddresses returns the CcRecipientAddresses field if non-nil, zero value otherwise.

### GetCcRecipientAddressesOk

`func (o *EmailAllOf) GetCcRecipientAddressesOk() (*[]string, bool)`

GetCcRecipientAddressesOk returns a tuple with the CcRecipientAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcRecipientAddresses

`func (o *EmailAllOf) SetCcRecipientAddresses(v []string)`

SetCcRecipientAddresses sets CcRecipientAddresses field to given value.

### HasCcRecipientAddresses

`func (o *EmailAllOf) HasCcRecipientAddresses() bool`

HasCcRecipientAddresses returns a boolean if a field has been set.

### SetCcRecipientAddressesNil

`func (o *EmailAllOf) SetCcRecipientAddressesNil(b bool)`

 SetCcRecipientAddressesNil sets the value for CcRecipientAddresses to be an explicit nil

### UnsetCcRecipientAddresses
`func (o *EmailAllOf) UnsetCcRecipientAddresses()`

UnsetCcRecipientAddresses ensures that no value is present for CcRecipientAddresses, not even an explicit nil
### GetBccRecipientAddresses

`func (o *EmailAllOf) GetBccRecipientAddresses() []string`

GetBccRecipientAddresses returns the BccRecipientAddresses field if non-nil, zero value otherwise.

### GetBccRecipientAddressesOk

`func (o *EmailAllOf) GetBccRecipientAddressesOk() (*[]string, bool)`

GetBccRecipientAddressesOk returns a tuple with the BccRecipientAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBccRecipientAddresses

`func (o *EmailAllOf) SetBccRecipientAddresses(v []string)`

SetBccRecipientAddresses sets BccRecipientAddresses field to given value.

### HasBccRecipientAddresses

`func (o *EmailAllOf) HasBccRecipientAddresses() bool`

HasBccRecipientAddresses returns a boolean if a field has been set.

### SetBccRecipientAddressesNil

`func (o *EmailAllOf) SetBccRecipientAddressesNil(b bool)`

 SetBccRecipientAddressesNil sets the value for BccRecipientAddresses to be an explicit nil

### UnsetBccRecipientAddresses
`func (o *EmailAllOf) UnsetBccRecipientAddresses()`

UnsetBccRecipientAddresses ensures that no value is present for BccRecipientAddresses, not even an explicit nil
### GetSentTimeSecs

`func (o *EmailAllOf) GetSentTimeSecs() int64`

GetSentTimeSecs returns the SentTimeSecs field if non-nil, zero value otherwise.

### GetSentTimeSecsOk

`func (o *EmailAllOf) GetSentTimeSecsOk() (*int64, bool)`

GetSentTimeSecsOk returns a tuple with the SentTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentTimeSecs

`func (o *EmailAllOf) SetSentTimeSecs(v int64)`

SetSentTimeSecs sets SentTimeSecs field to given value.

### HasSentTimeSecs

`func (o *EmailAllOf) HasSentTimeSecs() bool`

HasSentTimeSecs returns a boolean if a field has been set.

### SetSentTimeSecsNil

`func (o *EmailAllOf) SetSentTimeSecsNil(b bool)`

 SetSentTimeSecsNil sets the value for SentTimeSecs to be an explicit nil

### UnsetSentTimeSecs
`func (o *EmailAllOf) UnsetSentTimeSecs()`

UnsetSentTimeSecs ensures that no value is present for SentTimeSecs, not even an explicit nil
### GetReceivedTimeSecs

`func (o *EmailAllOf) GetReceivedTimeSecs() int64`

GetReceivedTimeSecs returns the ReceivedTimeSecs field if non-nil, zero value otherwise.

### GetReceivedTimeSecsOk

`func (o *EmailAllOf) GetReceivedTimeSecsOk() (*int64, bool)`

GetReceivedTimeSecsOk returns a tuple with the ReceivedTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedTimeSecs

`func (o *EmailAllOf) SetReceivedTimeSecs(v int64)`

SetReceivedTimeSecs sets ReceivedTimeSecs field to given value.

### HasReceivedTimeSecs

`func (o *EmailAllOf) HasReceivedTimeSecs() bool`

HasReceivedTimeSecs returns a boolean if a field has been set.

### SetReceivedTimeSecsNil

`func (o *EmailAllOf) SetReceivedTimeSecsNil(b bool)`

 SetReceivedTimeSecsNil sets the value for ReceivedTimeSecs to be an explicit nil

### UnsetReceivedTimeSecs
`func (o *EmailAllOf) UnsetReceivedTimeSecs()`

UnsetReceivedTimeSecs ensures that no value is present for ReceivedTimeSecs, not even an explicit nil
### GetProtectionGroupId

`func (o *EmailAllOf) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *EmailAllOf) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *EmailAllOf) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *EmailAllOf) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *EmailAllOf) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *EmailAllOf) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetStorageDomainId

`func (o *EmailAllOf) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *EmailAllOf) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *EmailAllOf) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *EmailAllOf) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *EmailAllOf) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *EmailAllOf) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetTenantId

`func (o *EmailAllOf) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *EmailAllOf) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *EmailAllOf) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *EmailAllOf) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *EmailAllOf) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *EmailAllOf) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


