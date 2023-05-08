package telemetry

import (
	"strconv"

	"github.com/Shopify/sarama"
	"github.com/go-logr/zerologr"
	"github.com/goTrooper/pkg/kafka"
	"golang.org/x/sys/windows"
)

// unsafe.Sizeof(windows.ProcessEntry32{})
const processEntrySize = 568

type TelemetryUtil struct {
	logger          zerologr.Logger
	kclient         *kafka.KafkaClient
	processSnapShot map[string]string
}

func NewTelemetryUtil(logr zerologr.Logger, client *kafka.KafkaClient) *TelemetryUtil {
	return &TelemetryUtil{
		logger:          logr,
		kclient:         client,
		processSnapShot: make(map[string]string),
	}
}

/*
* Single threaded telemetry collection
 */
func (tele *TelemetryUtil) StartCollectingTelemetry() {
	for {
		tele.enumProcessess()
	}
}

func (tele *TelemetryUtil) collectWindowsProcess() {
	//for {
	//enumProcessess()
	//}
}

/*
* Took from
* https://stackoverflow.com/questions/11356264/list-of-currently-running-process-in-golang-windows-version
 */
func (tele *TelemetryUtil) enumProcessess() {
	h, e := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	defer windows.CloseHandle(h)
	if e != nil {
		panic(e)
	}
	p := windows.ProcessEntry32{Size: processEntrySize}
	for {
		e := windows.Process32Next(h, &p)
		pid := strconv.FormatUint(uint64(p.ProcessID), 10)
		if e != nil {
			break
		}
		s := windows.UTF16ToString(p.ExeFile[:])
		if tele.processSnapShot[pid] == "" {
			tele.processSnapShot[pid] = s
			tele.submitTrooperEvent(s)
		}
	}
}

func (tele *TelemetryUtil) submitTrooperEvent(event string) {
	producer := tele.kclient.Producer
	if producer == nil {
		return
	}

	// Start Kafka transaction
	err := producer.BeginTxn()
	if err != nil {
		tele.logger.Error(err, "Kafa transtcion failied")
	}

	//eventsTopic := viper.GetString("kafka.producers[0].topic")

	producer.Input() <- &sarama.ProducerMessage{Topic: "trooper-events", Key: sarama.StringEncoder(event), Value: sarama.StringEncoder(event)}

	// commit transaction
	err = producer.CommitTxn()
	if err != nil {
		tele.logger.Error(err, "Producer: unable to commit tx")
		for {
			if producer.TxnStatus()&sarama.ProducerTxnFlagFatalError != 0 {
				// fatal error. need to recreate producer.
				tele.logger.Info("Producer: producer is in a fatal state, need to recreate it")
				break
			}
			// If producer is in abortable state, try to abort current transaction.
			if producer.TxnStatus()&sarama.ProducerTxnFlagAbortableError != 0 {
				err = producer.AbortTxn()
				if err != nil {
					// If an error occured just retry it.
					tele.logger.Error(err, "Producer: unable to abort transaction")
					continue
				}
				break
			}
			// if not you can retry
			err = producer.CommitTxn()
			if err != nil {
				tele.logger.Error(err, "Producer: unable to commit txn")
				continue
			}
		}
		return
	}
}
