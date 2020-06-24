package printer

import (
	"fmt"
	"log"

	"github.com/google/gousb"
)

type UsbPrinter struct {
	context  *gousb.Context
	endpoint *gousb.OutEndpoint
}

func (p *UsbPrinter) Write(buf []byte) (err error) {
	_, err = p.endpoint.Write(buf)
	return
}

func (p *UsbPrinter) Close() error {
	return p.context.Close()
}

// TODO: Use device class to select printer device
const EPSON_VENDOR = 0x4b8
const PRODUCT_ID = 0xE28

func getOutEndpoint(device *gousb.Device) (outEndpoint *gousb.OutEndpoint, err error) {
	inf, _, err := device.DefaultInterface()
	if err != nil {
		log.Fatalf("Could not get default interface: %v", err)
		return
	}

	for _, ep := range inf.Setting.Endpoints {
		if ep.Direction == gousb.EndpointDirectionOut {
			return inf.OutEndpoint(ep.Number)
		}
	}

	return
}

func GetPrinter() (printer *UsbPrinter, err error) {
	ctx := gousb.NewContext()

	device, err := ctx.OpenDeviceWithVIDPID(EPSON_VENDOR, PRODUCT_ID)
	if err == nil && device == nil {
		err = fmt.Errorf("Nil printer")
	}

	if err != nil {
		log.Fatalf("Could not open a device: %v", err)
		return
	}

	outEndpoint, err := getOutEndpoint(device)
	if err != nil {
		log.Fatalf("Could not get out endpoint: %v", err)
		return
	}

	printer = &UsbPrinter{
		context:  ctx,
		endpoint: outEndpoint,
	}

	return
}
