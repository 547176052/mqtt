package controller

import (
	"html/template"

	"github.com/547176052/mqtt/guard"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/template/types"
)

func (h *Handler) Example(ctx *context.Context) {
	var param = guard.GetExampleParam(ctx)
	h.HTML(ctx, types.Panel{
		Title:       "Example",
		Description: "example",
		Content:     template.HTML(param.Param),
	})
}
