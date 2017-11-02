package main

import (
	"log"

	"github.com/koropets/goav/avcodec"
	"github.com/koropets/goav/avdevice"
	"github.com/koropets/goav/avfilter"
	"github.com/koropets/goav/avformat"
	"github.com/koropets/goav/avutil"
	"github.com/koropets/goav/swresample"
	"github.com/koropets/goav/swscale"
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
