package pageactions

import (
	"strings"

	"github.com/fragmenta/router"
	"github.com/fragmenta/view"

	"github.com/send-to/server/src/lib/authorise"
	"github.com/send-to/server/src/pages"
)

// HandleIndex displays a list of pages
func HandleIndex(context router.Context) error {

	// Authorise
	err := authorise.Path(context)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Build a query
	q := pages.Query()

	// Order by required order, or default to id asc
	switch context.Param("order") {

	case "1":
		q.Order("created desc")

	case "2":
		q.Order("updated desc")

	case "3":
		q.Order("name asc")

	default:
		q.Order("id asc")

	}

	// Filter if necessary - this assumes name and summary cols
	filter := context.Param("filter")
	if len(filter) > 0 {
		filter = strings.Replace(filter, "&", "", -1)
		filter = strings.Replace(filter, " ", "", -1)
		filter = strings.Replace(filter, " ", " & ", -1)
		q.Where("( to_tsvector(name) || to_tsvector(summary) @@ to_tsquery(?) )", filter)
	}

	// Fetch the pages
	results, err := pages.FindAll(q)
	if err != nil {
		return router.InternalError(err)
	}

	// Render the template
	view := view.New(context)
	view.AddKey("filter", filter)
	view.AddKey("pages", results)
	return view.Render()

}
