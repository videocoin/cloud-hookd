package hookd

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	jobs_v1 "github.com/videocoin/cloud-api/jobs/v1"
	manager_v1 "github.com/videocoin/cloud-api/manager/v1"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// Common hook errors
var (
	ErrUnknownHook  = fmt.Errorf("unknown hook")
	ErrBadJobStatus = fmt.Errorf("invalid job status")
	ErrBadRequest   = echo.NewHTTPError(http.StatusBadRequest)
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
	default:
		return c.NoContent(http.StatusBadRequest)
	}

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Hook) handlePublish(ctx context.Context, r *http.Request) error {
	span, spanCtx := opentracing.StartSpanFromContext(ctx, "handlePublish")
	defer span.Finish()

	h.log.Info("handling publish hook")

	streamInfo, err := ParseStreamName(spanCtx, r.FormValue("name"))
	if err != nil {
		h.log.Warnf("failed to parse stream name [ %+v ]: %s", streamInfo, err)
		return ErrBadRequest
	}

	h.log = h.log.WithFields(logrus.Fields{
		"job_id": streamInfo.JobID,
	})

	span.SetTag("job_id", streamInfo.JobID)

	h.log.Infof("using job id: %s", streamInfo.JobID)

	if !h.prepared(streamInfo.JobID) {
		h.log.Error("invalid job status")
		return ErrBadJobStatus
	}

	h.log.Info("marking camera as on air")

	managerResp, err := h.manager.UpdateStatus(context.Background(), &manager_v1.UpdateJobRequest{
		Id:          streamInfo.JobID,
		Status:      jobs_v1.JobStatusPending,
		InputStatus: jobs_v1.InputStatusActive,
	})

	if err != nil {
		h.log.Errorf("failed to update stream status: %s", err.Error())
	}

	h.log.Debugf("manager response: %+v", managerResp)

	return nil
}

func (h *Hook) handlePublishDone(ctx context.Context, r *http.Request) error {
	span, spanCtx := opentracing.StartSpanFromContext(ctx, "handlePublishDone")
	defer span.Finish()

	h.log.Info("handling publish done hook")

	streamInfo, err := ParseStreamName(spanCtx, r.FormValue("name"))
	if err != nil {
		h.log.Warningf("failed to parse stream name: %s", err)
		return ErrBadRequest
	}

	h.log = h.log.WithFields(logrus.Fields{
		"job_id": streamInfo.JobID,
	})

	span.SetTag("job_id", streamInfo.JobID)

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

func (h *Hook) prepared(jobID string) bool {
	ticker := time.Tick(5 * time.Second)
	timeout := time.After(2 * time.Minute)
	for {
		select {
		case <-ticker:
			job, err := h.manager.Get(context.Background(), &manager_v1.JobRequest{
				Id: jobID,
			})
			if err != nil {
				h.log.Warnf("failed to get job: %s", err.Error())
			}
			if job.Status == jobs_v1.JobStatusPrepared {
				return true
			}
		case <-timeout:
			return false
		}
	}
}
