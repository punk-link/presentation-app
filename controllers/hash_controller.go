package controllers

import (
	"main/services/common"
	"strconv"

	"github.com/gin-gonic/gin"
	templates "github.com/punk-link/gin-generic-http-templates"
	"github.com/samber/do"
)

type HashController struct {
	hashCoder *common.HashCoder
}

func NewHashController(injector *do.Injector) (*HashController, error) {
	hashCoder := do.MustInvoke[*common.HashCoder](injector)

	return &HashController{
		hashCoder: hashCoder,
	}, nil
}

func (t *HashController) Decode(ctx *gin.Context) {
	hash := ctx.Param("target")
	templates.Ok(ctx, t.hashCoder.Decode(hash))
}

func (t *HashController) Encode(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("target"))
	if err != nil {
		templates.BadRequest(ctx, err.Error())
		return
	}

	templates.Ok(ctx, t.hashCoder.Encode(id))
}
