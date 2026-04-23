// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package files

import (
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// FileService contains methods and other services that help with interacting with
// the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileService] method instead.
type FileService struct {
	options    []option.RequestOption
	FileAssets FileAssetService
	Folders    FolderService
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r FileService) {
	r = FileService{}
	r.options = opts
	r.FileAssets = NewFileAssetService(opts...)
	r.Folders = NewFolderService(opts...)
	return
}

type CollectionResponseFile struct {
	Results []File        `json:"results" api:"required"`
	Paging  shared.Paging `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseFile) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseFolder struct {
	Results []Folder      `json:"results" api:"required"`
	Paging  shared.Paging `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseFolder) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type File struct {
	// File ID.
	ID string `json:"id" api:"required"`
	// If the file is deleted.
	Archived bool `json:"archived" api:"required"`
	// Creation time of the file object.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Timestamp of the latest update to the file.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// File access. Can be PUBLIC_INDEXABLE, PUBLIC_NOT_INDEXABLE, PRIVATE.
	//
	// Any of "HIDDEN_INDEXABLE", "HIDDEN_NOT_INDEXABLE", "HIDDEN_PRIVATE",
	// "HIDDEN_SENSITIVE", "PRIVATE", "PUBLIC_INDEXABLE", "PUBLIC_NOT_INDEXABLE",
	// "SENSITIVE".
	Access FileAccess `json:"access"`
	// Deletion time of the file object.
	ArchivedAt time.Time `json:"archivedAt" format:"date-time"`
	// Default hosting URL of the file. This will use one of HubSpot's provided URLs to
	// serve the file.
	DefaultHostingURL string `json:"defaultHostingUrl"`
	// Encoding of the file.
	Encoding  string `json:"encoding"`
	ExpiresAt int64  `json:"expiresAt"`
	// Extension of the file. ex: .jpg, .png, .gif, .pdf, etc.
	Extension string `json:"extension"`
	// The MD5 hash of the file.
	FileMd5 string `json:"fileMd5"`
	// For image and video files, the height of the content.
	Height int64 `json:"height"`
	// Previously "archied". Indicates if the file should be used when creating new
	// content like web pages.
	IsUsableInContent bool `json:"isUsableInContent"`
	// Name of the file.
	Name string `json:"name"`
	// ID of the folder the file is in.
	ParentFolderID string `json:"parentFolderId"`
	// Path of the file in the file manager.
	Path string `json:"path"`
	// Size of the file in bytes.
	Size int64 `json:"size"`
	// Any of "CONTENT", "CONVERSATIONS", "FORMS", "UI_EXTENSIONS", "UNKNOWN".
	SourceGroup FileSourceGroup `json:"sourceGroup"`
	// Type of the file. Can be IMG, DOCUMENT, AUDIO, MOVIE, or OTHER.
	Type string `json:"type"`
	// URL of the given file. This URL can change depending on the domain settings of
	// the account. Will use the select file hosting domain.
	URL string `json:"url"`
	// For image and video files, the width of the content.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Archived          respjson.Field
		CreatedAt         respjson.Field
		UpdatedAt         respjson.Field
		Access            respjson.Field
		ArchivedAt        respjson.Field
		DefaultHostingURL respjson.Field
		Encoding          respjson.Field
		ExpiresAt         respjson.Field
		Extension         respjson.Field
		FileMd5           respjson.Field
		Height            respjson.Field
		IsUsableInContent respjson.Field
		Name              respjson.Field
		ParentFolderID    respjson.Field
		Path              respjson.Field
		Size              respjson.Field
		SourceGroup       respjson.Field
		Type              respjson.Field
		URL               respjson.Field
		Width             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r File) RawJSON() string { return r.JSON.raw }
func (r *File) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File access. Can be PUBLIC_INDEXABLE, PUBLIC_NOT_INDEXABLE, PRIVATE.
type FileAccess string

const (
	FileAccessHiddenIndexable    FileAccess = "HIDDEN_INDEXABLE"
	FileAccessHiddenNotIndexable FileAccess = "HIDDEN_NOT_INDEXABLE"
	FileAccessHiddenPrivate      FileAccess = "HIDDEN_PRIVATE"
	FileAccessHiddenSensitive    FileAccess = "HIDDEN_SENSITIVE"
	FileAccessPrivate            FileAccess = "PRIVATE"
	FileAccessPublicIndexable    FileAccess = "PUBLIC_INDEXABLE"
	FileAccessPublicNotIndexable FileAccess = "PUBLIC_NOT_INDEXABLE"
	FileAccessSensitive          FileAccess = "SENSITIVE"
)

type FileSourceGroup string

const (
	FileSourceGroupContent       FileSourceGroup = "CONTENT"
	FileSourceGroupConversations FileSourceGroup = "CONVERSATIONS"
	FileSourceGroupForms         FileSourceGroup = "FORMS"
	FileSourceGroupUiExtensions  FileSourceGroup = "UI_EXTENSIONS"
	FileSourceGroupUnknown       FileSourceGroup = "UNKNOWN"
)

type FileActionResponse struct {
	// Time of completion of task.
	CompletedAt time.Time `json:"completedAt" api:"required" format:"date-time"`
	// Timestamp of when the task was started.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// Current status of the task.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status FileActionResponseStatus `json:"status" api:"required"`
	// ID of the requested task.
	TaskID string `json:"taskId" api:"required"`
	// Descriptive error messages.
	Errors []shared.StandardError `json:"errors"`
	// Link to check the status of the requested task.
	Links map[string]string `json:"links"`
	// Number of errors resulting from the task.
	NumErrors int64 `json:"numErrors"`
	// Timestamp of when the task was requested.
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
	Result      File      `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		TaskID      respjson.Field
		Errors      respjson.Field
		Links       respjson.Field
		NumErrors   respjson.Field
		RequestedAt respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileActionResponse) RawJSON() string { return r.JSON.raw }
func (r *FileActionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the task.
type FileActionResponseStatus string

const (
	FileActionResponseStatusCanceled   FileActionResponseStatus = "CANCELED"
	FileActionResponseStatusComplete   FileActionResponseStatus = "COMPLETE"
	FileActionResponseStatusPending    FileActionResponseStatus = "PENDING"
	FileActionResponseStatusProcessing FileActionResponseStatus = "PROCESSING"
)

// The property ClearExpires is required.
type FileUpdateInputParam struct {
	ClearExpires bool                 `json:"clearExpires" api:"required"`
	ExpiresAt    param.Opt[time.Time] `json:"expiresAt,omitzero" format:"date-time"`
	// Mark whether the file should be used in new content or not.
	IsUsableInContent param.Opt[bool] `json:"isUsableInContent,omitzero"`
	// New name for the file.
	Name param.Opt[string] `json:"name,omitzero"`
	// FolderId where the file should be moved to. folderId and folderPath parameters
	// cannot be set at the same time.
	ParentFolderID param.Opt[string] `json:"parentFolderId,omitzero"`
	// Folder path where the file should be moved to. folderId and folderPath
	// parameters cannot be set at the same time.
	ParentFolderPath param.Opt[string] `json:"parentFolderPath,omitzero"`
	// NONE: Do not run any duplicate validation. REJECT: Reject the upload if a
	// duplicate is found. RETURN_EXISTING: If a duplicate file is found, do not upload
	// a new file and return the found duplicate instead.
	//
	// Any of "HIDDEN_INDEXABLE", "HIDDEN_NOT_INDEXABLE", "HIDDEN_PRIVATE",
	// "HIDDEN_SENSITIVE", "PRIVATE", "PUBLIC_INDEXABLE", "PUBLIC_NOT_INDEXABLE",
	// "SENSITIVE".
	Access FileUpdateInputAccess `json:"access,omitzero"`
	paramObj
}

func (r FileUpdateInputParam) MarshalJSON() (data []byte, err error) {
	type shadow FileUpdateInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUpdateInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NONE: Do not run any duplicate validation. REJECT: Reject the upload if a
// duplicate is found. RETURN_EXISTING: If a duplicate file is found, do not upload
// a new file and return the found duplicate instead.
type FileUpdateInputAccess string

const (
	FileUpdateInputAccessHiddenIndexable    FileUpdateInputAccess = "HIDDEN_INDEXABLE"
	FileUpdateInputAccessHiddenNotIndexable FileUpdateInputAccess = "HIDDEN_NOT_INDEXABLE"
	FileUpdateInputAccessHiddenPrivate      FileUpdateInputAccess = "HIDDEN_PRIVATE"
	FileUpdateInputAccessHiddenSensitive    FileUpdateInputAccess = "HIDDEN_SENSITIVE"
	FileUpdateInputAccessPrivate            FileUpdateInputAccess = "PRIVATE"
	FileUpdateInputAccessPublicIndexable    FileUpdateInputAccess = "PUBLIC_INDEXABLE"
	FileUpdateInputAccessPublicNotIndexable FileUpdateInputAccess = "PUBLIC_NOT_INDEXABLE"
	FileUpdateInputAccessSensitive          FileUpdateInputAccess = "SENSITIVE"
)

type Folder struct {
	// ID of the folder.
	ID string `json:"id" api:"required"`
	// Marks whether the folder is deleted or not.
	Archived bool `json:"archived" api:"required"`
	// Timestamp of folder creation.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Timestamp of the latest update to the folder.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Timestamp of folder deletion.
	ArchivedAt time.Time `json:"archivedAt" format:"date-time"`
	// Name of the folder.
	Name string `json:"name"`
	// ID of the parent folder.
	ParentFolderID string `json:"parentFolderId"`
	// Path of the folder in the file manager.
	Path string `json:"path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Archived       respjson.Field
		CreatedAt      respjson.Field
		UpdatedAt      respjson.Field
		ArchivedAt     respjson.Field
		Name           respjson.Field
		ParentFolderID respjson.Field
		Path           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Folder) RawJSON() string { return r.JSON.raw }
func (r *Folder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderActionResponse struct {
	// When the requested changes have been completed.
	CompletedAt time.Time `json:"completedAt" api:"required" format:"date-time"`
	// Timestamp representing when the task was started at.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// Current status of the task.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status FolderActionResponseStatus `json:"status" api:"required"`
	// ID of the task.
	TaskID string `json:"taskId" api:"required"`
	// Detailed errors resulting from the task.
	Errors []shared.StandardError `json:"errors"`
	// Link to check the status of the task.
	Links map[string]string `json:"links"`
	// Number of errors resulting from the requested changes.
	NumErrors int64 `json:"numErrors"`
	// Timestamp representing when the task was requested.
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
	Result      Folder    `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		TaskID      respjson.Field
		Errors      respjson.Field
		Links       respjson.Field
		NumErrors   respjson.Field
		RequestedAt respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderActionResponse) RawJSON() string { return r.JSON.raw }
func (r *FolderActionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the task.
type FolderActionResponseStatus string

const (
	FolderActionResponseStatusCanceled   FolderActionResponseStatus = "CANCELED"
	FolderActionResponseStatusComplete   FolderActionResponseStatus = "COMPLETE"
	FolderActionResponseStatusPending    FolderActionResponseStatus = "PENDING"
	FolderActionResponseStatusProcessing FolderActionResponseStatus = "PROCESSING"
)

// The property Name is required.
type FolderInputParam struct {
	// Desired name for the folder.
	Name string `json:"name" api:"required"`
	// FolderId of the parent of the created folder. If not specified, the folder will
	// be created at the root level. parentFolderId and parentFolderPath cannot be set
	// at the same time.
	ParentFolderID param.Opt[string] `json:"parentFolderId,omitzero"`
	// Path of the parent of the created folder. If not specified the folder will be
	// created at the root level. parentFolderPath and parentFolderId cannot be set at
	// the same time.
	ParentPath param.Opt[string] `json:"parentPath,omitzero"`
	paramObj
}

func (r FolderInputParam) MarshalJSON() (data []byte, err error) {
	type shadow FolderInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FolderInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderUpdateInputParam struct {
	// New name. If specified the folder's name and fullPath will change. All children
	// of the folder will be updated accordingly.
	Name param.Opt[string] `json:"name,omitzero"`
	// New parent folderId. If changed, the folder and all it's children will be moved
	// into the specified folder. parentFolderId and parentFolderPath cannot be
	// specified at the same time.
	ParentFolderID param.Opt[int64] `json:"parentFolderId,omitzero"`
	paramObj
}

func (r FolderUpdateInputParam) MarshalJSON() (data []byte, err error) {
	type shadow FolderUpdateInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FolderUpdateInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type FolderUpdateInputWithIDParam struct {
	// The unique identifier of the folder to be updated.
	ID string `json:"id" api:"required"`
	// New name. If specified the folder's name and fullPath will change. All children
	// of the folder will be updated accordingly.
	Name param.Opt[string] `json:"name,omitzero"`
	// New parent folderId. If changed, the folder and all it's children will be moved
	// into the specified folder. parentFolderId and parentFolderPath cannot be
	// specified at the same time.
	ParentFolderID param.Opt[int64] `json:"parentFolderId,omitzero"`
	paramObj
}

func (r FolderUpdateInputWithIDParam) MarshalJSON() (data []byte, err error) {
	type shadow FolderUpdateInputWithIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FolderUpdateInputWithIDParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderUpdateTaskLocator struct {
	// ID of the task
	ID string `json:"id" api:"required"`
	// Links for where to check information related to the task. The `status` link
	// gives the URL for where to check the status of the task.
	Links map[string]string `json:"links"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderUpdateTaskLocator) RawJSON() string { return r.JSON.raw }
func (r *FolderUpdateTaskLocator) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Access, DuplicateValidationScope, DuplicateValidationStrategy,
// Overwrite are required.
type ImportFromURLInputParam struct {
	// PUBLIC_INDEXABLE: File is publicly accessible by anyone who has the URL. Search
	// engines can index the file. PUBLIC_NOT_INDEXABLE: File is publicly accessible by
	// anyone who has the URL. Search engines _can't_ index the file. PRIVATE: File is
	// NOT publicly accessible. Requires a signed URL to see content. Search engines
	// _can't_ index the file.
	//
	// Any of "HIDDEN_INDEXABLE", "HIDDEN_NOT_INDEXABLE", "HIDDEN_PRIVATE",
	// "HIDDEN_SENSITIVE", "PRIVATE", "PUBLIC_INDEXABLE", "PUBLIC_NOT_INDEXABLE",
	// "SENSITIVE".
	Access ImportFromURLInputAccess `json:"access,omitzero" api:"required"`
	// ENTIRE_PORTAL: Look for a duplicate file in the entire account. EXACT_FOLDER:
	// Look for a duplicate file in the provided folder.
	//
	// Any of "ENTIRE_PORTAL", "EXACT_FOLDER".
	DuplicateValidationScope ImportFromURLInputDuplicateValidationScope `json:"duplicateValidationScope,omitzero" api:"required"`
	// NONE: Do not run any duplicate validation. REJECT: Reject the upload if a
	// duplicate is found. RETURN_EXISTING: If a duplicate file is found, do not upload
	// a new file and return the found duplicate instead.
	//
	// Any of "NONE", "REJECT", "RETURN_EXISTING".
	DuplicateValidationStrategy ImportFromURLInputDuplicateValidationStrategy `json:"duplicateValidationStrategy,omitzero" api:"required"`
	// If true, will overwrite existing file if one with the same name and extension
	// exists in the given folder. The overwritten file will be deleted and the
	// uploaded file will take its place with a new ID. If unset or set as false, the
	// new file's name will be updated to prevent colliding with existing file if one
	// exists with the same path, name, and extension
	Overwrite bool `json:"overwrite" api:"required"`
	// Specifies the date and time when the file will expire.
	ExpiresAt param.Opt[time.Time] `json:"expiresAt,omitzero" format:"date-time"`
	// One of folderId or folderPath is required. Destination folderId for the uploaded
	// file.
	FolderID param.Opt[string] `json:"folderId,omitzero"`
	// One of folderPath or folderId is required. Destination folder path for the
	// uploaded file. If the folder path does not exist, there will be an attempt to
	// create the folder path.
	FolderPath param.Opt[string] `json:"folderPath,omitzero"`
	// Name to give the resulting file in the file manager.
	Name param.Opt[string] `json:"name,omitzero"`
	// Time to live. If specified the file will be deleted after the given time frame.
	// If left unset, the file will exist indefinitely
	Ttl param.Opt[string] `json:"ttl,omitzero"`
	// URL to download the new file from.
	URL param.Opt[string] `json:"url,omitzero"`
	paramObj
}

func (r ImportFromURLInputParam) MarshalJSON() (data []byte, err error) {
	type shadow ImportFromURLInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ImportFromURLInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PUBLIC_INDEXABLE: File is publicly accessible by anyone who has the URL. Search
// engines can index the file. PUBLIC_NOT_INDEXABLE: File is publicly accessible by
// anyone who has the URL. Search engines _can't_ index the file. PRIVATE: File is
// NOT publicly accessible. Requires a signed URL to see content. Search engines
// _can't_ index the file.
type ImportFromURLInputAccess string

const (
	ImportFromURLInputAccessHiddenIndexable    ImportFromURLInputAccess = "HIDDEN_INDEXABLE"
	ImportFromURLInputAccessHiddenNotIndexable ImportFromURLInputAccess = "HIDDEN_NOT_INDEXABLE"
	ImportFromURLInputAccessHiddenPrivate      ImportFromURLInputAccess = "HIDDEN_PRIVATE"
	ImportFromURLInputAccessHiddenSensitive    ImportFromURLInputAccess = "HIDDEN_SENSITIVE"
	ImportFromURLInputAccessPrivate            ImportFromURLInputAccess = "PRIVATE"
	ImportFromURLInputAccessPublicIndexable    ImportFromURLInputAccess = "PUBLIC_INDEXABLE"
	ImportFromURLInputAccessPublicNotIndexable ImportFromURLInputAccess = "PUBLIC_NOT_INDEXABLE"
	ImportFromURLInputAccessSensitive          ImportFromURLInputAccess = "SENSITIVE"
)

// ENTIRE_PORTAL: Look for a duplicate file in the entire account. EXACT_FOLDER:
// Look for a duplicate file in the provided folder.
type ImportFromURLInputDuplicateValidationScope string

const (
	ImportFromURLInputDuplicateValidationScopeEntirePortal ImportFromURLInputDuplicateValidationScope = "ENTIRE_PORTAL"
	ImportFromURLInputDuplicateValidationScopeExactFolder  ImportFromURLInputDuplicateValidationScope = "EXACT_FOLDER"
)

// NONE: Do not run any duplicate validation. REJECT: Reject the upload if a
// duplicate is found. RETURN_EXISTING: If a duplicate file is found, do not upload
// a new file and return the found duplicate instead.
type ImportFromURLInputDuplicateValidationStrategy string

const (
	ImportFromURLInputDuplicateValidationStrategyNone           ImportFromURLInputDuplicateValidationStrategy = "NONE"
	ImportFromURLInputDuplicateValidationStrategyReject         ImportFromURLInputDuplicateValidationStrategy = "REJECT"
	ImportFromURLInputDuplicateValidationStrategyReturnExisting ImportFromURLInputDuplicateValidationStrategy = "RETURN_EXISTING"
)

type ImportFromURLTaskLocator struct {
	// ID of the task
	ID string `json:"id" api:"required"`
	// Links for where to check information related to the task. The `status` link
	// gives the URL for where to check the status of the task.
	Links map[string]string `json:"links"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Links       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImportFromURLTaskLocator) RawJSON() string { return r.JSON.raw }
func (r *ImportFromURLTaskLocator) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SignedURL struct {
	// Timestamp of when the URL will no longer grant access to the file.
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	// Signed URL with access to the specified file. Anyone with this URL will be able
	// to access the file until it expires.
	URL string `json:"url" api:"required"`
	// Extension of the requested file.
	Extension string `json:"extension"`
	// For image and video files. The height of the file.
	Height int64 `json:"height"`
	// Name of the requested file.
	Name string `json:"name"`
	// Size in bytes of the requested file.
	Size int64 `json:"size"`
	// Type of the file. Can be IMG, DOCUMENT, AUDIO, MOVIE, or OTHER.
	Type string `json:"type"`
	// For image and video files. The width of the file.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		URL         respjson.Field
		Extension   respjson.Field
		Height      respjson.Field
		Name        respjson.Field
		Size        respjson.Field
		Type        respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SignedURL) RawJSON() string { return r.JSON.raw }
func (r *SignedURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
