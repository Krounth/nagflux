package collector

import "github.com/Krounth/nagflux/data"

type ResultQueues map[data.Target]chan Printable
