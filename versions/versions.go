package main

import (
	"log"

	"github.com/deepch/goav/avcodec"
	"github.com/deepch/goav/avdevice"
	"github.com/deepch/goav/avfilter"
	"github.com/deepch/goav/avformat"
	"github.com/deepch/goav/avutil"
	"github.com/deepch/goav/swresample"
	"github.com/deepch/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
