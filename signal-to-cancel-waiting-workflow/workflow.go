package signal_to_cancel_waiting_workflow

import (
	"go.temporal.io/sdk/workflow"
	"time"
)

func SignalToCancelWaiting(ctx workflow.Context) error {
	var err error
	done := false
	for !done {
		timerCtx, timerCancelFunc := workflow.WithCancel(ctx)

		workflow.NewSelector(ctx).
			AddReceive(signalCh, func(c workflow.ReceiveChannel, more bool) {
				if !more {
					// Ignores channel close event, if any. Note that it would still interrupt the current wait and
					// restart it with the same waiting period.
					return
				}

				var newWaitingPeriodS int
				c.Receive(ctx, &newWaitingPeriodS)
				// Overrides the waiting period.
				waitingPeriod = time.Duration(newWaitingPeriodS) * time.Second
				*lastSlackMessageTime = workflow.Now(ctx)
				*batchStartTime = lastSlackMessageTime.Add(waitingPeriod)
				w.sendBeforeBatchSlackMessage(ctx, spec, batches, batch, *batchStartTime, true)
			}).
			AddFuture(workflow.NewTimer(timerCtx, waitingPeriod), func(workflow.Future) {
				done = true
			}).
			AddReceive(ctx.Done(), func(_ workflow.ReceiveChannel, _ bool) {
				done = true
				err = ctx.Err()
			}).
			Select(ctx)

		// Always cancels the timer at the end of each loop.
		timerCancelFunc()
	}

	return err
}
