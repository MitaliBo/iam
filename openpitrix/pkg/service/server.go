// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package service

import (
	"net/http"

	"google.golang.org/grpc"
)

type Server struct {
	webServer  *http.Server
	grpcServer *grpc.Server
}

func OpenServer(dbtype, dbpath string) (*Server, error) {
	panic("TODO")
}

func (p *Server) Close() error {
	panic("TODO")
}
