package openShamrockApi

import (
	"encoding/json"
	adapter "github.com/ycvk/FreeClover/driver/openshamrock"
	entity "github.com/ycvk/FreeClover/entity/openshamrock"
)

// File 文件相关
type File struct {
	transceiver adapter.Transceiver
}

// SendPrivateFile 该接口用于上传私聊文件。
func (o File) SendPrivateFile(userId int64, fileUri string, name string) entity.SendPrivateFile {
	endpoint := "upload_private_file"
	values, err := json.Marshal(map[string]interface{}{
		"user_id": userId,
		"file":    fileUri,
		"name":    name,
	})
	if err != nil {
		return entity.SendPrivateFile{}
	}
	return *processJson[entity.SendPrivateFile](endpoint, values, o.transceiver)
}

// SendGroupFile 该接口用于上传群文件。
func (o File) SendGroupFile(groupId int64, fileUri string, name string) entity.SendGroupFile {
	endpoint := "upload_group_file"
	values, err := json.Marshal(map[string]interface{}{
		"group_id": groupId,
		"file":     fileUri,
		"name":     name,
	})
	if err != nil {
		return entity.SendGroupFile{}
	}
	return *processJson[entity.SendGroupFile](endpoint, values, o.transceiver)
}

// DeleteGroupFile 该接口用于删除群文件。
func (o File) DeleteGroupFile(groupId int64, fileId string, busid int32) entity.Common {
	endpoint := "delete_group_file"
	values, err := json.Marshal(map[string]interface{}{
		"group_id": groupId,
		"file_id":  fileId,
		"busid":    busid,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}

// CreateGroupFileFolder 该接口用于创建群文件文件夹。
func (o File) CreateGroupFileFolder(msgId int32) entity.Common {
	endpoint := "create_group_file_folder"
	values, err := json.Marshal(map[string]interface{}{
		"msg_id": msgId,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}

// DeleteGroupFolder 该接口用于删除群文件文件夹。
func (o File) DeleteGroupFolder(groupId int64, folderId string) entity.Common {
	endpoint := "delete_group_folder"
	values, err := json.Marshal(map[string]interface{}{
		"group_id":  groupId,
		"folder_id": folderId,
	})
	if err != nil {
		return entity.Common{}
	}
	return *processJson[entity.Common](endpoint, values, o.transceiver)
}

// GetGroupFileSystemInfo 该接口用于获取群文件系统信息。
func (o File) GetGroupFileSystemInfo(groupId int64) entity.GroupFileSystemInfo {
	endpoint := "get_group_file_system_info"
	values, err := json.Marshal(map[string]interface{}{
		"group_id": groupId,
	})
	if err != nil {
		return entity.GroupFileSystemInfo{}
	}
	return *processJson[entity.GroupFileSystemInfo](endpoint, values, o.transceiver)
}

// GetGroupRootFiles 该接口用于获取群根目录文件列表。
func (o File) GetGroupRootFiles(groupId int64) entity.GroupRootFileList {
	endpoint := "get_group_root_files"
	values, err := json.Marshal(map[string]interface{}{
		"group_id": groupId,
	})
	if err != nil {
		return entity.GroupRootFileList{}
	}
	return *processJson[entity.GroupRootFileList](endpoint, values, o.transceiver)
}

// GetGroupFilesFolder 该接口用于获取群子目录文件列表。
func (o File) GetGroupFilesFolder(groupId int64, folderId string) entity.GroupFiles {
	endpoint := "get_group_files_by_folder"
	values, err := json.Marshal(map[string]interface{}{
		"group_id":  groupId,
		"folder_id": folderId,
	})
	if err != nil {
		return entity.GroupFiles{}
	}
	return *processJson[entity.GroupFiles](endpoint, values, o.transceiver)
}

// GetGroupFileUrl 该接口用于获取群文件资源链接。
func (o File) GetGroupFileUrl(groupId int64, fileId string, busid int32) entity.GroupFileUrl {
	endpoint := "get_group_file_url"
	values, err := json.Marshal(map[string]interface{}{
		"group_id": groupId,
		"file_id":  fileId,
		"busid":    busid,
	})
	if err != nil {
		return entity.GroupFileUrl{}
	}
	return *processJson[entity.GroupFileUrl](endpoint, values, o.transceiver)
}
