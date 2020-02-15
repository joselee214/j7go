package imagesService

import (
	"context"
	"github.com/joselee214/j7f/components/errors"
	"j7go/errors"
	"j7go/proto/images"
	"go.uber.org/zap"
	"time"
	"j7go/models/images"
	"j7go/utils"
)

//单次批量插入最大值
const BATCH_INSERT_MAX = 100

//图片信息
type ImageInfo struct {
	ImageId   uint32
	ImageUrl  string
	CoverType int8
}

//获取图片
func GetImages(ctx context.Context, albumId uint, coverType int8) ([]*imagesModel.Image, error) {
	var imageInfos []*imagesModel.Image
	var err error

	if coverType == utils.IntZero {
		//获取所有coverType类型的图片
		imageInfos, err = imagesModel.GetAlbumImages(ctx, albumId)
		if err != nil {
			return nil, errors.NewFromCode(business_errors.ImagesError_GET_IMAGES_ERROR)
		}
	} else {
		//根据coverType获取图片
		imageInfos, err = imagesModel.GetAlbumImagesByCoverType(ctx, albumId, coverType)
		if err != nil {
			return nil, errors.NewFromCode(business_errors.ImagesError_GET_IMAGES_ERROR)
		}
	}
	return imageInfos, nil
}

//插入图片，存在albumId则在该album下插入图片，否则新生成album，并插入图片
func InsertImages(ctx context.Context, albumId uint32, imageInfos []*ImageInfo) (uint32, error) {
	if utils.IntZero == albumId {
		//没有相册需要新建相册
		newAlbum := imagesModel.Album{}
		newAlbum.IsDel = int8(utils.NOT_DELETED)
		newAlbum.UpdatedTime = uint(time.Now().Unix())
		newAlbum.CreatedTime = uint(time.Now().Unix())
		err := newAlbum.Insert(ctx)
		if err != nil {
			return utils.IntZero, err
		}
		albumId = uint32(newAlbum.ID)
	}

	//分批次批量插入, 防止单次大数据量插入
	imagesLen := len(imageInfos)
	for i := utils.IntZero; i < imagesLen; i = i + BATCH_INSERT_MAX {
		var addImageInfos []*imagesModel.Image
		var endIndex int
		if i+BATCH_INSERT_MAX > imagesLen {
			endIndex = imagesLen
			addImageInfos = make([]*imagesModel.Image, imagesLen-i)
		} else {
			endIndex = i + BATCH_INSERT_MAX
			addImageInfos = make([]*imagesModel.Image, BATCH_INSERT_MAX)
		}

		for index, imageRequest := range imageInfos[i:endIndex] {
			newImage := &imagesModel.Image{}
			newImage.ImageURL = imageRequest.ImageUrl

			if imageRequest.CoverType == utils.IntZero {
				imageRequest.CoverType = imagesModel.COVER_TYPE_GENERAL
			}
			newImage.CoverType = imageRequest.CoverType
			newImage.AlbumID = uint(albumId)
			newImage.IsDel = int8(utils.NOT_DELETED)
			newImage.UpdatedTime = uint(time.Now().Unix())
			newImage.CreatedTime = uint(time.Now().Unix())
			addImageInfos[index] = newImage
		}

		err := imagesModel.ImagesBatchInsert(ctx, addImageInfos)
		if err != nil {
			return utils.IntZero, err
		}
	}

	return albumId, nil
}

//更新相册图片
func UpdateAlbumImages(ctx context.Context, albumId uint, imageInfos []*ImageInfo) error {

	newImageInfos := make([]*imagesModel.Image, utils.IntZero, len(imageInfos))
	oldImageInfos, err := imagesModel.GetAlbumImagesIndexById(ctx, albumId)
	if err != nil {
		return err
	}

	utils.GetTraceLog(ctx).Debug("imageInfos", zap.Any("imageInfos", imageInfos))
	for _, imageInfo := range imageInfos {
		utils.GetTraceLog(ctx).Debug("imageInfo", zap.Any("imageInfo", imageInfo))
		//imageId为零值则插入
		if imageInfo.ImageId == utils.IntZero {

			newImage := imagesModel.Image{}
			newImage.ImageURL = imageInfo.ImageUrl

			if imageInfo.CoverType == utils.IntZero {
				imageInfo.CoverType = imagesModel.COVER_TYPE_GENERAL
			}
			newImage.CoverType = imageInfo.CoverType
			newImage.AlbumID = albumId
			newImage.IsDel = int8(utils.NOT_DELETED)
			newImage.UpdatedTime = uint(time.Now().Unix())
			newImage.CreatedTime = uint(time.Now().Unix())
			newImageInfos = append(newImageInfos, &newImage)
		} else {
			//imageId不存在，跳出循环
			oldImageInfo, ok := oldImageInfos[imageInfo.ImageId]
			if !ok {
				continue
			}
			//信息完全相同，不处理
			if oldImageInfo.CoverType == imageInfo.CoverType {
				//删掉完全不需要变更的图片信息
				delete(oldImageInfos, imageInfo.ImageId)
				continue
			} else {
				//cover_type变更先删后增处理
				newImage := imagesModel.Image{}
				newImage.ImageURL = imageInfo.ImageUrl
				newImage.CoverType = imageInfo.CoverType
				if imageInfo.CoverType == utils.IntZero {
					imageInfo.CoverType = imagesModel.COVER_TYPE_GENERAL
				}

				newImage.AlbumID = albumId
				newImage.IsDel = int8(utils.NOT_DELETED)
				newImage.UpdatedTime = uint(time.Now().Unix())
				newImage.CreatedTime = uint(time.Now().Unix())
				newImageInfos = append(newImageInfos, &newImage)
			}

		}
	}

	deleteIds := make([]uint32, utils.IntZero, len(oldImageInfos))
	for k := range oldImageInfos {
		deleteIds = append(deleteIds, k)
	}

	//删除图片
	if len(deleteIds) > utils.IntZero{
		err = imagesModel.BatchDeleteImages(ctx, deleteIds)
		if err != nil {
			return err
		}
	}

	//新图片插入, 批量插入，一次最多同时插入BATCH_INSERT_MAX条
	imagesLen := len(newImageInfos)
	if imagesLen > utils.IntZero {
		for i := utils.IntZero; i < imagesLen; i = i + BATCH_INSERT_MAX {
			var endIndex int
			if i+BATCH_INSERT_MAX > imagesLen {
				endIndex = imagesLen
			} else {
				endIndex = i + BATCH_INSERT_MAX
			}

			err = imagesModel.ImagesBatchInsert(ctx, newImageInfos[i:endIndex])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

//获取多相册指定coverType图片列表
func GetAlbumImagesByCoverType(ctx context.Context, coverType uint32, albumIds []uint32) (map[uint]*images.SingleAlbumResponse, error) {
	var err error
	albumImages, err := imagesModel.GetImages(ctx, int8(coverType), albumIds)
	if err != nil {
		return nil, err
	}

	utils.GetTraceLog(ctx).Debug("albumImages", zap.Any("albumImages", albumImages), zap.Error(err))

	imagesFormatted := make(map[uint]*images.SingleAlbumResponse, 0)

	for _, imageInfo := range albumImages {
		var singleImageInfo images.Image
		singleImageInfo.ImageId = uint32(imageInfo.ID)
		singleImageInfo.ImageUrl = imageInfo.ImageURL
		singleImageInfo.CoverType = uint32(imageInfo.CoverType)
		if _, ok := imagesFormatted[imageInfo.AlbumID]; !ok {
			imagesFormatted[imageInfo.AlbumID] = &images.SingleAlbumResponse{
				AlbumId: uint32(imageInfo.AlbumID)}
		}
		imagesFormatted[imageInfo.AlbumID].Images = append(imagesFormatted[imageInfo.AlbumID].Images, &singleImageInfo)

	}

	return imagesFormatted, err
}

func FormatImageInfo (imgs *images.CommonAlbumImageRequest) []*ImageInfo {
	imageInfos := make([]*ImageInfo, len(imgs.Images))
	for _, image := range imgs.Images {
		newImageIterm := ImageInfo{}
		newImageIterm.ImageUrl = image.ImageUrl
		newImageIterm.ImageId = uint32(image.ImageId)
		newImageIterm.CoverType = int8(image.CoverType)
		imageInfos = append(imageInfos, &newImageIterm)
	}

	return imageInfos
}