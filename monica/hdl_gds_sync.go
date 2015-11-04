package monica

import (
	"encoding/csv"
	"fmt"
	"github.com/cmingxu/monica/protogos/common"
	"github.com/cmingxu/monica/protogos/gds"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

type HdlGdsSync struct {
	session *Session
}

func NewGdsSyncHandler(session *Session) *HdlGdsSync {
	return &HdlGdsSync{
		session: session,
	}
}

func (handler *HdlGdsSync) HandlePackage(buf []byte) {
	gdsSync := new(common.GdsSync)
	proto.Unmarshal(buf, gdsSync)
	fmt.Println(gdsSync.GetVersion())
	fmt.Println(gdsSync.GetWhat())
	buildingcsv := filepath.Join(handler.session.Server.Config.GdsPath, "buildings.csv")
	log.Println(buildingcsv)
	csvContent, _ := ioutil.ReadFile(buildingcsv)

	r2 := csv.NewReader(strings.NewReader(string(csvContent)))
	ss, _ := r2.ReadAll()
	sz := len(ss)
	buildingGds := new(gds.BuildingGds)
	buildingGds.Buildings = []*gds.Building{}
	for i := 0; i < sz; i++ {
		building := new(gds.Building)
		building.Name = &(ss[i][0])
		intVal, _ := strconv.Atoi(ss[i][1])
		v := uint32(intVal)
		building.Level = &v
		buildingGds.Buildings = append(buildingGds.Buildings, building)
	}

	bytes, _ := proto.Marshal(buildingGds)
	log.Println("xxxxx", len(bytes))
	handler.session.WriteToClient(ProtoBuildingsGds, bytes)
}
