package reports

import (
    "encoding/json"
    "github.com/eoaliev/golang-hackathon-november2021/utils"
)

func WriteReportToJsonFile(report interface{}, filename string) (error) {
    reportJson, err := json.Marshal(report)
    if err != nil {
        return err
    }

    err = utils.WriteBytesToJsonFile("/staticfiles/"+filename, reportJson)
    if err != nil {
        return err
    }

    return nil
}
