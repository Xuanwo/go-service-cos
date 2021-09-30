// Code generated by go generate via cmd/definitions; DO NOT EDIT.
package cos

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"

	. "github.com/beyondstorage/go-storage/v4/pairs"
	"github.com/beyondstorage/go-storage/v4/pkg/httpclient"
	"github.com/beyondstorage/go-storage/v4/services"
	. "github.com/beyondstorage/go-storage/v4/types"
)

var (
	_ Storager
	_ services.ServiceError
	_ httpclient.Options
	_ time.Duration
	_ http.Request
	_ Error
)

// Type is the type for cos
const Type = "cos"

// ObjectSystemMetadata stores system metadata for object.
type ObjectSystemMetadata struct {
	ServerSideEncryption                  string
	ServerSideEncryptionCosKmsKeyID       string
	ServerSideEncryptionCustomerAlgorithm string
	ServerSideEncryptionCustomerKeyMd5    string
	StorageClass                          string
}

// GetObjectSystemMetadata will get ObjectSystemMetadata from Object.
//
// - This function should not be called by service implementer.
// - The returning ObjectServiceMetadata is read only and should not be modified.
func GetObjectSystemMetadata(o *Object) ObjectSystemMetadata {
	sm, ok := o.GetSystemMetadata()
	if ok {
		return sm.(ObjectSystemMetadata)
	}
	return ObjectSystemMetadata{}
}

// setObjectSystemMetadata will set ObjectSystemMetadata into Object.
//
// - This function should only be called once, please make sure all data has been written before set.
func setObjectSystemMetadata(o *Object, sm ObjectSystemMetadata) {
	o.SetSystemMetadata(sm)
}

// StorageSystemMetadata stores system metadata for object.
type StorageSystemMetadata struct {
	ServerSideEncryption                  string
	ServerSideEncryptionCosKmsKeyID       string
	ServerSideEncryptionCustomerAlgorithm string
	ServerSideEncryptionCustomerKeyMd5    string
	StorageClass                          string
}

// GetStorageSystemMetadata will get StorageSystemMetadata from Storage.
//
// - This function should not be called by service implementer.
// - The returning StorageServiceMetadata is read only and should not be modified.
func GetStorageSystemMetadata(s *StorageMeta) StorageSystemMetadata {
	sm, ok := s.GetSystemMetadata()
	if ok {
		return sm.(StorageSystemMetadata)
	}
	return StorageSystemMetadata{}
}

// setStorageSystemMetadata will set StorageSystemMetadata into Storage.
//
// - This function should only be called once, please make sure all data has been written before set.
func setStorageSystemMetadata(s *StorageMeta, sm StorageSystemMetadata) {
	s.SetSystemMetadata(sm)
}

// WithDefaultServicePairs will apply default_service_pairs value to Options.
func WithDefaultServicePairs(v DefaultServicePairs) Pair {
	return Pair{Key: "default_service_pairs", Value: v}
}

// WithDefaultStoragePairs will apply default_storage_pairs value to Options.
func WithDefaultStoragePairs(v DefaultStoragePairs) Pair {
	return Pair{Key: "default_storage_pairs", Value: v}
}

// WithEnableVirtualDir will apply enable_virtual_dir value to Options.
//
// virtual_dir feature is designed for a service that doesn't have native dir support but wants to
// provide simulated operations.
//
// - If this feature is disabled (the default behavior), the service will behave like it doesn't have
// any dir support.
// - If this feature is enabled, the service will support simulated dir behavior in create_dir, create,
// list, delete, and so on.
//
// This feature was introduced in GSP-109.
func WithEnableVirtualDir() Pair {
	return Pair{Key: "enable_virtual_dir", Value: true}
}

// WithServerSideEncryption will apply server_side_encryption value to Options.
//
// the server-side encryption algorithm used when storing this object. It can be `AES-256` for SSE-COS,
// and `cos/kms` for SSE-KMS.
func WithServerSideEncryption(v string) Pair {
	return Pair{Key: "server_side_encryption", Value: v}
}

// WithServerSideEncryptionContext will apply server_side_encryption_context value to Options.
//
// specifies the COS KMS Encryption Context to use for object encryption. The value of this header
// is a base64-encoded UTF-8 string holding JSON with the encryption context key-value pairs.
func WithServerSideEncryptionContext(v string) Pair {
	return Pair{Key: "server_side_encryption_context", Value: v}
}

// WithServerSideEncryptionCosKmsKeyID will apply server_side_encryption_cos_kms_key_id
// value to Options.
//
// specifies the COS KMS key ID to use for object encryption.
func WithServerSideEncryptionCosKmsKeyID(v string) Pair {
	return Pair{Key: "server_side_encryption_cos_kms_key_id", Value: v}
}

// WithServerSideEncryptionCustomerAlgorithm will apply server_side_encryption_customer_algorithm
// value to Options.
//
// specifies the algorithm to use to when encrypting the object. Now only `AES256` is supported.
func WithServerSideEncryptionCustomerAlgorithm(v string) Pair {
	return Pair{Key: "server_side_encryption_customer_algorithm", Value: v}
}

// WithServerSideEncryptionCustomerKey will apply server_side_encryption_customer_key value
// to Options.
//
// specifies the customer-provided encryption key to encrypt/decrypt the source object. It must
// be a 32-byte AES-256 key.
func WithServerSideEncryptionCustomerKey(v []byte) Pair {
	return Pair{Key: "server_side_encryption_customer_key", Value: v}
}

// WithServiceFeatures will apply service_features value to Options.
func WithServiceFeatures(v ServiceFeatures) Pair {
	return Pair{Key: "service_features", Value: v}
}

// WithStorageClass will apply storage_class value to Options.
func WithStorageClass(v string) Pair {
	return Pair{Key: "storage_class", Value: v}
}

// WithStorageFeatures will apply storage_features value to Options.
func WithStorageFeatures(v StorageFeatures) Pair {
	return Pair{Key: "storage_features", Value: v}
}

var pairMap = map[string]string{"content_md5": "string", "content_type": "string", "context": "context.Context", "continuation_token": "string", "credential": "string", "default_content_type": "string", "default_io_callback": "func([]byte)", "default_service_pairs": "DefaultServicePairs", "default_storage_pairs": "DefaultStoragePairs", "enable_virtual_dir": "bool", "endpoint": "string", "expire": "time.Duration", "http_client_options": "*httpclient.Options", "interceptor": "Interceptor", "io_callback": "func([]byte)", "list_mode": "ListMode", "location": "string", "multipart_id": "string", "name": "string", "object_mode": "ObjectMode", "offset": "int64", "server_side_encryption": "string", "server_side_encryption_context": "string", "server_side_encryption_cos_kms_key_id": "string", "server_side_encryption_customer_algorithm": "string", "server_side_encryption_customer_key": "[]byte", "service_features": "ServiceFeatures", "size": "int64", "storage_class": "string", "storage_features": "StorageFeatures", "work_dir": "string"}
var _ Servicer = &Service{}

type ServiceFeatures struct {
}

// pairServiceNew is the parsed struct
type pairServiceNew struct {
	pairs []Pair

	// Required pairs
	HasCredential bool
	Credential    string
	// Optional pairs
	HasDefaultServicePairs bool
	DefaultServicePairs    DefaultServicePairs
	HasEndpoint            bool
	Endpoint               string
	HasHTTPClientOptions   bool
	HTTPClientOptions      *httpclient.Options
	HasServiceFeatures     bool
	ServiceFeatures        ServiceFeatures
	// Enable features
}

// parsePairServiceNew will parse Pair slice into *pairServiceNew
func parsePairServiceNew(opts []Pair) (pairServiceNew, error) {
	result :=
		pairServiceNew{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		case "credential":
			if result.HasCredential {
				continue
			}
			result.HasCredential = true
			result.Credential = v.Value.(string)
		case "default_service_pairs":
			if result.HasDefaultServicePairs {
				continue
			}
			result.HasDefaultServicePairs = true
			result.DefaultServicePairs = v.Value.(DefaultServicePairs)
		case "endpoint":
			if result.HasEndpoint {
				continue
			}
			result.HasEndpoint = true
			result.Endpoint = v.Value.(string)
		case "http_client_options":
			if result.HasHTTPClientOptions {
				continue
			}
			result.HasHTTPClientOptions = true
			result.HTTPClientOptions = v.Value.(*httpclient.Options)
		case "service_features":
			if result.HasServiceFeatures {
				continue
			}
			result.HasServiceFeatures = true
			result.ServiceFeatures = v.Value.(ServiceFeatures)
		}
	}
	// Enable features

	// Default pairs

	if !result.HasCredential {
		return pairServiceNew{}, services.PairRequiredError{Keys: []string{"credential"}}
	}
	return result, nil
}

// DefaultServicePairs is default pairs for specific action
type DefaultServicePairs struct {
	Create []Pair
	Delete []Pair
	Get    []Pair
	List   []Pair
}
type pairServiceCreate struct {
	pairs []Pair
	// Required pairs
	HasLocation bool
	Location    string
	// Optional pairs
}

func (s *Service) parsePairServiceCreate(opts []Pair) (pairServiceCreate, error) {
	result :=
		pairServiceCreate{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		case "location":
			if result.HasLocation {
				continue
			}
			result.HasLocation = true
			result.Location = v.Value.(string)
		default:
			return pairServiceCreate{}, services.PairUnsupportedError{Pair: v}
		}
	}
	if !result.HasLocation {
		return pairServiceCreate{}, services.PairRequiredError{Keys: []string{"location"}}
	}
	return result, nil
}

type pairServiceDelete struct {
	pairs []Pair
	// Required pairs
	HasLocation bool
	Location    string
	// Optional pairs
}

func (s *Service) parsePairServiceDelete(opts []Pair) (pairServiceDelete, error) {
	result :=
		pairServiceDelete{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		case "location":
			if result.HasLocation {
				continue
			}
			result.HasLocation = true
			result.Location = v.Value.(string)
		default:
			return pairServiceDelete{}, services.PairUnsupportedError{Pair: v}
		}
	}
	if !result.HasLocation {
		return pairServiceDelete{}, services.PairRequiredError{Keys: []string{"location"}}
	}
	return result, nil
}

type pairServiceGet struct {
	pairs []Pair
	// Required pairs
	HasLocation bool
	Location    string
	// Optional pairs
}

func (s *Service) parsePairServiceGet(opts []Pair) (pairServiceGet, error) {
	result :=
		pairServiceGet{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		case "location":
			if result.HasLocation {
				continue
			}
			result.HasLocation = true
			result.Location = v.Value.(string)
		default:
			return pairServiceGet{}, services.PairUnsupportedError{Pair: v}
		}
	}
	if !result.HasLocation {
		return pairServiceGet{}, services.PairRequiredError{Keys: []string{"location"}}
	}
	return result, nil
}

type pairServiceList struct {
	pairs []Pair
	// Required pairs
	// Optional pairs
}

func (s *Service) parsePairServiceList(opts []Pair) (pairServiceList, error) {
	result :=
		pairServiceList{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		default:
			return pairServiceList{}, services.PairUnsupportedError{Pair: v}
		}
	}

	return result, nil
}
func (s *Service) Create(name string, pairs ...Pair) (store Storager, err error) {
	ctx := context.Background()
	return s.CreateWithContext(ctx, name, pairs...)
}
func (s *Service) CreateWithContext(ctx context.Context, name string, pairs ...Pair) (store Storager, err error) {
	defer func() {
		err =
			s.formatError("create", err, name)
	}()

	pairs = append(pairs, s.defaultPairs.Create...)
	var opt pairServiceCreate

	opt, err = s.parsePairServiceCreate(pairs)
	if err != nil {
		return
	}
	return s.create(ctx, name, opt)
}
func (s *Service) Delete(name string, pairs ...Pair) (err error) {
	ctx := context.Background()
	return s.DeleteWithContext(ctx, name, pairs...)
}
func (s *Service) DeleteWithContext(ctx context.Context, name string, pairs ...Pair) (err error) {
	defer func() {
		err =
			s.formatError("delete", err, name)
	}()

	pairs = append(pairs, s.defaultPairs.Delete...)
	var opt pairServiceDelete

	opt, err = s.parsePairServiceDelete(pairs)
	if err != nil {
		return
	}
	return s.delete(ctx, name, opt)
}
func (s *Service) Get(name string, pairs ...Pair) (store Storager, err error) {
	ctx := context.Background()
	return s.GetWithContext(ctx, name, pairs...)
}
func (s *Service) GetWithContext(ctx context.Context, name string, pairs ...Pair) (store Storager, err error) {
	defer func() {
		err =
			s.formatError("get", err, name)
	}()

	pairs = append(pairs, s.defaultPairs.Get...)
	var opt pairServiceGet

	opt, err = s.parsePairServiceGet(pairs)
	if err != nil {
		return
	}
	return s.get(ctx, name, opt)
}
func (s *Service) List(pairs ...Pair) (sti *StoragerIterator, err error) {
	ctx := context.Background()
	return s.ListWithContext(ctx, pairs...)
}
func (s *Service) ListWithContext(ctx context.Context, pairs ...Pair) (sti *StoragerIterator, err error) {
	defer func() {
		err =
			s.formatError("list", err, "")
	}()

	pairs = append(pairs, s.defaultPairs.List...)
	var opt pairServiceList

	opt, err = s.parsePairServiceList(pairs)
	if err != nil {
		return
	}
	return s.list(ctx, opt)
}

var (
	_ Direr       = &Storage{}
	_ Multiparter = &Storage{}
	_ Storager    = &Storage{}
)

type StorageFeatures struct { // virtual_dir feature is designed for a service that doesn't have native dir support but wants to
	// provide simulated operations.
	//
	// - If this feature is disabled (the default behavior), the service will behave like it doesn't have
	// any dir support.
	// - If this feature is enabled, the service will support simulated dir behavior in create_dir, create,
	// list, delete, and so on.
	//
	// This feature was introduced in GSP-109.
	VirtualDir bool
}

// pairStorageNew is the parsed struct
type pairStorageNew struct {
	pairs []Pair

	// Required pairs
	HasLocation bool
	Location    string
	HasName     bool
	Name        string
	// Optional pairs
	HasDefaultContentType  bool
	DefaultContentType     string
	HasDefaultIoCallback   bool
	DefaultIoCallback      func([]byte)
	HasDefaultStoragePairs bool
	DefaultStoragePairs    DefaultStoragePairs
	HasStorageFeatures     bool
	StorageFeatures        StorageFeatures
	HasWorkDir             bool
	WorkDir                string
	// Enable features
	hasEnableVirtualDir bool
	EnableVirtualDir    bool
}

// parsePairStorageNew will parse Pair slice into *pairStorageNew
func parsePairStorageNew(opts []Pair) (pairStorageNew, error) {
	result :=
		pairStorageNew{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		case "location":
			if result.HasLocation {
				continue
			}
			result.HasLocation = true
			result.Location = v.Value.(string)
		case "name":
			if result.HasName {
				continue
			}
			result.HasName = true
			result.Name = v.Value.(string)
		case "default_content_type":
			if result.HasDefaultContentType {
				continue
			}
			result.HasDefaultContentType = true
			result.DefaultContentType = v.Value.(string)
		case "default_io_callback":
			if result.HasDefaultIoCallback {
				continue
			}
			result.HasDefaultIoCallback = true
			result.DefaultIoCallback = v.Value.(func([]byte))
		case "default_storage_pairs":
			if result.HasDefaultStoragePairs {
				continue
			}
			result.HasDefaultStoragePairs = true
			result.DefaultStoragePairs = v.Value.(DefaultStoragePairs)
		case "storage_features":
			if result.HasStorageFeatures {
				continue
			}
			result.HasStorageFeatures = true
			result.StorageFeatures = v.Value.(StorageFeatures)
		case "work_dir":
			if result.HasWorkDir {
				continue
			}
			result.HasWorkDir = true
			result.WorkDir = v.Value.(string)
		case "enable_virtual_dir":
			if result.hasEnableVirtualDir {
				continue
			}
			result.hasEnableVirtualDir = true
			result.EnableVirtualDir = true
		}
	}
	// Enable features
	if result.hasEnableVirtualDir {
		result.HasStorageFeatures = true
		result.StorageFeatures.VirtualDir = true
	}
	// Default pairs
	if result.HasDefaultContentType {
		result.HasDefaultStoragePairs = true
		result.DefaultStoragePairs.CreateMultipart = append(result.DefaultStoragePairs.CreateMultipart, WithContentType(result.DefaultContentType))
		result.DefaultStoragePairs.Write = append(result.DefaultStoragePairs.Write, WithContentType(result.DefaultContentType))
	}
	if result.HasDefaultIoCallback {
		result.HasDefaultStoragePairs = true
		result.DefaultStoragePairs.Read = append(result.DefaultStoragePairs.Read, WithIoCallback(result.DefaultIoCallback))
		result.DefaultStoragePairs.Write = append(result.DefaultStoragePairs.Write, WithIoCallback(result.DefaultIoCallback))
	}
	if !result.HasLocation {
		return pairStorageNew{}, services.PairRequiredError{Keys: []string{"location"}}
	}
	if !result.HasName {
		return pairStorageNew{}, services.PairRequiredError{Keys: []string{"name"}}
	}
	return result, nil
}

// DefaultStoragePairs is default pairs for specific action
type DefaultStoragePairs struct {
	CompleteMultipart []Pair
	Create            []Pair
	CreateDir         []Pair
	CreateMultipart   []Pair
	Delete            []Pair
	List              []Pair
	ListMultipart     []Pair
	Metadata          []Pair
	Read              []Pair
	Stat              []Pair
	Write             []Pair
	WriteMultipart    []Pair
}
type pairStorageCompleteMultipart struct {
	pairs []Pair
	// Required pairs
	// Optional pairs
}

func (s *Storage) parsePairStorageCompleteMultipart(opts []Pair) (pairStorageCompleteMultipart, error) {
	result :=
		pairStorageCompleteMultipart{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		default:
			return pairStorageCompleteMultipart{}, services.PairUnsupportedError{Pair: v}
		}
	}

	return result, nil
}

type pairStorageCreate struct {
	pairs []Pair
	// Required pairs
	// Optional pairs
	HasMultipartID bool
	MultipartID    string
	HasObjectMode  bool
	ObjectMode     ObjectMode
}

func (s *Storage) parsePairStorageCreate(opts []Pair) (pairStorageCreate, error) {
	result :=
		pairStorageCreate{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		case "multipart_id":
			if result.HasMultipartID {
				continue
			}
			result.HasMultipartID = true
			result.MultipartID = v.Value.(string)
		case "object_mode":
			if result.HasObjectMode {
				continue
			}
			result.HasObjectMode = true
			result.ObjectMode = v.Value.(ObjectMode)
		default:
			return pairStorageCreate{}, services.PairUnsupportedError{Pair: v}
		}
	}

	return result, nil
}

type pairStorageCreateDir struct {
	pairs []Pair
	// Required pairs
	// Optional pairs
	HasStorageClass bool
	StorageClass    string
}

func (s *Storage) parsePairStorageCreateDir(opts []Pair) (pairStorageCreateDir, error) {
	result :=
		pairStorageCreateDir{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		case "storage_class":
			if result.HasStorageClass {
				continue
			}
			result.HasStorageClass = true
			result.StorageClass = v.Value.(string)
		default:
			return pairStorageCreateDir{}, services.PairUnsupportedError{Pair: v}
		}
	}

	return result, nil
}

type pairStorageCreateMultipart struct {
	pairs []Pair
	// Required pairs
	// Optional pairs
	HasContentType                           bool
	ContentType                              string
	HasServerSideEncryption                  bool
	ServerSideEncryption                     string
	HasServerSideEncryptionContext           bool
	ServerSideEncryptionContext              string
	HasServerSideEncryptionCosKmsKeyID       bool
	ServerSideEncryptionCosKmsKeyID          string
	HasServerSideEncryptionCustomerAlgorithm bool
	ServerSideEncryptionCustomerAlgorithm    string
	HasServerSideEncryptionCustomerKey       bool
	ServerSideEncryptionCustomerKey          []byte
	HasStorageClass                          bool
	StorageClass                             string
}

func (s *Storage) parsePairStorageCreateMultipart(opts []Pair) (pairStorageCreateMultipart, error) {
	result :=
		pairStorageCreateMultipart{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		case "content_type":
			if result.HasContentType {
				continue
			}
			result.HasContentType = true
			result.ContentType = v.Value.(string)
		case "server_side_encryption":
			if result.HasServerSideEncryption {
				continue
			}
			result.HasServerSideEncryption = true
			result.ServerSideEncryption = v.Value.(string)
		case "server_side_encryption_context":
			if result.HasServerSideEncryptionContext {
				continue
			}
			result.HasServerSideEncryptionContext = true
			result.ServerSideEncryptionContext = v.Value.(string)
		case "server_side_encryption_cos_kms_key_id":
			if result.HasServerSideEncryptionCosKmsKeyID {
				continue
			}
			result.HasServerSideEncryptionCosKmsKeyID = true
			result.ServerSideEncryptionCosKmsKeyID = v.Value.(string)
		case "server_side_encryption_customer_algorithm":
			if result.HasServerSideEncryptionCustomerAlgorithm {
				continue
			}
			result.HasServerSideEncryptionCustomerAlgorithm = true
			result.ServerSideEncryptionCustomerAlgorithm = v.Value.(string)
		case "server_side_encryption_customer_key":
			if result.HasServerSideEncryptionCustomerKey {
				continue
			}
			result.HasServerSideEncryptionCustomerKey = true
			result.ServerSideEncryptionCustomerKey = v.Value.([]byte)
		case "storage_class":
			if result.HasStorageClass {
				continue
			}
			result.HasStorageClass = true
			result.StorageClass = v.Value.(string)
		default:
			return pairStorageCreateMultipart{}, services.PairUnsupportedError{Pair: v}
		}
	}

	return result, nil
}

type pairStorageDelete struct {
	pairs []Pair
	// Required pairs
	// Optional pairs
	HasMultipartID bool
	MultipartID    string
	HasObjectMode  bool
	ObjectMode     ObjectMode
}

func (s *Storage) parsePairStorageDelete(opts []Pair) (pairStorageDelete, error) {
	result :=
		pairStorageDelete{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		case "multipart_id":
			if result.HasMultipartID {
				continue
			}
			result.HasMultipartID = true
			result.MultipartID = v.Value.(string)
		case "object_mode":
			if result.HasObjectMode {
				continue
			}
			result.HasObjectMode = true
			result.ObjectMode = v.Value.(ObjectMode)
		default:
			return pairStorageDelete{}, services.PairUnsupportedError{Pair: v}
		}
	}

	return result, nil
}

type pairStorageList struct {
	pairs []Pair
	// Required pairs
	// Optional pairs
	HasListMode bool
	ListMode    ListMode
}

func (s *Storage) parsePairStorageList(opts []Pair) (pairStorageList, error) {
	result :=
		pairStorageList{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		case "list_mode":
			if result.HasListMode {
				continue
			}
			result.HasListMode = true
			result.ListMode = v.Value.(ListMode)
		default:
			return pairStorageList{}, services.PairUnsupportedError{Pair: v}
		}
	}

	return result, nil
}

type pairStorageListMultipart struct {
	pairs []Pair
	// Required pairs
	// Optional pairs
}

func (s *Storage) parsePairStorageListMultipart(opts []Pair) (pairStorageListMultipart, error) {
	result :=
		pairStorageListMultipart{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		default:
			return pairStorageListMultipart{}, services.PairUnsupportedError{Pair: v}
		}
	}

	return result, nil
}

type pairStorageMetadata struct {
	pairs []Pair
	// Required pairs
	// Optional pairs
}

func (s *Storage) parsePairStorageMetadata(opts []Pair) (pairStorageMetadata, error) {
	result :=
		pairStorageMetadata{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		default:
			return pairStorageMetadata{}, services.PairUnsupportedError{Pair: v}
		}
	}

	return result, nil
}

type pairStorageRead struct {
	pairs []Pair
	// Required pairs
	// Optional pairs
	HasIoCallback                            bool
	IoCallback                               func([]byte)
	HasOffset                                bool
	Offset                                   int64
	HasServerSideEncryptionCustomerAlgorithm bool
	ServerSideEncryptionCustomerAlgorithm    string
	HasServerSideEncryptionCustomerKey       bool
	ServerSideEncryptionCustomerKey          []byte
	HasSize                                  bool
	Size                                     int64
}

func (s *Storage) parsePairStorageRead(opts []Pair) (pairStorageRead, error) {
	result :=
		pairStorageRead{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		case "io_callback":
			if result.HasIoCallback {
				continue
			}
			result.HasIoCallback = true
			result.IoCallback = v.Value.(func([]byte))
		case "offset":
			if result.HasOffset {
				continue
			}
			result.HasOffset = true
			result.Offset = v.Value.(int64)
		case "server_side_encryption_customer_algorithm":
			if result.HasServerSideEncryptionCustomerAlgorithm {
				continue
			}
			result.HasServerSideEncryptionCustomerAlgorithm = true
			result.ServerSideEncryptionCustomerAlgorithm = v.Value.(string)
		case "server_side_encryption_customer_key":
			if result.HasServerSideEncryptionCustomerKey {
				continue
			}
			result.HasServerSideEncryptionCustomerKey = true
			result.ServerSideEncryptionCustomerKey = v.Value.([]byte)
		case "size":
			if result.HasSize {
				continue
			}
			result.HasSize = true
			result.Size = v.Value.(int64)
		default:
			return pairStorageRead{}, services.PairUnsupportedError{Pair: v}
		}
	}

	return result, nil
}

type pairStorageStat struct {
	pairs []Pair
	// Required pairs
	// Optional pairs
	HasMultipartID                           bool
	MultipartID                              string
	HasObjectMode                            bool
	ObjectMode                               ObjectMode
	HasServerSideEncryptionCustomerAlgorithm bool
	ServerSideEncryptionCustomerAlgorithm    string
	HasServerSideEncryptionCustomerKey       bool
	ServerSideEncryptionCustomerKey          []byte
}

func (s *Storage) parsePairStorageStat(opts []Pair) (pairStorageStat, error) {
	result :=
		pairStorageStat{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		case "multipart_id":
			if result.HasMultipartID {
				continue
			}
			result.HasMultipartID = true
			result.MultipartID = v.Value.(string)
		case "object_mode":
			if result.HasObjectMode {
				continue
			}
			result.HasObjectMode = true
			result.ObjectMode = v.Value.(ObjectMode)
		case "server_side_encryption_customer_algorithm":
			if result.HasServerSideEncryptionCustomerAlgorithm {
				continue
			}
			result.HasServerSideEncryptionCustomerAlgorithm = true
			result.ServerSideEncryptionCustomerAlgorithm = v.Value.(string)
		case "server_side_encryption_customer_key":
			if result.HasServerSideEncryptionCustomerKey {
				continue
			}
			result.HasServerSideEncryptionCustomerKey = true
			result.ServerSideEncryptionCustomerKey = v.Value.([]byte)
		default:
			return pairStorageStat{}, services.PairUnsupportedError{Pair: v}
		}
	}

	return result, nil
}

type pairStorageWrite struct {
	pairs []Pair
	// Required pairs
	// Optional pairs
	HasContentMd5                            bool
	ContentMd5                               string
	HasContentType                           bool
	ContentType                              string
	HasIoCallback                            bool
	IoCallback                               func([]byte)
	HasServerSideEncryption                  bool
	ServerSideEncryption                     string
	HasServerSideEncryptionContext           bool
	ServerSideEncryptionContext              string
	HasServerSideEncryptionCosKmsKeyID       bool
	ServerSideEncryptionCosKmsKeyID          string
	HasServerSideEncryptionCustomerAlgorithm bool
	ServerSideEncryptionCustomerAlgorithm    string
	HasServerSideEncryptionCustomerKey       bool
	ServerSideEncryptionCustomerKey          []byte
	HasStorageClass                          bool
	StorageClass                             string
}

func (s *Storage) parsePairStorageWrite(opts []Pair) (pairStorageWrite, error) {
	result :=
		pairStorageWrite{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		case "content_md5":
			if result.HasContentMd5 {
				continue
			}
			result.HasContentMd5 = true
			result.ContentMd5 = v.Value.(string)
		case "content_type":
			if result.HasContentType {
				continue
			}
			result.HasContentType = true
			result.ContentType = v.Value.(string)
		case "io_callback":
			if result.HasIoCallback {
				continue
			}
			result.HasIoCallback = true
			result.IoCallback = v.Value.(func([]byte))
		case "server_side_encryption":
			if result.HasServerSideEncryption {
				continue
			}
			result.HasServerSideEncryption = true
			result.ServerSideEncryption = v.Value.(string)
		case "server_side_encryption_context":
			if result.HasServerSideEncryptionContext {
				continue
			}
			result.HasServerSideEncryptionContext = true
			result.ServerSideEncryptionContext = v.Value.(string)
		case "server_side_encryption_cos_kms_key_id":
			if result.HasServerSideEncryptionCosKmsKeyID {
				continue
			}
			result.HasServerSideEncryptionCosKmsKeyID = true
			result.ServerSideEncryptionCosKmsKeyID = v.Value.(string)
		case "server_side_encryption_customer_algorithm":
			if result.HasServerSideEncryptionCustomerAlgorithm {
				continue
			}
			result.HasServerSideEncryptionCustomerAlgorithm = true
			result.ServerSideEncryptionCustomerAlgorithm = v.Value.(string)
		case "server_side_encryption_customer_key":
			if result.HasServerSideEncryptionCustomerKey {
				continue
			}
			result.HasServerSideEncryptionCustomerKey = true
			result.ServerSideEncryptionCustomerKey = v.Value.([]byte)
		case "storage_class":
			if result.HasStorageClass {
				continue
			}
			result.HasStorageClass = true
			result.StorageClass = v.Value.(string)
		default:
			return pairStorageWrite{}, services.PairUnsupportedError{Pair: v}
		}
	}

	return result, nil
}

type pairStorageWriteMultipart struct {
	pairs []Pair
	// Required pairs
	// Optional pairs
	HasContentMd5 bool
	ContentMd5    string
}

func (s *Storage) parsePairStorageWriteMultipart(opts []Pair) (pairStorageWriteMultipart, error) {
	result :=
		pairStorageWriteMultipart{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		case "content_md5":
			if result.HasContentMd5 {
				continue
			}
			result.HasContentMd5 = true
			result.ContentMd5 = v.Value.(string)
		default:
			return pairStorageWriteMultipart{}, services.PairUnsupportedError{Pair: v}
		}
	}

	return result, nil
}
func (s *Storage) CompleteMultipart(o *Object, parts []*Part, pairs ...Pair) (err error) {
	ctx := context.Background()
	return s.CompleteMultipartWithContext(ctx, o, parts, pairs...)
}
func (s *Storage) CompleteMultipartWithContext(ctx context.Context, o *Object, parts []*Part, pairs ...Pair) (err error) {
	defer func() {
		err =
			s.formatError("complete_multipart", err)
	}()
	if !o.Mode.IsPart() {
		err = services.ObjectModeInvalidError{Expected: ModePart, Actual: o.Mode}
		return
	}
	pairs = append(pairs, s.defaultPairs.CompleteMultipart...)
	var opt pairStorageCompleteMultipart

	opt, err = s.parsePairStorageCompleteMultipart(pairs)
	if err != nil {
		return
	}
	return s.completeMultipart(ctx, o, parts, opt)
}
func (s *Storage) Create(path string, pairs ...Pair) (o *Object) {
	pairs = append(pairs, s.defaultPairs.Create...)
	var opt pairStorageCreate

	// Ignore error while handling local functions.
	opt, _ = s.parsePairStorageCreate(pairs)
	return s.create(path, opt)
}
func (s *Storage) CreateDir(path string, pairs ...Pair) (o *Object, err error) {
	ctx := context.Background()
	return s.CreateDirWithContext(ctx, path, pairs...)
}
func (s *Storage) CreateDirWithContext(ctx context.Context, path string, pairs ...Pair) (o *Object, err error) {
	defer func() {
		err =
			s.formatError("create_dir", err, path)
	}()

	pairs = append(pairs, s.defaultPairs.CreateDir...)
	var opt pairStorageCreateDir

	opt, err = s.parsePairStorageCreateDir(pairs)
	if err != nil {
		return
	}
	return s.createDir(ctx, strings.ReplaceAll(path, "\\", "/"), opt)
}
func (s *Storage) CreateMultipart(path string, pairs ...Pair) (o *Object, err error) {
	ctx := context.Background()
	return s.CreateMultipartWithContext(ctx, path, pairs...)
}
func (s *Storage) CreateMultipartWithContext(ctx context.Context, path string, pairs ...Pair) (o *Object, err error) {
	defer func() {
		err =
			s.formatError("create_multipart", err, path)
	}()

	pairs = append(pairs, s.defaultPairs.CreateMultipart...)
	var opt pairStorageCreateMultipart

	opt, err = s.parsePairStorageCreateMultipart(pairs)
	if err != nil {
		return
	}
	return s.createMultipart(ctx, strings.ReplaceAll(path, "\\", "/"), opt)
}
func (s *Storage) Delete(path string, pairs ...Pair) (err error) {
	ctx := context.Background()
	return s.DeleteWithContext(ctx, path, pairs...)
}
func (s *Storage) DeleteWithContext(ctx context.Context, path string, pairs ...Pair) (err error) {
	defer func() {
		err =
			s.formatError("delete", err, path)
	}()

	pairs = append(pairs, s.defaultPairs.Delete...)
	var opt pairStorageDelete

	opt, err = s.parsePairStorageDelete(pairs)
	if err != nil {
		return
	}
	return s.delete(ctx, strings.ReplaceAll(path, "\\", "/"), opt)
}
func (s *Storage) List(path string, pairs ...Pair) (oi *ObjectIterator, err error) {
	ctx := context.Background()
	return s.ListWithContext(ctx, path, pairs...)
}
func (s *Storage) ListWithContext(ctx context.Context, path string, pairs ...Pair) (oi *ObjectIterator, err error) {
	defer func() {
		err =
			s.formatError("list", err, path)
	}()

	pairs = append(pairs, s.defaultPairs.List...)
	var opt pairStorageList

	opt, err = s.parsePairStorageList(pairs)
	if err != nil {
		return
	}
	return s.list(ctx, strings.ReplaceAll(path, "\\", "/"), opt)
}
func (s *Storage) ListMultipart(o *Object, pairs ...Pair) (pi *PartIterator, err error) {
	ctx := context.Background()
	return s.ListMultipartWithContext(ctx, o, pairs...)
}
func (s *Storage) ListMultipartWithContext(ctx context.Context, o *Object, pairs ...Pair) (pi *PartIterator, err error) {
	defer func() {
		err =
			s.formatError("list_multipart", err)
	}()
	if !o.Mode.IsPart() {
		err = services.ObjectModeInvalidError{Expected: ModePart, Actual: o.Mode}
		return
	}
	pairs = append(pairs, s.defaultPairs.ListMultipart...)
	var opt pairStorageListMultipart

	opt, err = s.parsePairStorageListMultipart(pairs)
	if err != nil {
		return
	}
	return s.listMultipart(ctx, o, opt)
}
func (s *Storage) Metadata(pairs ...Pair) (meta *StorageMeta) {
	pairs = append(pairs, s.defaultPairs.Metadata...)
	var opt pairStorageMetadata

	// Ignore error while handling local functions.
	opt, _ = s.parsePairStorageMetadata(pairs)
	return s.metadata(opt)
}
func (s *Storage) Read(path string, w io.Writer, pairs ...Pair) (n int64, err error) {
	ctx := context.Background()
	return s.ReadWithContext(ctx, path, w, pairs...)
}
func (s *Storage) ReadWithContext(ctx context.Context, path string, w io.Writer, pairs ...Pair) (n int64, err error) {
	defer func() {
		err =
			s.formatError("read", err, path)
	}()

	pairs = append(pairs, s.defaultPairs.Read...)
	var opt pairStorageRead

	opt, err = s.parsePairStorageRead(pairs)
	if err != nil {
		return
	}
	return s.read(ctx, strings.ReplaceAll(path, "\\", "/"), w, opt)
}
func (s *Storage) Stat(path string, pairs ...Pair) (o *Object, err error) {
	ctx := context.Background()
	return s.StatWithContext(ctx, path, pairs...)
}
func (s *Storage) StatWithContext(ctx context.Context, path string, pairs ...Pair) (o *Object, err error) {
	defer func() {
		err =
			s.formatError("stat", err, path)
	}()

	pairs = append(pairs, s.defaultPairs.Stat...)
	var opt pairStorageStat

	opt, err = s.parsePairStorageStat(pairs)
	if err != nil {
		return
	}
	return s.stat(ctx, strings.ReplaceAll(path, "\\", "/"), opt)
}
func (s *Storage) Write(path string, r io.Reader, size int64, pairs ...Pair) (n int64, err error) {
	ctx := context.Background()
	return s.WriteWithContext(ctx, path, r, size, pairs...)
}
func (s *Storage) WriteWithContext(ctx context.Context, path string, r io.Reader, size int64, pairs ...Pair) (n int64, err error) {
	defer func() {
		err =
			s.formatError("write", err, path)
	}()

	pairs = append(pairs, s.defaultPairs.Write...)
	var opt pairStorageWrite

	opt, err = s.parsePairStorageWrite(pairs)
	if err != nil {
		return
	}
	return s.write(ctx, strings.ReplaceAll(path, "\\", "/"), r, size, opt)
}
func (s *Storage) WriteMultipart(o *Object, r io.Reader, size int64, index int, pairs ...Pair) (n int64, part *Part, err error) {
	ctx := context.Background()
	return s.WriteMultipartWithContext(ctx, o, r, size, index, pairs...)
}
func (s *Storage) WriteMultipartWithContext(ctx context.Context, o *Object, r io.Reader, size int64, index int, pairs ...Pair) (n int64, part *Part, err error) {
	defer func() {
		err =
			s.formatError("write_multipart", err)
	}()
	if !o.Mode.IsPart() {
		err = services.ObjectModeInvalidError{Expected: ModePart, Actual: o.Mode}
		return
	}
	pairs = append(pairs, s.defaultPairs.WriteMultipart...)
	var opt pairStorageWriteMultipart

	opt, err = s.parsePairStorageWriteMultipart(pairs)
	if err != nil {
		return
	}
	return s.writeMultipart(ctx, o, r, size, index, opt)
}
func init() {
	services.RegisterServicer(Type, NewServicer)
	services.RegisterStorager(Type, NewStorager)
	services.RegisterSchema(Type, pairMap)
}
