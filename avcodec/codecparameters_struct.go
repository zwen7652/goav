package avcodec

import "C"

func (cp *AvCodecParameters) CodecType() MediaType {
	return MediaType(cp.codec_type)
}

func (cp *AvCodecParameters) CodecId() CodecId {
	return CodecId(cp.codec_id)
}

func (cp *AvCodecParameters) SetCodecTag(tag uint) {
	cp.codec_tag = C.uint(tag)
}
