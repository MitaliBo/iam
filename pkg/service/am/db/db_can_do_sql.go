// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package db

const sqlCanDo = `
select distinct
	t3.url,
	t3.url_method
from
	role t1,
	role_module_binding t2,
	action2 t3,
	enable_action t4
where
	t1.role_id=t2.role_id and
	t2.module_id=t3.module_id and
	t1.role_id in (
		select
			t2.role_id
		from
			user_role_binding t1,
			role t2
		where
			t1.role_id=t2.role_id and
			t1.user_id=?
	) and
	t4.action_id=t3.action_id and
	t3.url=? and
	t3.url_method=?
`