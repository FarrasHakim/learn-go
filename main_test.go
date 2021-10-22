package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func Test_handleRequests(t *testing.T) {
	tests := []struct {
		name string
	}{
		// Todo
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleRequests()
		})
	}
}

func Test_getArticles(t *testing.T) {
	r, _ := http.NewRequest("GET", "api/articles", nil)
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test 1",
			args: args{
				w: httptest.NewRecorder(),
				r: r,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getArticles(tt.args.w, tt.args.r)
		})
	}
}

func Test_createArticle(t *testing.T) {
	param := url.Values{}
	param.Set("title", "Test Article")
	param.Set("description", "Test Desc")
	param.Set("author", "Test Author")
	payload := bytes.NewBufferString(param.Encode())

	r, _ := http.NewRequest("POST", "api/articles", payload)
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test 1",
			args: args{
				w: httptest.NewRecorder(),
				r: r,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createArticle(tt.args.w, tt.args.r)
		})
	}
}

func Test_getArticle(t *testing.T) {
	r, _ := http.NewRequest("GET", "api/article/1", nil)
	r.Body = nil
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		mock func()
	}{
		{
			name: "test 1",
			args: args{
				w: httptest.NewRecorder(),
				r: r,
			},
			mock: func() {
				articles = append(articles, Article{ID: "1", Title: "Article 1", Description: "Article 1 Description", Author: "Bobby", CreateDate: time.Date(2021, time.October, 10, 0, 0, 0, 0, time.UTC), LastUpdated: time.Date(2021, time.October, 10, 0, 0, 0, 0, time.UTC)})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			getArticle(tt.args.w, tt.args.r)
		})
	}
}

func Test_updateArticle(t *testing.T) {
	param := url.Values{}
	param.Set("title", "Test Update Article")
	param.Set("description", "Test Update Desc")
	param.Set("author", "Test Update Author")
	payload := bytes.NewBufferString(param.Encode())

	r, _ := http.NewRequest("PUT", "api/article/1", payload)
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test 1",
			args: args{
				w: httptest.NewRecorder(),
				r: r,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateArticle(tt.args.w, tt.args.r)
		})
	}
}

func Test_deleteArticle(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "api/article/1", nil)
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test 1",
			args: args{
				w: httptest.NewRecorder(),
				r: r,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteArticle(tt.args.w, tt.args.r)
		})
	}
}
