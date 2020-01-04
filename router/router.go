// Copyright 2019 HttpBox Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package router

import (
	"HttpBox/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initRouter(r *gin.Engine) {

	r.GET("/health", func(ctx *gin.Context) {
		_, _ = ctx.Writer.WriteString("Health")
		ctx.Writer.Flush()
	})

	r.GET("/echo", func(ctx *gin.Context) {
		req := ctx.Request
		_, _ = ctx.Writer.WriteString("Host: " + req.Host + "\n")
		_, _ = ctx.Writer.WriteString("Url: " + req.RequestURI + "\n")
		_, _ = ctx.Writer.WriteString("Method: " + req.Method + "\n")
		_, _ = ctx.Writer.WriteString("RequestHeaders: " + "\n")
		for key, value := range req.Header {
			_, _ = ctx.Writer.WriteString("  " + key + ":" + fmt.Sprintf("%v", value) + "\n")
		}
		_, _ = ctx.Writer.WriteString("Body: " + "\n")
		ctx.Writer.Flush()
	})

	r.GET("/json", func(ctx *gin.Context) {
		httpBox := model.HttpBox{
			Name:    "HttpBox",
			Desc:    "A mock server for test",
			Version: "v0.0.1",
			Tags:    []string{"Http", "MockServer"},
		}
		ctx.JSON(http.StatusOK, httpBox)
	})

}
