package group

import (
	"github.com/Fs02/grimoire/query"
)

func By(fields ...string) query.GroupClause {
	return query.Group(fields...)
}

func Fields(fields ...string) query.GroupClause {
	return By(fields...)
}