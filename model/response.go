package model

import "time"

// Link 网盘链接
type Link struct {
	Type     string `json:"type" sonic:"type"`
	URL      string `json:"url" sonic:"url"`
	Password string `json:"password" sonic:"password"`
}

// SearchResult 搜索结果
type SearchResult struct {
	MessageID string    `json:"message_id" sonic:"message_id"`
	UniqueID  string    `json:"unique_id" sonic:"unique_id"`     // 全局唯一ID
	Channel   string    `json:"channel" sonic:"channel"`
	Datetime  time.Time `json:"datetime" sonic:"datetime"`
	Title     string    `json:"title" sonic:"title"`
	Content   string    `json:"content" sonic:"content"`
	Links     []Link    `json:"links" sonic:"links"`
	Tags      []string  `json:"tags,omitempty" sonic:"tags,omitempty"`
}

// MergedLink 合并后的网盘链接
type MergedLink struct {
	URL      string    `json:"url" sonic:"url"`
	Password string    `json:"password" sonic:"password"`
	Note     string    `json:"note" sonic:"note"`
	Datetime time.Time `json:"datetime" sonic:"datetime"`
}

// MergedLinks 按网盘类型分组的合并链接
type MergedLinks map[string][]MergedLink

// SearchResponse 搜索响应
type SearchResponse struct {
	Total        int           `json:"total" sonic:"total"`
	Results      []SearchResult `json:"results" sonic:"results"`
	MergedByType MergedLinks   `json:"merged_by_type,omitempty" sonic:"merged_by_type,omitempty"`
}

// Response API通用响应
type Response struct {
	Code    int         `json:"code" sonic:"code"`
	Message string      `json:"message" sonic:"message"`
	Data    interface{} `json:"data,omitempty" sonic:"data,omitempty"`
}

// NewSuccessResponse 创建成功响应
func NewSuccessResponse(data interface{}) Response {
	return Response{
		Code:    0,
		Message: "success",
		Data:    data,
	}
}

// NewErrorResponse 创建错误响应
func NewErrorResponse(code int, message string) Response {
	return Response{
		Code:    code,
		Message: message,
	}
} 