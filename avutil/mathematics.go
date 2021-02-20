package avutil

//#include <libavutil/mathematics.h>
import "C"
import "github.com/zwen7652/goav/avcodec"

const (
	AV_ROUND_NEAR_INF    = C.AV_ROUND_NEAR_INF
	AV_ROUND_PASS_MINMAX = C.AV_ROUND_PASS_MINMAX
)

func AvRescaleQ(a int64, bq, cq avcodec.Rational) int64 {
	return (int64)(C.av_rescale_q(C.int64_t(a), (C.struct_AVRational)(bq), (C.struct_AVRational)(cq)))
}

func AvRescaleQRnd(a int64, bq, cq avcodec.Rational, flags uint32) int64 {
	return (int64)(C.av_rescale_q_rnd(C.int64_t(a), (C.struct_AVRational)(bq), (C.struct_AVRational)(cq), flags))
}
