// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package db

import (
	"context"
	"sort"

	"github.com/chai2010/template"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"openpitrix.io/iam/pkg/internal/strutil"
	"openpitrix.io/iam/pkg/pb/im"
	"openpitrix.io/iam/pkg/service/im/db_spec"
	"openpitrix.io/logger"
)

func (p *Database) getAllSubGroupIds(ctx context.Context, req *pbim.GroupIdList) ([]string, error) {
	const sqlTmpl = `
		SELECT * FROM user_group WHERE 1=0
			{{range $i, $v := .GroupId}}
				OR group_path LINK '%{{$v}}%'
				OR group_id='{{$v}}'
			{{end}}
	`
	var query, err = template.Render(sqlTmpl, req)
	if err != nil {
		err := status.Errorf(codes.Internal, "%v", err)
		logger.Errorf(ctx, "%+v", err)
		return nil, err
	}

	query = strutil.SimplifyString(query)
	logger.Infof(ctx, "query: %s", query)

	var rows []db_spec.UserGroup
	p.DB.Raw(query).Find(&rows)
	if err := p.DB.Error; err != nil {
		logger.Warnf(ctx, "%+v", err)
		return nil, err
	}

	var allGroupId []string
	for _, v := range rows {
		allGroupId = append(allGroupId, v.GroupId)
	}

	sort.Strings(allGroupId)
	return allGroupId, nil
}