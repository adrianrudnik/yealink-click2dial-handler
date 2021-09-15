package yealink

func call(device *Device, targetNumber, outgoingAccount string) error {
	// @todo account support, we just assume the same as the device IP right now

	_, err := device.GetCallEndpointUrl()
	if err != nil {
		return err
	}

	// #http://10.3.20.10/servlet?key=number=1234&outgoing_uri=1006@10.2.1.48
	return nil
}
