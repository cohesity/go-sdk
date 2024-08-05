# SfdcAuroraClusterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuroraClusterArn** | **NullableString** | Arn of the Aurora cluster. | 
**DatabaseAccessIAMRoleArn** | **NullableString** | Contains the Arn of the IAM role of the which has access to the db user allocated to the tenant. | 
**DatabasePort** | **NullableString** | Database port to access the dbs on the Aurora cluster. | 
**DatabaseSchema** | Pointer to **NullableString** | Database schema to access the dbs on the Aurora cluster. | [optional] 
**DatabaseUser** | **NullableString** | Database user to access the dbs on the Aurora cluster. | 
**ReaderEndpoint** | Pointer to **NullableString** | Reader endpoint of the Aurora cluster. | [optional] 
**RegionId** | **NullableString** | Specifies the region id of the Aurora cluster. | 
**S3AccessIAMRoleArn** | **NullableString** | Contains the Arn of the IAM role which has read and write access to the tenant&#39;s s3 bucket. | 
**S3BucketName** | **NullableString** | Contains the tenant&#39;s S3 bucket. | 
**S3BucketPrefix** | **NullableString** | S3Bucket prefix for the intermediate. | 
**TenantId** | **NullableString** | Contains the Id of the tenant. | 
**WriterEndpoint** | **NullableString** | Writer endpoint of the Aurora cluster. | 

## Methods

### NewSfdcAuroraClusterInfo

`func NewSfdcAuroraClusterInfo(auroraClusterArn NullableString, databaseAccessIAMRoleArn NullableString, databasePort NullableString, databaseUser NullableString, regionId NullableString, s3AccessIAMRoleArn NullableString, s3BucketName NullableString, s3BucketPrefix NullableString, tenantId NullableString, writerEndpoint NullableString, ) *SfdcAuroraClusterInfo`

NewSfdcAuroraClusterInfo instantiates a new SfdcAuroraClusterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSfdcAuroraClusterInfoWithDefaults

`func NewSfdcAuroraClusterInfoWithDefaults() *SfdcAuroraClusterInfo`

NewSfdcAuroraClusterInfoWithDefaults instantiates a new SfdcAuroraClusterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuroraClusterArn

`func (o *SfdcAuroraClusterInfo) GetAuroraClusterArn() string`

GetAuroraClusterArn returns the AuroraClusterArn field if non-nil, zero value otherwise.

### GetAuroraClusterArnOk

`func (o *SfdcAuroraClusterInfo) GetAuroraClusterArnOk() (*string, bool)`

GetAuroraClusterArnOk returns a tuple with the AuroraClusterArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuroraClusterArn

`func (o *SfdcAuroraClusterInfo) SetAuroraClusterArn(v string)`

SetAuroraClusterArn sets AuroraClusterArn field to given value.


### SetAuroraClusterArnNil

`func (o *SfdcAuroraClusterInfo) SetAuroraClusterArnNil(b bool)`

 SetAuroraClusterArnNil sets the value for AuroraClusterArn to be an explicit nil

### UnsetAuroraClusterArn
`func (o *SfdcAuroraClusterInfo) UnsetAuroraClusterArn()`

UnsetAuroraClusterArn ensures that no value is present for AuroraClusterArn, not even an explicit nil
### GetDatabaseAccessIAMRoleArn

`func (o *SfdcAuroraClusterInfo) GetDatabaseAccessIAMRoleArn() string`

GetDatabaseAccessIAMRoleArn returns the DatabaseAccessIAMRoleArn field if non-nil, zero value otherwise.

### GetDatabaseAccessIAMRoleArnOk

`func (o *SfdcAuroraClusterInfo) GetDatabaseAccessIAMRoleArnOk() (*string, bool)`

GetDatabaseAccessIAMRoleArnOk returns a tuple with the DatabaseAccessIAMRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseAccessIAMRoleArn

`func (o *SfdcAuroraClusterInfo) SetDatabaseAccessIAMRoleArn(v string)`

SetDatabaseAccessIAMRoleArn sets DatabaseAccessIAMRoleArn field to given value.


### SetDatabaseAccessIAMRoleArnNil

`func (o *SfdcAuroraClusterInfo) SetDatabaseAccessIAMRoleArnNil(b bool)`

 SetDatabaseAccessIAMRoleArnNil sets the value for DatabaseAccessIAMRoleArn to be an explicit nil

### UnsetDatabaseAccessIAMRoleArn
`func (o *SfdcAuroraClusterInfo) UnsetDatabaseAccessIAMRoleArn()`

UnsetDatabaseAccessIAMRoleArn ensures that no value is present for DatabaseAccessIAMRoleArn, not even an explicit nil
### GetDatabasePort

`func (o *SfdcAuroraClusterInfo) GetDatabasePort() string`

GetDatabasePort returns the DatabasePort field if non-nil, zero value otherwise.

### GetDatabasePortOk

`func (o *SfdcAuroraClusterInfo) GetDatabasePortOk() (*string, bool)`

GetDatabasePortOk returns a tuple with the DatabasePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabasePort

`func (o *SfdcAuroraClusterInfo) SetDatabasePort(v string)`

SetDatabasePort sets DatabasePort field to given value.


### SetDatabasePortNil

`func (o *SfdcAuroraClusterInfo) SetDatabasePortNil(b bool)`

 SetDatabasePortNil sets the value for DatabasePort to be an explicit nil

### UnsetDatabasePort
`func (o *SfdcAuroraClusterInfo) UnsetDatabasePort()`

UnsetDatabasePort ensures that no value is present for DatabasePort, not even an explicit nil
### GetDatabaseSchema

`func (o *SfdcAuroraClusterInfo) GetDatabaseSchema() string`

GetDatabaseSchema returns the DatabaseSchema field if non-nil, zero value otherwise.

### GetDatabaseSchemaOk

`func (o *SfdcAuroraClusterInfo) GetDatabaseSchemaOk() (*string, bool)`

GetDatabaseSchemaOk returns a tuple with the DatabaseSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseSchema

`func (o *SfdcAuroraClusterInfo) SetDatabaseSchema(v string)`

SetDatabaseSchema sets DatabaseSchema field to given value.

### HasDatabaseSchema

`func (o *SfdcAuroraClusterInfo) HasDatabaseSchema() bool`

HasDatabaseSchema returns a boolean if a field has been set.

### SetDatabaseSchemaNil

`func (o *SfdcAuroraClusterInfo) SetDatabaseSchemaNil(b bool)`

 SetDatabaseSchemaNil sets the value for DatabaseSchema to be an explicit nil

### UnsetDatabaseSchema
`func (o *SfdcAuroraClusterInfo) UnsetDatabaseSchema()`

UnsetDatabaseSchema ensures that no value is present for DatabaseSchema, not even an explicit nil
### GetDatabaseUser

`func (o *SfdcAuroraClusterInfo) GetDatabaseUser() string`

GetDatabaseUser returns the DatabaseUser field if non-nil, zero value otherwise.

### GetDatabaseUserOk

`func (o *SfdcAuroraClusterInfo) GetDatabaseUserOk() (*string, bool)`

GetDatabaseUserOk returns a tuple with the DatabaseUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUser

`func (o *SfdcAuroraClusterInfo) SetDatabaseUser(v string)`

SetDatabaseUser sets DatabaseUser field to given value.


### SetDatabaseUserNil

`func (o *SfdcAuroraClusterInfo) SetDatabaseUserNil(b bool)`

 SetDatabaseUserNil sets the value for DatabaseUser to be an explicit nil

### UnsetDatabaseUser
`func (o *SfdcAuroraClusterInfo) UnsetDatabaseUser()`

UnsetDatabaseUser ensures that no value is present for DatabaseUser, not even an explicit nil
### GetReaderEndpoint

`func (o *SfdcAuroraClusterInfo) GetReaderEndpoint() string`

GetReaderEndpoint returns the ReaderEndpoint field if non-nil, zero value otherwise.

### GetReaderEndpointOk

`func (o *SfdcAuroraClusterInfo) GetReaderEndpointOk() (*string, bool)`

GetReaderEndpointOk returns a tuple with the ReaderEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReaderEndpoint

`func (o *SfdcAuroraClusterInfo) SetReaderEndpoint(v string)`

SetReaderEndpoint sets ReaderEndpoint field to given value.

### HasReaderEndpoint

`func (o *SfdcAuroraClusterInfo) HasReaderEndpoint() bool`

HasReaderEndpoint returns a boolean if a field has been set.

### SetReaderEndpointNil

`func (o *SfdcAuroraClusterInfo) SetReaderEndpointNil(b bool)`

 SetReaderEndpointNil sets the value for ReaderEndpoint to be an explicit nil

### UnsetReaderEndpoint
`func (o *SfdcAuroraClusterInfo) UnsetReaderEndpoint()`

UnsetReaderEndpoint ensures that no value is present for ReaderEndpoint, not even an explicit nil
### GetRegionId

`func (o *SfdcAuroraClusterInfo) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *SfdcAuroraClusterInfo) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *SfdcAuroraClusterInfo) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### SetRegionIdNil

`func (o *SfdcAuroraClusterInfo) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *SfdcAuroraClusterInfo) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetS3AccessIAMRoleArn

`func (o *SfdcAuroraClusterInfo) GetS3AccessIAMRoleArn() string`

GetS3AccessIAMRoleArn returns the S3AccessIAMRoleArn field if non-nil, zero value otherwise.

### GetS3AccessIAMRoleArnOk

`func (o *SfdcAuroraClusterInfo) GetS3AccessIAMRoleArnOk() (*string, bool)`

GetS3AccessIAMRoleArnOk returns a tuple with the S3AccessIAMRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3AccessIAMRoleArn

`func (o *SfdcAuroraClusterInfo) SetS3AccessIAMRoleArn(v string)`

SetS3AccessIAMRoleArn sets S3AccessIAMRoleArn field to given value.


### SetS3AccessIAMRoleArnNil

`func (o *SfdcAuroraClusterInfo) SetS3AccessIAMRoleArnNil(b bool)`

 SetS3AccessIAMRoleArnNil sets the value for S3AccessIAMRoleArn to be an explicit nil

### UnsetS3AccessIAMRoleArn
`func (o *SfdcAuroraClusterInfo) UnsetS3AccessIAMRoleArn()`

UnsetS3AccessIAMRoleArn ensures that no value is present for S3AccessIAMRoleArn, not even an explicit nil
### GetS3BucketName

`func (o *SfdcAuroraClusterInfo) GetS3BucketName() string`

GetS3BucketName returns the S3BucketName field if non-nil, zero value otherwise.

### GetS3BucketNameOk

`func (o *SfdcAuroraClusterInfo) GetS3BucketNameOk() (*string, bool)`

GetS3BucketNameOk returns a tuple with the S3BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3BucketName

`func (o *SfdcAuroraClusterInfo) SetS3BucketName(v string)`

SetS3BucketName sets S3BucketName field to given value.


### SetS3BucketNameNil

`func (o *SfdcAuroraClusterInfo) SetS3BucketNameNil(b bool)`

 SetS3BucketNameNil sets the value for S3BucketName to be an explicit nil

### UnsetS3BucketName
`func (o *SfdcAuroraClusterInfo) UnsetS3BucketName()`

UnsetS3BucketName ensures that no value is present for S3BucketName, not even an explicit nil
### GetS3BucketPrefix

`func (o *SfdcAuroraClusterInfo) GetS3BucketPrefix() string`

GetS3BucketPrefix returns the S3BucketPrefix field if non-nil, zero value otherwise.

### GetS3BucketPrefixOk

`func (o *SfdcAuroraClusterInfo) GetS3BucketPrefixOk() (*string, bool)`

GetS3BucketPrefixOk returns a tuple with the S3BucketPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3BucketPrefix

`func (o *SfdcAuroraClusterInfo) SetS3BucketPrefix(v string)`

SetS3BucketPrefix sets S3BucketPrefix field to given value.


### SetS3BucketPrefixNil

`func (o *SfdcAuroraClusterInfo) SetS3BucketPrefixNil(b bool)`

 SetS3BucketPrefixNil sets the value for S3BucketPrefix to be an explicit nil

### UnsetS3BucketPrefix
`func (o *SfdcAuroraClusterInfo) UnsetS3BucketPrefix()`

UnsetS3BucketPrefix ensures that no value is present for S3BucketPrefix, not even an explicit nil
### GetTenantId

`func (o *SfdcAuroraClusterInfo) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SfdcAuroraClusterInfo) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SfdcAuroraClusterInfo) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### SetTenantIdNil

`func (o *SfdcAuroraClusterInfo) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SfdcAuroraClusterInfo) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetWriterEndpoint

`func (o *SfdcAuroraClusterInfo) GetWriterEndpoint() string`

GetWriterEndpoint returns the WriterEndpoint field if non-nil, zero value otherwise.

### GetWriterEndpointOk

`func (o *SfdcAuroraClusterInfo) GetWriterEndpointOk() (*string, bool)`

GetWriterEndpointOk returns a tuple with the WriterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriterEndpoint

`func (o *SfdcAuroraClusterInfo) SetWriterEndpoint(v string)`

SetWriterEndpoint sets WriterEndpoint field to given value.


### SetWriterEndpointNil

`func (o *SfdcAuroraClusterInfo) SetWriterEndpointNil(b bool)`

 SetWriterEndpointNil sets the value for WriterEndpoint to be an explicit nil

### UnsetWriterEndpoint
`func (o *SfdcAuroraClusterInfo) UnsetWriterEndpoint()`

UnsetWriterEndpoint ensures that no value is present for WriterEndpoint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


