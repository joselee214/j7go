package image

import (
	"go.7yes.com/j7f/components/errors"
	"j7go/proto/images"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"j7go/services/images"
	"j7go/utils"
)

func Init(g *grpc.Server) {
	s := &imageService{}
	images.RegisterImageServerServer(g, s)

}

type imageService struct{}

//获取单相册图片
func (s *imageService) GetAlbumImages(server images.ImageServer_GetAlbumImagesServer) error {
	for {
		request, err := server.Recv()
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("recv", zap.String("grpc", err.Error()))
			return err
		}

		imageInfos, err := imagesService.GetImages(server.Context(), uint(request.AlbumId), int8(request.CoverType))
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("image", zap.String("update_album_images", err.Error()))
		}

		imagesResponse := &images.GetImagesResponse{}
		for _, image := range imageInfos {
			newImageIterm := images.Image{}
			newImageIterm.ImageUrl = image.ImageURL
			newImageIterm.ImageId = uint32(image.ID)
			newImageIterm.CoverType = uint32(image.CoverType)
			imagesResponse.Images = append(imagesResponse.Images, &newImageIterm)
		}

		imagesResponse.Status = errors.GetResHeader(err)

		err = server.Send(imagesResponse)
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("send", zap.String("grpc", err.Error()))
			return err
		}
	}

}

//更新单个相册下的图片
func (s *imageService) UpdateAlbumImages(server images.ImageServer_UpdateAlbumImagesServer) error {
	for {
		request, err := server.Recv()
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("recv", zap.String("grpc", err.Error()))
			return err
		}

		imageInfos := make([]*imagesService.ImageInfo, 0)
		for _, image := range request.Images {
			imageInfo := &imagesService.ImageInfo{}
			imageInfo.CoverType = int8(image.CoverType)
			imageInfo.ImageUrl = image.ImageUrl
			imageInfo.ImageId = image.ImageId
			imageInfos = append(imageInfos, imageInfo)
		}
		//更新相册图片
		err = imagesService.UpdateAlbumImages(server.Context(), uint(request.AlbumId), imageInfos)
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("image", zap.String("update_album_images", err.Error()))
		}

		updateResponse := images.ImagesResponse{
			Status: errors.GetResHeader(err),
		}
		err = server.Send(&updateResponse)
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("send", zap.String("grpc", err.Error()))
			return err
		}
	}

}

//新增图片
func (s *imageService) AddImages(server images.ImageServer_AddImagesServer) error {
	for {
		request, err := server.Recv()
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("recv", zap.String("grpc", err.Error()))
			return err
		}
		utils.GetTraceLog(server.Context()).Debug("addImages", zap.Any("request", request))

		var imageInfos []*imagesService.ImageInfo
		for _, singleImage := range request.Images {
			imageInfo := &imagesService.ImageInfo{}
			imageInfo.CoverType = int8(singleImage.CoverType)
			imageInfo.ImageUrl = singleImage.ImageUrl
			imageInfos = append(imageInfos, imageInfo)
		}
		albumId, err := imagesService.InsertImages(server.Context(), request.AlbumId, imageInfos)
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("image", zap.String("add_images", err.Error()))
		}

		insertResponse := images.ImagesResponse{
			Status: errors.GetResHeader(err),
		}
		utils.GetTraceLog(server.Context()).Debug("albumid", zap.Any("albumid", albumId))
		insertResponse.AlbumId = albumId
		err = server.Send(&insertResponse)
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("send", zap.String("grpc", err.Error()))
			return err
		}
	}
}

//批量获取特定coverType的图片
func (s *imageService) GetAlbumImagesByCoverType(server images.ImageServer_GetAlbumImagesByCoverTypeServer) error {
	for {
		ImagesRequestInfo, err := server.Recv()
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("recv", zap.String("grpc", err.Error()))
			return err
		}
		formattedImages, err := imagesService.GetAlbumImagesByCoverType(server.Context(), ImagesRequestInfo.CoverType, ImagesRequestInfo.AlbumIds)

		var singleAlbumImages []*images.SingleAlbumResponse

		for _, albumInfo := range formattedImages {
			albumImages := images.SingleAlbumResponse{}
			albumImages.AlbumId = albumInfo.AlbumId
			albumImages.Images = albumInfo.Images
			singleAlbumImages = append(singleAlbumImages, &albumImages)
		}

		albumImagesResponse := images.AlbumsImagesResponse{
			Status:            errors.GetResHeader(err),
			SingleAlbumImages: singleAlbumImages,
		}
		err = server.Send(&albumImagesResponse)
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("send", zap.String("grpc", err.Error()))
			return err
		}
	}
}
