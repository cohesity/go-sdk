# MssqlConnectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostIdentifier** | **string** | Specifies the unique identifier to locate the SQL node or cluster. The host identifier can be IP address or FQDN. | 

## Methods

### NewMssqlConnectionParams

`func NewMssqlConnectionParams(hostIdentifier string, ) *MssqlConnectionParams`

NewMssqlConnectionParams instantiates a new MssqlConnectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMssqlConnectionParamsWithDefaults

`func NewMssqlConnectionParamsWithDefaults() *MssqlConnectionParams`

NewMssqlConnectionParamsWithDefaults instantiates a new MssqlConnectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostIdentifier

`func (o *MssqlConnectionParams) GetHostIdentifier() string`

GetHostIdentifier returns the HostIdentifier field if non-nil, zero value otherwise.

### GetHostIdentifierOk

`func (o *MssqlConnectionParams) GetHostIdentifierOk() (*string, bool)`

GetHostIdentifierOk returns a tuple with the HostIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostIdentifier

`func (o *MssqlConnectionParams) SetHostIdentifier(v string)`

SetHostIdentifier sets HostIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


