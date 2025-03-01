// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package rtcp

import "errors"

var (
	errWrongMarshalSize         = errors.New("rtcp: wrong marshal size")
	errInvalidTotalLost         = errors.New("rtcp: invalid total lost count")
	errInvalidHeader            = errors.New("rtcp: invalid header")
	errEmptyCompound            = errors.New("rtcp: empty compound packet")
	errBadFirstPacket           = errors.New("rtcp: first packet in compound must be SR or RR")
	errMissingCNAME             = errors.New("rtcp: compound missing SourceDescription with CNAME")
	errPacketBeforeCNAME        = errors.New("rtcp: feedback packet seen before CNAME")
	errTooManyReports           = errors.New("rtcp: too many reports")
	errTooManyChunks            = errors.New("rtcp: too many chunks")
	errTooManySources           = errors.New("rtcp: too many sources")
	errPacketTooShort           = errors.New("rtcp: packet too short")
	errWrongType                = errors.New("rtcp: wrong packet type")
	errSDESTextTooLong          = errors.New("rtcp: sdes must be < 255 octets long")
	errSDESMissingType          = errors.New("rtcp: sdes item missing type")
	errReasonTooLong            = errors.New("rtcp: reason must be < 255 octets long")
	errBadVersion               = errors.New("rtcp: invalid packet version")
	errWrongPadding             = errors.New("rtcp: invalid padding value")
	errWrongFeedbackType        = errors.New("rtcp: wrong feedback message type")
	errWrongPayloadType         = errors.New("rtcp: wrong payload type")
	errHeaderTooSmall           = errors.New("rtcp: header length is too small")
	errSSRCMustBeZero           = errors.New("rtcp: media SSRC must be 0")
	errMissingREMBidentifier    = errors.New("missing REMB identifier")
	errSSRCNumAndLengthMismatch = errors.New("SSRC num and length do not match")
	errInvalidSizeOrStartIndex  = errors.New("invalid size or startIndex")
	errInvalidBitrate           = errors.New("invalid bitrate")
	errWrongChunkType           = errors.New("rtcp: wrong chunk type")
	errBadStructMemberType      = errors.New("rtcp: struct contains unexpected member type")
	errBadReadParameter         = errors.New("rtcp: cannot read into non-pointer")
)
