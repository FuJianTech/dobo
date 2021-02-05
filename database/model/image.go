package model

import (
	"bytes"
	"dobo/utils"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"io"
	"os"
)

func GetImageList() ([]types.ImageSummary, error) {
	images, err := utils.Cli.ImageList(utils.Ctx, types.ImageListOptions{})
	if err != nil {
		return nil, err
	}
	return images, nil
}

func PullImage(refStr string) (io.ReadCloser, error) {
	var options types.ImagePullOptions
	return utils.Cli.ImagePull(utils.Ctx, refStr, options)
}

func PushImage(refStr string) (io.ReadCloser, error) {
	var options types.ImagePushOptions
	return utils.Cli.ImagePush(utils.Ctx, refStr, options)
}

// 获取镜像信息
func GetImageInfo(imageId string) (types.ImageInspect, error) {
	raw, _, err := utils.Cli.ImageInspectWithRaw(utils.Ctx, imageId)
	if err != nil {
		return types.ImageInspect{}, err
	}
	return raw, nil
}

func TagImage(source string, target string) error {
	return utils.Cli.ImageTag(utils.Ctx, source, target)
}

func DeleteImage(imageId string, forge bool) error {
	removeOption := types.ImageRemoveOptions{Force: forge, PruneChildren: true}
	_, err := utils.Cli.ImageRemove(utils.Ctx, imageId, removeOption)
	if err != nil {
		return err
	}
	return nil
}

/** 导出镜像 */
func SaveImage(imageTag string) (io.ReadCloser, error) {
	return utils.Cli.ImageSave(utils.Ctx, []string{imageTag})
}

func PruneImage() (types.ImagesPruneReport, error) {
	var filter filters.Args
	return utils.Cli.ImagesPrune(utils.Ctx, filter)
}

/** 导入镜像 */
func ImportImage(filePath string) (string, error) {
	open, err := os.OpenFile(filePath, os.O_RDWR, 666)
	if err != nil {
		return "", errors.New("读取文件失败")
	}

	load, err := utils.Cli.ImageLoad(utils.Ctx, open, false)

	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(load.Body)

	return buf.String(), err
}

