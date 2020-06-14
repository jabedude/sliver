package rpc

import (
	"context"

	"github.com/bishopfox/sliver/protobuf/sliverpb"
)

/*
	Sliver Implant Framework
	Copyright (C) 2019  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

func (rpc *SliverServer) NamedPipes(ctx context.Context, req *sliverpb.NamedPipesReq) (*sliverpb.NamedPipes, error) {
	resp := &sliverpb.NamedPipes{}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rpc *SliverServer) TCPListener(ctx context.Context, req *sliverpb.TCPPivotReq) (*sliverpb.TCPPivot, error) {
	resp := &sliverpb.TCPPivot{}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
