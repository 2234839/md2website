package main

import (
	"bytes"
	"html/template"
	"path"

	store "github.com/2234839/md2website/src/store"
	"github.com/2234839/md2website/src/util"
)

// HTMLtemplate 包含一些模板
var HTMLtemplate = template.Must(template.ParseGlob(path.Join(TemplateDir, "./*.html")))

func init() {
	template.Must(HTMLtemplate.ParseGlob(path.Join(TemplateDir, "./*/*.html")))
}
func unescaped(x string) interface{} { return template.HTML(x) }

// ExecTemplate 执行模板
/**
 * 现有模板 menu
 */
func ExecTemplate(t *template.Template, data interface{}) string {
	buf := new(bytes.Buffer)
	t.Execute(buf, data)
	return buf.String()
}

var globalF = template.FuncMap{"unescaped": unescaped}

var articleTemplate = HTMLtemplate.New("article").Funcs(globalF)
var menuTemplate = HTMLtemplate.New("menu").Funcs(globalF)
var embeddedBlockTemplate = HTMLtemplate.New("embeddedBlock").Funcs(globalF)
var blockRefTemplate = HTMLtemplate.New("blockRef").Funcs(globalF)

type sonEntityI struct {
	WebPath string
	IsDir   bool
	Name    string
}

// ArticleInfo 结构
type ArticleInfo struct {
	PageTitle string
	Content   interface{}
	LevelRoot string
}

// MenuInfo 菜单结构
type MenuInfo struct {
	SonEntityList []sonEntityI
	PageTitle     string
	LevelRoot     string
}

// ArticleRender 渲染文章html
func ArticleRender(info ArticleInfo) string {
	return ExecTemplate(articleTemplate, info)
}

// MenuRender 渲染菜单
func MenuRender(info MenuInfo) string {
	return ExecTemplate(menuTemplate, info)
}

// EmbeddedBlockRender 渲染嵌入块
func EmbeddedBlockRender(info store.EmbeddedBlockInfo) string {
	return ExecTemplate(embeddedBlockTemplate, info)
}

// BlockRefRender 渲染块引用
func BlockRefRender(info store.BlockRefInfo) string {
	return ExecTemplate(blockRefTemplate, info)
}

// TemplateRender 将数据通过模板进行渲染 目前支持 EmbeddedBlockInfo BlockRefInfo 的处理
func TemplateRender(info interface{}) string {
	EmbeddedBlock, ok := info.(store.EmbeddedBlockInfo)
	if ok {
		return EmbeddedBlockRender(EmbeddedBlock)
	}

	BlockRef, ok := info.(store.BlockRefInfo)
	if ok {
		return BlockRefRender(BlockRef)
	}
	util.Warn("没有找到对应的 template render", info)
	return "[渲染错误]没有找到对应的 template render"
}
