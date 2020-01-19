package rdtp

func (p *Packet) computeChecksum() uint16 {
	csum := uint16(0)
	csum += p.SrcPort
	csum += p.DstPort
	csum += p.Length
	for i := 0; i < len(p.Payload); i++ {
		csum += uint16(p.Payload[i])
	}
	return ^csum
}

func (p *Packet) verifyChecksum() bool {
	return p.Checksum == p.computeChecksum()
}