package fileactions

import (
	"fmt"
	"net/http"

	"github.com/fragmenta/router"

	"github.com/send-to/server/src/files"
	"github.com/send-to/server/src/lib/authorise"
)

// HandleDownload sends a single file
func HandleDownload(context router.Context) error {

	// Find the file
	file, err := files.Find(context.ParamInt("id"))
	if err != nil {
		return router.InternalError(err)
	}

	// Authorise access to this file - only the file owners can access their own files
	err = authorise.Resource(context, file)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// If we are permitted, send the file to the user
	//w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	//http.DetectContentType(data []byte) string
	h := context.Header()
	h.Set("Content-Type", "application/pgp-encrypted")
	h.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", file.Name()))
	h.Set("Content-Transfer-Encoding", "binary")

	http.ServeFile(context, context.Request(), file.Path)
	return nil
}
