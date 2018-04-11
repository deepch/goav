// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/frame.h>
//#include <stdlib.h>
import "C"
import (
	"image"
	"unsafe"
)

type (
	AvBuffer     C.struct_AVBuffer
	AvBufferRef  C.struct_AVBufferRef
	AvBufferPool C.struct_AVBufferPool
	//	Frame               C.struct_AVFrame
	AvFrameSideData     C.struct_AVFrameSideData
	AvFrameSideDataType C.enum_AVFrameSideDataType
)

func AvFrameGetBestEffortTimestamp(f *Frame) int64 {
	return int64(C.av_frame_get_best_effort_timestamp((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

func AvprivFrameGetMetadatap(f *Frame) **Dictionary {
	return (**Dictionary)(unsafe.Pointer(C.avpriv_frame_get_metadatap((*C.struct_AVFrame)(unsafe.Pointer(f)))))
}

func AvFrameSetQpTable(f *Frame, b *AvBufferRef, s, q int) int {
	return int(C.av_frame_set_qp_table((*C.struct_AVFrame)(unsafe.Pointer(f)), (*C.struct_AVBufferRef)(unsafe.Pointer(b)), C.int(s), C.int(q)))
}

func AvFrameGetQpTable(f *Frame, s, t *int) int8 {
	return int8(*C.av_frame_get_qp_table((*C.struct_AVFrame)(unsafe.Pointer(f)), (*C.int)(unsafe.Pointer(s)), (*C.int)(unsafe.Pointer(t))))
}

//Allocate an Frame and set its fields to default values.
func AvFrameAlloc() *Frame {
	return (*Frame)(unsafe.Pointer(C.av_frame_alloc()))
}

//Free the frame and any dynamically allocated objects in it, e.g.
func AvFrameFree(f *Frame) {
	var ptr *C.struct_AVFrame = (*C.struct_AVFrame)(unsafe.Pointer(f))
	C.av_frame_free((**C.struct_AVFrame)(&ptr))
}

//Allocate new buffer(s) for audio or video data.
func AvFrameGetBuffer(f *Frame, a int) int {
	return int(C.av_frame_get_buffer((*C.struct_AVFrame)(unsafe.Pointer(f)), C.int(a)))
}

//Setup a new reference to the data described by an given frame.
func AvFrameRef(d, s *Frame) int {
	return int(C.av_frame_ref((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(s))))
}

//Create a new frame that references the same data as src.
func AvFrameClone(f *Frame) *Frame {
	return (*Frame)(unsafe.Pointer(C.av_frame_clone((*C.struct_AVFrame)(unsafe.Pointer(f)))))
}

//Unreference all the buffers referenced by frame and reset the frame fields.
func AvFrameUnref(f *Frame) {
	cf := (*C.struct_AVFrame)(unsafe.Pointer(f))
	C.av_frame_unref(cf)
}

//Move everythnig contained in src to dst and reset src.
func AvFrameMoveRef(d, s *Frame) {
	C.av_frame_move_ref((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(s)))
}

//Check if the frame data is writable.
func AvFrameIsWritable(f *Frame) int {
	return int(C.av_frame_is_writable((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

//Ensure that the frame data is writable, avoiding data copy if possible.
func AvFrameMakeWritable(f *Frame) int {
	return int(C.av_frame_make_writable((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

//Copy only "metadata" fields from src to dst.
func AvFrameCopyProps(d, s *Frame) int {
	return int(C.av_frame_copy_props((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(s))))
}

//Get the buffer reference a given data plane is stored in.
func AvFrameGetPlaneBuffer(f *Frame, p int) *AvBufferRef {
	return (*AvBufferRef)(unsafe.Pointer(C.av_frame_get_plane_buffer((*C.struct_AVFrame)(unsafe.Pointer(f)), C.int(p))))
}

//Add a new side data to a frame.
func AvFrameNewSideData(f *Frame, d AvFrameSideDataType, s int) *AvFrameSideData {
	return (*AvFrameSideData)(unsafe.Pointer(C.av_frame_new_side_data((*C.struct_AVFrame)(unsafe.Pointer(f)), (C.enum_AVFrameSideDataType)(d), C.int(s))))
}

func AvFrameGetSideData(f *Frame, t AvFrameSideDataType) *AvFrameSideData {
	return (*AvFrameSideData)(unsafe.Pointer(C.av_frame_get_side_data((*C.struct_AVFrame)(unsafe.Pointer(f)), (C.enum_AVFrameSideDataType)(t))))
}

func Data(f *Frame) *uint8 {
	return (*uint8)(unsafe.Pointer((*C.uint8_t)(unsafe.Pointer(&f.Data))))
}

func Linesize(f *Frame) int {
	return int(*(*C.int)(unsafe.Pointer(&f.Linesize)))
}

// Width returns the frame width
func (f *Frame) Width() int {
	return int(f.width)
}

// SetWidth sets the frame width
func (f *Frame) SetWidth(w int) {
	f.width = int32(w)
}

// Height returns the frame height
func (f *Frame) Height() int {
	return int(f.height)
}

// SetHeight setes the frame height
func (f *Frame) SetHeight(h int) {
	f.height = int32(h)
}

// Format returns the frame's format.
func (f *Frame) Format() int {
	return int(f.format)
}

// SetFormat sets the frame's format.
func (f *Frame) SetFormat(format int) {
	f.format = int32(format)
}

//SetPicture creates a YCbCr image from the frame
func SetPicture(f *Frame, img *image.YCbCr) {
	f.Data[0] = &img.Y[0]
	f.Data[1] = &img.Cb[0]
	f.Data[2] = &img.Cr[0]
	f.Linesize[0] = int32(img.YStride)
	f.Linesize[1] = int32(img.CStride)
	f.Linesize[2] = int32(img.CStride)
}

//GetPicture creates a YCbCr image from the frame
func GetPicture(f *Frame) (img *image.YCbCr, err error) {
	// For 4:4:4, CStride == YStride/1 && len(Cb) == len(Cr) == len(Y)/1.
	// For 4:2:2, CStride == YStride/2 && len(Cb) == len(Cr) == len(Y)/2.
	// For 4:2:0, CStride == YStride/2 && len(Cb) == len(Cr) == len(Y)/4.
	// For 4:4:0, CStride == YStride/1 && len(Cb) == len(Cr) == len(Y)/2.
	// For 4:1:1, CStride == YStride/4 && len(Cb) == len(Cr) == len(Y)/4.
	// For 4:1:0, CStride == YStride/4 && len(Cb) == len(Cr) == len(Y)/8.

	w := int(f.Linesize[0])
	h := int(f.height)
	r := image.Rectangle{image.Point{0, 0}, image.Point{w, h}}
	// TODO: Use the sub sample ratio from the input image 'f.format'
	img = image.NewYCbCr(r, image.YCbCrSubsampleRatio420)
	// convert the frame data data to a Go byte array
	img.Y = C.GoBytes(unsafe.Pointer(f.Data[0]), C.int(w*h))

	wCb := int(f.Linesize[1])
	if unsafe.Pointer(f.Data[1]) != nil {
		img.Cb = C.GoBytes(unsafe.Pointer(f.Data[1]), C.int(wCb*h/2))
	}

	wCr := int(f.Linesize[2])
	if unsafe.Pointer(f.Data[2]) != nil {
		img.Cr = C.GoBytes(unsafe.Pointer(f.Data[2]), C.int(wCr*h/2))
	}
	return
}

/*
func GetPicture(f *Frame) (img *image.YCbCr, err error) {
	// For 4:4:4, CStride == YStride/1 && len(Cb) == len(Cr) == len(Y)/1.
	// For 4:2:2, CStride == YStride/2 && len(Cb) == len(Cr) == len(Y)/2.
	// For 4:2:0, CStride == YStride/2 && len(Cb) == len(Cr) == len(Y)/4.
	// For 4:4:0, CStride == YStride/1 && len(Cb) == len(Cr) == len(Y)/2.
	// For 4:1:1, CStride == YStride/4 && len(Cb) == len(Cr) == len(Y)/4.
	// For 4:1:0, CStride == YStride/4 && len(Cb) == len(Cr) == len(Y)/8.

	w := int(f.linesize[0])
	h := int(f.height)
	r := image.Rectangle{image.Point{0, 0}, image.Point{w, h}}
	// TODO: Use the sub sample ratio from the input image 'f.format'
	img = image.NewYCbCr(r, image.YCbCrSubsampleRatio420)
	// convert the frame data data to a Go byte array

	yb_tmp := C.GoBytes(unsafe.Pointer(f.data[0]), C.int(w*h))
	yb := make([]uint8, len(yb_tmp))
	copy(yb, yb_tmp)
	img.Y = yb
	wCb := int(f.linesize[1])
	if unsafe.Pointer(f.data[1]) != nil {

		//img.Cb = C.GoBytes(unsafe.Pointer(f.data[1]), C.int(wCb*h/2))
		cb_tmp := C.GoBytes(unsafe.Pointer(f.data[1]), C.int(wCb*h/2))
		cb := make([]uint8, len(cb_tmp))
		copy(cb, cb_tmp)
		img.Cb = cb
	}

	wCr := int(f.linesize[2])
	if unsafe.Pointer(f.data[2]) != nil {
		//img.Cr = C.GoBytes(unsafe.Pointer(f.data[2]), C.int(wCr*h/2))

		cr_tmp := C.GoBytes(unsafe.Pointer(f.data[2]), C.int(wCr*h/2))
		cr := make([]uint8, len(cr_tmp))
		copy(cr, cr_tmp)
		img.Cr = cr
	}
	return
}
*/
func GetPictureRGB(f *Frame) (img *image.RGBA, err error) {
	w := int(f.Linesize[0])
	h := int(f.height)
	r := image.Rectangle{image.Point{0, 0}, image.Point{w, h}}
	// TODO: Use the sub sample ratio from the input image 'f.format'
	img = image.NewRGBA(r)
	// convert the frame data data to a Go byte array
	img.Pix = C.GoBytes(unsafe.Pointer(f.Data[0]), C.int(w*h))
	img.Stride = w
	return
}

// //static int get_video_buffer (Frame *frame, int align)
// func GetVideoBuffer(f *Frame, a int) int {
// 	return int(C.get_video_buffer(f, C.int(a)))
// }

// //static int get_audio_buffer (Frame *frame, int align)
// func GetAudioBuffer(f *Frame, a int) int {
// 	return C.get_audio_buffer(f, C.int(a))
// }

// //static void get_frame_defaults (Frame *frame)
// func GetFrameDefaults(f *Frame) {
// 	C.get_frame_defaults(*C.struct_AVFrame(f))
// }
