package controllers

type JsonResult struct {
	Success bool        `json:"Result"`
	Content interface{} `json:"Info"`
}
