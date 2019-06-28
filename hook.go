package hookd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/opentracing/opentracing-go"
	manager_v1 "github.com/videocoin/cloud-api/manager/v1"
	workorder_v1 "github.com/videocoin/cloud-api/workorder/v1"

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
	ctx := req.Context()
	err := req.ParseForm()
	if err != nil {
		h.log.Errorf("failed to parse form: %s", err)
		return ErrBadRequest
	}

	h.log.Debugf("handle hook %+v", req.Form)

	call := req.FormValue("call")
	switch call {
	case "publish":
		err = h.handlePublish(ctx, req)
	case "publish_done":
		err = h.handlePublishDone(ctx, req)
	case "record":
		err = h.handleRecord(ctx, req)
	case "record_done":
		err = h.handleRecordDone(ctx, req)
	default:
		return c.NoContent(http.StatusBadRequest)
	}

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Hook) handlePublish(ctx context.Context, r *http.Request) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "handlePublish")
	defer span.Finish()

	h.log.Info("handling publish hook")

	streamInfo, err := ParseStreamName(r.FormValue("name"))
	if err != nil {
		h.log.Warnf("failed to parse stream name [ %+v ]: %s", streamInfo, err)
		return ErrBadRequest
	}

	h.log = h.log.WithFields(logrus.Fields{
		"job_id": streamInfo.JobID,
	})

	span.LogKV(
		"job_id", streamInfo.JobID,
	)

	h.log.Infof("using job id: %s", streamInfo.JobID)

	h.log.Info("getting user profile")

	h.log.Info("marking camera as on air")

	managerResp, err := h.manager.UpdateStatus(context.Background(), &manager_v1.UpdateJobRequest{
		Id:           streamInfo.JobID,
		Status:       workorder_v1.WorkOrderStatusPending,
		IngestStatus: workorder_v1.IngestStatusActive,
	})

	if err != nil {
		h.log.Errorf("failed to update stream status: %s", err.Error())
	}

	h.log.Debugf("manager response: %+v", managerResp)

	return nil
}

func (h *Hook) handlePublishDone(ctx context.Context, r *http.Request) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "handlePublishDone")
	defer span.Finish()

	h.log.Info("handling publish done hook")

	streamInfo, err := ParseStreamName(r.FormValue("name"))
	if err != nil {
		h.log.Warningf("failed to parse stream name: %s", err)
		return ErrBadRequest
	}

	h.log = h.log.WithFields(logrus.Fields{
		"job_id": streamInfo.JobID,
	})

	span.LogKV(
		"job_id", streamInfo.JobID,
	)

	h.log.Info("marking stream as offline")

	managerResp, err := h.manager.Stop(context.Background(), &manager_v1.JobRequest{
		Id: streamInfo.JobID,
	})

	if err != nil {
		h.log.Errorf("failed to stop stream: %s", err.Error())
	}

	h.log.Debugf("manager response: %+v", managerResp)

	return nil

}

func (h *Hook) handleRecord(ctx context.Context, r *http.Request) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "handleRecord")
	defer span.Finish()
	return nil
}

func (h *Hook) handleRecordDone(ctx context.Context, r *http.Request) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "handleRecordDone")
	defer span.Finish()
	return nil
}
