package rpc

import (
	"context"
	"time"

	"github.com/rs/xid"
	uuid "github.com/satori/go.uuid"
	v1 "github.com/seayuns/pbnj/api/v1"
	"github.com/seayuns/pbnj/pkg/logging"
	"github.com/seayuns/pbnj/pkg/task"
)

// MachineService for doing power and device actions.
type MachineService struct {
	// Timeout is how long a task should be run
	// before it is cancelled. This is for use in a
	// TaskRunner.Execute function that runs all BMC
	// interactions in the background.
	Timeout    time.Duration
	TaskRunner task.Task
	v1.UnimplementedMachineServer
}

// BootDevice sets the next boot device of a machine.
func (m *MachineService) BootDevice(ctx context.Context, in *v1.DeviceRequest) (*v1.DeviceResponse, error) {
	l := logging.ExtractLogr(ctx)
	taskID := xid.New().String()
	l = l.WithValues("taskID", taskID)

	l.Info(
		"start BootDevice request",
		"username", in.Authn.GetDirectAuthn().GetUsername(),
		"vendor", in.Vendor.GetName(),
		"bootDevice", in.BootDevice.String(),
		"persistent", in.Persistent,
		"efiBoot", in.EfiBoot,
	)

	// execFunc := func(s chan string) (string, error) {
	// 	mbd, err := machine.NewBootDeviceSetter(
	// 		machine.WithDeviceRequest(in),
	// 		machine.WithLogger(l),
	// 		machine.WithStatusMessage(s),
	// 	)
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	// Because this is a background task, we want to pass through the span context, but not be
	// 	// a child context. This allows us to correctly plumb otel into the background task.
	// 	c := trace.ContextWithSpanContext(context.Background(), trace.SpanContextFromContext(ctx))
	// 	taskCtx, cancel := context.WithTimeout(c, m.Timeout)
	// 	defer cancel()
	// 	return mbd.BootDeviceSet(taskCtx, in.BootDevice.String(), in.Persistent, in.EfiBoot)
	// }
	// m.TaskRunner.Execute(ctx, l, "setting boot device", taskID, execFunc)

	// fake api
	taskMockId := uuid.NewV4().String()
	if in.GetBootDevice() == v1.BootDevice_BOOT_DEVICE_PXE {
		time.Sleep(10 * time.Second)
		return &v1.DeviceResponse{TaskId: "fake-task-pxe" + taskMockId}, nil
	}
	if in.GetBootDevice() == v1.BootDevice_BOOT_DEVICE_DISK {
		time.Sleep(10 * time.Second)
		return &v1.DeviceResponse{TaskId: "fake-task-disk" + taskMockId}, nil
	} else if in.GetBootDevice() == v1.BootDevice_BOOT_DEVICE_CDROM {
		time.Sleep(10 * time.Second)
		return &v1.DeviceResponse{TaskId: "fake-task-cdrom" + taskMockId}, nil
	} else if in.GetBootDevice() == v1.BootDevice_BOOT_DEVICE_BIOS {
		time.Sleep(10 * time.Second)
		return &v1.DeviceResponse{TaskId: "fake-task-bios" + taskMockId}, nil
	}
	return &v1.DeviceResponse{TaskId: taskID}, nil
}

// Power does a power action against a BMC.
func (m *MachineService) Power(ctx context.Context, in *v1.PowerRequest) (*v1.PowerResponse, error) {
	l := logging.ExtractLogr(ctx)
	taskID := xid.New().String()
	l = l.WithValues("taskID", taskID, "bmcIP", in.Authn.GetDirectAuthn().GetHost().GetHost())
	l.Info(
		"start Power request",
		"username", in.Authn.GetDirectAuthn().GetUsername(),
		"vendor", in.Vendor.GetName(),
		"powerAction", in.GetPowerAction().String(),
		"softTimeout", in.SoftTimeout,
		"OffDuration", in.OffDuration,
	)

	//fake api
	// execFunc := func(s chan string) (string, error) {
	// 	mp, err := machine.NewPowerSetter(
	// 		machine.WithPowerRequest(in),
	// 		machine.WithLogger(l),
	// 		machine.WithStatusMessage(s),
	// 	)
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	// Because this is a background task, we want to pass through the span context, but not be
	// 	// a child context. This allows us to correctly plumb otel into the background task.
	// 	c := trace.ContextWithSpanContext(context.Background(), trace.SpanContextFromContext(ctx))
	// 	taskCtx, cancel := context.WithTimeout(c, m.Timeout)
	// 	defer cancel()
	// 	return mp.PowerSet(taskCtx, in.PowerAction.String())
	// }
	// m.TaskRunner.Execute(ctx, l, "power action: "+in.GetPowerAction().String(), taskID, execFunc)

	//fake api
	taskMockId := uuid.NewV4().String()
	if in.GetPowerAction() == v1.PowerAction_POWER_ACTION_OFF {
		time.Sleep(10 * time.Second)
		return &v1.PowerResponse{
			TaskId: "fake-task-off" + taskMockId,
		}, nil
	} else if in.GetPowerAction() == v1.PowerAction_POWER_ACTION_ON {
		time.Sleep(10 * time.Second)
		return &v1.PowerResponse{
			TaskId: "fake-task-on" + taskMockId,
		}, nil
	} else if in.GetPowerAction() == v1.PowerAction_POWER_ACTION_RESET {
		time.Sleep(10 * time.Second)
		return &v1.PowerResponse{
			TaskId: "fake-task-reset" + taskMockId,
		}, nil
	} else if in.GetPowerAction() == v1.PowerAction_POWER_ACTION_CYCLE {
		time.Sleep(10 * time.Second)
		return &v1.PowerResponse{
			TaskId: "fake-task-cycle" + taskMockId,
		}, nil
	}
	return &v1.PowerResponse{TaskId: taskID}, nil
}
