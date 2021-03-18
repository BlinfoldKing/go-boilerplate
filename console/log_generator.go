package console

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	"go-boilerplate/helper"
	"go-boilerplate/modules/sensor"
	sensorlog "go-boilerplate/modules/sensor_log"
	"math/rand"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var generator = &cobra.Command{
	Use:   "generate-sensor-log",
	Short: "generate sensor log",
	Long:  `This subcommand used to generate sensor log`,
	Run:   generateSensorLog,
}

func init() {
	Root.AddCommand(generator)
}

func generateSensorLog(cmd *cobra.Command, args []string) {
	adapters, err := adapters.Init()
	if err != nil {
		helper.Logger.Fatal(err)
	}

	sensorrepo := sensor.CreatePosgresRepository(adapters.Postgres)
	sensorservice := sensor.CreateService(sensorrepo)

	logrepo := sensorlog.CreatePosgresRepository(adapters.Postgres)
	logservice := sensorlog.CreateService(logrepo)

	sensors, _, err := sensorservice.GetList(entity.OffsetPagination{})
	if err != nil {
		helper.Logger.Fatal(err)
	}

	next := time.Now()
	for true {
		now := time.Now()
		if now.After(next) {
			for _, sensor := range sensors {
				switch sensor.SensorType {
				case entity.SensorTemperature:
					val := rand.Intn(10)
					valstr := strconv.Itoa(val)
					logservice.CreateSensorLog(sensor.ID, "Celcius Degree", "{}", valstr)
				case entity.SensorCam:
					val := rand.Intn(1)
					valstr := strconv.Itoa(val)
					logservice.CreateSensorLog(sensor.ID, "Motion Detected", `{ "url": "https://www.youtube.com/watch?v=d1UeEv0RkgU" }`, valstr)
				case entity.SensorDoor, entity.SensorGate, entity.SensorSwitch:
					val := rand.Intn(1)
					valstr := strconv.Itoa(val)
					logservice.CreateSensorLog(sensor.ID, "Open/Close", `{}`, valstr)
				case entity.SensorGyro:
					val := rand.Intn(10)
					valstr := strconv.Itoa(val)
					logservice.CreateSensorLog(sensor.ID, "Degree", `{}`, valstr)
				case entity.SensorFence:
					val := rand.Intn(1)
					valstr := strconv.Itoa(val)
					logservice.CreateSensorLog(sensor.ID, "Intrusion", `{}`, valstr)
				case entity.SensorHumidty:
					val := rand.Intn(10)
					valstr := strconv.Itoa(val)
					logservice.CreateSensorLog(sensor.ID, "Percent", `{}`, valstr)
				case entity.SensorModem, entity.SensorRouter:
					val := rand.Intn(1000)
					valstr := strconv.Itoa(val)
					logservice.CreateSensorLog(sensor.ID, "kb (Incoming)", `{}`, valstr)
					logservice.CreateSensorLog(sensor.ID, "kb (Outgoing)", `{}`, valstr)
				}
			}

			next = now.Add(30 * time.Minute)
		}
	}
}
