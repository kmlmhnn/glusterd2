package rebalance

import (
	"fmt"
	"net/http"

	"github.com/gluster/glusterd2/glusterd2/gdctx"
	restutils "github.com/gluster/glusterd2/glusterd2/servers/rest/utils"
	"github.com/gluster/glusterd2/glusterd2/transaction"
	"github.com/gluster/glusterd2/glusterd2/volume"
	"github.com/gluster/glusterd2/pkg/api"
	rebalanceapi "github.com/gluster/glusterd2/plugins/rebalance/api"
	"github.com/gorilla/mux"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func createRebalanceInfo(volname string) *rebalanceapi.RebalanceInfo {
	var commit uint64
	return &rebalanceapi.RebalanceInfo{
		Volname:     volname,
		RebalanceID: uuid.NewRandom(),
		Status:      rebalanceapi.GfDefragStatusStarted,
		Cmd:         rebalanceapi.GfDefragCmdStart,
		CommitHash:  setCommitHash(commit),
	}
}

func rebalanceStart(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	logger := gdctx.GetReqLogger(ctx)

	// collect inputs from url
	volname := mux.Vars(r)["volname"]

	var req rebalanceapi.RebalanceStartReq

	// Unmarshal Request so to handle fix-layout and start force
	if err := restutils.UnmarshalRequest(r, &req); err != nil {
		restutils.SendHTTPError(ctx, w, http.StatusUnprocessableEntity, "Unable to parse the request", api.ErrCodeDefault)
		return
	}

	vol, err := volume.GetVolume(volname)
	if err != nil {
		restutils.SendHTTPError(r.Context(), w, http.StatusNotFound, "Invalid volume", api.ErrCodeDefault)
		return
	}

	if vol.State != volume.VolStarted {
		restutils.SendHTTPError(r.Context(), w, http.StatusBadRequest, "Volume not started", api.ErrCodeDefault)
		return
	}

	if vol.DistCount == 1 {
		restutils.SendHTTPError(r.Context(), w, http.StatusBadRequest, "Volume is not distributed volume or contain only 1 brick",
			api.ErrCodeDefault)
		return
	}

	// Check for remove- brick pending
	//TODO

	//A simple transaction to start rebalance
	txn := transaction.NewTxn(ctx)
	defer txn.Cleanup()
	lock, unlock, err := transaction.CreateLockSteps(volname)
	if err != nil {
		restutils.SendHTTPError(r.Context(), w, http.StatusInternalServerError, err.Error(), api.ErrCodeDefault)
		return
	}

	txn.Nodes = vol.Nodes()
	txn.Steps = []*transaction.Step{
		lock,
		{
			DoFunc: "rebal-start.Commit",
			Nodes:  txn.Nodes,
		},
		{
			DoFunc: "rebal-start.StoreVolume",
			Nodes:  []uuid.UUID{gdctx.MyUUID},
		},
		unlock,
	}

	err = txn.Ctx.Set("volname", volname)
	if err != nil {
		logger.WithError(err).Error("failed to set request in transaction context")
		restutils.SendHTTPError(ctx, w, http.StatusInternalServerError, err.Error(), api.ErrCodeDefault)
		return
	}

	rebal := createRebalanceInfo(volname)
	if rebal == nil {
		logger.WithError(err).Error("failed to create Rebalance info")
		restutils.SendHTTPError(ctx, w, http.StatusInternalServerError, err.Error(), api.ErrCodeDefault)
		return
	}

	err = txn.Ctx.Set("rinfo", rebal)
	if err != nil {
		logger.WithError(err).Error("failed to set rebalance info in transaction context")
		restutils.SendHTTPError(ctx, w, http.StatusInternalServerError, err.Error(), api.ErrCodeDefault)
		return
	}

	err = txn.Do()
	if err != nil {
		logger.WithFields(log.Fields{
			"error":   err.Error(),
			"volname": volname,
		}).Error("failed to start rebalance on volume")
		restutils.SendHTTPError(r.Context(), w, http.StatusInternalServerError, err.Error(), api.ErrCodeDefault)
		return
	}

	// TODO: glusterdVolResetStats resets few variables to zero
	// These variables updates in dht code and then need to be updated in Rebalance info struct
	txn.Ctx.Logger().WithField("volname", rebal.Volname).Info("rebalance started")
	restutils.SendHTTPResponse(ctx, w, http.StatusOK, rebal)
}

func rebalanceStop(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := gdctx.GetReqLogger(ctx)

	// collect inputs from url
	volname := mux.Vars(r)["volname"]

	// Validate rebalance command
	vol, err := volume.GetVolume(volname)
	if err != nil {
		restutils.SendHTTPError(r.Context(), w, http.StatusNotFound, "Invalid volume", api.ErrCodeDefault)
		return
	}

	rebalinfo, err := GetRebalanceInfo(volname)
	if err != nil {
		restutils.SendHTTPError(ctx, w, http.StatusNotFound, err.Error(), api.ErrCodeDefault)
		return
	}

	// Check rebalance process is started or not
	if rebalinfo.Status != rebalanceapi.GfDefragStatusStarted {
		restutils.SendHTTPError(r.Context(), w, http.StatusBadRequest, "Rebalance process is not started",
			api.ErrCodeDefault)
		return
	}

	// Check remove brick operation is running
	//TODO

	//A simple transaction to stop rebalance
	txn := transaction.NewTxn(ctx)
	defer txn.Cleanup()
	lock, unlock, err := transaction.CreateLockSteps(volname)
	if err != nil {
		restutils.SendHTTPError(r.Context(), w, http.StatusInternalServerError, err.Error(), api.ErrCodeDefault)
		return
	}

	txn.Nodes = vol.Nodes()
	txn.Steps = []*transaction.Step{
		lock,
		{
			DoFunc: "rebal-stop.Commit",
			Nodes:  txn.Nodes,
		},
		{
			DoFunc: "rebal-stop.StoreVolume",
			Nodes:  []uuid.UUID{gdctx.MyUUID},
		},
		unlock,
	}

	err = txn.Ctx.Set("volname", volname)
	if err != nil {
		logger.WithError(err).Error("failed to set request in transaction context")
		restutils.SendHTTPError(ctx, w, http.StatusInternalServerError, err.Error(), api.ErrCodeDefault)
		return
	}

	rebalinfo.Volname = volname
	rebalinfo.Status = rebalanceapi.GfDefragStatusStopped
	rebalinfo.Cmd = rebalanceapi.GfDefragCmdStop

	err = txn.Ctx.Set("rinfo", rebalinfo)
	if err != nil {
		logger.WithError(err).Error("failed to set rebalance info in transaction context")
		restutils.SendHTTPError(ctx, w, http.StatusInternalServerError, err.Error(), api.ErrCodeDefault)
		return
	}

	err = txn.Do()
	if err != nil {
		logger.WithFields(log.Fields{
			"error":   err.Error(),
			"volname": volname,
		}).Error("failed to stop rebalance on volume")
		restutils.SendHTTPError(r.Context(), w, http.StatusInternalServerError, err.Error(), api.ErrCodeDefault)
		return
	}

	txn.Ctx.Logger().WithField("volname", rebalinfo.Volname).Info("rebalance stopped")
	restutils.SendHTTPResponse(r.Context(), w, http.StatusOK, "Rebalance Stop")
}
func rebalanceStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// collect inputs from url
	volname := mux.Vars(r)["volname"]

	// Validate rebalance command
	vol, err := volume.GetVolume(volname)
	if err != nil {
		restutils.SendHTTPError(r.Context(), w, http.StatusNotFound, "Invalid volume", api.ErrCodeDefault)
		return
	}

	if vol.DistCount == 1 {
		restutils.SendHTTPError(r.Context(), w, http.StatusBadRequest, "Volume is not distributed volume or contain only 1 brick",
			api.ErrCodeDefault)
		return
	}

	rebal, err := GetRebalanceInfo(volname)
	if err != nil {
		restutils.SendHTTPError(ctx, w, http.StatusNotFound, err.Error(), api.ErrCodeDefault)
		return
	}

	if rebal.Status == rebalanceapi.GfDefragStatusNotStarted {
		restutils.SendHTTPError(r.Context(), w, http.StatusBadRequest, "Rebalance process is not started on particular volume",
			api.ErrCodeDefault)
		return
	}

	// TODO: Need to provide list of nodes where the process is running, currently provides only localhost
	NodeID := gdctx.MyUUID
	rebalStatus := fmt.Sprintf("Rebalance Status:%s-success Running on nodes:%s Rebalance Status:%+v", rebal.Volname, NodeID, rebal)
	restutils.SendHTTPResponse(r.Context(), w, http.StatusOK, rebalStatus)
}
