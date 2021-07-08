# Vault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaTrustedCertificate** | Pointer to **NullableString** | Specifies the CA (certificate authority) trusted certificate. | [optional] 
**ClientCertificate** | Pointer to **NullableString** | Specifies the client CA  certificate. This certificate is in pem format. | [optional] 
**ClientPrivateKey** | Pointer to **NullableString** | Specifies the client private key. This certificate is in pem format. | [optional] 
**CompressionPolicy** | Pointer to **NullableString** | Specifies whether to send data to the Vault in a compressed format. &#39;kCompressionNone&#39; indicates that data is not compressed. &#39;kCompressionLow&#39; indicates that data is compressed using LZ4 or Snappy. &#39;kCompressionHigh&#39; indicates that data is compressed in Gzip. | [optional] 
**Config** | Pointer to [**VaultConfig**](VaultConfig.md) |  | [optional] 
**CustomerManagingEncryptionKeys** | Pointer to **NullableBool** | Specifies whether to manage the encryption key manually or let the Cohesity Cluster manage it. If true, you must get the encryption key store it outside the Cluster, before disaster strikes such as the source local Cohesity Cluster being down. You can get the encryption key by downloading it using the Cohesity Dashboard or by calling the GET /public/vaults/encryptionKey/{id} operation. | [optional] 
**DedupEnabled** | Pointer to **NullableBool** | Specifies whether to deduplicate data before sending it to the Vault. | [optional] 
**DeleteVaultError** | Pointer to **NullableString** | Specifies the error message when deleting a vault. | [optional] 
**Description** | Pointer to **NullableString** | Specifies a description about the Vault. | [optional] 
**DesiredWalLocation** | Pointer to **NullableString** | Desired location for write ahead logs(wal). &#39;kHomePartition&#39; indicates desired wal location to be the home partition. &#39;kDisk&#39; indicates desired wal location to be the same disk as chunk repo. &#39;kScribe&#39; indicates desired wal location to be scribe. &#39;kScribeTable&#39; indicates chunk repository state is kept as key-value pairs in scribe. | [optional] 
**EncryptionKeyFileDownloaded** | Pointer to **NullableBool** | Specifies if the encryption key file has been downloaded using the Cohesity Dashboard (Cohesity UI). If true, the encryption key has been downloaded using the Cohesity Dashboard. An encryption key can only be downloaded once using the Cohesity Dashboard. | [optional] 
**EncryptionPolicy** | Pointer to **NullableString** | Specifies whether to send and store data in an encrypted format. &#39;kEncryptionNone&#39; indicates the data is not encrypted. &#39;kEncryptionStrong&#39; indicates the data is encrypted. | [optional] 
**ExternalTargetType** | Pointer to **NullableString** | Specifies the type of Vault. &#39;kNearline&#39; indicates a Google Nearline Vault. &#39;kGlacier&#39; indicates an AWS Glacier Vault. &#39;kS3&#39; indicates an AWS S3 Vault. &#39;kAzureStandard&#39; indicates a Microsoft Azure Standard Vault. &#39;kS3Compatible&#39; indicates an S3 Compatible Vault. (See the online help for supported types.) &#39;kQStarTape&#39; indicates a QStar Tape Vault. &#39;kGoogleStandard&#39; indicates a Google Standard Vault. &#39;kGoogleDRA&#39; indicates a Google DRA Vault. &#39;kAmazonS3StandardIA&#39; indicates an Amazon S3 Standard-IA Vault. &#39;kAWSGovCloud&#39; indicates an AWS Gov Cloud Vault. &#39;kNAS&#39; indicates a NAS Vault. &#39;kColdline&#39; indicates a Google Coldline Vault. &#39;kAzureGovCloud&#39; indicates a Microsoft Azure Gov Cloud Vault. &#39;kAzureArchive&#39; indicates an Azure Archive Vault. &#39;kAzure&#39; indicates an Azure Vault. &#39;kGoogle&#39; indicates a Google Vault. &#39;kAmazon&#39; indicates an Amazon Vault. &#39;kOracle&#39; indicates an Oracle Vault. &#39;kOracleTierStandard&#39; indicates an Oracle Tier Standard Vault. &#39;kOracleTierArchive&#39; indicates an Oracle Tier Archive Vault. &#39;kAmazonC2S&#39; indicates an Amazon Commercial Cloud Services Vault. | [optional] 
**FullArchiveIntervalDays** | Pointer to **NullableInt64** | Specifies the number days between full archives to the Vault. The current default is 90 days. This is only meaningful when incrementalArchivesEnabled is true and the Vault usage type is kArchival. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies an id that identifies the Vault. | [optional] 
**IncrementalArchivesEnabled** | Pointer to **NullableBool** | Specifies whether to perform incremental archival when sending data to the Vault. If false, only full backups are performed. If true, incremental backups are performed between the full backups. | [optional] 
**IsPasswordEncrypted** | Pointer to **NullableBool** | Specifies if given password is not encrypted or not in the cluster config. | [optional] 
**KeyFileDownloadTimeUsecs** | Pointer to **NullableInt64** | Specifies the time (in microseconds) when the encryption key file was downloaded from the Cohesity Dashboard (Cohesity UI). An encryption key can only be downloaded once using the Cohesity Dashboard. | [optional] 
**KeyFileDownloadUser** | Pointer to **NullableString** | Specifies the user who downloaded the encryption key from the Cohesity Dashboard (Cohesity UI). This field is only populated if encryption is enabled for the Vault and customerManagingEncryptionKeys is true. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the Vault. | [optional] 
**RemovalState** | Pointer to **NullableString** | Specifies the state of the vault to be removed. &#39;kDontRemove&#39; means the state of object is functional and it is not being removed. &#39;kMarkedForRemoval&#39; means the object is being removed. &#39;kOkToRemove&#39; means the object has been removed on the Cohesity Cluster and if the object is physical, it can be removed from the Cohesity Cluster. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of Vault. This field is deprecated. This field is split into ExternalTargetType in and TierType in respective credentials. Initialize those fields instead. deprecated: true &#39;kNearline&#39; indicates a Google Nearline Vault. &#39;kGlacier&#39; indicates an AWS Glacier Vault. &#39;kS3&#39; indicates an AWS S3 Vault. &#39;kAzureStandard&#39; indicates a Microsoft Azure Standard Vault. &#39;kS3Compatible&#39; indicates an S3 Compatible Vault. (See the online help for supported types.) &#39;kQStarTape&#39; indicates a QStar Tape Vault. &#39;kGoogleStandard&#39; indicates a Google Standard Vault. &#39;kGoogleDRA&#39; indicates a Google DRA Vault. &#39;kAmazonS3StandardIA&#39; indicates an Amazon S3 Standard-IA Vault. &#39;kAWSGovCloud&#39; indicates an AWS Gov Cloud Vault. &#39;kNAS&#39; indicates a NAS Vault. &#39;kColdline&#39; indicates a Google Coldline Vault. &#39;kAzureGovCloud&#39; indicates a Microsoft Azure Gov Cloud Vault. &#39;kAzureArchive&#39; indicates an Azure Archive Vault. &#39;kAzure&#39; indicates an Azure Vault. &#39;kGoogle&#39; indicates a Google Vault. &#39;kAmazon&#39; indicates an Amazon Vault. &#39;kOracle&#39; indicates an Oracle Vault. &#39;kOracleTierStandard&#39; indicates an Oracle Tier Standard Vault. &#39;kOracleTierArchive&#39; indicates an Oracle Tier Archive Vault. &#39;kAmazonC2S&#39; indicates an Amazon Commercial Cloud Services Vault. | [optional] 
**UsageType** | Pointer to **NullableString** | Specifies the usage type of the Vault. &#39;kArchival&#39; indicates the Vault provides archive storage for backup data. &#39;kCloudSpill&#39; indicates the Vault provides additional storage for cold data. | [optional] 
**VaultBandwidthLimits** | Pointer to [**VaultBandwidthLimits**](VaultBandwidthLimits.md) |  | [optional] 

## Methods

### NewVault

`func NewVault() *Vault`

NewVault instantiates a new Vault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultWithDefaults

`func NewVaultWithDefaults() *Vault`

NewVaultWithDefaults instantiates a new Vault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaTrustedCertificate

`func (o *Vault) GetCaTrustedCertificate() string`

GetCaTrustedCertificate returns the CaTrustedCertificate field if non-nil, zero value otherwise.

### GetCaTrustedCertificateOk

`func (o *Vault) GetCaTrustedCertificateOk() (*string, bool)`

GetCaTrustedCertificateOk returns a tuple with the CaTrustedCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaTrustedCertificate

`func (o *Vault) SetCaTrustedCertificate(v string)`

SetCaTrustedCertificate sets CaTrustedCertificate field to given value.

### HasCaTrustedCertificate

`func (o *Vault) HasCaTrustedCertificate() bool`

HasCaTrustedCertificate returns a boolean if a field has been set.

### SetCaTrustedCertificateNil

`func (o *Vault) SetCaTrustedCertificateNil(b bool)`

 SetCaTrustedCertificateNil sets the value for CaTrustedCertificate to be an explicit nil

### UnsetCaTrustedCertificate
`func (o *Vault) UnsetCaTrustedCertificate()`

UnsetCaTrustedCertificate ensures that no value is present for CaTrustedCertificate, not even an explicit nil
### GetClientCertificate

`func (o *Vault) GetClientCertificate() string`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *Vault) GetClientCertificateOk() (*string, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *Vault) SetClientCertificate(v string)`

SetClientCertificate sets ClientCertificate field to given value.

### HasClientCertificate

`func (o *Vault) HasClientCertificate() bool`

HasClientCertificate returns a boolean if a field has been set.

### SetClientCertificateNil

`func (o *Vault) SetClientCertificateNil(b bool)`

 SetClientCertificateNil sets the value for ClientCertificate to be an explicit nil

### UnsetClientCertificate
`func (o *Vault) UnsetClientCertificate()`

UnsetClientCertificate ensures that no value is present for ClientCertificate, not even an explicit nil
### GetClientPrivateKey

`func (o *Vault) GetClientPrivateKey() string`

GetClientPrivateKey returns the ClientPrivateKey field if non-nil, zero value otherwise.

### GetClientPrivateKeyOk

`func (o *Vault) GetClientPrivateKeyOk() (*string, bool)`

GetClientPrivateKeyOk returns a tuple with the ClientPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrivateKey

`func (o *Vault) SetClientPrivateKey(v string)`

SetClientPrivateKey sets ClientPrivateKey field to given value.

### HasClientPrivateKey

`func (o *Vault) HasClientPrivateKey() bool`

HasClientPrivateKey returns a boolean if a field has been set.

### SetClientPrivateKeyNil

`func (o *Vault) SetClientPrivateKeyNil(b bool)`

 SetClientPrivateKeyNil sets the value for ClientPrivateKey to be an explicit nil

### UnsetClientPrivateKey
`func (o *Vault) UnsetClientPrivateKey()`

UnsetClientPrivateKey ensures that no value is present for ClientPrivateKey, not even an explicit nil
### GetCompressionPolicy

`func (o *Vault) GetCompressionPolicy() string`

GetCompressionPolicy returns the CompressionPolicy field if non-nil, zero value otherwise.

### GetCompressionPolicyOk

`func (o *Vault) GetCompressionPolicyOk() (*string, bool)`

GetCompressionPolicyOk returns a tuple with the CompressionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionPolicy

`func (o *Vault) SetCompressionPolicy(v string)`

SetCompressionPolicy sets CompressionPolicy field to given value.

### HasCompressionPolicy

`func (o *Vault) HasCompressionPolicy() bool`

HasCompressionPolicy returns a boolean if a field has been set.

### SetCompressionPolicyNil

`func (o *Vault) SetCompressionPolicyNil(b bool)`

 SetCompressionPolicyNil sets the value for CompressionPolicy to be an explicit nil

### UnsetCompressionPolicy
`func (o *Vault) UnsetCompressionPolicy()`

UnsetCompressionPolicy ensures that no value is present for CompressionPolicy, not even an explicit nil
### GetConfig

`func (o *Vault) GetConfig() VaultConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Vault) GetConfigOk() (*VaultConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Vault) SetConfig(v VaultConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Vault) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCustomerManagingEncryptionKeys

`func (o *Vault) GetCustomerManagingEncryptionKeys() bool`

GetCustomerManagingEncryptionKeys returns the CustomerManagingEncryptionKeys field if non-nil, zero value otherwise.

### GetCustomerManagingEncryptionKeysOk

`func (o *Vault) GetCustomerManagingEncryptionKeysOk() (*bool, bool)`

GetCustomerManagingEncryptionKeysOk returns a tuple with the CustomerManagingEncryptionKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerManagingEncryptionKeys

`func (o *Vault) SetCustomerManagingEncryptionKeys(v bool)`

SetCustomerManagingEncryptionKeys sets CustomerManagingEncryptionKeys field to given value.

### HasCustomerManagingEncryptionKeys

`func (o *Vault) HasCustomerManagingEncryptionKeys() bool`

HasCustomerManagingEncryptionKeys returns a boolean if a field has been set.

### SetCustomerManagingEncryptionKeysNil

`func (o *Vault) SetCustomerManagingEncryptionKeysNil(b bool)`

 SetCustomerManagingEncryptionKeysNil sets the value for CustomerManagingEncryptionKeys to be an explicit nil

### UnsetCustomerManagingEncryptionKeys
`func (o *Vault) UnsetCustomerManagingEncryptionKeys()`

UnsetCustomerManagingEncryptionKeys ensures that no value is present for CustomerManagingEncryptionKeys, not even an explicit nil
### GetDedupEnabled

`func (o *Vault) GetDedupEnabled() bool`

GetDedupEnabled returns the DedupEnabled field if non-nil, zero value otherwise.

### GetDedupEnabledOk

`func (o *Vault) GetDedupEnabledOk() (*bool, bool)`

GetDedupEnabledOk returns a tuple with the DedupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupEnabled

`func (o *Vault) SetDedupEnabled(v bool)`

SetDedupEnabled sets DedupEnabled field to given value.

### HasDedupEnabled

`func (o *Vault) HasDedupEnabled() bool`

HasDedupEnabled returns a boolean if a field has been set.

### SetDedupEnabledNil

`func (o *Vault) SetDedupEnabledNil(b bool)`

 SetDedupEnabledNil sets the value for DedupEnabled to be an explicit nil

### UnsetDedupEnabled
`func (o *Vault) UnsetDedupEnabled()`

UnsetDedupEnabled ensures that no value is present for DedupEnabled, not even an explicit nil
### GetDeleteVaultError

`func (o *Vault) GetDeleteVaultError() string`

GetDeleteVaultError returns the DeleteVaultError field if non-nil, zero value otherwise.

### GetDeleteVaultErrorOk

`func (o *Vault) GetDeleteVaultErrorOk() (*string, bool)`

GetDeleteVaultErrorOk returns a tuple with the DeleteVaultError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteVaultError

`func (o *Vault) SetDeleteVaultError(v string)`

SetDeleteVaultError sets DeleteVaultError field to given value.

### HasDeleteVaultError

`func (o *Vault) HasDeleteVaultError() bool`

HasDeleteVaultError returns a boolean if a field has been set.

### SetDeleteVaultErrorNil

`func (o *Vault) SetDeleteVaultErrorNil(b bool)`

 SetDeleteVaultErrorNil sets the value for DeleteVaultError to be an explicit nil

### UnsetDeleteVaultError
`func (o *Vault) UnsetDeleteVaultError()`

UnsetDeleteVaultError ensures that no value is present for DeleteVaultError, not even an explicit nil
### GetDescription

`func (o *Vault) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Vault) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Vault) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Vault) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Vault) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Vault) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDesiredWalLocation

`func (o *Vault) GetDesiredWalLocation() string`

GetDesiredWalLocation returns the DesiredWalLocation field if non-nil, zero value otherwise.

### GetDesiredWalLocationOk

`func (o *Vault) GetDesiredWalLocationOk() (*string, bool)`

GetDesiredWalLocationOk returns a tuple with the DesiredWalLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredWalLocation

`func (o *Vault) SetDesiredWalLocation(v string)`

SetDesiredWalLocation sets DesiredWalLocation field to given value.

### HasDesiredWalLocation

`func (o *Vault) HasDesiredWalLocation() bool`

HasDesiredWalLocation returns a boolean if a field has been set.

### SetDesiredWalLocationNil

`func (o *Vault) SetDesiredWalLocationNil(b bool)`

 SetDesiredWalLocationNil sets the value for DesiredWalLocation to be an explicit nil

### UnsetDesiredWalLocation
`func (o *Vault) UnsetDesiredWalLocation()`

UnsetDesiredWalLocation ensures that no value is present for DesiredWalLocation, not even an explicit nil
### GetEncryptionKeyFileDownloaded

`func (o *Vault) GetEncryptionKeyFileDownloaded() bool`

GetEncryptionKeyFileDownloaded returns the EncryptionKeyFileDownloaded field if non-nil, zero value otherwise.

### GetEncryptionKeyFileDownloadedOk

`func (o *Vault) GetEncryptionKeyFileDownloadedOk() (*bool, bool)`

GetEncryptionKeyFileDownloadedOk returns a tuple with the EncryptionKeyFileDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyFileDownloaded

`func (o *Vault) SetEncryptionKeyFileDownloaded(v bool)`

SetEncryptionKeyFileDownloaded sets EncryptionKeyFileDownloaded field to given value.

### HasEncryptionKeyFileDownloaded

`func (o *Vault) HasEncryptionKeyFileDownloaded() bool`

HasEncryptionKeyFileDownloaded returns a boolean if a field has been set.

### SetEncryptionKeyFileDownloadedNil

`func (o *Vault) SetEncryptionKeyFileDownloadedNil(b bool)`

 SetEncryptionKeyFileDownloadedNil sets the value for EncryptionKeyFileDownloaded to be an explicit nil

### UnsetEncryptionKeyFileDownloaded
`func (o *Vault) UnsetEncryptionKeyFileDownloaded()`

UnsetEncryptionKeyFileDownloaded ensures that no value is present for EncryptionKeyFileDownloaded, not even an explicit nil
### GetEncryptionPolicy

`func (o *Vault) GetEncryptionPolicy() string`

GetEncryptionPolicy returns the EncryptionPolicy field if non-nil, zero value otherwise.

### GetEncryptionPolicyOk

`func (o *Vault) GetEncryptionPolicyOk() (*string, bool)`

GetEncryptionPolicyOk returns a tuple with the EncryptionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionPolicy

`func (o *Vault) SetEncryptionPolicy(v string)`

SetEncryptionPolicy sets EncryptionPolicy field to given value.

### HasEncryptionPolicy

`func (o *Vault) HasEncryptionPolicy() bool`

HasEncryptionPolicy returns a boolean if a field has been set.

### SetEncryptionPolicyNil

`func (o *Vault) SetEncryptionPolicyNil(b bool)`

 SetEncryptionPolicyNil sets the value for EncryptionPolicy to be an explicit nil

### UnsetEncryptionPolicy
`func (o *Vault) UnsetEncryptionPolicy()`

UnsetEncryptionPolicy ensures that no value is present for EncryptionPolicy, not even an explicit nil
### GetExternalTargetType

`func (o *Vault) GetExternalTargetType() string`

GetExternalTargetType returns the ExternalTargetType field if non-nil, zero value otherwise.

### GetExternalTargetTypeOk

`func (o *Vault) GetExternalTargetTypeOk() (*string, bool)`

GetExternalTargetTypeOk returns a tuple with the ExternalTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTargetType

`func (o *Vault) SetExternalTargetType(v string)`

SetExternalTargetType sets ExternalTargetType field to given value.

### HasExternalTargetType

`func (o *Vault) HasExternalTargetType() bool`

HasExternalTargetType returns a boolean if a field has been set.

### SetExternalTargetTypeNil

`func (o *Vault) SetExternalTargetTypeNil(b bool)`

 SetExternalTargetTypeNil sets the value for ExternalTargetType to be an explicit nil

### UnsetExternalTargetType
`func (o *Vault) UnsetExternalTargetType()`

UnsetExternalTargetType ensures that no value is present for ExternalTargetType, not even an explicit nil
### GetFullArchiveIntervalDays

`func (o *Vault) GetFullArchiveIntervalDays() int64`

GetFullArchiveIntervalDays returns the FullArchiveIntervalDays field if non-nil, zero value otherwise.

### GetFullArchiveIntervalDaysOk

`func (o *Vault) GetFullArchiveIntervalDaysOk() (*int64, bool)`

GetFullArchiveIntervalDaysOk returns a tuple with the FullArchiveIntervalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullArchiveIntervalDays

`func (o *Vault) SetFullArchiveIntervalDays(v int64)`

SetFullArchiveIntervalDays sets FullArchiveIntervalDays field to given value.

### HasFullArchiveIntervalDays

`func (o *Vault) HasFullArchiveIntervalDays() bool`

HasFullArchiveIntervalDays returns a boolean if a field has been set.

### SetFullArchiveIntervalDaysNil

`func (o *Vault) SetFullArchiveIntervalDaysNil(b bool)`

 SetFullArchiveIntervalDaysNil sets the value for FullArchiveIntervalDays to be an explicit nil

### UnsetFullArchiveIntervalDays
`func (o *Vault) UnsetFullArchiveIntervalDays()`

UnsetFullArchiveIntervalDays ensures that no value is present for FullArchiveIntervalDays, not even an explicit nil
### GetId

`func (o *Vault) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Vault) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Vault) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Vault) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Vault) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Vault) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIncrementalArchivesEnabled

`func (o *Vault) GetIncrementalArchivesEnabled() bool`

GetIncrementalArchivesEnabled returns the IncrementalArchivesEnabled field if non-nil, zero value otherwise.

### GetIncrementalArchivesEnabledOk

`func (o *Vault) GetIncrementalArchivesEnabledOk() (*bool, bool)`

GetIncrementalArchivesEnabledOk returns a tuple with the IncrementalArchivesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalArchivesEnabled

`func (o *Vault) SetIncrementalArchivesEnabled(v bool)`

SetIncrementalArchivesEnabled sets IncrementalArchivesEnabled field to given value.

### HasIncrementalArchivesEnabled

`func (o *Vault) HasIncrementalArchivesEnabled() bool`

HasIncrementalArchivesEnabled returns a boolean if a field has been set.

### SetIncrementalArchivesEnabledNil

`func (o *Vault) SetIncrementalArchivesEnabledNil(b bool)`

 SetIncrementalArchivesEnabledNil sets the value for IncrementalArchivesEnabled to be an explicit nil

### UnsetIncrementalArchivesEnabled
`func (o *Vault) UnsetIncrementalArchivesEnabled()`

UnsetIncrementalArchivesEnabled ensures that no value is present for IncrementalArchivesEnabled, not even an explicit nil
### GetIsPasswordEncrypted

`func (o *Vault) GetIsPasswordEncrypted() bool`

GetIsPasswordEncrypted returns the IsPasswordEncrypted field if non-nil, zero value otherwise.

### GetIsPasswordEncryptedOk

`func (o *Vault) GetIsPasswordEncryptedOk() (*bool, bool)`

GetIsPasswordEncryptedOk returns a tuple with the IsPasswordEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordEncrypted

`func (o *Vault) SetIsPasswordEncrypted(v bool)`

SetIsPasswordEncrypted sets IsPasswordEncrypted field to given value.

### HasIsPasswordEncrypted

`func (o *Vault) HasIsPasswordEncrypted() bool`

HasIsPasswordEncrypted returns a boolean if a field has been set.

### SetIsPasswordEncryptedNil

`func (o *Vault) SetIsPasswordEncryptedNil(b bool)`

 SetIsPasswordEncryptedNil sets the value for IsPasswordEncrypted to be an explicit nil

### UnsetIsPasswordEncrypted
`func (o *Vault) UnsetIsPasswordEncrypted()`

UnsetIsPasswordEncrypted ensures that no value is present for IsPasswordEncrypted, not even an explicit nil
### GetKeyFileDownloadTimeUsecs

`func (o *Vault) GetKeyFileDownloadTimeUsecs() int64`

GetKeyFileDownloadTimeUsecs returns the KeyFileDownloadTimeUsecs field if non-nil, zero value otherwise.

### GetKeyFileDownloadTimeUsecsOk

`func (o *Vault) GetKeyFileDownloadTimeUsecsOk() (*int64, bool)`

GetKeyFileDownloadTimeUsecsOk returns a tuple with the KeyFileDownloadTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFileDownloadTimeUsecs

`func (o *Vault) SetKeyFileDownloadTimeUsecs(v int64)`

SetKeyFileDownloadTimeUsecs sets KeyFileDownloadTimeUsecs field to given value.

### HasKeyFileDownloadTimeUsecs

`func (o *Vault) HasKeyFileDownloadTimeUsecs() bool`

HasKeyFileDownloadTimeUsecs returns a boolean if a field has been set.

### SetKeyFileDownloadTimeUsecsNil

`func (o *Vault) SetKeyFileDownloadTimeUsecsNil(b bool)`

 SetKeyFileDownloadTimeUsecsNil sets the value for KeyFileDownloadTimeUsecs to be an explicit nil

### UnsetKeyFileDownloadTimeUsecs
`func (o *Vault) UnsetKeyFileDownloadTimeUsecs()`

UnsetKeyFileDownloadTimeUsecs ensures that no value is present for KeyFileDownloadTimeUsecs, not even an explicit nil
### GetKeyFileDownloadUser

`func (o *Vault) GetKeyFileDownloadUser() string`

GetKeyFileDownloadUser returns the KeyFileDownloadUser field if non-nil, zero value otherwise.

### GetKeyFileDownloadUserOk

`func (o *Vault) GetKeyFileDownloadUserOk() (*string, bool)`

GetKeyFileDownloadUserOk returns a tuple with the KeyFileDownloadUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFileDownloadUser

`func (o *Vault) SetKeyFileDownloadUser(v string)`

SetKeyFileDownloadUser sets KeyFileDownloadUser field to given value.

### HasKeyFileDownloadUser

`func (o *Vault) HasKeyFileDownloadUser() bool`

HasKeyFileDownloadUser returns a boolean if a field has been set.

### SetKeyFileDownloadUserNil

`func (o *Vault) SetKeyFileDownloadUserNil(b bool)`

 SetKeyFileDownloadUserNil sets the value for KeyFileDownloadUser to be an explicit nil

### UnsetKeyFileDownloadUser
`func (o *Vault) UnsetKeyFileDownloadUser()`

UnsetKeyFileDownloadUser ensures that no value is present for KeyFileDownloadUser, not even an explicit nil
### GetName

`func (o *Vault) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Vault) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Vault) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Vault) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Vault) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Vault) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRemovalState

`func (o *Vault) GetRemovalState() string`

GetRemovalState returns the RemovalState field if non-nil, zero value otherwise.

### GetRemovalStateOk

`func (o *Vault) GetRemovalStateOk() (*string, bool)`

GetRemovalStateOk returns a tuple with the RemovalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalState

`func (o *Vault) SetRemovalState(v string)`

SetRemovalState sets RemovalState field to given value.

### HasRemovalState

`func (o *Vault) HasRemovalState() bool`

HasRemovalState returns a boolean if a field has been set.

### SetRemovalStateNil

`func (o *Vault) SetRemovalStateNil(b bool)`

 SetRemovalStateNil sets the value for RemovalState to be an explicit nil

### UnsetRemovalState
`func (o *Vault) UnsetRemovalState()`

UnsetRemovalState ensures that no value is present for RemovalState, not even an explicit nil
### GetType

`func (o *Vault) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Vault) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Vault) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Vault) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Vault) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Vault) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUsageType

`func (o *Vault) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *Vault) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *Vault) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *Vault) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### SetUsageTypeNil

`func (o *Vault) SetUsageTypeNil(b bool)`

 SetUsageTypeNil sets the value for UsageType to be an explicit nil

### UnsetUsageType
`func (o *Vault) UnsetUsageType()`

UnsetUsageType ensures that no value is present for UsageType, not even an explicit nil
### GetVaultBandwidthLimits

`func (o *Vault) GetVaultBandwidthLimits() VaultBandwidthLimits`

GetVaultBandwidthLimits returns the VaultBandwidthLimits field if non-nil, zero value otherwise.

### GetVaultBandwidthLimitsOk

`func (o *Vault) GetVaultBandwidthLimitsOk() (*VaultBandwidthLimits, bool)`

GetVaultBandwidthLimitsOk returns a tuple with the VaultBandwidthLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultBandwidthLimits

`func (o *Vault) SetVaultBandwidthLimits(v VaultBandwidthLimits)`

SetVaultBandwidthLimits sets VaultBandwidthLimits field to given value.

### HasVaultBandwidthLimits

`func (o *Vault) HasVaultBandwidthLimits() bool`

HasVaultBandwidthLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


