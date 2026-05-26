import (
	"context"
	"errors"
	"fmt"
	{{if .time}}"time"{{end}}
	{{if .containsDbSql}}"database/sql"{{end}}
	"github.com/polpo-space/gorm-zero/gormc"
    "github.com/polpo-space/gorm-zero/gormc/batchx"
	"github.com/polpo-space/gorm-zero/gormc/pagex"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"

	{{if .third}}{{.third}}{{end}}
)
