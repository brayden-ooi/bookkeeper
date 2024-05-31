// tag is an underlying internal structure used to associate accounts to different attributes
// one good use is add account archtypes such as "assets", "liabilities" and associate the relevant accounts to the tags
// on create user, "assets", "liabilities", "equity" will be created
// every account must be created with one archtype account tag
package account_tag

import (
	"context"
	"database/sql"
	"errors"

	"github.com/brayden-ooi/bookkeeper/internal/database"
	"github.com/brayden-ooi/bookkeeper/internal/service"
	"github.com/brayden-ooi/bookkeeper/internal/utils"
)

type acc_tag_service struct {
	user_id int64
	ctx     context.Context
}

func Init(ctx context.Context) *acc_tag_service {
	return &acc_tag_service{
		user_id: utils.MyContext{ctx}.GetUserID(),
		ctx:     ctx,
	}
}

// write operations
type CreateAccountTagFormKey = string

const (
	Tag_id          CreateAccountTagFormKey = "id"
	Tag_name        CreateAccountTagFormKey = "name"
	Tag_description CreateAccountTagFormKey = "description"
)

func (srv *acc_tag_service) Create(id, name, description string) (database.AccountTag, error) {
	if id == "" || name == "" {
		return database.AccountTag{}, errors.New("invalid create account tag arguments")
	}

	return service.DB.CreateAccountTag(srv.ctx, database.CreateAccountTagParams{
		ID:   id,
		Name: name,
		Description: sql.NullString{
			String: description,
			Valid:  len(description) != 0,
		},
		UserID: srv.user_id,
	})
}

// read operations
func (srv *acc_tag_service) GetByID(tag_id string) (database.AccountTag, error) {
	// fetch individual tag data
	tag, err := service.DB.GetAccountTagByUserAndID(srv.ctx, database.GetAccountTagByUserAndIDParams{
		UserID: srv.user_id,
		ID:     tag_id,
	})

	if err != nil {
		return database.AccountTag{}, err
	}

	return tag, nil
}

func (srv *acc_tag_service) List() ([]database.AccountTag, error) {
	return service.DB.ListAccountTagsByUser(srv.ctx, srv.user_id)
}
