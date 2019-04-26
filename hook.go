package hookd

import (
	"context"
	"fmt"
	"net/http"

	manager_v1 "github.com/VideoCoin/cloud-api/manager/v1"
	workorder_v1 "github.com/VideoCoin/cloud-api/workorder/v1"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// Common hook errors
var (
	ErrUnknownHook = fmt.Errorf("unknown hook")
	ErrBadRequest  = echo.NewHTTPError(http.StatusBadRequest)
)

// Hook struct used for managing hooks
type Hook struct {
	e       *echo.Echo
	log     *logrus.Entry
	manager manager_v1.ManagerServiceClient
}

// NewHook returns new hook reference
func NewHook(
	e *echo.Echo,
	prefix string,
	manager manager_v1.ManagerServiceClient,
	log *logrus.Entry,
) (*Hook, error) {
	hook := &Hook{
		e:       e,
		manager: manager,
		log:     log,
	}
	hook.e.Any(prefix, hook.handleHook)
	return hook, nil
}

func (h *Hook) handleHook(c echo.Context) error {
	req := c.Request()

	err := req.ParseForm()
	if err != nil {
		h.log.Errorf("failed to parse form: %s", err)
		return ErrBadRequest
	}

	h.log.Debugf("handle hook %+v", req.Form)

	call := req.FormValue("call")
	switch call {
	case "publish":
		err = h.handlePublish(req)
	case "publish_done":
		err = h.handlePublishDone(req)
	case "record":
		err = h.handleRecord(req)
	case "record_done":
		err = h.handleRecordDone(req)
	default:
		return c.NoContent(http.StatusBadRequest)
	}

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Hook) handlePublish(r *http.Request) error {
	h.log.Info("handling publish hook")

	streamInfo, err := ParseStreamName(r.FormValue("name"))
	if err != nil {
		h.log.Warnf("failed to parse stream name: %s", err)
		return ErrBadRequest
	}

	h.log = h.log.WithFields(logrus.Fields{
		"wallet_address": streamInfo.WalletAddress,
		"stream_id":      streamInfo.StreamID,
	})

	h.log.Info("getting user profile")

	ctx := context.Background()

	h.log.Info("marking camera as on air")

	managerResp, err := h.manager.UpdateStreamStatus(ctx, &manager_v1.StreamStatusRequest{
		StreamId:     streamInfo.StreamID,
		Status:       workorder_v1.WorkOrderStatusPending.String(),
		IngestStatus: workorder_v1.IngestStatusActive,
	})

	if err != nil {
		h.log.Errorf("failed to update stream status: %s", err.Error())
	}

	h.log.Debugf("manager response: %+v", managerResp)

	return nil
}

func (h *Hook) handlePublishDone(r *http.Request) error {

	h.log.Info("handling publish done hook")

	streamInfo, err := ParseStreamName(r.FormValue("name"))
	if err != nil {
		h.log.Warningf("failed to parse stream name: %s", err)
		return ErrBadRequest
	}

	h.log = h.log.WithFields(logrus.Fields{
		"wallet_address": streamInfo.WalletAddress,
		"stream_id":      streamInfo.StreamID,
	})

	ctx := context.Background()

	h.log.Info("marking stream as offline")

	managerResp, err := h.manager.StopStream(ctx, &manager_v1.StreamRequest{
		StreamId: streamInfo.StreamID,
	})

	if err != nil {
		h.log.Errorf("failed to stop stream: %s", err.Error())
	}

	h.log.Debugf("manager response: %+v", managerResp)

	return nil

}

func (h *Hook) handleRecord(r *http.Request) error {
	return nil
}

func (h *Hook) handleRecordDone(r *http.Request) error {
	return nil
}
