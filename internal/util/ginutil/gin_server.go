// Copyright 2019 The Vearch Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package ginutil

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vearch/vearch/internal/util"
)

type GinServer struct {
	Server *http.Server
}

func NewGinServer(ginRouter *gin.Engine, ip string, port uint16) *GinServer {
	if port == 0 {
		panic(errors.New("can not found module.http-port in config"))
	}

	if strings.Compare(ip, "127.0.0.1") == 0 || strings.Compare(ip, "localhost") == 0 {
		ip = ""
	}

	server := &http.Server{
		Addr:         util.BuildAddr(ip, port),
		Handler:      ginRouter,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	//server.SetKeepAlivesEnabled(false)
	return &GinServer{Server: server}
}

//func (this *GinServer) Export() error {
//	docService, err := document.NewDocService()
//	if err != nil {
//		return err
//	}
//
//	document.StartHttp(this.server.Handler.(*gin.Engine), *docService)
//
//	return nil
//}

func (this *GinServer) Run() error {
	if err := this.Server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (this GinServer) Stop() {
	if this.Server != nil {
		this.Server.Close()
	}
}
