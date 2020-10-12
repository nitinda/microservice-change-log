package channels

// ValidateChannel to validate channel status
// return true when validated successfully.
// return false when validated unsuccessfully.
func ValidateChannel(done chan bool) bool {
	select {
	case ok := <-done:
		if ok {
			return true
		}
	}
	return false
}
