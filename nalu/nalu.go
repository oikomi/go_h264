package nalu

type NALU struct {
	StartCodePrefixLen int //! 4 for parameter sets and first slice in picture, 3 for everything else (suggested)
	Len                int //! Length of the NAL unit (Excluding the start code, which does not belong to the NALU)
	ForbiddenBit       int //! should be always FALSE
	NalReferenceIdc    int //! NALU_PRIORITY_xxxx
	NalUnitType        int //! NALU_TYPE_xxxx
}


