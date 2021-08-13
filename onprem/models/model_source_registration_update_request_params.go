/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package onprem
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/onprem/utils"
	"fmt"
)

var _ = NullableBool{}

// SourceRegistrationUpdateRequestParams Specifies the Source registration Update request parameters.
type SourceRegistrationUpdateRequestParams struct {
	// Specifies the environment type of the Protection Source.
	Environment NullableString `json:"environment"`
	// Specifies if credentials are encrypted by internal key.
	IsInternalEncrypted NullableBool `json:"isInternalEncrypted,omitempty"`
	// Specifies the key that user has encrypted the credential with.
	EncryptionKey NullableString `json:"encryptionKey,omitempty"`
	// Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user.
	ConnectionId NullableInt64 `json:"connectionId,omitempty"`
	// Specfies the list of connections for the source.
	Connections []ConnectionConfig `json:"connections,omitempty"`
	VmwareParams *VmwareSourceRegistrationParams `json:"vmwareParams,omitempty"`
	PhysicalParams *PhysicalSourceRegistrationParams `json:"physicalParams,omitempty"`
	GenericNasParams *GenericNasRegistrationParams `json:"genericNasParams,omitempty"`
	IsilonParams *IsilonRegistrationParams `json:"isilonParams,omitempty"`
	NetappParams *NetappRegistrationParams `json:"netappParams,omitempty"`
	ElastifileParams *ElastifileRegistrationParams `json:"elastifileParams,omitempty"`
	FlashbladeParams *FlashbladeRegistrationParams `json:"flashbladeParams,omitempty"`
	GpfsParams *GpfsRegistrationParams `json:"gpfsParams,omitempty"`
	CassandraParams *CassandraSourceRegistrationParams `json:"cassandraParams,omitempty"`
	MongodbParams *MongoDBSourceRegistrationParams `json:"mongodbParams,omitempty"`
	CouchbaseParams *CouchbaseSourceRegistrationParams `json:"couchbaseParams,omitempty"`
	HdfsParams *HdfsSourceRegistrationParams `json:"hdfsParams,omitempty"`
	HbaseParams *HbaseSourceRegistrationParams `json:"hbaseParams,omitempty"`
	HiveParams *HiveSourceRegistrationParams `json:"hiveParams,omitempty"`
	UdaParams *UdaSourceRegistrationParams `json:"udaParams,omitempty"`
	Office365Params *Office365SourceRegistrationParams `json:"office365Params,omitempty"`
	AwsParams *AwsSourceRegistrationParams `json:"awsParams,omitempty"`
	HypervParams *HyperVSourceRegistrationParams `json:"hypervParams,omitempty"`
}

// NewSourceRegistrationUpdateRequestParams instantiates a new SourceRegistrationUpdateRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceRegistrationUpdateRequestParams(environment NullableString) *SourceRegistrationUpdateRequestParams {
	this := SourceRegistrationUpdateRequestParams{}
	this.Environment = environment
	return &this
}

// NewSourceRegistrationUpdateRequestParamsWithDefaults instantiates a new SourceRegistrationUpdateRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceRegistrationUpdateRequestParamsWithDefaults() *SourceRegistrationUpdateRequestParams {
	this := SourceRegistrationUpdateRequestParams{}
	return &this
}

// GetEnvironment returns the Environment field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SourceRegistrationUpdateRequestParams) GetEnvironment() string {
	if o == nil || o.Environment.Get() == nil {
		var ret string
		return ret
	}

	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceRegistrationUpdateRequestParams) GetEnvironmentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// SetEnvironment sets field value
func (o *SourceRegistrationUpdateRequestParams) SetEnvironment(v string) {
	o.Environment.Set(&v)
}

// GetIsInternalEncrypted returns the IsInternalEncrypted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceRegistrationUpdateRequestParams) GetIsInternalEncrypted() bool {
	if o == nil || o.IsInternalEncrypted.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsInternalEncrypted.Get()
}

// GetIsInternalEncryptedOk returns a tuple with the IsInternalEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceRegistrationUpdateRequestParams) GetIsInternalEncryptedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsInternalEncrypted.Get(), o.IsInternalEncrypted.IsSet()
}

// HasIsInternalEncrypted returns a boolean if a field has been set.
func (o *SourceRegistrationUpdateRequestParams) HasIsInternalEncrypted() bool {
	if o != nil && o.IsInternalEncrypted.IsSet() {
		return true
	}

	return false
}

// SetIsInternalEncrypted gets a reference to the given NullableBool and assigns it to the IsInternalEncrypted field.
func (o *SourceRegistrationUpdateRequestParams) SetIsInternalEncrypted(v bool) {
	o.IsInternalEncrypted.Set(&v)
}
// SetIsInternalEncryptedNil sets the value for IsInternalEncrypted to be an explicit nil
func (o *SourceRegistrationUpdateRequestParams) SetIsInternalEncryptedNil() {
	o.IsInternalEncrypted.Set(nil)
}

// UnsetIsInternalEncrypted ensures that no value is present for IsInternalEncrypted, not even an explicit nil
func (o *SourceRegistrationUpdateRequestParams) UnsetIsInternalEncrypted() {
	o.IsInternalEncrypted.Unset()
}

// GetEncryptionKey returns the EncryptionKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceRegistrationUpdateRequestParams) GetEncryptionKey() string {
	if o == nil || o.EncryptionKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.EncryptionKey.Get()
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceRegistrationUpdateRequestParams) GetEncryptionKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EncryptionKey.Get(), o.EncryptionKey.IsSet()
}

// HasEncryptionKey returns a boolean if a field has been set.
func (o *SourceRegistrationUpdateRequestParams) HasEncryptionKey() bool {
	if o != nil && o.EncryptionKey.IsSet() {
		return true
	}

	return false
}

// SetEncryptionKey gets a reference to the given NullableString and assigns it to the EncryptionKey field.
func (o *SourceRegistrationUpdateRequestParams) SetEncryptionKey(v string) {
	o.EncryptionKey.Set(&v)
}
// SetEncryptionKeyNil sets the value for EncryptionKey to be an explicit nil
func (o *SourceRegistrationUpdateRequestParams) SetEncryptionKeyNil() {
	o.EncryptionKey.Set(nil)
}

// UnsetEncryptionKey ensures that no value is present for EncryptionKey, not even an explicit nil
func (o *SourceRegistrationUpdateRequestParams) UnsetEncryptionKey() {
	o.EncryptionKey.Unset()
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceRegistrationUpdateRequestParams) GetConnectionId() int64 {
	if o == nil || o.ConnectionId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ConnectionId.Get()
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceRegistrationUpdateRequestParams) GetConnectionIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConnectionId.Get(), o.ConnectionId.IsSet()
}

// HasConnectionId returns a boolean if a field has been set.
func (o *SourceRegistrationUpdateRequestParams) HasConnectionId() bool {
	if o != nil && o.ConnectionId.IsSet() {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given NullableInt64 and assigns it to the ConnectionId field.
func (o *SourceRegistrationUpdateRequestParams) SetConnectionId(v int64) {
	o.ConnectionId.Set(&v)
}
// SetConnectionIdNil sets the value for ConnectionId to be an explicit nil
func (o *SourceRegistrationUpdateRequestParams) SetConnectionIdNil() {
	o.ConnectionId.Set(nil)
}

// UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
func (o *SourceRegistrationUpdateRequestParams) UnsetConnectionId() {
	o.ConnectionId.Unset()
}

// GetConnections returns the Connections field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceRegistrationUpdateRequestParams) GetConnections() []ConnectionConfig {
	if o == nil  {
		var ret []ConnectionConfig
		return ret
	}
	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceRegistrationUpdateRequestParams) GetConnectionsOk() (*[]ConnectionConfig, bool) {
	if o == nil || o.Connections == nil {
		return nil, false
	}
	return &o.Connections, true
}

// HasConnections returns a boolean if a field has been set.
func (o *SourceRegistrationUpdateRequestParams) HasConnections() bool {
	if o != nil && o.Connections != nil {
		return true
	}

	return false
}

// SetConnections gets a reference to the given []ConnectionConfig and assigns it to the Connections field.
func (o *SourceRegistrationUpdateRequestParams) SetConnections(v []ConnectionConfig) {
	o.Connections = v
}

// GetVmwareParams returns the VmwareParams field value if set, zero value otherwise.
func (o *SourceRegistrationUpdateRequestParams) GetVmwareParams() VmwareSourceRegistrationParams {
	if o == nil || o.VmwareParams == nil {
		var ret VmwareSourceRegistrationParams
		return ret
	}
	return *o.VmwareParams
}

// GetVmwareParamsOk returns a tuple with the VmwareParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceRegistrationUpdateRequestParams) GetVmwareParamsOk() (*VmwareSourceRegistrationParams, bool) {
	if o == nil || o.VmwareParams == nil {
		return nil, false
	}
	return o.VmwareParams, true
}

// HasVmwareParams returns a boolean if a field has been set.
func (o *SourceRegistrationUpdateRequestParams) HasVmwareParams() bool {
	if o != nil && o.VmwareParams != nil {
		return true
	}

	return false
}

// SetVmwareParams gets a reference to the given VmwareSourceRegistrationParams and assigns it to the VmwareParams field.
func (o *SourceRegistrationUpdateRequestParams) SetVmwareParams(v VmwareSourceRegistrationParams) {
	o.VmwareParams = &v
}

// GetPhysicalParams returns the PhysicalParams field value if set, zero value otherwise.
func (o *SourceRegistrationUpdateRequestParams) GetPhysicalParams() PhysicalSourceRegistrationParams {
	if o == nil || o.PhysicalParams == nil {
		var ret PhysicalSourceRegistrationParams
		return ret
	}
	return *o.PhysicalParams
}

// GetPhysicalParamsOk returns a tuple with the PhysicalParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceRegistrationUpdateRequestParams) GetPhysicalParamsOk() (*PhysicalSourceRegistrationParams, bool) {
	if o == nil || o.PhysicalParams == nil {
		return nil, false
	}
	return o.PhysicalParams, true
}

// HasPhysicalParams returns a boolean if a field has been set.
func (o *SourceRegistrationUpdateRequestParams) HasPhysicalParams() bool {
	if o != nil && o.PhysicalParams != nil {
		return true
	}

	return false
}

// SetPhysicalParams gets a reference to the given PhysicalSourceRegistrationParams and assigns it to the PhysicalParams field.
func (o *SourceRegistrationUpdateRequestParams) SetPhysicalParams(v PhysicalSourceRegistrationParams) {
	o.PhysicalParams = &v
}

// GetGenericNasParams returns the GenericNasParams field value if set, zero value otherwise.
func (o *SourceRegistrationUpdateRequestParams) GetGenericNasParams() GenericNasRegistrationParams {
	if o == nil || o.GenericNasParams == nil {
		var ret GenericNasRegistrationParams
		return ret
	}
	return *o.GenericNasParams
}

// GetGenericNasParamsOk returns a tuple with the GenericNasParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceRegistrationUpdateRequestParams) GetGenericNasParamsOk() (*GenericNasRegistrationParams, bool) {
	if o == nil || o.GenericNasParams == nil {
		return nil, false
	}
	return o.GenericNasParams, true
}

// HasGenericNasParams returns a boolean if a field has been set.
func (o *SourceRegistrationUpdateRequestParams) HasGenericNasParams() bool {
	if o != nil && o.GenericNasParams != nil {
		return true
	}

	return false
}

// SetGenericNasParams gets a reference to the given GenericNasRegistrationParams and assigns it to the GenericNasParams field.
func (o *SourceRegistrationUpdateRequestParams) SetGenericNasParams(v GenericNasRegistrationParams) {
	o.GenericNasParams = &v
}

// GetIsilonParams returns the IsilonParams field value if set, zero value otherwise.
func (o *SourceRegistrationUpdateRequestParams) GetIsilonParams() IsilonRegistrationParams {
	if o == nil || o.IsilonParams == nil {
		var ret IsilonRegistrationParams
		return ret
	}
	return *o.IsilonParams
}

// GetIsilonParamsOk returns a tuple with the IsilonParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceRegistrationUpdateRequestParams) GetIsilonParamsOk() (*IsilonRegistrationParams, bool) {
	if o == nil || o.IsilonParams == nil {
		return nil, false
	}
	return o.IsilonParams, true
}

// HasIsilonParams returns a boolean if a field has been set.
func (o *SourceRegistrationUpdateRequestParams) HasIsilonParams() bool {
	if o != nil && o.IsilonParams != nil {
		return true
	}

	return false
}

// SetIsilonParams gets a reference to the given IsilonRegistrationParams and assigns it to the IsilonParams field.
func (o *SourceRegistrationUpdateRequestParams) SetIsilonParams(v IsilonRegistrationParams) {
	o.IsilonParams = &v
}

// GetNetappParams returns the NetappParams field value if set, zero value otherwise.
func (o *SourceRegistrationUpdateRequestParams) GetNetappParams() NetappRegistrationParams {
	if o == nil || o.NetappParams == nil {
		var ret NetappRegistrationParams
		return ret
	}
	return *o.NetappParams
}

// GetNetappParamsOk returns a tuple with the NetappParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceRegistrationUpdateRequestParams) GetNetappParamsOk() (*NetappRegistrationParams, bool) {
	if o == nil || o.NetappParams == nil {
		return nil, false
	}
	return o.NetappParams, true
}

// HasNetappParams returns a boolean if a field has been set.
func (o *SourceRegistrationUpdateRequestParams) HasNetappParams() bool {
	if o != nil && o.NetappParams != nil {
		return true
	}

	return false
}

// SetNetappParams gets a reference to the given NetappRegistrationParams and assigns it to the NetappParams field.
func (o *SourceRegistrationUpdateRequestParams) SetNetappParams(v NetappRegistrationParams) {
	o.NetappParams = &v
}

// GetElastifileParams returns the ElastifileParams field value if set, zero value otherwise.
func (o *SourceRegistrationUpdateRequestParams) GetElastifileParams() ElastifileRegistrationParams {
	if o == nil || o.ElastifileParams == nil {
		var ret ElastifileRegistrationParams
		return ret
	}
	return *o.ElastifileParams
}

// GetElastifileParamsOk returns a tuple with the ElastifileParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceRegistrationUpdateRequestParams) GetElastifileParamsOk() (*ElastifileRegistrationParams, bool) {
	if o == nil || o.ElastifileParams == nil {
		return nil, false
	}
	return o.ElastifileParams, true
}

// HasElastifileParams returns a boolean if a field has been set.
func (o *SourceRegistrationUpdateRequestParams) HasElastifileParams() bool {
	if o != nil && o.ElastifileParams != nil {
		return true
	}

	return false
}

// SetElastifileParams gets a reference to the given ElastifileRegistrationParams and assigns it to the ElastifileParams field.
func (o *SourceRegistrationUpdateRequestParams) SetElastifileParams(v ElastifileRegistrationParams) {
	o.ElastifileParams = &v
}

// GetFlashbladeParams returns the FlashbladeParams field value if set, zero value otherwise.
func (o *SourceRegistrationUpdateRequestParams) GetFlashbladeParams() FlashbladeRegistrationParams {
	if o == nil || o.FlashbladeParams == nil {
		var ret FlashbladeRegistrationParams
		return ret
	}
	return *o.FlashbladeParams
}

// GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceRegistrationUpdateRequestParams) GetFlashbladeParamsOk() (*FlashbladeRegistrationParams, bool) {
	if o == nil || o.FlashbladeParams == nil {
		return nil, false
	}
	return o.FlashbladeParams, true
}

// HasFlashbladeParams returns a boolean if a field has been set.
func (o *SourceRegistrationUpdateRequestParams) HasFlashbladeParams() bool {
	if o != nil && o.FlashbladeParams != nil {
		return true
	}

	return false
}

// SetFlashbladeParams gets a reference to the given FlashbladeRegistrationParams and assigns it to the FlashbladeParams field.
func (o *SourceRegistrationUpdateRequestParams) SetFlashbladeParams(v FlashbladeRegistrationParams) {
	o.FlashbladeParams = &v
}

// GetGpfsParams returns the GpfsParams field value if set, zero value otherwise.
func (o *SourceRegistrationUpdateRequestParams) GetGpfsParams() GpfsRegistrationParams {
	if o == nil || o.GpfsParams == nil {
		var ret GpfsRegistrationParams
		return ret
	}
	return *o.GpfsParams
}

// GetGpfsParamsOk returns a tuple with the GpfsParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceRegistrationUpdateRequestParams) GetGpfsParamsOk() (*GpfsRegistrationParams, bool) {
	if o == nil || o.GpfsParams == nil {
		return nil, false
	}
	return o.GpfsParams, true
}

// HasGpfsParams returns a boolean if a field has been set.
func (o *SourceRegistrationUpdateRequestParams) HasGpfsParams() bool {
	if o != nil && o.GpfsParams != nil {
		return true
	}

	return false
}

// SetGpfsParams gets a reference to the given GpfsRegistrationParams and assigns it to the GpfsParams field.
func (o *SourceRegistrationUpdateRequestParams) SetGpfsParams(v GpfsRegistrationParams) {
	o.GpfsParams = &v
}

// GetCassandraParams returns the CassandraParams field value if set, zero value otherwise.
func (o *SourceRegistrationUpdateRequestParams) GetCassandraParams() CassandraSourceRegistrationParams {
	if o == nil || o.CassandraParams == nil {
		var ret CassandraSourceRegistrationParams
		return ret
	}
	return *o.CassandraParams
}

// GetCassandraParamsOk returns a tuple with the CassandraParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceRegistrationUpdateRequestParams) GetCassandraParamsOk() (*CassandraSourceRegistrationParams, bool) {
	if o == nil || o.CassandraParams == nil {
		return nil, false
	}
	return o.CassandraParams, true
}

// HasCassandraParams returns a boolean if a field has been set.
func (o *SourceRegistrationUpdateRequestParams) HasCassandraParams() bool {
	if o != nil && o.CassandraParams != nil {
		return true
	}

	return false
}

// SetCassandraParams gets a reference to the given CassandraSourceRegistrationParams and assigns it to the CassandraParams field.
func (o *SourceRegistrationUpdateRequestParams) SetCassandraParams(v CassandraSourceRegistrationParams) {
	o.CassandraParams = &v
}

// GetMongodbParams returns the MongodbParams field value if set, zero value otherwise.
func (o *SourceRegistrationUpdateRequestParams) GetMongodbParams() MongoDBSourceRegistrationParams {
	if o == nil || o.MongodbParams == nil {
		var ret MongoDBSourceRegistrationParams
		return ret
	}
	return *o.MongodbParams
}

// GetMongodbParamsOk returns a tuple with the MongodbParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceRegistrationUpdateRequestParams) GetMongodbParamsOk() (*MongoDBSourceRegistrationParams, bool) {
	if o == nil || o.MongodbParams == nil {
		return nil, false
	}
	return o.MongodbParams, true
}

// HasMongodbParams returns a boolean if a field has been set.
func (o *SourceRegistrationUpdateRequestParams) HasMongodbParams() bool {
	if o != nil && o.MongodbParams != nil {
		return true
	}

	return false
}

// SetMongodbParams gets a reference to the given MongoDBSourceRegistrationParams and assigns it to the MongodbParams field.
func (o *SourceRegistrationUpdateRequestParams) SetMongodbParams(v MongoDBSourceRegistrationParams) {
	o.MongodbParams = &v
}

// GetCouchbaseParams returns the CouchbaseParams field value if set, zero value otherwise.
func (o *SourceRegistrationUpdateRequestParams) GetCouchbaseParams() CouchbaseSourceRegistrationParams {
	if o == nil || o.CouchbaseParams == nil {
		var ret CouchbaseSourceRegistrationParams
		return ret
	}
	return *o.CouchbaseParams
}

// GetCouchbaseParamsOk returns a tuple with the CouchbaseParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceRegistrationUpdateRequestParams) GetCouchbaseParamsOk() (*CouchbaseSourceRegistrationParams, bool) {
	if o == nil || o.CouchbaseParams == nil {
		return nil, false
	}
	return o.CouchbaseParams, true
}

// HasCouchbaseParams returns a boolean if a field has been set.
func (o *SourceRegistrationUpdateRequestParams) HasCouchbaseParams() bool {
	if o != nil && o.CouchbaseParams != nil {
		return true
	}

	return false
}

// SetCouchbaseParams gets a reference to the given CouchbaseSourceRegistrationParams and assigns it to the CouchbaseParams field.
func (o *SourceRegistrationUpdateRequestParams) SetCouchbaseParams(v CouchbaseSourceRegistrationParams) {
	o.CouchbaseParams = &v
}

// GetHdfsParams returns the HdfsParams field value if set, zero value otherwise.
func (o *SourceRegistrationUpdateRequestParams) GetHdfsParams() HdfsSourceRegistrationParams {
	if o == nil || o.HdfsParams == nil {
		var ret HdfsSourceRegistrationParams
		return ret
	}
	return *o.HdfsParams
}

// GetHdfsParamsOk returns a tuple with the HdfsParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceRegistrationUpdateRequestParams) GetHdfsParamsOk() (*HdfsSourceRegistrationParams, bool) {
	if o == nil || o.HdfsParams == nil {
		return nil, false
	}
	return o.HdfsParams, true
}

// HasHdfsParams returns a boolean if a field has been set.
func (o *SourceRegistrationUpdateRequestParams) HasHdfsParams() bool {
	if o != nil && o.HdfsParams != nil {
		return true
	}

	return false
}

// SetHdfsParams gets a reference to the given HdfsSourceRegistrationParams and assigns it to the HdfsParams field.
func (o *SourceRegistrationUpdateRequestParams) SetHdfsParams(v HdfsSourceRegistrationParams) {
	o.HdfsParams = &v
}

// GetHbaseParams returns the HbaseParams field value if set, zero value otherwise.
func (o *SourceRegistrationUpdateRequestParams) GetHbaseParams() HbaseSourceRegistrationParams {
	if o == nil || o.HbaseParams == nil {
		var ret HbaseSourceRegistrationParams
		return ret
	}
	return *o.HbaseParams
}

// GetHbaseParamsOk returns a tuple with the HbaseParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceRegistrationUpdateRequestParams) GetHbaseParamsOk() (*HbaseSourceRegistrationParams, bool) {
	if o == nil || o.HbaseParams == nil {
		return nil, false
	}
	return o.HbaseParams, true
}

// HasHbaseParams returns a boolean if a field has been set.
func (o *SourceRegistrationUpdateRequestParams) HasHbaseParams() bool {
	if o != nil && o.HbaseParams != nil {
		return true
	}

	return false
}

// SetHbaseParams gets a reference to the given HbaseSourceRegistrationParams and assigns it to the HbaseParams field.
func (o *SourceRegistrationUpdateRequestParams) SetHbaseParams(v HbaseSourceRegistrationParams) {
	o.HbaseParams = &v
}

// GetHiveParams returns the HiveParams field value if set, zero value otherwise.
func (o *SourceRegistrationUpdateRequestParams) GetHiveParams() HiveSourceRegistrationParams {
	if o == nil || o.HiveParams == nil {
		var ret HiveSourceRegistrationParams
		return ret
	}
	return *o.HiveParams
}

// GetHiveParamsOk returns a tuple with the HiveParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceRegistrationUpdateRequestParams) GetHiveParamsOk() (*HiveSourceRegistrationParams, bool) {
	if o == nil || o.HiveParams == nil {
		return nil, false
	}
	return o.HiveParams, true
}

// HasHiveParams returns a boolean if a field has been set.
func (o *SourceRegistrationUpdateRequestParams) HasHiveParams() bool {
	if o != nil && o.HiveParams != nil {
		return true
	}

	return false
}

// SetHiveParams gets a reference to the given HiveSourceRegistrationParams and assigns it to the HiveParams field.
func (o *SourceRegistrationUpdateRequestParams) SetHiveParams(v HiveSourceRegistrationParams) {
	o.HiveParams = &v
}

// GetUdaParams returns the UdaParams field value if set, zero value otherwise.
func (o *SourceRegistrationUpdateRequestParams) GetUdaParams() UdaSourceRegistrationParams {
	if o == nil || o.UdaParams == nil {
		var ret UdaSourceRegistrationParams
		return ret
	}
	return *o.UdaParams
}

// GetUdaParamsOk returns a tuple with the UdaParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceRegistrationUpdateRequestParams) GetUdaParamsOk() (*UdaSourceRegistrationParams, bool) {
	if o == nil || o.UdaParams == nil {
		return nil, false
	}
	return o.UdaParams, true
}

// HasUdaParams returns a boolean if a field has been set.
func (o *SourceRegistrationUpdateRequestParams) HasUdaParams() bool {
	if o != nil && o.UdaParams != nil {
		return true
	}

	return false
}

// SetUdaParams gets a reference to the given UdaSourceRegistrationParams and assigns it to the UdaParams field.
func (o *SourceRegistrationUpdateRequestParams) SetUdaParams(v UdaSourceRegistrationParams) {
	o.UdaParams = &v
}

// GetOffice365Params returns the Office365Params field value if set, zero value otherwise.
func (o *SourceRegistrationUpdateRequestParams) GetOffice365Params() Office365SourceRegistrationParams {
	if o == nil || o.Office365Params == nil {
		var ret Office365SourceRegistrationParams
		return ret
	}
	return *o.Office365Params
}

// GetOffice365ParamsOk returns a tuple with the Office365Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceRegistrationUpdateRequestParams) GetOffice365ParamsOk() (*Office365SourceRegistrationParams, bool) {
	if o == nil || o.Office365Params == nil {
		return nil, false
	}
	return o.Office365Params, true
}

// HasOffice365Params returns a boolean if a field has been set.
func (o *SourceRegistrationUpdateRequestParams) HasOffice365Params() bool {
	if o != nil && o.Office365Params != nil {
		return true
	}

	return false
}

// SetOffice365Params gets a reference to the given Office365SourceRegistrationParams and assigns it to the Office365Params field.
func (o *SourceRegistrationUpdateRequestParams) SetOffice365Params(v Office365SourceRegistrationParams) {
	o.Office365Params = &v
}

// GetAwsParams returns the AwsParams field value if set, zero value otherwise.
func (o *SourceRegistrationUpdateRequestParams) GetAwsParams() AwsSourceRegistrationParams {
	if o == nil || o.AwsParams == nil {
		var ret AwsSourceRegistrationParams
		return ret
	}
	return *o.AwsParams
}

// GetAwsParamsOk returns a tuple with the AwsParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceRegistrationUpdateRequestParams) GetAwsParamsOk() (*AwsSourceRegistrationParams, bool) {
	if o == nil || o.AwsParams == nil {
		return nil, false
	}
	return o.AwsParams, true
}

// HasAwsParams returns a boolean if a field has been set.
func (o *SourceRegistrationUpdateRequestParams) HasAwsParams() bool {
	if o != nil && o.AwsParams != nil {
		return true
	}

	return false
}

// SetAwsParams gets a reference to the given AwsSourceRegistrationParams and assigns it to the AwsParams field.
func (o *SourceRegistrationUpdateRequestParams) SetAwsParams(v AwsSourceRegistrationParams) {
	o.AwsParams = &v
}

// GetHypervParams returns the HypervParams field value if set, zero value otherwise.
func (o *SourceRegistrationUpdateRequestParams) GetHypervParams() HyperVSourceRegistrationParams {
	if o == nil || o.HypervParams == nil {
		var ret HyperVSourceRegistrationParams
		return ret
	}
	return *o.HypervParams
}

// GetHypervParamsOk returns a tuple with the HypervParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceRegistrationUpdateRequestParams) GetHypervParamsOk() (*HyperVSourceRegistrationParams, bool) {
	if o == nil || o.HypervParams == nil {
		return nil, false
	}
	return o.HypervParams, true
}

// HasHypervParams returns a boolean if a field has been set.
func (o *SourceRegistrationUpdateRequestParams) HasHypervParams() bool {
	if o != nil && o.HypervParams != nil {
		return true
	}

	return false
}

// SetHypervParams gets a reference to the given HyperVSourceRegistrationParams and assigns it to the HypervParams field.
func (o *SourceRegistrationUpdateRequestParams) SetHypervParams(v HyperVSourceRegistrationParams) {
	o.HypervParams = &v
}

func (o SourceRegistrationUpdateRequestParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["environment"] = o.Environment.Get()
	}
	if o.IsInternalEncrypted.IsSet() {
		toSerialize["isInternalEncrypted"] = o.IsInternalEncrypted.Get()
	}
	if o.EncryptionKey.IsSet() {
		toSerialize["encryptionKey"] = o.EncryptionKey.Get()
	}
	if o.ConnectionId.IsSet() {
		toSerialize["connectionId"] = o.ConnectionId.Get()
	}
	if o.Connections != nil {
		toSerialize["connections"] = o.Connections
	}
	if o.VmwareParams != nil {
		toSerialize["vmwareParams"] = o.VmwareParams
	}
	if o.PhysicalParams != nil {
		toSerialize["physicalParams"] = o.PhysicalParams
	}
	if o.GenericNasParams != nil {
		toSerialize["genericNasParams"] = o.GenericNasParams
	}
	if o.IsilonParams != nil {
		toSerialize["isilonParams"] = o.IsilonParams
	}
	if o.NetappParams != nil {
		toSerialize["netappParams"] = o.NetappParams
	}
	if o.ElastifileParams != nil {
		toSerialize["elastifileParams"] = o.ElastifileParams
	}
	if o.FlashbladeParams != nil {
		toSerialize["flashbladeParams"] = o.FlashbladeParams
	}
	if o.GpfsParams != nil {
		toSerialize["gpfsParams"] = o.GpfsParams
	}
	if o.CassandraParams != nil {
		toSerialize["cassandraParams"] = o.CassandraParams
	}
	if o.MongodbParams != nil {
		toSerialize["mongodbParams"] = o.MongodbParams
	}
	if o.CouchbaseParams != nil {
		toSerialize["couchbaseParams"] = o.CouchbaseParams
	}
	if o.HdfsParams != nil {
		toSerialize["hdfsParams"] = o.HdfsParams
	}
	if o.HbaseParams != nil {
		toSerialize["hbaseParams"] = o.HbaseParams
	}
	if o.HiveParams != nil {
		toSerialize["hiveParams"] = o.HiveParams
	}
	if o.UdaParams != nil {
		toSerialize["udaParams"] = o.UdaParams
	}
	if o.Office365Params != nil {
		toSerialize["office365Params"] = o.Office365Params
	}
	if o.AwsParams != nil {
		toSerialize["awsParams"] = o.AwsParams
	}
	if o.HypervParams != nil {
		toSerialize["hypervParams"] = o.HypervParams
	}
	return json.Marshal(toSerialize)
}

type NullableSourceRegistrationUpdateRequestParams struct {
	value *SourceRegistrationUpdateRequestParams
	isSet bool
}

func (v NullableSourceRegistrationUpdateRequestParams) Get() *SourceRegistrationUpdateRequestParams {
	return v.value
}

func (v *NullableSourceRegistrationUpdateRequestParams) Set(val *SourceRegistrationUpdateRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceRegistrationUpdateRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceRegistrationUpdateRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceRegistrationUpdateRequestParams(val *SourceRegistrationUpdateRequestParams) *NullableSourceRegistrationUpdateRequestParams {
	return &NullableSourceRegistrationUpdateRequestParams{value: val, isSet: true}
}

func (v NullableSourceRegistrationUpdateRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceRegistrationUpdateRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o SourceRegistrationUpdateRequestParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}