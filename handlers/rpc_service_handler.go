package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"project/internal/models"
	"project/internal/services"
)

type RpcServiceHandler struct {
	service *services.RpcServiceService
}

func NewRpcServiceController(service *services.RpcServiceService) *RpcServiceController {
	return &RpcServiceController{service: service}
}

func (c *RpcServiceHandler) CreateRpcService(ctx *gin.Context) {
	var rpcService models.RpcService
	if err := ctx.ShouldBindJSON(&rpcService); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateRpcService(&rpcService); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create rpc service"})
		return
	}

	ctx.JSON(http.StatusCreated, rpcService)
}

func (c *RpcServiceHandler) GetRpcService(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	rpcService, err := c.service.GetRpcService(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Rpc service not found"})
		return
	}

	ctx.JSON(http.StatusOK, rpcService)
}

func (c *RpcServiceHandler) UpdateRpcService(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var rpcService models.RpcService
	if err := ctx.ShouldBindJSON(&rpcService); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rpcService.ID = uint(id)

	if err := c.service.UpdateRpcService(uint(id), &rpcService); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update rpc service"})
		return
	}

	ctx.JSON(http.StatusOK, rpcService)
}

func (c *RpcServiceHandler) DeleteRpcService(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.service.DeleteRpcService(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete rpc service"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
