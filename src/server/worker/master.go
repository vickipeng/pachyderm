package worker

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gogo/protobuf/types"

	"github.com/pachyderm/pachyderm/src/client"
	"github.com/pachyderm/pachyderm/src/client/auth"
	"github.com/pachyderm/pachyderm/src/client/pfs"
	"github.com/pachyderm/pachyderm/src/client/pkg/pbutil"
	"github.com/pachyderm/pachyderm/src/client/pps"
	pfsserver "github.com/pachyderm/pachyderm/src/server/pfs"
	"github.com/pachyderm/pachyderm/src/server/pkg/backoff"
	col "github.com/pachyderm/pachyderm/src/server/pkg/collection"
	"github.com/pachyderm/pachyderm/src/server/pkg/dlock"
	"github.com/pachyderm/pachyderm/src/server/pkg/errutil"
	"github.com/pachyderm/pachyderm/src/server/pkg/obj"
	"github.com/pachyderm/pachyderm/src/server/pkg/ppsconsts"
	"github.com/pachyderm/pachyderm/src/server/pkg/ppsutil"
	filesync "github.com/pachyderm/pachyderm/src/server/pkg/sync"
	pfs_sync "github.com/pachyderm/pachyderm/src/server/pkg/sync"
	"github.com/pachyderm/pachyderm/src/server/pkg/watch"
	"github.com/pachyderm/pachyderm/src/server/worker/common"
	"github.com/pachyderm/pachyderm/src/server/worker/logs"
)

const (
	// maximumRetriesPerDatum is the maximum number of times each datum
	// can failed to be processed before we declare that the job has failed.
	maximumRetriesPerDatum = 3

	masterLockPath = "_master_worker_lock"

	// The number of datums the master caches
	numCachedDatums = 1000000

	ttl = int64(30)
)

func (a *APIServer) master(masterType string, spawner func(*client.APIClient) error) {
	logger := logs.NewMasterLogger(a.pipelineInfo)
	masterLock := dlock.NewDLock(a.etcdClient, path.Join(a.etcdPrefix, masterLockPath, a.pipelineInfo.Pipeline.Name, a.pipelineInfo.Salt))
	b := backoff.NewInfiniteBackOff()
	// Setting a high backoff so that when this master fails, the other
	// workers are more likely to become the master.
	// Also, we've observed race conditions where StopPipeline would cause
	// a master to restart before it's deleted.  PPS would then get confused
	// by the restart and create the workers again, because the restart would
	// bring the pipeline state from PAUSED to RUNNING.  By setting a high
	// retry interval, the master would be deleted before it gets a chance
	// to restart.
	b.InitialInterval = 10 * time.Second
	backoff.RetryNotify(func() error {
		// We use pachClient.Ctx here because it contains auth information.
		ctx, cancel := context.WithCancel(a.pachClient.Ctx())
		defer cancel() // make sure that everything this loop might spawn gets cleaned up
		ctx, err := masterLock.Lock(ctx)
		pachClient := a.pachClient.WithCtx(ctx)
		if err != nil {
			return err
		}
		defer masterLock.Unlock(ctx)
		logger.Logf("Launching %v master process", masterType)
		return spawner(pachClient)
	}, b, func(err error, d time.Duration) error {
		if auth.IsErrNotAuthorized(err) {
			logger.Logf("failing %q due to auth rejection", a.pipelineInfo.Pipeline.Name)
			return ppsutil.FailPipeline(a.pachClient.Ctx(), a.etcdClient, a.pipelines,
				a.pipelineInfo.Pipeline.Name, "worker master could not access output "+
					"repo to watch for new commits")
		}
		logger.Logf("master: error running the %v master process: %v; retrying in %v", masterType, err, d)
		return nil
	})
}
