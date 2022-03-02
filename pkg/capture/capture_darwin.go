package capture

/*
#cgo CFLAGS: -x objective-c -Wdeprecated-declarations
#cgo LDFLAGS: -framework Cocoa -framework AVFoundation -framework CoreMedia
#import <AVFoundation/AVFoundation.h>
#include <stdio.h>
#include "types.h"

capture_device *device_at(capture_device **ptr, int idx) {
    return ptr[idx];
}

capture_format *format_at(capture_device *ptr, int idx) {
    return ptr->formats[idx];
}

int list_devices(struct devices *dev) {
	if (@available(macOS 10.15, *)) {
        AVCaptureDeviceDiscoverySession *discoverySession = [
                AVCaptureDeviceDiscoverySession
                discoverySessionWithDeviceTypes:
                    @[AVCaptureDeviceTypeBuiltInWideAngleCamera, AVCaptureDeviceTypeExternalUnknown]
                mediaType:NULL
                position:AVCaptureDevicePositionUnspecified
            ];
        int dcount = [[discoverySession devices] count];
        dev->count = dcount;
        if (dcount == 0) {
            return 0;
        }
        dev->devices = (capture_device **)malloc(sizeof(capture_device *) * dcount);
        for (int ci = 0; ci < dcount; ci++) {
            dev->devices[ci] = (capture_device *)malloc(sizeof(capture_device));

            AVCaptureDevice *device = discoverySession.devices[ci];
            dev->devices[ci]->name = [[device localizedName] UTF8String];
            dev->devices[ci]->model = [[device modelID] UTF8String];
            dev->devices[ci]->manufacturer = [[device manufacturer] UTF8String];

            int fcount = [device.formats count];
            dev->devices[ci]->format_count = fcount;
			dev->devices[ci]->formats = (capture_format **)malloc(sizeof(capture_format *) * fcount);
 			for (int fi = 0; fi < dev->devices[ci]->format_count; fi++) {
				dev->devices[ci]->formats[fi] = (capture_format *)malloc(sizeof(capture_format));
	 			AVCaptureDeviceFormat *format = device.formats[fi];
				CMFormatDescriptionRef fd = [format formatDescription];
				CMVideoDimensions dimension = CMVideoFormatDescriptionGetDimensions(fd);
                dev->devices[ci]->formats[fi]->width = dimension.width;
                dev->devices[ci]->formats[fi]->height = dimension.height;
                dev->devices[ci]->formats[fi]->codec = CMFormatDescriptionGetMediaSubType(fd);

				dev->devices[ci]->formats[fi]->min_frame_rate = 10000;
				dev->devices[ci]->formats[fi]->max_frame_rate = 0;

				for (AVFrameRateRange *range in format.videoSupportedFrameRateRanges) {
                    if (range.maxFrameRate > dev->devices[ci]->formats[fi]->max_frame_rate) {
						dev->devices[ci]->formats[fi]->max_frame_rate = range.maxFrameRate;
					}
                    if (range.minFrameRate < dev->devices[ci]->formats[fi]->min_frame_rate) {
						dev->devices[ci]->formats[fi]->min_frame_rate = range.minFrameRate;
					}
				}
			}
		}
    } else {
    	// TODO: do we really need support older versions
    	//NSArray *devices = [AVCaptureDevice devicesWithMediaType:AVMediaTypeVideo];
		//for (AVCaptureDevice *device in devices) {
    	//        const char *name = [[device localizedName] UTF8String];
    	//        index            = [devices indexOfObject:device];
    	//        printf("[%d] %s\n", index, name);
    	//}
		return -1;
	}
	return 0;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// ListCameras returns camera and its resolution
func ListCameras() ([]Device, error){
	var tmp *C.struct_devices
	dev := (*C.struct_devices)(C.malloc(C.size_t(unsafe.Sizeof(tmp))))
	ret := C.list_devices(dev)
	defer C.free(unsafe.Pointer(dev))
	defer C.free(unsafe.Pointer(dev.devices))

	if ret < 0 {
		return nil, fmt.Errorf("os version requires 10.15 at least")
	}

	retDevices := make([]Device, int(dev.count))
	for i := 0; i < int(dev.count); i++ {
		fmt.Println(i)

		device := C.device_at(dev.devices, C.int(i))
		defer C.free(unsafe.Pointer(device))
		defer C.free(unsafe.Pointer(device.formats))

		retDevices[i].Name = C.GoString(device.name)
		retDevices[i].Model = C.GoString(device.model)
		retDevices[i].Manufacturer = C.GoString(device.manufacturer)
		retDevices[i].Formats = make([]Format, int(device.format_count))

		for j := 0; j < int(device.format_count); j++ {
			format := C.format_at(device, C.int(j))
			defer C.free(unsafe.Pointer(format))

			retDevices[i].Formats[j].Codec = int(format.codec)
			retDevices[i].Formats[j].Width = uint(format.width)
			retDevices[i].Formats[j].Height = uint(format.height)
			retDevices[i].Formats[j].MinFrameRate = float32(format.min_frame_rate)
			retDevices[i].Formats[j].MaxFrameRate = float32(format.max_frame_rate)
		}
	}
	fmt.Println(retDevices)
	return retDevices, nil
}