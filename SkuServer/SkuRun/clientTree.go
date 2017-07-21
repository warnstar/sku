package SkuRun

type ClientTreeNode struct {
	Label string `json:"label"`
	Fd int64 `json:"fd"`
	Children [] ClientTreeNode `json:"children"`
}