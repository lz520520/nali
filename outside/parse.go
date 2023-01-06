package outside

import (
	"github.com/zu1k/nali/internal/db"
	"github.com/zu1k/nali/pkg/dbif"
)

func DBFind(typ dbif.QueryType, query string) string {
	return db.Find(typ, query)
}
