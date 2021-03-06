// Package versioncommands implements the version ReST end point
package versioncommands

import (
	"net/http"

	"github.com/gluster/glusterd2/pkg/api"
	restutils "github.com/gluster/glusterd2/servers/rest/utils"
	"github.com/gluster/glusterd2/version"
)

func getVersionHandler(w http.ResponseWriter, r *http.Request) {
	resp := api.VersionResp{
		GlusterdVersion: version.GlusterdVersion,
		APIVersion:      version.APIVersion,
	}
	restutils.SendHTTPResponse(w, http.StatusOK, resp)
}
