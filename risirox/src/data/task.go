package data

import iserver "risirox/risirox/src/server"

type Task struct {
	Handler func(message iserver.IRequest)

	Param iserver.IRequest
}
