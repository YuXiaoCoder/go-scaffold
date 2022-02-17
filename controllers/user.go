package controllers

import (
	"errors"
	"go-scaffold/dao/rds"
	"go-scaffold/logic"
	"go-scaffold/models"
	"go-scaffold/pkg/configs"

	"gopkg.in/guregu/null.v4"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignInHandler 登录
// @Summary 登录
// @Description 登录
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param object body models.BodySignIn true "请求体参数"
// @Success 200 {object} models.ResponseSignIn
// @Router /sign_in [post]
func SignInHandler(ctx *gin.Context) {
	// 获取参数并进行检验
	params := new(models.BodySignIn)
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("failed to controllers.SignInHandler", zap.Error(err))
		// 判断错误的类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		} else {
			ResponseErrorWithMessage(ctx, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
			return
		}
	}

	// 业务逻辑处理
	user, token, err := logic.SignIn(params)
	if err != nil {
		zap.L().Error("failed to logic.SignIn", zap.Error(err))
		ResponseError(ctx, CodeInvalidUser)
		return
	}

	// 返回响应
	ResponseSuccess(ctx, models.DataSignIn{
		ID:       user.ID,
		Email:    user.Email,
		Nickname: user.Nickname,
		Gender:   user.Gender,
		Token:    token,
	})
	return
}

// SignUPHandler 注册
// @Summary 注册
// @Description 注册
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param object body models.BodySignUP true "请求体参数"
// @Success 200 {object} models.ResponseCommon
// @Router /sign_up [post]
func SignUPHandler(ctx *gin.Context) {
	// 获取参数并进行检验
	params := new(models.BodySignUP)
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("failed to controllers.SignUPHandler", zap.Error(err))
		// 判断错误的类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		} else {
			ResponseErrorWithMessage(ctx, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
			return
		}
	}

	// 业务逻辑处理
	if err := logic.SignUP(params); err != nil {
		zap.L().Error("failed to logic.SignUP", zap.Error(err))
		// 用户已存在
		if errors.Is(err, rds.ErrorUserExist) {
			ResponseError(ctx, CodeUserExist)
			return
		} else {
			ResponseError(ctx, CodeServerBusy)
			return
		}
	}

	// 返回响应
	ResponseSuccess(ctx, nil)
	return
}

// ListUserHandler 用户列表
// @Summary 用户列表
// @Description 用户列表
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query models.QueryUserInfoList true "查询参数"
// @Success 200 {object} models.ResponseUserInfoList
// @Router /users [get]
func ListUserHandler(ctx *gin.Context) {
	// 获取参数并进行检验
	params := &models.QueryUserInfoListNull{
		Limit:  null.NewInt(configs.AllConfig.Pager.Limit, true),
		Offset: null.NewInt(configs.AllConfig.Pager.Offset, true),
	}
	if err := ctx.ShouldBindQuery(params); err != nil {
		zap.L().Error("failed to controllers.ListUserHandler", zap.Error(err))
		// 判断错误的类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		} else {
			ResponseErrorWithMessage(ctx, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
			return
		}
	}

	// 业务逻辑处理
	data, err := logic.ListUser(params)
	if err != nil {
		zap.L().Error("failed to logic.ListUser", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}

	// 返回响应
	ResponseSuccess(ctx, data)
	return
}

// GetUserByIDHandler 获取指定用户
// @Summary 获取指定用户
// @Description 获取指定用户
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path int true "User ID"
// @Success 200 {object} models.ResponseUserInfo
// @Router /users/{id} [get]
func GetUserByIDHandler(ctx *gin.Context) {
	// 获取参数并进行检验
	id, err := getInt64Param(ctx, "id")
	if err != nil {
		ResponseError(ctx, CodeInvalidParam)
		return
	}

	// 业务逻辑处理
	data, err := logic.GetUserByID(id)
	if err != nil {
		zap.L().Error("failed to logic.GetUserByID", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}

	// 返回响应
	ResponseSuccess(ctx, data)
	return
}

// UpdateUserByIDHandler 更新指定用户
// @Summary 更新指定用户
// @Description 更新指定用户
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path int true "User ID"
// @Param object body models.BodyUpdateUser false "请求体参数"
// @Success 200 {object} models.ResponseUserInfo
// @Router /users/{id} [post]
func UpdateUserByIDHandler(ctx *gin.Context) {
	// 获取参数并进行检验
	id, err := getInt64Param(ctx, "id")
	if err != nil {
		ResponseError(ctx, CodeInvalidParam)
		return
	}

	params := new(models.BodyUpdateUserNull)
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("failed to controllers.UpdateUserByIDHandler", zap.Error(err))
		// 判断错误的类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		} else {
			ResponseErrorWithMessage(ctx, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
			return
		}
	}

	// 业务逻辑处理
	data, err := logic.UpdateUserByID(id, params)
	if err != nil {
		zap.L().Error("failed to logic.UpdateUserByID", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}

	// 返回响应
	ResponseSuccess(ctx, data)
	return
}

// DeleteUserByIDHandler 删除指定用户
// @Summary 删除指定用户
// @Description 删除指定用户
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path int true "User ID"
// @Success 200 {object} models.ResponseCommon
// @Router /users/{id} [delete]
func DeleteUserByIDHandler(ctx *gin.Context) {
	// 获取参数并进行检验
	id, err := getInt64Param(ctx, "id")
	if err != nil {
		ResponseError(ctx, CodeInvalidParam)
		return
	}

	// 业务逻辑处理
	if err = logic.DeleteUserByID(id); err != nil {
		zap.L().Error("failed to logic.DeleteUserByID", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}

	// 返回响应
	ResponseSuccess(ctx, nil)
	return
}
