package balancer

import (
	// "crypto/sha256"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/SemenchenkoVitaliy/GDC-CP/netutils"
)

var maxInt uint

func init() {
	maxInt--
}

type ServerInfo struct {
	Info   netutils.AuthData
	Server netutils.Server
}

type ServerInfoC struct {
	Info        netutils.AuthData
	Server      netutils.Server
	connections int
}

/******************************************************************************
	Round-robin balancing
******************************************************************************/

type RoundRobin struct {
	servers []ServerInfo
	ids     map[int]int
	idGener int
	idCur   int
}

func NewRoundRobin() (balancer *RoundRobin) {
	return &RoundRobin{
		servers: []ServerInfo{},
		ids:     make(map[int]int),
	}
}

func (b *RoundRobin) Add(server netutils.Server, data netutils.AuthData) (id int) {
	id = b.idGener
	b.idGener++
	b.idCur = 0
	b.ids[id] = len(b.servers)
	b.servers = append(b.servers, ServerInfo{
		Info:   data,
		Server: server,
	})
	return id
}

func (b *RoundRobin) Remove(id int) {
	b.idCur = 0
	index := b.ids[id]
	delete(b.ids, id)
	for k, v := range b.ids {
		if v > index {
			b.ids[k]--
		}
	}
	b.servers = append(b.servers[:index], b.servers[index+1:]...)
}

func (b *RoundRobin) GetAll() (servers []ServerInfo, ok bool) {
	if len(b.servers) == 0 {
		return b.servers, false
	}
	return b.servers, true
}

func (b *RoundRobin) GetOne() (server ServerInfo, ok bool) {
	if len(b.servers) == 0 {
		return server, false
	}
	server = b.servers[b.idCur]
	b.idCur++
	if b.idCur == len(b.servers) {
		b.idCur = 0
	}
	return server, true
}

/******************************************************************************
	Random balancing
******************************************************************************/

type Random struct {
	servers []ServerInfo
	ids     map[int]int
	idGener int
}

func NewRandom() (balancer *Random) {
	return &Random{
		servers: []ServerInfo{},
		ids:     make(map[int]int),
	}
}

func (b *Random) Add(server netutils.Server, data netutils.AuthData) (id int) {
	id = b.idGener
	b.idGener++
	b.ids[id] = len(b.servers)
	b.servers = append(b.servers, ServerInfo{
		Info:   data,
		Server: server,
	})
	return id
}

func (b *Random) Remove(id int) {
	index := b.ids[id]
	delete(b.ids, id)
	for k, v := range b.ids {
		if v > index {
			b.ids[k]--
		}
	}
	b.servers = append(b.servers[:index], b.servers[index+1:]...)
}

func (b *Random) GetAll() (servers []ServerInfo, ok bool) {
	if len(b.servers) == 0 {
		return b.servers, false
	}
	return b.servers, true
}

func (b *Random) GetOne(r *http.Request) (server ServerInfo, ok bool) {
	if len(b.servers) == 0 {
		return server, false
	}
	randNum := rand.Int()
	return b.servers[randNum%len(b.servers)], true
}

/******************************************************************************
	IP hash balancing
******************************************************************************/

type IPHash struct {
	servers []ServerInfo
	ids     map[int]int
	idGener int
}

func NewIPHash() (balancer *IPHash) {
	return &IPHash{
		servers: []ServerInfo{},
		ids:     make(map[int]int),
	}
}

func (b *IPHash) Add(server netutils.Server, data netutils.AuthData) (id int) {
	id = b.idGener
	b.idGener++
	b.ids[id] = len(b.servers)
	b.servers = append(b.servers, ServerInfo{
		Info:   data,
		Server: server,
	})
	return id
}

func (b *IPHash) Remove(id int) {
	index := b.ids[id]
	delete(b.ids, id)
	for k, v := range b.ids {
		if v > index {
			b.ids[k]--
		}
	}
	b.servers = append(b.servers[:index], b.servers[index+1:]...)
}

func (b *IPHash) GetAll() (servers []ServerInfo, ok bool) {
	if len(b.servers) == 0 {
		return b.servers, false
	}
	return b.servers, true
}

func (b *IPHash) GetOne(r *http.Request) (server ServerInfo, ok bool) {
	if len(b.servers) == 0 {
		return server, false
	}

	hashNum, _ := strconv.Atoi(strings.Replace(r.URL.Host, ":", "", -1))
	return b.servers[hashNum%len(b.servers)], true
}

/*****************************************************************************
	Url path balancing
*****************************************************************************/

type UrlHash struct {
	servers []ServerInfo
	ids     map[int]int
	idGener int
}

func NewUrlHash() (balancer *UrlHash) {
	return &UrlHash{
		servers: []ServerInfo{},
		ids:     make(map[int]int),
	}
}

func (b *UrlHash) Add(server netutils.Server, data netutils.AuthData) (id int) {
	id = b.idGener
	b.idGener++
	b.ids[id] = len(b.servers)
	b.servers = append(b.servers, ServerInfo{
		Info:   data,
		Server: server,
	})
	return id
}

func (b *UrlHash) Remove(id int) {
	index := b.ids[id]
	delete(b.ids, id)
	for k, v := range b.ids {
		if v > index {
			b.ids[k]--
		}
	}
	b.servers = append(b.servers[:index], b.servers[index+1:]...)
}

func (b *UrlHash) GetAll() (servers []ServerInfo, ok bool) {
	if len(b.servers) == 0 {
		return b.servers, false
	}
	return b.servers, true
}

// func (b *UrlHash) GetOne(r *http.Request) (server ServerInfo, ok bool) {
// 	if len(b.servers) == 0 {
// 		return server, false
// 	}
// 	hashNum, _ := strconv.ParseInt(string(sha256.Sum256([]byte(r.URL.Path))), 16, 0)
// 	return b.servers[hashNum%len(b.servers)], true
// }

/*****************************************************************************
	Least connections balancing
*****************************************************************************/

type LeastConnections struct {
	servers []ServerInfoC
	ids     map[int]int
	idGener int
}

func NewLeastConnections() (balancer *LeastConnections) {
	return &LeastConnections{
		servers: []ServerInfoC{},
		ids:     make(map[int]int),
	}
}

func (b *LeastConnections) Add(server netutils.Server, data netutils.AuthData) (id int) {
	id = b.idGener
	b.idGener++
	b.ids[id] = len(b.servers)
	b.servers = append(b.servers, ServerInfoC{
		Info:        data,
		Server:      server,
		connections: 0,
	})
	return id
}

func (b *LeastConnections) Remove(id int) {
	index := b.ids[id]
	delete(b.ids, id)
	for k, v := range b.ids {
		if v > index {
			b.ids[k]--
		}
	}
	b.servers = append(b.servers[:index], b.servers[index+1:]...)
}

func (b *LeastConnections) GetAll() (servers []ServerInfoC, ok bool) {
	if len(b.servers) == 0 {
		return b.servers, false
	}
	return b.servers, true
}

func (b *LeastConnections) GetOne() (server ServerInfoC, id int, ok bool) {
	if len(b.servers) == 0 {
		return server, id, false
	}
	var min, iMin int

	for k, v := range b.ids {
		if b.servers[v].connections == 0 {
			id = k
			iMin = v
			break
		} else if b.servers[v].connections < min {
			id = k
			iMin = v
		}
	}

	b.servers[iMin].connections++
	return b.servers[iMin], id, true
}

func (b *LeastConnections) ReleaseOne(id int) {
	if index, ok := b.ids[id]; ok {
		b.servers[index].connections--
	}
}

/****************************************************************************
	Distributed balancing
****************************************************************************/

type Distributed struct {
	servers map[int]ServerInfo
	ids     map[int]uint
	idArr   []int
	idCur   int
}

func NewDistributed() (balancer *Distributed) {
	return &Distributed{
		servers: make(map[int]ServerInfo),
		ids:     make(map[int]uint),
		idArr:   []int{},
	}
}

func (b *Distributed) refreshIdArr() {
	b.idArr = make([]int, len(b.ids))
	i := 0
	for k := range b.ids {
		b.idArr[i] = k
		i++
	}
}

func (b *Distributed) Add(server netutils.Server, data netutils.AuthData) (id int, ok bool) {
	id = data.Id
	if _, ok := b.ids[id]; ok {
		return id, false
	}

	b.ids[id] = 0
	b.servers[id] = ServerInfo{
		Info:   data,
		Server: server,
	}
	b.refreshIdArr()
	return id, true
}

func (b *Distributed) Remove(id int) {
	delete(b.ids, id)
	delete(b.servers, id)
	b.refreshIdArr()
}

func (b *Distributed) GetAll() (servers map[int]ServerInfo, ok bool) {
	if len(b.servers) == 0 {
		return b.servers, false
	}
	return b.servers, true
}

func (b *Distributed) GetNIds(n int) (ids []int, ok bool) {
	size := len(b.ids)
	if size == 0 {
		return ids, false
	} else if size <= n {
		return b.idArr, true
	}

	idCur := b.idCur
	ids = make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = idCur
		if idCur++; idCur >= size {
			idCur = 0
		}
	}

	if b.idCur++; b.idCur >= size {
		b.idCur = 0
	}
	return ids, true
}

func (b *Distributed) GetOneFrom(ids []int) (server ServerInfo, ok bool) {
	if len(b.servers) == 0 || len(ids) == 0 {
		return server, false
	}

	var (
		min   uint = maxInt
		idMin int  = -1
	)

	for _, id := range ids {
		if load, ok := b.ids[id]; ok && load < min {
			idMin = id
			min = load
		}
	}

	if idMin == -1 {
		return server, false
	}

	b.ids[idMin]++
	return b.servers[idMin], true
}

func (b *Distributed) GetOne(id int) (server ServerInfo, ok bool) {
	server, ok = b.servers[id]
	return server, ok
}
