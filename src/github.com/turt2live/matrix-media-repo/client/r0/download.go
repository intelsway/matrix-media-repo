package r0

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/turt2live/matrix-media-repo/client"
	"github.com/turt2live/matrix-media-repo/config"
	"github.com/turt2live/matrix-media-repo/media_handler"
	"github.com/turt2live/matrix-media-repo/storage"
	"github.com/turt2live/matrix-media-repo/util"
)

type DownloadMediaResponse struct {
	ContentType string
	Filename    string
	SizeBytes   int64
	Location    string
}

func DownloadMedia(w http.ResponseWriter, r *http.Request, db storage.Database, c config.MediaRepoConfig, log *logrus.Entry) interface{} {
	if !ValidateUserCanDownload(r, db, c) {
		return client.AuthFailed()
	}

	params := mux.Vars(r)

	server := params["server"]
	mediaId := params["mediaId"]
	filename := params["filename"]

	log = log.WithFields(logrus.Fields{
		"mediaId":  mediaId,
		"server":   server,
		"filename": filename,
	})

	media, err := media_handler.FindMedia(r.Context(), server, mediaId, c, db, log)
	if err != nil {
		if err == media_handler.ErrMediaNotFound {
			return client.NotFoundError()
		} else if err == media_handler.ErrMediaTooLarge {
			return client.RequestTooLarge()
		}
		log.Error("Unexpected error locating media: " + err.Error())
		return client.InternalServerError("Unexpected Error")
	}

	if filename == "" {
		filename = media.UploadName
	}

	return &DownloadMediaResponse{
		ContentType: media.ContentType,
		Filename:    filename,
		SizeBytes:   media.SizeBytes,
		Location:    media.Location,
	}
}

func ValidateUserCanDownload(r *http.Request, db storage.Database, c config.MediaRepoConfig) (bool) {
	hs := util.GetHomeserverConfig(r.Host, c)
	if !hs.DownloadRequiresAuth {
		return true // no auth required == can access
	}

	accessToken := util.GetAccessTokenFromRequest(r)
	userId, err := util.GetUserIdFromToken(r.Context(), r.Host, accessToken, c)
	return userId != "" && err != nil
}
