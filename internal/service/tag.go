package service

import (
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/WeCanRun/gin-blog/internal/dto"
	"github.com/WeCanRun/gin-blog/internal/model"
	"github.com/WeCanRun/gin-blog/internal/server"
	"github.com/WeCanRun/gin-blog/internal/service/cache_service"
	"github.com/WeCanRun/gin-blog/pkg/export"
	"github.com/WeCanRun/gin-blog/pkg/file"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/setting"
	"github.com/WeCanRun/gin-blog/pkg/util"
	"github.com/jinzhu/gorm"
	"github.com/tealeg/xlsx"
	"io"
	"strconv"
	"time"
)

func GetTags(c *server.Context, req *dto.GetTagsRequest) (resp *dto.GetTagsResponse, err error) {
	pageNum := util.GetPage(c)
	pageSize := setting.APP.PageSize

	tags, err := model.GetTags(pageNum, pageSize)
	if err != nil {
		logging.Error("services#GetTags fail %v", err)
		return
	}
	var ids []uint
	var names []string
	for _, tag := range tags {
		ids = append(ids, tag.ID)
		names = append(names, tag.Name)
	}
	resp = &dto.GetTagsResponse{
		Pager: dto.Pager{
			PageNum:   pageNum,
			PageSize:  pageSize,
			TotalRows: uint(len(tags)),
		},
		IDs:   ids,
		Names: names,
	}
	return
}

func GetTag(ctx *server.Context, id uint) (resp dto.GetTagResponse, err error) {
	tag, err := cache_service.GetTagCacheById(id)
	if tag.ID > 0 {
		resp.ID = id
		resp.Name = tag.Name
		logging.Info("GetTag | get tag from cache, id:%d", id)
		return
	}

	tag, err = model.GetTagById(id)
	if err != nil {
		logging.Error("models#GetTagById fail, %v", err)
		return
	}

	if err := cache_service.SetTagCacheById(tag); err != nil {
		logging.Error("GetTag | cache_service.SetTagCacheById fail, err:%v", err)
	}

	resp.ID = id
	resp.Name = tag.Name
	return
}

func AddTag(ctx *server.Context, req *dto.AddTagRequest) (err error) {
	if req.State != 0 && req.State != 1 {
		err = errors.New("services#AddTag state 只能为 0 或 1")
		logging.Error("err: %v", err)
		return
	}
	// todo 其他敏感词校验
	err = model.AddTag(model.Tag{
		Name:      req.Name,
		CreatedBy: req.CreatedBy,
		State:     req.State,
	})
	return
}

func DeleteTag(ctx *server.Context, id uint) (err error) {
	// todo 各种校验
	err = model.DeleteTag(id)
	if err != nil {
		logging.Error("DeleteTag | model.DeleteTag fail, err:%v", err)
		return
	}

	err = cache_service.DeleteTagCacheById(id)
	if err != nil {
		logging.Error("DeleteTag | cache_service.DeleteTagCacheById fail, err:%v", err)

	}

	return nil
}

func EditTag(ctx *server.Context, request *dto.EditRequest) (err error) {
	// todo 各种校验
	err = model.EditTag(model.Tag{
		Model:     gorm.Model{ID: request.ID},
		Name:      request.Name,
		UpdatedBy: request.UpdatedBy,
		State:     request.State,
	})
	if err != nil {
		logging.Error("EditTag | model.EditTag fail, err:%v", err)
		return
	}

	tag, err := model.GetTagById(request.ID)
	if err != nil {
		logging.Error("EditTag | model.GetTagById fail, err:%v", err)
	}

	err = cache_service.SetTagCacheById(tag)
	if err != nil {
		logging.Error("EditTag | cache_service.SetTagCacheById fail, err%v", err)
	}

	return nil
}

// 根据名字导出标签
func ExportTags(name string, state int) (dto.ExportTagsResponse, error) {
	tags, err := model.GetTagsByName(name)
	if err != nil {
		logging.Error("ExportTags |  model.GetTagsByName fail, err:%v", err)
		return dto.ExportTagsResponse{}, err
	}

	excel := xlsx.NewFile()
	sheet, err := excel.AddSheet("标签信息")
	if err != nil {
		logging.Error("ExportTags | excel.AddSheet fail, err%v", err)
		return dto.ExportTagsResponse{}, err
	}
	// 写入标题行
	titles := []string{"ID", "名称", "创建人", "创建时间", "修改人", "修改时间"}
	row := sheet.AddRow()
	var cell *xlsx.Cell
	for _, title := range titles {
		cell = row.AddCell()
		cell.Value = title
	}

	// 写入内容
	for _, tag := range tags {
		values := []string{
			strconv.Itoa(int(tag.ID)),
			tag.Name,
			tag.CreatedBy,
			strconv.Itoa(int(tag.CreatedAt.Unix())),
			tag.UpdatedBy,
			strconv.Itoa(int(tag.UpdatedAt.Unix())),
		}
		row := sheet.AddRow()
		for _, value := range values {
			row.AddCell().Value = value
		}
	}

	// 创建保存的目录
	excelPath := export.GetExcelRealDir()
	err = file.IsNotExitMKDir(excelPath)
	if err != nil {
		logging.Error("ExportTags | file.IsNotExitMKDir fail, err%v", err)
		return dto.ExportTagsResponse{}, err
	}

	// 构建文件名
	exportName := export.ExportExcelName(name, time.Now())
	fullName := export.GetExcelRealPath(exportName)
	err = excel.Save(fullName)
	if err != nil {
		logging.Error("ExportTags | excel.Save fail, err%v", err)
	}

	response := dto.ExportTagsResponse{
		ExportFullUrl: export.GetExcelFullUrl(exportName),
		ExportSaveUrl: export.GetExcelSaveUrl(exportName),
	}
	return response, err
}

// 导入标签信息
func ImportTags(r io.Reader) error {
	excel, err := excelize.OpenReader(r)
	if err != nil {
		logging.Error("ImportTags | excelize.OpenReader fail, err:%v", err)
		return err
	}
	rows, err := excel.GetRows("标签信息")
	if err != nil {
		logging.Error("ImportTags | GetRows fail, err:%v ", err)
	}

	for i, row := range rows {
		if i > 0 {
			var data []string
			for _, cell := range row {
				data = append(data, cell)
			}
			if len(data) < 2 {
				logging.Error("ImportTags | 数据解析错误, data:%v", data)
				continue
			}
			model.AddTag(model.Tag{
				Name:      data[1],
				CreatedBy: data[2],
				State:     1,
			})
		}
	}
	return nil
}
