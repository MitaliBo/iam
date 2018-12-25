// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package service

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"openpitrix.io/iam/openpitrix/pkg/pb"
)

func (p *Database) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	sql, values := pkgBuildSql_InsertInto(
		dbSpec.GroupTableName,
		req.GetValue(),
	)
	if len(values) == 0 {
		err := status.Errorf(codes.InvalidArgument, "empty field")
		return nil, err
	}
	_, err := p.DB.ExecContext(ctx, sql, values...)
	if err != nil {
		return nil, err
	}

	reply := &pb.CreateGroupResponse{
		Head: &pb.ResponseHeader{
			UserId:     req.GetHead().GetUserId(),
			OwnerPath:  "", // TODO
			AccessPath: "", // TODO
		},
		GroupId: req.GetValue().GetGroupId(),
	}

	return reply, nil
}

func (p *Database) DeleteGroups(ctx context.Context, req *pb.DeleteGroupsRequest) (*pb.DeleteGroupsResponse, error) {
	sql := pkgBuildSql_Delete(
		dbSpec.GroupTableName, dbSpec.GroupPrimaryKeyName,
		req.GroupId...,
	)

	_, err := p.DB.ExecContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	reply := &pb.DeleteGroupsResponse{
		Head: &pb.ResponseHeader{
			UserId:     req.GetHead().GetUserId(),
			OwnerPath:  "", // TODO
			AccessPath: "", // TODO
		},
		GroupId: req.GroupId,
	}

	return reply, nil
}
func (p *Database) ModifyGroup(ctx context.Context, req *pb.ModifyGroupRequest) (*pb.ModifyGroupResponse, error) {
	sql, values := pkgBuildSql_Update(
		dbSpec.GroupTableName, req.GetValue(),
		dbSpec.GroupPrimaryKeyName,
	)

	_, err := p.DB.ExecContext(ctx, sql, values...)
	if err != nil {
		return nil, err
	}

	reply := &pb.ModifyGroupResponse{
		Head: &pb.ResponseHeader{
			UserId:     req.GetHead().GetUserId(),
			OwnerPath:  "", // TODO
			AccessPath: "", // TODO
		},
		GroupId: req.GetValue().GetGroupId(),
	}

	return reply, nil
}
func (p *Database) GetGroup(ctx context.Context, req *pb.GetGroupRequest) (*pb.GetGroupResponse, error) {
	var sql = fmt.Sprintf(
		"SELECT * FROM %s WHERE %s=$1",
		dbSpec.GroupTableName,
		dbSpec.GroupPrimaryKeyName,
	)
	var value pb.Group
	err := p.DB.Get(&value, sql, req.GetGroupId())
	if err != nil {
		return nil, err
	}

	reply := &pb.GetGroupResponse{
		Head: &pb.ResponseHeader{
			UserId:     req.GetHead().GetUserId(),
			OwnerPath:  "", // TODO
			AccessPath: "", // TODO
		},
		Value: &value,
	}

	return reply, nil
}
func (p *Database) DescribeGroups(context.Context, *pb.DescribeGroupsRequest) (*pb.DescribeGroupsResponse, error) {
	panic("TODO")
}
