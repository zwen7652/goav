package avcodec

//#include <libavcodec/avcodec.h>
import "C"

func AvcodecParametersToContext(codecContext *Context, codecParameters *AvCodecParameters) int {
	return int(C.avcodec_parameters_to_context((*C.struct_AVCodecContext)(codecContext), (*C.struct_AVCodecParameters)(codecParameters)))
}

func AvcodecParametersFromContext(codecParameters *AvCodecParameters, codecContext *Context) int {
	return int(C.avcodec_parameters_from_context((*C.struct_AVCodecParameters)(codecParameters), (*C.struct_AVCodecContext)(codecContext)))
}

func AvcodecParametersCopy(out, in *AvCodecParameters) int {
	return int(C.avcodec_parameters_copy((*C.struct_AVCodecParameters)(out), (*C.struct_AVCodecParameters)(in)))
}
