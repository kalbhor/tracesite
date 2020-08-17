package tracesite

const DEFAULT_PORT = 33434
const DEFAULT_PACKET_SIZE = 52
const DEFAULT_TIMEOUT_MS = 4000
const DEFAULT_RETRIES = 3
const DEFAULT_DEST = "google.com"
const DEFAULT_MAX_HOPS = 24
const DEFAULT_START_TTL = 1

type Options struct {
	dest       string
	port       int
	packetSize int
	retries    int
	timeoutMs  int
	maxHops    int
	startTTL   int
}

func (o *Options) Dest() string {

	if o.dest == "" {
		o.dest = DEFAULT_DEST
	}

	return o.dest
}

func (o *Options) SetDest(dest string) {

	o.dest = dest
}

func (o *Options) StartTTL() int {

	if o.startTTL == 0 {
		o.startTTL = DEFAULT_START_TTL
	}

	return o.startTTL
}

func (o *Options) SetStartTTL(startTTL int) {
	o.startTTL = startTTL
}

func (o *Options) Port() int {

	if o.port == 0 {
		o.port = DEFAULT_PORT
	}

	return o.port
}

func (o *Options) SetPort(port int) {
	o.port = port
}

func (o *Options) PacketSize() int {

	if o.packetSize == 0 {
		o.packetSize = DEFAULT_PACKET_SIZE
	}

	return o.packetSize
}

func (o *Options) SetPacketSize(packetSize int) {
	o.packetSize = packetSize
}

func (o *Options) Retries() int {

	if o.retries == 0 {
		o.retries = DEFAULT_RETRIES
	}

	return o.retries
}

func (o *Options) SetRetries(retries int) {
	o.retries = retries
}

func (o *Options) TimeoutMs() int {

	if o.timeoutMs == 0 {
		o.timeoutMs = DEFAULT_TIMEOUT_MS
	}

	return o.timeoutMs
}

func (o *Options) SetTimeoutMs(timeoutMs int) {
	o.timeoutMs = timeoutMs
}

func (o *Options) MaxHops() int {

	if o.maxHops == 0 {
		o.maxHops = DEFAULT_MAX_HOPS
	}

	return o.timeoutMs
}

func (o *Options) SetMaxHops(maxHops int) {
	o.maxHops = maxHops
}
