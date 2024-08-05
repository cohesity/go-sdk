# ConnectorAwsMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsInstanceId** | **NullableString** | AWS EC2 instance id. | 

## Methods

### NewConnectorAwsMetadata

`func NewConnectorAwsMetadata(awsInstanceId NullableString, ) *ConnectorAwsMetadata`

NewConnectorAwsMetadata instantiates a new ConnectorAwsMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorAwsMetadataWithDefaults

`func NewConnectorAwsMetadataWithDefaults() *ConnectorAwsMetadata`

NewConnectorAwsMetadataWithDefaults instantiates a new ConnectorAwsMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsInstanceId

`func (o *ConnectorAwsMetadata) GetAwsInstanceId() string`

GetAwsInstanceId returns the AwsInstanceId field if non-nil, zero value otherwise.

### GetAwsInstanceIdOk

`func (o *ConnectorAwsMetadata) GetAwsInstanceIdOk() (*string, bool)`

GetAwsInstanceIdOk returns a tuple with the AwsInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsInstanceId

`func (o *ConnectorAwsMetadata) SetAwsInstanceId(v string)`

SetAwsInstanceId sets AwsInstanceId field to given value.


### SetAwsInstanceIdNil

`func (o *ConnectorAwsMetadata) SetAwsInstanceIdNil(b bool)`

 SetAwsInstanceIdNil sets the value for AwsInstanceId to be an explicit nil

### UnsetAwsInstanceId
`func (o *ConnectorAwsMetadata) UnsetAwsInstanceId()`

UnsetAwsInstanceId ensures that no value is present for AwsInstanceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


