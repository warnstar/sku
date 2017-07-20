package SkuRun


func (s *Server) CheckPiAllTimeSync() bool {
	timeSyncNum := 0
	for _, onePi := range s.Pis {
		if onePi.IsTimeSync {
			timeSyncNum++
		}
	}

	if timeSyncNum >= s.PiMaxNum {
		s.IsAllTimeSync = true

		return true
	}

	return false
}

func (s *Server) CheckPiAllWriteKb() bool {
	Num := 0
	for _, onePi := range s.Pis {
		if onePi.IsWriteKbFinish {
			Num++
		}
	}

	if Num >= s.PiMaxNum {
		s.IsAllWriteKb = true
		return true
	}

	return false
}

func (s *Server) CheckPiAllSendResult() bool {
	Num := 0
	for _, onePi := range s.Pis {
		if onePi.IsSendResult {
			Num++
		}
	}

	if Num >= s.PiMaxNum {
		s.IsAllSendResult = true
		return true
	}

	return false
}


func (s *Server) CheckPiAllConnected() bool {
	return s.IsAllConnected
}