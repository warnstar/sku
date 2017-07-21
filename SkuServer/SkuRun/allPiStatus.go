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
	} else {
		s.IsAllTimeSync = false
	}

	return s.IsAllTimeSync
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
	} else {
		s.IsAllWriteKb = false
	}

	return s.IsAllWriteKb
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
	} else {
		s.IsAllSendResult = false
	}

	return s.IsAllSendResult
}

func (s *Server) CheckPiAllConnected() bool {
	if s.PiCurNum >= s.PiMaxNum {
		s.IsAllConnected = true
	} else {
		s.IsAllConnected = false
	}
	return s.IsAllConnected
}