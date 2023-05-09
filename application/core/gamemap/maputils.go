package gamemap

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type AABB struct {
	Min Corner
	Max Corner
}

type Ray struct {
	Origin, Direction Corner
}

type JsonAABB struct {
	Min string `json:"Min"`
	Max string `json:"Max"`
}

type JsonEntity struct {
	AABB JsonAABB `json:"AABB"`
}

type Corner struct {
	X float32
	Y float32
	Z float32
}
type Entity struct {
	AABB AABB
}
type JsonData struct {
	Entities map[string]Entity
}

var Maputils *JsonData
var MapFilePath = "application/conf/map.json"

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (gmap *JsonData) LoadMap() {
	var entities map[string]JsonEntity
	if confFileExists, _ := PathExists(MapFilePath); confFileExists != true {
		//fmt.Println("Config File ", g.ConfFilePath , " is not exist!!")
		return
	}
	data, err := os.ReadFile(MapFilePath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &entities)
	if err != nil {
		panic(err)
	}

	for key, entity := range entities {
		min := strings.Split(entity.AABB.Min, " ")
		max := strings.Split(entity.AABB.Max, " ")

		xMin := parseFloat(min[0], "x")
		yMin := parseFloat(min[1], "y")
		zMin := parseFloat(min[2], "z")

		xMax := parseFloat(max[0], "x")
		yMax := parseFloat(max[1], "y")
		zMax := parseFloat(max[2], "z")

		AABB1 := Entity{
			AABB{
				Min: Corner{xMin, yMin, zMin},
				Max: Corner{xMax, yMax, zMax},
			},
		}
		gmap.Entities[key] = AABB1
	}
}
func parseFloat(str string, ctype string) float32 {
	var f float32
	switch ctype {
	case "x":
		_, err := fmt.Sscanf(str, "X=%f", &f)
		if err != nil {
			panic(err)
		}
		return f
	case "y":
		_, err := fmt.Sscanf(str, "Y=%f", &f)
		if err != nil {
			panic(err)
		}
		return f
	case "z":
		_, err := fmt.Sscanf(str, "Z=%f", &f)
		if err != nil {
			panic(err)
		}
		return f
	}
	return f
}
func (gmap *JsonData) NewRay(Origin, Direction Corner) Ray {
	return Ray{Origin: Origin, Direction: Direction}
}

func (gmap *JsonData) NewVector(X float32, Y float32, Z float32) Corner {
	return Corner{X: X, Y: Y, Z: Z}
}
func (gmap *JsonData) ActorRayIntersect(Name string, RayTrace Ray) bool {
	return gmap.RayAABBIntersect(RayTrace, gmap.Entities[Name].AABB)
}
func (gmap *JsonData) RayAABBIntersect(ray Ray, aabb AABB) bool {
	// 检查光线方向向量是否为零向量

	// 计算入射和出射时间
	var tmin, tmax float32

	if ray.Direction.X >= 0 {
		tmin = (aabb.Min.X - ray.Origin.X) / ray.Direction.X
		tmax = (aabb.Max.X - ray.Origin.X) / ray.Direction.X
	} else {
		tmin = (aabb.Max.X - ray.Origin.X) / ray.Direction.X
		tmax = (aabb.Min.X - ray.Origin.X) / ray.Direction.X
	}

	if ray.Direction.Y >= 0 {
		tymin := (aabb.Min.Y - ray.Origin.Y) / ray.Direction.Y
		tymax := (aabb.Max.Y - ray.Origin.Y) / ray.Direction.Y

		if tymin > tmin {
			tmin = tymin
		}

		if tymax < tmax {
			tmax = tymax
		}
	} else {
		tymin := (aabb.Max.Y - ray.Origin.Y) / ray.Direction.Y
		tymax := (aabb.Min.Y - ray.Origin.Y) / ray.Direction.Y

		if tymin > tmin {
			tmin = tymin
		}

		if tymax < tmax {
			tmax = tymax
		}
	}

	if ray.Direction.Z >= 0 {
		tzmin := (aabb.Min.Z - ray.Origin.Z) / ray.Direction.Z
		tzmax := (aabb.Max.Z - ray.Origin.Z) / ray.Direction.Z

		if tzmin > tmin {
			tmin = tzmin
		}

		if tzmax < tmax {
			tmax = tzmax
		}
	} else {
		tzmin := (aabb.Max.Z - ray.Origin.Z) / ray.Direction.Z
		tzmax := (aabb.Min.Z - ray.Origin.Z) / ray.Direction.Z

		if tzmin > tmin {
			tmin = tzmin
		}

		if tzmax < tmax {
			tmax = tzmax
		}
	}

	// 检查是否有相交
	if tmin <= tmax && tmax > 0 {
		return true
	} else {
		return false
	}
}

func init() {
	Maputils = &JsonData{
		Entities: make(map[string]Entity),
	}
	Maputils.LoadMap()
}
