package handler

import (
	"log"
	"net/http"

	"github.com/brayden-ooi/bookkeeper/internal/database"
	"github.com/brayden-ooi/bookkeeper/internal/service/account_tag"
	"github.com/brayden-ooi/bookkeeper/internal/utils"
	"github.com/brayden-ooi/bookkeeper/internal/view/pages/pages_account_tags"
)

func CreateTag(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	switch r.Method {
	case http.MethodPost:
		// grab Form
		tag_id := r.FormValue(account_tag.Tag_id)
		tag_name := r.FormValue(account_tag.Tag_name)
		tag_description := r.FormValue(account_tag.Tag_description)

		if _, err := account_tag.Init(ctx).Create(tag_id, tag_name, tag_description); err != nil {
			log.Fatal(err)
		}

		http.Redirect(w, r, "/account-tags", http.StatusSeeOther)
	case http.MethodGet:
		pages_account_tags.Create().Render(ctx, w)
	}
}

func ListTags(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tags, err := account_tag.Init(ctx).List()

	if err != nil {
		log.Fatal(err)
	}

	// render the page
	pages_account_tags.Index(utils.Filter(tags, func(acc database.AccountTag) bool {
		return string(acc.ID[0]) != "_"
	})).Render(ctx, w)
}

func GetTag(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	ctx := r.Context()

	if id == "" {
		log.Fatal("invalid read account tag argument")
	}

	tag, err := account_tag.Init(ctx).GetByID(id)

	if err != nil {
		log.Fatal(err)
	}

	pages_account_tags.Index([]database.AccountTag{tag}).Render(ctx, w)
}
