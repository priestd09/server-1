package fileactions

import (
	"github.com/fragmenta/router"

	"github.com/send-to/server/src/files"
	"github.com/send-to/server/src/lib/authorise"
)

// HandleDestroy handles a DESTROY request for files
func HandleDestroy(context router.Context) error {

	// Find the file
	file, err := files.Find(context.ParamInt("id"))
	if err != nil {
		return router.NotFoundError(err)
	}

	// Authorise destroy file
	err = authorise.ResourceAndAuthenticity(context, file)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Destroy the file
	file.Destroy()

	// Redirect to files root
	return router.Redirect(context, "/files")
}
